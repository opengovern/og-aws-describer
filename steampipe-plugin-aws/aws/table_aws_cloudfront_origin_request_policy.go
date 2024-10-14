package aws

import (
	"context"

	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsCloudFrontOriginRequestPolicy(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_cloudfront_origin_request_policy",
		Description: "AWS CloudFront Origin Request Policy",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"NoSuchOriginRequestPolicy", "InvalidParameter"}),
			},
			Hydrate: opengovernance.GetCloudFrontOriginRequestPolicy,
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListCloudFrontOriginRequestPolicy,
		},
		Columns: awsKaytuColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "A unique name to identify the origin request policy.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.OriginRequestPolicy.OriginRequestPolicy.OriginRequestPolicyConfig.Name")},
			{
				Name:        "id",
				Description: "The ID for the origin request policy.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.OriginRequestPolicy.OriginRequestPolicy.Id"),
			},
			{
				Name:        "comment",
				Description: "The comment for this origin request policy.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.OriginRequestPolicy.OriginRequestPolicy.OriginRequestPolicyConfig.Comment")},
			{
				Name:        "etag",
				Description: "The current version of the origin request policy.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.OriginRequestPolicy.ETag")},
			{
				Name:        "last_modified_time",
				Description: "The date and time when the origin request policy was last modified.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.OriginRequestPolicy.OriginRequestPolicy.LastModifiedTime")},
			{
				Name:        "cookies_config",
				Description: "The cookies from viewer requests to include in origin requests.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.OriginRequestPolicy.OriginRequestPolicy.OriginRequestPolicyConfig.CookiesConfig")},
			{
				Name:        "headers_config",
				Description: "The HTTP headers to include in origin requests.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.OriginRequestPolicy.OriginRequestPolicy.OriginRequestPolicyConfig.HeadersConfig")},
			{
				Name:        "query_strings_config",
				Description: "The URL query strings from viewer requests to include in origin requests.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.OriginRequestPolicy.OriginRequestPolicy.OriginRequestPolicyConfig.QueryStringsConfig")},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.OriginRequestPolicy.OriginRequestPolicy.OriginRequestPolicyConfig.Name")},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("ARN").Transform(arnToAkas)},
		}),
	}
}
