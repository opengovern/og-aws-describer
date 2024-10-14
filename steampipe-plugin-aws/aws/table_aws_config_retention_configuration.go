package aws

import (
	"context"
	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAwsConfigRetentionConfiguration(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_config_retention_configuration",
		Description: "AWS Config Retention Configuration",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListConfigRetentionConfiguration,
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the retention configuration object.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.RetentionConfiguration.Name"),
			},
			{
				Name:        "retention_period_in_days",
				Description: "Number of days Config stores your historical information.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.RetentionConfiguration.RetentionPeriodInDays"),
			},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.RetentionConfiguration.Name"),
			},
		}),
	}
}
