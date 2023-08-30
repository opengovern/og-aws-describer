package aws

import (
	"context"
	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsAccountAlternateContact(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_account_alternate_contact",
		Description: "AWS Account Alternate Contact",
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListAccountAlternateContact,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"ResourceNotFoundException"}),
			},
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:       "linked_account_id",
					Require:    plugin.Optional,
					CacheMatch: "exact",
				},
				{
					Name:    "contact_type",
					Require: plugin.Optional,
				},
			},
		},
		Columns: awsKaytuColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The name associated with this alternate contact.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AlternateContact.Name")},
			{
				Name:        "contact_type",
				Description: "The type of alternate contact.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AlternateContact.AlternateContactType")},
			{
				Name:        "email_address",
				Description: "The email address associated with this alternate contact.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AlternateContact.EmailAddress")},
			{
				Name:        "phone_number",
				Description: "The phone number associated with this alternate contact.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AlternateContact.PhoneNumber")},
			{
				Name:        "contact_title",
				Description: "The title associated with this alternate contact.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AlternateContact.Title")},
			{
				Name:        "linked_account_id",
				Description: "Account ID to get alternate contact details for.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.LinkedAccountID")},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AlternateContact.Name")},
		}),
	}
}
