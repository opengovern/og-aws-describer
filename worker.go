package kaytu_aws_describer

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/go-errors/errors"
	"github.com/kaytu-io/kaytu-aws-describer/aws"
	"github.com/kaytu-io/kaytu-aws-describer/aws/describer"
	"github.com/kaytu-io/kaytu-aws-describer/aws/model"
	"gitlab.com/keibiengine/keibi-engine/pkg/describe"
	"gitlab.com/keibiengine/keibi-engine/pkg/describe/api"
	"gitlab.com/keibiengine/keibi-engine/pkg/describe/es"
	"gitlab.com/keibiengine/keibi-engine/pkg/describe/proto/src/golang"
	"gitlab.com/keibiengine/keibi-engine/pkg/kafka"
	"gitlab.com/keibiengine/keibi-engine/pkg/source"
	"gitlab.com/keibiengine/keibi-engine/pkg/steampipe"
	"gitlab.com/keibiengine/keibi-engine/pkg/vault"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func doDescribeAWS(ctx context.Context, job describe.DescribeJob, config map[string]any,
	logger *zap.Logger, client *golang.DescribeServiceClient) ([]kafka.Doc, []string, error) {

	var resourceIDs []string
	creds, err := aws.AccountConfigFromMap(config)
	if err != nil {
		return nil, nil, fmt.Errorf("aws account credentials: %w", err)
	}

	var clientStream *describer.StreamSender
	if client != nil {
		stream, err := (*client).DeliverAWSResources(context.Background())
		if err != nil {
			return nil, nil, err
		}

		f := func(resource describer.Resource) error {
			descriptionJSON, err := json.Marshal(resource.Description)
			if err != nil {
				return err
			}

			return stream.Send(&golang.AWSResource{
				Arn:             resource.ARN,
				Id:              resource.ID,
				Name:            resource.Name,
				Account:         resource.Account,
				Region:          resource.Region,
				Partition:       resource.Partition,
				Type:            resource.Type,
				DescriptionJson: string(descriptionJSON),
				Job: &golang.DescribeJob{
					JobId:         uint32(job.JobID),
					ScheduleJobId: uint32(job.ScheduleJobID),
					ParentJobId:   uint32(job.ParentJobID),
					ResourceType:  job.ResourceType,
					SourceId:      job.SourceID,
					AccountId:     job.AccountID,
					DescribedAt:   job.DescribedAt,
					SourceType:    string(job.SourceType),
					ConfigReg:     job.CipherText,
					TriggerType:   string(job.TriggerType),
					RetryCounter:  uint32(job.RetryCounter),
				},
			})
		}
		clientStream = (*describer.StreamSender)(&f)
	}

	output, err := aws.GetResources(
		ctx,
		job.ResourceType,
		job.TriggerType,
		creds.AccountID,
		creds.Regions,
		creds.AccessKey,
		creds.SecretKey,
		creds.SessionToken,
		creds.AssumeRoleARN,
		false,
		clientStream,
	)
	if err != nil {
		return nil, nil, fmt.Errorf("AWS: %w", err)
	}

	var errs []string
	for region, err := range output.Errors {
		if err != "" {
			errs = append(errs, fmt.Sprintf("region (%s): %s", region, err))
		}
	}

	var msgs []kafka.Doc

	for _, resources := range output.Resources {
		for _, resource := range resources {
			if resource.Description == nil {
				continue
			}

			awsMetadata := model.Metadata{
				Name:         resource.Name,
				AccountID:    resource.Account,
				SourceID:     job.SourceID,
				Region:       resource.Region,
				Partition:    resource.Name,
				ResourceType: strings.ToLower(resource.Type),
			}
			awsMetadataBytes, err := json.Marshal(awsMetadata)
			if err != nil {
				errs = append(errs, fmt.Sprintf("marshal metadata: %v", err.Error()))
				continue
			}
			metadata := make(map[string]string)
			err = json.Unmarshal(awsMetadataBytes, &metadata)
			if err != nil {
				errs = append(errs, fmt.Sprintf("unmarshal metadata: %v", err.Error()))
				continue
			}

			kafkaResource := es.Resource{
				ID:            resource.UniqueID(),
				Name:          resource.Name,
				SourceType:    source.CloudAWS,
				ResourceType:  strings.ToLower(job.ResourceType),
				ResourceGroup: "",
				Location:      resource.Region,
				SourceID:      job.SourceID,
				ResourceJobID: job.JobID,
				SourceJobID:   job.ParentJobID,
				ScheduleJobID: job.ScheduleJobID,
				CreatedAt:     job.DescribedAt,
				Description:   resource.Description,
				Metadata:      metadata,
			}
			lookupResource := es.LookupResource{
				ResourceID:    resource.UniqueID(),
				Name:          resource.Name,
				SourceType:    source.CloudAWS,
				ResourceType:  strings.ToLower(job.ResourceType),
				ResourceGroup: "",
				Location:      resource.Region,
				SourceID:      job.SourceID,
				ResourceJobID: job.JobID,
				SourceJobID:   job.ParentJobID,
				ScheduleJobID: job.ScheduleJobID,
				CreatedAt:     job.DescribedAt,
			}
			resourceIDs = append(resourceIDs, resource.UniqueID())
			pluginTableName := steampipe.ExtractTableName(job.ResourceType)
			desc, err := steampipe.ConvertToDescription(job.ResourceType, kafkaResource)
			if err != nil {
				errs = append(errs, fmt.Sprintf("convertToDescription: %v", err.Error()))
				continue
			}
			cells, err := steampipe.AWSDescriptionToRecord(desc, pluginTableName)
			if err != nil {
				errs = append(errs, fmt.Sprintf("awsdescriptionToRecord: %v", err.Error()))
				continue
			}
			for name, v := range cells {
				if name == "title" || name == "name" {
					kafkaResource.Metadata["name"] = v.GetStringValue()
				}
			}

			tags, err := steampipe.ExtractTags(job.ResourceType, kafkaResource)
			if err != nil {
				errs = append(errs, fmt.Sprintf("failed to build tags for service: %v", err.Error()))
				tags = map[string]string{}
			}
			lookupResource.Tags = tags

			msgs = append(msgs, kafkaResource)
			msgs = append(msgs, lookupResource)
		}
	}

	logger.Info(fmt.Sprintf("job[%d] parent[%d] resourceType[%s]\n",
		job.JobID, job.ParentJobID, job.ResourceType))

	// For AWS resources, since they are queries independently per region,
	// if there is an error in some regions, return those errors. For the regions
	// with no error, return the list of resources.
	if len(errs) > 0 {
		err = fmt.Errorf("AWS: [%s]", strings.Join(errs, ","))
	} else {
		err = nil
	}

	return msgs, resourceIDs, err
}

func Do(ctx context.Context,
	vlt *vault.KMSVaultSourceConfig,
	job describe.DescribeJob,
	keyARN string,
	logger *zap.Logger,
	describeDeliverEndpoint *string) error {
	logger.Info("Starting DescribeJob",
		zap.Uint("jobID", job.JobID),
		zap.Uint("parentJobID", job.ParentJobID),
		zap.String("resourceType", job.ResourceType),
		zap.String("sourceID", job.SourceID),
		zap.String("accountID", job.AccountID),
		zap.Int64("describedAt", job.DescribedAt),
		zap.String("sourceType", string(job.SourceType)),
		zap.String("cipherText", job.CipherText),
		zap.String("triggerType", string(job.TriggerType)),
		zap.Uint("retryCounter", job.RetryCounter))

	if job.SourceType != source.CloudAWS {
		return fmt.Errorf("unsupported source type %s", job.SourceType)
	}

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("paniced with error:", err)
			fmt.Println(errors.Wrap(err, 2).ErrorStack())
		}
	}()

	// Assume it succeeded unless it fails somewhere
	var (
		status               = api.DescribeResourceJobSucceeded
		firstErr    error    = nil
		resourceIDs []string = nil
	)

	fail := func(err error) {
		status = api.DescribeResourceJobFailed
		if firstErr == nil {
			firstErr = err
		}
	}

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	if conn, err := grpc.Dial(*describeDeliverEndpoint); err == nil {
		defer conn.Close()
		client := golang.NewDescribeServiceClient(conn)

		if config, err := vlt.Decrypt(job.CipherText, keyARN); err == nil {
			_, resourceIDs, err = doDescribeAWS(ctx, job, config, logger, &client)
			if err != nil {
				// Don't return here. In certain cases, such as AWS, resources might be
				// available for some regions while there was failures in other regions.
				// Instead, continue to write whatever you can to kafka.
				fail(fmt.Errorf("describe resources: %w", err))
			}
		} else if config == nil {
			fail(fmt.Errorf("config is null! path is: %s", job.CipherText))
		} else {
			fail(fmt.Errorf("resource source config: %w", err))
		}

		errMsg := ""
		if firstErr != nil {
			errMsg = firstErr.Error()
		}

		_, err := client.DeliverResult(ctx, &golang.DeliverResultRequest{
			JobId:       uint32(job.JobID),
			ParentJobId: uint32(job.ParentJobID),
			Status:      string(status),
			Error:       errMsg,
			DescribeJob: &golang.DescribeJob{
				JobId:         uint32(job.JobID),
				ScheduleJobId: uint32(job.ScheduleJobID),
				ParentJobId:   uint32(job.ParentJobID),
				ResourceType:  job.ResourceType,
				SourceId:      job.SourceID,
				AccountId:     job.AccountID,
				DescribedAt:   job.DescribedAt,
				SourceType:    string(job.SourceType),
				ConfigReg:     job.CipherText,
				TriggerType:   string(job.TriggerType),
				RetryCounter:  uint32(job.RetryCounter),
			},
			DescribedResourceIds: resourceIDs,
		})
		if err != nil {
			return fmt.Errorf("DeliverResult: %v", err)
		}
		return nil
	} else {
		return fmt.Errorf("grpc: %v", err)
	}
}
