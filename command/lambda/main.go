package main

import (
	"context"
	"fmt"
	"github.com/kaytu-io/kaytu-aws-describer/proto/src/golang"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	kaytu_aws_describer "github.com/kaytu-io/kaytu-aws-describer"
	"github.com/kaytu-io/kaytu-aws-describer/pkg/describe"
	"github.com/kaytu-io/kaytu-aws-describer/pkg/vault"
	"go.uber.org/zap"
)

const (
	DescribeResourceJobFailed    string = "FAILED"
	DescribeResourceJobSucceeded string = "SUCCEEDED"
)

func DescribeHandler(ctx context.Context, input describe.LambdaDescribeWorkerInput) error {
	logger, err := zap.NewProduction()
	if err != nil {
		return err
	}
	logger.Info(fmt.Sprintf("%v", input))

	kmsVault, err := vault.NewKMSVaultSourceConfig(ctx, "", "", input.KeyRegion)
	if err != nil {
		return fmt.Errorf("failed to initialize KMS vault: %w", err)
	}

	resourceIds, err := kaytu_aws_describer.Do(
		ctx,
		kmsVault,
		logger,
		input.DescribeJob,
		input.KeyARN,
		input.DescribeEndpoint,
	)

	errMsg := ""
	status := DescribeResourceJobSucceeded
	if err != nil {
		errMsg = err.Error()
		status = DescribeResourceJobFailed
	}

	for retry := 0; retry < 5; retry++ {
		conn, err := grpc.Dial(input.DescribeEndpoint, grpc.WithTransportCredentials(credentials.NewTLS(nil)))
		if err != nil {
			logger.Error("[result delivery] connection failure:", zap.Error(err))
			time.Sleep(1 * time.Second)
			continue
		}
		client := golang.NewDescribeServiceClient(conn)

		_, err = client.DeliverResult(ctx, &golang.DeliverResultRequest{
			JobId:       uint32(input.DescribeJob.JobID),
			ParentJobId: uint32(input.DescribeJob.ParentJobID),
			Status:      status,
			Error:       errMsg,
			DescribeJob: &golang.DescribeJob{
				JobId:         uint32(input.DescribeJob.JobID),
				ScheduleJobId: uint32(input.DescribeJob.ScheduleJobID),
				ParentJobId:   uint32(input.DescribeJob.ParentJobID),
				ResourceType:  input.DescribeJob.ResourceType,
				SourceId:      input.DescribeJob.SourceID,
				AccountId:     input.DescribeJob.AccountID,
				DescribedAt:   input.DescribeJob.DescribedAt,
				SourceType:    string(input.DescribeJob.SourceType),
				ConfigReg:     input.DescribeJob.CipherText,
				TriggerType:   string(input.DescribeJob.TriggerType),
				RetryCounter:  uint32(input.DescribeJob.RetryCounter),
			},
			DescribedResourceIds: resourceIds,
		})
		if err != nil {
			logger.Error("[result delivery] rpc failed:", zap.Error(err))
			time.Sleep(1 * time.Second)
			continue
		}
		break
	}

	return err
}

func main() {
	lambda.Start(DescribeHandler)
}
