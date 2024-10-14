package aws

import (
	"context"
	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsVpcVerifiedAccessTrustProvider(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_vpc_verified_access_trust_provider",
		Description: "AWS VPC Verified Access Trust Provider",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListEC2VerifiedAccessTrustProvider,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"InvalidParameterValue"}),
			},
			KeyColumns: []*plugin.KeyColumn{
				{Name: "verified_access_trust_provider_id", Require: plugin.Optional},
			},
		},

		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "verified_access_trust_provider_id",
				Description: "The ID of the AWS Verified Access trust provider.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.VerifiedAccountGroup.VerifiedAccessTrustProviderId")},
			{
				Name:        "creation_time",
				Description: "The creation time.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.VerifiedAccountGroup.CreationTime")},
			{
				Name:        "description",
				Description: "A description for the AWS Verified Access trust provider.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.VerifiedAccountGroup.Description")},
			{
				Name:        "device_trust_provider_type",
				Description: "The type of device-based trust provider.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.VerifiedAccountGroup.DeviceTrustProviderType")},
			{
				Name:        "last_updated_time",
				Description: "The last updated time.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.VerifiedAccountGroup.LastUpdatedTime")},
			{
				Name:        "policy_reference_name",
				Description: "The identifier to be used when working with policy rules.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.VerifiedAccountGroup.PolicyReferenceName")},
			{
				Name:        "trust_provider_type",
				Description: "The type of Verified Access trust provider.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.VerifiedAccountGroup.TrustProviderType")},
			{
				Name:        "user_trust_provider_type",
				Description: "The type of user-based trust provider.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.VerifiedAccountGroup.UserTrustProviderType")},
			{
				Name:        "oidc_options",
				Description: "The OpenID Connect details for an oidc-type, user-identity based trust provider.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.VerifiedAccountGroup.OidcOptions")},
			{
				Name:        "tags_src",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.VerifiedAccountGroup.Tags")},

			// Steampipe standard columns

			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.From(trustProviderTitle),
			},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.VerifiedAccountGroup.Tags").Transform(trustProviderTurbotTags),
			},
		}),
	}
}

//// TRANSFORM FUNCTIONS

func trustProviderTurbotTags(_ context.Context, d *transform.TransformData) (interface{}, error) {
	accessPoint := d.HydrateItem.(opengovernance.EC2VerifiedAccessTrustProvider).Description.VerifiedAccessTrustProvider

	// Get the resource tags
	var turbotTagsMap map[string]string
	if accessPoint.Tags != nil {
		turbotTagsMap = map[string]string{}
		for _, i := range accessPoint.Tags {
			turbotTagsMap[*i.Key] = *i.Value
		}
	}

	return turbotTagsMap, nil
}

func trustProviderTitle(_ context.Context, d *transform.TransformData) (interface{}, error) {
	accessPoint := d.HydrateItem.(opengovernance.EC2VerifiedAccessTrustProvider).Description.VerifiedAccessTrustProvider
	title := accessPoint.VerifiedAccessTrustProviderId

	if accessPoint.Tags != nil {
		for _, i := range accessPoint.Tags {
			if *i.Key == "Name" {
				title = i.Value
			}
		}
	}

	return title, nil
}
