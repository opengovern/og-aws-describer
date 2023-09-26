package aws

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/ecs/types"
	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsEcsService(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_ecs_service",
		Description: "AWS ECS Service",
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListECSService,
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "service_name",
				Description: "The name of the service.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Service.ServiceName"),
			},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) specifying the service.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Service.ServiceArn"),
			},
			{
				Name:        "status",
				Description: "The status of the service. Valid values are: ACTIVE, DRAINING, or INACTIVE.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Service.Status"),
			},
			{
				Name:        "cluster_arn",
				Description: "The Amazon Resource Name (ARN) of the cluster that hosts the service.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Service.ClusterArn"),
			},
			{
				Name:        "task_definition",
				Description: "The task definition to use for tasks in the service.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Service.TaskDefinition"),
			},
			{
				Name:        "created_at",
				Description: "The date and time when the service was created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Service.CreatedAt"),
			},
			{
				Name:        "created_by",
				Description: "The principal that created the service.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Service.CreatedBy"),
			},
			{
				Name:        "deployment_controller_type",
				Description: "The deployment controller type to use. Possible values are: ECS, CODE_DEPLOY, and EXTERNAL.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Service.DeploymentController.Type"),
			},
			{
				Name:        "desired_count",
				Description: "The desired number of instantiations of the task definition to keep running on the service.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Service.DesiredCount"),
			},
			{
				Name:        "enable_ecs_managed_tags",
				Description: "Specifies whether to enable Amazon ECS managed tags for the tasks in the service.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Service.EnableECSManagedTags"),
			},
			{
				Name:        "enable_execute_command",
				Description: "Indicates whether or not the execute command functionality is enabled for the service.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Service.EnableExecuteCommand"),
			},
			{
				Name:        "health_check_grace_period_seconds",
				Description: "The period of time, in seconds, that the Amazon ECS service scheduler ignores unhealthy Elastic Load Balancing target health checks after a task has first started.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Service.HealthCheckGracePeriodSeconds"),
			},
			{
				Name:        "launch_type",
				Description: "The launch type on which your service is running. If no value is specified, it will default to EC2.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Service.LaunchType"),
			},
			{
				Name:        "pending_count",
				Description: "The number of tasks in the cluster that are in the PENDING state.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Service.PendingCount"),
			},
			{
				Name:        "platform_family",
				Description: "The operating system that your tasks in the service run on.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Service.PlatformFamily"),
			},
			{
				Name:        "platform_version",
				Description: "The platform version on which to run your service.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Service.PlatformVersion"),
			},
			{
				Name:        "propagate_tags",
				Description: "Specifies whether to propagate the tags from the task definition or the service to the task. If no value is specified, the tags are not propagated.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Service.PropagateTags"),
			},
			{
				Name:        "role_arn",
				Description: "The ARN of the IAM role associated with the service that allows the Amazon ECS container agent to register container instances with an Elastic Load Balancing load balancer.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Service.RoleArn"),
			},
			{
				Name:        "running_count",
				Description: "The number of tasks in the cluster that are in the RUNNING state.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Service.RunningCount"),
			},
			{
				Name:        "scheduling_strategy",
				Description: "The scheduling strategy to use for the service.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Service.SchedulingStrategy"),
			},
			{
				Name:        "capacity_provider_strategy",
				Description: "The capacity provider strategy associated with the service.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Service.CapacityProviderStrategy"),
			},
			{
				Name:        "deployment_configuration",
				Description: "Optional deployment parameters that control how many tasks run during the deployment and the ordering of stopping and starting tasks.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Service.DeploymentConfiguration"),
			},
			{
				Name:        "deployments",
				Description: "The current state of deployments for the service.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Service.Deployments"),
			},
			{
				Name:        "events",
				Description: "The event stream for your service. A maximum of 100 of the latest events are displayed.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Service.Events"),
			},
			{
				Name:        "load_balancers",
				Description: "A list of Elastic Load Balancing load balancer objects, containing the load balancer name, the container name (as it appears in a container definition), and the container port to access from the load balancer.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Service.LoadBalancers"),
			},
			{
				Name:        "network_configuration",
				Description: "The VPC subnet and security group configuration for tasks that receive their own elastic network interface by using the awsvpc networking mode.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Service.NetworkConfiguration"),
			},
			{
				Name:        "placement_constraints",
				Description: "The placement constraints for the tasks in the service.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Service.PlacementConstraints"),
			},
			{
				Name:        "placement_strategy",
				Description: "The placement strategy that determines how tasks for the service are placed.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Service.PlacementStrategy"),
			},
			{
				Name:        "service_registries",
				Description: "The details of the service discovery registries to assign to this service.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Service.ServiceRegistries"),
			},
			{
				Name:        "task_sets",
				Description: "Information about a set of Amazon ECS tasks in either an AWS CodeDeploy or an EXTERNAL deployment.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Service.TaskSets"),
			},
			{
				Name:        "tags_src",
				Description: "The metadata that you apply to the service to help you categorize and organize them.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Tags"),
			},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Service.ServiceName"),
			},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Tags").Transform(getEcsServiceTurbotTags),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Service.ServiceArn").Transform(transform.EnsureStringArray),
			},
		}),
	}
}

//// LIST FUNCTION

//// TRANSFORM FUNCTIONS

func getEcsServiceTurbotTags(_ context.Context, d *transform.TransformData) (interface{},
	error) {
	tags := d.HydrateItem.([]types.Tag)

	if len(tags) == 0 {
		return nil, nil
	}

	turbotTagsMap := map[string]string{}
	for _, i := range tags {
		turbotTagsMap[*i.Key] = *i.Value
	}
	return turbotTagsMap, nil
}
