package aws

import (
	"context"
	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsWafRuleGroup(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_waf_rule_group",
		Description: "AWS WAF Rule Group",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"rule_group_id"}),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"NonexistentItemException", "WAFNonexistentItemException"}),
			},
			Hydrate: kaytu.GetWAFRuleGroup,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListWAFRuleGroup,
		},
		Columns: awsKaytuColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the rule group.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.RuleGroupSummary.Name")},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the entity.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.ARN")},
			{
				Name:        "rule_group_id",
				Description: "A unique identifier for the rule group.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.RuleGroupSummary.RuleGroupId")},
			{
				Name:        "metric_name",
				Description: "A friendly name or description for the metrics for this RuleGroup.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.RuleGroup.RuleGroup.MetricName")},
			{
				Name:        "activated_rules",
				Description: "A list of activated rules associated with the resource.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ActivatedRules")},
			{
				Name:        "tags_src",
				Description: "A list of tags associated with the resource.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Tags")},

			// steampipe standard columns

			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.RuleGroupSummary.Name")},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				//Transform:   transform.FromField("Description.Tags").Transform(classicRuleGroupTagListToTurbotTags),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ARN").Transform(transform.EnsureStringArray),
			},
		}),
	}
}

//
//func classicRuleGroupTagListToTurbotTags(ctx context.Context, d *transform.TransformData) (interface{}, error) {
//	data := d.HydrateItem.(kaytu.WAFRuleGroup).Description.Tags
//
//	if data == nil || len(data) < 1 {
//		return nil, nil
//	}
//
//	// Mapping the resource tags inside turbotTags
//	var turbotTagsMap map[string]string
//	if data != nil {
//		turbotTagsMap = map[string]string{}
//		for _, i := range data {
//			turbotTagsMap[*i.Key] = *i.Value
//		}
//	}
//
//	return turbotTagsMap, nil
//}
