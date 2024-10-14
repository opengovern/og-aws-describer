package aws

import (
	"context"
	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsKmsKeyRotation(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_kms_key_rotation",
		Description: "AWS KMS Key Rotation",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListKMSKeyRotation,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"NotFoundException"}),
			},
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "key_id",
				Description: "Unique identifier of the key.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.KeyId"),
			},
			{
				Name:        "key_arn",
				Description: "ARN of the key.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.KeyArn"),
			},
			{
				Name:        "rotation_type",
				Description: "Identifies whether the key material rotation was a scheduled automatic rotation or an on-demand rotation.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.RotationType"),
			},
			{
				Name:        "rotation_date",
				Description: "Date and time that the key material rotation completed.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.RotationDate"),
			},

			/// Standard columns for all tables
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.KeyId"),
			},
		}),
	}
}
