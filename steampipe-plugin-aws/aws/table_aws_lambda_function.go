package aws

import (
	"context"
	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAwsLambdaFunction(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_lambda_function",
		Description: "AWS Lambda Function",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("name"),
			Hydrate:    kaytu.GetLambdaFunction,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListLambdaFunction,
		},

		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the function.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Function.Configuration.FunctionName")},
			{
				Name:        "arn",
				Description: "The function's Amazon Resource Name (ARN).",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Function.Configuration.FunctionArn"),
			},
			{
				Name:        "code_sha_256",
				Description: "The SHA256 hash of the function's deployment package.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Function.Configuration.CodeSha256")},
			{
				Name:        "code_size",
				Description: "The size of the function's deployment package, in bytes.",
				Type:        proto.ColumnType_INT,

				Transform: transform.FromField("Description.Function.Configuration.CodeSize")},
			{
				Name:        "dead_letter_config_target_arn",
				Description: "The Amazon Resource Name (ARN) of an Amazon SQS queue or Amazon SNS topic.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Function.Configuration.DeadLetterConfig.TargetArn")},
			{
				Name:        "description",
				Description: "The function's description.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Function.Configuration.Description")},
			{
				Name:        "handler",
				Description: "The function that Lambda calls to begin executing your function.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Function.Configuration.Handler")},
			{
				Name:        "kms_key_arn",
				Description: "The KMS key that's used to encrypt the function's environment variables. This key is only returned if you've configured a customer managed CMK.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Function.Configuration.KMSKeyArn")},
			{
				Name:        "last_modified",
				Description: "The date and time that the function was last updated.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Function.Configuration.LastModified")},
			{
				Name:        "timeout",
				Description: "The amount of time in seconds that Lambda allows a function to run before stopping it.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Function.Configuration.Timeout")},
			{
				Name:        "version",
				Description: "The version of the Lambda function.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Function.Configuration.Version")},
			{
				Name:        "package_type",
				Description: "The type of deployment package.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Function.Configuration.PackageType"),
			},
			{
				Name:        "master_arn",
				Description: "For Lambda@Edge functions, the ARN of the master function.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Function.Configuration.MasterArn")},
			{
				Name:        "memory_size",
				Description: "The memory that's allocated to the function.",
				Type:        proto.ColumnType_INT,

				Transform: transform.FromField("Description.Function.Configuration.MemorySize")},
			{
				Name:        "revision_id",
				Description: "The latest updated revision of the function or alias.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Policy.RevisionId")},
			{
				Name:        "role",
				Description: "The function's execution role.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Function.Configuration.Role")},
			{
				Name:        "runtime",
				Description: "The runtime environment for the Lambda function.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Function.Configuration.Runtime")},
			{
				Name:        "state",
				Description: "The current state of the function.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Function.Configuration.State")},
			{
				Name:        "state_reason",
				Description: "The reason for the function's current state.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Function.Configuration.StateReason")},
			{
				Name:        "state_reason_code",
				Description: "The reason code for the function's current state.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Function.Configuration.StateReasonCode")},
			{
				Name:        "last_update_status",
				Description: "The status of the last update that was performed on the function.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Function.Configuration.LastUpdateStatus")},
			{
				Name:        "last_update_status_reason",
				Description: "The reason for the last update that was performed on the function.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Function.Configuration.LastUpdateStatusReason")},
			{
				Name:        "last_update_status_reason_code",
				Description: "The reason code for the last update that was performed on the function.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Function.Configuration.LastUpdateStatusReasonCode")},
			{
				Name:        "reserved_concurrent_executions",
				Description: "The number of concurrent executions that are reserved for this function.",
				Type:        proto.ColumnType_INT,

				Transform: transform.FromField("Description.Function.Concurrency.ReservedConcurrentExecutions")},
			{
				Name:        "vpc_id",
				Description: "The VPC ID that is attached to Lambda function.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Function.Configuration.VpcConfig.VpcId")},
			{
				Name:        "architectures",
				Description: "The instruction set architecture that the function supports. Architecture is a string array with one of the valid values.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Function.Configuration.Architectures")},
			{
				Name:        "code",
				Description: "The deployment package of the function or version.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Function.Code")},
			{
				Name:        "environment_variables",
				Description: "The environment variables that are accessible from function code during execution.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Function.Configuration.Environment.Variables")},
			{
				Name:        "file_system_configs",
				Description: "Connection settings for an Amazon EFS file system.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Function.Configuration.FileSystemConfigs")},
			{
				Name:        "policy",
				Description: "The resource-based iam policy of Lambda function.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Policy.Policy"),
			},
			{
				Name:        "policy_std",
				Description: "Contains the policy in a canonical form for easier searching.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Policy.Policy").Transform(policyToCanonical),
			},
			{
				Name:        "tracing_config",
				Description: "The function's X-Ray tracing configuration.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Function.Configuration.TracingConfig")},
			{
				Name:        "snap_start",
				Description: "Set ApplyOn to PublishedVersions to create a snapshot of the initialized execution environment when you publish a function version.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Function.Configuration.SnapStart")},
			{
				Name:        "url_config",
				Description: "The function URL configuration details of the function.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.UrlConfig")},
			{
				Name:        "vpc_security_group_ids",
				Description: "A list of VPC security groups IDs attached to Lambda function.",
				Type:        proto.ColumnType_JSON,

				Transform: transform.FromField("Description.Function.Configuration.VpcConfig.SecurityGroupIds")},
			{
				Name:        "vpc_subnet_ids",
				Description: "A list of VPC subnet IDs attached to Lambda function.",
				Type:        proto.ColumnType_JSON,

				Transform: transform.FromField("Description.Function.Configuration.VpcConfig.SubnetIds")},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Function.Configuration.FunctionName")},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,

				Transform: transform.FromField("Description.Function.Tags")},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Function.Configuration.FunctionArn").Transform(arnToAkas),
			},
		}),
	}
}

//// LIST FUNCTION

func listAwsLambdaFunctions(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {

	// Create service
	svc, err := LambdaClient(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("aws_lambda_function.listAwsLambdaFunctions", "connection_error", err)
		return nil, err
	}
	if svc == nil {
		// Unsupported region, return no data
		return nil, nil
	}

	maxItems := int32(10000)
	input := lambda.ListFunctionsInput{}

	// Reduce the basic request limit down if the user has only requested a small number of rows
	if d.QueryContext.Limit != nil {
		limit := int32(*d.QueryContext.Limit)
		if limit < maxItems {
			if limit < 1 {
				maxItems = int32(1)
			} else {
				maxItems = int32(limit)
			}
		}
	}

	input.MaxItems = aws.Int32(maxItems)
	paginator := lambda.NewListFunctionsPaginator(svc, &input, func(o *lambda.ListFunctionsPaginatorOptions) {
		o.Limit = maxItems
		o.StopOnDuplicateToken = true
	})

	for paginator.HasMorePages() {
		output, err := paginator.NextPage(ctx)
		if err != nil {
			plugin.Logger(ctx).Error("aws_lambda_function.listAwsLambdaFunctions", "api_error", err)
			return nil, err
		}

		for _, function := range output.Functions {
			d.StreamListItem(ctx, function)

			// Context may get cancelled due to manual cancellation or if the limit has been reached
			if d.RowsRemaining(ctx) == 0 {
				return nil, nil
			}
		}
	}

	return nil, nil
}

//// HYDRATE FUNCTIONS
