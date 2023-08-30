package aws

import (
	"context"

	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsSecurityLakeDataLake(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_securitylake_data_lake",
		Description: "AWS Security Lake Data Lake",
		List: &plugin.ListConfig{
			Hydrate: kaytu.GetSecurityLakeDataLake,
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "encryption_key",
				Description: "The type of encryption key used by Security Lake to encrypt the lake configuration.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DataLake.EncryptionKey")},
			{
				Name:        "replication_role_arn",
				Description: "This parameter uses the IAM role created by you that is managed by Security Lake, to ensure the replication setting is correct.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DataLake.ReplicationRoleArn")},
			{
				Name:        "s3_bucket_arn",
				Description: "Amazon Resource Names (ARNs) uniquely identify Amazon Web Services resources.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DataLake.S3BucketArn")},
			{
				Name:        "status",
				Description: "Retrieves the status of the configuration operation for an account in Amazon Security Lake.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DataLake.Status")},
			{
				Name:        "replication_destination_regions",
				Description: "Replication enables automatic, asynchronous copying of objects across Amazon S3 buckets.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DataLake.ReplicationDestinationRegions")},
			{
				Name:        "retention_settings",
				Description: "Retention settings for the destination Amazon S3 buckets.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DataLake.RetentionSettings")},

			// Steampipe standard columns
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DataLake.TagsMap")},
		}),
	}
}
