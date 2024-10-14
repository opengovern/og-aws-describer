package aws

import (
	"context"
	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAwsWAFRule(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_waf_rule",
		Description: "AWS WAF Rule",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("rule_id"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"WAFNonexistentItemException"}),
			},
			Hydrate: opengovernance.GetWAFRule,
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListWAFRule,
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
				Description: "The name of the metric for the Rule.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Rule.MetricName")},
			{
				Name:        "predicates",
				Description: "The Predicates object contains one Predicate element for each ByteMatchSet,IPSet, or SqlInjectionMatchSet object that you want to include in a Rule.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Rule.Predicates")},
			{
				Name:        "tags_src",
				Description: "A list of tags assigned to the Rule.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Tags")},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Rule.Name")},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(wafRuleTagListToTurbotTags),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("ARN").Transform(arnToAkas)},
		}),
	}
}

func wafRuleTagListToTurbotTags(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	plugin.Logger(ctx).Trace("tagListToTurbotTags")
	tagList := d.HydrateItem.(opengovernance.WAFRule).Description.Tags

	var turbotTagsMap map[string]string
	if tagList != nil {
		turbotTagsMap = map[string]string{}
		for _, i := range tagList {
			turbotTagsMap[*i.Key] = *i.Value
		}
	}

	return turbotTagsMap, nil
}
