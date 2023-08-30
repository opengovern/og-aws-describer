package aws

import (
	"context"

	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAwsOpsWorksCMServer(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_opsworkscm_server",
		Description: "AWS OpsWorksCM Server",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("server_name"),
			Hydrate:    kaytu.GetOpsWorksCMServer,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListOpsWorksCMServer,
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "server_name",
				Description: "The name of the server.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Server.ServerName")},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the server",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Server.ServerArn")},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Server.ServerName")},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(getOpsWorksCMServerTurbotTags),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Server.ServerArn").Transform(arnToAkas),
			},
		}),
	}
}

//// TRANSFORM FUNCTIONS

func getOpsWorksCMServerTurbotTags(_ context.Context, d *transform.TransformData) (interface{}, error) {
	tags := d.HydrateItem.(kaytu.OpsWorksCMServer).Description.Tags
	return opsWorksCMV2TagsToMap(tags)
}
