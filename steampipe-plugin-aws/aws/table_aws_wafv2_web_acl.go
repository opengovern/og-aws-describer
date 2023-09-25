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

func tableAwsWafv2WebAcl(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_wafv2_web_acl",
		Description: "AWS WAFv2 Web ACL",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"id", "name", "scope"}),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"WAFNonexistentItemException", "WAFInvalidParameterException"}),
			}, Hydrate: kaytu.GetWAFv2WebACL,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListWAFv2WebACL,
		},

		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the Web ACL. You cannot change the name of a Web ACL after you create it.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.WebACL.Name")},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the entity.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.WebACL.ARN"),
			},
			{
				Name:        "associated_resources",
				Description: "The array of Amazon Resource Names (ARNs) of the associated resources.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.AssociatedResources"),
			},
			{
				Name:        "id",
				Description: "The unique identifier for the Web ACL.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.WebACL.Id")},
			{
				Name:        "scope",
				Description: "Specifies the scope of the Web ACL. Possibles values are: 'REGIONAL' and 'CLOUDFRONT'.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.From(webAclLocation),
			},
			{
				Name:        "description",
				Description: "A description of the Web ACL that helps with identification.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.WebACL.Description")},
			{
				Name:        "capacity",
				Description: "The Web ACL capacity units(WCUs) currently being used by this resource.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.WebACL.Capacity")},
			{
				Name:        "lock_token",
				Description: "A token used for optimistic locking.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.LockToken")},
			{
				Name:        "managed_by_firewall_manager",
				Description: "Indicates whether this web ACL is managed by AWS Firewall Manager.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.WebACL.ManagedByFirewallManager")},
			{
				Name:        "default_action",
				Description: "The action to perform if none of the Rules contained in the Web ACL match.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.WebACL.DefaultAction")},
			{
				Name:        "logging_configuration",
				Description: "The logging configuration for the specified web ACL.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.LoggingConfiguration")},
			{
				Name:        "pre_process_firewall_manager_rule_groups",
				Description: "The first set of rules for AWS WAF to process in the web ACL.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.WebACL.PreProcessFirewallManagerRuleGroups")},
			{
				Name:        "post_process_firewall_manager_rule_groups",
				Description: "The last set of rules for AWS WAF to process in the web ACL.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.WebACL.PostProcessFirewallManagerRuleGroups")},
			{
				Name:        "rules",
				Description: "The Rule statements used to identify the web requests that you want to allow, block, or count.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.WebACL.Rules")},
			{
				Name:        "visibility_config",
				Description: "Defines and enables Amazon CloudWatch metrics and web request sample collection.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.WebACL.VisibilityConfig")},
			{
				Name:        "tags_src",
				Description: "A list of tags associated with the resource.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.TagInfoForResource.TagList")},
			// Steampipe standard columns

			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.WebACL.Name")},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.TagInfoForResource.TagList").Transform(webAclTagListToTurbotTags),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.WebACL.ARN").Transform(transform.EnsureStringArray),
			},
			// AWS standard columns
			{
				Name:        "region",
				Description: "The AWS Region in which the resource is located.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.From(webAclRegion),
			},
		}),
	}
}

//// LIST FUNCTION

//// HYDRATE FUNCTIONS

//// TRANSFORM FUNCTIONS

func webAclLocation(_ context.Context, d *transform.TransformData) (interface{}, error) {
	data := d.HydrateItem.(kaytu.WAFv2WebACL)
	loc := strings.Split(strings.Split(*data.Description.WebACL.ARN, ":")[5], "/")[0]
	if loc == "regional" {
		return "REGIONAL", nil
	}
	return "CLOUDFRONT", nil
}

func webAclTagListToTurbotTags(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	plugin.Logger(ctx).Trace("webAclTagListToTurbotTags")
	data := d.HydrateItem.(kaytu.WAFv2WebACL).Description

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

func webAclRegion(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	data := d.HydrateItem.(kaytu.WAFv2WebACL).Description
	loc := strings.Split(strings.Split(*data.WebACL.ARN, ":")[5], "/")[0]
	if loc == "global" {
		return "global", nil
	}
	region := d.HydrateItem.(kaytu.WAFv2WebACL).Metadata.Region
	return region, nil
}
