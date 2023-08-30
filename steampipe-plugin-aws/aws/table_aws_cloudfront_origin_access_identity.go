package aws

import (
	"context"

	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsCloudFrontOriginAccessIdentity(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_cloudfront_origin_access_identity",
		Description: "AWS CloudFront Origin Access Identity",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"NoSuchCloudFrontOriginAccessIdentity"}),
			},
			Hydrate: kaytu.GetCloudFrontOriginAccessIdentity,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListCloudFrontOriginAccessIdentity,
		},
		Columns: awsKaytuColumns([]*plugin.Column{
			{
				Name:        "id",
				Description: "The ID for the origin access identity.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.OriginAccessIdentity.CloudFrontOriginAccessIdentity.Id"),
			},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) specifying the origin access identity.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("ARN").Transform(arnToAkas),
			},
			{
				Name:        "s3_canonical_user_id",
				Description: "The Amazon S3 canonical user ID for the origin access identity, which you use when giving the origin access identity read permission to an object in Amazon S3.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.OriginAccessIdentity.CloudFrontOriginAccessIdentity.S3CanonicalUserId")},
			{
				Name:        "caller_reference",
				Description: "A unique value that ensures that the request can't be replayed.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.OriginAccessIdentity.CloudFrontOriginAccessIdentity.CloudFrontOriginAccessIdentityConfig.CallerReference")},
			{
				Name:        "comment",
				Description: "The comment for this origin access identity.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.OriginAccessIdentity.CloudFrontOriginAccessIdentity.CloudFrontOriginAccessIdentityConfig.Comment")},
			{
				Name:        "etag",
				Description: "The current version of the origin access identity's information.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.OriginAccessIdentity.ETag")},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.OriginAccessIdentity.ResultMetadata")},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("ARN").Transform(transform.EnsureStringArray),
			},
		}),
	}
}
