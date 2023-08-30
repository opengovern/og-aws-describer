package aws

import (
	"context"
	"fmt"

	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"
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
			Hydrate:    kaytu.GetCloudFrontOriginAccessControl,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListCloudFrontOriginAccessControl,
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
				Transform:   transform.From(getCloudFrontOriginAccessControlArn)},
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
				Transform:   transform.From(getCloudFrontOriginAccessControlArn).Transform(arnToAkas),
			},
		}),
	}
}

//// TRANSFORM FUNCTIONS

func getCloudFrontOriginAccessControlTurbotTags(_ context.Context, d *transform.TransformData) (interface{}, error) {
	tags := d.HydrateItem.(kaytu.CloudFrontOriginAccessControl).Description.Tags
	return cloudfrontV2TagsToMap(tags)
}

func getCloudFrontOriginAccessControlArn(_ context.Context, d *transform.TransformData) (interface{}, error) {
	originAccessControl := d.HydrateItem.(kaytu.CloudFrontOriginAccessControl).Description.OriginAccessControl
	metadata := d.HydrateItem.(kaytu.CloudFrontOriginAccessControl).Metadata

	arn := fmt.Sprintf("arn:%s:cloudfront::%s:origin-access-control/%s", metadata.Partition, metadata.AccountID, *originAccessControl.Id) //TODO: this is fake ARN, find out the real one's format
	return arn, nil
}
