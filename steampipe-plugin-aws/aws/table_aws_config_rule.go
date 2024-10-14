package aws

import (
	"context"

	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsConfigRule(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_config_rule",
		Description: "AWS Config Rule",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("name"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"NoSuchConfigRuleException", "ResourceNotFoundException", "ValidationException"}),
			},
			Hydrate: opengovernance.GetConfigRule,
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListConfigRule,
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:    "name",
					Require: plugin.Optional,
				},
			},
		},

		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The name that you assign to the AWS Config rule.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Rule.ConfigRuleName")},
			{
				Name:        "rule_id",
				Description: "The ID of the AWS Config rule.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Rule.ConfigRuleId")},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the AWS Config rule.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("ConfigRuleArn"),
			},
			{
				Name:        "rule_state",
				Description: "It indicate the evaluation status for the AWS Config rule.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Rule.ConfigRuleState")},
			{
				Name:        "created_by",
				Description: "Service principal name of the service that created the rule.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Rule.CreatedBy")},
			{
				Name:        "description",
				Description: "The description that you provide for the AWS Config rule.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Rule.Description")},
			{
				Name:        "maximum_execution_frequency",
				Description: "The maximum frequency with which AWS Config runs evaluations for a rule.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Rule.MaximumExecutionFrequency")},
			{
				Name:        "compliance_by_config_rule",
				Description: "The compliance information of the config rule.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Compliance")},
			{
				Name:        "evaluation_modes",
				Description: "The modes the Config rule can be evaluated in. The valid values are distinct objects. By default, the value is Detective evaluation mode only.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Rule.EvaluationModes")},
			{
				Name:        "input_parameters",
				Description: "A string, in JSON format, that is passed to the AWS Config rule Lambda function.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Rule.InputParameters")},
			{
				Name:        "scope",
				Description: "Defines which resources can trigger an evaluation for the rule. The scope can include one or more resource types, a combination of one resource type and one resource ID, or a combination of a tag key and value. Specify a scope to constrain the resources that can trigger an evaluation for the rule. If you do not specify a scope, evaluations are triggered when any resource in the recording group changes.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Rule.Scope")},
			{
				Name:        "source",
				Description: "Provides the rule owner (AWS or customer), the rule identifier, and the notifications that cause the function to evaluate your AWS resources.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Rule.Source")},
			{
				Name:        "tags_src",
				Description: "A list of tags assigned to the rule.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Tags")},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Rule.ConfigRuleName")},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(configRuleTurbotTags),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Rule.ConfigRuleArn").Transform(transform.EnsureStringArray),
			},
		}),
	}
}

//// TRANSFORM FUNCTIONS

func configRuleTurbotTags(_ context.Context, d *transform.TransformData) (interface{}, error) {
	data := d.HydrateItem.(opengovernance.ConfigRule).Description

	if data.Tags == nil {
		return nil, nil
	}

	// Mapping the resource tags inside turbotTags
	turbotTagsMap := map[string]string{}
	for _, i := range data.Tags {
		turbotTagsMap[*i.Key] = *i.Value
	}

	return turbotTagsMap, nil
}
