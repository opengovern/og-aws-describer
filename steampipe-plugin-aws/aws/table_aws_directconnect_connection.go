package aws

import (
	"context"
	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAwsDirectConnectConnection(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_directconnect_connection",
		Description: "AWS DirectConnect Connection",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("connection_id"),
			Hydrate:    opengovernance.GetDirectConnectConnection,
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListDirectConnectConnection,
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "connection_id",
				Description: "The id of the connection.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Connection.ConnectionId")},
			{
				Name:        "name",
				Description: "The name of the connection.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Connection.ConnectionName")},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the connection",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("ARN")},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Connection.ConnectionName")},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(getDirectConnectConnectionTurbotTags),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("ARN").Transform(arnToAkas),
			},
		}),
	}
}

//// TRANSFORM FUNCTIONS

func getDirectConnectConnectionTurbotTags(_ context.Context, d *transform.TransformData) (interface{}, error) {
	tags := d.HydrateItem.(opengovernance.DirectConnectConnection).Description.Connection.Tags
	return directConnectV2TagsToMap(tags)
}
