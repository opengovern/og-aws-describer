package aws

import (
	"context"

	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"

	"github.com/aws/aws-sdk-go-v2/service/guardduty"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

type detectorInfo = struct {
	guardduty.GetDetectorOutput
	DetectorID string
}

//// TABLE DEFINITION

func tableAwsGuardDutyDetector(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_guardduty_detector",
		Description: "AWS GuardDuty Detector",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("detector_id"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"InvalidInputException", "BadRequestException"}),
			},
			Hydrate: kaytu.GetGuardDutyDetector,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListGuardDutyDetector,
		},

		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "detector_id",
				Description: "The ID of the detector.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.DetectorId")},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) specifying the detector.",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getGuardDutyDetectorARN,
				Transform:   transform.FromValue(),
			},
			{
				Name:        "status",
				Description: "The detector status.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Detector.Status")},
			{
				Name:        "created_at",
				Description: "The timestamp of when the detector was created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Detector.CreatedAt")},
			{
				Name:        "finding_publishing_frequency",
				Description: "The publishing frequency of the finding.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Detector.FindingPublishingFrequency")},
			{
				Name:        "service_role",
				Description: "The GuardDuty service role.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Detector.ServiceRole")},
			{
				Name:        "updated_at",
				Description: "The last-updated timestamp for the detector.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Detector.UpdatedAt")},
			{
				Name:        "data_sources",
				Description: "Describes which data sources are enabled for the detector.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Detector.DataSources")},

			{
				Name:        "master_account",
				Description: "Contains information about the administrator account and invitation.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromValue(),
				//	we don't have this field in struct
			},

			// Standard columns

			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DetectorId")},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Detector.Tags")},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Hydrate:     getGuardDutyDetectorARN,
				Transform:   transform.FromValue().Transform(transform.EnsureStringArray),
			},
		}),
	}
}

//// LIST FUNCTION

//// HYDRATE FUNCTIONS

func getGuardDutyDetectorARN(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("getGuardDutyDetectorARN")
	region := d.EqualsQualString(matrixKeyRegion)
	data := h.Item.(kaytu.GuardDutyDetector).Description
	metadata := h.Item.(kaytu.GuardDutyDetector).Metadata

	arn := "arn:" + metadata.Partition + ":guardduty:" + region + ":" + metadata.AccountID + ":detector/" + data.DetectorId

	return arn, nil
}
