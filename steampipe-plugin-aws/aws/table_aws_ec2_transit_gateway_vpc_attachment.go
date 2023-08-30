package aws

import (
	"context"

	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAwsEc2TransitGatewayVpcAttachment(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name: "aws_ec2_transit_gateway_vpc_attachment",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("transit_gateway_attachment_id"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"InvalidTransitGatewayAttachmentID.NotFound", "InvalidTransitGatewayAttachmentID.Unavailable", "InvalidTransitGatewayAttachmentID.Malformed", "InvalidAction"}),
			},
			Hydrate: kaytu.GetEC2TransitGatewayAttachment,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListEC2TransitGatewayAttachment,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "association_state", Require: plugin.Optional},
				{Name: "association_transit_gateway_route_table_id", Require: plugin.Optional},
				{Name: "resource_id", Require: plugin.Optional},
				{Name: "resource_owner_id", Require: plugin.Optional},
				{Name: "resource_type", Require: plugin.Optional},
				{Name: "state", Require: plugin.Optional},
				{Name: "transit_gateway_id", Require: plugin.Optional},
				{Name: "transit_gateway_owner_id", Require: plugin.Optional},
			},
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"InvalidAction"}),
			},
		},

		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "transit_gateway_attachment_id",
				Description: "The ID of the transit gateway attachment.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.TransitGatewayAttachment.TransitGatewayAttachmentId")},
			{
				Name:        "transit_gateway_id",
				Description: "The ID of the transit gateway.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.TransitGatewayAttachment.TransitGatewayId")},
			{
				Name:        "transit_gateway_owner_id",
				Description: "The ID of the AWS account that owns the transit gateway.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.TransitGatewayAttachment.TransitGatewayOwnerId")},
			{
				Name:        "state",
				Description: "The attachment state of the transit gateway attachment.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.TransitGatewayAttachment.State")},
			{
				Name:        "creation_time",
				Description: "The creation time of the transit gateway attachment.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.TransitGatewayAttachment.CreationTime")},
			{
				Name:        "resource_id",
				Description: "The ID of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.TransitGatewayAttachment.ResourceId")},
			{
				Name:        "resource_type",
				Description: "The resource type of the transit gateway attachment.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.TransitGatewayAttachment.ResourceType")},
			{
				Name:        "resource_owner_id",
				Description: "The ID of the AWS account that owns the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.TransitGatewayAttachment.ResourceOwnerId")},
			{
				Name:        "association_state",
				Description: "The state of the association.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.TransitGatewayAttachment.Association.State")},
			{
				Name:        "association_transit_gateway_route_table_id",
				Description: "The ID of the route table for the transit gateway.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.TransitGatewayAttachment.Association.TransitGatewayRouteTableId")},
			{
				Name:        "tags_src",
				Description: "A list of tags assigned.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.TransitGatewayAttachment.Tags")},

			/// Standard columns
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(transitGatewayAttachmentRawTagsToTurbotTags),
			},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.From(getEc2TransitGatewayAttachmentTitle),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("ARN").Transform(arnToAkas)},
		}),
	}
}

//// TRANSFORM FUNCTIONS

func transitGatewayAttachmentRawTagsToTurbotTags(_ context.Context, d *transform.TransformData) (interface{}, error) {
	data := d.HydrateItem.(kaytu.EC2TransitGatewayAttachment).Description.TransitGatewayAttachment
	var turbotTagsMap map[string]string
	if data.Tags == nil {
		return nil, nil
	}

	turbotTagsMap = map[string]string{}
	for _, i := range data.Tags {
		turbotTagsMap[*i.Key] = *i.Value
	}

	return &turbotTagsMap, nil
}

func getEc2TransitGatewayAttachmentTitle(_ context.Context, d *transform.TransformData) (interface{}, error) {
	data := d.HydrateItem.(kaytu.EC2TransitGatewayAttachment).Description.TransitGatewayAttachment
	title := data.TransitGatewayAttachmentId
	if data.Tags != nil {
		for _, i := range data.Tags {
			if *i.Key == "Name" {
				title = i.Value
			}
		}
	}
	return title, nil
}
