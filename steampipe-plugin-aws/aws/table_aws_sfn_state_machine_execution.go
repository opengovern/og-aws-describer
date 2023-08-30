package aws

import (
	"context"

	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAwsStepFunctionsStateMachineExecution(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_sfn_state_machine_execution",
		Description: "AWS Step Functions State Machine Execution",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("execution_arn"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"InvalidParameter", "ExecutionDoesNotExist", "InvalidArn"}),
			},
			Hydrate: kaytu.GetStepFunctionsStateMachineExecution,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListStepFunctionsStateMachineExecution,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "status", Require: plugin.Optional},
				{Name: "state_machine_arn", Require: plugin.Optional},
			},
		},

		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the execution.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Execution.Name")},
			{
				Name:        "execution_arn",
				Description: "The Amazon Resource Name (ARN) that identifies the execution.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Execution.ExecutionArn")},
			{
				Name:        "status",
				Description: "The current status of the execution.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Execution.Status")},
			{
				Name:        "input",
				Description: "The string that contains the JSON input data of the execution.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Execution.Input")},
			{
				Name:        "output",
				Description: "The JSON output data of the execution.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Execution.Output")},
			{
				Name:        "start_date",
				Description: "The date the execution started.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Execution.StartDate")},
			{
				Name:        "state_machine_arn",
				Description: "The Amazon Resource Name (ARN) of the executed state machine.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Execution.StateMachineArn")},
			{
				Name:        "stop_date",
				Description: "If the execution already ended, the date the execution stopped.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Execution.StopDate")},
			{
				Name:        "trace_header",
				Description: "The AWS X-Ray trace header that was passed to the execution.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Execution.TraceHeader")},
			{
				Name:        "input_details",
				Description: "Provides details about execution input or output.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Execution.InputDetails")},
			{
				Name:        "output_details",
				Description: "Provides details about execution input or output.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Execution.OutputDetails")},

			// Standard columns for all tables
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Execution.Name")},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Execution.ExecutionArn").Transform(transform.EnsureStringArray),
			},
		}),
	}
}
