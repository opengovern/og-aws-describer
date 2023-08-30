package aws

import (
	"context"
	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAwsSSMInventory(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_ssm_inventory",
		Description: "AWS SSM Inventory",
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListSSMInventory,
			KeyColumns: plugin.KeyColumnSlice{
				{Name: "id", Require: plugin.Optional, Operators: []string{"=", "<>"}},
				{Name: "type_name", Require: plugin.Optional},
			},
		},

		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "id",
				Description: "ID of the inventory result entity.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Id")},
			{
				Name:        "type_name",
				Description: "The type of inventory item returned by the request.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.TypeName")},
			{
				Name:        "capture_time",
				Description: "The time that inventory information was collected for the managed node(s).",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.CaptureTime")},
			{
				Name:        "schema_version",
				Description: "The inventory schema version used by the managed node(s).",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.SchemaVersion")},
			{
				Name:        "content",
				Description: "Contains all the inventory data of the item type. Results include attribute names and values.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Content")},
			{
				Name:        "schema",
				Description: "The inventory item schema definition. Users can use this to compose inventory query filters.",
				Type:        proto.ColumnType_JSON,

				Transform: transform.FromField("Description.Schemas")},

			// Steampipe standard columns

			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Id")},
		}),
	}
}
