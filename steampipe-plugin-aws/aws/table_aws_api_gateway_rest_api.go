package aws

import (
	"context"
	"encoding/json"
	go_kit_packs "github.com/turbot/go-kit/types"
	"net/url"

	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsAPIGatewayRestAPI(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_api_gateway_rest_api",
		Description: "AWS API Gateway Rest API ",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("api_id"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"NotFoundException"}),
			},
			Hydrate: kaytu.GetApiGatewayRestAPI,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListApiGatewayRestAPI,
		},

		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The API's name",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.RestAPI.Name")},
			{
				Name:        "api_id",
				Description: "The API's identifier. This identifier is unique across all of APIs in API Gateway",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.RestAPI.Id")},
			{
				Name:        "version",
				Description: "A version identifier for the API",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.RestAPI.Version")},
			{
				Name:        "api_key_source",
				Description: "The source of the API key for metering requests according to a usage plan",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.RestAPI.ApiKeySource")},
			{
				Name:        "created_date",
				Description: "The timestamp when the API was created",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.RestAPI.CreatedDate")},
			{
				Name:        "description",
				Description: "The API's description",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.RestAPI.Description")},
			{
				Name:        "minimum_compression_size",
				Description: "A nullable integer that is used to enable compression (with non-negative between 0 and 10485760 (10M) bytes, inclusive) or disable compression (with a null value) on an API. When compression is enabled, compression or decompression is not applied on the payload if the payload size is smaller than this value",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.RestAPI.MinimumCompressionSize")},
			{
				Name:        "policy",
				Description: "A stringified JSON policy document that applies to this RestApi regardless of the caller and Method configuration",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.RestAPI.Policy").Transform(unmarshalJSON).Transform(transform.UnmarshalYAML),
			},
			{
				Name:        "policy_std",
				Description: "Contains the policy in a canonical form for easier searching.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.RestAPI.Policy").Transform(unmarshalJSON).Transform(policyToCanonical),
			},
			{
				Name:        "binary_media_types",
				Description: "The list of binary media types supported by the RestApi. By default, the RestApi supports only UTF-8-encoded text payloads",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.RestAPI.BinaryMediaTypes")},
			{
				Name:        "endpoint_configuration_types",
				Description: "The endpoint configuration of this RestApi showing the endpoint types of the API",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.RestAPI.EndpointConfiguration.Types")},
			{
				Name:        "endpoint_configuration_vpc_endpoint_ids",
				Description: "The endpoint configuration of this RestApi showing the endpoint types of the API",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.RestAPI.EndpointConfiguration.VpcEndpointIds")},
			{
				Name:        "warnings",
				Description: "The warning messages reported when failonwarnings is turned on during API import",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.RestAPI.Warnings")},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.RestAPI.Tags")},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.RestAPI.Name")},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("ARN").Transform(arnToAkas)},
		}),
	}
}

//// TRANSFORM FUNCTION

// unmarshalJSON :: parse the yaml-encoded data and return the result
func unmarshalJSON(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	inputStr := go_kit_packs.SafeString(d.Value)
	var result interface{}
	if inputStr != "" {
		// Resource IAM policy for aws_api_gateway_rest_api is stored as stringified json object after removing the double quotes from end
		decoded, err := url.QueryUnescape("\"" + inputStr + "\"")
		if err != nil {
			plugin.Logger(ctx).Error("aws_api_gateway_rest_api.unmarshalJSON", err)
			return nil, err
		}

		err = json.Unmarshal([]byte(decoded), &result)
		if err != nil {
			plugin.Logger(ctx).Error("aws_api_gateway_rest_api.unmarshalJSON", err)
			return nil, err
		}
	}

	return result, nil
}
