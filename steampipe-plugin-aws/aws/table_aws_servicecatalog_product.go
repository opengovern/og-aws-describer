package aws

import (
	"context"
	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"

	servicecatalogv1 "github.com/aws/aws-sdk-go/service/servicecatalog"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsServicecatalogProduct(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_servicecatalog_product",
		Description: "AWS Service Catalog Product",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("product_id"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"ResourceNotFoundException"}),
			},
			Hydrate: kaytu.GetServiceCatalogProduct,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListServiceCatalogProduct,
			KeyColumns: plugin.KeyColumnSlice{
				{
					Name:    "accept_language",
					Require: plugin.Optional,
				},
				{
					Name:    "full_text_search",
					Require: plugin.Optional,
				},
				{
					Name:    "owner",
					Require: plugin.Optional,
				},
				{
					Name:    "type",
					Require: plugin.Optional,
				},
				{
					Name:    "source_product_id",
					Require: plugin.Optional,
				},
			},
		},
		GetMatrixItemFunc: SupportedRegionMatrix(servicecatalogv1.EndpointsID),
		Columns: awsRegionalColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the product.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ProductViewSummary.Name")},
			{
				Name:        "id",
				Description: "The product view identifier.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ProductViewSummary.Id")},
			{
				Name:        "product_id",
				Description: "The product identifier.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ProductViewSummary.ProductId")},
			{
				Name:        "source_product_id",
				Description: "The source product identifier.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromQual("source_product_id"),
			},
			{
				Name:        "distributor",
				Description: "The distributor of the product. Contact the product administrator for the significance of this value.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ProductViewSummary.Distributor")},
			{
				Name:        "accept_language",
				Description: "The language code.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromQual("accept_language"),
			},
			{
				Name:        "full_text_search",
				Description: "The full text for the product.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromQual("full_text_search"),
			},
			{
				Name:        "has_default_path",
				Description: "Indicates whether the product has a default path. If the product does not have a default path, call ListLaunchPaths to disambiguate between paths.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.ProductViewSummary.HasDefaultPath")},
			{
				Name:        "owner",
				Description: "The owner of the product. Contact the product administrator for the significance of this value.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ProductViewSummary.Owner")},
			{
				Name:        "short_description",
				Description: "Short description of the product.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ProductViewSummary.ShortDescription")},
			{
				Name:        "support_description",
				Description: "The description of the support for this product.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ProductViewSummary.SupportDescription")},
			{
				Name:        "support_email",
				Description: "The email contact information to obtain support for this product.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ProductViewSummary.SupportEmail")},
			{
				Name:        "support_url",
				Description: "The URL information to obtain support for this product.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ProductViewSummary.SupportUrl")},
			{
				Name:        "type",
				Description: "The product type. Contact the product administrator for the significance of this value. If this value is MARKETPLACE, the product was created by Amazon Web Services Marketplace.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ProductViewSummary.Type")},
			{
				Name:        "budgets",
				Description: "Information about the associated budgets.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Budgets"),
			},
			{
				Name:        "launch_paths",
				Description: "Information about the associated launch paths.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.LunchPaths"),
			},
			{
				Name:        "provisioning_artifacts",
				Description: "Information about the provisioning artifacts for the specified product.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ProvisioningArtifacts"),
			},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ProductViewSummary.Name")},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("ARN").Transform(transform.EnsureStringArray),
			},
		}),
	}
}
