package aws

import (
	"context"
	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAwsOpenSearchServerlessCollection(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_opensearchserverless_collection",
		Description: "AWS OpenSearchServerless Collection",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("arn"), //TODO: change this to the primary key columns in model.go
			Hydrate:    kaytu.GetOpenSearchServerlessCollection,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListOpenSearchServerlessCollection,
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the collection.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Collection.Name")},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the collection",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.CollectionSummary.Arn")},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.CollectionSummary.Name")},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Collection.Tags"), // probably needs a transform function
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.CollectionSummary.Arn").Transform(arnToAkas),
			},
		}),
	}
}
