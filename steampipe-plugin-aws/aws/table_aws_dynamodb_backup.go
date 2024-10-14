package aws

import (
	"context"
	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAwsDynamoDBBackup(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_dynamodb_backup",
		Description: "AWS DynamoDB Backup",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("arn"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"ValidationException"}),
			},
			Hydrate: opengovernance.GetDynamoDbBackup,
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListDynamoDbBackup,
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:    "backup_type",
					Require: plugin.Optional,
				},
				{
					Name:    "arn",
					Require: plugin.Optional,
				},
				{
					Name:    "table_name",
					Require: plugin.Optional,
				},
			},
		},

		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "Name of the backup.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Backup.BackupName")},
			{
				Name:        "arn",
				Description: "Amazon Resource Name associated with the backup.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Backup.BackupArn"),
			},
			{
				Name:        "table_name",
				Description: "Unique identifier for the table to which backup belongs.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Backup.TableName")},
			{
				Name:        "table_arn",
				Description: "Name of the table to which backup belongs.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Backup.TableArn")},
			{
				Name:        "table_id",
				Description: "ARN associated with the table to which backup belongs.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Backup.TableId")},
			{
				Name:        "backup_status",
				Description: "Current status of the backup. Backup can be in one of the following states: CREATING, ACTIVE, DELETED.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Backup.BackupStatus")},
			{
				Name:        "backup_type",
				Description: "Backup type (USER | SYSTEM | AWS_BACKUP).",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Backup.BackupType")},
			{
				Name:        "backup_creation_datetime",
				Description: "Time at which the backup was created.",
				Type:        proto.ColumnType_TIMESTAMP,

				Transform: transform.FromField("Description.Backup.BackupCreationDateTime")},
			{
				Name:        "backup_expiry_datetime",
				Description: "Time at which the automatic on-demand backup created by DynamoDB will expire.",
				Type:        proto.ColumnType_TIMESTAMP,

				Transform: transform.FromField("Description.Backup.BackupExpiryDateTime")},
			{
				Name:        "backup_size_bytes",
				Description: "Size of the backup in bytes.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Backup.BackupSizeBytes")},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Backup.BackupName")},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Backup.BackupArn").Transform(arnToAkas),
			},
		}),
	}
}
