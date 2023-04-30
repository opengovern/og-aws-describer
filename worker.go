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
	"go.uber.org/zap"
)

func Do(ctx context.Context,
	vlt *vault.KMSVaultSourceConfig,
	logger *zap.Logger,
	job describe.DescribeJob,
	keyARN string,
	describeDeliverEndpoint string,
	describeDeliverToken string) (resourceIDs []string, err error) {
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

	return doDescribeAWS(ctx, logger, job, config, describeDeliverEndpoint, describeDeliverToken)
}

func doDescribeAWS(ctx context.Context, logger *zap.Logger, job describe.DescribeJob, config map[string]any, describeEndpoint string, describeToken string) ([]string, error) {
	rs, err := NewResourceSender(describeEndpoint, describeToken, job.JobID, logger)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to resource sender: %w", err)
	}

	creds, err := aws.AccountConfigFromMap(config)
	if err != nil {
		return nil, fmt.Errorf("aws account credentials: %w", err)
	}

	f := func(resource describer.Resource) error {
		descriptionJSON, err := json.Marshal(resource.Description)
		if err != nil {
			return err
		}

		rs.Send(&golang.AWSResource{
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
		return nil
	}
	clientStream := (*describer.StreamSender)(&f)

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

	return rs.GetResourceIDs(), err
}
