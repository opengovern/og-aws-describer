package aws

import (
	"context"
	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAwsCloudFrontOriginAccessControl(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_cloudfront_origin_access_control",
		Description: "AWS CloudFront OriginAccessControl",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    opengovernance.GetCloudFrontOriginAccessControl,
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListCloudFrontOriginAccessControl,
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "id",
				Description: "The id of the origin access control.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.OriginAccessControl.Id")},
			{
				Name:        "name",
				Description: "The name of the origin access control.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.OriginAccessControl.Name")},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the origin access control",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("ARN")},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.OriginAccessControl.Name")},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(getCloudFrontOriginAccessControlTurbotTags),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("ARN").Transform(arnToAkas),
			},
		}),
	}
}

//// TRANSFORM FUNCTIONS

func getCloudFrontOriginAccessControlTurbotTags(_ context.Context, d *transform.TransformData) (interface{}, error) {
	tags := d.HydrateItem.(opengovernance.CloudFrontOriginAccessControl).Description.Tags
	return cloudfrontV2TagsToMap(tags)
}
