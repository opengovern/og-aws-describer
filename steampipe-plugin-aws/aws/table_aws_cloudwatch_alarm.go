package aws

import (
	"context"
	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsCloudWatchAlarm(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_cloudwatch_alarm",
		Description: "AWS CloudWatch Alarm",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("name"),
			Hydrate:    kaytu.GetCloudWatchAlarm,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListCloudWatchAlarm,
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:    "name",
					Require: plugin.Optional,
				},
				{
					Name:    "state_value",
					Require: plugin.Optional,
				},
			},
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the alarm.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.MetricAlarm.AlarmName"),
			},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the alarm.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.MetricAlarm.AlarmArn"),
			},
			{
				Name:        "state_value",
				Description: "The state value for the alarm.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.MetricAlarm.StateValue"),
			},
			{
				Name:        "actions_enabled",
				Description: "Indicates whether actions should be executed during any changes to the alarm state.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.MetricAlarm.ActionsEnabled"),
			},
			{
				Name:        "alarm_configuration_updated_timestamp",
				Description: "The time stamp of the last update to the alarm configuration.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.MetricAlarm.AlarmConfigurationUpdatedTimestamp"),
			},
			{
				Name:        "alarm_description",
				Description: "The description of the alarm.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.MetricAlarm.AlarmDescription"),
			},
			{
				Name:        "comparison_operator",
				Description: "The arithmetic operation to use when comparing the specified statistic and threshold. The specified statistic value is used as the first operand.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.MetricAlarm.ComparisonOperator"),
			},
			{
				Name:        "datapoints_to_alarm",
				Description: "The number of data points that must be breaching to trigger the alarm.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.MetricAlarm.DatapointsToAlarm"),
			},
			{
				Name:        "evaluate_low_sample_count_percentile",
				Description: "Used only for alarms based on percentiles.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.MetricAlarm.EvaluateLowSampleCountPercentile"),
			},
			{
				Name:        "evaluation_periods",
				Description: "The number of periods over which data is compared to the specified threshold.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.MetricAlarm.EvaluationPeriods"),
			},
			{
				Name:        "extended_statistic",
				Description: "The percentile statistic for the metric associated with the alarm. Specify a value between p0.0 and p100.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.MetricAlarm.ExtendedStatistic"),
			},
			{
				Name:        "metric_name",
				Description: "The name of the metric associated with the alarm, if this is an alarm based on a single metric.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.MetricAlarm.MetricName"),
			},
			{
				Name:        "namespace",
				Description: "The namespace of the metric associated with the alarm.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.MetricAlarm.Namespace"),
			},
			{
				Name:        "period",
				Description: "The period, in seconds, over which the statistic is applied.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.MetricAlarm.EvaluationPeriods"),
			},
			{
				Name:        "state_reason",
				Description: "An explanation for the alarm state, in text format.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.MetricAlarm.StateReason"),
			},
			{
				Name:        "state_reason_data",
				Description: "An explanation for the alarm state, in JSON format.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.MetricAlarm.StateReasonData"),
			},
			{
				Name:        "state_updated_timestamp",
				Description: "The time stamp of the last update to the alarm state.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.MetricAlarm.StateUpdatedTimestamp"),
			},
			{
				Name:        "statistic",
				Description: "The statistic for the metric associated with the alarm, other than percentile.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.MetricAlarm.Statistic"),
			},
			{
				Name:        "threshold",
				Description: "The value to compare with the specified statistic.",
				Type:        proto.ColumnType_DOUBLE,
				Transform:   transform.FromField("Description.MetricAlarm.Threshold"),
			},
			{
				Name:        "threshold_metric_id",
				Description: "In an alarm based on an anomaly detection model, this is the ID of the ANOMALY_DETECTION_BAND function used as the threshold for the alarm.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.MetricAlarm.ThresholdMetricId"),
			},
			{
				Name:        "treat_missing_data",
				Description: "Sets how this alarm is to handle missing data points. If this parameter is omitted, the default behavior of missing is used.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.MetricAlarm.TreatMissingData"),
			},
			{
				Name:        "unit",
				Description: "The unit of the metric associated with the alarm.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.MetricAlarm.Unit"),
			},
			{
				Name:        "alarm_actions",
				Description: "The actions to execute when this alarm transitions to the ALARM state from any other state. Each action is specified as an Amazon Resource Name (ARN).",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.MetricAlarm.AlarmActions"),
			},
			{
				Name:        "dimensions",
				Description: "The dimensions for the metric associated with the alarm.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.MetricAlarm.Dimensions"),
			},
			{
				Name:        "insufficient_data_actions",
				Description: "The actions to execute when this alarm transitions to the INSUFFICIENT_DATA state from any other state. Each action is specified as an Amazon Resource Name (ARN).",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.MetricAlarm.InsufficientDataActions"),
			},
			{
				Name:        "metrics",
				Description: "An array of MetricDataQuery structures, used in an alarm based on a metric math expression.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.MetricAlarm.Metrics"),
			},
			{
				Name:        "ok_actions",
				Description: "The actions to execute when this alarm transitions to the OK state from any other state. Each action is specified as an Amazon Resource Name (ARN).",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.MetricAlarm.OKActions"),
			},
			{
				Name:        "tags_src",
				Description: "The list of tag keys and values associated with alarm.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Tags"),
			},

			// Standard columns
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.MetricAlarm.AlarmName"),
			},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(getAwsCloudWatchAlarmTurbotTags),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.MetricAlarm.AlarmArn").Transform(arnToAkas),
			},
		}),
	}
}

//// LIST FUNCTION

//// HYDRATE FUNCTIONS

//// TRANSFORM FUNCTIONS

func getAwsCloudWatchAlarmTurbotTags(_ context.Context, d *transform.TransformData) (interface{}, error) {
	cloudWatchAlarm := d.HydrateItem.(kaytu.CloudWatchAlarm).Description

	if cloudWatchAlarm.Tags == nil {
		return nil, nil
	}

	turbotTagsMap := map[string]string{}
	for _, i := range cloudWatchAlarm.Tags {
		turbotTagsMap[*i.Key] = *i.Value
	}
	return turbotTagsMap, nil
}
