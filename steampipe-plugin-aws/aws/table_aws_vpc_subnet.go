package aws

import (
	"context"

	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAwsVpcSubnet(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_vpc_subnet",
		Description: "AWS VPC Subnet",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("subnet_id"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"InvalidSubnetID.Malformed", "InvalidSubnetID.NotFound"}),
			},
			Hydrate: kaytu.GetEC2Subnet,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListEC2Subnet,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "availability_zone", Require: plugin.Optional},
				{Name: "availability_zone_id", Require: plugin.Optional},
				{Name: "available_ip_address_count", Require: plugin.Optional},
				{Name: "cidr_block", Require: plugin.Optional},
				{Name: "default_for_az", Require: plugin.Optional},
				{Name: "outpost_arn", Require: plugin.Optional},
				{Name: "owner_id", Require: plugin.Optional},
				{Name: "state", Require: plugin.Optional},
				{Name: "subnet_arn", Require: plugin.Optional},
				{Name: "vpc_id", Require: plugin.Optional},
			},
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "subnet_id",
				Description: "Contains the unique ID to specify a subnet.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Subnet.SubnetId"),
			},
			{
				Name:        "subnet_arn",
				Description: "Contains the Amazon Resource Name (ARN) of the subnet.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Subnet.SubnetArn"),
			},
			{
				Name:        "vpc_id",
				Description: "ID of the VPC, the subnet is in.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Subnet.VpcId"),
			},
			{
				Name:        "cidr_block",
				Description: "Contains the IPv4 CIDR block assigned to the subnet.",
				Type:        proto.ColumnType_CIDR,
				Transform:   transform.FromField("Description.Subnet.CidrBlock"),
			},
			{
				Name:        "state",
				Description: "Current state of the subnet.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Subnet.State"),
			},
			{
				Name:        "owner_id",
				Description: "Contains the AWS account that own the subnet.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Subnet.OwnerId"),
			},
			{
				Name:        "assign_ipv6_address_on_creation",
				Description: "Indicates whether a network interface created in this subnet (including a network interface created by RunInstances) receives an IPv6 address.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Subnet.AssignIpv6AddressOnCreation"),
			},
			{
				Name:        "available_ip_address_count",
				Description: "The number of unused private IPv4 addresses in the subnet. The IPv4 addresses for any stopped instances are considered unavailable.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Subnet.AvailableIpAddressCount"),
			},
			{
				Name:        "availability_zone",
				Description: "The Availability Zone of the subnet.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Subnet.AvailabilityZone"),
			},
			{
				Name:        "availability_zone_id",
				Description: "The AZ ID of the subnet.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Subnet.AvailabilityZoneId"),
			},
			{
				Name:        "customer_owned_ipv4_pool",
				Description: "The customer-owned IPv4 address pool associated with the subnet.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Subnet.CustomerOwnedIpv4Pool"),
			},
			{
				Name:        "default_for_az",
				Description: "Indicates whether this is the default subnet for the Availability Zone.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Subnet.DefaultForAz"),
			},
			{
				Name:        "map_customer_owned_ip_on_launch",
				Description: "Indicates whether a network interface created in this subnet (including a network interface created by RunInstances) receives a customer-owned IPv4 address.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Subnet.MapCustomerOwnedIpOnLaunch"),
			},
			{
				Name:        "map_public_ip_on_launch",
				Description: "Indicates whether instances launched in this subnet receive a public IPv4 address.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Subnet.MapPublicIpOnLaunch"),
			},
			{
				Name:        "outpost_arn",
				Description: "The Amazon Resource Name (ARN) of the Outpost. Available only if subnet is on an outpost.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Subnet.OutpostArn"),
			},
			{
				Name:        "ipv6_cidr_block_association_set",
				Description: "A list of IPv6 CIDR blocks associated with the subnet.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Subnet.Ipv6CidrBlockAssociationSet"),
			},
			{
				Name:        "tags_src",
				Description: "A list of tags that are attached to the subnet.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Subnet.Tags"),
			},

			// Standard columns for all tables
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Subnet.Tags").Transform(subnetTagListToTurbotTags),
			},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.From(getSubnetTurbotTitle),
			},
			{
				//error we don't have subnetArn
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("SubnetArn").Transform(arnToAkas),
			},
		}),
	}
}

//// LIST FUNCTION

//// HYDRATE FUNCTIONS

//// TRANSFORM FUNCTIONS

func subnetTagListToTurbotTags(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	if tagList, ok := d.Value.([]types.Tag); ok {
		// Mapping the resource tags inside turbotTags
		var turbotTagsMap map[string]string
		if tagList != nil {
			turbotTagsMap = map[string]string{}
			for _, i := range tagList {
				turbotTagsMap[*i.Key] = *i.Value
			}
		}
		return turbotTagsMap, nil
	} else {
		return nil, nil
	}
}

func getSubnetTurbotTitle(_ context.Context, d *transform.TransformData) (interface{}, error) {
	subnet := d.HydrateItem.(kaytu.EC2Subnet).Description.Subnet
	var title string
	tags := subnet.Tags
	for _, i := range tags {
		if *i.Key == "Name" {
			title = *i.Value
		}
	}

	if title == "" {
		title = *subnet.SubnetId
	}

	return title, nil
}
