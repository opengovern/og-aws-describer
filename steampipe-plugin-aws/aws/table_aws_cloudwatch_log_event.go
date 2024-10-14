package aws

import (
	"context"
	"encoding/json"
	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"
	"strings"

	"github.com/turbot/go-kit/types"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAwsCloudwatchLogEventListKeyColumns() []*plugin.KeyColumn {
	return []*plugin.KeyColumn{
		{Name: "log_group_name"},
		{Name: "log_stream_name", Require: plugin.Optional},
		{Name: "filter", Require: plugin.Optional, CacheMatch: "exact"},
		{Name: "region", Require: plugin.Optional},
		{Name: "timestamp", Operators: []string{">", ">=", "=", "<", "<="}, Require: plugin.Optional},
	}
}

func tableAwsCloudwatchLogEvent(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_cloudwatch_log_event",
		Description: "AWS CloudWatch Log Event",
		List: &plugin.ListConfig{
			Hydrate:    opengovernance.ListCloudWatchLogEvent,
			KeyColumns: tableAwsCloudwatchLogEventListKeyColumns(),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"ResourceNotFoundException"}),
			},
		},

		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			// Top columns
			{
				Name:        "log_group_name",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.LogGroupName"),
				Description: "The name of the log group to which this event belongs.",
			},
			{
				Name:        "log_stream_name",
				Type:        proto.ColumnType_STRING,
				Description: "The name of the log stream to which this event belongs.",
				Transform:   transform.FromField("Description.LogEvent.LogStreamName")},
			{
				Name:        "event_id",
				Type:        proto.ColumnType_STRING,
				Description: "The ID of the event.",
				Transform:   transform.FromField("Description.LogEvent.EventId")},
			{
				Name:      "timestamp",
				Type:      proto.ColumnType_TIMESTAMP,
				Transform: transform.FromField("Description.LogEvent.Timestamp").Transform(transform.UnixMsToTimestamp), Description: "The time when the event occurred.",
			},
			{
				Name:      "ingestion_time",
				Type:      proto.ColumnType_TIMESTAMP,
				Transform: transform.FromField("Description.LogEvent.IngestionTime").Transform(transform.UnixMsToTimestamp), Description: "The time when the event was ingested.",
			},
			{
				Name:        "filter",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromQual("filter"),
				Description: "Filter pattern for the search.",
			},
			{
				Name:      "message",
				Type:      proto.ColumnType_STRING,
				Transform: transform.FromField("Description.LogEvent.Message").Transform(trim), Description: "The data contained in the log event.",
			},
			{
				Name:      "message_json",
				Type:      proto.ColumnType_JSON,
				Transform: transform.FromField("Description.LogEvent.Message").Transform(trim).Transform(cloudwatchLogsMesssageJson), Description: "The data contained in the log event in json format. Only if data is valid json string.",
			},
		}),
	}
}

func trim(_ context.Context, d *transform.TransformData) (interface{}, error) {
	if d.Value == nil {
		return nil, nil
	}
	s := types.ToString(d.Value)
	return strings.TrimSpace(s), nil
}

//// TRANSFORM FUNCTIONS

func cloudwatchLogsMesssageJson(_ context.Context, d *transform.TransformData) (interface{}, error) {
	event := d.HydrateItem.(opengovernance.CloudWatchLogEvent).Description.LogEvent
	var eventMessage interface{}
	err := json.Unmarshal([]byte(*event.Message), &eventMessage)
	if err != nil {
		return nil, nil
	}
	return eventMessage, nil
}
