package aws

import (
	"context"

	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsSecurityHubInsight(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_securityhub_insight",
		Description: "AWS Securityhub Insight",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("arn"),
			Hydrate:    opengovernance.GetSecurityHubInsight,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"ResourceNotFoundException", "InvalidInputException"}),
			},
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListSecurityHubInsight,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"ResourceNotFoundException"}),
			},
		},

		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The name of a Security Hub insight.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Insight.Name")},
			{
				Name:        "arn",
				Description: "The ARN of a Security Hub insight.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Insight.InsightArn"),
			},
			{
				Name:        "group_by_attribute",
				Description: "The grouping attribute for the insight's findings. Indicates how to group the matching findings,and identifies the type of item that the insight applies to. For example, if an insight is grouped by resource identifier, then the insight produces a list of resource identifiers.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Insight.GroupByAttribute")},
			{
				Name:        "filters",
				Description: "One or more attributes used to filter the findings included in the insight. The insight only includes findings that match the criteria defined in the filters.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Insight.Filters")},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Insight.Name")},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Insight.InsightArn").Transform(transform.EnsureStringArray),
			},
		}),
	}
}
