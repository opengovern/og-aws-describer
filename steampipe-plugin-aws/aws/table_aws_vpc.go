package aws

import (
	"context"
	"fmt"
	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAwsVpc(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_vpc",
		Description: "AWS VPC",

		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("vpc_id"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"NotFoundException", "InvalidVpcID.NotFound"}),
			},
			Hydrate: kaytu.ListEC2Vpc,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListEC2Vpc,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "cidr_block", Require: plugin.Optional},
				{Name: "dhcp_options_id", Require: plugin.Optional},
				{Name: "is_default", Require: plugin.Optional, Operators: []string{"=", "<>"}},
				{Name: "owner_id", Require: plugin.Optional},
				{Name: "state", Require: plugin.Optional},
			},
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "vpc_id",
				Description: "The ID of the VPC.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Vpc.VpcId"),
			},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) specifying the vpc.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("ARN"),
			},
			{
				Name:        "cidr_block",
				Description: "The primary IPv4 CIDR block for the VPC.",
				Type:        proto.ColumnType_CIDR,
				Transform:   transform.FromField("Description.Vpc.CidrBlock"),
			},
			{
				Name:        "state",
				Description: "Contains the current state of the VPC.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Vpc.State"),
			},
			{
				Name:        "is_default",
				Description: "Indicates whether the VPC is the default VPC.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Vpc.IsDefault"),
			},
			{
				Name:        "dhcp_options_id",
				Description: "Contains the ID of the set of DHCP options, associated with the VPC.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Vpc.DhcpOptionsId"),
			},
			{
				Name:        "instance_tenancy",
				Description: "The allowed tenancy of instances launched into the VPC.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Vpc.InstanceTenancy"),
			},
			{
				Name:        "owner_id",
				Description: "Contains ID of the AWS account that owns the VPC.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Vpc.OwnerId"),
			},
			{
				Name:        "cidr_block_association_set",
				Description: "Information about the IPv4 CIDR blocks associated with the VPC.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Vpc.CidrBlockAssociationSet"),
			},
			{
				Name:        "ipv6_cidr_block_association_set",
				Description: "Information about the IPv6 CIDR blocks associated with the VPC.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Vpc.Ipv6CidrBlockAssociationSet"),
			},
			{
				Name:        "tags_src",
				Description: "A list of tags that are attached with the VPC.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Vpc.Tags"),
			},

			// Standard columns for all tables
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.From(getVpcTurbotTitle),
			},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(getVpcTurbotTags),
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

func getVpcTurbotTags(_ context.Context, d *transform.TransformData) (interface{}, error) {
	vpc := d.HydrateItem.(kaytu.EC2Vpc).Description.Vpc
	if vpc.Tags != nil {
		// Mapping the resource tags inside turbotTags
		var turbotTagsMap map[string]string
		turbotTagsMap = map[string]string{}
		for _, i := range vpc.Tags {
			turbotTagsMap[*i.Key] = *i.Value
		}
		return turbotTagsMap, nil
	} else {
		return nil, nil
	}
}

func getVpcTurbotTitle(_ context.Context, d *transform.TransformData) (interface{}, error) {
	vpc := d.HydrateItem.(kaytu.EC2Vpc).Description.Vpc

	if vpc.Tags != nil {
		for _, i := range vpc.Tags {
			if *i.Key == "Name" {
				return *i.Value, nil
			}
		}
	}
	return vpc.VpcId, nil
}

//// UTILITY FUNCTION

// Build vpc resources filter parameter

type VpcFilterKeyMap struct {
	ColumnName string
	FilterName string
	ColumnType string
}

func buildVpcResourcesFilterParameter(keyMap []VpcFilterKeyMap, quals plugin.KeyColumnQualMap) []types.Filter {
	filters := make([]types.Filter, 0)
	for _, keyDetail := range keyMap {
		if quals[keyDetail.ColumnName] != nil {
			filter := &types.Filter{
				Name: aws.String(keyDetail.FilterName),
			}

			value := getQualsValueByColumn(quals, keyDetail.ColumnName, keyDetail.ColumnType)
			switch keyDetail.ColumnType {
			case "string":
				val, ok := value.(string)
				if ok {
					filter.Values = []string{val}
				}
			case "int64":
				val, ok := value.(int64)
				if ok {
					filter.Values = []string{fmt.Sprint(val)}
				}
			case "double":
				val, ok := value.(float64)
				if ok {
					filter.Values = []string{fmt.Sprint(val)}
				}
			case "ipaddr":
				val, ok := value.(string)
				if ok {
					filter.Values = []string{fmt.Sprint(val)}
				}
			case "cidr": // Ip address with mask
				val, ok := value.(string)
				if ok {
					filter.Values = []string{fmt.Sprint(val)}
				}
			case "time":
				val, ok := value.(time.Time)
				if ok {
					v := val.Format(time.RFC3339)
					filter.Values = []string{fmt.Sprint(v)}
				}
			case "boolean":
				filter.Values = []string{fmt.Sprint(value)}
			}

			filters = append(filters, *filter)
		}
	}

	return filters
}
