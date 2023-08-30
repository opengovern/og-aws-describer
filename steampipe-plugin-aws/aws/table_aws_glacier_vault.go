package aws

import (
	"context"

	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsGlacierVault(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_glacier_vault",
		Description: "AWS Glacier Vault",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("vault_name"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"ResourceNotFoundException", "InvalidParameter"}),
			},
			Hydrate: kaytu.GetGlacierVault,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListGlacierVault,
		},

		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "vault_name",
				Description: "The name of the vault.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Vault.VaultName")},
			{
				Name:        "vault_arn",
				Description: "The Amazon Resource Name (ARN) of the vault.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Vault.VaultARN")},
			{
				Name:        "creation_date",
				Description: "The Universal Coordinated Time (UTC) date when the vault was created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Vault.CreationDate")},
			{
				Name:        "last_inventory_date",
				Description: "The Universal Coordinated Time (UTC) date when Amazon S3 Glacier completed the last vault inventory.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Vault.LastInventoryDate")},
			{
				Name:        "number_of_archives",
				Description: "The number of archives in the vault as of the last inventory date.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Vault.NumberOfArchives")},
			{
				Name:        "size_in_bytes",
				Description: "Total size, in bytes, of the archives in the vault as of the last inventory date.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Vault.SizeInBytes")},
			{
				Name:        "policy",
				Description: "Contains the returned vault access policy as a JSON string.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.AccessPolicy.Policy")},
			{
				Name:        "policy_std",
				Description: "Contains the policy in a canonical form for easier searching.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.AccessPolicy.Policy").Transform(unescape).Transform(policyToCanonical),
			},
			{
				Name:        "vault_lock_policy",
				Description: "The vault lock policy.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.LockPolicy.Policy")},
			{
				Name:        "vault_lock_policy_std",
				Description: "Contains the policy in a canonical form for easier searching.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.LockPolicy.Policy").Transform(unescape).Transform(policyToCanonical),
			},
			{
				Name:        "tags_src",
				Description: "A list of tags associated with the vault.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Tags")},

			// Standard columns
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Vault.VaultName")},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Tags")},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Vault.VaultARN").Transform(arnToAkas),
			},
		}),
	}
}
