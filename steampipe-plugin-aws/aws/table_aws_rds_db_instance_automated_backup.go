package aws

import (
	"context"
	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsRDSDBInstanceAutomatedBackup(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_rds_db_instance_automated_backup",
		Description: "AWS RDS DB Instance Automated Backup",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("arn"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"DBInstanceAutomatedBackupNotFound"}),
			},
			Hydrate: opengovernance.GetRDSDBInstanceAutomatedBackup,
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListRDSDBInstanceAutomatedBackup,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"InvalidParameterValue"}),
			},
			KeyColumns: []*plugin.KeyColumn{
				{Name: "db_instance_identifier", Require: plugin.Optional},
				{Name: "dbi_resource_id", Require: plugin.Optional},
				{Name: "status", Require: plugin.Optional},
			},
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "db_instance_identifier",
				Description: "The friendly name to identify the DB Instance.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.InstanceAutomatedBackup.DBInstanceIdentifier"),
			},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) for the replicated automated backups.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.InstanceAutomatedBackup.DBInstanceAutomatedBackupsArn"),
			},
			{
				Name:        "db_instance_arn",
				Description: "The Amazon Resource Name (ARN) for the automated backups.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.InstanceAutomatedBackup.DBInstanceArn"),
			},
			{
				Name:        "status",
				Description: "Specifies the current state of this database.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.InstanceAutomatedBackup.Status"),
			},
			{
				Name:        "allocated_storage",
				Description: "Specifies the allocated storage size in gibibytes (GiB).",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.InstanceAutomatedBackup.AllocatedStorage"),
			},
			{
				Name:        "availability_zone",
				Description: "The Availability Zone that the automated backup was created in.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.InstanceAutomatedBackup.AvailabilityZone"),
			},
			{
				Name:        "backup_retention_period",
				Description: "The retention period for the automated backups.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.InstanceAutomatedBackup.BackupRetentionPeriod"),
			},
			{
				Name:        "backup_target",
				Description: "Specifies where automated backups are stored: Amazon Web Services Outposts or the Amazon Web Services Region.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.InstanceAutomatedBackup.BackupTarget"),
			},
			{
				Name:        "dbi_resource_id",
				Description: "The identifier for the source DB instance, which can't be changed and which is unique to an Amazon Web Services Region.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.InstanceAutomatedBackup.DbiResourceId"),
			},
			{
				Name:        "encrypted",
				Description: "Specifies whether the automated backup is encrypted.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.InstanceAutomatedBackup.Encrypted"),
			},
			{
				Name:        "engine",
				Description: "The name of the database engine for this automated backup.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.InstanceAutomatedBackup.Engine"),
			},
			{
				Name:        "engine_version",
				Description: "The version of the database engine for the automated backup.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.InstanceAutomatedBackup.EngineVersion"),
			},
			{
				Name:        "iam_database_authentication_enabled",
				Description: "True if mapping of Amazon Web Services Identity and Access Management (IAM) accounts to database accounts is enabled, and otherwise false.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.InstanceAutomatedBackup.IAMDatabaseAuthenticationEnabled"),
			},
			{
				Name:        "instance_create_time",
				Description: "True if mapping of Amazon Web Services Identity and Access Management (IAM) accounts to database accounts is enabled, and otherwise false.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.InstanceAutomatedBackup.InstanceCreateTime"),
			},
			{
				Name:        "iops",
				Description: "True if mapping of Amazon Web Services Identity and Access Management (IAM) accounts to database accounts is enabled, and otherwise false.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.InstanceAutomatedBackup.Iops"),
			},
			{
				Name:        "kms_key_id",
				Description: "The Amazon Web Services KMS key ID for an automated backup. The Amazon Web Services KMS key identifier is the key ARN, key ID, alias ARN, or alias name for the KMS key.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.InstanceAutomatedBackup.KmsKeyId"),
			},
			{
				Name:        "license_model",
				Description: "The Amazon Web Services KMS key ID for an automated backup. The Amazon Web Services KMS key identifier is the key ARN, key ID, alias ARN, or alias name for the KMS key.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.InstanceAutomatedBackup.LicenseModel"),
			},
			{
				Name:        "master_username",
				Description: "The license model of an automated backup.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.InstanceAutomatedBackup.MasterUsername"),
			},
			{
				Name:        "option_group_name",
				Description: "The option group the automated backup is associated with. If omitted, the default option group for the engine specified is used.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.InstanceAutomatedBackup.OptionGroupName"),
			},
			{
				Name:        "port",
				Description: "The port number that the automated backup used for connections. Default: Inherits from the source DB instance Valid Values: 1150-65535.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.InstanceAutomatedBackup.Port"),
			},
			{
				Name:        "storage_throughput",
				Description: "Specifies the storage throughput for the automated backup.",
				Type:        proto.ColumnType_INT,
				//Transform:   transform.FromField("Description.InstanceAutomatedBackup.StorageType"),

			},
			{
				Name:        "storage_type",
				Description: "Specifies the storage type associated with the automated backup.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.InstanceAutomatedBackup.StorageType"),
			},
			{
				Name:        "tde_credential_arn",
				Description: "The ARN from the key store with which the automated backup is associated for TDE encryption.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.InstanceAutomatedBackup.TdeCredentialArn"),
			},
			{
				Name:        "timezone",
				Description: "The time zone of the automated backup.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.InstanceAutomatedBackup.Timezone"),
			},
			{
				Name:        "vpc_id",
				Description: "Provides the VPC ID associated with the DB instance.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.InstanceAutomatedBackup.VpcId"),
			},
			{
				Name:        "db_instance_automated_backups_replications",
				Description: "The list of replications to different Amazon Web Services Regions associated with the automated backup.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.InstanceAutomatedBackup.DBInstanceAutomatedBackupsReplications"),
			},
			{
				Name:        "restore_window",
				Description: "Earliest and latest time an instance can be restored to.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.InstanceAutomatedBackup.RestoreWindow"),
			},

			// Stteampipe standard columns
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.InstanceAutomatedBackup.DBInstanceIdentifier"),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.InstanceAutomatedBackup.DBInstanceAutomatedBackupsArn").Transform(arnToAkas),
			},
		}),
	}
}
