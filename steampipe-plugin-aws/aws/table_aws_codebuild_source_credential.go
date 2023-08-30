package aws

import (
	"context"
	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"
	"strings"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsCodeBuildSourceCredential(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_codebuild_source_credential",
		Description: "AWS CodeBuild Source Credential",
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListCodeBuildSourceCredential,
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the token.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.SourceCredentialsInfo.Arn")},
			{
				Name:        "auth_type",
				Description: "The type of authentication used by the credentials. Possible values are: OAUTH, BASIC_AUTH, or PERSONAL_ACCESS_TOKEN.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.SourceCredentialsInfo.AuthType")},
			{
				Name:        "server_type",
				Description: "The type of source provider. Possible values are: GITHUB, GITHUB_ENTERPRISE, or BITBUCKET.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.SourceCredentialsInfo.ServerType")},
			// Steampipe standard columns

			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.From(codebuildSourceCredentialTitle),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.SourceCredentialsInfo.Arn").Transform(transform.EnsureStringArray),
			},
		}),
	}
}

//// LIST FUNCTION

//// TRANSFORM FUNCTIONS

func codebuildSourceCredentialTitle(_ context.Context, d *transform.TransformData) (interface{},
	error) {
	data := d.HydrateItem.(kaytu.CodeBuildSourceCredential).Description.SourceCredentialsInfo

	splitPart := strings.Split(*data.Arn, ":")
	title := string(data.ServerType) + " - " + splitPart[3]

	return title, nil
}
