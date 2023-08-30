package aws

import (
	"context"
	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsAPIGatewayV2Api(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_api_gatewayv2_api",
		Description: "AWS API Gateway Version 2 API",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("api_id"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"NotFoundException"}),
			},
			Hydrate: kaytu.GetApiGatewayV2API,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListApiGatewayV2API,
		},

		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the API",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.API.Name")},
			{
				Name:        "api_id",
				Description: "The API ID",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.API.ApiId")},
			{
				Name:        "api_endpoint",
				Description: "The URI of the API, of the form {api-id}.execute-api.{region}.amazonaws.com",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.API.ApiEndpoint")},
			{
				Name:        "protocol_type",
				Description: "The API protocol",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.API.ProtocolType")},
			{
				Name:        "api_key_selection_expression",
				Description: "An API key selection expression. Supported only for WebSocket APIs",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.API.ApiKeySelectionExpression")},
			{
				Name:        "disable_execute_api_endpoint",
				Description: "Specifies whether clients can invoke your API by using the default execute-api endpoint.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.API.DisableExecuteApiEndpoint"),
			},
			{
				Name:        "route_selection_expression",
				Description: "The route selection expression for the API. For HTTP APIs, the routeSelectionExpression must be ${request.method} ${request.path}. If not provided, this will be the default for HTTP APIs",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.API.RouteSelectionExpression")},
			{
				Name:        "created_date",
				Description: "The timestamp when the API was created",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.API.CreatedDate")},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.API.Tags")},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.API.Name")},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("ARN").Transform(arnToAkas)},
		}),
	}
}
