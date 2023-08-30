package aws

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch/types"
	cloudwatchv1 "github.com/aws/aws-sdk-go/service/cloudwatch"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsCloudWatchMetricStatisticDataPoint(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_cloudwatch_metric_statistic_data_point",
		Description: "AWS CloudWatch Metric Statistic Data Point",
		List: &plugin.ListConfig{
			Hydrate: nil,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"InvalidParameterValue"}),
			},
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:    "metric_name",
					Require: plugin.Required,
				},
				{
					Name:    "namespace",
					Require: plugin.Required,
				},
				{
					Name:       "timestamp",
					Operators:  []string{">", ">=", "=", "<", "<="},
					Require:    plugin.Optional,
					CacheMatch: "exact",
				},
				{
					Name:       "period",
					Require:    plugin.Optional,
					CacheMatch: "exact",
				},
				{
					Name:       "dimensions",
					Require:    plugin.Optional,
					CacheMatch: "exact",
				},
				{
					Name:       "unit",
					Require:    plugin.Optional,
					CacheMatch: "exact",
				},
			},
		},
		GetMatrixItemFunc: SupportedRegionMatrix(cloudwatchv1.EndpointsID),
		Columns: awsRegionalColumns(cwMetricColumns([]*plugin.Column{
			{
				Name:        "label",
				Description: "A label for the specified metric.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "period",
				Description: "The granularity, in seconds, of the returned data points.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "dimensions",
				Description: "The dimensions for the metric.",
				Type:        proto.ColumnType_JSON,
			},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Label"),
			},
		})),
	}
}

type MetricStatistics struct {
	MetricName *string
	Namespace  *string
	Period     *int32
	Label      *string
	Dimensions []types.Dimension
	types.Datapoint
}
