package aws

import (
	"context"

	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsCodeArtifactRepository(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_codeartifact_repository",
		Description: "AWS CodeArtifact Repository",
		Get: &plugin.GetConfig{
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:    "name",
					Require: plugin.Required,
				},
				{
					Name:    "domain_name",
					Require: plugin.Required,
				},
				{
					Name:    "domain_owner",
					Require: plugin.Optional,
				},
			},
			Hydrate: kaytu.GetCodeArtifactRepository,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListCodeArtifactRepository,
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the repository.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Repository.Name")},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the repository",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Repository.Arn")},
			{
				Name:        "domain_name",
				Description: "The name of the domain that contains the repository.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Repository.DomainName")},
			{
				Name:        "administrator_account",
				Description: "The Amazon Web Services account ID that manages the repository.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Repository.AdministratorAccount")},
			{
				Name:        "description",
				Description: "The description of the repository.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Repository.Description")},
			{
				Name:        "domain_owner",
				Description: "The 12-digit account number of the Amazon Web Services account that owns the repository. It does not include dashes or spaces.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Repository.DomainOwner")},
			{
				Name:        "external_connections",
				Description: "An array of external connections associated with the repository.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Description.ExternalConnections")},
			{
				Name:        "policy",
				Description: "An CodeArtifact resource policy that contains a resource ARN, document details, and a revision.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Policy.ResourceArn")},
			{
				Name:        "policy_std",
				Description: "Contains the contents of the resource-based policy in a canonical form for easier searching.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Policy")},
			{
				Name:        "repository_endpoint",
				Description: "A string that specifies the URL of the returned endpoint.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Endpoints"),
			},
			{
				Name:        "upstreams",
				Description: "A list of upstream repositories to associate with the repository.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Description.Upstreams")},
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
				Transform:   transform.FromField("Description.Repository.Name")},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(getCodeArtifactRepositoryTurbotTags),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Repository.Arn").Transform(arnToAkas),
			},
		}),
	}
}

//// TRANSFORM FUNCTIONS

func getCodeArtifactRepositoryTurbotTags(_ context.Context, d *transform.TransformData) (interface{}, error) {
	tags := d.HydrateItem.(kaytu.CodeArtifactRepository).Description.Tags
	return codeArtifactV2TagsToMap(tags)
}
