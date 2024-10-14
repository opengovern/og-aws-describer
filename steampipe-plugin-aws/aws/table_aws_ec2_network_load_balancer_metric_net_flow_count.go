package aws

import (
	"context"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsEc2NetworkLoadBalancerMetricNetFlowCount(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_ec2_network_load_balancer_metric_net_flow_count",
		Description: "AWS EC2 Network Load Balancer Metrics - Net Flow Count",
		List:        &plugin.ListConfig{
			//Hydrate: opengovernance.ListNetworkLoadBalancerMetricNetFlowCount,
		},

		Columns: awsKaytuRegionalColumns(kaytuCwMetricColumns(
			[]*plugin.Column{
				{
					Name:        "name",
					Description: "The friendly name of the Load Balancer.",
					Type:        proto.ColumnType_STRING,
					Transform:   transform.FromField("Description.DimensionValue"),
				},
			})),
	}
}
