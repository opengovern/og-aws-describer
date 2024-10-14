package aws

import (
	"context"
	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"

	"github.com/aws/aws-sdk-go-v2/service/ecs/types"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsEcsCluster(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_ecs_cluster",
		Description: "AWS ECS Cluster",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("cluster_arn"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"ResourceNotFoundException", "InvalidParameterException"}),
			},
			Hydrate: opengovernance.GetECSCluster,
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListECSCluster,
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "cluster_arn",
				Description: "The Amazon Resource Name (ARN) that identifies the cluster.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Cluster.ClusterArn"),
			},
			{
				Name:        "cluster_name",
				Description: "A user-generated string that you use to identify your cluster.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Cluster.ClusterName"),
			},
			{
				Name:        "active_services_count",
				Description: "The number of services that are running on the cluster in an ACTIVE state.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Cluster.ActiveServicesCount"),
			},
			{
				Name:        "pending_tasks_count",
				Description: "The number of tasks in the cluster that are in the PENDING state.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Cluster.PendingTasksCount"),
			},
			{
				Name:        "registered_container_instances_count",
				Description: "The number of container instances registered into the cluster. This includes container instances in both ACTIVE and DRAINING status.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Cluster.RegisteredContainerInstancesCount"),
			},
			{
				Name:        "running_tasks_count",
				Description: "The number of tasks in the cluster that are in the RUNNING state.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Cluster.RunningTasksCount"),
			},
			{
				Name:        "status",
				Description: "The status of the cluster.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Cluster.Status"),
			},
			{
				Name:        "attachments_status",
				Description: "The status of the capacity providers associated with the cluster.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Cluster.AttachmentsStatus"),
			},
			{
				Name:        "attachments",
				Description: "The resources attached to a cluster. When using a capacity provider with a cluster, the Auto Scaling plan that is created will be returned as a cluster attachment.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Cluster.Attachments"),
			},
			{
				Name:        "capacity_providers",
				Description: "The capacity providers associated with the cluster.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Cluster.CapacityProviders"),
			},
			{
				Name:        "default_capacity_provider_strategy",
				Description: "The default capacity provider strategy for the cluster.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Cluster.DefaultCapacityProviderStrategy"),
			},
			{
				Name:        "settings",
				Description: "The settings for the cluster. This parameter indicates whether CloudWatch Container Insights is enabled or disabled for a cluster.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Cluster.Settings"),
			},
			{
				Name:        "statistics",
				Description: "Additional information about your clusters that are separated by launch type.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Cluster.Statistics"),
			},
			{
				Name:        "tags_src",
				Description: "A list of tags attached to the cluster.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Cluster.Tags"),
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
				Transform:   transform.FromField("Description.Cluster.Tags").Transform(getAwsEcsClusterTurbotTags),
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

//// TRANSFORM FUNCTIONS

func getAwsEcsClusterTurbotTags(_ context.Context, d *transform.TransformData) (interface{}, error) {
	tags := d.Value.([]types.Tag)

	if tags != nil {
		turbotTagsMap := map[string]string{}
		for _, i := range tags {
			turbotTagsMap[*i.Key] = *i.Value
		}
		return turbotTagsMap, nil
	}

	return nil, nil
}
