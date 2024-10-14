package aws

import (
	"context"
	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAwsSSMDocumentPermission(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_ssm_document_permission",
		Description: "AWS SSM Document Permission",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListSSMDocumentPermission,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "document_name", Require: plugin.Required},
			},
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "document_name",
				Description: "The name of the Systems Manager document.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Document.Document.Name")},
			{
				Name:        "shared_account_id",
				Description: "The Amazon Web Services account ID where the current document is shared.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Permissions.AccountSharingInfoList.AccountId")},
			{
				Name:        "shared_document_version",
				Description: "The version of the current document shared with the account.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Permissions.AccountSharingInfoList.SharedDocumentVersion")},
			{
				Name:        "account_ids",
				Description: "The account IDs that have permission to use this document. The ID can be either an AWS account or All.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Permissions.AccountIds")},
			// Steampipe standard columns
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Permissions.AccountSharingInfoList.SharedDocumentVersion")},
		}),
	}
}
