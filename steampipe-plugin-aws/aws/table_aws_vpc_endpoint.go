package aws

import (
	"context"
	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAwsVpcEndpoint(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_vpc_endpoint",
		Description: "AWS VPC Endpoint",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("vpc_endpoint_id"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"InvalidVpcEndpointId.NotFound", "InvalidVpcEndpointId.Malformed"}),
			},
			Hydrate: kaytu.GetEC2VPCEndpoint,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListEC2VPCEndpoint,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "service_name", Require: plugin.Optional},
				{Name: "vpc_id", Require: plugin.Optional},
				{Name: "state", Require: plugin.Optional},
			},
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "vpc_endpoint_id",
				Description: "The ID of the VPC endpoint.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.VpcEndpoint.VpcEndpointId"),
			},
			{
				Name:        "service_name",
				Description: "The name of the service to which the endpoint is associated.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.VpcEndpoint.ServiceName"),
			},
			{
				Name:        "owner_id",
				Description: "The ID of the AWS account that owns the VPC endpoint.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.VpcEndpoint.OwnerId"),
			},
			{
				Name:        "vpc_id",
				Description: "The ID of the VPC to which the endpoint is associated.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.VpcEndpoint.VpcId"),
			},
			{
				Name:        "vpc_endpoint_type",
				Description: "The type of endpoint.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.VpcEndpoint.VpcEndpointType"),
			},
			{
				Name:        "state",
				Description: "The state of the VPC endpoint.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.VpcEndpoint.State"),
			},
			{
				Name:        "private_dns_enabled",
				Description: "Indicates whether the VPC is associated with a private hosted zone.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.VpcEndpoint.PrivateDnsEnabled"),
			},
			{
				Name:        "requester_managed",
				Description: "Indicates whether the VPC endpoint is being managed by its service.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.VpcEndpoint.RequesterManaged"),
			},
			{
				Name:        "policy",
				Description: "The policy document associated with the endpoint, if applicable.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.VpcEndpoint.PolicyDocument").Transform(transform.UnmarshalYAML),
			},
			{
				Name:        "policy_std",
				Description: "Contains the policy in a canonical form for easier searching.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.VpcEndpoint.PolicyDocument").Transform(unescape).Transform(policyToCanonical),
			},
			{
				Name:        "subnet_ids",
				Description: "One or more subnets in which the endpoint is located.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.VpcEndpoint.SubnetIds"),
			},
			{
				Name:        "route_table_ids",
				Description: "One or more route tables associated with the endpoint.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.VpcEndpoint.RouteTableIds"),
			},
			{
				Name:        "groups",
				Description: "Information about the security groups that are associated with the network interface.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.VpcEndpoint.Groups"),
			},
			{
				Name:        "network_interface_ids",
				Description: "One or more network interfaces for the endpoint.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.VpcEndpoint.NetworkInterfaceIds"),
			},
			{
				Name:        "dns_entries",
				Description: "The DNS entries for the endpoint.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.VpcEndpoint.DnsEntries"),
			},
			{
				Name:        "creation_timestamp",
				Description: "The date and time that the VPC endpoint was created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.VpcEndpoint.CreationTimestamp"),
			},
			{
				Name:        "tags_src",
				Description: "A list of tags assigned to the VPC endpoint.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.VpcEndpoint.Tags"),
			},

			//standard columns
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromP(getVpcEndpointTurbotData, "Tags"),
			},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromP(getVpcEndpointTurbotData, "Title"),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(getVpcEndpointAkas),
			},
		}),
	}
}

//// LIST FUNCTION

//// HYDRATE FUNCTIONS

//// TRANSFORM FUNCTIONS

func getVpcEndpointAkas(_ context.Context, d *transform.TransformData) (interface{}, error) {
	vpcEndpoint := d.HydrateItem.(kaytu.EC2VPCEndpoint).Description.VpcEndpoint
	metadata := d.HydrateItem.(kaytu.EC2VPCEndpoint).Metadata

	akas := []string{"arn:" + metadata.Partition + ":ec2:" + metadata.Region + ":" + metadata.AccountID + ":vpc-endpoint/" + *vpcEndpoint.VpcEndpointId}

	return akas, nil
}

//// TRANSFORM FUNCTIONS

func getVpcEndpointTurbotData(_ context.Context, d *transform.TransformData) (interface{}, error) {
	vpcEndpoint := d.HydrateItem.(kaytu.EC2VPCEndpoint).Description.VpcEndpoint
	param := d.Param.(string)

	// Get resource title
	title := vpcEndpoint.VpcEndpointId

	// Get the resource tags
	var turbotTagsMap map[string]string
	if vpcEndpoint.Tags != nil {
		turbotTagsMap = map[string]string{}
		for _, i := range vpcEndpoint.Tags {
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
