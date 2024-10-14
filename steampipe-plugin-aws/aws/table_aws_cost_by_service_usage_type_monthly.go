package aws

import (
	"context"

	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAwsCostByServiceUsageTypeMonthly(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_cost_by_service_usage_type_monthly",
		Description: "AWS Cost Explorer - Cost by Service and Usage Type (Monthly)",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListCostExplorerByServiceUsageTypeMonthly,
			KeyColumns: plugin.KeyColumnSlice{
				{Name: "service", Operators: []string{"=", "<>"}, Require: plugin.Optional},
				{Name: "usage_type", Operators: []string{"=", "<>"}, Require: plugin.Optional},
			},
		},
		Columns: awsKaytuColumns(
			kaytuCostExplorerColumns([]*plugin.Column{
				{
					Name:        "service",
					Description: "The name of the AWS service.",
					Type:        proto.ColumnType_STRING,
					Transform:   transform.FromField("Dimension1"),
				},
				{
					Name:        "usage_type",
					Description: "The usage type of this metric.",
					Type:        proto.ColumnType_STRING,
					Transform:   transform.FromField("Dimension2"),
				},
			}),
		),
	}
}
