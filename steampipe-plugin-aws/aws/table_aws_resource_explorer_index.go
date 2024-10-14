package aws

import (
	"context"
	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

//// TABLE DEFINITION

func tableAWSResourceExplorerIndex(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_resource_explorer_index",
		Description: "AWS Resource Explorer Index",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListResourceExplorer2Index,
			IgnoreConfig: &plugin.IgnoreConfig{
				// ValidationException error thrown for below cases in the table
				// 1. Type of the index type passed as input is not a valid value
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"ValidationException"}),
			},
			KeyColumns: plugin.KeyColumnSlice{
				{Name: "type", Require: plugin.Optional},
				{Name: "region", Require: plugin.Optional},
			},
		},
		Columns: awsKaytuDefaultColumns([]*plugin.Column{
			{
				Name:        "arn",
				Description: "The Amazon resource name (ARN) of the index.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Index.Arn")},
			{
				Name:        "type",
				Description: "The type of index. It can be one of the following values: LOCAL, AGGREGATOR.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Index.Type")},
			{
				Name:        "region",
				Description: "The AWS Region in which the index exists.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Index.Region")},
		}),
	}
}
