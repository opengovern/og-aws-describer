package aws

import (
	"context"
	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"

	"github.com/aws/aws-sdk-go-v2/service/wafregional"
	"github.com/aws/aws-sdk-go-v2/service/wafregional/types"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsWafRegionalRuleGroup(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_wafregional_rule_group",
		Description: "AWS WAF Regional Rule Group",
		Get: &plugin.GetConfig{
			//KeyColumns: plugin.AllColumns([]string{"rule_group_id"}),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"NonexistentItemException", "WAFNonexistentItemException"}),
			},
			Hydrate: kaytu.GetWAFRegionalRuleGroup,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListWAFRegionalRuleGroup,
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the rule group.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.RuleGroup.Name")},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the entity.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ARN")},
			{
				Name:        "rule_group_id",
				Description: "A unique identifier for the rule group.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.RuleGroup.RuleGroupId"),
			},
			{
				Name:        "metric_name",
				Description: "A friendly name or description for the metrics for this RuleGroup.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.RuleGroup.MetricName"),
			},
			{
				Name:        "activated_rules",
				Description: "A list of activated rules associated with the resource.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ActivatedRules"),
			},
			{
				Name:        "tags_src",
				Description: "A list of tags associated with the resource.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Tags"),
			},

			// steampipe standard columns
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.RuleGroup.Name"),
			},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Tags"),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ARN"),
			},
		}),
	}
}

//// LIST FUNCTION

func listWafRegionalRuleGroups(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// Create session
	svc, err := WAFRegionalClient(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("aws_wafregional_rule_group.listWafRegionalRuleGroups", "client_error", err)
		return nil, err
	}

	// Unsupported region, return no data
	if svc == nil {
		return nil, nil
	}

	// List all rule groups
	pagesLeft := true
	params := &wafregional.ListRuleGroupsInput{
		Limit: int32(100),
	}

	// Reduce the basic request limit down if the user has only requested a small number of rows
	limit := d.QueryContext.Limit
	if d.QueryContext.Limit != nil {
		if *limit < int64(params.Limit) {
			params.Limit = int32(*limit)
		}
	}

	for pagesLeft {
		response, err := svc.ListRuleGroups(ctx, params)
		if err != nil {
			plugin.Logger(ctx).Error("aws_wafregional_rule_group.listWafRegionalRuleGroups", "api_error", err)
			return nil, err
		}

		for _, ruleGroups := range response.RuleGroups {
			d.StreamListItem(ctx, ruleGroups)

			// Context may get cancelled due to manual cancellation or if the limit has been reached
			if d.RowsRemaining(ctx) == 0 {
				return nil, nil
			}
		}

		if response.NextMarker != nil {
			pagesLeft = true
			params.NextMarker = response.NextMarker
		} else {
			pagesLeft = false
		}
	}

	return nil, nil
}

//// HYDRATE FUNCTIONS

func getWafRegionalRuleGroup(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	var id string
	if h.Item != nil {
		dataMap, err := getWafRegionalRuleGroupData(ctx, d, h)
		data := dataMap.(map[string]string)
		if err != nil {
			return err, nil
		}
		id = data["rule_group_id"]
	} else {
		id = d.EqualsQuals["rule_group_id"].GetStringValue()
		if id == "" {
			return nil, nil
		}
	}

	// Create session
	svc, err := WAFRegionalClient(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("aws_wafregional_rule_group.getWafRegionalRuleGroup", "client_error", err)
		return nil, err
	}

	// Unsupported region, return no data
	if svc == nil {
		return nil, nil
	}

	params := &wafregional.GetRuleGroupInput{
		RuleGroupId: &id,
	}

	op, err := svc.GetRuleGroup(ctx, params)
	if err != nil {
		plugin.Logger(ctx).Error("aws_wafregional_rule_group.getWafRegionalRuleGroup", "api_error", err)
		return nil, err
	}

	return op.RuleGroup, nil
}

func getWafRegionalRuleGroupData(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	data := map[string]string{}

	commonData, err := getCommonColumns(ctx, d, h)
	if err != nil {
		return nil, err
	}
	commonColumnData := commonData.(*awsCommonColumnData)
	region := d.EqualsQualString(matrixKeyRegion)
	switch item := h.Item.(type) {
	case *types.RuleGroup:
		data["rule_group_id"] = *item.RuleGroupId
		data["rule_group_arn"] = "arn:" + commonColumnData.Partition + ":waf-regional:" + region + ":" + commonColumnData.AccountId + ":rulegroup/" + *item.RuleGroupId

	case types.RuleGroupSummary:
		data["rule_group_id"] = *item.RuleGroupId
		data["rule_group_arn"] = "arn:" + commonColumnData.Partition + ":waf-regional:" + region + ":" + commonColumnData.AccountId + ":rulegroup/" + *item.RuleGroupId
	}

	return data, nil
}
