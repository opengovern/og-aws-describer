package aws

import (
	"context"

	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsSecurityHubFindingAggregator(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_securityhub_finding_aggregator",
		Description: "AWS Security Hub Finding Aggregator",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("arn"),
			Hydrate:    opengovernance.GetSecurityHubFindingAggregator,
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListSecurityHubFindingAggregator,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"InvalidAccessException"}),
			},
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the finding aggregator.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.FindingAggregator.FindingAggregatorArn"),
			},
			{
				Name:        "finding_aggregation_region",
				Description: "The aggregation Region.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.FindingAggregator.FindingAggregationRegion")},
			{
				Name:        "region_linking_mode",
				Description: "Indicates whether to link all Regions, all Regions except for a list of excluded Regions, or a list of included Regions.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.FindingAggregator.RegionLinkingMode")},
			{
				Name:        "regions",
				Description: "The list of excluded Regions or included Regions.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.FindingAggregator.Regions")},

			/// Steampipe standard columns
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.FindingAggregator.FindingAggregatorArn").Transform(transform.EnsureStringArray),
			},
		}),
	}
}
