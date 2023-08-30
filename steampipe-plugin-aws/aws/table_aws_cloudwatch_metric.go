package aws

import (
	"context"

	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsCloudWatchMetric(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_cloudwatch_metric",
		Description: "AWS CloudWatch Metric",
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListCloudWatchMetric,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"InvalidParameterValue"}),
			},
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:    "metric_name",
					Require: plugin.Optional,
				},
				{
					Name:    "namespace",
					Require: plugin.Optional,
				},
				{
					Name:       "dimensions_filter",
					Require:    plugin.Optional,
					CacheMatch: "exact",
				},
			},
		},

		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "metric_name",
				Description: "The name of the metric.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Metric.MetricName")},
			{
				Name:        "namespace",
				Description: "The namespace for the metric.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Metric.Namespace")},
			{
				Name:        "dimensions_filter",
				Description: "The dimensions to filter against.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromQual("dimensions_filter"),
			},
			{
				Name:        "dimensions",
				Description: "The dimensions for the metric.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Metric.Dimensions")},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Metric.MetricName")},
		}),
	}
}
