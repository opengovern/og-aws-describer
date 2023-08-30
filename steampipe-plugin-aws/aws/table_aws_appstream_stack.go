package aws

import (
	"context"

	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAwsAppstreamStack(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_appstream_stack",
		Description: "AWS AppStream Stack",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("name"),
			Hydrate:    kaytu.GetAppStreamStack,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListAppStreamStack,
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the stack.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Stack.Name")},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the stack",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Stack.Arn")},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Stack.Name")},
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
				Transform:   transform.FromField("Description.Stack.Arn").Transform(arnToAkas),
			},
		}),
	}
}
