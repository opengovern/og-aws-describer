package aws

import (
	"context"

	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsDaxCluster(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_dax_cluster",
		Description: "AWS DAX Cluster",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("cluster_name"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"ClusterNotFoundFault", "ServiceLinkedRoleNotFoundFault"}),
			},
			Hydrate: kaytu.GetDAXCluster,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListDAXCluster,
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:    "cluster_name",
					Require: plugin.Optional,
				},
			},
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "cluster_name",
				Description: "The name of the DAX cluster.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Cluster.ClusterName")},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) that uniquely identifies the cluster.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Cluster.ClusterArn")},
			{
				Name:        "description",
				Description: "The description of the cluster.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Cluster.Description")},
			{
				Name:        "active_nodes",
				Description: "The number of nodes in the cluster that are active.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Cluster.ActiveNodes")},
			{
				Name:        "iam_role_arn",
				Description: "A valid Amazon Resource Name (ARN) that identifies an IAM role.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Cluster.IamRoleArn")},
			{
				Name:        "node_type",
				Description: "The node type for the nodes in the cluster.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Cluster.NodeType")},
			{
				Name:        "preferred_maintenance_window",
				Description: "A range of time when maintenance of DAX cluster software will be performed.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Cluster.PreferredMaintenanceWindow")},
			{
				Name:        "status",
				Description: "The current status of the cluster.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Cluster.Status")},
			{
				Name:        "subnet_group",
				Description: "The subnet group where the DAX cluster is running.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Cluster.SubnetGroup")},
			{
				Name:        "total_nodes",
				Description: "The total number of nodes in the cluster.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Cluster.TotalNodes")},
			{
				Name:        "cluster_discovery_endpoint",
				Description: "The configuration endpoint for this DAX cluster, consisting of a DNS name and a port number.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Cluster.ClusterDiscoveryEndpoint")},
			{
				Name:        "node_ids_to_remove",
				Description: "A list of nodes to be removed from the cluster.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Cluster.NodeIdsToRemove")},
			{
				Name:        "nodes",
				Description: "A list of nodes that are currently in the cluster.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Cluster.Nodes")},
			{
				Name:        "notification_configuration",
				Description: "Describes a notification topic and its status.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Cluster.NotificationConfiguration")},
			{
				Name:        "parameter_group",
				Description: "The parameter group being used by nodes in the cluster.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Cluster.ParameterGroup")},
			{
				Name:        "sse_description",
				Description: "The description of the server-side encryption status on the specified DAX cluster.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Cluster.SSEDescription")},
			{
				Name:        "security_groups",
				Description: "A list of security groups, and the status of each, for the nodes in the cluster.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Cluster.SecurityGroups")},
			{
				Name:        "tags_src",
				Description: "A list of tags currently associated with the DAX cluster.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Tags"),
			},

			// Standard columns
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Cluster.ClusterName"),
			},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(daxClusterTurbotData),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Cluster.ClusterArn").Transform(arnToAkas),
			},
		}),
	}
}

//// LIST FUNCTION

//// HYDRATE FUNCTIONS

//// TRANSFORM FUNCTION

func daxClusterTurbotData(_ context.Context, d *transform.TransformData) (interface{}, error) {
	data := d.HydrateItem.(kaytu.DAXCluster).Description
	if data.Tags == nil {
		return nil, nil
	}

	// Mapping the resource tags inside turbotTags
	if data.Tags != nil {
		turbotTagsMap := map[string]string{}
		for _, i := range data.Tags {
			turbotTagsMap[*i.Key] = *i.Value
		}
		return turbotTagsMap, nil
	}
	return nil, nil
}
