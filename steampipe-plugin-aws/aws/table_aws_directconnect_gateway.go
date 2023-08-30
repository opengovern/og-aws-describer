package aws

import (
	"context"
	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAwsDirectConnectGateway(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_directconnect_gateway",
		Description: "AWS DirectConnect Gateway",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("direct_connect_gateway_id"),
			Hydrate:    kaytu.GetDirectConnectGateway,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListDirectConnectGateway,
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "direct_connect_gateway_id",
				Description: "The id of the gateway.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Gateway.DirectConnectGatewayId")},
			{
				Name:        "name",
				Description: "The name of the gateway.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Gateway.DirectConnectGatewayName")},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the gateway",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("ARN")},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Gateway.DirectConnectGatewayName")},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(getDirectConnectGatewayTurbotTags),
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

func getDirectConnectGatewayTurbotTags(_ context.Context, d *transform.TransformData) (interface{}, error) {
	tags := d.HydrateItem.(kaytu.DirectConnectGateway).Description.Tags
	return directConnectV2TagsToMap(tags)
}
