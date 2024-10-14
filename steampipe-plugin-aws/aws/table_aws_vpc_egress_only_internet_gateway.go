package aws

import (
	"context"
	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"

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
			Hydrate: opengovernance.GetEC2EgressOnlyInternetGateway,
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListEC2EgressOnlyInternetGateway,
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
				Transform:   transform.FromField("ARN").Transform(arnToAkas)},
		}),
	}
}

//// TRANSFORM FUNCTIONS

func egressOnlyIGWApiDataToTurbotData(_ context.Context, d *transform.TransformData) (interface{}, error) {
	egw := d.HydrateItem.(opengovernance.EC2EgressOnlyInternetGateway).Description.EgressOnlyInternetGateway
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
