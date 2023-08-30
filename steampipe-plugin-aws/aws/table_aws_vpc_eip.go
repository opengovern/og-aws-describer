package aws

import (
	"context"

	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAwsVpcEip(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_vpc_eip",
		Description: "AWS VPC Elastic IP",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("allocation_id"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"InvalidAllocationID.NotFound", "InvalidAllocationID.Malformed"}),
			},
			Hydrate: kaytu.GetEC2EIP,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListEC2EIP,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "association_id", Require: plugin.Optional},
				{Name: "domain", Require: plugin.Optional},
				{Name: "instance_id", Require: plugin.Optional},
				{Name: "network_border_group", Require: plugin.Optional},
				{Name: "network_interface_id", Require: plugin.Optional},
				{Name: "network_interface_owner_id", Require: plugin.Optional},
				{Name: "private_ip_address", Require: plugin.Optional},
				{Name: "public_ip", Require: plugin.Optional},
			},
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "allocation_id",
				Description: "Contains the ID representing the allocation of the address for use with EC2-VPC.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Address.AllocationId"),
			},
			{
				// EIPs in EC2-Classic have no valid ARN due to no allocation ID
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) specifying the VPC EIP.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("ARN"),
			},
			{
				Name:        "public_ip",
				Description: "Contains the Elastic IP address.",
				Type:        proto.ColumnType_IPADDR,
				Transform:   transform.FromField("Description.Address.PublicIp"),
			},
			{
				Name:        "public_ipv4_pool",
				Description: "The ID of an address pool.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Address.PublicIpv4Pool"),
			},
			{
				Name:        "domain",
				Description: "Indicates whether Elastic IP address is for use with instances in EC2-Classic(standard) or instances in a VPC (vpc).",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Address.Domain"),
			},
			{
				Name:        "association_id",
				Description: "Contains the ID representing the association of the address with an instance in a VPC.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Address.AssociationId"),
			},
			{
				Name:        "carrier_ip",
				Description: "The carrier IP address associated. This option is only available for network interfaces which reside in a subnet in a Wavelength Zone (for example an EC2 instance).",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Address.CarrierIp"),
			},
			{
				Name:        "customer_owned_ip",
				Description: "The customer-owned IP address.",
				Type:        proto.ColumnType_IPADDR,
				Transform:   transform.FromField("Description.Address.CustomerOwnedIp"),
			},
			{
				Name:        "customer_owned_ipv4_pool",
				Description: "The ID of the customer-owned address pool.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Address.CustomerOwnedIpv4Pool"),
			},
			{
				Name:        "instance_id",
				Description: "Contains the ID of the instance that the address is associated with.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Address.InstanceId"),
			},
			{
				Name:        "network_border_group",
				Description: "The name of the unique set of Availability Zones, Local Zones, or Wavelength Zones from which AWS advertises IP addresses.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Address.NetworkBorderGroup"),
			},
			{
				Name:        "network_interface_id",
				Description: "The ID of the network interface.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Address.NetworkInterfaceId"),
			},
			{
				Name:        "network_interface_owner_id",
				Description: "The ID of the AWS account that owns the network interface.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Address.NetworkInterfaceOwnerId"),
			},
			{
				Name:        "private_ip_address",
				Description: "The private IP address associated with the Elastic IP address.",
				Type:        proto.ColumnType_IPADDR,
				Transform:   transform.FromField("Description.Address.PrivateIpAddress"),
			},
			{
				Name:        "tags_src",
				Description: "A list of tags that are attached to the vpc.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Address.Tags"),
			},

			// Standard columns for all tables
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(getVpcEipTurbotTags),
			},
			{
				// Fallback to public IP for EIPs in EC2-Classic
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Address.AssociationId"),
			},
			{
				// EIPs in EC2-Classic have no valid ARN, so no valid AKAs either
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("ARN").Transform(transform.EnsureStringArray),
			},
		}),
	}
}

//// TRANSFORM FUNCTIONS

func getVpcEipTurbotTags(_ context.Context, d *transform.TransformData) (interface{}, error) {
	eip := d.HydrateItem.(kaytu.EC2EIP).Description.Address
	return ec2V2TagsToMap(eip.Tags)
}
