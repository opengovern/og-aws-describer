package aws

import (
	"context"
	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsRDSDBCluster(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_rds_db_cluster",
		Description: "AWS RDS DB Cluster",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("db_cluster_identifier"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"DBClusterNotFoundFault"}),
			},
			Hydrate: kaytu.GetRDSDBCluster,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListRDSDBCluster,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "clone_group_id", Require: plugin.Optional},
				{Name: "engine", Require: plugin.Optional},
			},
		},

		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "db_cluster_identifier",
				Description: "The friendly name to identify the DB Cluster.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DBCluster.DBClusterIdentifier")},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) for the DB Cluster.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("ARN"),
			},
			{
				Name:        "status",
				Description: "Specifies the status of this DB Cluster.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DBCluster.Status")},
			{
				Name:        "resource_id",
				Description: "The AWS Region-unique, immutable identifier for the DB cluster.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DBCluster.DbClusterResourceId")},
			{
				Name:        "create_time",
				Description: "Specifies the time when the DB cluster was created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.DBCluster.ClusterCreateTime")},
			{
				Name:        "activity_stream_kinesis_stream_name",
				Description: "The name of the Amazon Kinesis data stream used for the database activity stream.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DBCluster.ActivityStreamKinesisStreamName")},
			{
				Name:        "activity_stream_kms_key_id",
				Description: "The AWS KMS key identifier used for encrypting messages in the database activity stream.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DBCluster.ActivityStreamKmsKeyId")},
			{
				Name:        "activity_stream_mode",
				Description: "The mode of the database activity stream.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DBCluster.ActivityStreamMode")},
			{
				Name:        "activity_stream_status",
				Description: "The status of the database activity stream.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DBCluster.ActivityStreamStatus")},
			{
				Name:        "allocated_storage",
				Description: "Specifies the allocated storage size in gibibytes (GiB).",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.DBCluster.AllocatedStorage")},
			{
				Name:        "backtrack_consumed_change_records",
				Description: "The number of change records stored for Backtrack.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.DBCluster.BacktrackConsumedChangeRecords")},
			{
				Name:        "backtrack_window",
				Description: "The target backtrack window, in seconds.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.DBCluster.BacktrackWindow")},
			{
				Name:        "backup_retention_period",
				Description: "Specifies the number of days for which automatic DB snapshots are retained.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.DBCluster.BackupRetentionPeriod")},
			{
				Name:        "capacity",
				Description: "The current capacity of an Aurora Serverless DB cluster.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.DBCluster.Capacity")},
			{
				Name:        "character_set_name",
				Description: "Specifies the name of the character set that this cluster is associated with.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DBCluster.CharacterSetName")},
			{
				Name:        "clone_group_id",
				Description: "Identifies the clone group to which the DB cluster is associated.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DBCluster.CloneGroupId")},
			{
				Name:        "copy_tags_to_snapshot",
				Description: "Specifies whether tags are copied from the DB cluster to snapshots of the DB cluster, or not.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.DBCluster.CopyTagsToSnapshot")},
			{
				Name:        "cross_account_clone",
				Description: "Specifies whether the DB cluster is a clone of a DB cluster owned by a different AWS account, or not.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.DBCluster.CrossAccountClone")},
			{
				Name:        "database_name",
				Description: "Contains the name of the initial database of this DB cluster that was provided at create time.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DBCluster.DatabaseName")},
			{
				Name:        "db_cluster_parameter_group",
				Description: "Specifies the name of the DB cluster parameter group for the DB cluster.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DBCluster.DBClusterParameterGroup")},
			{
				Name:        "db_subnet_group",
				Description: "Specifies information on the subnet group associated with the DB cluster.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DBCluster.DBSubnetGroup")},
			{
				Name:        "deletion_protection",
				Description: "Specifies whether the DB cluster has deletion protection enabled, or not.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.DBCluster.DeletionProtection")},
			{
				Name:        "earliest_backtrack_time",
				Description: "The earliest time to which a DB cluster can be backtracked.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.DBCluster.EarliestBacktrackTime")},
			{
				Name:        "earliest_restorable_time",
				Description: "The earliest time to which a database can be restored with point-in-time restore.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.DBCluster.EarliestRestorableTime")},
			{
				Name:        "endpoint",
				Description: "Specifies the connection endpoint for the primary instance of the DB cluster.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DBCluster.Endpoint")},
			{
				Name:        "engine",
				Description: "The name of the database engine to be used for this DB cluster.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DBCluster.Engine")},
			{
				Name:        "engine_mode",
				Description: "The DB engine mode of the DB cluster.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DBCluster.EngineMode")},
			{
				Name:        "engine_version",
				Description: "Indicates the database engine version.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DBCluster.EngineVersion")},
			{
				Name:        "global_write_forwarding_requested",
				Description: "Specifies whether you have requested to enable write forwarding for a secondary cluster in an Aurora global database, or not.",
				Type:        proto.ColumnType_BOOL,
				Default:     false,
				Transform:   transform.FromField("Description.DBCluster.GlobalWriteForwardingRequested")},
			{
				Name:        "global_write_forwarding_status",
				Description: "Specifies whether a secondary cluster in an Aurora global database has write forwarding enabled, or not.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DBCluster.GlobalWriteForwardingStatus")},
			{
				Name:        "hosted_zone_id",
				Description: "Specifies the ID that Amazon Route 53 assigns when you create a hosted zone.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DBCluster.HostedZoneId")},
			{
				Name:        "http_endpoint_enabled",
				Description: "Specifies whether the HTTP endpoint for an Aurora Serverless DB cluster is enabled, or not.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.DBCluster.HttpEndpointEnabled")},
			{
				Name:        "iam_database_authentication_enabled",
				Description: "Specifies whether the the mapping of AWS IAM accounts to database accounts is enabled, or not.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.DBCluster.IAMDatabaseAuthenticationEnabled")},
			{
				Name:        "kms_key_id",
				Description: "The AWS KMS key identifier for the encrypted DB cluster.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DBCluster.KmsKeyId")},
			{
				Name:        "latest_restorable_time",
				Description: "Specifies the latest time to which a database can be restored with point-in-time restore.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.DBCluster.LatestRestorableTime")},
			{
				Name:        "master_user_name",
				Description: "Contains the master username for the DB cluster.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DBCluster.MasterUsername")},
			{
				Name:        "multi_az",
				Description: "Specifies whether the DB cluster has instances in multiple Availability Zones, or not.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.DBCluster.MultiAZ")},
			{
				Name:        "percent_progress",
				Description: "Specifies the progress of the operation as a percentage.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DBCluster.PercentProgress")},
			{
				Name:        "port",
				Description: "Specifies the port that the database engine is listening on.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.DBCluster.Port")},
			{
				Name:        "preferred_backup_window",
				Description: "Specifies the daily time range during which automated backups are created.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DBCluster.PreferredBackupWindow")},
			{
				Name:        "preferred_maintenance_window",
				Description: "Specifies the weekly time range during which system maintenance can occur",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DBCluster.PreferredMaintenanceWindow")},
			{
				Name:        "reader_endpoint",
				Description: "The reader endpoint for the DB cluster.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DBCluster.ReaderEndpoint")},
			{
				Name:        "storage_encrypted",
				Description: "Specifies whether the DB cluster is encrypted, or not.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.DBCluster.StorageEncrypted")},
			{
				Name:        "associated_roles",
				Description: "A list of AWS IAM roles that are associated with the DB cluster.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DBCluster.AssociatedRoles")},
			{
				Name:        "availability_zones",
				Description: "A list of Availability Zones (AZs) where instances in the DB cluster can be created.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DBCluster.AvailabilityZones")},
			{
				Name:        "custom_endpoints",
				Description: "A list of all custom endpoints associated with the cluster.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DBCluster.CustomEndpoints")},
			{
				Name:        "members",
				Description: "A list of instances that make up the DB cluster.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DBCluster.DBClusterMembers")},
			{
				Name:        "option_group_memberships",
				Description: "A list of option group memberships for this DB cluster.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DBCluster.DBClusterOptionGroupMemberships")},
			{
				Name:        "domain_memberships",
				Description: "A list of Active Directory Domain membership records associated with the DB cluster.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DBCluster.DomainMemberships")},
			{
				Name:        "enabled_cloudwatch_logs_exports",
				Description: "A list of log types that this DB cluster is configured to export to CloudWatch Logs.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DBCluster.EnabledCloudwatchLogsExports")},
			{
				Name:        "read_replica_identifiers",
				Description: "A list of identifiers of the read replicas associated with this DB cluster.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DBCluster.ReadReplicaIdentifiers")},
			{
				Name:        "vpc_security_groups",
				Description: "A list of VPC security groups that the DB cluster belongs to.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DBCluster.VpcSecurityGroups")},
			{
				Name:        "tags_src",
				Description: "A list of tags attached to the DB Cluster.",
				Type:        proto.ColumnType_JSON,

				Transform: transform.FromField("Description.DBCluster.TagList")},
			// Standard columns
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(getRDSDBClusterTurbotTags),
			},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DBCluster.DBClusterIdentifier")},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DBCluster.DBClusterArn").Transform(arnToAkas),
			},
		}),
	}
}

//// LIST FUNCTION

//// HYDRATE FUNCTIONS

//// TRANSFORM FUNCTIONS

func getRDSDBClusterTurbotTags(_ context.Context, d *transform.TransformData) (interface{}, error) {
	dbCluster := d.HydrateItem.(kaytu.RDSDBCluster).Description.DBCluster

	if dbCluster.TagList != nil {
		turbotTagsMap := map[string]string{}
		for _, i := range dbCluster.TagList {
			turbotTagsMap[*i.Key] = *i.Value
		}
		return turbotTagsMap, nil
	}
	return nil, nil
}
