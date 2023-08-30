package aws

import (
	"context"
	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsEc2TransitGateway(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_ec2_transit_gateway",
		Description: "AWS EC2 Transit Gateway",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("transit_gateway_id"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"InvalidTransitGatewayID.NotFound", "InvalidTransitGatewayID.Unavailable", "InvalidTransitGatewayID.Malformed"}),
			},
			Hydrate: kaytu.GetEC2TransitGateway,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListEC2TransitGateway,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "propagation_default_route_table_id", Require: plugin.Optional},
				{Name: "amazon_side_asn", Require: plugin.Optional},
				{Name: "association_default_route_table_id", Require: plugin.Optional},
				{Name: "auto_accept_shared_attachments", Require: plugin.Optional},
				{Name: "default_route_table_association", Require: plugin.Optional},
				{Name: "default_route_table_propagation", Require: plugin.Optional},
				{Name: "dns_support", Require: plugin.Optional},
				{Name: "vpn_ecmp_support", Require: plugin.Optional},
				{Name: "owner_id", Require: plugin.Optional},
				{Name: "state", Require: plugin.Optional},
			},
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"InvalidAction"}),
			},
		},

		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "transit_gateway_id",
				Description: "The ID of the transit gateway.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.TransitGateway.TransitGatewayId")},
			{
				Name:        "transit_gateway_arn",
				Description: "The Amazon Resource Name (ARN) of the transit gateway.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.TransitGateway.TransitGatewayArn")},
			{
				Name:        "state",
				Description: "The state of the transit gateway.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.TransitGateway.State")},
			{
				Name:        "owner_id",
				Description: "The ID of the AWS account ID that owns the transit gateway.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.TransitGateway.OwnerId")},
			{
				Name:        "description",
				Description: "The description of the transit gateway.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.TransitGateway.Description")},
			{
				Name:        "creation_time",
				Description: "The date and time when transit gateway was created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.TransitGateway.CreationTime")},
			{
				Name:        "amazon_side_asn",
				Description: "A private Autonomous System Number (ASN) for the Amazon side of a BGP session. The range is 64512 to 65534 for 16-bit ASNs and 4200000000 to 4294967294 for 32-bit ASNs.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.TransitGateway.Options.AmazonSideAsn")},
			{
				Name:        "association_default_route_table_id",
				Description: "The ID of the default association route table.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.TransitGateway.Options.AssociationDefaultRouteTableId")},
			{
				Name:        "auto_accept_shared_attachments",
				Description: "Indicates whether attachment requests are automatically accepted.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.TransitGateway.Options.AutoAcceptSharedAttachments")},
			{
				Name:        "default_route_table_association",
				Description: "Indicates whether resource attachments are automatically associated with the default association route table.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.TransitGateway.Options.DefaultRouteTableAssociation")},
			{
				Name:        "default_route_table_propagation",
				Description: "Indicates whether resource attachments are automatically associated with the default association route table.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.TransitGateway.Options.DefaultRouteTablePropagation")},
			{
				Name:        "dns_support",
				Description: "Indicates whether DNS support is enabled.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.TransitGateway.Options.DnsSupport")},
			{
				Name:        "multicast_support",
				Description: "Indicates whether multicast is enabled on the transit gateway.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.TransitGateway.Options.MulticastSupport")},
			{
				Name:        "propagation_default_route_table_id",
				Description: "The ID of the default propagation route table.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.TransitGateway.Options.PropagationDefaultRouteTableId")},
			{
				Name:        "vpn_ecmp_support",
				Description: "Indicates whether Equal Cost Multipath Protocol support is enabled.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.TransitGateway.Options.VpnEcmpSupport")},
			{
				Name:        "cidr_blocks",
				Description: "A list of transit gateway CIDR blocks.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.TransitGateway.Options.TransitGatewayCidrBlocks")},
			{
				Name:        "tags_src",
				Description: "A list of tags that are assigned to the transit gateway.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.TransitGateway.Tags")},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(getEc2TransitGatewayTurbotTags),
			},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.From(getEc2TransitGatewayTurbotTitle),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.TransitGateway.TransitGatewayArn").Transform(arnToAkas),
			},
		}),
	}
}

// Create Session

// List call

//// HYDRATE FUNCTIONS

// create service

//// TRANSFORM FUNCTIONS

func getEc2TransitGatewayTurbotTags(_ context.Context, d *transform.TransformData) (interface{}, error) {
	data := d.HydrateItem.(kaytu.EC2TransitGateway).Description.TransitGateway
	return ec2V2TagsToMap(data.Tags)
}

func getEc2TransitGatewayTurbotTitle(_ context.Context, d *transform.TransformData) (interface{}, error) {
	data := d.HydrateItem.(kaytu.EC2TransitGateway).Description.TransitGateway
	title := data.TransitGatewayId
	if data.Tags != nil {
		for _, i := range data.Tags {
			if *i.Key == "Name" {
				title = i.Value
			}
		}
	}
	return title, nil
}
