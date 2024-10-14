package aws

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsEbsVolumeMetricReadOps(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_ebs_volume_metric_read_ops",
		Description: "AWS EBS Volume Cloudwatch Metrics - Read Ops",
		List:        &plugin.ListConfig{
			//Hydrate: opengovernance.ListEbsVolumeMetricReadOps,
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
