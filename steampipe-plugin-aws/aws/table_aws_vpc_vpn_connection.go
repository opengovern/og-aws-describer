package aws

import (
	"context"
	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsVpcVpnConnection(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_vpc_vpn_connection",
		Description: "AWS VPC VPN Connection",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("vpn_connection_id"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"InvalidVpnConnectionID.NotFound", "InvalidVpnConnectionID.Malformed"}),
			},
			Hydrate: opengovernance.GetEC2VPNConnection,
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListEC2VPNConnection,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"InvalidParameterValue"}),
			},
			KeyColumns: []*plugin.KeyColumn{
				{Name: "customer_gateway_configuration", Require: plugin.Optional},
				{Name: "customer_gateway_id", Require: plugin.Optional},
				{Name: "state", Require: plugin.Optional},
				{Name: "type", Require: plugin.Optional},
				{Name: "vpn_gateway_id", Require: plugin.Optional},
				{Name: "transit_gateway_id", Require: plugin.Optional},
			},
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "vpn_connection_id",
				Description: "The ID of the VPN connection.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.VpnConnection.VpnConnectionId"),
			},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) specifying the VPN connection.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("ARN"),
			},
			{
				Name:        "state",
				Description: "The current state of the VPN connection.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.VpnConnection.State"),
			},
			{
				Name:        "type",
				Description: "The type of VPN connection.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.VpnConnection.Type"),
			},
			{
				Name:        "category",
				Description: "The category of the VPN connection. A value of VPN indicates an AWS VPN connection.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.VpnConnection.Category"),
			},
			{
				Name:        "vpn_gateway_id",
				Description: "The ID of the virtual private gateway at the AWS side of the VPN connection.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.VpnConnection.VpnGatewayId"),
			},
			{
				Name:        "customer_gateway_id",
				Description: "The ID of the customer gateway at your end of the VPN connection.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.VpnConnection.CustomerGatewayId")},
			{
				Name:        "customer_gateway_configuration",
				Description: "The configuration information for the VPN connection's customer gateway.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.VpnConnection.CustomerGatewayConfiguration"),
			},
			{
				Name:        "transit_gateway_id",
				Description: "The ID of the transit gateway associated with the VPN connection.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.VpnConnection.TransitGatewayId"),
			},
			{
				Name:        "options",
				Description: "The VPN connection options.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.VpnConnection.Options"),
			},
			{
				Name:        "routes",
				Description: "The static routes associated with the VPN connection.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.VpnConnection.Routes"),
			},
			{
				Name:        "vgw_telemetry",
				Description: "Information about the VPN tunnel.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.VpnConnection.VgwTelemetry"),
			},
			{
				Name:        "tags_src",
				Description: "A list of tags that are attached to VPN gateway.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.VpnConnection.Tags"),
			},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromP(vpnConnectionTurbotData, "Title"),
			},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromP(vpnConnectionTurbotData, "Tags"),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("ARN").Transform(transform.EnsureStringArray),
			},
		}),
	}
}

//// TRANSFORM FUNCTIONS

func vpnConnectionTurbotData(_ context.Context, d *transform.TransformData) (interface{}, error) {
	vpnConnection := d.HydrateItem.(opengovernance.EC2VPNConnection).Description.VpnConnection
	param := d.Param.(string)

	// Get resource title
	title := vpnConnection.VpnConnectionId

	// Get the resource tags
	var turbotTagsMap map[string]string
	if vpnConnection.Tags != nil {
		turbotTagsMap = map[string]string{}
		for _, i := range vpnConnection.Tags {
			turbotTagsMap[*i.Key] = *i.Value
			if *i.Key == "Name" {
				title = i.Value
			}
		}
	}

	if param == "Tags" {
		if vpnConnection.Tags == nil {
			return nil, nil
		}
		return turbotTagsMap, nil
	}

	return title, nil
}
