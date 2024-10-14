package aws

import (
	"context"

	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAwsFsxTask(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_fsx_task",
		Description: "AWS FSX Task",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("task_id"),
			Hydrate:    opengovernance.GetFSXTask,
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListFSXTask,
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "task_id",
				Description: "The id of the task.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Task.TaskId")},
			{
				Name:        "resource_arn",
				Description: "The Amazon Resource Name (ARN) of the task",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Task.ResourceARN")},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Task.TaskId")},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(getFsxTaskTurbotTags),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Task.ResourceARN").Transform(arnToAkas),
			},
		}),
	}
}

//// TRANSFORM FUNCTIONS

func getFsxTaskTurbotTags(_ context.Context, d *transform.TransformData) (interface{}, error) {
	tags := d.HydrateItem.(opengovernance.FSXTask).Description.Task.Tags
	return fsxV2TagsToMap(tags)
}
