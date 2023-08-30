package aws

import (
	"context"

	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAwsCostForecastDaily(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_cost_forecast_daily",
		Description: "AWS Cost Explorer - Cost Forecast (Daily)",
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListCostExplorerForcastDaily,
		},
		Columns: awsKaytuColumns([]*plugin.Column{
			{
				Name:        "period_start",
				Description: "Start timestamp for this cost metric",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.PeriodStart"),
			},
			{
				Name:        "period_end",
				Description: "End timestamp for this cost metric",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.PeriodEnd"),
			},
			{
				Name:        "mean_value",
				Description: "Average forecasted value",
				Type:        proto.ColumnType_DOUBLE,
				Transform:   transform.FromField("Description.MeanValue"),
			},
		},
		),
	}
}
