package aws

import (
	"context"
	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAwsVpcNatGateway(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_vpc_nat_gateway",
		Description: "AWS VPC Network Address Translation Gateway",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("nat_gateway_id"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"NatGatewayMalformed", "NatGatewayNotFound"}),
			},
			Hydrate: kaytu.GetEC2NatGateway,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListEC2NatGateway,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "state", Require: plugin.Optional},
				{Name: "subnet_id", Require: plugin.Optional},
				{Name: "vpc_id", Require: plugin.Optional},
			},
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "nat_gateway_id",
				Description: "The ID of the NAT gateway.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.NatGateway.NatGatewayId"),
			},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) specifying the NAT gateway.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.From(getVpcNatGatewayARN),
			},
			{
				Name:        "nat_gateway_addresses",
				Description: "Information about the IP addresses and network interface associated with the NAT gateway.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.NatGateway.NatGatewayAddresses"),
			},
			{
				Name:        "state",
				Description: "The current state of the NAT gateway (pending | failed | available | deleting | deleted).",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.NatGateway.State"),
			},
			{
				Name:        "create_time",
				Description: "The date and time the NAT gateway was created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.NatGateway.CreateTime"),
			},
			{
				Name:        "vpc_id",
				Description: "The ID of the VPC in which the NAT gateway is located.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.NatGateway.VpcId"),
			},
			{
				Name:        "subnet_id",
				Description: "The ID of the subnet in which the NAT gateway is located.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.NatGateway.SubnetId"),
			},
			{
				Name:        "delete_time",
				Description: "The date and time the NAT gateway was deleted, if applicable.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.NatGateway.DeleteTime"),
			},
			{
				Name:        "failure_code",
				Description: "If the NAT gateway could not be created, specifies the error code for the failure. (InsufficientFreeAddressesInSubnet | Gateway.NotAttached | InvalidAllocationID.NotFound | Resource.AlreadyAssociated | InternalError | InvalidSubnetID.NotFound).",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.NatGateway.FailureCode"),
			},
			{
				Name:        "failure_message",
				Description: "If the NAT gateway could not be created, specifies the error message for the failure.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.NatGateway.FailureMessage"),
			},
			{
				Name:        "provisioned_bandwidth",
				Description: "Reserved. If you need to sustain traffic greater than the documented limits (https://docs.aws.amazon.com/vpc/latest/userguide/vpc-nat-gateway.html).",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.NatGateway.ProvisionedBandwidth"),
			},
			{
				Name:        "tags_src",
				Description: "A list of tags that are attached to NAT gateway.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.NatGateway.Tags"),
			},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromP(getVpcNatGatewayTurbotData, "Tags"),
			},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromP(getVpcNatGatewayTurbotData, "Title"),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(getVpcNatGatewayARN).Transform(transform.EnsureStringArray),
			},
		}),
	}
}

//// LIST FUNCTION

//// HYDRATE FUNCTIONS

//// TRANSFORM FUNCTIONS

func getVpcNatGatewayARN(_ context.Context, d *transform.TransformData) (interface{}, error) {
	natGateway := d.HydrateItem.(kaytu.EC2NatGateway).Description.NatGateway
	metadata := d.HydrateItem.(kaytu.EC2NatGateway).Metadata

	arn := "arn:" + metadata.Partition + ":ec2:" + metadata.Region + ":" + metadata.AccountID + ":natgateway/" + *natGateway.NatGatewayId
	return arn, nil
}

//// TRANSFORM FUNCTIONS

func getVpcNatGatewayTurbotData(_ context.Context, d *transform.TransformData) (interface{}, error) {
	natGateway := d.HydrateItem.(kaytu.EC2NatGateway).Description.NatGateway
	param := d.Param.(string)

	// Get resource title
	title := natGateway.NatGatewayId

	// Get the resource tags
	var turbotTagsMap map[string]string
	if natGateway.Tags != nil {
		turbotTagsMap = map[string]string{}
		for _, i := range natGateway.Tags {
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
