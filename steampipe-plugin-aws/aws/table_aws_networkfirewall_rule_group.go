package aws

import (
	"context"

	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsNetworkFirewallRuleGroup(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_networkfirewall_rule_group",
		Description: "AWS Network Firewall Rule Group",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AnyColumn([]string{"arn", "rule_group_name"}),
			Hydrate:    kaytu.GetNetworkFirewallRuleGroup,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListNetworkFirewallRuleGroup,
		},

		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "rule_group_name",
				Description: "The descriptive name of the rule group.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.RuleGroupResponse.RuleGroupName")},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the rule group.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.RuleGroupResponse.RuleGroupArn"),
			},
			{
				Name:        "capacity",
				Description: "The maximum operating resources that this rule group can use. Rule group capacity is fixed at creation. When you update a rule group, you are limited to this capacity. When you reference a rule group from a firewall policy, Network Firewall reserves this capacity for the rule group.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.RuleGroupResponse.Capacity")},
			{
				Name:        "consumed_capacity",
				Description: "The number of capacity units currently consumed by the rule group rules.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.RuleGroupResponse.ConsumedCapacity")},
			{
				Name:        "description",
				Description: "A description of the rule group.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.RuleGroupResponse.Description")},
			{
				Name:        "number_of_associations",
				Description: "The number of firewall policies that use this rule group.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.RuleGroupResponse.NumberOfAssociations")},
			{
				Name:        "rule_group_id",
				Description: "The unique identifier for the rule group.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.RuleGroupResponse.RuleGroupId")},
			{
				Name:        "rule_group_status",
				Description: "Detailed information about the current status of a rule group.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.RuleGroupResponse.RuleGroupStatus")},
			{
				Name:        "rule_variables",
				Description: "Settings that are available for use in the rules in the rule group. You can only use these for stateful rule groups.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.RuleGroup.RuleVariables")},
			{
				Name:        "rules_source",
				Description: "The stateful rules or stateless rules for the rule group.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.RuleGroup.RulesSource")},
			{
				Name:        "stateful_rule_options",
				Description: "Additional options governing how Network Firewall handles the rule group. You can only use these for stateful rule groups.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.RuleGroup.StatefulRuleOptions")},
			{
				Name:        "type",
				Description: "Indicates whether the rule group is stateless or stateful. If the rule group is stateless, it contains stateless rules. If it is stateful, it contains stateful rules.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.RuleGroupResponse.Type")},
			{
				Name:        "tags_src",
				Description: "A list of tags assigned to the resource.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.RuleGroupResponse.Tags")},

			// Steampipe standard columns

			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.RuleGroupResponse.RuleGroupName")},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.RuleGroupResponse.Tags").Transform(networkFirewallRuleGroupTurbotTags),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.RuleGroup.RuleGroupArn").Transform(transform.EnsureStringArray),
			},
		}),
	}
}

func networkFirewallRuleGroupTurbotTags(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	ruleGroup := d.HydrateItem.(kaytu.NetworkFirewallRuleGroup).Description

	// Mapping the resource tags inside turbotTags
	var turbotTagsMap map[string]string
	if ruleGroup.RuleGroupResponse.Tags != nil {
		turbotTagsMap = map[string]string{}
		for _, i := range ruleGroup.RuleGroupResponse.Tags {
			turbotTagsMap[*i.Key] = *i.Value
		}
	}

	return turbotTagsMap, nil
}
