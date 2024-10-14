package aws

import (
	"context"

	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAwsLambdaAlias(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_lambda_alias",
		Description: "AWS Lambda Alias",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"name", "function_name", "region"}),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"InvalidParameter", "ResourceNotFoundException"}),
			},
			Hydrate: opengovernance.GetLambdaAlias,
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListLambdaAlias,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "function_version", Require: plugin.Optional},
				{Name: "function_name", Require: plugin.Optional},
			},
		},

		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the alias.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Alias.Name")},
			{
				Name:        "function_name",
				Description: "The name of the function.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.FunctionName")},
			{
				Name:        "alias_arn",
				Description: "The Amazon Resource Name (ARN) of the alias.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Alias.AliasArn")},
			{
				Name:        "function_version",
				Description: "The function version that the alias invokes.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Alias.FunctionVersion")},
			{
				Name:        "revision_id",
				Description: "A unique identifier that changes when you update the alias.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Alias.RevisionId")},
			{
				Name:        "description",
				Description: "A description of the alias.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Alias.Description")},
			{
				Name:        "policy",
				Description: "Contains the resource-based policy.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Policy")},
			{
				Name:        "policy_std",
				Description: "Contains the contents of the resource-based policy in a canonical form for easier searching.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Policy").Transform(unescape).Transform(policyToCanonical),
			},
			{
				Name:        "url_config",
				Description: "The function URL configuration details of the alias.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.UrlConfig")},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Alias.Name")},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Alias.AliasArn").Transform(arnToAkas),
			},
		}),
	}
}
