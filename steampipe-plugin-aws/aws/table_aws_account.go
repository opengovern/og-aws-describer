package aws

import (
	"context"
	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsAccount(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_account",
		Description: "AWS Account",
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListIAMAccount,
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "account_aliases",
				Description: "A list of aliases associated with the account, if applicable.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Aliases"),
			},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) specifying the account.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.From(accountARN),
			},
			{
				Name:        "organization_id",
				Description: "The unique identifier (ID) of an organization, if applicable.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Organization.Id"),
			},
			{
				Name:        "organization_arn",
				Description: "The Amazon Resource Name (ARN) of an organization.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Organization.Arn"),
			},
			{
				Name:        "organization_feature_set",
				Description: "Specifies the functionality that currently is available to the organization. If set to \"ALL\", then all features are enabled and policies can be applied to accounts in the organization. If set to \"CONSOLIDATED_BILLING\", then only consolidated billing functionality is available.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Organization.FeatureSet"),
			},
			{
				Name:        "organization_master_account_arn",
				Description: "The Amazon Resource Name (ARN) of the account that is designated as the management account for the organization",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Organization.MasterAccountArn"),
			},
			{
				Name:        "organization_master_account_email",
				Description: "The email address that is associated with the AWS account that is designated as the management account for the organization",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Organization.MasterAccountEmail"),
			},
			{
				Name:        "organization_master_account_id",
				Description: "The unique identifier (ID) of the management account of an organization",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Organization.MasterAccountId"),
			},
			{
				Name:        "organization_available_policy_types",
				Description: "The Region opt-in status. The possible values are opt-in-not-required, opted-in, and not-opted-in",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Organization.AvailablePolicyTypes"),
			},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.From(accountDataToTitle),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(accountARN).Transform(transform.EnsureStringArray),
			},
		}),
	}
}

//// LIST FUNCTION

//// HYDRATE FUNCTIONS

//// Transform Functions

func accountDataToTitle(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	plugin.Logger(ctx).Trace("getAwsAccountAkas")
	accountInfo := d.HydrateItem.(kaytu.IAMAccount)

	if len(accountInfo.Description.Aliases) > 0 {
		return accountInfo.Description.Aliases[0], nil
	}

	return accountInfo.Metadata.AccountID, nil
}

func accountARN(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	plugin.Logger(ctx).Trace("accountARN")
	metadata := d.HydrateItem.(kaytu.IAMAccount).Metadata

	arn := "arn:" + metadata.Partition + ":::" + metadata.AccountID

	return arn, nil
}
