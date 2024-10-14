package aws

import (
	"context"

	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAwsKeyspacesTable(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_keyspaces_table",
		Description: "AWS Keyspaces Table",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("table_name"),
			Hydrate:    opengovernance.GetKeyspacesTable,
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListKeyspacesTable,
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "table_name",
				Description: "The name of the table.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Table.TableName")},
			{
				Name:        "resource_arn",
				Description: "The Amazon Resource Name (ARN) of the table",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Table.ResourceArn")},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Table.TableName")},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(getKeyspacesTableTurbotTags),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Table.ResourceArn").Transform(arnToAkas),
			},
		}),
	}
}

//// TRANSFORM FUNCTIONS

func getKeyspacesTableTurbotTags(_ context.Context, d *transform.TransformData) (interface{}, error) {
	tags := d.HydrateItem.(opengovernance.KeyspacesTable).Description.Tags
	return keyspacesV2TagsToMap(tags)
}
