package aws

import (
	"context"
	"strings"

	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsWafv2RuleGroup(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_wafv2_rule_group",
		Description: "AWS WAFv2 Rule Group",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"id", "name", "scope"}),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"WAFInvalidParameterException", "WAFNonexistentItemException", "ValidationException", "InvalidParameter"}),
			},
			Hydrate: kaytu.GetWAFv2RuleGroup,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListWAFv2RuleGroup,
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the rule group.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.RuleGroup.Name")},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the entity.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.RuleGroup.ARN")},
			{
				Name:        "id",
				Description: "A unique identifier for the rule group.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.RuleGroup.Id"),
			},
			{
				Name:        "scope",
				Description: "Specifies the scope of the rule group. Possible values are: 'REGIONAL' and 'CLOUDFRONT'.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.From(ruleGroupLocation),
			},
			{
				Name:        "capacity",
				Description: "The web ACL capacity units (WCUs) required for this rule group.",
				Type:        proto.ColumnType_INT,

				Transform: transform.FromField("Description.RuleGroup.Capacity")},
			{
				Name:        "description",
				Description: "A description of the rule group that helps with identification.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.RuleGroup.Description")},
			{
				Name:        "lock_token",
				Description: "A token used for optimistic locking.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.RuleGroupSummary.LockToken")},
			{
				Name:        "rules",
				Description: "The Rule statements used to identify the web requests that you want to allow, block, or count.",
				Type:        proto.ColumnType_JSON,

				Transform: transform.FromField("Description.RuleGroup.Rules")},
			{
				Name:        "visibility_config",
				Description: "Defines and enables Amazon CloudWatch metrics and web request sample collection.",
				Type:        proto.ColumnType_JSON,

				Transform: transform.FromField("Description.RuleGroup.VisibilityConfig")},
			{
				Name:        "tags_src",
				Description: "A list of tags associated with the resource.",
				Type:        proto.ColumnType_JSON,

				Transform: transform.FromField("Description.Tags.TagInfoForResource.TagList")},

			// steampipe standard columns

			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.RuleGroup.Name")},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,

				Transform: transform.FromField("Description.Tags").Transform(ruleGroupTagListToTurbotTags),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,

				Transform: transform.FromField("Description.RuleGroup.ARN").Transform(arnToAkas),

				// AWS standard columns
			},

			{
				Name:        "partition",
				Description: "The AWS partition in which the resource is located (aws, aws-cn, or aws-us-gov).",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Metadata.Partition")},
			{
				Name:        "region",
				Description: "The AWS Region in which the resource is located.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.From(ruleGroupRegion),
			},
			{
				Name:        "account_id",
				Description: "The AWS Account ID in which the resource is located.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Metadata.AccountID")},
		}),
	}
}

//// TRANSFORM FUNCTIONS

func ruleGroupLocation(_ context.Context, d *transform.TransformData) (interface{}, error) {
	data := d.HydrateItem.(kaytu.WAFv2RuleGroup)
	loc := strings.Split(strings.Split(*data.Description.RuleGroup.ARN, ":")[5], "/")[0]
	if loc == "regional" {
		return "REGIONAL", nil
	}
	return "CLOUDFRONT", nil
}

func ruleGroupTagListToTurbotTags(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	data := d.HydrateItem.(kaytu.WAFv2RuleGroup).Description.Tags

	if data.TagInfoForResource.TagList == nil || len(data.TagInfoForResource.TagList) < 1 {
		return nil, nil
	}

	// Mapping the resource tags inside turbotTags
	var turbotTagsMap map[string]string
	if data.TagInfoForResource.TagList != nil {
		turbotTagsMap = map[string]string{}
		for _, i := range data.TagInfoForResource.TagList {
			turbotTagsMap[*i.Key] = *i.Value
		}
	}

	return turbotTagsMap, nil
}

func ruleGroupRegion(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	data := d.HydrateItem.(kaytu.WAFv2RuleGroup)
	loc := strings.Split(strings.Split(*data.Description.RuleGroup.ARN, ":")[5], "/")[0]

	if loc == "global" {
		return "global", nil
	}
	return data.Metadata.Region, nil
}
