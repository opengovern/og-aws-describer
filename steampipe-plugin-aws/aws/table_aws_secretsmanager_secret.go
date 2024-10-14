package aws

import (
	"context"
	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"

	"github.com/aws/aws-sdk-go-v2/service/secretsmanager/types"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsSecretsManagerSecret(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_secretsmanager_secret",
		Description: "AWS Secrets Manager Secret",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("arn"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"ValidationException", "InvalidParameter", "ResourceNotFoundException"}),
			},
			Hydrate: opengovernance.GetSecretsManagerSecret,
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListSecretsManagerSecret,
		},

		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The friendly name of the secret.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Secret.Name")},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the secret.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Secret.ARN"),
			},
			{
				Name:        "created_date",
				Description: "The date and time when a secret was created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Secret.CreatedDate")},
			{
				Name:        "description",
				Description: "The user-provided description of the secret.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Secret.Description")},
			{
				Name:        "kms_key_id",
				Description: "The ARN or alias of the AWS KMS customer master key (CMK) used to encrypt the SecretString and SecretBinary fields in each version of the secret.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Secret.KmsKeyId")},
			{
				Name:        "deleted_date",
				Description: "The date and time the deletion of the secret occurred.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Secret.DeletedDate")},
			{
				Name:        "last_accessed_date",
				Description: "The last date that this secret was accessed.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Secret.LastAccessedDate")},
			{
				Name:        "last_changed_date",
				Description: "The last date and time that this secret was modified in any way.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Secret.LastChangedDate")},
			{
				Name:        "last_rotated_date",
				Description: "The most recent date and time that the Secrets Manager rotation process was successfully completed.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Secret.LastRotatedDate")},
			{
				Name:        "owning_service",
				Description: "Returns the name of the service that created the secret.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Secret.OwningService")},
			{
				Name:        "primary_region",
				Description: "The Region where Secrets Manager originated the secret.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Secret.PrimaryRegion")},
			{
				Name:        "policy",
				Description: "A JSON-formatted string that describes the permissions that are associated with the attached secret.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ResourcePolicy")},
			{
				Name:        "policy_std",
				Description: "Contains the permissions that are associated with the attached secret in a canonical form for easier searching.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ResourcePolicy").Transform(unescape).Transform(policyToCanonical),
			},
			{
				Name:        "replication_status",
				Description: "Describes a list of replication status objects as InProgress, Failed or InSync.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Secret.ReplicationStatus")},
			{
				Name:        "rotation_enabled",
				Description: "Indicates whether automatic, scheduled rotation is enabled for this secret.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Secret.RotationEnabled")},
			{
				Name:        "rotation_lambda_arn",
				Description: "The ARN of an AWS Lambda function invoked by Secrets Manager to rotate and expire the secret either automatically per the schedule or manually by a call to RotateSecret.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Secret.RotationLambdaARN"),
			},
			{
				Name:        "rotation_rules",
				Description: "A structure that defines the rotation configuration for the secret.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Secret.RotationRules")},
			{
				Name:        "secret_versions_to_stages",
				Description: "A list of all of the currently assigned SecretVersionStage staging labels and the SecretVersionId attached to each one.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Secret.VersionIdsToStages")},
			{
				Name:        "tags_src",
				Description: "The list of user-defined tags associated with the secret.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Secret.Tags")},
			// Standard columns for all tables

			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Secret.Name")},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Secret.Tags").Transform(secretsManagerSecretTagListToTurbotTags),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Secret.ARN").Transform(arnToAkas),
			},
		}),
	}
}

//// LIST FUNCTION

//// HYDRATE FUNCTIONS

//// TRANSFORM FUNCTION

func secretsManagerSecretTagListToTurbotTags(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	plugin.Logger(ctx).Trace("secretsManagerSecretTagListToTurbotTags")
	tagList := d.Value.([]types.Tag)

	// Mapping the resource tags inside turbotTags
	var turbotTagsMap map[string]string
	if tagList != nil {
		turbotTagsMap = map[string]string{}
		for _, i := range tagList {
			turbotTagsMap[*i.Key] = *i.Value
		}
	}

	return turbotTagsMap, nil
}
