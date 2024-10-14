package aws

import (
	"context"
	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAwsCloudwatchLogMetricFilter(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_cloudwatch_log_metric_filter",
		Description: "AWS CloudWatch Log Metric Filter",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("name"),
			Hydrate:    opengovernance.GetCloudWatchLogsMetricFilter,
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListCloudWatchLogsMetricFilter,
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:    "name",
					Require: plugin.Optional,
				},
				{
					Name:    "log_group_name",
					Require: plugin.Optional,
				},
				{
					Name:    "metric_transformation_name",
					Require: plugin.Optional,
				},
				{
					Name:    "metric_transformation_namespace",
					Require: plugin.Optional,
				},
			},
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the metric filter",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.MetricFilter.FilterName")},
			{
				Name:        "log_group_name",
				Description: "The name of the log group",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.MetricFilter.LogGroupName")},
			{
				Name:        "creation_time",
				Description: "The creation time of the metric filter",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.MetricFilter.CreationTime").Transform(convertTimestamp),
			},
			{
				Name:        "filter_pattern",
				Description: "A symbolic description of how CloudWatch Logs should interpret the data in each log event",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.MetricFilter.FilterPattern")},
			{
				Name:        "metric_transformation_name",
				Description: "The name of the CloudWatch metric",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromP(logMetricTransformationsData, "MetricName"),
			},
			{
				Name:        "metric_transformation_namespace",
				Description: "A custom namespace to contain metric in CloudWatch. Namespaces are used to group together metrics that are similar",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromP(logMetricTransformationsData, "MetricNamespace"),
			},
			{
				Name:        "metric_transformation_value",
				Description: "The value to publish to the CloudWatch metric when a filter pattern matches a log event",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromP(logMetricTransformationsData, "MetricValue"),
			},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.MetricFilter.FilterName")},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(getCloudwatchLogMetricFilterAkas),
			},
		}),
	}
}

//// LIST FUNCTION

//// HYDRATE FUNCTIONS

//// TRANSFORM FUNCTIONS

func getCloudwatchLogMetricFilterAkas(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	metricFilter := d.HydrateItem.(opengovernance.CloudWatchLogsMetricFilter).Description.MetricFilter
	metadata := d.HydrateItem.(opengovernance.CloudWatchLogsMetricFilter).Metadata

	// Get data for turbot defined properties
	akas := []string{"arn:" + metadata.Partition + ":logs:" + metadata.Region + ":" + metadata.AccountID + ":log-group:" + *metricFilter.LogGroupName + ":metric-filter:" + *metricFilter.FilterName}

	return akas, nil
}

func logMetricTransformationsData(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	plugin.Logger(ctx).Trace("logMetricTransformationsData")
	metricFilterData := d.HydrateItem.(opengovernance.CloudWatchLogsMetricFilter).Description.MetricFilter

	if metricFilterData.MetricTransformations != nil && len(metricFilterData.MetricTransformations) > 0 {
		if d.Param.(string) == "MetricName" {
			return metricFilterData.MetricTransformations[0].MetricName, nil
		} else if d.Param.(string) == "MetricNamespace" {
			return metricFilterData.MetricTransformations[0].MetricNamespace, nil
		} else {
			return metricFilterData.MetricTransformations[0].MetricValue, nil
		}
	}
	return nil, nil
}
