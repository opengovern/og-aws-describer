package aws

import (
	"context"

	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAwsEcsTaskSet(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_ecs_task_set",
		Description: "AWS ECS TaskSet",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    kaytu.GetECSTaskSet,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListECSTaskSet,
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "id",
				Description: "The id of the task set.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.TaskSet.Id")},
			{
				Name:        "task_set_arn",
				Description: "The Amazon Resource Name (ARN) of the task set",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.TaskSet.TaskSetArn")},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.TaskSet.Id")},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(getEcsTaskSetTurbotTags),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.TaskSet.TaskSetArn").Transform(arnToAkas),
			},
		}),
	}
}

func getEcsTaskSetTurbotTags(_ context.Context, d *transform.TransformData) (interface{}, error) {
	tags := d.HydrateItem.(kaytu.ECSTaskSet).Description.TaskSet.Tags
	return ecsV2TagsToMap(tags)
}
