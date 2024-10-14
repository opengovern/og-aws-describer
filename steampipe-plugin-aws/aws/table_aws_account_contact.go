package aws

import (
	"context"
	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsAccountContact(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_account_contact",
		Description: "AWS Account Contact",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListAccountContact,
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:       "linked_account_id",
					Require:    plugin.Optional,
					CacheMatch: "exact",
				},
			},
		},
		Columns: awsKaytuColumns([]*plugin.Column{
			{
				Name:        "full_name",
				Description: "The full name of the primary contact address.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AlternateContact.FullName")},
			{
				Name:        "address_line_1",
				Description: "The first line of the primary contact address",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AlternateContact.AddressLine1")},
			{
				Name:        "address_line_2",
				Description: "The second line of the primary contact address, if any.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AlternateContact.AddressLine2")},
			{
				Name:        "address_line_3",
				Description: "The third line of the primary contact address, if any.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AlternateContact.AddressLine3")},
			{
				Name:        "company_name",
				Description: "The name of the company associated with the primary contact information, if any.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AlternateContact.CompanyName")},
			{
				Name:        "city",
				Description: "The city of the primary contact address.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AlternateContact.City")},
			{
				Name:        "country_code",
				Description: "The ISO-3166 two-letter country code for the primary contact address.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AlternateContact.CountryCode")},
			{
				Name:        "district_or_county",
				Description: "The district or county of the primary contact address, if any.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AlternateContact.DistrictOrCounty")},
			{
				Name:        "phone_number",
				Description: "The phone number of the primary contact information. The number will be validated and, in some countries, checked for activation.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AlternateContact.PhoneNumber")},
			{
				Name:        "postal_code",
				Description: "The postal code of the primary contact address.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AlternateContact.PostalCode")},
			{
				Name:        "state_or_region",
				Description: "The state or region of the primary contact address. This field is required in selected countries.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AlternateContact.StateOrRegion")},
			{
				Name:        "website_url",
				Description: "The URL of the website associated with the primary contact information, if any.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AlternateContact.WebsiteUrl")},
			{
				Name:        "linked_account_id",
				Description: "Account ID to get contact details for.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.LinkedAccountID")},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AlternateContact.FullName")},
		}),
	}
}
