package aws

import (
	"context"

	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsKinesisVideoStream(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_kinesis_video_stream",
		Description: "AWS Kinesis Video Stream",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("stream_name"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"ResourceNotFoundException"}),
			},
			Hydrate: opengovernance.GetKinesisVideoStream,
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListKinesisVideoStream,
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "stream_name",
				Description: "The name of the stream.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Stream.StreamName")},
			{
				Name:        "stream_arn",
				Description: "The Amazon Resource Name (ARN) of the stream.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Stream.StreamARN")},
			{
				Name:        "status",
				Description: "The status of the stream.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Stream.Status")},
			{
				Name:        "version",
				Description: "The version of the stream.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Stream.Version")},
			{
				Name:        "kms_key_id",
				Description: "The ID of the AWS Key Management Service (AWS KMS) key that Kinesis Video Streams uses to encrypt data on the stream.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Stream.KmsKeyId")},
			{
				Name:        "creation_time",
				Description: "A time stamp that indicates when the stream was created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Stream.CreationTime")},
			{
				Name:        "data_retention_in_hours",
				Description: "How long the stream retains data, in hours.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Stream.DataRetentionInHours")},
			{
				Name:        "device_name",
				Description: "The name of the device that is associated with the stream.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Stream.DeviceName")},
			{
				Name:        "media_type",
				Description: "The MediaType of the stream.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Stream.MediaType")},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Tags")},

			// Standard columns for all tables
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Stream.StreamName")},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Stream.StreamARN").Transform(arnToAkas),
			},
		}),
	}
}
