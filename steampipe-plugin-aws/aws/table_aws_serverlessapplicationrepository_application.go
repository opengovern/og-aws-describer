package aws

import (
	"context"

	"github.com/kaytu-io/og-aws-describer/pkg/opengovernance-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAwsServerlessApplicationRepositoryApplication(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_serverlessapplicationrepository_application",
		Description: "AWS Serverless Application Repository Application",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("arn"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"InvalidParameter", "NotFoundException"}),
			},
			Hydrate: opengovernance.GetServerlessApplicationRepositoryApplication,
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListServerlessApplicationRepositoryApplication,
		},

		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the application.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Application.Name")},
			{
				Name:        "arn",
				Description: "The application Amazon Resource Name (ARN).",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Application.ApplicationId"),
			},
			{
				Name:        "author",
				Description: "The name of the author publishing the app.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Application.Author")},
			{
				Name:        "creation_time",
				Description: "The date and time this resource was created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Application.CreationTime")},
			{
				Name:        "description",
				Description: "The description of the application.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Application.Description")},
			{
				Name:        "home_page_url",
				Description: "A URL with more information about the application.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Application.HomePageUrl")},
			{
				Name:        "is_verified_author",
				Description: "Whether the author is verified.",
				Type:        proto.ColumnType_BOOL,
				Default:     false,
				Transform:   transform.FromField("Description.Application.IsVerifiedAuthor")},
			{
				Name:        "license_url",
				Description: "The URL of the license.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Application.LicenseUrl")},
			{
				Name:        "readme_url",
				Description: "The URL of the Readme.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Application.ReadmeUrl")},
			{
				Name:        "spdx_license_id",
				Description: "A valid identifier from https://spdx.org/licenses/.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Application.SpdxLicenseId")},
			{
				Name:        "verified_author_url",
				Description: "The URL of the verified author.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Application.VerifiedAuthorUrl")},
			{
				Name:        "labels",
				Description: "Labels to improve discovery of apps in search results.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Application.Labels")},
			{
				Name:        "statements",
				Description: "The contents of the access policy.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Statements")},
			{
				Name:        "version",
				Description: "The policy statement of the application.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Application.Version")},

			// Standard columns for all tables
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Application.Name")},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Application.ApplicationId").Transform(transform.EnsureStringArray),
			},
		}),
	}
}
