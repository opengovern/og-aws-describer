package aws

import (
	"context"

	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsRedshiftCluster(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_redshift_cluster",
		Description: "AWS Redshift Cluster",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("cluster_identifier"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"ClusterNotFound"}),
			},
			Hydrate: opengovernance.GetRedshiftCluster,
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListRedshiftCluster,
		},

		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "cluster_identifier",
				Description: "The unique identifier of the cluster.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Cluster.ClusterIdentifier")},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) specifying the cluster.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("ARN"),
			},
			{
				Name:        "cluster_namespace_arn",
				Description: "The namespace Amazon Resource Name (ARN) of the cluster.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Cluster.ClusterNamespaceArn")},
			{
				Name:        "allow_version_upgrade",
				Description: "A boolean value that, if true, indicates that major version upgrades will be applied automatically to the cluster during the maintenance window.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Cluster.AllowVersionUpgrade")},
			{
				Name:        "automated_snapshot_retention_period",
				Description: "The number of days that automatic cluster snapshots are retained.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Cluster.AutomatedSnapshotRetentionPeriod")},
			{
				Name:        "availability_zone",
				Description: "The name of the Availability Zone in which the cluster is located.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Cluster.AvailabilityZone")},
			{
				Name:        "availability_zone_relocation_status",
				Description: "Describes the status of the Availability Zone relocation operation.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Cluster.AvailabilityZoneRelocationStatus")},
			{
				Name:        "cluster_availability_status",
				Description: "The availability status of the cluster for queries.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Cluster.ClusterAvailabilityStatus")},
			{
				Name:        "cluster_create_time",
				Description: "The date and time that the cluster was created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Cluster.ClusterCreateTime")},
			{
				Name:        "cluster_nodes",
				Description: "The nodes in the cluster.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Cluster.ClusterNodes")},
			{
				Name:        "cluster_parameter_groups",
				Description: "The list of cluster parameter groups that are associated with this cluster. Each parameter group in the list is returned with its status.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Cluster.ClusterParameterGroups")},
			{
				Name:        "cluster_public_key",
				Description: "The public key for the cluster.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Cluster.ClusterPublicKey")},
			{
				Name:        "cluster_revision_number",
				Description: "The specific revision number of the database in the cluster.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Cluster.ClusterRevisionNumber")},
			{
				Name:        "cluster_security_groups",
				Description: "A list of cluster security group that are associated with the cluster. Each security group is represented by an element that contains ClusterSecurityGroup.Name and ClusterSecurityGroup.Status subelements. Cluster security groups are used when the cluster is not created in an Amazon Virtual Private Cloud (VPC). Clusters that are created in a VPC use VPC security groups, which are listed by the VpcSecurityGroups parameter.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Cluster.ClusterSecurityGroups")},
			{
				Name:        "cluster_snapshot_copy_status",
				Description: "A value that returns the destination region and retention period that are configured for cross-region snapshot copy.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Cluster.ClusterSnapshotCopyStatus")},
			{
				Name:        "cluster_status",
				Description: "The current state of the cluster.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Cluster.ClusterStatus")},
			{
				Name:        "cluster_subnet_group_name",
				Description: "The name of the subnet group that is associated with the cluster. This parameter is valid only when the cluster is in a VPC.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Cluster.ClusterSubnetGroupName")},
			{
				Name:        "cluster_version",
				Description: "The version ID of the Amazon Redshift engine that is running on the cluster.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Cluster.ClusterVersion")},
			{
				Name:        "data_transfer_progress",
				Description: "Describes the status of a cluster while it is in the process of resizing with an incremental resize.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Cluster.DataTransferProgress")},
			{
				Name:        "db_name",
				Description: "The name of the initial database that was created when the cluster was created. This same name is returned for the life of the cluster. If an initial database was not specified, a database named devdev was created by default.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Cluster.DBName")},
			{
				Name:        "deferred_maintenance_windows",
				Description: "Describes a group of DeferredMaintenanceWindow objects.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Cluster.DeferredMaintenanceWindows")},
			{
				Name:        "elastic_ip_status",
				Description: "The status of the elastic IP (EIP) address.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Cluster.ElasticIpStatus")},
			{
				Name:        "elastic_resize_number_of_node_options",
				Description: "The number of nodes that you can resize the cluster to with the elastic resize method.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Cluster.ElasticResizeNumberOfNodeOptions")},
			{
				Name:        "encrypted",
				Description: "A boolean value that, if true, indicates that data in the cluster is encrypted at rest.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Cluster.Encrypted")},
			{
				Name:        "endpoint",
				Description: "The connection endpoint.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Cluster.Endpoint")},
			{
				Name:        "enhanced_vpc_routing",
				Description: "An option that specifies whether to create the cluster with enhanced VPC routing enabled. To create a cluster that uses enhanced VPC routing, the cluster must be in a VPC. If this option is true, enhanced VPC routing is enabled.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Cluster.EnhancedVpcRouting")},
			{
				Name:        "expected_next_snapshot_schedule_time",
				Description: "The date and time when the next snapshot is expected to be taken for clusters with a valid snapshot schedule and backups enabled.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Cluster.ExpectedNextSnapshotScheduleTime")},
			{
				Name:        "expected_next_snapshot_schedule_time_status",
				Description: "The status of next expected snapshot for clusters having a valid snapshot schedule and backups enabled.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Cluster.ExpectedNextSnapshotScheduleTimeStatus")},
			{
				Name:        "hsm_status",
				Description: "A value that reports whether the Amazon Redshift cluster has finished applying any hardware security module (HSM) settings changes specified in a modify cluster command.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Cluster.HsmStatus")},
			{
				Name:        "iam_roles",
				Description: "A list of AWS Identity and Access Management (IAM) roles that can be used by the cluster to access other AWS services.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Cluster.IamRoles")},
			{
				Name:        "kms_key_id",
				Description: "The AWS Key Management Service (AWS KMS) key ID of the encryption key used to encrypt data in the cluster.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Cluster.KmsKeyId")},
			{
				Name:        "maintenance_track_name",
				Description: "The name of the maintenance track for the cluster.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Cluster.MaintenanceTrackName")},
			{
				Name:        "manual_snapshot_retention_period",
				Description: "The default number of days to retain a manual snapshot. If the value is -1, the snapshot is retained indefinitely. This setting doesn't change the retention period of existing snapshots. The value must be either -1 or an integer between 1 and 3,653.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Cluster.ManualSnapshotRetentionPeriod")},
			{
				Name:        "master_username",
				Description: "The master user name for the cluster. This name is used to connect to the database that is specified in the DBName parameter.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Cluster.MasterUsername")},
			{
				Name:        "modify_status",
				Description: "The status of a modify operation, if any, initiated for the cluster.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Cluster.ModifyStatus")},
			{
				Name:        "next_maintenance_window_start_time",
				Description: "The date and time in UTC when system maintenance can begin.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Cluster.NextMaintenanceWindowStartTime")},
			{
				Name:        "node_type",
				Description: "The node type for the nodes in the cluster.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Cluster.NodeType")},
			{
				Name:        "number_of_nodes",
				Description: "The number of compute nodes in the cluster.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Cluster.NumberOfNodes")},
			{
				Name:        "pending_actions",
				Description: "Cluster operations that are waiting to be started.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Cluster.PendingActions")},
			{
				Name:        "pending_modified_values",
				Description: "A value that, if present, indicates that changes to the cluster are pending. Specific pending changes are identified by subelements.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Cluster.PendingModifiedValues")},
			{
				Name:        "preferred_maintenance_window",
				Description: "The weekly time range, in Universal Coordinated Time (UTC), during which system maintenance can occur.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Cluster.PreferredMaintenanceWindow")},
			{
				Name:        "publicly_accessible",
				Description: "A boolean value that, if true, indicates that the cluster can be accessed from a public network.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Cluster.PubliclyAccessible")},
			{
				Name:        "resize_info",
				Description: "Describes a resize operation.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Cluster.ResizeInfo")},
			{
				Name:        "restore_status",
				Description: "A value that describes the status of a cluster restore action. This parameter returns null if the cluster was not created by restoring a snapshot.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Cluster.RestoreStatus")},
			{
				Name:        "snapshot_schedule_identifier",
				Description: "A unique identifier for the cluster snapshot schedule.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Cluster.SnapshotScheduleIdentifier")},
			{
				Name:        "snapshot_schedule_state",
				Description: "The current state of the cluster snapshot schedule.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Cluster.SnapshotScheduleState")},
			{
				Name:        "vpc_id",
				Description: "The identifier of the VPC the cluster is in, if the cluster is in a VPC.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Cluster.VpcId")},
			{
				Name:        "vpc_security_groups",
				Description: "A list of Amazon Virtual Private Cloud (Amazon VPC) security groups that are associated with the cluster. This parameter is returned only if the cluster is in a VPC.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Cluster.VpcSecurityGroups")},
			{
				Name:        "logging_status",
				Description: "Describes the status of logging for a cluster.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.LoggingStatus")},
			{
				Name:        "scheduled_actions",
				Description: "A list of scheduled actions for specified cluster.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ScheduledActions")},
			{
				Name:        "tags_src",
				Description: "The list of tags for the cluster.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Cluster.Tags")},
			// Steampipe standard columns

			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Cluster.ClusterIdentifier")},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(getRedshiftClusterTurbotTags),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("ARN").Transform(transform.EnsureStringArray),
			},
		}),
	}
}

//// LIST FUNCTION

//// TRANSFORM FUNCTIONS

func getRedshiftClusterTurbotTags(_ context.Context, d *transform.TransformData) (interface{}, error) {
	cluster := d.HydrateItem.(opengovernance.RedshiftCluster).Description.Cluster

	if cluster.Tags != nil {
		turbotTagsMap := map[string]string{}
		for _, i := range cluster.Tags {
			turbotTagsMap[*i.Key] = *i.Value
		}
		return turbotTagsMap, nil
	}
	return nil, nil
}
