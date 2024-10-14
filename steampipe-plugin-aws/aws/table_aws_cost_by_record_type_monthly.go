package aws

import (
	"context"

	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAwsCostByRecordTypeMonthly(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_cost_by_record_type_monthly",
		Description: "AWS Cost Explorer - Cost by Record Type (Monthly)",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListCostExplorerByRecordTypeMonthly,
		},
		Columns: awsKaytuColumns(
			kaytuCostExplorerColumns([]*plugin.Column{
				{
					Name:        "linked_account_id",
					Description: "The linked AWS Account ID.",
					Type:        proto.ColumnType_STRING,
					Transform:   transform.FromField("Description.Dimension1"),
				},
				{
					Name:        "record_type",
					Description: "The different types of charges such as RI fees, usage, costs, tax refunds, and credits.",
					Type:        proto.ColumnType_STRING,
					Transform:   transform.FromField("Description.Dimension2"),
				},
			}),
		),
	}
}
