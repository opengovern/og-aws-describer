package aws

import (
	"context"
	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAwsCloudFrontStreamingDistribution(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_cloudfront_streamingdistribution",
		Description: "AWS CloudFront StreamingDistribution",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("arn"), //TODO: change this to the primary key columns in model.go
			Hydrate:    kaytu.GetCloudFrontStreamingDistribution,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListCloudFrontStreamingDistribution,
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "id",
				Description: "The id of the streamingdistribution.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.StreamingDistribution.Id")},
			{
				Name:        "name",
				Description: "The name of the streamingdistribution.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.StreamingDistribution.Name")},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the streamingdistribution",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.StreamingDistribution.ARN")}, // or generate it below
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.StreamingDistribution.Name")},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.StreamingDistribution.Tags"), // probably needs a transform function
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.StreamingDistribution.ARN").Transform(arnToAkas), // or generate it below (keep the Transform(arnToTurbotAkas) or use Transform(transform.EnsureStringArray))
			},
		}),
	}
}

//// TRANSFORM FUNCTIONS
