package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
	kaytu_aws_describer "github.com/kaytu-io/kaytu-aws-describer"
	"github.com/kaytu-io/kaytu-aws-describer/pkg/describe"
	"github.com/kaytu-io/kaytu-aws-describer/pkg/vault"
	"go.uber.org/zap"
)

func DescribeHandler(ctx context.Context, input describe.LambdaDescribeWorkerInput) error {
	logger, err := zap.NewProduction()
	if err != nil {
		return err
	}

	logger.Info(fmt.Sprintf("%v", input))

	kmsVault, err := vault.NewKMSVaultSourceConfig(ctx)
	if err != nil {
		return fmt.Errorf("failed to initialize KMS vault: %w", err)
	}

	return kaytu_aws_describer.Do(
		ctx,
		kmsVault,
		input.DescribeJob,
		input.KeyARN,
		logger,
		&input.DescribeEndpoint,
	)
}

func main() {
	lambda.Start(DescribeHandler)
}
