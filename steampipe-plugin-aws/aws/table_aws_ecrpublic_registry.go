package aws

import (
	"context"

	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAwsEcrpublicRegistry(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_ecrpublic_registry",
		Description: "AWS ECRPublic Registry}",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("registry_id"),
			Hydrate:    kaytu.GetECRPublicRegistry,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListECRPublicRegistry,
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "registry_id",
				Description: "The id of the public registry.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.PublicRegistry.RegistryId")},
			{
				Name:        "registry_arn",
				Description: "The Amazon Resource Name (ARN) of the public registry",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.PublicRegistry.RegistryArn")},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.PublicRegistry.RegistryId")},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(getEcrpublicFleetTurbotTags),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.PublicRegistry.RegistryArn").Transform(arnToAkas),
			},
		}),
	}
}

//// TRANSFORM FUNCTIONS

func getEcrpublicFleetTurbotTags(_ context.Context, d *transform.TransformData) (interface{}, error) {
	tags := d.HydrateItem.(kaytu.ECRPublicRegistry).Description.Tags
	return ecrpublicV2TagsToMap(tags)
}
