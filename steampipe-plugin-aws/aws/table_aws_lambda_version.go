package aws

import (
	"context"

	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"

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
			Hydrate: opengovernance.GetLambdaFunctionVersion,
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListLambdaFunctionVersion,
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
				Name:        "kms_key_arn",
				Description: "The KMS key that's used to encrypt the function's environment variables.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.FunctionVersion.KMSKeyArn"),
			},
			{
				Name:        "role",
				Description: "The function's execution role.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.FunctionVersion.Role")},
			{
				Name:        "signing_job_arn",
				Description: "The ARN of the signing job.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.FunctionVersion.SigningJobArn")},
			{
				Name:        "signing_profile_version_arn",
				Description: "The ARN of the signing profile version.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.FunctionVersion.SigningProfileVersionArn")},
			{
				Name:        "state_reason",
				Description: "The reason for the function's current state.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.FunctionVersion.StateReason")},
			{
				Name:        "state_reason_code",
				Description: "The reason code for the function's current state.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.FunctionVersion.StateReasonCode")},
			{
				Name:        "ephemeral_storage_size",
				Description: "The size of the function's /tmp directory in MB.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.FunctionVersion.EphemeralStorage.Size")},

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

			{
				Name:        "architectures",
				Description: "The instruction set architecture that the function supports.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.FunctionVersion.Architectures")},
			{
				Name:        "dead_letter_config",
				Description: "The function's dead letter queue configuration.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.FunctionVersion.DeadLetterConfig")},
			{
				Name:        "file_system_configs",
				Description: "Connection settings for an Amazon EFS file system.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.FunctionVersion.FileSystemConfigs")},
			{
				Name:        "image_config_response",
				Description: "The function's image configuration values.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.FunctionVersion.ImageConfigResponse")},
			{
				Name:        "layers",
				Description: "The function's layers.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.FunctionVersion.Layers")},
			{
				Name:        "logging_config",
				Description: "The function's Amazon CloudWatch Logs configuration settings.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.FunctionVersion.LoggingConfig")},
			{
				Name:        "runtime_version_config",
				Description: "The ARN of the runtime and any errors that occurred.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.FunctionVersion.RuntimeVersionConfig")},
			{
				Name:        "snap_start",
				Description: "Configuration for creating a snapshot of the initialized execution environment.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.FunctionVersion.SnapStart")},
			{
				Name:        "tracing_config",
				Description: "The function's X-Ray tracing configuration.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.FunctionVersion.TracingConfig")},

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
