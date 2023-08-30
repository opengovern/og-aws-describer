package aws

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

// check why the list and get did commit
func tableAwsIdentityStoreUser(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_identitystore_user",
		Description: "AWS Identity Store User",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"identity_store_id", "id"}),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"ResourceNotFoundException", "ValidationException"}),
			},
			//Hydrate: kaytu.GetIdentityStoreUser,
		},
		List: &plugin.ListConfig{
			KeyColumns: plugin.AllColumns([]string{"identity_store_id"}),
			//Hydrate:    kaytu.ListIdentityStoreUser,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"ResourceNotFoundException"}),
			},
		},

		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "identity_store_id",
				Description: "The globally unique identifier for the identity store.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.User.IdentityStoreId")},
			{
				Name:        "name",
				Description: "Contains the user’s display name value.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.User.UserName")},
			{
				Name:        "id",
				Description: "The identifier for a user in the identity store.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("User.UserId"),
			},

			// Standard columns for all tables
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.User.UserName")},
		}),
	}
}
