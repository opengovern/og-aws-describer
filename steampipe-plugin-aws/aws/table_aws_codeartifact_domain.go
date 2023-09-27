package aws

import (
	"context"

	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsCodeArtifactDomain(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_codeartifact_domain",
		Description: "AWS Code Artifact Domain",
		Get: &plugin.GetConfig{
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:    "name",
					Require: plugin.Required,
				},
				{
					Name:    "owner",
					Require: plugin.Optional,
				},
			},
			Hydrate: kaytu.GetCodeArtifactDomain,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListCodeArtifactDomain,
		},

		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the domain.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Domain.Name")},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) specifying the domain.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Domain.Arn")},
			{
				Name:        "asset_size_bytes",
				Description: "The total size of all assets in the domain.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Domain.AssetSizeBytes")},
			{
				Name:        "created_time",
				Description: "A timestamp that contains the date and time the domain was created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Domain.CreatedTime")},
			{
				Name:        "encryption_key",
				Description: "The key used to encrypt the domain.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Domain.EncryptionKey")},
			{
				Name:        "owner",
				Description: "The 12-digit account number of the Amazon Web Services account that owns the domain. It does not include dashes or spaces.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Domain.Owner")},
			{
				Name:        "repository_count",
				Description: "The number of repositories in the domain.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Domain.RepositoryCount")},
			{
				Name:        "s3_bucket_arn",
				Description: "The Amazon Resource Name (ARN) of the Amazon S3 bucket that is used to store package assets in the domain.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Domain.S3BucketArn")},
			{
				Name:        "status",
				Description: "A string that contains the status of the domain.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Domain.Status")},
			{
				Name:        "policy",
				Description: "An CodeArtifact resource policy that contains a resource ARN, document details, and a revision.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Policy")},
			{
				Name:        "policy_std",
				Description: "Contains the contents of the resource-based policy in a canonical form for easier searching.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Policy").Transform(unescape).Transform(policyToCanonical),
			},
			{
				Name:        "tags_src",
				Description: "A list of tags assigned to the resource.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Tags")},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Domain.Name")},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(codeArtifactDomainTurbotTags),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Domain.Arn").Transform(transform.EnsureStringArray),
			},
		}),
	}
}

//// TRANSFORM FUNCTIONS

func codeArtifactDomainTurbotTags(_ context.Context, d *transform.TransformData) (interface{}, error) {
	tags := d.HydrateItem.(kaytu.CodeArtifactDomain).Description.Tags

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
