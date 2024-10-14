package aws

import (
	"context"
	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsEc2InstanceMetricCpuUtilizationDaily(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_ec2_instance_metric_cpu_utilization_daily",
		Description: "AWS EC2 Instance Cloudwatch Metrics - CPU Utilization (Daily)",
		List: &plugin.ListConfig{
			ParentHydrate: opengovernance.ListEC2Instance,
			Hydrate:       listEc2InstanceMetricCpuUtilizationDaily,
		},

		Columns: awsKaytuRegionalColumns(cwMetricColumns(
			[]*plugin.Column{
				{
					Name:        "instance_id",
					Description: "The ID of the instance.",
					Type:        proto.ColumnType_STRING,
					Transform:   transform.FromField("DimensionValue"),
				},
			})),
	}
}

func listEc2InstanceMetricCpuUtilizationDaily(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	instance := h.Item.(opengovernance.EC2Instance).Description.Instance
	return listCWMetricStatistics(ctx, d, "DAILY", "AWS/EC2", "CPUUtilization", "InstanceId", *instance.InstanceId)
}
