package aws

import (
	"context"
	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAwsTimestreamDatabase(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_timestream_database",
		Description: "AWS Timestream Database",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("arn"), //TODO: change this to the primary key columns in model.go
			Hydrate:    opengovernance.GetTimestreamDatabase,
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListTimestreamDatabase,
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the database.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Database.DatabaseName")},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the database",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Database.Arn")},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Database.DatabaseName")},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Tags"), // probably needs a transform function
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Database.Arn").Transform(arnToAkas),
			},
		}),
	}
}
