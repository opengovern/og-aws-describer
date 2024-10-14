package aws

import (
	"context"

	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsEc2TransitGatewayRoute(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_ec2_transit_gateway_route",
		Description: "AWS EC2 Transit Gateway Route",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListEC2TransitGatewayRoute,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "prefix_list_id", Require: plugin.Optional},
				{Name: "state", Require: plugin.Optional},
				{Name: "type", Require: plugin.Optional},
			},
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"InvalidAction"}),
			},
		},

		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "transit_gateway_route_table_id",
				Description: "The ID of the transit gateway route table.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.TransitGatewayRouteTableId")},
			{
				Name:        "destination_cidr_block",
				Description: "The CIDR block used for destination matches.",
				Type:        proto.ColumnType_CIDR,

				Transform: transform.FromField("Description.TransitGatewayRoute.DestinationCidrBlock")},
			{
				Name:        "prefix_list_id",
				Description: "The ID of the prefix list used for destination matches.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.TransitGatewayRoute.PrefixListId")},
			{
				Name:        "state",
				Description: "The state of the route.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.TransitGatewayRoute.State")},
			{
				Name:        "type",
				Description: "The route type.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.TransitGatewayRoute.Type")},
			{
				Name:        "transit_gateway_attachments",
				Description: "The attachments.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.TransitGatewayRoute.TransitGatewayAttachments")},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.TransitGatewayRoute.DestinationCidrBlock")},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("ARN").Transform(arnToAkas)},
		}),
	}
}
