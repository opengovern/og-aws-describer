package aws

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsEc2NetworkLoadBalancerMetricNetFlowCountDaily(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_ec2_network_load_balancer_metric_net_flow_count_daily",
		Description: "AWS EC2 Network Load Balancer Metrics - Net Flow Count (Daily)",
		List:        &plugin.ListConfig{
			//Hydrate: opengovernance.ListNetworkLoadBalancerMetricNetFlowCountDaily,
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
