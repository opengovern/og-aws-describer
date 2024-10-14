package aws

import (
	"context"

	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsKmsKey(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_kms_key",
		Description: "AWS KMS Key",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"NotFoundException", "InvalidParameter"}),
			},
			Hydrate: opengovernance.GetKMSKey,
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListKMSKey,
		},

		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "id",
				Description: "Unique identifier of the key.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Metadata.KeyId"),
			},
			{
				Name:        "arn",
				Description: "ARN of the key.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Metadata.Arn"),
			},
			{
				Name:        "key_manager",
				Description: "The manager of the CMK. CMKs in your AWS account are either customer managed or AWS managed.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Metadata.KeyManager")},
			{
				Name:        "creation_date",
				Description: "The date and time when the CMK was created.",
				Type:        proto.ColumnType_TIMESTAMP,

				Transform: transform.FromField("Description.Metadata.CreationDate")},
			{
				Name:        "aws_account_id",
				Description: "The twelve-digit account ID of the AWS account that owns the CMK.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Metadata.AWSAccountId")},
			{
				Name:        "enabled",
				Description: "Specifies whether the CMK is enabled. When KeyState is Enabled this value is true, otherwise it is false.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Metadata.Enabled")},
			{
				Name:        "customer_master_key_spec",
				Description: "Describes the type of key material in the CMK.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Metadata.CustomerMasterKeySpec")},
			{
				Name:        "description",
				Description: "The description of the CMK.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Metadata.Description")},
			{
				Name:        "deletion_date",
				Description: "The date and time after which AWS KMS deletes the CMK.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Metadata.DeletionDate")},
			{
				Name:        "key_state",
				Description: "The current status of the CMK. For more information about how key state affects the use of a CMK, see [Key state: Effect on your CMK](https://docs.aws.amazon.com/kms/latest/developerguide/key-state.html).",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Metadata.KeyState")},
			{
				Name:        "key_usage",
				Description: "The [cryptographic operations](https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#cryptographic-operations) for which you can use the CMK.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Metadata.KeyUsage")},
			{
				Name:        "origin",
				Description: "The source of the CMK's key material. When this value is AWS_KMS, AWS KMS created the key material. When this value is EXTERNAL, the key material was imported from your existing key management infrastructure or the CMK lacks key material. When this value is AWS_CLOUDHSM, the key material was created in the AWS CloudHSM cluster associated with a custom key store.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Metadata.Origin")},
			{
				Name:        "valid_to",
				Description: "The time at which the imported key material expires.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Metadata.ValidTo")},
			{
				Name:        "aliases",
				Description: "A list of aliases for the key.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Aliases"),
			},
			{
				Name:        "key_rotation_enabled",
				Description: "A Boolean value that specifies whether key rotation is enabled.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.KeyRotationEnabled")},
			{
				Name:        "policy",
				Description: "A key policy document in JSON format.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Policy").Transform(transform.UnmarshalYAML),
			},
			{
				Name:        "policy_std",
				Description: "Contains the policy in a canonical form for easier searching.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Policy").Transform(unescape).Transform(policyToCanonical),
			},
			{
				Name:        "tags_src",
				Description: "A list of tags attached to key.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Tags")},
			/// Standard columns for all tables

			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Title")},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Tags")},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Metadata.Arn").Transform(arnToAkas),
			},
		}),
	}
}

//// LIST FUNCTION

//// HYDRATE FUNCTIONS
