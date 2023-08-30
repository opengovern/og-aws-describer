package aws

import (
	"context"
	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAwsVpcInternetGateway(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_vpc_internet_gateway",
		Description: "AWS VPC Internet Gateway",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("internet_gateway_id"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"InvalidInternetGatewayID.NotFound", "InvalidInternetGatewayID.Malformed"}),
			},
			Hydrate: kaytu.GetEC2InternetGateway,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListEC2InternetGateway,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "owner_id", Require: plugin.Optional},
			},
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "internet_gateway_id",
				Description: "The ID of the internet gateway.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.InternetGateway.InternetGatewayId"),
			},
			{
				Name:        "owner_id",
				Description: "The ID of the AWS account that owns the internet gateway.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.InternetGateway.OwnerId"),
			},
			{
				Name:        "attachments",
				Description: "Any VPCs attached to the internet gateway.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.InternetGateway.Attachments"),
			},
			{
				Name:        "tags_src",
				Description: "tags assigned to the internet gateway.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.InternetGateway.Tags"),
			},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromP(getVpcInternetGatewayTurbotData, "Tags"),
			},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromP(getVpcInternetGatewayTurbotData, "Title"),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(getVpcInternetGatewayTurbotAkas),
			},
		}),
	}
}

//// LIST FUNCTION

//// HYDRATE FUNCTIONS

//// TRANSFORM FUNCTIONS

func getVpcInternetGatewayTurbotAkas(_ context.Context, d *transform.TransformData) (interface{}, error) {
	internetGateway := d.HydrateItem.(kaytu.EC2InternetGateway).Description.InternetGateway
	metadata := d.HydrateItem.(kaytu.EC2InternetGateway).Metadata

	// Get data for turbot defined properties
	akas := []string{"arn:" + metadata.Partition + ":ec2:" + metadata.Region + ":" + metadata.AccountID + ":internet-gateway/" + *internetGateway.InternetGatewayId}

	return akas, nil
}

//// TRANSFORM FUNCTIONS

func getVpcInternetGatewayTurbotData(_ context.Context, d *transform.TransformData) (interface{}, error) {
	internetGateway := d.HydrateItem.(kaytu.EC2InternetGateway).Description.InternetGateway
	param := d.Param.(string)

	// Get resource title
	title := internetGateway.InternetGatewayId

	// Get the resource tags
	var turbotTagsMap map[string]string
	if internetGateway.Tags != nil {
		turbotTagsMap = map[string]string{}
		for _, i := range internetGateway.Tags {
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
