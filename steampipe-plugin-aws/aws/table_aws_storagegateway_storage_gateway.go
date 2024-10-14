package aws

import (
	"context"

	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAwsStorageGatewayStorageGateway(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_storagegateway_storage_gateway",
		Description: "AWS StorageGateway StorageGateway",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("gateway_id"),
			Hydrate:    opengovernance.GetStorageGatewayStorageGateway,
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListStorageGatewayStorageGateway,
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "gateway_id",
				Description: "The id of the storage gateway.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.StorageGateway.GatewayId")},
			{
				Name:        "name",
				Description: "The name of the storage gateway.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.StorageGateway.GatewayName")},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the storage gateway",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.StorageGateway.GatewayARN")},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.StorageGateway.GatewayName")},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(getStorageGatewayStorageGatewayTurbotTags),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.StorageGateway.GatewayARN").Transform(arnToAkas),
			},
		}),
	}
}

//// TRANSFORM FUNCTIONS

func getStorageGatewayStorageGatewayTurbotTags(_ context.Context, d *transform.TransformData) (interface{}, error) {
	tags := d.HydrateItem.(opengovernance.StorageGatewayStorageGateway).Description.Tags
	return storageGatewayV2TagsToMap(tags)
}
