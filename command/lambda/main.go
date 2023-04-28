package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	kaytu_aws_describer "github.com/kaytu-io/kaytu-aws-describer"
	"gitlab.com/keibiengine/keibi-engine/pkg/describe"
	"gitlab.com/keibiengine/keibi-engine/pkg/vault"
	"go.uber.org/zap"
)

func DescribeHandler(ctx context.Context, req events.APIGatewayProxyRequest) error {
	logger, err := zap.NewProduction()
	if err != nil {
		return err
	}

	logger.Info(req.Body)

	var input describe.LambdaDescribeWorkerInput
	err = json.Unmarshal([]byte(req.Body), &input)
	if err != nil {
		logger.Error("Failed to unmarshal input", zap.Error(err))
		return err
	}

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
