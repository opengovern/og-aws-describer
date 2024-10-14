package aws

import (
	"context"
	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsCloudwatchLogStream(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_cloudwatch_log_stream",
		Description: "AWS CloudWatch Log Stream",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"name"}),
			Hydrate:    opengovernance.GetCloudWatchLogStream,
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListCloudWatchLogStream,
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
				Description: "The name of the log stream.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.LogStream.LogStreamName")},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the log stream.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.LogStream.Arn"),
			},
			{
				Name:        "log_group_name",
				Description: "The name of the log group, in which the log stream belongs.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.LogGroupName")},
			{
				Name:        "creation_time",
				Description: "The creation time of the log stream.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.LogStream.CreationTime").Transform(transform.UnixMsToTimestamp),
			},
			{
				Name:        "first_event_timestamp",
				Description: "The time of the first event.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.LogStream.FirstEventTimestamp").Transform(transform.UnixMsToTimestamp),
			},
			{
				Name:        "last_event_timestamp",
				Description: "The time of the most recent log event in the log stream in CloudWatch Logs.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.LogStream.LastEventTimestamp").Transform(transform.UnixMsToTimestamp),
			},
			{
				Name:        "last_ingestion_time",
				Description: "Specifies the last log ingestion time.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.LogStream.LastIngestionTime").Transform(transform.UnixMsToTimestamp),
			},
			{
				Name:        "upload_sequence_token",
				Description: "Specifies the log upload sequence token.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.LogStream.UploadSequenceToken")},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.LogStream.LogStreamName")},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.LogStream.Arn").Transform(arnToAkas),
			},
		}),
	}
}
