package aws

import (
	"context"
	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsBackupVault(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_backup_vault",
		Description: "AWS Backup Vault",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AnyColumn([]string{"name"}),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"InvalidParameter", "AccessDeniedException"}),
			},
			Hydrate: kaytu.GetBackupVault,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListBackupVault,
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The name of a logical container where backups are stored.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.BackupVault.BackupVaultName"),
			},
			{
				Name:        "arn",
				Description: "An Amazon Resource Name (ARN) that uniquely identifies a backup vault.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.BackupVault.BackupVaultArn"),
			},
			{
				Name:        "creation_date",
				Description: "The date and time a resource backup is created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.BackupVault.CreationDate"),
			},
			{
				Name:        "creator_request_id",
				Description: "An unique string that identifies the request and allows failed requests to be retried without the risk of running the operation twice.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.BackupVault.CreatorRequestId"),
			},
			{
				Name:        "encryption_key_arn",
				Description: "The server-side encryption key that is used to protect your backups.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.BackupVault.EncryptionKeyArn"),
			},
			{
				Name:        "number_of_recovery_points",
				Description: "The number of recovery points that are stored in a backup vault.",
				Type:        proto.ColumnType_DOUBLE,
				Transform:   transform.FromField("Description.BackupVault.NumberOfRecoveryPoints"),
			},
			{
				Name:        "sns_topic_arn",
				Description: "An ARN that uniquely identifies an Amazon Simple Notification Service.",
				Type:        proto.ColumnType_STRING,
				//Hydrate:     getAwsBackupVaultNotification, //TODO-Saleh
				Transform: transform.FromField("Description.SNSTopicArn"),
			},
			{
				Name:        "policy",
				Description: "The backup vault access policy document in JSON format.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Policy"),
			},
			{
				Name:        "policy_std",
				Description: "Contains the backup vault access policy document in a canonical form for easier searching.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Policy").Transform(unescape).Transform(policyToCanonical),
			},
			{
				Name:        "backup_vault_events",
				Description: "An array of events that indicate the status of jobs to back up resources to the backup vault.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.BackupVaultEvents"),
			},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.BackupVault.BackupVaultName"),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.BackupVault.BackupVaultArn").Transform(arnToAkas),
			},
		}),
	}
}

//// LIST FUNCTION

//// HYDRATE FUNCTIONS
