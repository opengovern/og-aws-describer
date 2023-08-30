package aws

import (
	"context"
	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAwsVpcNetworkACL(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_vpc_network_acl",
		Description: "AWS VPC Network ACL",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("network_acl_id"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"InvalidNetworkAclID.NotFound"}),
			},
			Hydrate: kaytu.GetEC2NetworkAcl,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListEC2NetworkAcl,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "is_default", Require: plugin.Optional, Operators: []string{"=", "<>"}},
				{Name: "owner_id", Require: plugin.Optional},
				{Name: "vpc_id", Require: plugin.Optional},
			},
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "network_acl_id",
				Description: "The ID of the network ACL.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.NetworkAcl.NetworkAclId"),
			},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) specifying the network ACL.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.From(getVpcNetworkACLARN),
			},
			{
				Name:        "is_default",
				Description: "Indicates whether this is the default network ACL for the VPC.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.NetworkAcl.IsDefault"),
			},
			{
				Name:        "vpc_id",
				Description: "The ID of the VPC for the network ACL.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.NetworkAcl.VpcId"),
			},
			{
				Name:        "owner_id",
				Description: "The ID of the AWS account that owns the network ACL.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.NetworkAcl.OwnerId"),
			},
			{
				Name:        "associations",
				Description: "Any associations between the network ACL and one or more subnets.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.NetworkAcl.Associations"),
			},
			{
				Name:        "entries",
				Description: "One or more entries (rules) in the network ACL.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.NetworkAcl.Entries"),
			},
			{
				Name:        "tags_src",
				Description: "A list of tags that are attached to Network ACL.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.NetworkAcl.Tags"),
			},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromP(getVpcNetworkACLTurbotData, "Tags"),
			},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromP(getVpcNetworkACLTurbotData, "Title"),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(getVpcNetworkACLARN).Transform(transform.EnsureStringArray),
			},
		}),
	}
}

//// LIST FUNCTION

//// HYDRATE FUNCTIONS

//// TRANSFORM FUNCTIONS

func getVpcNetworkACLARN(_ context.Context, d *transform.TransformData) (interface{}, error) {
	networkACL := d.HydrateItem.(kaytu.EC2NetworkAcl).Description.NetworkAcl
	metadata := d.HydrateItem.(kaytu.EC2NetworkAcl).Metadata

	// Get data for turbot defined properties
	arn := "arn:" + metadata.Partition + ":ec2:" + metadata.Region + ":" + metadata.AccountID + ":network-acl/" + *networkACL.NetworkAclId

	return arn, nil
}

//// TRANSFORM FUNCTIONS

func getVpcNetworkACLTurbotData(_ context.Context, d *transform.TransformData) (interface{}, error) {
	networkACL := d.HydrateItem.(kaytu.EC2NetworkAcl).Description.NetworkAcl
	param := d.Param.(string)

	// Get resource title
	title := networkACL.NetworkAclId

	// Get the resource tags
	turbotTagsMap := map[string]string{}
	if networkACL.Tags != nil {
		for _, i := range networkACL.Tags {
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
