package aws

import (
	"context"
	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsEcrpublicRepository(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_ecrpublic_repository",
		Description: "AWS ECR Public Repository",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("repository_name"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"RepositoryNotFoundException", "RepositoryPolicyNotFoundException", "LifecyclePolicyNotFoundException"}),
			},
			Hydrate: opengovernance.GetECRPublicRepository,
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListECRPublicRepository,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "registry_id", Require: plugin.Optional},
			},
		},

		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "repository_name",
				Description: "The name of the repository.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.PublicRepository.RepositoryName")},
			{
				Name:        "registry_id",
				Description: "The AWS account ID associated with the public registry that contains the repository.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.PublicRepository.RegistryId")},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) that identifies the repository.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.PublicRepository.RepositoryArn"),
			},
			{
				Name:        "repository_uri",
				Description: "The URI for the repository.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.PublicRepository.RepositoryUri")},
			{
				Name:        "created_at",
				Description: "The date and time, in JavaScript date format, when the repository was created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.PublicRepository.CreatedAt")},
			{
				Name:        "image_details",
				Description: "A list of ImageDetail objects that contain data about the image.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ImageDetails")},
			{
				Name:        "policy",
				Description: "The JSON repository policy text associated with the repository.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Policy.PolicyText")},
			{
				Name:        "policy_std",
				Description: "Contains the policy in a canonical form for easier searching.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Policy.PolicyText").Transform(unescape).Transform(policyToCanonical),
			},
			{
				Name:        "tags_src",
				Description: "A list of tags assigned to the repository.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Tags")},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.PublicRepository.RepositoryName")},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Tags").Transform(ecrpublicTagListToTurbotTags),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.PublicRepository.RepositoryArn").Transform(arnToAkas),
			},
		}),
	}
}

//// TRANSFORM FUNCTIONS

func ecrpublicTagListToTurbotTags(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	plugin.Logger(ctx).Trace("ecrpublicTagListToTurbotTags")
	tags := d.HydrateItem.(opengovernance.ECRPublicRepository).Description.Tags

	if tags == nil {
		return nil, nil
	}

	// Mapping the resource tags inside turbotTags
	var turbotTagsMap map[string]string
	if tags != nil {
		turbotTagsMap = map[string]string{}
		for _, i := range tags {
			turbotTagsMap[*i.Key] = *i.Value
		}
	}

	return turbotTagsMap, nil
}
