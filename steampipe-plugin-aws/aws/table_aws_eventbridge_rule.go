package aws

import (
	"context"

	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAwsEventBridgeRule(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_eventbridge_rule",
		Description: "AWS EventBridge Rule",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("name"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"ResourceNotFoundException", "ValidationException"}),
			},
			Hydrate: kaytu.GetEventBridgeRule,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListEventBridgeRule,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "event_bus_name", Require: plugin.Optional},
				{Name: "name_prefix", Require: plugin.Optional},
			},
		},

		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the rule.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Rule.Name")},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the rule.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Rule.Arn")},
			{
				Name:        "description",
				Description: "The description of the rule.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Rule.Description")},
			{
				Name:        "state",
				Description: "The state of the rule.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Rule.State")},
			{
				Name:        "event_bus_name",
				Description: "The name or ARN of the event bus associated with the rule.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Rule.EventBusName")},
			{
				Name:        "created_by",
				Description: "The account ID of the user that created the rule.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Rule.CreatedBy")},
			{
				Name:        "managed_by",
				Description: "If this is a managed rule, created by an AWS service on your behalf, this field displays the principal name of the AWS service that created the rule.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Rule.ManagedBy")},
			{
				Name:        "event_pattern",
				Description: "The event pattern of the rule.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Rule.EventPattern")},
			{
				Name:        "name_prefix",
				Description: "Specifying this limits the results to only those event rules with names that start with the specified prefix.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Rule.ScheduleExpression")},
			{
				Name:        "targets",
				Description: "The targets assigned to the rule.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Targets")},
			{
				Name:        "tags_src",
				Description: "A list of tags assigned to the rule.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Tags")},

			// Standard columns for all tables
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Rule.Name")},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Tags").Transform(eventBridgeTagListToTurbotTags),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Rule.Arn").Transform(arnToAkas),
			},
		}),
	}
}

//// TRANSFORM FUNCTIONS

func eventBridgeTagListToTurbotTags(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	plugin.Logger(ctx).Trace("eventBridgeTagListToTurbotTags")
	tagList := d.HydrateItem.(kaytu.EventBridgeRule).Description.Tags

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
