package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/guardduty/types"

	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

type DestinationInfo = struct {
	DestinationId                   *string
	DestinationType                 types.DestinationType
	DestinationArn                  *string
	KmsKeyArn                       *string
	Status                          types.PublishingStatus
	PublishingFailureStartTimestamp *int64
	DetectorId                      string
}

func tableAwsGuardDutyPublishingDestination(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_guardduty_publishing_destination",
		Description: "AWS GuardDuty Publishing Destination",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"detector_id", "destination_id"}),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"InvalidInputException", "NoSuchEntityException", "BadRequestException"}),
			},
			Hydrate: kaytu.GetGuardDutyPublishingDestination,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListGuardDutyPublishingDestination,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "detector_id", Require: plugin.Optional},
			},
		},

		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "destination_id",
				Description: "The ID of the publishing destination.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.PublishingDestination.DestinationId")},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) for the publishing destination.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("ARN").Transform(arnToAkas),
			},
			{
				Name:        "detector_id",
				Description: "The ID of the detector.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DetectorId")},
			{
				Name:        "destination_arn",
				Description: "The ARN of the resource to publish to.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.PublishingDestination.DestinationProperties.DestinationArn")},
			{
				Name:        "destination_type",
				Description: "The type of publishing destination. Currently, only Amazon S3 buckets are supported.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.PublishingDestination.DestinationType")},
			{
				Name:        "kms_key_arn",
				Description: "The ARN of the KMS key to use for encryption.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.PublishingDestination.DestinationProperties.KmsKeyArn")},
			{
				Name:        "publishing_failure_start_timestamp",
				Description: "The time, in epoch millisecond format, at which GuardDuty was first unable to publish findings to the destination.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.PublishingDestination.PublishingFailureStartTimestamp").Transform(transform.UnixMsToTimestamp),
			},
			{
				Name:        "status",
				Description: "The status of the publishing destination.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.PublishingDestination.Status")},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.PublishingDestination.DestinationId")},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("ARN").Transform(transform.EnsureStringArray),
			},
		}),
	}
}
