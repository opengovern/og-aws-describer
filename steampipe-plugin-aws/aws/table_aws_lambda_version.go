package aws

import (
	"context"

	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAwsLambdaVersion(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_lambda_version",
		Description: "AWS Lambda Version",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"version", "function_name"}),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"InvalidParameter", "ResourceNotFoundException"}),
			},
			Hydrate: kaytu.GetLambdaFunctionVersion,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListLambdaFunctionVersion,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "function_name", Require: plugin.Optional},
			},
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "version",
				Description: "The version of the Lambda function.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.FunctionVersion.Version")},
			{
				Name:        "function_name",
				Description: "The name of the function.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.FunctionVersion.FunctionName")},
			{
				Name:        "arn",
				Description: "The function's Amazon Resource Name (ARN).",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.FunctionVersion.FunctionArn"),
			},
			{
				Name:        "master_arn",
				Description: "For Lambda@Edge functions, the ARN of the master function.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.FunctionVersion.MasterArn")},
			{
				Name:        "state",
				Description: "The current state of the function.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.FunctionVersion.State")},
			{
				Name:        "code_sha_256",
				Description: "The SHA256 hash of the function's deployment package.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.FunctionVersion.CodeSha256")},
			{
				Name:        "code_size",
				Description: "The size of the function's deployment package, in bytes.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.FunctionVersion.CodeSize")},
			{
				Name:        "description",
				Description: "The function's description.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.FunctionVersion.Description")},
			{
				Name:        "handler",
				Description: "The function that Lambda calls to begin executing your function.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.FunctionVersion.Handler")},
			{
				Name:        "last_modified",
				Description: "The date and time that the function was last updated, in ISO-8601 format.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.FunctionVersion.LastModified")},
			{
				Name:        "last_update_status",
				Description: "The status of the last update that was performed on the function.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.FunctionVersion.LastUpdateStatus")},
			{
				Name:        "last_update_status_reason",
				Description: "The reason for the last update that was performed on the function.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.FunctionVersion.LastUpdateStatusReason")},
			{
				Name:        "last_update_status_reason_code",
				Description: "The reason code for the last update that was performed on the function.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.FunctionVersion.LastUpdateStatusReasonCode")},
			{
				Name:        "memory_size",
				Description: "The memory that's allocated to the function.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.FunctionVersion.MemorySize")},
			{
				Name:        "revision_id",
				Description: "The latest updated revision of the function or alias.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Policy.RevisionId")},
			{
				Name:        "runtime",
				Description: "The runtime environment for the Lambda function.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.FunctionVersion.Runtime")},
			{
				Name:        "timeout",
				Description: "The amount of time in seconds that Lambda allows a function to run before stopping it.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.FunctionVersion.Timeout")},
			{
				Name:        "vpc_id",
				Description: "The ID of the VPC.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.FunctionVersion.VpcConfig.VpcId")},
			{
				Name:        "environment_variables",
				Description: "The environment variables that are accessible from function code during execution.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.FunctionVersion.Environment.Variables")},
			{
				Name:        "policy",
				Description: "Contains the resource-based policy.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Policy")},
			{
				Name:        "policy_std",
				Description: "Contains the contents of the resource-based policy in a canonical form for easier searching.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Policy").Transform(unescape).Transform(policyToCanonical),
			},
			{
				Name:        "vpc_security_group_ids",
				Description: "A list of VPC security groups IDs attached to Lambda function.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.FunctionVersion.VpcConfig.SecurityGroupIds")},
			{
				Name:        "vpc_subnet_ids",
				Description: "A list of VPC subnet IDs attached to Lambda function.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.FunctionVersion.VpcConfig.SubnetIds")},

			// Standard columns for all tables
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.FunctionVersion")},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.FunctionVersion.FunctionArn").Transform(arnToAkas),
			},
		}),
	}
}
