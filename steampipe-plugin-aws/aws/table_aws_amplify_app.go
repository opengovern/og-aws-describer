package aws

import (
	"context"
	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsAmplifyApp(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_amplify_app",
		Description: "AWS Amplify App",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("app_id"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"ValidationException", "NotFoundException"}),
			},
			Hydrate: kaytu.GetAmplifyApp,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListAmplifyApp,
		},

		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "app_id",
				Description: "The unique ID of the Amplify app.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.App.AppId")},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the Amplify app.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.App.AppArn"),
			},
			{
				Name:        "name",
				Description: "The name for the Amplify app.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.App.Name")},
			{
				Name:        "description",
				Description: "The description for the Amplify app.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.App.Description")},
			{
				Name:        "create_time",
				Description: "Creates a date and time for the Amplify app.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.App.CreateTime")},
			{
				Name:        "update_time",
				Description: "Updates the date and time for the Amplify app.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.App.UpdateTime")},
			{
				Name:        "basic_auth_credentials",
				Description: "The basic authorization credentials for branches for the Amplify app. You must base64-encode the authorization credentials and provide them in the format user:password.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.App.BasicAuthCredentials")},
			{
				Name:        "custom_headers",
				Description: "Describes the custom HTTP headers for the Amplify app.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.App.CustomHeaders")},
			{
				Name:        "default_domain",
				Description: "The default domain for the Amplify app.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.App.DefaultDomain")},
			{
				Name:        "enable_auto_branch_creation",
				Description: "Enables automated branch creation for the Amplify app.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.App.EnableAutoBranchCreation")},
			{
				Name:        "enable_basic_auth",
				Description: "Enables basic authorization for the Amplify app's branches.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.App.EnableBasicAuth")},
			{
				Name:        "enable_branch_auto_build",
				Description: "Enables the auto-building of branches for the Amplify app.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.App.EnableBranchAutoBuild")},
			{
				Name:        "enable_branch_auto_deletion",
				Description: "Automatically disconnect a branch in the Amplify Console when you delete a branch from your Git repository.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.App.EnableBranchAutoDeletion")},
			{
				Name:        "iam_service_role_arn",
				Description: "The AWS Identity and Access Management (IAM) service role for the Amazon Resource Name (ARN) of the Amplify app.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.App.IamServiceRoleArn")},
			{
				Name:        "platform",
				Description: "The platform for the Amplify app.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.App.Platform")},
			{
				Name:        "repository",
				Description: "The Git repository for the Amplify app.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.App.Repository")},
			{
				Name:        "repository_clone_method",
				Description: "The Amplify service uses this parameter to specify the authentication protocol to use to access the Git repository for an Amplify app. Amplify specifies TOKEN for a GitHub repository, SIGV4 for an AWS CodeCommit repository, and SSH for GitLab and Bitbucket repositories.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.App.RepositoryCloneMethod")},
			{
				Name:        "auto_branch_creation_config",
				Description: "Describes the automated branch creation configuration for the Amplify app.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.App.AutoBranchCreationConfig")},
			{
				Name:        "auto_branch_creation_patterns",
				Description: "Describes the automated branch creation glob patterns for the Amplify app.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.App.AutoBranchCreationPatterns")},
			{
				Name:        "build_spec",
				Description: "Describes the content of the build specification (build spec) for the Amplify app.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.App.BuildSpec").Transform(transform.UnmarshalYAML),
			},
			{
				Name:        "custom_rules",
				Description: "Describes the custom redirect and rewrite rules for the Amplify app.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.App.CustomRules")},
			{
				Name:        "environment_variables",
				Description: "The environment variables for the Amplify app.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.App.EnvironmentVariables")},
			{
				Name:        "production_branch",
				Description: "Describes the information about a production branch of the Amplify app.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.App.ProductionBranch")},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.App.Name")},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.App.Tags")},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.App.AppArn").Transform(transform.EnsureStringArray),
			},
		}),
	}
}
