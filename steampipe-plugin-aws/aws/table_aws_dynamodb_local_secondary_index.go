package aws

import (
	"context"

	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAwsDynamoDbLocalSecondaryIndex(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_dynamodb_local_secondary_index",
		Description: "AWS DynamoDb LocalSecondaryIndex",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("index_arn"),
			Hydrate:    kaytu.GetDynamoDbLocalSecondaryIndex,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListDynamoDbLocalSecondaryIndex,
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "index_name",
				Description: "The name of the local secondary index.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.LocalSecondaryIndex.IndexName")},
			{
				Name:        "index_arn",
				Description: "The Amazon Resource Name (ARN) of the local secondary index",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.LocalSecondaryIndex.IndexArn")},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.LocalSecondaryIndex.IndexName")},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.LocalSecondaryIndex.IndexArn").Transform(arnToAkas),
			},
		}),
	}
}
