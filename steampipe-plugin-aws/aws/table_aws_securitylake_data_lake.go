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
				Name:        "kms_key_id",
				Description: "The id of KMS encryption key used by Amazon Security Lake to encrypt the Security Lake object.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DataLake.EncryptionConfiguration.KmsKeyId")},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) created by you to provide to the subscriber.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DataLake.DataLakeArn")},
			{
				Name:        "create_status",
				Description: "Retrieves the status of the configuration operation for an account in Amazon Security Lake.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DataLake.CreateStatus")},
			{
				Name:        "replication_role_arn",
				Description: "This parameter uses the IAM role created by you that is managed by Security Lake, to ensure the replication setting is correct.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DataLake.ReplicationConfiguration.RoleArn")},
			{
				Name:        "s3_bucket_arn",
				Description: "Amazon Resource Names (ARNs) uniquely identify Amazon Web Services resources.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DataLake.ReplicationConfiguration.S3BucketArn")},
			{
				Name:        "replication_configuration",
				Description: "Provides replication details of Amazon Security Lake object.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DataLake.ReplicationConfiguration")},
			{
				Name:        "lifecycle_configuration",
				Description: "Provides lifecycle details of Amazon Security Lake object.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DataLake.LifecycleConfiguration")},
			{
				Name:        "update_status",
				Description: "The status of the last UpdateDataLake or DeleteDataLake API request.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DataLake.UpdateStatus")},

			// Steampipe standard columns
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DataLake.DataLakeArn").Transform(transform.EnsureStringArray),
			},
		}),
	}
}
