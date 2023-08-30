package aws

import (
	"context"
	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAwsWAFRegionalRule(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_wafregional_rule",
		Description: "AWS WAF Regional Rule",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("rule_id"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"WAFNonexistentItemException"}),
			},
			Hydrate: kaytu.GetWAFRegionalRule,
		},
		List: &plugin.ListConfig{
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"WAFNonexistentItemException"}),
			},
			Hydrate: kaytu.ListWAFRegionalRule,
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "rule_id",
				Description: "The id of the rule.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Rule.RuleId")},
			{
				Name:        "name",
				Description: "The name of the rule.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Rule.Name")},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the rule",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("ARN")},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Rule.Name")},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(getWAFRegionalRuleTurbotTags),
			},
			{
				Name:        "metric_name",
				Description: "A friendly name or description for the metrics for this Rule.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Rule.MetricName")},
			{
				Name:        "predicates",
				Description: "The Predicates object contains one Predicate element for each ByteMatchSet,IPSet, or SqlInjectionMatchSet object that you want to include in a Rule.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Rule.Predicates")},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("ARN").Transform(arnToAkas),
			},
		}),
	}
}

//// TRANSFORM FUNCTIONS

func getWAFRegionalRuleTurbotTags(_ context.Context, d *transform.TransformData) (interface{}, error) {
	tags := d.HydrateItem.(kaytu.WAFRegionalRule).Description.Tags
	return wafRegionalV2TagsToMap(tags)
}
