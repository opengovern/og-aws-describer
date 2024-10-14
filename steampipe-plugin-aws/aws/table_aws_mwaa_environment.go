package aws

import (
	"context"

	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAwsMWAAEnvironment(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_mwaa_environment",
		Description: "AWS MWAA Environment",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("name"),
			Hydrate:    opengovernance.GetMWAAEnvironment,
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListMWAAEnvironment,
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the environment.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Environment.Name")},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the environment",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Environment.Arn")},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Environment.Name")},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Environment.Tags"),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Environment.Arn").Transform(arnToAkas),
			},
		}),
	}
}
