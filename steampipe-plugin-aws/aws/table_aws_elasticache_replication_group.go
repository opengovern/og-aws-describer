package aws

import (
	"context"

	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsElastiCacheReplicationGroup(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_elasticache_replication_group",
		Description: "AWS ElastiCache Replication Group",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("replication_group_id"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"ReplicationGroupNotFoundFault", "InvalidParameterValue"}),
			},
			Hydrate: kaytu.GetElastiCacheReplicationGroup,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListElastiCacheReplicationGroup,
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "replication_group_id",
				Description: "The identifier for the replication group.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ReplicationGroup.ReplicationGroupId")},
			{
				Name:        "arn",
				Description: "The ARN (Amazon Resource Name) of the replication group.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ReplicationGroup.ARN"),
			},
			{
				Name:        "description",
				Description: "The user supplied description of the replication group.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ReplicationGroup.Description")},
			{
				Name:        "at_rest_encryption_enabled",
				Description: "A flag that enables encryption at-rest when set to true.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.ReplicationGroup.AtRestEncryptionEnabled")},
			{
				Name:        "kms_key_id",
				Description: "The ID of the KMS key used to encrypt the disk in the cluster.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ReplicationGroup.KmsKeyId")},
			{
				Name:        "auth_token_enabled",
				Description: "A flag that enables using an AuthToken (password) when issuing Redis commands.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.ReplicationGroup.AuthTokenEnabled")},
			{
				Name:        "auth_token_last_modified_date",
				Description: "The date when the auth token was last modified.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.ReplicationGroup.AuthTokenLastModifiedDate")},
			{
				Name:        "automatic_failover",
				Description: "Indicates the status of automatic failover for this Redis replication group.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ReplicationGroup.AutomaticFailover")},
			{
				Name:        "cache_node_type",
				Description: "The name of the compute and memory capacity node type for each node in the replication group.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ReplicationGroup.CacheNodeType")},
			{
				Name:        "cluster_enabled",
				Description: "A flag indicating whether or not this replication group is cluster enabled.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.ReplicationGroup.ClusterEnabled")},
			{
				Name:        "multi_az",
				Description: "A flag indicating if you have Multi-AZ enabled to enhance fault tolerance.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.ReplicationGroup.MultiAZ")},
			{
				Name:        "snapshot_retention_limit",
				Description: "The number of days for which ElastiCache retains automatic cluster snapshots before deleting them.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.ReplicationGroup.SnapshotRetentionLimit")},
			{
				Name:        "snapshot_window",
				Description: "The daily time range (in UTC) during which ElastiCache begins taking a daily snapshot of your node group (shard).",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ReplicationGroup.SnapshotWindow")},
			{
				Name:        "snapshotting_cluster_id",
				Description: "The cluster ID that is used as the daily snapshot source for the replication group.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ReplicationGroup.SnapshottingClusterId")},
			{
				Name:        "status",
				Description: "The current state of this replication group - creating, available, modifying, deleting, create-failed, snapshotting.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ReplicationGroup.Status")},
			{
				Name:        "transit_encryption_enabled",
				Description: "A flag that enables in-transit encryption when set to true.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.ReplicationGroup.TransitEncryptionEnabled")},
			{
				Name:        "configuration_endpoint",
				Description: "The configuration endpoint for this replication group.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ReplicationGroup.ConfigurationEndpoint")},
			{
				Name:        "global_replication_group_info",
				Description: "The name of the Global Datastore and role of this replication group in the Global Datastore.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ReplicationGroup.GlobalReplicationGroupInfo")},
			{
				Name:        "member_clusters",
				Description: "The names of all the cache clusters that are part of this replication group.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ReplicationGroup.MemberClusters")},
			{
				Name:        "member_clusters_outpost_arns",
				Description: "The outpost ARNs of the replication group's member clusters.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ReplicationGroup.MemberClustersOutpostArns")},
			{
				Name:        "node_groups",
				Description: "A list of node groups in this replication group.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ReplicationGroup.NodeGroups")},
			{
				Name:        "pending_modified_values",
				Description: "A group of settings to be applied to the replication group, either immediately or during the next maintenance window.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ReplicationGroup.PendingModifiedValues")},
			{
				Name:        "user_group_ids",
				Description: "The list of user group IDs that have access to the replication group.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ReplicationGroup.UserGroupIds")},
			// Standard columns

			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.ReplicationGroup.ReplicationGroupId")},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ReplicationGroup.ARN").Transform(arnToAkas),
			},
		}),
	}
}

//// LIST FUNCTION

//// HYDRATE FUNCTION
