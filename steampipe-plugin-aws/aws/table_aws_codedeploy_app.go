package aws

import (
	"context"
	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsCodeDeployApplication(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_codedeploy_app",
		Description: "AWS CodeDeploy Application",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("application_name"),
			Hydrate:    kaytu.GetCodeDeployApplication,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListCodeDeployApplication,
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "application_id",
				Description: "The application ID.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Application.ApplicationId"),
			},
			{
				Name:        "application_name",
				Description: "The application name.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Application.ApplicationName")},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) specifying the application.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("ARN")},
			{
				Name:        "compute_platform",
				Description: "The destination platform type for deployment of the application (Lambda or Server).",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Application.ComputePlatform")},
			{
				Name:        "create_time",
				Description: "The time at which the application was created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Application.CreateTime")},
			{
				Name:        "github_account_name",
				Description: "The name for a connection to a GitHub account.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Application.GitHubAccountName")},
			{
				Name:        "linked_to_github",
				Description: "True if the user has authenticated with GitHub for the specified application. Otherwise, false.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Application.LinkedToGitHub")},
			{
				Name:        "tags_src",
				Description: "A list of tag key and value pairs associated with this application.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Tags")},

			// Steampipe standard columns
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(getCodeDeployApplicationTurbotTags),
			},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Application.ApplicationName")},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("ARN").Transform(arnToAkas),
			},
		}),
	}
}

//// TRANSFORM FUNCTIONS

func getCodeDeployApplicationTurbotTags(_ context.Context, d *transform.TransformData) (interface{}, error) {
	tags := d.HydrateItem.(kaytu.CodeDeployApplication).Description.Tags
	return codeDeployV2TagsToMap(tags)
}
