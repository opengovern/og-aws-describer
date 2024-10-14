package aws

import (
	"context"
	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

// check why the list and get did commit
func tableAwsIdentityStoreUser(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_identitystore_user",
		Description: "AWS Identity Store User",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListIdentityStoreUser,
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
				Transform:   transform.FromField("Description.User.DisplayName")},
			{
				Name:        "user_name",
				Description: "Contains the user’s display name value.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.User.UserName")},
			{
				Name:        "email",
				Description: "Contains the user’s display name value.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.PrimaryEmail")},
			{
				Name:        "id",
				Description: "The identifier for a user in the identity store.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.User.UserId"),
			},
			{
				Name:        "external_ids",
				Description: "The identifier for a group in the identity store.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.User.ExternalIds"),
			},

			// Standard columns for all tables
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.User.Title")},
		}),
	}
}
