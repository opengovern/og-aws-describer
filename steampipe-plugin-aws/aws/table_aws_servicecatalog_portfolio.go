package aws

import (
	"context"
	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsServicecatalogPortfolio(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_servicecatalog_portfolio",
		Description: "AWS Service Catalog Portfolio",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"ResourceNotFoundException"}),
			},
			Hydrate: kaytu.GetServiceCatalogPortfolio,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListServiceCatalogPortfolio,
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "display_name",
				Description: "The name to use for display purposes.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Portfolio.DisplayName")},
			{
				Name:        "id",
				Description: "The portfolio identifier.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Portfolio.Id")},
			{
				Name:        "arn",
				Description: "The ARN assigned to the portfolio.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Portfolio.ARN")},
			{
				Name:        "created_time",
				Description: "The UTC timestamp of the creation time.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Portfolio.CreatedTime")},
			{
				Name:        "description",
				Description: "The description of the portfolio.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Portfolio.Description")},
			{
				Name:        "provider_name",
				Description: "The name of the portfolio provider.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Portfolio.ProviderName")},
			{
				Name:        "budgets",
				Description: "Information about the associated budgets.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Budgets")},
			{
				Name:        "tag_options",
				Description: "Information about the tag options associated with the portfolio.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.TagOptions")},
			{
				Name:        "tags_src",
				Description: "Information about the tags associated with the portfolio.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Tag")},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Portfolio.DisplayName")},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Tag")},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Portfolio.ARN").Transform(transform.EnsureStringArray)},
		}),
	}
}
