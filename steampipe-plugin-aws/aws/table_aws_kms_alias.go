package aws

import (
	"context"

	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsKmsAlias(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_kms_alias",
		Description: "AWS KMS Alias",
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListKMSAlias,
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "alias_name",
				Description: "String that contains the alias. This value begins with alias/.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Alias.AliasName")},
			{
				Name:        "arn",
				Description: "String that contains the key ARN.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Alias.AliasArn"),
			},
			{
				Name:        "target_key_id",
				Description: "String that contains the key identifier of the KMS key associated with the alias.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Alias.TargetKeyId")},
			{
				Name:        "creation_date",
				Description: "Date and time that the alias was most recently created in the account and Region.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Alias.CreationDate")},
			{
				Name:        "last_updated_date",
				Description: "Date and time that the alias was most recently associated with a KMS key in the account and Region.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Alias.LastUpdatedDate")},

			/// Standard columns for all tables
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Alias.AliasName")},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Alias.AliasArn").Transform(arnToAkas),
			},
		}),
	}
}
