package aws

import (
	"context"

	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsEksAddonVersion(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_eks_addon_version",
		Description: "AWS EKS Addon Version",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListEKSAddonVersion,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "addon_name", Require: plugin.Optional},
			},
		},

		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "addon_name",
				Description: "The name of the add-on.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AddonName")},
			{
				Name:        "addon_version",
				Description: "The version of the add-on.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AddonVersion.AddonVersion")},
			{
				Name:        "type",
				Description: "The type of the add-on.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AddonType")},
			{
				Name:        "addon_configuration",
				Description: "The configuration for the add-on.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.AddonConfiguration")},
			{
				Name:        "architecture",
				Description: "The architectures that the version supports.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.AddonVersion.Architecture")},
			{
				Name:        "compatibilities",
				Description: "An object that represents the compatibilities of a version.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.AddonVersion.Compatibilities")},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AddonVersion.AddonVersion")},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("ARN").Transform(arnToAkas)},
		}),
	}
}
