package aws

import (
	"context"
	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsEcrRepository(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_ecr_repository",
		Description: "AWS ECR Repository",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("repository_name"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"RepositoryNotFoundException", "RepositoryPolicyNotFoundException", "LifecyclePolicyNotFoundException"}),
			},
			Hydrate: kaytu.GetECRRepository,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListECRRepository,
		},

		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "repository_name",
				Description: "The name of the repository.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Repository.RepositoryName")},
			{
				Name:        "registry_id",
				Description: "The AWS account ID associated with the registry that contains the repositories to be described.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Repository.RegistryId")},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) that identifies the repository.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Repository.RepositoryArn"),
			},
			{
				Name:        "repository_uri",
				Description: "The URI for the repository.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Repository.RepositoryUri")},
			{
				Name:        "created_at",
				Description: "The date and time, in JavaScript date format, when the repository was created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Repository.CreatedAt")},
			{
				Name:        "image_tag_mutability",
				Description: "The tag mutability setting for the repository.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Repository.ImageTagMutability")},
			{
				Name:        "last_evaluated_at",
				Description: "The time stamp of the last time that the lifecycle policy was run.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.LifecyclePolicy.LastEvaluatedAt")},
			{
				Name:        "max_results",
				Description: "The maximum number of repository results returned by DescribeRepositories.",
				Hydrate:     kaytu.GetECRRepository,
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "encryption_configuration",
				Description: "The encryption configuration for the repository.",
				Transform:   transform.FromField("Description.Repository.EncryptionConfiguration"),
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "image_details",
				Description: "A list of ImageDetail objects that contain data about the image.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ImageDetails")},
			{
				Name:        "repository_scanning_configuration",
				Description: "Gets the scanning configuration for one or more repositories.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.RepositoryScanningConfiguration")},
			{
				Name:        "image_scanning_configuration",
				Description: "The image scanning configuration for a repository.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Repository.ImageScanningConfiguration")},
			{
				Name:        "image_scanning_findings",
				Description: "Scan findings for an image.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ImageScanFinding")},
			{
				Name:        "lifecycle_policy",
				Description: "The JSON lifecycle policy text.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.LifecyclePolicy.LifecyclePolicyText")},
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
				Description: "A list of tags assigned to the Repository.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Tags")},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Tags").Transform(ecrTagListToTurbotTags),
			},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Repository.RepositoryName")},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Repository.RepositoryArn").Transform(arnToAkas),
			},
		}),
	}
}

//// TRANSFORM FUNCTIONS

func ecrTagListToTurbotTags(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	plugin.Logger(ctx).Trace("ecrTagListToTurbotTags")
	tags := d.HydrateItem.(kaytu.ECRRepository).Description.Tags

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
