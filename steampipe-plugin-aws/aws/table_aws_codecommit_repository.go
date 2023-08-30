package aws

import (
	"context"

	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsCodeCommitRepository(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_codecommit_repository",
		Description: "AWS CodeCommit Repository",
		List: &plugin.ListConfig{
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"InvalidParameter"}),
			},
			Hydrate: kaytu.ListCodeCommitRepository,
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "repository_name",
				Description: "The repository's name.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Repository.RepositoryName")},
			{
				Name:        "repository_id",
				Description: "The ID of the repository.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Repository.RepositoryId")},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the repository.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Repository.Arn")},
			{
				Name:        "description",
				Description: "A comment or description about the repository.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Repository.RepositoryDescription")},
			{
				Name:        "creation_date",
				Description: "The date and time the repository was created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Repository.CreationDate")},
			{
				Name:        "clone_url_http",
				Description: "The URL to use for cloning the repository over HTTPS.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Repository.CloneUrlHttp")},
			{
				Name:        "clone_url_ssh",
				Description: "The URL to use for cloning the repository over SSH.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Repository.CloneUrlSsh")},
			{
				Name:        "default_branch",
				Description: "The repository's default branch name.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Repository.DefaultBranch")},
			{
				Name:        "last_modified_date",
				Description: "The date and time the repository was last modified.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Repository.LastModifiedDate")},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Repository.RepositoryName")},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Tags")},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Repository.Arn").Transform(arnToAkas),
			},
		}),
	}
}
