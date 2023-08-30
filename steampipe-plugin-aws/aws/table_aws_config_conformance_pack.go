package aws

import (
	"context"

	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAwsConfigConformancePack(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_config_conformance_pack",
		Description: "AWS Config Conformance Pack",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("name"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"NoSuchConformancePackException"}),
			},
			Hydrate: kaytu.GetConfigConformancePack,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListConfigConformancePack,
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:    "name",
					Require: plugin.Optional,
				},
			},
		},

		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "Name of the conformance pack.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ConformancePack.ConformancePackName")},
			{
				Name:        "arn",
				Description: "Amazon Resource Name (ARN) of the conformance pack.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ConformancePack.ConformancePackArn"),
			},
			{
				Name:        "conformance_pack_id",
				Description: "ID of the conformance pack.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ConformancePack.ConformancePackId")},
			{
				Name:        "created_by",
				Description: "AWS service that created the conformance pack.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ConformancePack.CreatedBy")},
			{
				Name:        "delivery_s3_bucket",
				Description: "Amazon S3 bucket where AWS Config stores conformance pack templates.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ConformancePack.DeliveryS3Bucket")},
			{
				Name:        "delivery_s3_key_prefix",
				Description: "The prefix for the Amazon S3 delivery bucket.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ConformancePack.DeliveryS3KeyPrefix")},
			{
				Name:        "last_update_requested_time",
				Description: "Last update to the conformance pack.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.ConformancePack.LastUpdateRequestedTime")},
			{
				Name:        "input_parameters",
				Description: "A list of ConformancePackInputParameter objects.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ConformancePack.ConformancePackInputParameters")},

			// Standard columns
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ConformancePack.ConformancePackArn").Transform(arnToAkas),
			},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ConformancePack.ConformancePackName")},
		}),
	}
}
