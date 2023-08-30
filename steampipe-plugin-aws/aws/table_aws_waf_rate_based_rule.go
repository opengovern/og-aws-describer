package aws

import (
	"context"

	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsWafRateBasedRule(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_waf_rate_based_rule",
		Description: "AWS WAF Rate Based Rule",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("rule_id"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"WAFNonexistentItemException", "ValidationException"}),
			},
			Hydrate: kaytu.GetWAFRateBasedRule,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListWAFRateBasedRule,
		},
		Columns: awsKaytuColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The name for the rule.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Rule.Name")},
			{
				Name:        "rule_id",
				Description: "The ID of the Rule.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Rule.RuleId")},
			{
				Name:        "metric_name",
				Description: "The name or description for the metrics for a RateBasedRule.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Rule.MetricName")},
			{
				Name:        "rate_key",
				Description: "The field that AWS WAF uses to determine if requests are likely arriving from single source and thus subject to rate monitoring.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Rule.RateKey")},
			{
				Name:        "rate_limit",
				Description: "The maximum number of requests, which have an identical value in the field specified by the RateKey, allowed in a five-minute period.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Rule.RateLimit")},
			{
				Name:        "predicates",
				Description: "The Predicates object contains one Predicate element for each ByteMatchSet, IPSet or SqlInjectionMatchSet object that you want to include in a RateBasedRule.",
				Type:        proto.ColumnType_JSON,

				Transform: transform.FromField("Description.Rule.MatchPredicates")},
			{
				Name:        "tags_src",
				Description: "A list of tags assigned to the Rule.",
				Type:        proto.ColumnType_JSON,

				Transform: transform.FromField("Description.Tags.TagList")},

			// Steampipe standard columns

			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Rule.Name")},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,

				Transform: transform.From(wafRateBasedRuletagListToTurbotTags),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ARN")},

			//// TRANSFORM FUNCTIONS
		}),
	}
}

func wafRateBasedRuletagListToTurbotTags(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	tagList := d.HydrateItem.(kaytu.WAFRateBasedRule).Description.Tags
	if tagList == nil {
		return nil, nil
	}

	// Mapping the resource tags inside turbotTags
	var turbotTagsMap map[string]string
	if tagList.TagList != nil {
		turbotTagsMap = map[string]string{}
		for _, i := range tagList.TagList {
			turbotTagsMap[*i.Key] = *i.Value
		}
	}
	return turbotTagsMap, nil
}
