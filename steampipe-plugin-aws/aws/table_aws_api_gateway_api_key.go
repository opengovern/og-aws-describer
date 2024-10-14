package aws

import (
	"context"

	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsAPIGatewayAPIKey(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_api_gateway_api_key",
		Description: "AWS API Gateway API Key",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"NotFoundException"}),
			},
			Hydrate: opengovernance.GetApiGatewayApiKey,
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListApiGatewayApiKey,
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:    "customer_id",
					Require: plugin.Optional,
				},
			},
		},

		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the API Key",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ApiKey.Name")},
			{
				Name:        "id",
				Description: "The identifier of the API Key",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ApiKey.Id")},
			{
				Name:        "enabled",
				Description: "Specifies whether the API Key can be used by callers",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.ApiKey.Enabled")},
			{
				Name:        "created_date",
				Description: "The timestamp when the API Key was created",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.ApiKey.CreatedDate")},
			{
				Name:        "last_updated_date",
				Description: "The timestamp when the API Key was last updated",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.ApiKey.LastUpdatedDate")},
			{
				Name:        "customer_id",
				Description: "An AWS Marketplace customer identifier , when integrating with the AWS SaaS Marketplace",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ApiKey.CustomerId")},
			{
				Name:        "description",
				Description: "The description of the API Key",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ApiKey.Description")},
			{
				Name:        "value",
				Description: "The value of the API Key",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ApiKey.Value")},
			{
				Name:        "stage_keys",
				Description: "A list of Stage resources that are associated with the ApiKey resource",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "tags_src",
				Description: "A list of tags attached to API key",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Tags"),
			},

			// Standard columns
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Name"),
			},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("ARN").Transform(arnToAkas)},
		}),
	}
}
