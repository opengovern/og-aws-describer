package aws

import (
	"context"
	"fmt"
	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"

	"github.com/aws/aws-sdk-go-v2/service/ec2/types"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAwsVpcDhcpOptions(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_vpc_dhcp_options",
		Description: "AWS VPC DHCP Options",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("dhcp_options_id"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"InvalidDhcpOptionID.NotFound"}),
			},
			Hydrate: kaytu.GetEC2DhcpOptions,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListEC2DhcpOptions,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "owner_id", Require: plugin.Optional},
			},
		},

		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "dhcp_options_id",
				Description: "The ID of the set of DHCP options.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DhcpOptions.DhcpOptionsId")},
			{
				Name:        "owner_id",
				Description: "The ID of the AWS account that owns the DHCP options set.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DhcpOptions.OwnerId")},
			{
				Name:        "domain_name",
				Description: "The domain name for instances. This value is used to complete unqualified DNS hostnames.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DhcpOptions.DhcpConfigurations").TransformP(dhcpConfigurationToStringSlice, "domain-name"),
			},
			{
				Name:        "domain_name_servers",
				Description: "The IP addresses of up to four domain name servers, or AmazonProvidedDNS.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DhcpOptions.DhcpConfigurations").TransformP(dhcpConfigurationToStringSlice, "domain-name-servers"),
			},
			{
				Name:        "netbios_name_servers",
				Description: "The IP addresses of up to four NetBIOS name servers.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DhcpOptions.DhcpConfigurations").TransformP(dhcpConfigurationToStringSlice, "netbios-name-servers"),
			},
			{
				Name:        "netbios_node_type",
				Description: "The NetBIOS node type (1, 2, 4, or 8).",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DhcpOptions.DhcpConfigurations").TransformP(dhcpConfigurationToStringSlice, "netbios-node-type"),
			},
			{
				Name:        "ntp_servers",
				Description: "The IP addresses of up to four Network Time Protocol (NTP) servers.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DhcpOptions.DhcpConfigurations").TransformP(dhcpConfigurationToStringSlice, "ntp-servers"),
			},
			{
				Name:        "tags_src",
				Description: "A list of tags that are attached to vpc dhcp options.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DhcpOptions.Tags")},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromP(vpcDhcpOptionsAPIDataToTurbotData, "Title"),
			},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromP(vpcDhcpOptionsAPIDataToTurbotData, "Tags"),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(getEc2DhcpOptionsArn).Transform(arnToAkas)},
		}),
	}
}

//// TRANSFORM FUNCTIONS

func getEc2DhcpOptionsArn(_ context.Context, d *transform.TransformData) (interface{}, error) {
	dhcpOptions := d.HydrateItem.(kaytu.EC2DhcpOptions).Description.DhcpOptions
	metadata := d.HydrateItem.(kaytu.EC2DhcpOptions).Metadata

	arn := fmt.Sprintf("arn:%s:ec2:%s:%s:dhcp-options/%s", metadata.Partition, metadata.Region, metadata.AccountID, *dhcpOptions.DhcpOptionsId)
	return arn, nil
}

func dhcpConfigurationToStringSlice(_ context.Context, d *transform.TransformData) (interface{}, error) {
	if d.Value == nil {
		return nil, nil
	}
	dhcpConfigurations := d.Value.([]types.DhcpConfiguration)

	var values []*string
	for _, configuration := range dhcpConfigurations {
		if *configuration.Key == d.Param.(string) {
			values = mapString(configuration.Values)
		}
	}
	return values, nil
}

func vpcDhcpOptionsAPIDataToTurbotData(_ context.Context, d *transform.TransformData) (interface{}, error) {
	vpcDhcpOptions := d.HydrateItem.(kaytu.EC2DhcpOptions).Description.DhcpOptions
	param := d.Param.(string)

	// Get resource title
	title := *vpcDhcpOptions.DhcpOptionsId

	// Get the resource tags
	var turbotTagsMap map[string]string
	if vpcDhcpOptions.Tags != nil {
		turbotTagsMap = map[string]string{}
		for _, i := range vpcDhcpOptions.Tags {
			turbotTagsMap[*i.Key] = *i.Value
			if *i.Key == "Name" {
				title = *i.Value
			}
		}
	}

	if param == "Tags" {
		return turbotTagsMap, nil
	}

	return title, nil
}

func mapString(l []types.AttributeValue) []*string {
	var values []*string
	for _, v := range l {
		values = append(values, v.Value)
	}
	return values
}