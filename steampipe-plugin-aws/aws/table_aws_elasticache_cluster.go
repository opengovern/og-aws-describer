package aws

import (
	"context"

	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsElastiCacheCluster(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_elasticache_cluster",
		Description: "AWS ElastiCache Cluster",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("cache_cluster_id"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"CacheClusterNotFound", "InvalidParameterValue"}),
			},
			Hydrate: opengovernance.GetElastiCacheCluster,
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListElastiCacheCluster,
		},

		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "cache_cluster_id",
				Description: "An unique identifier for ElastiCache cluster.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Cluster.CacheClusterId")},
			{
				Name:        "arn",
				Description: "The ARN (Amazon Resource Name) of the cache cluster.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Cluster.ARN"),
			},
			{
				Name:        "cache_node_type",
				Description: "The name of the compute and memory capacity node type for the cluster.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Cluster.CacheNodeType")},
			{
				Name:        "cache_cluster_status",
				Description: "The current state of this cluster, one of the following values: available, creating, deleted, deleting, incompatible-network, modifying, rebooting cluster nodes, restore-failed, or snapshotting.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Cluster.CacheClusterStatus")},
			{
				Name:        "at_rest_encryption_enabled",
				Description: "A flag that enables encryption at-rest when set to true.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Cluster.AtRestEncryptionEnabled")},
			{
				Name:        "auth_token_enabled",
				Description: "A flag that enables using an AuthToken (password) when issuing Redis commands.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Cluster.AuthTokenEnabled")},
			{
				Name:        "auto_minor_version_upgrade",
				Description: "This parameter is currently disabled.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Cluster.AutoMinorVersionUpgrade")},
			{
				Name:        "cache_cluster_create_time",
				Description: "The date and time when the cluster was created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Cluster.CacheClusterCreateTime")},
			{
				Name:        "cache_subnet_group_name",
				Description: "The name of the cache subnet group associated with the cluster.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Cluster.CacheSubnetGroupName")},
			{
				Name:        "client_download_landing_page",
				Description: "The URL of the web page where you can download the latest ElastiCache client library.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Cluster.ClientDownloadLandingPage")},
			{
				Name:        "configuration_endpoint",
				Description: "Represents a Memcached cluster endpoint which can be used by an application to connect to any node in the cluster.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Cluster.ConfigurationEndpoint")},
			{
				Name:        "engine",
				Description: "The name of the cache engine (memcached or redis) to be used for this cluster.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Cluster.Engine")},
			{
				Name:        "engine_version",
				Description: "The version of the cache engine that is used in this cluster.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Cluster.EngineVersion")},
			{
				Name:        "num_cache_nodes",
				Description: "The number of cache nodes in the cluster.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Cluster.NumCacheNodes")},
			{
				Name:        "preferred_availability_zone",
				Description: "The name of the Availability Zone in which the cluster is located or 'Multiple' if the cache nodes are located in different Availability Zones.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Cluster.PreferredAvailabilityZone")},
			{
				Name:        "preferred_maintenance_window",
				Description: "Specifies the weekly time range during which maintenance on the cluster is performed.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Cluster.PreferredMaintenanceWindow")},
			{
				Name:        "replication_group_id",
				Description: "The replication group to which this cluster belongs.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Cluster.ReplicationGroupId")},
			{
				Name:        "snapshot_retention_limit",
				Description: "The number of days for which ElastiCache retains automatic cluster snapshots before deleting them.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Cluster.SnapshotRetentionLimit")},
			{
				Name:        "snapshot_window",
				Description: "The daily time range (in UTC) during which ElastiCache begins taking a daily snapshot of your cluster.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Cluster.SnapshotWindow")},
			{
				Name:        "transit_encryption_enabled",
				Description: "A flag that enables in-transit encryption when set to true.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Cluster.TransitEncryptionEnabled")},
			{
				Name:        "auth_token_last_modified_date",
				Description: "The date the auth token was last modified.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Cluster.AuthTokenLastModifiedDate")},
			{
				Name:        "ip_discovery",
				Description: "The network type associated with the cluster, either ipv4 or ipv6.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Cluster.IpDiscovery")},
			{
				Name:        "network_type",
				Description: "Must be either ipv4, ipv6, or dual_stack.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Cluster.NetworkType")},
			{
				Name:        "preferred_outpost_arn",
				Description: "The outpost ARN in which the cache cluster is created.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Cluster.PreferredOutpostArn")},
			{
				Name:        "replication_group_log_delivery_enabled",
				Description: "A boolean value indicating whether log delivery is enabled for the replication group.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Cluster.ReplicationGroupLogDeliveryEnabled")},
			{
				Name:        "transit_encryption_mode",
				Description: "A setting that allows you to migrate your clients to use in-transit encryption, with no downtime.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Cluster.TransitEncryptionMode")},
			{
				Name:        "cache_parameter_group",
				Description: "Status of the cache parameter group.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Cluster.CacheParameterGroup")},
			{
				Name:        "notification_configuration",
				Description: "Describes a notification topic and its status.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Cluster.NotificationConfiguration")},
			{
				Name:        "pending_modified_values",
				Description: "A group of settings that are applied to the cluster in the future, or that are currently being applied.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Cluster.PendingModifiedValues")},
			{
				Name:        "security_groups",
				Description: "A list of VPC Security Groups associated with the cluster.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Cluster.SecurityGroups")},
			{
				Name:        "cache_nodes",
				Description: "A list of cache nodes that are members of the cluster.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Cluster.CacheNodes")},
			{
				Name:        "cache_security_groups",
				Description: "A list of cache security group elements, composed of name and status sub-elements.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Cluster.CacheSecurityGroups")},
			{
				Name:        "log_delivery_configurations",
				Description: "Returns the destination, format, and type of the logs.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Cluster.LogDeliveryConfigurations")},
			{
				Name:        "tags_src",
				Description: "A list of tags associated with the cluster.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.TagList")},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Cluster.CacheClusterId")},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.TagList").Transform(clusterTagListToTurbotTags),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Cluster.ARN").Transform(arnToAkas),
			},
		}),
	}
}

//// TRANSFORM FUNCTIONS

func clusterTagListToTurbotTags(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	plugin.Logger(ctx).Trace("clusterTagListToTurbotTags")
	clusterTag := d.HydrateItem.(opengovernance.ElastiCacheCluster).Description

	if clusterTag.TagList == nil {
		return nil, nil
	}

	// Mapping the resource tags inside turbotTags
	var turbotTagsMap map[string]string
	if clusterTag.TagList != nil {
		turbotTagsMap = map[string]string{}
		for _, i := range clusterTag.TagList {
			turbotTagsMap[*i.Key] = *i.Value
		}
	}

	return turbotTagsMap, nil
}
