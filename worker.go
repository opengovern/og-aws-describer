package kaytu_aws_describer

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/go-errors/errors"
	"github.com/kaytu-io/kaytu-aws-describer/aws"
	"github.com/kaytu-io/kaytu-aws-describer/aws/describer"
	"github.com/kaytu-io/kaytu-aws-describer/pkg/describe"
	"github.com/kaytu-io/kaytu-aws-describer/pkg/source"
	"github.com/kaytu-io/kaytu-aws-describer/pkg/vault"
	"github.com/kaytu-io/kaytu-aws-describer/proto/src/golang"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const (
	DescribeResourceJobFailed    string = "FAILED"
	DescribeResourceJobSucceeded string = "SUCCEEDED"
)

func doDescribeAWS(ctx context.Context, job describe.DescribeJob, config map[string]any, client *golang.DescribeServiceClient) ([]string, error) {
	var resourceIDs []string
	creds, err := aws.AccountConfigFromMap(config)
	if err != nil {
		return nil, fmt.Errorf("aws account credentials: %w", err)
	}

	var clientStream *describer.StreamSender
	if client != nil {
		grpcCtx := context.Background()
		grpcCtx = context.WithValue(grpcCtx, "resourceJobID", job.JobID)
		stream, err := (*client).DeliverAWSResources(grpcCtx)
		if err != nil {
			return nil, err
		}

		f := func(resource describer.Resource) error {
			descriptionJSON, err := json.Marshal(resource.Description)
			if err != nil {
				return err
			}

			resourceIDs = append(resourceIDs, resource.UniqueID())
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
		defer stream.CloseAndRecv()
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
		return nil, fmt.Errorf("AWS: %w", err)
	}

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

	return resourceIDs, err
}

func Do(ctx context.Context,
	vlt *vault.KMSVaultSourceConfig,
	job describe.DescribeJob,
	keyARN string,
	describeDeliverEndpoint *string) error {
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
		status               = DescribeResourceJobSucceeded
		firstErr    error    = nil
		resourceIDs []string = nil
	)

	fail := func(err error) {
		status = DescribeResourceJobFailed
		if firstErr == nil {
			firstErr = err
		}
	}

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	if conn, err := grpc.Dial(*describeDeliverEndpoint, grpc.WithTransportCredentials(credentials.NewTLS(nil))); err == nil {
		defer conn.Close()
		client := golang.NewDescribeServiceClient(conn)

		if config, err := vlt.Decrypt(job.CipherText, keyARN); err == nil {
			resourceIDs, err = doDescribeAWS(ctx, job, config, &client)
			if err != nil {
				// Don't return here. In certain cases, such as AWS, resources might be
				// available for some regions while there was failures in other regions.
				// Instead, continue to write whatever you can to kafka.
				fail(fmt.Errorf("describe resources: %w", err))
			}
		} else if err != nil {
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
