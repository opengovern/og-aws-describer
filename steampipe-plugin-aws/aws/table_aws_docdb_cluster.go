package aws

import (
	"context"

	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsDocDBCluster(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_docdb_cluster",
		Description: "AWS DocumentDB Cluster",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("db_cluster_identifier"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"DBClusterNotFoundFault"}),
			},
			Hydrate: opengovernance.GetDocDBCluster,
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListDocDBCluster,
		},

		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "db_cluster_identifier",
				Description: "Contains a user-supplied cluster identifier. This identifier is the unique key that identifies a cluster.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DBCluster.DBClusterIdentifier")},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) for the Cluster.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DBCluster.DBClusterArn"),
			},
			{
				Name:        "status",
				Description: "Specifies the current state of this cluster.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DBCluster.Status")},
			{
				Name:        "cluster_create_time",
				Description: "Specifies the time when the cluster was created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.DBCluster.ClusterCreateTime")},
			{
				Name:        "backup_retention_period",
				Description: "Specifies the number of days for which automatic snapshots are retained.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.DBCluster.BackupRetentionPeriod")},
			{
				Name:        "clone_group_id",
				Description: "Identifies the clone group to which the DB cluster is associated.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DBCluster.CloneGroupId")},
			{
				Name:        "db_cluster_parameter_group",
				Description: "Specifies the name of the cluster parameter group for the cluster.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DBCluster.DBClusterParameterGroup")},
			{
				Name:        "db_cluster_resource_id",
				Description: "The Region-unique, immutable identifier for the cluster.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DBCluster.DbClusterResourceId")},
			{
				Name:        "db_subnet_group",
				Description: "Specifies information on the subnet group associated with the cluster.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DBCluster.DBSubnetGroup")},
			{
				Name:        "deletion_protection",
				Description: "Specifies whether the cluster has deletion protection enabled, or not.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.DBCluster.DeletionProtection")},
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
				Name:        "engine_version",
				Description: "Indicates the database engine version.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DBCluster.EngineVersion")},
			{
				Name:        "hosted_zone_id",
				Description: "Specifies the ID that Amazon Route 53 assigns when you create a hosted zone.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DBCluster.HostedZoneId")},
			{
				Name:        "kms_key_id",
				Description: "The AWS KMS key identifier for the encrypted cluster.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DBCluster.KmsKeyId")},
			{
				Name:        "latest_restorable_time",
				Description: "Specifies the latest time to which a database can be restored with point-in-time restore.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.DBCluster.LatestRestorableTime")},
			{
				Name:        "master_user_name",
				Description: "Contains the master username for the cluster.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DBCluster.MasterUsername")},
			{
				Name:        "multi_az",
				Description: "Specifies whether the cluster has instances in multiple Availability Zones, or not.",
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
				Name:        "replication_source_identifier",
				Description: "Contains the identifier of the source cluster if this cluster is a secondary cluster.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DBCluster.ReplicationSourceIdentifier")},
			{
				Name:        "storage_encrypted",
				Description: "Specifies whether the cluster is encrypted, or not.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.DBCluster.StorageEncrypted")},
			{
				Name:        "associated_roles",
				Description: "A list of AWS IAM roles that are associated with the cluster.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DBCluster.AssociatedRoles")},
			{
				Name:        "availability_zones",
				Description: "A list of Availability Zones (AZs) where instances in the cluster can be created.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DBCluster.AvailabilityZones")},
			{
				Name:        "enabled_cloudwatch_logs_exports",
				Description: "A list of log types that this cluster is configured to export to Amazon CloudWatch Logs.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DBCluster.EnabledCloudwatchLogsExports")},
			{
				Name:        "members",
				Description: "A list of instances that make up the cluster.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DBCluster.DBClusterMembers")},
			{
				Name:        "read_replica_identifiers",
				Description: "A list of identifiers of the read replicas associated with this cluster.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DBCluster.ReadReplicaIdentifiers")},
			{
				Name:        "vpc_security_groups",
				Description: "A list of VPC security groups that the DB cluster belongs to.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DBCluster.VpcSecurityGroups")},
			{
				Name:        "tags_src",
				Description: "A list of tags attached to the Cluster.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Tags")},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Tags").Transform(docDBClusterTagListToTurbotTags),
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
				Transform:   transform.FromField("Description.DBCluster.DBClusterArn").Transform(transform.EnsureStringArray),
			},
		}),
	}
}

func docDBClusterTagListToTurbotTags(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	plugin.Logger(ctx).Trace("docDBClusterTagListToTurbotTags")
	tagList := d.HydrateItem.(opengovernance.DocDBCluster).Description.Tags

	// Mapping the resource tags inside turbotTags
	var turbotTagsMap map[string]string
	if tagList != nil {
		turbotTagsMap = map[string]string{}
		for _, i := range tagList {
			turbotTagsMap[*i.Key] = *i.Value
		}
	}

	return turbotTagsMap, nil
}
