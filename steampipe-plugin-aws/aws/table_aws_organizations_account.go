package aws

import (
	"context"

	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAwsOrganizationsAccount(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_organizations_account",
		Description: "AWS Organizations Account",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"AccountNotFoundException", "InvalidInputException"}),
			},
			Hydrate: opengovernance.GetOrganizationsAccount,
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListOrganizationsAccount,
		},
		Columns: awsKaytuColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The description of the permission set.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Account.Name")},
			// This description has added text for better clarification on ID type
			{
				Name:        "id",
				Description: "The unique identifier (account ID) of the member account.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Account.Id")},
			{
				Name:        "id",
				Description: "The unique identifier (account ID) of the member account.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Account.Id")},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the account.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Account.Arn")},
			{
				Name:        "status",
				Description: "The status of the account in the organization.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Account.Status")},
			{
				Name:        "email",
				Description: "The email address associated with the AWS account.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Account.Email")},
			{
				Name:        "joined_method",
				Description: "The method by which the account joined the organization.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Account.JoinedMethod")},
			{
				Name:        "joined_timestamp",
				Description: "The date the account became a part of the organization.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Account.JoinedTimestamp")},
			{
				Name:      "tags_src",
				Type:      proto.ColumnType_JSON,
				Transform: transform.FromField("Description.Tags")},
			// Standard columns for all tables
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(getOrganizationsResourceTurbotTags),
			},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Account.Name")},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Account.Arn").Transform(transform.EnsureStringArray),
				//// LIST FUNCTION
			},
		}),
	}
}

//// TRANSFORM FUNCTIONS

func getOrganizationsResourceTurbotTags(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	tags := d.HydrateItem.(opengovernance.OrganizationsAccount).Description.Tags
	tagsMap := map[string]string{}

	for _, tag := range tags {
		tagsMap[*tag.Key] = *tag.Value
	}

	return tagsMap, nil
}
