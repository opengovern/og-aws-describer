package aws

import (
	"context"
	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsEc2InstanceMetricCpuUtilizationHourly(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_ec2_instance_metric_cpu_utilization_hourly",
		Description: "AWS EC2 Instance Cloudwatch Metrics - CPU Utilization (Hourly)",
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListEC2InstanceMetricCpuUtilizationHourly,
		},

		Columns: awsKaytuRegionalColumns(cwMetricColumns(
			[]*plugin.Column{
				{
					Name:        "instance_id",
					Description: "The ID of the instance.",
					Type:        proto.ColumnType_STRING,
					Transform:   transform.FromField("Description.InstanceId"),
				},
				{
					Name:        "timestamp",
					Description: "Datapoint Timestamp.",
					Type:        proto.ColumnType_TIMESTAMP,
					Transform:   transform.FromField("Description.Timestamp"),
				},
				{
					Name:        "average",
					Description: "Sample Average.",
					Type:        proto.ColumnType_DOUBLE,
					Transform:   transform.FromField("Description.Average"),
				},
				{
					Name:        "sum",
					Description: "Sample Sum.",
					Type:        proto.ColumnType_DOUBLE,
					Transform:   transform.FromField("Description.Sum"),
				},
				{
					Name:        "maximum",
					Description: "Sample Maximum.",
					Type:        proto.ColumnType_DOUBLE,
					Transform:   transform.FromField("Description.Maximum"),
				},
				{
					Name:        "minimum",
					Description: "Sample Minimum.",
					Type:        proto.ColumnType_DOUBLE,
					Transform:   transform.FromField("Description.Minimum"),
				},
				{
					Name:        "sample_count",
					Description: "Sample Count.",
					Type:        proto.ColumnType_DOUBLE,
					Transform:   transform.FromField("Description.SampleCount"),
				},
				{
					Name:        "account_id",
					Description: "Account Id of the resource",
					Type:        proto.ColumnType_STRING,
					Transform:   transform.FromField("Account"),
				},
			})),
	}
}
