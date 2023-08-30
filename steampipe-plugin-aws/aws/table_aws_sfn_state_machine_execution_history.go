package aws

import (
	"context"

	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAwsStepFunctionsStateMachineExecutionHistory(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_sfn_state_machine_execution_history",
		Description: "AWS Step Functions State Machine Execution History",
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListStepFunctionsStateMachineExecutionHistories,
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "id",
				Description: "The id of the event.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ExecutionHistory.Id")},
			{
				Name:        "execution_arn",
				Description: "The Amazon Resource Name (ARN) that identifies the execution.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ARN")},
			{
				Name:        "previous_event_id",
				Description: "The id of the previous event.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ExecutionHistory.PreviousEventId")},
			{
				Name:        "timestamp",
				Description: "The date and time the event occurred.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.ExecutionHistory.Timestamp")},
			{
				Name:        "type",
				Description: "The type of the event.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ExecutionHistory.Type")},
			{
				Name:        "activity_failed_event_details",
				Description: "Contains details about an activity that failed during an execution.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ExecutionHistory.ActivityFailedEventDetails")},
			{
				Name:        "activity_schedule_failed_event_details",
				Description: "Contains details about an activity schedule event that failed during an execution.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ExecutionHistory.ActivityScheduleFailedEventDetails")},
			{
				Name:        "activity_scheduled_event_details",
				Description: "Contains details about an activity scheduled during an execution.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ExecutionHistory.ActivityScheduledEventDetails")},
			{
				Name:        "activity_started_event_details",
				Description: "Contains details about the start of an activity during an execution.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ExecutionHistory.ActivityStartedEventDetails")},
			{
				Name:        "activity_succeeded_event_details",
				Description: "Contains details about an activity that successfully terminated during an execution.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ExecutionHistory.ActivitySucceededEventDetails")},
			{
				Name:        "activity_timed_out_event_details",
				Description: "Contains details about an activity timeout that occurred during an execution.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ExecutionHistory.ActivityTimedOutEventDetails")},
			{
				Name:        "execution_aborted_event_details",
				Description: "Contains details about an abort of an execution.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ExecutionHistory.ExecutionAbortedEventDetails")},
			{
				Name:        "execution_failed_event_details",
				Description: "Contains details about an execution failure event.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ExecutionHistory.ExecutionFailedEventDetails")},
			{
				Name:        "execution_started_event_details",
				Description: "Contains details about the start of the execution.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ExecutionHistory.ExecutionStartedEventDetails")},
			{
				Name:        "execution_succeeded_event_details",
				Description: "Contains details about the successful termination of the execution.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ExecutionHistory.ExecutionSucceededEventDetails")},
			{
				Name:        "execution_timed_out_event_details",
				Description: "Contains details about the execution timeout that occurred during the execution.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ExecutionHistory.ExecutionTimedOutEventDetails")},
			{
				Name:        "lambda_function_failed_event_details",
				Description: "Contains details about a lambda function that failed during an execution.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ExecutionHistory.LambdaFunctionFailedEventDetails")},
			{
				Name:        "lambda_function_schedule_failed_event_details",
				Description: "Contains details about a failed lambda function schedule event that occurred during an execution.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ExecutionHistory.LambdaFunctionScheduleFailedEventDetails")},
			{
				Name:        "lambda_function_scheduled_event_details",
				Description: "Contains details about a lambda function scheduled during an execution.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ExecutionHistory.LambdaFunctionScheduledEventDetails")},
			{
				Name:        "lambda_function_start_failed_event_details",
				Description: "Contains details about a lambda function that failed to start during an execution.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ExecutionHistory.LambdaFunctionStartFailedEventDetails")},
			{
				Name:        "lambda_function_succeeded_event_details",
				Description: "Contains details about a lambda function that terminated successfully during an execution.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ExecutionHistory.LambdaFunctionSucceededEventDetails")},
			{
				Name:        "lambda_function_timed_out_event_details",
				Description: "Contains details about a lambda function timeout that occurred during an execution.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ExecutionHistory.LambdaFunctionTimedOutEventDetails")},
			{
				Name:        "map_iteration_aborted_event_details",
				Description: "Contains details about an iteration of a Map state that was aborted.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ExecutionHistory.MapIterationAbortedEventDetails")},
			{
				Name:        "map_iteration_failed_event_details",
				Description: "Contains details about an iteration of a Map state that failed.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ExecutionHistory.MapIterationFailedEventDetails")},
			{
				Name:        "map_iteration_started_event_details",
				Description: "Contains details about an iteration of a Map state that was started.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ExecutionHistory.MapIterationStartedEventDetails")},
			{
				Name:        "map_iteration_succeeded_event_details",
				Description: "Contains details about an iteration of a Map state that succeeded.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ExecutionHistory.MapIterationSucceededEventDetails")},
			{
				Name:        "map_state_started_event_details",
				Description: "Contains details about Map state that was started.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ExecutionHistory.MapStateStartedEventDetails")},
			{
				Name:        "state_entered_event_details",
				Description: "Contains details about a state entered during an execution.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ExecutionHistory.StateEnteredEventDetails")},
			{
				Name:        "state_exited_event_details",
				Description: "Contains details about an exit from a state during an execution.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ExecutionHistory.StateExitedEventDetails")},
			{
				Name:        "task_failed_event_details",
				Description: "Contains details about the failure of a task.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ExecutionHistory.TaskFailedEventDetails")},
			{
				Name:        "task_scheduled_event_details",
				Description: "Contains details about a task that was scheduled.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ExecutionHistory.TaskScheduledEventDetails")},
			{
				Name:        "task_start_failed_event_details",
				Description: "Contains details about a task that failed to start.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ExecutionHistory.TaskStartFailedEventDetails")},
			{
				Name:        "task_started_event_details",
				Description: "Contains details about a task that was started.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ExecutionHistory.TaskStartedEventDetails")},
			{
				Name:        "task_submit_failed_event_details",
				Description: "Contains details about a task that where the submit failed.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ExecutionHistory.TaskSubmitFailedEventDetails")},
			{
				Name:        "task_submitted_event_details",
				Description: "Contains details about a submitted task.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ExecutionHistory.TaskSubmittedEventDetails")},
			{
				Name:        "task_succeeded_event_details",
				Description: "Contains details about a task that succeeded.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ExecutionHistory.TaskSucceededEventDetails")},
			{
				Name:        "task_timed_out_event_details",
				Description: "Contains details about a task that timed out.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ExecutionHistory.TaskTimedOutEventDetails")},

			// Standard columns for all tables

			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ExecutionHistory.Id")},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ARN")},
		}),
	}
}
