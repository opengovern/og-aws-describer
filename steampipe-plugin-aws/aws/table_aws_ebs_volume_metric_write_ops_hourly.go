package aws

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsEbsVolumeMetricWriteOpsHourly(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_ebs_volume_metric_write_ops_hourly",
		Description: "AWS EBS Volume Cloudwatch Metrics - Write Ops (Hourly)",
		List:        &plugin.ListConfig{
			//Hydrate: opengovernance.ListEbsVolumeMetricReadOpsHourly,
		},

		Columns: awsKaytuRegionalColumns(kaytuCwMetricColumns(
			[]*plugin.Column{
				{
					Name:        "volume_id",
					Description: "The EBS Volume ID.",
					Type:        proto.ColumnType_STRING,
					Transform:   transform.FromField("Description.DimensionValue"),
				},
			})),
	}
}
