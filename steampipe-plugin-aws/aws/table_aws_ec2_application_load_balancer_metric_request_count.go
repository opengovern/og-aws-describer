package aws

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsEc2ApplicationLoadBalancerMetricRequestCount(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_ec2_application_load_balancer_metric_request_count",
		Description: "AWS EC2 Application Load Balancer Metrics - Request Count",
		List:        &plugin.ListConfig{
			//Hydrate: opengovernance.ListApplicationLoadBalancerMetricRequestCount,
		},

		Columns: awsKaytuRegionalColumns(kaytuCwMetricColumns(
			[]*plugin.Column{
				{
					Name:        "name",
					Description: "The friendly name of the Load Balancer that was provided during resource creation.",
					Type:        proto.ColumnType_STRING,
					Transform:   transform.FromField("Description.DimensionValue"),
				},
			})),
	}
}
