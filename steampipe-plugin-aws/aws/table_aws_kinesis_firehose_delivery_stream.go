package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/firehose/types"
	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsKinesisFirehoseDeliveryStream(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_kinesis_firehose_delivery_stream",
		Description: "AWS Kinesis Firehose Delivery Stream",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("delivery_stream_name"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"ValidationException", "InvalidParameter", "ResourceNotFoundException"}),
			},
			Hydrate: opengovernance.GetFirehoseDeliveryStream,
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListFirehoseDeliveryStream,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "delivery_stream_type", Require: plugin.Optional},
			},
		},

		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "delivery_stream_name",
				Description: "The name of the delivery stream.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DeliveryStream.DeliveryStreamName")},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the delivery stream.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DeliveryStream.DeliveryStreamARN"),
			},
			{
				Name:        "delivery_stream_status",
				Description: "The server-side encryption type used on the stream.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DeliveryStream.DeliveryStreamStatus")},
			{
				Name:        "delivery_stream_type",
				Description: "The delivery stream type.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DeliveryStream.DeliveryStreamType")},
			{
				Name:        "version_id",
				Description: "The version id of the stream. Each time the destination is updated for a delivery stream, the version ID is changed, and the current version ID is required when updating the destination",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DeliveryStream.VersionId")},
			{
				Name:        "create_timestamp",
				Description: "The date and time that the delivery stream was created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.DeliveryStream.CreateTimestamp")},
			{
				Name:        "has_more_destinations",
				Description: "Indicates whether there are more destinations available to list.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.DeliveryStream.HasMoreDestinations")},
			{
				Name:        "last_update_timestamp",
				Description: "The date and time that the delivery stream was last updated.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.DeliveryStream.LastUpdateTimestamp")},
			{
				Name:        "delivery_stream_encryption_configuration",
				Description: "Indicates the server-side encryption (SSE) status for the delivery stream.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DeliveryStream.DeliveryStreamEncryptionConfiguration")},
			{
				Name:        "destinations",
				Description: "The destinations for the stream.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DeliveryStream.Destinations")},
			{
				Name:        "failure_description",
				Description: "Provides details in case one of the following operations fails due to an error related to KMS: CreateDeliveryStream, DeleteDeliveryStream, StartDeliveryStreamEncryption,StopDeliveryStreamEncryption.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DeliveryStream.FailureDescription")},
			{
				Name:        "source",
				Description: "If the DeliveryStreamType parameter is KinesisStreamAsSource, a SourceDescription object describing the source Kinesis data stream.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DeliveryStream.Source")},
			{
				Name:        "tags_src",
				Description: "A list of tags associated with the delivery stream.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Tags")},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DeliveryStream.DeliveryStreamName")},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Tags").Transform(kinesisFirehoseTagListToTurbotTags),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DeliveryStream.DeliveryStreamARN").Transform(transform.EnsureStringArray),
			},
		}),
	}
}

//// TRANSFORM FUNCTIONS

func kinesisFirehoseTagListToTurbotTags(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	tagList := d.Value.([]types.Tag)

	if tagList == nil {
		return nil, nil
	}

	// Mapping the resource tags inside turbotTags
	var turbotTagsMap map[string]string
	if tagList != nil {
		turbotTagsMap = map[string]string{}
		for _, i := range tagList {
			turbotTagsMap[*i.Key] = *i.Value
		}
	}

	return turbotTagsMap, nil
}
