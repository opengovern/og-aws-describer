package aws

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/cloudwatch/types"
	cloudwatchv1 "github.com/aws/aws-sdk-go/service/cloudwatch"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsCloudWatchMetricDataPoint(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_cloudwatch_metric_data_point",
		Description: "AWS CloudWatch Metric Data Point",
		List: &plugin.ListConfig{
			Hydrate: nil,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"InvalidParameterValue"}),
			},
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:    "id",
					Require: plugin.Required,
				},
				{
					Name:       "source_account_id",
					Require:    plugin.Optional,
					CacheMatch: "exact",
				},
				{
					Name:       "expression",
					Require:    plugin.Optional,
					CacheMatch: "exact",
				},
				{
					Name:       "metric_stat",
					Require:    plugin.Optional,
					CacheMatch: "exact",
				},
				{
					Name:       "period",
					Require:    plugin.Optional,
					CacheMatch: "exact",
				},
				{
					Name:    "scan_by",
					Require: plugin.Optional,
				},
				{
					Name:       "timestamp",
					Operators:  []string{">", ">=", "=", "<", "<="},
					Require:    plugin.Optional,
					CacheMatch: "exact",
				},
				{
					Name:       "timezone",
					Require:    plugin.Optional,
					CacheMatch: "exact",
				},
			},
		},
		GetMatrixItemFunc: SupportedRegionMatrix(cloudwatchv1.EndpointsID),
		Columns: awsRegionalColumns([]*plugin.Column{
			{
				Name:        "id",
				Description: "The short name you specified to represent this metric.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.MetricDataQuery.Id"),
			},
			{
				Name:        "label",
				Description: "The human-readable label associated with the data.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.MetricDataQuery.Label"),
			},
			{
				Name:        "period",
				Description: "The granularity, in seconds, of the returned data points.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.MetricDataQuery.Period"),
			},
			{
				Name:        "status_code",
				Description: "The status of the returned data. Complete indicates that all data points in the requested time range were returned. PartialData means that an incomplete set of data points were returned.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.MetricDataResult.StatusCode"),
			},
			{
				Name:        "timestamp",
				Description: "The timestamp for the data points, formatted in Unix timestamp format.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Timestamp"),
			},
			{
				Name:        "value",
				Description: "The data point for the metric corresponding to Timestamp.",
				Type:        proto.ColumnType_DOUBLE,
				Transform:   transform.FromField("Description.Value"),
			},
			{
				Name:        "expression",
				Description: "This field can contain either a Metrics Insights query, or a metric math expression to be performed on the returned data.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.MetricDataResult.Expression"),
			},
			{
				Name:        "scan_by",
				Description: "The order in which data points should be returned. TimestampDescending returns the newest data first and paginates when the MaxDatapoints limit is reached. TimestampAscending returns the oldest data first and paginates when the MaxDatapoints limit is reached.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.MetricDataResult.Expression"),
			},
			{
				Name:        "timezone",
				Description: "You can use timezone to specify your time zone so that the labels of returned data display the correct time for your time zone.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.MetricDataResult.Expression"),
			},
			{
				Name:        "source_account_id",
				Description: "The ID of the account where the metrics are located.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.MetricDataResult.AccountId"),
			},
			{
				Name:        "metric_stat",
				Description: "The metric to be returned, along with statistics, period, and units. Use this parameter only if this object is retrieving a metric and not performing a math expression on returned data.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.MetricDataResult.MetricStat"),
			},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.MetricDataResult.Label"),
			},
		}),
	}
}

type MetricDataPoint struct {
	Id         *string
	Label      *string
	StatusCode types.StatusCode
	Period     *int32
	Timestamp  time.Time
	Value      float64
}
