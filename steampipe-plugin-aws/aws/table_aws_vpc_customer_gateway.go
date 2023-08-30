package aws

import (
	"context"

	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAwsVpcCustomerGateway(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_vpc_customer_gateway",
		Description: "AWS VPC Customer Gateway",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("customer_gateway_id"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"InvalidCustomerGatewayID.NotFound", "InvalidCustomerGatewayID.Malformed"}),
			},
			Hydrate: kaytu.GetEC2CustomerGateway,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListEC2CustomerGateway,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "ip_address", Require: plugin.Optional},
				{Name: "bgp_asn", Require: plugin.Optional},
				{Name: "state", Require: plugin.Optional},
				{Name: "type", Require: plugin.Optional},
			},
		},

		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "customer_gateway_id",
				Description: "The ID of the customer gateway.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.CustomerGateway.CustomerGatewayId")},
			{
				Name:        "type",
				Description: "The type of VPN connection the customer gateway supports (ipsec.1).",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.CustomerGateway.Type")},
			{
				Name:        "state",
				Description: "The current state of the customer gateway (pending | available | deleting | deleted).",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.CustomerGateway.State")},
			{
				Name:        "bgp_asn",
				Description: "The customer gateway's Border Gateway Protocol (BGP) Autonomous System Number (ASN).",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.CustomerGateway.BgpAsn")},
			{
				Name:        "certificate_arn",
				Description: "The Amazon Resource Name (ARN) for the customer gateway certificate.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.CustomerGateway.CertificateArn")},
			{
				Name:        "device_name",
				Description: "The name of customer gateway device.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.CustomerGateway.DeviceName")},
			{
				Name:        "ip_address",
				Description: "The Internet-routable IP address of the customer gateway's outside interface.",
				Type:        proto.ColumnType_IPADDR,
				Transform:   transform.FromField("Description.CustomerGateway.IpAddress")},
			{
				Name:        "tags_src",
				Description: "A list of tags that are attached to customer gateway.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.CustomerGateway.Tags")},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromP(getVpcCustomerGatewayTurbotData, "Tags"),
			},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromP(getVpcCustomerGatewayTurbotData, "Title"),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,

				Transform: transform.FromField("Description.CustomerGateway.Tags")},

			//// LIST FUNCTION
		}),
	}
}

//// TRANSFORM FUNCTIONS

func getVpcCustomerGatewayTurbotData(_ context.Context, d *transform.TransformData) (interface{}, error) {
	customerGateway := d.HydrateItem.(kaytu.EC2CustomerGateway).Description.CustomerGateway
	param := d.Param.(string)

	// Get resource title
	title := customerGateway.CustomerGatewayId

	// Get the resource tags
	var turbotTagsMap map[string]string
	if customerGateway.Tags != nil {
		turbotTagsMap = map[string]string{}
		for _, i := range customerGateway.Tags {
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
