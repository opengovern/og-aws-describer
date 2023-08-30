package aws

import (
	"context"
	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsBackupRecoveryPoint(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_backup_recovery_point",
		Description: "AWS Backup Recovery Point",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"backup_vault_name", "recovery_point_arn"}),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"NotFoundException", "AccessDeniedException"}),
			},
			Hydrate: kaytu.GetBackupRecoveryPoint,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListBackupRecoveryPoint,
			KeyColumns: []*plugin.KeyColumn{
				// {
				// 	Name:    "recovery_point_arn",
				// 	Require: plugin.Optional,
				// },
				{
					Name:    "resource_type",
					Require: plugin.Optional,
				},
				{
					Name:    "completion_date",
					Require: plugin.Optional,
				},
			},
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "backup_vault_name",
				Description: "The name of a logical container where backups are stored.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.RecoveryPoint.BackupVaultName"),
			},
			{
				Name:        "recovery_point_arn",
				Description: "An ARN that uniquely identifies a recovery point.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.RecoveryPoint.RecoveryPointArn"),
			},
			{
				Name:        "resource_type",
				Description: "The type of Amazon Web Services resource to save as a recovery point.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.RecoveryPoint.ResourceType"),
			},
			{
				Name:        "status",
				Description: "A status code specifying the state of the recovery point.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.RecoveryPoint.Status"),
			},
			{
				Name:        "backup_size_in_bytes",
				Description: "The size, in bytes, of a backup.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.RecoveryPoint.BackupSizeInBytes"),
			},
			{
				Name:        "backup_vault_arn",
				Description: "An ARN that uniquely identifies a backup vault.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.RecoveryPoint.BackupVaultArn"),
			},
			{
				Name:        "creation_date",
				Description: "The date and time that a recovery point is created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.RecoveryPoint.CreationDate"),
			},
			{
				Name:        "completion_date",
				Description: "The date and time that a job to create a recovery point is completed.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.RecoveryPoint.CompletionDate"),
			},
			{
				Name:        "encryption_key_arn",
				Description: "The server-side encryption key used to protect your backups.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.RecoveryPoint.EncryptionKeyArn"),
			},
			{
				Name:        "iam_role_arn",
				Description: "Specifies the IAM role ARN used to create the target recovery point.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.RecoveryPoint.IamRoleArn"),
			},
			{
				Name:        "is_encrypted",
				Description: "A Boolean value that is returned as TRUE if the specified recovery point is encrypted, or FALSE if the recovery point is not encrypted.",
				Type:        proto.ColumnType_BOOL,
				Default:     false,
				Transform:   transform.FromField("Description.RecoveryPoint.IsEncrypted"),
			},
			{
				Name:        "last_restore_time",
				Description: "The date and time that a recovery point was last restored.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.RecoveryPoint.LastRestoreTime"),
			},
			{
				Name:        "resource_arn",
				Description: "An ARN that uniquely identifies a saved resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.RecoveryPoint.ResourceArn"),
			},
			{
				Name:        "source_backup_vault_arn",
				Description: "An Amazon Resource Name (ARN) that uniquely identifies the source vault where the resource was originally backed up in.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.RecoveryPoint.SourceBackupVaultArn"),
			},
			{
				Name:        "status_message",
				Description: "A status message explaining the reason for the recovery point deletion failure.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.RecoveryPoint.StatusMessage"),
			},
			{
				Name:        "storage_class",
				Description: "Specifies the storage class of the recovery point. Valid values are WARM or COLD.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.RecoveryPoint.StorageClass"),
			},
			{
				Name:        "calculated_lifecycle",
				Description: "An object containing DeleteAt and MoveToColdStorageAt timestamps.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.RecoveryPoint.CalculatedLifecycle"),
			},
			{
				Name:        "created_by",
				Description: "Contains identifying information about the creation of a recovery point, including the BackupPlanArn, BackupPlanId, BackupPlanVersion, and BackupRuleId of the backup plan used to create it.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.RecoveryPoint.CreatedBy"),
			},
			{
				Name:        "lifecycle",
				Description: "The lifecycle defines when a protected resource is transitioned to cold storage and when it expires.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.RecoveryPoint.Lifecycle"),
			},

			// Steampipe standard columns
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.RecoveryPoint.ResourceArn").Transform(arnToAkas),
			},
		}),
	}
}

//// LIST FUNCTION

//// HYDRATE FUNCTION
