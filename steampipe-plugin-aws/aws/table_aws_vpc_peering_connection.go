package aws

import (
	"context"
	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

//// TABLE DEFINITION

func tableAwsVpcPeeringConnection(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_vpc_peering_connection",
		Description: "AWS VPC Peering Connection",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListEC2VpcPeeringConnection,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "accepter_cidr_block", Require: plugin.Optional},
				{Name: "accepter_owner_id", Require: plugin.Optional},
				{Name: "accepter_vpc_id", Require: plugin.Optional},
				{Name: "expiration_time", Require: plugin.Optional},
				{Name: "requester_cidr_block", Require: plugin.Optional},
				{Name: "requester_owner_id", Require: plugin.Optional},
				{Name: "requester_vpc_id", Require: plugin.Optional},
				{Name: "status_message", Require: plugin.Optional},
				{Name: "id", Require: plugin.Optional},
			},
		},

		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "id",
				Description: "The ID of the VPC peering connection.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.VpcPeeringConnection.VpcPeeringConnectionId"),
			},
			{
				Name:        "status_code",
				Description: "The status of the VPC peering connection. Possible values include: 'pending-acceptance', 'failed', 'expired', 'provisioning', 'active', 'deleting', 'deleted' or 'rejected'.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.VpcPeeringConnection.Status.Code")},
			{
				Name:        "accepter_cidr_block",
				Description: "The IPv4 CIDR block for the accepter VPC.",
				Type:        proto.ColumnType_CIDR,
				Transform:   transform.FromField("Description.VpcPeeringConnection.AccepterVpcInfo.CidrBlock")},
			{
				Name:        "accepter_owner_id",
				Description: "The ID of the Amazon Web Services account that owns the accepter VPC.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.VpcPeeringConnection.AccepterVpcInfo.OwnerId")},
			{
				Name:        "accepter_region",
				Description: "The Region in which the accepter VPC is located.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.VpcPeeringConnection.AccepterVpcInfo.Region")},
			{
				Name:        "accepter_vpc_id",
				Description: "The ID of the accepter VPC.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.VpcPeeringConnection.AccepterVpcInfo.VpcId")},
			{
				Name:        "expiration_time",
				Description: "The time that an unaccepted VPC peering connection will expire.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.VpcPeeringConnection.ExpirationTime")},
			{
				Name:        "requester_cidr_block",
				Description: "The IPv4 CIDR block for the requester VPC.",
				Type:        proto.ColumnType_CIDR,
				Transform:   transform.FromField("Description.VpcPeeringConnection.RequesterVpcInfo.CidrBlock")},
			{
				Name:        "requester_owner_id",
				Description: "The ID of the Amazon Web Services account that owns the requester VPC.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.VpcPeeringConnection.RequesterVpcInfo.OwnerId")},
			{
				Name:        "requester_region",
				Description: "The Region in which the requester VPC is located.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.VpcPeeringConnection.RequesterVpcInfo.Region")},
			{
				Name:        "requester_vpc_id",
				Description: "The ID of the requester VPC.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.VpcPeeringConnection.RequesterVpcInfo.VpcId")},
			{
				Name:        "status_message",
				Description: "A message that provides more information about the status, if applicable.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.VpcPeeringConnection.Status.Message")},
			{
				Name:        "accepter_cidr_block_set",
				Description: "Information about the IPv4 CIDR blocks for the accepter VPC.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.VpcPeeringConnection.AccepterVpcInfo.CidrBlockSet")},
			{
				Name:        "accepter_ipv6_cidr_block_set",
				Description: "The IPv6 CIDR block for the accepter VPC.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.VpcPeeringConnection.AccepterVpcInfo.Ipv6CidrBlockSet")},
			{
				Name:        "accepter_peering_options",
				Description: "Information about the VPC peering connection options for the accepter VPC.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.VpcPeeringConnection.AccepterVpcInfo.PeeringOptions")},
			{
				Name:        "requester_cidr_block_set",
				Description: "Information about the IPv4 CIDR blocks for the requester VPC.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.VpcPeeringConnection.RequesterVpcInfo.CidrBlockSet")},
			{
				Name:        "requester_ipv6_cidr_block_set",
				Description: "The IPv6 CIDR block for the requester VPC.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.VpcPeeringConnection.RequesterVpcInfo.Ipv6CidrBlockSet")},
			{
				Name:        "requester_peering_options",
				Description: "Information about the VPC peering connection options for the requester VPC.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.VpcPeeringConnection.RequesterVpcInfo.PeeringOptions")},
			{
				Name:        "tags_src",
				Description: "The tags assigned to the resource.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.VpcPeeringConnection.Tags")},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.VpcPeeringConnection.VpcPeeringConnectionId")},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(vpcPeeringConnectionTags),
			},
		}),
	}
}

//// TRANSFORM FUNCTION

func vpcPeeringConnectionTags(_ context.Context, d *transform.TransformData) (interface{}, error) {
	connection := d.HydrateItem.(opengovernance.EC2VpcPeeringConnection).Description.VpcPeeringConnection

	var turbotTagsMap map[string]string
	if connection.Tags != nil {
		turbotTagsMap = map[string]string{}
		for _, i := range connection.Tags {
			turbotTagsMap[*i.Key] = *i.Value
		}
		return turbotTagsMap, nil
	}
	return nil, nil
}
