package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/kinesis/types"
	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsKinesisStream(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_kinesis_stream",
		Description: "AWS Kinesis Stream",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("stream_name"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"ResourceNotFoundException", "InvalidParameter"}),
			},
			Hydrate: opengovernance.GetKinesisStream,
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListKinesisStream,
		},

		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "stream_name",
				Description: "The name of the stream being described.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Stream.StreamName")},
			{
				Name:        "stream_arn",
				Description: "The Amazon Resource Name (ARN) for the stream being described.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Stream.StreamARN")},
			{
				Name:        "stream_status",
				Description: "The current status of the stream being described.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Stream.StreamStatus")},
			{
				Name:        "stream_creation_timestamp",
				Description: "The approximate time that the stream was created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Stream.StreamCreationTimestamp")},
			{
				Name:        "encryption_type",
				Description: "The server-side encryption type used on the stream.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Stream.EncryptionType")},
			{
				Name:        "key_id",
				Description: "The GUID for the customer-managed AWS KMS key to use for encryption.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Stream.KeyId")},
			{
				Name:        "retention_period_hours",
				Description: "The current retention period, in hours.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Stream.RetentionPeriodHours")},
			{
				Name:        "consumer_count",
				Description: "The number of enhanced fan-out consumers registered with the stream.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.DescriptionSummary.ConsumerCount")},
			{
				Name:        "open_shard_count",
				Description: "The number of open shards in the stream.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.DescriptionSummary.OpenShardCount")},
			{
				Name:        "has_more_shards",
				Description: "If set to true, more shards in the stream are available to describe.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Stream.HasMoreShards")},
			{
				Name:        "shards",
				Description: "The shards that comprise the stream.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Stream.Shards")},
			{
				Name:        "enhanced_monitoring",
				Description: "Represents the current enhanced monitoring settings of the stream.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Stream.EnhancedMonitoring")},
			{
				Name:        "tags_src",
				Description: "A list of tags associated with the stream.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Tags")},
			// Standard columns for all tables
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Stream.StreamName")},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Tags").Transform(kinesisTagListToTurbotTags),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Stream.StreamARN").Transform(arnToAkas),
			},
		}),
	}
}

//// TRANSFORM FUNCTIONS

func kinesisTagListToTurbotTags(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	plugin.Logger(ctx).Trace("kinesisTagListToTurbotTags")
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
