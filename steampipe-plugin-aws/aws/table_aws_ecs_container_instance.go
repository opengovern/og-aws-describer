package aws

import (
	"context"

	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAwsEcsContainerInstance(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_ecs_container_instance",
		Description: "AWS ECS Container Instance",
		List: &plugin.ListConfig{
			ParentHydrate: kaytu.ListECSCluster,
			Hydrate:       kaytu.ListECSContainerInstance,
		},
		Columns: awsKaytuColumns([]*plugin.Column{
			{
				Name:        "arn",
				Description: "The namespace Amazon Resource Name (ARN) of the container instance.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ContainerInstance.ContainerInstanceArn"),
			},
			{
				Name:        "ec2_instance_id",
				Description: "The EC2 instance ID of the container instance.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ContainerInstance.Ec2InstanceId")},
			{
				Name:        "cluster_arn",
				Description: "The ARN of the cluster.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Cluster.ClusterArn")},
			{
				Name:        "agent_connected",
				Description: "True if the agent is connected to Amazon ECS.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.ContainerInstance.AgentConnected")},
			{
				Name:        "agent_update_status",
				Description: "The status of the most recent agent update.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ContainerInstance.AgentUpdateStatus")},
			{
				Name:        "attachments",
				Description: "The resources attached to a container instance, such as elastic network interfaces.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ContainerInstance.Attachments")},
			{
				Name:        "attributes",
				Description: "The attributes set for the container instance.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ContainerInstance.Attributes")},
			{
				Name:        "capacity_provider_name",
				Description: "The capacity provider associated with the container instance.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ContainerInstance.CapacityProviderName")},
			{
				Name:        "pending_tasks_count",
				Description: "The number of tasks on the container instance that are in the PENDING status.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.ContainerInstance.PendingTasksCount")},
			{
				Name:        "registered_at",
				Description: "The Unix timestamp for when the container instance was registered.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.ContainerInstance.RegisteredAt")},
			{
				Name:        "registered_resources",
				Description: "CPU and memory that can be allocated on this container instance to tasks.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ContainerInstance.RegisteredResources")},
			{
				Name:        "remaining_resources",
				Description: "CPU and memory that is available for new tasks.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ContainerInstance.RemainingResources")},
			{
				Name:        "running_tasks_count",
				Description: "CPU and memory that is available for new tasks.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.ContainerInstance.RunningTasksCount")},
			{
				Name:        "status",
				Description: "The status of the container instance.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ContainerInstance.Status")},
			{
				Name:        "status_reason",
				Description: "The reason that the container instance reached its current status.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ContainerInstance.StatusReason")},
			{
				Name:        "version",
				Description: "The reason that the container instance reached its current status.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.ContainerInstance.Version")},
			{
				Name:        "version_info",
				Description: "Version information for the Amazon ECS container agent and Docker daemon running on the container instance.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ContainerInstance.VersionInfo")},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(containerInstanceTurbotTags),
			},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ContainerInstance.ContainerInstanceArn")},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ContainerInstance.ContainerInstanceArn").Transform(arnToAkas),
			},
		}),
	}
}

//// TRANSFORM FUNCTIONS

func containerInstanceTurbotTags(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	plugin.Logger(ctx).Trace("containerInstanceTagsToTurbotTags")
	tags := d.HydrateItem.(kaytu.ECSContainerInstance).Description.ContainerInstance.Tags

	// Mapping the resource tags inside turbotTags
	var turbotTagsMap map[string]string
	if tags != nil {
		turbotTagsMap = map[string]string{}
		for _, i := range tags {
			turbotTagsMap[*i.Key] = *i.Value
		}
	}

	return turbotTagsMap, nil
}
