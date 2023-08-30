package aws

import (
	"context"
	"strings"

	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAwsEcsTask(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_ecs_task",
		Description: "AWS ECS Task",
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListECSTask,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"ClusterNotFoundException", "ServiceNotFoundException", "InvalidParameterException"}),
			},
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:    "container_instance_arn",
					Require: plugin.Optional,
				},
				{
					Name:    "desired_status",
					Require: plugin.Optional,
				},
				{
					Name:    "launch_type",
					Require: plugin.Optional,
				},
				{
					Name:       "service_name",
					Require:    plugin.Optional,
					CacheMatch: "exact",
				},
			},
		},

		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "task_arn",
				Description: "The Amazon Resource Name (ARN) of the task.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Task.TaskArn")},
			{
				Name:        "container_instance_arn",
				Description: "The ARN of the container instances that host the task.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Task.ContainerInstanceArn")},
			{
				Name:        "cluster_name",
				Description: "A user-generated string that you use to identify your cluster.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.From(extractClusterName),
			},
			{
				Name:        "desired_status",
				Description: "The desired status of the task.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Task.DesiredStatus")},
			{
				Name:        "launch_type",
				Description: "The infrastructure on which your task is running.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Task.LaunchType")},

			{
				Name:        "availability_zone",
				Description: "The availability zone of the task.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Task.AvailabilityZone")},
			{
				Name:        "capacity_provider_name",
				Description: "The capacity provider associated with the task.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Task.CapacityProviderName")},
			{
				Name:        "cluster_arn",
				Description: "The ARN of the cluster that hosts the task.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Task.ClusterArn")},
			{
				Name:        "connectivity",
				Description: "The connectivity status of a task.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Task.Connectivity")},
			{
				Name:        "connectivity_at",
				Description: "The Unix timestamp for when the task last went into CONNECTED status.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Task.ConnectivityAt")},
			{
				Name:        "cpu",
				Description: "The number of CPU units used by the task as expressed in a task definition.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.NO_MATCH_WAS_FOUND")},
			{
				Name:        "created_at",
				Description: "The Unix timestamp for when the task was created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Task.CreatedAt")},
			{
				Name:        "enable_execute_command",
				Description: "Whether or not execute command functionality is enabled for this task. If true, this enables execute command functionality on all containers in the task.",
				Type:        proto.ColumnType_BOOL,
				Default:     false,
				Transform:   transform.FromField("Description.Task.EnableExecuteCommand")},
			{
				Name:        "execution_stopped_at",
				Description: "The Unix timestamp for when the task execution stopped.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Task.ExecutionStoppedAt")},
			{
				Name:        "group",
				Description: "The name of the task group associated with the task.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Task.Group")},
			{
				Name:        "health_status",
				Description: "The health status for the task, which is determined by the health of the essential containers in the task. If all essential containers in the task are reporting as HEALTHY, then the task status also reports as HEALTHY.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Task.HealthStatus")},
			{
				Name:        "last_status",
				Description: "The last known status of the task.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Task.LastStatus")},
			{
				Name:        "memory",
				Description: "The amount of memory (in MiB) used by the task as expressed in a task definition.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Task.Memory")},
			{
				Name:        "platform_version",
				Description: "The platform version on which your task is running.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Task.PlatformVersion")},
			{
				Name:        "pull_started_at",
				Description: "The Unix timestamp for when the container image pull began.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Task.PullStartedAt")},
			{
				Name:        "pull_stopped_at",
				Description: "The Unix timestamp for when the container image pull completed.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Task.PullStoppedAt")},
			{
				Name:        "service_name",
				Description: "The name of the service.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromQual("service_name"),
			},
			{
				Name:        "started_at",
				Description: "The Unix timestamp for when the task started.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Task.StartedAt")},
			{
				Name:        "started_by",
				Description: "The tag specified when a task is started.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Task.StartedBy")},
			{
				Name:        "stop_code",
				Description: "The stop code indicating why a task was stopped.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Task.StopCode")},
			{
				Name:        "stopped_at",
				Description: "The Unix timestamp for when the task was stopped.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Task.StoppedAt")},
			{
				Name:        "stopped_reason",
				Description: "The reason that the task was stopped.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Task.StoppedReason")},
			{
				Name:        "stopping_at",
				Description: "The Unix timestamp for when the task stops.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Task.StoppingAt")},
			{
				Name:        "task_definition_arn",
				Description: "The ARN of the task definition that creates the task.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Task.TaskDefinitionArn")},
			{
				Name:        "version",
				Description: "The version counter for the task.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Task.Version")},
			{
				Name:        "attachments",
				Description: "The Elastic Network Adapter associated with the task if the task uses the awsvpc network mode.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Task.Attachments")},
			{
				Name:        "attributes",
				Description: "The attributes of the task.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Task.Attributes")},
			{
				Name:        "containers",
				Description: "The containers associated with the task.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Task.Containers")},
			{
				Name:        "ephemeral_storage",
				Description: "The ephemeral storage settings for the task.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Task.EphemeralStorage")},
			{
				Name:        "inference_accelerators",
				Description: "The Elastic Inference accelerator associated with the task.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Task.InferenceAccelerators")},
			{
				Name:        "overrides",
				Description: "One or more container overrides.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Task.Overrides")},
			{
				Name:        "protection",
				Description: "Protection status of task in an Amazon ECS service.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.TaskProtection")},
			{
				Name:        "tags_src",
				Description: "A list of tags associated with task.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Task.Tags")},

			// Standard columns for all tables
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(ecsTaskTags),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Task.TaskArn").Transform(transform.EnsureStringArray),
			},
		}),
	}
}

//// TRANSFORM FUNCTIONS

func extractClusterName(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	task := d.HydrateItem.(kaytu.ECSTask).Description.Task
	clusterName := strings.Split(string(*task.ClusterArn), "/")[1]

	return clusterName, nil
}

func ecsTaskTags(_ context.Context, d *transform.TransformData) (interface{}, error) {
	task := d.HydrateItem.(kaytu.ECSTask).Description.Task

	var turbotTagsMap map[string]string
	if len(task.Tags) > 0 {
		turbotTagsMap = map[string]string{}
		for _, i := range task.Tags {
			turbotTagsMap[*i.Key] = *i.Value
		}
	}

	return turbotTagsMap, nil
}
