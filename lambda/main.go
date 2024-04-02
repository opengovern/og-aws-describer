package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/kaytu-io/kaytu-aws-describer/describer"
	"github.com/kaytu-io/kaytu-util/pkg/describe"
	"go.uber.org/zap"
	"os"
	"strings"
)

func lambdaHandler(ctx context.Context, input describe.DescribeWorkerInput) error {
	fmt.Printf("Input: %s", zap.Any("input", input).String)
	logger := zap.NewNop()
	if val, ok := os.LookupEnv("DEBUG"); ok && strings.ToLower(val) == "true" {
		logger, _ = zap.NewProduction()
	}
	return describer.DescribeHandler(ctx, logger, describer.TriggeredByAWSLambda, input)
}

func main() {
	lambda.Start(lambdaHandler)
}
