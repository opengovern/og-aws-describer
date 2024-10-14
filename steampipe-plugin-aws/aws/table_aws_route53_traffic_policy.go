package aws

import (
	"context"

	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAwsRoute53TrafficPolicy(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_route53_traffic_policy",
		Description: "AWS Route53 Traffic Policy",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"id", "version"}),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"NoSuchTrafficPolicy"}),
			},
			Hydrate: opengovernance.GetRoute53TrafficPolicy,
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListRoute53TrafficPolicy,
		},
		Columns: awsKaytuColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The name that you specified when traffic policy was created.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.TrafficPolicy.Name")},
			{
				Name:        "id",
				Description: "The ID that Amazon Route 53 assigned to a traffic policy when it was created.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.TrafficPolicy.Id")},
			{
				Name:        "type",
				Description: "The DNS type of the resource record sets that Amazon Route 53 creates when you use a traffic policy to create a traffic policy instance.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.TrafficPolicy.Type")},
			{
				Name:        "version",
				Description: "The version number that Amazon Route 53 assigns to a traffic policy.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.TrafficPolicy.Version")},
			{
				Name:        "comment",
				Description: "The comment that you specified when traffic policy was created.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.TrafficPolicy.Comment")},
			{
				Name:        "document",
				Description: "The definition of a traffic policy in JSON format.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.TrafficPolicy.Document")},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.TrafficPolicy.Name")},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("ARN").Transform(arnToAkas)},
		}),
	}
}
