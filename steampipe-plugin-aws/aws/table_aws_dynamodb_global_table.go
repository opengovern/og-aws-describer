package aws

import (
	"context"

	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAwsDynamoDBGlobalTable(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_dynamodb_global_table",
		Description: "AWS DynamoDB Global Table",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("global_table_name"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"ResourceNotFoundException"}),
			},
			Hydrate: opengovernance.GetDynamoDbGlobalTable,
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListDynamoDbGlobalTable,
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:    "global_table_name",
					Require: plugin.Optional,
				},
			},
		},

		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "global_table_name",
				Description: "The global table name.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.GlobalTable.GlobalTableName")},
			{
				Name:        "global_table_arn",
				Description: "The unique identifier of the global table.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.GlobalTable.GlobalTableArn")},
			{
				Name:        "global_table_status",
				Description: "The current state of the global table.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.GlobalTable.GlobalTableStatus")},
			{
				Name:        "creation_date_time",
				Description: "The creation time of the global table.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.GlobalTable.CreationDateTime")},
			{
				Name:        "replication_group",
				Description: "The Regions where the global table has replicas.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.GlobalTable.ReplicationGroup")},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.GlobalTable.GlobalTableName")},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.GlobalTable.GlobalTableArn").Transform(arnToAkas),
			},
		}),
	}
}
