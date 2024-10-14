package aws

import (
	"context"

	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAwsRAMResourceAssociation(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_ram_resource_association",
		Description: "AWS RAM Resource Association",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListRamResourceAssociation,
		},

		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "resource_share_name",
				Description: "The name of the resource share.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ResourceAssociation.ResourceShareName")},
			{
				Name:        "resource_share_arn",
				Description: "The Amazon Resoure Name (ARN) of the resource share.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ResourceAssociation.ResourceShareArn")},
			{
				Name:        "status",
				Description: "The current status of the association.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ResourceAssociation.Status")},
			{
				Name:        "associated_entity",
				Description: "The Amazon Resoure Name (ARN) of the associated resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ResourceAssociation.AssociatedEntity")},
			{
				Name:        "association_type",
				Description: "The type of entity included in this association.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ResourceAssociation.AssociationType")},
			{
				Name:        "creation_time",
				Description: "The date and time when the association was created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.ResourceAssociation.CreationTime")},
			{
				Name:        "external",
				Description: "Indicates whether the principal belongs to the same organization in Organizations as the Amazon Web Services account that owns the resource share.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.ResourceAssociation.External")},
			{
				Name:        "last_updated_time",
				Description: "The date and time when the association was last updated..",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.ResourceAssociation.LastUpdatedTime")},
			{
				Name:        "status_message",
				Description: "A message about the status of the association.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ResourceAssociation.StatusMessage")},
			{
				Name:        "resource_share_permission",
				Description: "Information about an RAM permission that is associated with a resource share and any of its resources of a specified type.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ResourceSharePermission")},

			// Standard columns for all tables
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ResourceAssociation.ResourceShareName")},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ResourceAssociation.ResourceShareArn").Transform(transform.EnsureStringArray),
			},
		}),
	}
}
