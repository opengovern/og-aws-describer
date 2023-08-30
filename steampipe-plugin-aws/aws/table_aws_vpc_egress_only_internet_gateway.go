package aws

import (
	"context"
	"fmt"
	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAwsVpcEgressOnlyIGW(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_vpc_egress_only_internet_gateway",
		Description: "AWS VPC Egress Only Internet Gateway",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"InvalidEgressOnlyInternetGatewayId.NotFound", "InvalidEgressOnlyInternetGatewayId.Malformed"}),
			},
			Hydrate: kaytu.GetEC2EgressOnlyInternetGateway,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListEC2EgressOnlyInternetGateway,
		},

		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "id",
				Description: "The ID of the egress-only internet gateway.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.EgressOnlyInternetGateway.EgressOnlyInternetGatewayId"),
			},
			{
				Name:        "attachments",
				Description: "Information about the attachment of the egress-only internet gateway.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.EgressOnlyInternetGateway.Attachments")},
			{
				Name:        "tags_src",
				Description: "A list of tags that are attached to egress only internet gateway.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.EgressOnlyInternetGateway.Tags")},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromP(egressOnlyIGWApiDataToTurbotData, "Tags"),
			},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromP(egressOnlyIGWApiDataToTurbotData, "Title"),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(getEc2EgressOnlyIGWArn).Transform(arnToAkas)},
		}),
	}
}

//// TRANSFORM FUNCTIONS

func getEc2EgressOnlyIGWArn(_ context.Context, d *transform.TransformData) (interface{}, error) {
	egressOnlyInternetGateway := d.HydrateItem.(kaytu.EC2EgressOnlyInternetGateway).Description.EgressOnlyInternetGateway
	metadata := d.HydrateItem.(kaytu.EC2EgressOnlyInternetGateway).Metadata

	arn := fmt.Sprintf("arn:%s:ec2:%s:%s:egress-only-internet-gateway/%s", metadata.Partition, metadata.Region, metadata.AccountID, *egressOnlyInternetGateway.EgressOnlyInternetGatewayId)
	return arn, nil
}

func egressOnlyIGWApiDataToTurbotData(_ context.Context, d *transform.TransformData) (interface{}, error) {
	egw := d.HydrateItem.(kaytu.EC2EgressOnlyInternetGateway).Description.EgressOnlyInternetGateway
	param := d.Param.(string)

	// Get resource title
	title := egw.EgressOnlyInternetGatewayId

	// Get the resource tags
	var turbotTagsMap map[string]string
	if egw.Tags != nil {
		turbotTagsMap = map[string]string{}
		for _, i := range egw.Tags {
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
