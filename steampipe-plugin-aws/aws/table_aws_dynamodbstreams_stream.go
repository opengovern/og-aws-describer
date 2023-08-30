package aws

import (
	"context"

	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAwsDynamodbstreamsStream(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_dynamodbstreams_stream",
		Description: "AWS DynamoDb Stream",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("stream_arn"),
			Hydrate:    kaytu.GetDynamoDbStream,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListDynamoDbStream,
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "stream_label",
				Description: "The label of the stream.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Stream.StreamLabel")},
			{
				Name:        "stream_arn",
				Description: "The Amazon Resource Name (ARN) of the stream",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Stream.StreamArn")},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Stream.StreamLabel")},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Stream.StreamArn").Transform(arnToAkas),
			},
		}),
	}
}
