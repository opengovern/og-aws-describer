package aws

import (
	"context"
	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"

	ssmv1 "github.com/aws/aws-sdk-go/service/ssm"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAwsSSMInventoryEntry(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_ssm_inventory_entry",
		Description: "AWS SSM Inventory Entry",
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListSSMInventoryEntry,
			KeyColumns: plugin.KeyColumnSlice{
				{Name: "instance_id", Require: plugin.Optional},
				{Name: "type_name", Require: plugin.Optional},
			},
		},
		GetMatrixItemFunc: SupportedRegionMatrix(ssmv1.EndpointsID),
		Columns: awsRegionalColumns([]*plugin.Column{
			{
				Name:        "instance_id",
				Description: "The managed node ID targeted by the request to query inventory information.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.InstanceId")},
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
				Name:        "entries",
				Description: "The inventory items on the managed node(s).",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Entries")},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.InstanceId")},
		}),
	}
}
