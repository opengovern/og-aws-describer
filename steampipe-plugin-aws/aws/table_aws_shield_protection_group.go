package aws

import (
	"context"

	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAwsShieldProtectionGroup(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_shield_protection_group",
		Description: "AWS Shield ProtectionGroup",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("protection_group_id"),
			Hydrate:    opengovernance.GetShieldProtectionGroup,
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListShieldProtectionGroup,
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "protection_group_id",
				Description: "The id of the protection group.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ProtectionGroup.ProtectionGroupId")},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the protection group",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ProtectionGroup.ProtectionGroupArn")},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ProtectionGroup.ProtectionGroupId")},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(getShieldProtectionGroupTurbotTags),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ProtectionGroup.ProtectionGroupArn").Transform(arnToAkas),
			},
		}),
	}
}

//// TRANSFORM FUNCTIONS

func getShieldProtectionGroupTurbotTags(_ context.Context, d *transform.TransformData) (interface{}, error) {
	tags := d.HydrateItem.(opengovernance.ShieldProtectionGroup).Description.Tags
	return shieldV2TagsToMap(tags)
}
