package aws

import (
	"context"

	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsIamSamlProvider(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_iam_saml_provider",
		Description: "AWS IAM Saml Provider",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"arn"}),
			Hydrate:    opengovernance.GetIAMSamlProvider,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"NoSuchEntity"}),
			},
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListIAMSamlProvider,
		},
		Columns: awsKaytuColumns([]*plugin.Column{
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) specifying the IAM policy.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("ARN")},
			{
				Name:        "create_date",
				Description: "The date and time when the SAML provider was created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.SamlProvider.CreateDate")},
			{
				Name:        "valid_until",
				Description: "The expiration date and time for the SAML provider.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.SamlProvider.ValidUntil")},
			{
				Name:        "saml_metadata_document",
				Description: "The XML metadata document that includes information about an identity provider.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.SamlProvider.SAMLMetadataDocument")},
			{
				Name:        "tags_src",
				Description: "A list of tags that are attached to the specified IAM SAML provider.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.SamlProvider.Tags")},

			// Steampipe standard columns
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(samlProviderTurbotTags),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("ARN").Transform(transform.EnsureStringArray),
			},
		}),
	}
}

//// TRANSFORM FUNCTION

func samlProviderTurbotTags(_ context.Context, d *transform.TransformData) (interface{}, error) {
	provider := d.HydrateItem.(opengovernance.IAMSamlProvider).Description.SamlProvider
	if len(provider.Tags) == 0 {
		return nil, nil
	}

	turbotTagsMap := map[string]string{}
	for _, i := range provider.Tags {
		turbotTagsMap[*i.Key] = *i.Value
	}
	return turbotTagsMap, nil
}
