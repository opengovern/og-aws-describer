package aws

import (
	"context"
	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"
	"strings"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsEcsTaskDefinition(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_ecs_task_definition",
		Description: "AWS ECS Task Definition",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("task_definition_arn"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"InvalidParameterException", "ClientException"}),
			},
			Hydrate: kaytu.GetECSTaskDefinition,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListECSTaskDefinition,
		},

		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "task_definition_arn",
				Description: "The Amazon Resource Name (ARN) that identifies the task definition.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.TaskDefinition.TaskDefinitionArn")},
			{
				Name:        "cpu",
				Description: "The number of cpu units used by the task.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.TaskDefinition.Cpu"),
			},
			{
				Name:        "status",
				Description: "The status of the task definition.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.TaskDefinition.Status")},
			{
				Name:        "execution_role_arn",
				Description: "The Amazon Resource Name (ARN) of the task execution role that grants the Amazon ECS container agent permission to make AWS API calls on your behalf.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.TaskDefinition.ExecutionRoleArn")},
			{
				Name:        "family",
				Description: "The name of a family that this task definition is registered to.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.TaskDefinition.Family")},
			{
				Name:        "ipc_mode",
				Description: "The IPC resource namespace to use for the containers in the task.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.TaskDefinition.IpcMode")},
			{
				Name:        "memory",
				Description: "The amount (in MiB) of memory used by the task.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.TaskDefinition.Memory")},
			{
				Name:        "network_mode",
				Description: "The Docker networking mode to use for the containers in the task.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.TaskDefinition.NetworkMode")},
			{
				Name:        "pid_mode",
				Description: "The process namespace to use for the containers in the task.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.TaskDefinition.PidMode")},
			{
				Name:        "revision",
				Description: "The revision of the task in a particular family.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.TaskDefinition.Revision")},
			{
				Name:        "task_role_arn",
				Description: "The short name or full Amazon Resource Name (ARN) of the AWS Identity and Access Management (IAM) role that grants containers in the task permission to call AWS APIs on your behalf.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.TaskDefinition.TaskRoleArn")},
			{
				Name:        "registered_at",
				Description: "The Unix timestamp for when the task definition was registered.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.TaskDefinition.RegisteredAt")},
			{
				Name:        "registered_by",
				Description: "The principal that registered the task definition.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.TaskDefinition.RegisteredBy")},
			{
				Name:        "container_definitions",
				Description: "A list of container definitions in JSON format that describe the different containers that make up your task.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.TaskDefinition.ContainerDefinitions")},
			{
				Name:        "compatibilities",
				Description: "The launch type to use with your task.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.TaskDefinition.Compatibilities")},
			{
				Name:        "inference_accelerators",
				Description: "The Elastic Inference accelerator associated with the task.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.TaskDefinition.InferenceAccelerators")},
			{
				Name:        "placement_constraints",
				Description: "An array of placement constraint objects to use for tasks.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.TaskDefinition.PlacementConstraints")},
			{
				Name:        "proxy_configuration",
				Description: "The configuration details for the App Mesh proxy.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.TaskDefinition.ProxyConfiguration")},
			{
				Name:        "requires_attributes",
				Description: "The container instance attributes required by your task.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.TaskDefinition.RequiresAttributes")},
			{
				Name:        "requires_compatibilities",
				Description: "The launch type the task requires. If no value is specified, it will default to EC2. Valid values include EC2 and FARGATE.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.TaskDefinition.RequiresCompatibilities")},
			{
				Name:        "volumes",
				Description: "The list of volume definitions for the task.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.TaskDefinition.Volumes")},
			{
				Name:        "tags_src",
				Description: "A list of tags associated with task.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Tags")},
			// Standard columns

			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromP(getAwsEcsTaskDefinitionTurbotData, "Title"),
			},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromP(getAwsEcsTaskDefinitionTurbotData, "Tags"),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.TaskDefinition.TaskDefinitionArn").Transform(arnToAkas),
			},
		}),
	}
}

//// LIST FUNCTION

//// HYDRATE FUNCTIONS

//// TRANSFORM FUNCTIONS

func getAwsEcsTaskDefinitionTurbotData(_ context.Context, d *transform.TransformData) (interface{},
	error) {
	param := d.Param.(string)
	ecsTaskDefinition := d.HydrateItem.(kaytu.ECSTaskDefinition).Description

	// Get resource title
	arn := ecsTaskDefinition.TaskDefinition.TaskDefinitionArn
	splitArn := strings.Split(*arn, "/")
	title := splitArn[len(splitArn)-1]

	if param == "Tags" {
		if ecsTaskDefinition.Tags == nil {
			return nil, nil
		}

		// Get the resource tags
		if ecsTaskDefinition.Tags != nil {
			turbotTagsMap := map[string]string{}
			for _, i := range ecsTaskDefinition.Tags {
				turbotTagsMap[*i.Key] = *i.Value
			}
			return turbotTagsMap, nil
		}
	}

	return title, nil
}
