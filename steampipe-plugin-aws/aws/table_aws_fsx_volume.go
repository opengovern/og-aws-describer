package aws

import (
	"context"

	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAwsFsxVolume(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_fsx_volume",
		Description: "AWS FSX Volume",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("volume_id"),
			Hydrate:    kaytu.GetFSXVolume,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListFSXVolume,
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "volume_id",
				Description: "The id of the volume.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Volume.VolumeId")},
			{
				Name:        "name",
				Description: "The name of the volume.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Volume.Name")},
			{
				Name:        "resource_arn",
				Description: "The Amazon Resource Name (ARN) of the volume",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Volume.ResourceARN")},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Volume.VolumeId")},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(getFsxVolumeTurbotTags),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Volume.ResourceARN").Transform(arnToAkas),
			},
		}),
	}
}

//// TRANSFORM FUNCTIONS

func getFsxVolumeTurbotTags(_ context.Context, d *transform.TransformData) (interface{}, error) {
	tags := d.HydrateItem.(kaytu.FSXVolume).Description.Volume.Tags
	return fsxV2TagsToMap(tags)
}
