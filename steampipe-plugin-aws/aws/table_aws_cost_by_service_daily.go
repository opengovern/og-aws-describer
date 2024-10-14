package aws

import (
	"context"

	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAwsCostByServiceDaily(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_cost_by_service_daily",
		Description: "AWS Cost Explorer - Cost by Service (Daily)",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListCostExplorerByServiceDaily,
		},
		Columns: awsKaytuColumns(
			kaytuCostExplorerColumns([]*plugin.Column{
				{
					Name:        "service",
					Description: "The name of the AWS service.",
					Type:        proto.ColumnType_STRING,
					Transform:   transform.FromField("Description.Dimension1"),
				},
				{
					Name:        "cost_source",
					Description: "The source of the cost.",
					Type:        proto.ColumnType_STRING,
					Transform:   transform.FromField("Description.Dimension2"),
				},
			}),
		),
	}
}
