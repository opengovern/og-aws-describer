package aws

import (
	"context"

	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAwsStepFunctionsStateMachine(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_sfn_state_machine",
		Description: "AWS Step Functions State Machine",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("arn"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"ResourceNotFoundException", "StateMachineDoesNotExist", "InvalidArn"}),
			},
			Hydrate: kaytu.GetStepFunctionsStateMachine,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListStepFunctionsStateMachine,
		},

		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the state machine.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.StateMachine.Name")},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) that identifies the state machine.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.StateMachine.StateMachineArn"),
			},
			{
				Name:        "status",
				Description: "The current status of the state machine.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.StateMachine.Status")},
			{
				Name:        "type",
				Description: "The type of the state machine.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.StateMachine.Type")},
			{
				Name:        "creation_date",
				Description: "The date the state machine is created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.StateMachine.CreationDate")},
			{
				Name:        "definition",
				Description: "The Amazon States Language definition of the state machine.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.StateMachine.Definition")},
			{
				Name:        "role_arn",
				Description: "The Amazon Resource Name (ARN) of the IAM role used when creating this state machine.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.StateMachine.RoleArn")},
			{
				Name:        "logging_configuration",
				Description: "The LoggingConfiguration data type is used to set CloudWatch Logs options.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.StateMachine.LoggingConfiguration")},
			{
				Name:        "tags_src",
				Description: "The list of tags associated with the state machine.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Tags")},
			{
				Name:        "tracing_configuration",
				Description: "Selects whether AWS X-Ray tracing is enabled.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.StateMachine.TracingConfiguration")},

			// Standard columns for all tables
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(stateMachineTagsToTurbotTags),
			},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.StateMachine.Name")},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.StateMachine.StateMachineArn").Transform(transform.EnsureStringArray),
			},
		}),
	}
}

//// TRANSFORM FUNCTIONS

func stateMachineTagsToTurbotTags(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	tags := d.HydrateItem.(kaytu.StepFunctionsStateMachine).Description.Tags

	if tags == nil {
		return nil, nil
	}
	// Mapping the resource tags inside turbotTags
	var turbotTagsMap map[string]string
	if tags != nil {
		turbotTagsMap = map[string]string{}
		for _, i := range tags {
			turbotTagsMap[*i.Key] = *i.Value
		}
	}

	return turbotTagsMap, nil
}
