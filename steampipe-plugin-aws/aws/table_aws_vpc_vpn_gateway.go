package aws

import (
	"context"
	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAwsVpcVpnGateway(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_vpc_vpn_gateway",
		Description: "AWS VPC VPN Gateway",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("vpn_gateway_id"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"InvalidVpnGatewayID.NotFound", "InvalidVpnGatewayID.Malformed"}),
			},
			Hydrate: opengovernance.GetEC2VPNGateway,
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListEC2VPNGateway,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "amazon_side_asn", Require: plugin.Optional},
				{Name: "availability_zone", Require: plugin.Optional},
				{Name: "state", Require: plugin.Optional},
				{Name: "type", Require: plugin.Optional},
			},
		},

		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "vpn_gateway_id",
				Description: "The ID of the virtual private gateway.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.VPNGateway.VpnGatewayId")},
			{
				Name:        "state",
				Description: "The current state of the virtual private gateway.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.VPNGateway.State")},
			{
				Name:        "type",
				Description: "The type of VPN connection the virtual private gateway supports.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.VPNGateway.Type")},
			{
				Name:        "amazon_side_asn",
				Description: "The private Autonomous System Number (ASN) for the Amazon side of a BGP session.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.VPNGateway.AmazonSideAsn")},
			{
				Name:        "availability_zone",
				Description: "The Availability Zone where the virtual private gateway was created, if applicable. This field may be empty or not returned.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.VPNGateway.AvailabilityZone")},
			{
				Name:        "vpc_attachments",
				Description: "Any VPCs attached to the virtual private gateway.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.VPNGateway.VpcAttachments")},
			{
				Name:        "tags_src",
				Description: "A list of tags that are attached to VPN gateway.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.VPNGateway.Tags")},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromP(getVpcVpnGatewayTurbotData, "Tags"),
			},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromP(getVpcVpnGatewayTurbotData, "Title"),
			},
			{
				//error we don't have this
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("ARN").Transform(arnToAkas),
			},
		}),
	}
}

//// TRANSFORM FUNCTIONS

func getVpcVpnGatewayTurbotData(_ context.Context, d *transform.TransformData) (interface{}, error) {
	vpnGateway := d.HydrateItem.(opengovernance.EC2VPNGateway).Description.VPNGateway
	param := d.Param.(string)

	// Get resource title
	title := vpnGateway.VpnGatewayId

	// Get the resource tags
	var turbotTagsMap map[string]string
	if vpnGateway.Tags != nil {
		turbotTagsMap = map[string]string{}
		for _, i := range vpnGateway.Tags {
			turbotTagsMap[*i.Key] = *i.Value
			if *i.Key == "Name" {
				title = i.Value
			}
		}
	}

	if param == "Tags" {
		return turbotTagsMap, nil
	}

	return title, nil
}
