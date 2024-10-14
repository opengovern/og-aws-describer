package aws

import (
	"context"

	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

//// TABLE DEFINITION

func tableAwsGlueDataCatalogEncryptionSettings(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_glue_data_catalog_encryption_settings",
		Description: "AWS Glue Data Catalog Encryption Settings",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListGlueDataCatalogEncryptionSettings,
		},

		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "encryption_at_rest",
				Description: "A list of public keys to be used by the DataCatalogEncryptionSettingss for authentication.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DataCatalogEncryptionSettings.EncryptionAtRest")},
			{
				Name:        "connection_password_encryption",
				Description: "A list of security group identifiers used in this DataCatalogEncryptionSettings.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DataCatalogEncryptionSettings.ConnectionPasswordEncryption")},
		}),
	}
}
