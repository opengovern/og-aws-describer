package aws

import (
	"context"

	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsCloudFrontFunction(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_cloudfront_function",
		Description: "AWS CloudFront Function",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("name"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"NoSuchFunctionExists"}),
			},
			Hydrate: kaytu.GetCloudFrontFunction,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListCloudFrontFunction,
		},

		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the CloudFront function.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Function.FunctionSummary.Name")},
			{
				Name:        "arn",
				Description: "The version identifier for the current version of the CloudFront function.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Function.FunctionSummary.FunctionMetadata.FunctionARN"),
			},
			{
				Name:        "status",
				Description: "The status of the CloudFront function.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Function.FunctionSummary.Status")},
			{
				Name:        "e_tag",
				Description: "The version identifier for the current version of the CloudFront function.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Function.ETag")},
			{
				Name:        "function_config",
				Description: "Contains configuration information about a CloudFront function.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Function.FunctionSummary.FunctionConfig")},
			{
				Name:        "function_metadata",
				Description: "Contains metadata about a CloudFront function.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Function.FunctionSummary.FunctionMetadata")},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Function.FunctionSummary.Name")},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Function.FunctionSummary.FunctionMetadata.FunctionARN").Transform(transform.EnsureStringArray),
			},
		}),
	}
}
