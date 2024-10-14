package aws

import (
	"context"
	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsAPIGatewayAuthorizer(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_api_gateway_authorizer",
		Description: "AWS API Gateway Authorizer",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"rest_api_id", "id"}),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"NotFoundException"}),
			},
			Hydrate: opengovernance.GetApiGatewayAuthorizer,
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListApiGatewayAuthorizer,
		},

		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "id",
				Description: "The identifier for the authorizer resource",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Authorizer.Id"),
			},
			{
				Name:        "name",
				Description: "The name of the authorizer",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Authorizer.Name")},
			{
				Name:        "rest_api_id",
				Description: "The id of the rest api",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.RestApiId")},
			{
				Name:        "auth_type",
				Description: "Optional customer-defined field, used in OpenAPI imports and exports without functional impact",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Authorizer.AuthType")},
			{
				Name:        "authorizer_credentials",
				Description: "Specifies the required credentials as an IAM role for API Gateway to invoke the authorizer",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Authorizer.AuthorizerCredentials")},
			{
				Name:        "authorizer_uri",
				Description: "Specifies the authorizer's Uniform Resource Identifier (URI). For TOKEN or REQUEST authorizers, this must be a well-formed Lambda function URI",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Authorizer.AuthorizerUri")},
			{
				Name:        "identity_validation_expression",
				Description: "A validation expression for the incoming identity token. For TOKEN authorizers, this value is a regular expression",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Authorizer.IdentityValidationExpression")},
			{
				Name:        "identity_source",
				Description: "The identity source for which authorization is requested",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Authorizer.IdentitySource")},
			{
				Name:        "provider_arns",
				Description: "A list of the Amazon Cognito user pool ARNs for the COGNITO_USER_POOLS authorizer",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Authorizer.ProviderARNs")},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Authorizer.Name")},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("ARN").Transform(arnToAkas)},
		}),
	}
}
