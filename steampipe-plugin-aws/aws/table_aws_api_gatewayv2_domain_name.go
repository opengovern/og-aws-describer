package aws

import (
	"context"

	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsAPIGatewayV2DomainName(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_api_gatewayv2_domain_name",
		Description: "AWS API Gateway Version 2 Domain Name",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"domain_name"}),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"NotFoundException"}),
			},
			Hydrate: opengovernance.GetApiGatewayV2DomainName,
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListApiGatewayV2DomainName,
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "domain_name",
				Description: "The name of the DomainName resource",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DomainName.DomainName"),
			},
			{
				Name:        "api_mapping_selection_expression",
				Description: "The API mapping selection expression.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ApiMappingSelectionExpression"),
			},
			{
				Name:        "domain_name_configurations",
				Description: "The domain name configurations",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DomainName.DomainNameConfigurations")},
			{
				Name:        "mutual_tls_authentication",
				Description: "The mutual TLS authentication configuration for a custom domain name",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DomainName.MutualTlsAuthentication")},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DomainName.Tags")},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DomainName.DomainName")},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("ARN").Transform(arnToAkas)},
		}),
	}
}
