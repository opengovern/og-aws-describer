package aws

import (
	"context"

	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"

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
			Hydrate: opengovernance.GetGuardDutyDetector,
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListGuardDutyDetector,
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
				Transform:   transform.FromField("ARN"),
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
				Name:        "features",
				Description: "Describes the features that have been enabled for the detector.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Detector.Features"),
			},
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
				Transform:   transform.FromField("ARN").Transform(transform.EnsureStringArray),
			},
		}),
	}
}
