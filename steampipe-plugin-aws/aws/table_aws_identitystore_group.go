package aws

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAwsIdentityStoreGroup(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_identitystore_group",
		Description: "AWS Identity Store Group",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"identity_store_id", "id"}),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"ResourceNotFoundException", "ValidationException"}),
			},
			//Hydrate: kaytu.GetIdentityStoreGroup,
		},
		List: &plugin.ListConfig{
			KeyColumns: plugin.AllColumns([]string{"identity_store_id"}),
			//Hydrate:    kaytu.ListIdentityStoreGroup,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"ResourceNotFoundException"}),
			},
		},

		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "identity_store_id",
				Description: "The globally unique identifier for the identity store.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Group.IdentityStoreId")},
			{
				Name:        "name",
				Description: "Contains the group's display name value.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Group.DisplayName")},
			{
				Name:        "id",
				Description: "The identifier for a group in the identity store.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Group.GroupId"),
			},

			// Standard columns for all tables
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Group.DisplayName")},

			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Group.GroupId").Transform(arnToAkas),
			},
		}),
	}
}
