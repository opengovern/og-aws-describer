package aws

import (
	"context"

	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAwsEventBridgeBus(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_eventbridge_bus",
		Description: "AWS EventBridge Bus",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("arn"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"InvalidParameter", "ResourceNotFoundException", "ValidationException"}),
			},
			Hydrate: kaytu.GetEventBridgeBus,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListEventBridgeBus,
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the event bus.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Bus.Name")},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the account permitted to write events to the current account.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Bus.Arn")},
			{
				Name:        "policy",
				Description: "The policy that enables the external account to send events to your account.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Bus.Policy")},
			{
				Name:        "policy_std",
				Description: "Contains the policy that enables the external account to send events to your account in a canonical form for easier searching.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Bus.Policy").Transform(unescape).Transform(policyToCanonical),
			},
			{
				Name:        "tags_src",
				Description: "A list of tags assigned to the bus.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Tags")},

			// Standard columns for all tables
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Bus.Name")},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Tags").Transform(eventBridgeBusTagListToTurbotTags),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Bus.Arn").Transform(transform.EnsureStringArray),
			},
		}),
	}
}

//// TRANSFORM FUNCTION

func eventBridgeBusTagListToTurbotTags(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	plugin.Logger(ctx).Trace("eventBridgeBusTagListToTurbotTags")
	tagList := d.HydrateItem.(kaytu.EventBridgeBus).Description.Tags

	if tagList == nil {
		return nil, nil
	}

	// Mapping the resource tags inside turbotTags
	var turbotTagsMap map[string]string
	if tagList != nil {
		turbotTagsMap = map[string]string{}
		for _, i := range tagList {
			turbotTagsMap[*i.Key] = *i.Value
		}
	}

	return turbotTagsMap, nil
}
