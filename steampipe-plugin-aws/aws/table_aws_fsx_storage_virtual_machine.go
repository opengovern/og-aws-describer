package aws

import (
	"context"

	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAwsFsxStorageVirtualMachine(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_fsx_storage_virtual_machine",
		Description: "AWS FSX StorageVirtualMachine",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("storage_virtual_machine_id"),
			Hydrate:    opengovernance.GetFSXStorageVirtualMachine,
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListFSXStorageVirtualMachine,
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "storage_virtual_machine_id",
				Description: "The id of the storage virtual machine.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.StorageVirtualMachine.StorageVirtualMachineId")},
			{
				Name:        "name",
				Description: "The name of the storage virtual machine.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.StorageVirtualMachine.Name")},
			{
				Name:        "resource_arn",
				Description: "The Amazon Resource Name (ARN) of the storage virtual machine",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.StorageVirtualMachine.ResourceARN")},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.StorageVirtualMachine.StorageVirtualMachineId")},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(getFsxStorageVirtualMachineTurbotTags),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.StorageVirtualMachine.ResourceARN").Transform(arnToAkas),
			},
		}),
	}
}

//// TRANSFORM FUNCTIONS

func getFsxStorageVirtualMachineTurbotTags(_ context.Context, d *transform.TransformData) (interface{}, error) {
	tags := d.HydrateItem.(opengovernance.FSXStorageVirtualMachine).Description.StorageVirtualMachine.Tags
	return fsxV2TagsToMap(tags)
}
