package aws

import (
	"context"
	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAwsCodeStarProject(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_codestar_project",
		Description: "AWS CodeStar Project",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    kaytu.GetCodeStarProject,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListCodeStarProject,
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "id",
				Description: "The id of the project.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Project.Id")},
			{
				Name:        "name",
				Description: "The name of the project.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Project.Name")},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the project",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Project.Arn")},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Project.Name")},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Tags"),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Project.Arn").Transform(arnToAkas),
			},
		}),
	}
}
