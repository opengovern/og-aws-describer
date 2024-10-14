package aws

import (
	"context"

	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsAPIGatewayV2Integration(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_api_gatewayv2_integration",
		Description: "AWS API Gateway Version 2 Integration",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"integration_id", "api_id"}),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"NotFoundException", "TooManyRequestsException"}),
			},
			Hydrate: opengovernance.GetApiGatewayV2Integration,
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListApiGatewayV2Integration,
		},

		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "integration_id",
				Description: "Represents the identifier of an integration.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Integration.IntegrationId")},
			{
				Name:        "api_id",
				Description: "Represents the identifier of an API.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ApiId")},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) specifying the integration.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("ARN").Transform(arnToAkas),
			},
			{
				Name:        "description",
				Description: "Represents the description of an integration.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Integration.Description")},
			{
				Name:        "integration_method",
				Description: "Specifies the integration's HTTP method type.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Integration.IntegrationMethod")},
			{
				Name:        "integration_type",
				Description: "Represents an API method integration type.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Integration.IntegrationType")},
			{
				Name:        "integration_uri",
				Description: "A string representation of a URI with a length between [1-2048]. For a Lambda integration, specify the URI of a Lambda function. For an HTTP integration, specify a fully-qualified URL.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Integration.IntegrationUri")},
			{
				Name:        "api_gateway_managed",
				Description: "Specifies whether an integration is managed by API Gateway. If you created an API using using quick create, the resulting integration is managed by API Gateway. You can update a managed integration, but you can't delete it.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Integration.ApiGatewayManaged")},
			{
				Name:        "connection_id",
				Description: "The ID of the VPC link for a private integration. Supported only for HTTP APIs.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Integration.ConnectionId")},
			{
				Name:        "connection_type",
				Description: "Represents a connection type.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Integration.ConnectionType")},
			{
				Name:        "content_handling_strategy",
				Description: "Specifies how to handle response payload content type conversions. Supported only for WebSocket APIs.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Integration.ContentHandlingStrategy")},
			{
				Name:        "credentials_arn",
				Description: "Specifies the credentials required for the integration, if any. For AWS integrations, three options are available. To specify an IAM Role for API Gateway to assume, use the role's Amazon Resource Name (ARN).",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Integration.CredentialsArn")},
			{
				Name:        "integration_response_selection_expression",
				Description: "An expression used to extract information at runtime. See Selection Expressions(https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-apikey-selection-expressions for more information.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Integration.IntegrationResponseSelectionExpression")},
			{
				Name:        "integration_subtype",
				Description: "A string with a length between [1-128].",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Integration.IntegrationSubtype")},
			{
				Name:        "passthrough_behavior",
				Description: "Represents passthrough behavior for an integration response. Supported only for WebSocket APIs.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Integration.PassthroughBehavior")},
			{
				Name:        "payload_format_version",
				Description: "Specifies the format of the payload sent to an integration. Required for HTTP APIs.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Integration.PayloadFormatVersion")},
			{
				Name:        "template_selection_expression",
				Description: "The template selection expression for the integration. Supported only for WebSocket APIs.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Integration.TemplateSelectionExpression")},
			{
				Name:        "timeout_in_millis",
				Description: "Indicates custom timeout between 50 and 29,000 milliseconds for WebSocket APIs and between 50 and 30,000 milliseconds for HTTP APIs. The default timeout is 29 seconds for WebSocket APIs and 30 seconds for HTTP APIs.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Integration.TimeoutInMillis")},
			{
				Name:        "request_parameters",
				Description: "For HTTP API itegrations, without a specified integrationSubtype request parameters are a key-value map specifying how to transform HTTP requests before sending them to backend integrations. The key should follow the pattern <action>:<header|querystring|path>.<location>. The action can be append, overwrite or remove. For values, you can provide static values, or map request data, stage variables, or context variables that are evaluated at runtime. To learn more, see Transforming API requests and responses (https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-parameter-mapping.html).",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Integration.RequestParameters")},
			{
				Name:        "request_templates",
				Description: "Represents a map of Velocity templates that are applied on the request payload based on the value of the Content-Type header sent by the client. The content type value is the key in this map, and the template (as a String) is the value. Supported only for WebSocket APIs.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Integration.RequestTemplates")},
			{
				Name:        "response_parameters",
				Description: "API requests and responses (https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-parameter-mapping.html).",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Integration.ResponseParameters")},
			{
				Name:        "tls_config",
				Description: "The TLS configuration for a private integration. If you specify a TLS configuration, private integration traffic uses the HTTPS protocol. Supported only for HTTP APIs.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Integration.TlsConfig")},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Integration.IntegrationId")},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("ARN").Transform(transform.EnsureStringArray),
			},
		}),
	}
}
