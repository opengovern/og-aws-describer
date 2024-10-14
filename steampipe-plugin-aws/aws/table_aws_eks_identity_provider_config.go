package aws

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsEksIdentityProviderConfig(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_eks_identity_provider_config",
		Description: "AWS EKS Identity Provider Config",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"name", "type", "cluster_name"}),
			//Hydrate:    opengovernance.GetEKSIdentityProviderConfig,
		},
		List: &plugin.ListConfig{
			//Hydrate: opengovernance.ListEKSIdentityProviderConfig,
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the identity provider configuration.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ConfigName")},
			{
				Name:        "type",
				Description: "The type of the identity provider configuration.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ConfigType")},
			{
				Name:        "client_id",
				Description: "This is also known as audience. The ID of the client application that makes authentication requests to the OIDC identity provider.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.IdentityProviderConfig.ClientId")},
			{
				Name:        "cluster_name",
				Description: "The name of the cluster.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.IdentityProviderConfig.ClusterName")},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the configuration.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.IdentityProviderConfig.IdentityProviderConfigArn"),
			},
			{
				Name:        "groups_claim",
				Description: "The JSON web token (JWT) claim that the provider uses to return your groups.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.IdentityProviderConfig.GroupsClaim")},
			{
				Name:        "groups_prefix",
				Description: "The prefix that is prepended to group claims to prevent clashes with existing names (such as system: groups).",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.IdentityProviderConfig.GroupsPrefix")},
			{
				Name:        "issuer_url",
				Description: "The URL of the OIDC identity provider that allows the API server to discover public signing keys for verifying tokens.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.IdentityProviderConfig.IssuerUrl")},
			{
				Name:        "username_claim",
				Description: "The JSON Web token (JWT) claim that is used as the username.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.IdentityProviderConfig.UsernameClaim")},
			{
				Name:        "status",
				Description: "The status of the OIDC identity provider.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.IdentityProviderConfig.Status")},
			{
				Name:        "username_prefix",
				Description: "The prefix that is prepended to username claims to prevent clashes with existing names.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.IdentityProviderConfig.UsernamePrefix")},
			{
				Name:        "required_claims",
				Description: "The key-value pairs that describe required claims in the identity token.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.IdentityProviderConfig.RequiredClaims")},
			{
				Name:        "tags_src",
				Description: "The metadata to apply to the provider configuration to assist with categorization and organization.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.IdentityProviderConfig.Tags"),
			},

			// Standard columns for all tables
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ConfigName")},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.IdentityProviderConfig.Tags")},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.IdentityProviderConfig.IdentityProviderConfigArn").Transform(arnToAkas),
			},
		}),
	}
}

//// LIST FUNCTION

//// HYDRATE FUNCTIONS
