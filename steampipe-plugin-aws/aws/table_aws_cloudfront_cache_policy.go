package aws

import (
	"context"

	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsCloudFrontCachePolicy(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_cloudfront_cache_policy",
		Description: "AWS CloudFront Cache Policy",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"NoSuchCachePolicy"}),
			},
			Hydrate: opengovernance.GetCloudFrontCachePolicy,
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListCloudFrontCachePolicy,
		},
		Columns: awsKaytuColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "A unique name to identify the cache policy.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.CachePolicy.CachePolicy.CachePolicyConfig.Name")},
			{
				Name:        "id",
				Description: "The unique identifier for the cache policy.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.CachePolicy.CachePolicy.Id"),
			},
			{
				Name:        "comment",
				Description: "A comment to describe the cache policy.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.CachePolicy.CachePolicy.CachePolicyConfig.Comment")},
			{
				Name:        "default_ttl",
				Description: "The default amount of time, in seconds, that you want objects to stay in the CloudFront cache before CloudFront sends another request to the origin to see if the object has been updated.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.CachePolicy.CachePolicy.CachePolicyConfig.DefaultTTL")},
			{
				Name:        "etag",
				Description: "The current version of the cache policy.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.CachePolicy.ETag")},
			{
				Name:        "max_ttl",
				Description: "The maximum amount of time, in seconds, that you want objects to stay in the CloudFront cache before CloudFront sends another request to the origin to see if the object has been updated.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.CachePolicy.CachePolicy.CachePolicyConfig.MaxTTL")},
			{
				Name:        "min_ttl",
				Description: "The minimum amount of time, in seconds, that you want objects to stay in the CloudFront cache before CloudFront sends another request to the origin to see if the object has been updated.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.CachePolicy.CachePolicy.CachePolicyConfig.MinTTL")},
			{
				Name:        "last_modified_time",
				Description: "The date and time when the cache policy was last modified.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.CachePolicy.CachePolicy.LastModifiedTime")},
			{
				Name:        "parameters_in_cache_key_and_forwarded_to_origin",
				Description: "The HTTP headers, cookies, and URL query strings to include in the cache key. The values included in the cache key are automatically included in requests that CloudFront sends to the origin.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.CachePolicy.CachePolicy.CachePolicyConfig.ParametersInCacheKeyAndForwardedToOrigin")},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.CachePolicy.CachePolicy.CachePolicyConfig.Name")},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("ARN").Transform(arnToAkas)},
		}),
	}
}
