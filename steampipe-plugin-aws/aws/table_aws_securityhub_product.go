package aws

import (
	"context"

	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsSecurityhubProduct(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_securityhub_product",
		Description: "AWS Securityhub Product",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("product_arn"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"InvalidInputException"}),
			},
			Hydrate: kaytu.GetSecurityHubProduct,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListSecurityHubProduct,
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the product.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Product.ProductName")},
			{
				Name:        "product_arn",
				Description: "The ARN assigned to the product.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Product.ProductArn")},
			{
				Name:        "activation_url",
				Description: "The URL used to activate the product.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Product.ActivationUrl")},
			{
				Name:        "company_name",
				Description: "The name of the company that provides the product.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Product.CompanyName")},
			{
				Name:        "description",
				Description: "A description of the product.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Product.Description")},
			{
				Name:        "marketplace_url",
				Description: "The URL for the page that contains more information about the product.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Product.MarketplaceUrl")},
			{
				Name:        "categories",
				Description: "The categories assigned to the product.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Product.Categories")},
			{
				Name:        "integration_types",
				Description: "The types of integration that the product supports.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Product.IntegrationTypes")},
			{
				Name:        "product_subscription_resource_policy",
				Description: "The resource policy associated with the product.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Product.ProductSubscriptionResourcePolicy")},

			// Standard columns
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Product.ProductName")},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Product.ProductArn").Transform(arnToAkas),
			},
		}),
	}
}
