package aws

import (
	"context"
	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

// // TABLE DEFINITION
// error
// check the function the end of the codes
func tableAwsNeptuneDBCluster(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_neptune_db_cluster",
		Description: "AWS Neptune DB Cluster",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("db_cluster_identifier"),
			Hydrate:    kaytu.GetNeptuneDatabaseCluster,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"DBClusterNotFoundFault"}),
			},
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListNeptuneDatabaseCluster,
		},

		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "db_cluster_identifier",
				Description: "Contains a user-supplied DB cluster identifier. This identifier is the unique key that identifies a DB cluster.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Cluster.DBClusterIdentifier")},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) for the DB cluster.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Cluster.DBClusterArn"),
			},
			{
				Name:        "status",
				Description: "Specifies the current state of this DB cluster.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Cluster.Status")},
			{
				Name:        "cluster_create_time",
				Description: "Specifies the time when the DB cluster was created, in Universal Coordinated Time (UTC).",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Cluster.ClusterCreateTime")},
			{
				Name:        "allocated_storage",
				Description: "AllocatedStorage always returns 1, because Neptune DB cluster storage size is not fixed, but instead automatically adjusts as needed.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Cluster.AllocatedStorage")},
			{
				Name:        "automatic_restart_time",
				Description: "Time at which the DB cluster will be automatically restarted.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Cluster.AutomaticRestartTime")},
			{
				Name:        "backup_retention_period",
				Description: "Specifies the number of days for which automatic DB snapshots are retained.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Cluster.BackupRetentionPeriod")},
			{
				Name:        "clone_group_id",
				Description: "Identifies the clone group to which the DB cluster is associated.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Cluster.CloneGroupId")},
			{
				Name:        "copy_tags_to_snapshot",
				Description: "If set to true, tags are copied to any snapshot of the DB cluster that is created.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Cluster.CopyTagsToSnapshot")},
			{
				Name:        "cross_account_clone",
				Description: "If set to true, the DB cluster can be cloned across accounts.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Cluster.CrossAccountClone")},
			{
				Name:        "db_cluster_parameter_group",
				Description: "Specifies the name of the DB cluster parameter group for the DB cluster.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Cluster.DBClusterParameterGroup")},
			{
				Name:        "db_subnet_group",
				Description: "Specifies information on the subnet group associated with the DB cluster.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Cluster.DBSubnetGroup")},
			{
				Name:        "database_name",
				Description: "Contains the name of the initial database of this DB cluster that was provided.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Cluster.DatabaseName")},
			{
				Name:        "db_cluster_resource_id",
				Description: "The Amazon Region-unique, immutable identifier for the DB cluster.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Cluster.DbClusterResourceId")},
			{
				Name:        "deletion_protection",
				Description: "Indicates whether or not the DB cluster has deletion protection enabled.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Cluster.DeletionProtection")},
			{
				Name:        "earliest_restorable_time",
				Description: "Specifies the earliest time to which a database can be restored with point-in-time restore.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Cluster.EarliestRestorableTime")},
			{
				Name:        "endpoint",
				Description: "Specifies the connection endpoint for the primary instance of the DB cluster.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Cluster.Endpoint")},
			{
				Name:        "engine",
				Description: "Provides the name of the database engine to be used for this DB cluster.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Cluster.Engine")},
			{
				Name:        "engine_version",
				Description: "Indicates the database engine version.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Cluster.EngineVersion")},
			{
				Name:        "hosted_zone_id",
				Description: "Specifies the ID that Amazon Route 53 assigns when you create a hosted zone.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Cluster.HostedZoneId")},
			{
				Name:        "iam_database_authentication_enabled",
				Description: "True if mapping of Amazon Identity and Access Management (IAM) accounts to database accounts is enabled, and otherwise false.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Cluster.IAMDatabaseAuthenticationEnabled")},
			{
				Name:        "kms_key_id",
				Description: "If StorageEncrypted is true, the Amazon KMS key identifier for the encrypted DB cluster.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Cluster.KmsKeyId")},
			{
				Name:        "latest_restorable_time",
				Description: "Specifies the latest time to which a database can be restored with point-in-time restore.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Cluster.LatestRestorableTime")},
			{
				Name:        "multi_az",
				Description: "Specifies whether the DB cluster has instances in multiple Availability Zones.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Cluster.MultiAZ")},
			{
				Name:        "percent_progress",
				Description: "Specifies the progress of the operation as a percentage.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Cluster.PercentProgress")},
			{
				Name:        "port",
				Description: "Specifies the port that the database engine is listening on.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Cluster.Port")},
			{
				Name:        "preferred_backup_window",
				Description: "Specifies the daily time range during which automated backups are created if automated backups are enabled, as determined by the BackupRetentionPeriod.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Cluster.PreferredBackupWindow")},
			{
				Name:        "preferred_maintenance_window",
				Description: "Specifies the weekly time range during which system maintenance can occur, in Universal Coordinated Time (UTC).",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Cluster.PreferredMaintenanceWindow")},
			{
				Name:        "reader_endpoint",
				Description: "The reader endpoint for the DB cluster.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Cluster.ReaderEndpoint")},
			{
				Name:        "storage_encrypted",
				Description: "Specifies whether the DB cluster is encrypted.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Cluster.StorageEncrypted")},
			{
				Name:        "associated_roles",
				Description: "Provides a list of the Amazon Identity and Access Management (IAM) roles.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Cluster.AssociatedRoles")},
			{
				Name:        "availability_zones",
				Description: "Provides the list of EC2 Availability Zones that instances in the DB cluster can be created in.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Cluster.AvailabilityZones")},
			{
				Name:        "db_cluster_members",
				Description: "Provides the list of instances that make up the DB cluster.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Cluster.DBClusterMembers")},
			{
				Name:        "enabled_cloudwatch_logs_exports",
				Description: "A list of log types that this DB cluster is configured to export to CloudWatch Logs.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Cluster.EnabledCloudwatchLogsExports")},
			{
				Name:        "read_replica_identifiers",
				Description: "Contains one or more identifiers of the Read Replicas associated with this DB cluster.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Cluster.ReadReplicaIdentifiers")},
			{
				Name:        "vpc_security_groups",
				Description: "Provides a list of VPC security groups that the DB cluster belongs to.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Cluster.VpcSecurityGroups")},

			{
				Name:        "tags_src",
				Description: "A list of tags assigned to the resource.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Tags")},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Cluster.DBClusterIdentifier")},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Tags").Transform(neptuneDBClusterTurbotTags),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,

				Transform: transform.

					//// LIST FUNCTION
					FromField("Description.Cluster.DBClusterArn").Transform(transform.EnsureStringArray),
			},
		}),
	}
}

func neptuneDBClusterTurbotTags(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	tagsDetails := d.HydrateItem.(kaytu.NeptuneDatabaseCluster).Description.Tags

	// Mapping the resource tags inside turbotTags
	var turbotTagsMap map[string]string
	if tagsDetails != nil {
		turbotTagsMap = map[string]string{}
		for _, i := range tagsDetails {
			turbotTagsMap[*i.Key] = *i.Value
		}
	}

	return turbotTagsMap, nil
}
