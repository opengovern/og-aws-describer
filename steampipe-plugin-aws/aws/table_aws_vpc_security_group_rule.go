package aws

import (
	"context"
	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAwsVpcSecurityGroupRule(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_vpc_security_group_rule",
		Description: "AWS VPC Security Group Rule",
		// TODO -- get call returning a list of items
		//DefaultIgnoreConfig: &plugin.IgnoreConfig{
		//	ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"InvalidGroup.NotFound", "InvalidSecurityGroupRuleId.Malformed", "InvalidSecurityGroupRuleId.NotFound"}),
		//},
		//Get: &plugin.GetConfig{
		//	KeyColumns: plugin.SingleColumn("security_group_rule_id"),
		//	IgnoreConfig: &plugin.IgnoreConfig{
		//		ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"InvalidGroup.NotFound", "InvalidSecurityGroupRuleId.Malformed", "InvalidSecurityGroupRuleId.NotFound"}),
		//	},
		//	Hydrate: getSecurityGroupRule,
		//},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListEC2SecurityGroupRule,
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:    "group_id",
					Require: plugin.Optional,
				},
			},
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "security_group_rule_id",
				Description: "The ID of the security group rule.",
				Type:        proto.ColumnType_STRING,
				//error
			},
			{
				Name:        "group_name",
				Description: "The name of the security group to which rule belongs.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Group.GroupName")},
			{
				Name:        "group_id",
				Description: "The ID of the security group to which rule belongs.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Group.GroupId")},
			{
				Name:        "is_egress",
				Description: "Indicates whether the security group rule is an outbound rule.",
				Type:        proto.ColumnType_BOOL,
				//error
			},
			{
				Name:        "type",
				Description: "Type of the rule ( ingress | egress).",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Type")},
			{
				Name:        "vpc_id",
				Description: "The ID of the VPC for the security group.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Group.VpcId")},
			{
				Name:        "owner_id",
				Description: "The AWS account ID of the owner of the security group to which rule belongs.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Group.OwnerId")},
			{
				Name:        "group_owner_id",
				Description: "The ID of the Amazon Web Services account that owns the security group.",
				Type:        proto.ColumnType_STRING,
				//error
			},
			{
				Name:        "description",
				Description: "The security group rule description.",
				Type:        proto.ColumnType_STRING,
				//error
			},
			{
				Name:        "ip_protocol",
				Description: "The IP protocol name (tcp, udp, icmp, icmpv6) or number [see Protocol Numbers ](http://www.iana.org/assignments/protocol-numbers/protocol-numbers.xhtml). Use -1 to specify all protocols. When authorizing security group rules, specifying -1 or a protocol number other than tcp, udp, icmp, or icmpv6 allows traffic on all ports, regardless of any port range specified. For tcp, udp, and icmp, a port range is specified. For icmpv6, the port range is optional. If port range is omitted, traffic for all types and codes is allowed.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Permission.IpProtocol")},
			{
				Name:        "from_port",
				Description: "The start of port range for the TCP and UDP protocols, or an ICMP/ICMPv6 type number. A value of -1 indicates all ICMP/ICMPv6 types.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Permission.FromPort")},
			{
				Name:        "to_port",
				Description: "The end of port range for the TCP and UDP protocols, or an ICMP/ICMPv6 code. A value of -1 indicates all ICMP/ICMPv6 codes.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Permission.ToPort")},
			{
				Name:        "cidr_ip",
				Description: "The IPv4 CIDR range. It can be either a CIDR range or a source security group, not both. A single IPv4 address is denoted by /32 prefix length.",
				Type:        proto.ColumnType_CIDR,
				Transform:   transform.FromField("Description.IPRange.CidrIp")},

			{
				Name:        "cidr_ipv4",
				Description: "The IPv4 CIDR range.",
				Type:        proto.ColumnType_CIDR,
				//error
			},
			{
				Name:        "cidr_ipv6",
				Description: "The IPv6 CIDR range. It can be either a CIDR range or a source security group, not both. A single IPv6 address is denoted by /128 prefix length.",
				Type:        proto.ColumnType_CIDR,
				Transform:   transform.FromField("Description.Ipv6Range.CidrIpv6")},
			{
				Name:        "pair_group_id",
				Description: "The ID of the security group that references this user ID group pair.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.UserIDGroupPair.GroupId")},

			{
				Name:        "referenced_group_id",
				Description: "The ID of the referenced security group.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("ReferencedGroupInfo.GroupId"),
				//error
			},
			{
				Name:        "pair_group_name",
				Description: "The name of the security group that references this user ID group pair.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.UserIDGroupPair.GroupName")},
			{
				Name:        "pair_peering_status",
				Description: "The status of a VPC peering connection, if applicable.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.UserIDGroupPair.PeeringStatus")},
			{
				Name:        "referenced_peering_status",
				Description: "The status of a VPC peering connection, if applicable.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("ReferencedGroupInfo.PeeringStatus"),
				//error
			},
			{
				Name:        "pair_user_id",
				Description: "The ID of an AWS account. For a referenced security group in another VPC, the account ID of the referenced security group is returned in the response. If the referenced security group is deleted, this value is not returned.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.UserIDGroupPair.UserId")},
			{
				Name:        "referenced_user_id",
				Description: "The ID of an AWS account. For a referenced security group in another VPC, the account ID of the referenced security group is returned in the response. If the referenced security group is deleted, this value is not returned.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("ReferencedGroupInfo.UserId"),
				//error
			},
			{
				Name:        "pair_vpc_id",
				Description: "The ID of the VPC for the referenced security group, if applicable.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.UserIDGroupPair.VpcId")},
			{
				Name:        "referenced_vpc_id",
				Description: "The ID of the VPC for the referenced security group, if applicable.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("ReferencedGroupInfo.VpcId"),
				//error
			},
			{
				Name:        "pair_vpc_peering_connection_id",
				Description: "The ID of the VPC peering connection, if applicable.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.UserIDGroupPair.VpcPeeringConnectionId")},

			{
				Name:        "referenced_vpc_peering_connection_id",
				Description: "The ID of the VPC peering connection, if applicable.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.UserIDGroupPair.VpcPeeringConnectionId")},
			//error
			{
				Name:        "prefix_list_id",
				Description: "The ID of the referenced prefix list.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.PrefixListId.PrefixListId")},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Hydrate:     getSecurityGroupRuleTurbotData,
				Transform:   transform.FromField("Title"),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Hydrate:     getSecurityGroupRuleTurbotData,
				Transform:   transform.FromField("Akas"),
			},
		}),
	}
}

//// HYDRATE FUNCTIONS

func getSecurityGroupRuleTurbotData(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	sgRule := h.Item.(kaytu.EC2SecurityGroupRule).Description
	metadata := h.Item.(kaytu.EC2SecurityGroupRule).Metadata

	// To create uninque aka
	hashCode := sgRule.Type + "_" + *sgRule.Permission.IpProtocol
	if sgRule.Permission.FromPort != nil {
		hashCode = hashCode + "_" + string(*(sgRule.Permission.FromPort)) + "_" + string(*(sgRule.Permission.ToPort))
	}

	if sgRule.IPRange != nil && sgRule.IPRange.CidrIp != nil {
		hashCode = hashCode + "_" + *sgRule.IPRange.CidrIp
	} else if sgRule.Ipv6Range != nil && sgRule.Ipv6Range.CidrIpv6 != nil {
		hashCode = hashCode + "_" + *sgRule.Ipv6Range.CidrIpv6
	} else if sgRule.UserIDGroupPair != nil && *sgRule.UserIDGroupPair.GroupId == *sgRule.Group.GroupId {
		hashCode = hashCode + "_" + *sgRule.Group.GroupId
	} else if sgRule.PrefixListId != nil && sgRule.PrefixListId.PrefixListId != nil {
		hashCode = hashCode + "_" + *sgRule.PrefixListId.PrefixListId
	}

	// generate aka for the rule
	akas := []string{"arn:" + metadata.Partition + ":ec2:" + metadata.Region + ":" + *sgRule.Group.OwnerId + ":security-group/" + *sgRule.Group.GroupId + ":" + hashCode}

	title := *sgRule.Group.GroupId + "_" + hashCode

	turbotData := map[string]interface{}{
		"Akas":  akas,
		"Title": title,
	}

	return turbotData, nil
}
