package describer

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	awsmodel "github.com/kaytu-io/kaytu-aws-describer/aws/model"
	"github.com/kaytu-io/kaytu-aws-describer/pkg/steampipe"

	"github.com/go-errors/errors"
	"github.com/kaytu-io/kaytu-aws-describer/aws"
	"github.com/kaytu-io/kaytu-aws-describer/aws/describer"
	"github.com/kaytu-io/kaytu-util/pkg/describe"
	"github.com/kaytu-io/kaytu-util/pkg/source"
	"github.com/kaytu-io/kaytu-util/pkg/vault"
	"github.com/kaytu-io/kaytu-util/proto/src/golang"
	"go.uber.org/zap"
)

type KaytuError struct {
	ErrCode string

	error
}

func Do(ctx context.Context,
	vlt *vault.KMSVaultSourceConfig,
	logger *zap.Logger,
	job describe.DescribeJob,
	keyARN string,
	describeDeliverEndpoint string,
	describeDeliverToken string,
	ingestionPipelineEndpoint string,
	useOpenSearch bool,
	kafkaTopic string,
	workspaceId string,
	workspaceName string) (resourceIDs []string, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("paniced with error: %v", r)
			logger.Error("paniced with error", zap.Error(err), zap.String("stackTrace", errors.Wrap(r, 2).ErrorStack()))
		}
	}()

	if job.SourceType != source.CloudAWS {
		return nil, fmt.Errorf("unsupported source type %s", job.SourceType)
	}

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	config, err := vlt.Decrypt(job.CipherText, keyARN)
	if err != nil {
		return nil, fmt.Errorf("decrypt error: %w", err)
	}

	return doDescribeAWS(ctx, logger, job, config, workspaceId, workspaceName, describeDeliverEndpoint, ingestionPipelineEndpoint, describeDeliverToken, kafkaTopic, useOpenSearch)
}

func doDescribeAWS(ctx context.Context, logger *zap.Logger, job describe.DescribeJob, config map[string]any, workspaceId, workspaceName string, describeEndpoint, ingestionPipelineEndpoint string, describeToken string, kafkaTopic string, useOpenSearch bool) ([]string, error) {
	rs, err := NewResourceSender(workspaceId, workspaceName, describeEndpoint, ingestionPipelineEndpoint, describeToken, job.JobID, kafkaTopic, useOpenSearch, logger)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to resource sender: %w", err)
	}

	plg := steampipe.Plugin()
	creds, err := aws.AccountConfigFromMap(config)
	if err != nil {
		return nil, fmt.Errorf("aws account credentials: %w", err)
	}

	f := func(resource describer.Resource) error {
		if resource.Description == nil {
			return nil
		}
		descriptionJSON, err := json.Marshal(resource.Description)
		if err != nil {
			return err
		}
		partition, _ := aws.PartitionOf(resource.Region)
		if partition == "" {
			partition = "aws"
		}
		resource.Account = job.AccountID
		resource.Type = strings.ToLower(job.ResourceType)
		resource.Partition = partition
		awsMetadata := awsmodel.Metadata{
			Name:         resource.Name,
			AccountID:    job.AccountID,
			SourceID:     job.SourceID,
			Region:       resource.Region,
			Partition:    partition,
			ResourceType: strings.ToLower(job.ResourceType),
		}

		awsMetadataBytes, err := json.Marshal(awsMetadata)
		if err != nil {
			return fmt.Errorf("marshal metadata: %v", err.Error())
		}

		metadata := make(map[string]string)
		err = json.Unmarshal(awsMetadataBytes, &metadata)
		if err != nil {
			return fmt.Errorf("unmarshal metadata: %v", err.Error())
		}

		kafkaResource := Resource{
			ID:            resource.UniqueID(),
			ARN:           resource.ARN,
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

		tags, name, err := steampipe.ExtractTagsAndNames(plg, job.ResourceType, kafkaResource)
		if err != nil {
			return fmt.Errorf("failed to build tags for service: %v", err.Error())
		}
		if len(name) > 0 {
			kafkaResource.Metadata["name"] = name
		}

		rs.Send(&golang.AWSResource{
			UniqueId:        resource.UniqueID(),
			Arn:             resource.ARN,
			Id:              resource.ID,
			Name:            resource.Name,
			Account:         job.AccountID,
			Region:          resource.Region,
			Partition:       partition,
			Type:            job.ResourceType,
			DescriptionJson: string(descriptionJSON),
			Metadata:        metadata,
			Tags:            tags,
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
		return nil
	}
	clientStream := (*describer.StreamSender)(&f)

	output, err := aws.GetResources(
		ctx, logger,
		job.ResourceType, job.TriggerType,
		job.AccountID,
		creds.Regions, creds.AccountID, creds.AccessKey, creds.SecretKey, creds.SessionToken, creds.AssumeRoleName, creds.AssumeAdminRoleName, creds.ExternalID,
		false, clientStream)
	if err != nil {
		return nil, fmt.Errorf("AWS: %w", err)
	}

	rs.Finish()

	var errs []string
	for region, err := range output.Errors {
		if err != "" {
			errs = append(errs, fmt.Sprintf("region (%s): %s", region, err))
		}
	}

	// For AWS resources, since they are queries independently per region,
	// if there is an error in some regions, return those errors. For the regions
	// with no error, return the list of resources.
	if len(errs) > 0 {
		err = fmt.Errorf("AWS: [%s]", strings.Join(errs, ","))
	} else {
		err = nil
	}

	var kerr error
	if err != nil {
		kerr = KaytuError{
			ErrCode: output.ErrorCode,
			error:   err,
		}
	}

	return rs.GetResourceIDs(), kerr
}
