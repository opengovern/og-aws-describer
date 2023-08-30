package aws

import (
	"context"

	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsWafWebAcl(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_waf_web_acl",
		Description: "AWS WAF Web ACL",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"web_acl_id"}),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"WAFNonexistentItemException", "WAFInvalidParameterException"}),
			},
			Hydrate: kaytu.GetWAFWebAcl,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListWAFWebAcl,
		},
		Columns: awsKaytuColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the Web ACL. You cannot change the name of a Web ACL after you create it.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.WebACL.Name")},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the entity.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.WebACL.WebACLArn")},
			{
				Name:        "web_acl_id",
				Description: "The unique identifier for the Web ACL.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.WebACL.WebACLId")},
			{
				Name:        "default_action",
				Description: "The action to perform if none of the Rules contained in the WebACL match.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.WebACL.DefaultAction.Type")},
			{
				Name:        "metric_name",
				Description: "A friendly name or description for the metrics for this WebACL.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.WebACL.MetricName")},
			{
				Name:        "logging_configuration",
				Description: "The logging configuration for the specified web ACL.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.LoggingConfiguration")},
			{
				Name:        "rules",
				Description: "The Rule statements used to identify the web requests that you want to allow, block, or count.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.WebACL.Rules")},
			{
				Name:        "tags_src",
				Description: "A list of tags associated with the resource.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Tags.TagList")},

			// Steampipe standard columns

			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.WebACL.Name")},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Tags.TagList").Transform(classicWebAclTagListToTurbotTags),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.WebACL.WebACLArn").Transform(transform.EnsureStringArray),
			},
		}),
	}
}

func classicWebAclTagListToTurbotTags(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	data := d.HydrateItem.(kaytu.WAFWebAcl).Description.Tags

	if data == nil {
		return nil, nil
	}

	if data.TagList == nil || len(data.TagList) < 1 {
		return nil, nil
	}

	// Mapping the resource tags inside turbotTags
	var turbotTagsMap map[string]string
	if data != nil {
		turbotTagsMap = map[string]string{}
		for _, i := range data.TagList {
			turbotTagsMap[*i.Key] = *i.Value
		}
	}

	return turbotTagsMap, nil
}
