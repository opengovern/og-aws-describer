package aws

import (
	"context"

	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsCloudFrontResponseHeadersPolicy(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_cloudfront_response_headers_policy",
		Description: "AWS Cloudfront Response Headers Policy",
		List: &plugin.ListConfig{
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:    "type",
					Require: plugin.Optional,
				},
			},
			Hydrate: opengovernance.ListCloudFrontResponseHeadersPolicy,
		},

		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the response headers policy.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ResponseHeadersPolicy.ResponseHeadersPolicy.ResponseHeadersPolicyConfig.Name")},
			{
				Name:        "id",
				Description: "The identifier for the response headers policy.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ResponseHeadersPolicy.ResponseHeadersPolicy.Id"),
			},
			{
				Name:        "arn",
				Description: "The version identifier for the current version of the response headers policy.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("ARN").Transform(arnToAkas),
			},
			{
				Name:        "last_modified_time",
				Description: "The date and time when the response headers policy was last modified.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.ResponseHeadersPolicy.ResponseHeadersPolicy.LastModifiedTime")},
			{
				Name:        "type",
				Description: "The type of response headers policy, either managed (created by AWS) or custom (created in this AWS account).",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ResponseHeadersPolicy.ResponseHeadersPolicy.ResponseHeadersPolicyConfig.SecurityHeadersConfig.ContentTypeOptions")},
			{
				Name:        "etag",
				Description: "The version identifier for the current version of the response headers policy.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ResponseHeadersPolicy.ETag")},
			{
				Name:        "response_headers_policy_config",
				Description: "A response headers policy contains information about a set of HTTP response headers and their values. CloudFront adds the headers in the policy to HTTP responses that it sends for requests that match a cache behavior that’s associated with the policy.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ResponseHeadersPolicy.ResponseHeadersPolicy.ResponseHeadersPolicyConfig")},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ResponseHeadersPolicy.ResponseHeadersPolicy.ResponseHeadersPolicyConfig.Name")},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("ARN").Transform(transform.EnsureStringArray),
			},
		}),
	}
}
