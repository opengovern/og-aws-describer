package describer

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/apigateway/types"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/apigateway"
	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2"
	typesv2 "github.com/aws/aws-sdk-go-v2/service/apigatewayv2/types"
	"github.com/kaytu-io/kaytu-aws-describer/aws/model"
)

func ApiGatewayStage(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := apigateway.NewFromConfig(cfg)
	paginator := apigateway.NewGetRestApisPaginator(client, &apigateway.GetRestApisInput{})

	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, restItem := range page.Items {
			out, err := client.GetStages(ctx, &apigateway.GetStagesInput{
				RestApiId: restItem.Id,
			})
			if err != nil {
				return nil, err
			}

			for _, stageItem := range out.Item {
				arn := "arn:" + describeCtx.Partition + ":apigateway:" + describeCtx.Region + "::/restapis/" + *restItem.Id + "/stages/" + *stageItem.StageName
				resource := Resource{
					Region: describeCtx.KaytuRegion,
					ARN:    arn,
					Name:   *restItem.Name,
					Description: model.ApiGatewayStageDescription{
						RestApiId: restItem.Id,
						Stage:     stageItem,
					},
				}
				if stream != nil {
					m := *stream
					err := m(resource)
					if err != nil {
						return nil, err
					}
				} else {
					values = append(values, resource)
				}
			}
		}
	}
	return values, nil
}

func GetApiGatewayStage(ctx context.Context, cfg aws.Config, fields map[string]string) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	restAPIID := fields["id"]
	client := apigateway.NewFromConfig(cfg)

	out, err := client.GetStages(ctx, &apigateway.GetStagesInput{
		RestApiId: &restAPIID,
	})
	if err != nil {
		return nil, err
	}

	var values []Resource
	for _, stageItem := range out.Item {
		arn := "arn:" + describeCtx.Partition + ":apigateway:" + describeCtx.Region + "::/restapis/" + restAPIID + "/stages/" + *stageItem.StageName
		values = append(values, Resource{
			Region: describeCtx.KaytuRegion,
			ARN:    arn,
			Name:   *stageItem.StageName,
			Description: model.ApiGatewayStageDescription{
				RestApiId: &restAPIID,
				Stage:     stageItem,
			},
		})
	}
	return values, nil
}

func ApiGatewayV2Stage(ctx context.Context, cfg aws.Config) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := apigatewayv2.NewFromConfig(cfg)

	var apis []typesv2.Api
	err := PaginateRetrieveAll(func(prevToken *string) (nextToken *string, err error) {
		output, err := client.GetApis(ctx, &apigatewayv2.GetApisInput{
			NextToken: prevToken,
		})
		if err != nil {
			return nil, err
		}

		apis = append(apis, output.Items...)
		return output.NextToken, nil
	})
	if err != nil {
		return nil, err
	}

	var values []Resource
	for _, api := range apis {
		var stages []typesv2.Stage
		err := PaginateRetrieveAll(func(prevToken *string) (nextToken *string, err error) {
			output, err := client.GetStages(ctx, &apigatewayv2.GetStagesInput{
				ApiId:     api.ApiId,
				NextToken: prevToken,
			})
			if err != nil {
				return nil, err
			}

			stages = append(stages, output.Items...)
			return output.NextToken, nil
		})
		if err != nil {
			return nil, err
		}

		for _, stage := range stages {
			values = append(values, Resource{
				Region: describeCtx.KaytuRegion,
				ID:     CompositeID(*api.ApiId, *stage.StageName),
				Name:   *api.Name,
				Description: model.ApiGatewayV2StageDescription{
					ApiId: api.ApiId,
					Stage: stage,
				},
			})
		}
	}

	return values, nil
}

func ApiGatewayRestAPI(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := apigateway.NewFromConfig(cfg)
	paginator := apigateway.NewGetRestApisPaginator(client, &apigateway.GetRestApisInput{})

	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			if isErr(err, "NotFoundException") {
				continue
			}
			return nil, err
		}

		for _, restItem := range page.Items {
			arn := fmt.Sprintf("arn:%s:apigateway:%s::/restapis/%s", describeCtx.Partition, describeCtx.Region, *restItem.Id)
			resource := Resource{
				Region: describeCtx.KaytuRegion,
				ARN:    arn,
				Name:   *restItem.Name,
				Description: model.ApiGatewayRestAPIDescription{
					RestAPI: restItem,
				},
			}
			if stream != nil {
				m := *stream
				err := m(resource)
				if err != nil {
					return nil, err
				}
			} else {
				values = append(values, resource)
			}
		}
	}
	return values, nil
}

func GetApiGatewayRestAPI(ctx context.Context, cfg aws.Config, fields map[string]string) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	id := fields["id"]
	client := apigateway.NewFromConfig(cfg)

	out, err := client.GetRestApi(ctx, &apigateway.GetRestApiInput{
		RestApiId: &id,
	})
	if err != nil {
		if isErr(err, "NotFoundException") {
			return nil, nil
		}
		return nil, err
	}

	var values []Resource
	arn := fmt.Sprintf("arn:%s:apigateway:%s::/restapis/%s", describeCtx.Partition, describeCtx.Region, *out.Id)
	values = append(values, Resource{
		Region: describeCtx.KaytuRegion,
		ARN:    arn,
		Name:   *out.Name,
		Description: model.ApiGatewayRestAPIDescription{
			RestAPI: types.RestApi{
				ApiKeySource:              out.ApiKeySource,
				BinaryMediaTypes:          out.BinaryMediaTypes,
				CreatedDate:               out.CreatedDate,
				Description:               out.Description,
				DisableExecuteApiEndpoint: out.DisableExecuteApiEndpoint,
				EndpointConfiguration:     out.EndpointConfiguration,
				Id:                        out.Id,
				MinimumCompressionSize:    out.MinimumCompressionSize,
				Name:                      out.Name,
				Policy:                    out.Policy,
				Tags:                      out.Tags,
				Version:                   out.Version,
				Warnings:                  out.Warnings,
			},
		},
	})
	return values, nil
}

func ApiGatewayApiKey(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := apigateway.NewFromConfig(cfg)
	paginator := apigateway.NewGetApiKeysPaginator(client, &apigateway.GetApiKeysInput{})

	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			if isErr(err, "NotFoundException") {
				continue
			}
			return nil, err
		}

		for _, apiKey := range page.Items {
			arn := fmt.Sprintf("arn:%s:apigateway:%s::/apikeys/%s", describeCtx.Partition, describeCtx.Region, *apiKey.Id)
			resource := Resource{
				Region: describeCtx.KaytuRegion,
				ID:     *apiKey.Id,
				ARN:    arn,
				Name:   *apiKey.Name,
				Description: model.ApiGatewayApiKeyDescription{
					ApiKey: apiKey,
				},
			}
			if stream != nil {
				m := *stream
				err := m(resource)
				if err != nil {
					return nil, err
				}
			} else {
				values = append(values, resource)
			}
		}
	}
	return values, nil
}

func ApiGatewayUsagePlan(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := apigateway.NewFromConfig(cfg)
	paginator := apigateway.NewGetUsagePlansPaginator(client, &apigateway.GetUsagePlansInput{})

	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			if isErr(err, "NotFoundException") {
				continue
			}
			return nil, err
		}

		for _, usagePlan := range page.Items {
			arn := fmt.Sprintf("arn:%s:apigateway:%s::/usageplans/%s", describeCtx.Partition, describeCtx.Region, *usagePlan.Id)
			resource := Resource{
				Region: describeCtx.KaytuRegion,
				ID:     *usagePlan.Id,
				ARN:    arn,
				Name:   *usagePlan.Name,
				Description: model.ApiGatewayUsagePlanDescription{
					UsagePlan: usagePlan,
				},
			}
			if stream != nil {
				m := *stream
				err := m(resource)
				if err != nil {
					return nil, err
				}
			} else {
				values = append(values, resource)
			}
		}
	}
	return values, nil
}

func ApiGatewayAuthorizer(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := apigateway.NewFromConfig(cfg)
	paginator := apigateway.NewGetRestApisPaginator(client, &apigateway.GetRestApisInput{})

	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			if isErr(err, "NotFoundException") {
				continue
			}
			return nil, err
		}

		for _, api := range page.Items {
			authorizers, err := client.GetAuthorizers(ctx, &apigateway.GetAuthorizersInput{
				RestApiId: api.Id,
			})
			if err != nil {
				return nil, err
			}
			for _, authorizer := range authorizers.Items {
				arn := fmt.Sprintf("arn:%s:apigateway:%s::/restapis/%s/authorizer/%s", describeCtx.Partition, describeCtx.Region, *api.Id, *authorizer.Id)
				resource := Resource{
					Region: describeCtx.KaytuRegion,
					ID:     *authorizer.Id,
					ARN:    arn,
					Name:   *api.Name,
					Description: model.ApiGatewayAuthorizerDescription{
						Authorizer: authorizer,
						RestApiId:  *api.Id,
					},
				}
				if stream != nil {
					m := *stream
					err := m(resource)
					if err != nil {
						return nil, err
					}
				} else {
					values = append(values, resource)
				}
			}
		}
	}
	return values, nil
}

func ApiGatewayV2API(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)

	client := apigatewayv2.NewFromConfig(cfg)

	var values []Resource
	err := PaginateRetrieveAll(func(prevToken *string) (nextToken *string, err error) {
		output, err := client.GetApis(ctx, &apigatewayv2.GetApisInput{
			NextToken: prevToken,
		})
		if err != nil {
			if isErr(err, "NotFoundException") {
				return nil, nil
			}
			return nil, err
		}

		for _, api := range output.Items {
			arn := fmt.Sprintf("arn:%s:apigateway:%s::/apis/%s", describeCtx.Partition, describeCtx.Region, *api.ApiId)
			resource := Resource{
				Region: describeCtx.KaytuRegion,
				ARN:    arn,
				Name:   *api.Name,
				Description: model.ApiGatewayV2APIDescription{
					API: api,
				},
			}
			if stream != nil {
				if err := (*stream)(resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, resource)
			}
		}
		return output.NextToken, nil
	})
	if err != nil {
		if isErr(err, "NotFoundException") {
			return nil, nil
		}
		return nil, err
	}

	return values, nil
}

func GetApiGatewayV2API(ctx context.Context, cfg aws.Config, fields map[string]string) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	apiID := fields["id"]

	client := apigatewayv2.NewFromConfig(cfg)

	var values []Resource
	out, err := client.GetApi(ctx, &apigatewayv2.GetApiInput{
		ApiId: &apiID,
	})
	if err != nil {
		if isErr(err, "NotFoundException") {
			return nil, nil
		}
		return nil, err
	}

	arn := fmt.Sprintf("arn:%s:apigateway:%s::/apis/%s", describeCtx.Partition, describeCtx.Region, apiID)
	values = append(values, Resource{
		Region: describeCtx.KaytuRegion,
		ARN:    arn,
		Name:   *out.Name,
		Description: model.ApiGatewayV2APIDescription{
			API: typesv2.Api{
				Name:                      out.Name,
				ProtocolType:              out.ProtocolType,
				RouteSelectionExpression:  out.RouteSelectionExpression,
				ApiEndpoint:               out.ApiEndpoint,
				ApiGatewayManaged:         out.ApiGatewayManaged,
				ApiId:                     out.ApiId,
				ApiKeySelectionExpression: out.ApiKeySelectionExpression,
				CorsConfiguration:         out.CorsConfiguration,
				CreatedDate:               out.CreatedDate,
				Description:               out.Description,
				DisableExecuteApiEndpoint: out.DisableExecuteApiEndpoint,
				DisableSchemaValidation:   out.DisableSchemaValidation,
				ImportInfo:                out.ImportInfo,
				Tags:                      out.Tags,
				Version:                   out.Version,
				Warnings:                  out.Warnings,
			},
		},
	})

	return values, nil
}

func ApiGatewayV2DomainName(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)

	client := apigatewayv2.NewFromConfig(cfg)
	var values []Resource
	err := PaginateRetrieveAll(func(prevToken *string) (nextToken *string, err error) {
		output, err := client.GetDomainNames(ctx, &apigatewayv2.GetDomainNamesInput{
			NextToken: prevToken,
		})
		if err != nil {
			if isErr(err, "NotFoundException") {
				return nil, nil
			}
			return nil, err
		}

		for _, domainName := range output.Items {
			arn := fmt.Sprintf("arn:%s:apigateway:%s::/domainnames/%s", describeCtx.Partition, describeCtx.Region, *domainName.DomainName)
			resource := Resource{
				Region: describeCtx.KaytuRegion,
				ARN:    arn,
				Name:   *domainName.DomainName,
				Description: model.ApiGatewayV2DomainNameDescription{
					DomainName: domainName,
				},
			}
			if stream != nil {
				if err := (*stream)(resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, resource)
			}
		}
		return output.NextToken, nil
	})
	if err != nil {
		if isErr(err, "NotFoundException") {
			return nil, nil
		}
		return nil, err
	}

	return values, nil
}

func GetApiGatewayV2DomainName(ctx context.Context, cfg aws.Config, fields map[string]string) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	domainName := fields["domain_name"]

	client := apigatewayv2.NewFromConfig(cfg)
	var values []Resource
	out, err := client.GetDomainName(ctx, &apigatewayv2.GetDomainNameInput{
		DomainName: &domainName,
	})
	if err != nil {
		if isErr(err, "NotFoundException") {
			return nil, nil
		}
		return nil, err
	}
	arn := fmt.Sprintf("arn:%s:apigateway:%s::/domainnames/%s", describeCtx.Partition, describeCtx.Region, domainName)
	values = append(values, Resource{
		Region: describeCtx.KaytuRegion,
		ARN:    arn,
		Name:   *out.DomainName,
		Description: model.ApiGatewayV2DomainNameDescription{
			DomainName: typesv2.DomainName{
				DomainName:                    out.DomainName,
				ApiMappingSelectionExpression: out.ApiMappingSelectionExpression,
				DomainNameConfigurations:      out.DomainNameConfigurations,
				MutualTlsAuthentication:       out.MutualTlsAuthentication,
				Tags:                          out.Tags,
			},
		},
	})
	return values, nil
}

func ApiGatewayV2Integration(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)

	client := apigatewayv2.NewFromConfig(cfg)

	var values []Resource
	err := PaginateRetrieveAll(func(prevToken *string) (*string, error) {
		output, err := client.GetApis(ctx, &apigatewayv2.GetApisInput{
			NextToken: prevToken,
		})
		if err != nil {
			if isErr(err, "NotFoundException") {
				return nil, nil
			}
			return nil, err
		}

		err = PaginateRetrieveAll(func(prevToken *string) (nextToken *string, err error) {
			for _, api := range output.Items {
				output, err := client.GetIntegrations(ctx, &apigatewayv2.GetIntegrationsInput{
					ApiId:     aws.String(*api.ApiId),
					NextToken: prevToken,
				})

				if err != nil {
					return nil, err
				}

				for _, integration := range output.Items {
					arn := fmt.Sprintf("arn:%s:apigateway:%s::/apis/%s/integrations/%s", describeCtx.Partition, describeCtx.Region, *api.ApiId, *integration.IntegrationId)
					resource := Resource{
						Region: describeCtx.KaytuRegion,
						ARN:    arn,
						ID:     *integration.IntegrationId,
						Description: model.ApiGatewayV2IntegrationDescription{
							Integration: integration,
							ApiId:       *api.ApiId,
						},
					}
					if stream != nil {
						if err := (*stream)(resource); err != nil {
							return nil, err
						}
					} else {
						values = append(values, resource)
					}
				}
				if err != nil {
					return nil, err
				}
				return output.NextToken, nil
			}
			return output.NextToken, nil
		})

		if err != nil {
			if isErr(err, "NotFoundException") || isErr(err, "TooManyRequestsException") {
				return nil, nil
			}
			return nil, err
		}
		return output.NextToken, nil
	})
	if err != nil {
		if isErr(err, "NotFoundException") {
			return nil, nil
		}
		return nil, err
	}

	return values, nil
}

func GetApiGatewayV2Integration(ctx context.Context, cfg aws.Config, fields map[string]string) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	apiId := fields["api_id"]
	integrationID := fields["id"]

	client := apigatewayv2.NewFromConfig(cfg)

	var values []Resource
	api, err := client.GetApi(ctx, &apigatewayv2.GetApiInput{
		ApiId: &apiId,
	})
	if err != nil {
		if isErr(err, "NotFoundException") {
			return nil, nil
		}
		return nil, err
	}

	integration, err := client.GetIntegration(ctx, &apigatewayv2.GetIntegrationInput{
		ApiId:         aws.String(*api.ApiId),
		IntegrationId: &integrationID,
	})

	if err != nil {
		return nil, err
	}

	arn := fmt.Sprintf("arn:%s:apigateway:%s::/apis/%s/integrations/%s", describeCtx.Partition, describeCtx.Region, *api.ApiId, *integration.IntegrationId)
	values = append(values, Resource{
		Region: describeCtx.KaytuRegion,
		ARN:    arn,
		ID:     *integration.IntegrationId,
		Description: model.ApiGatewayV2IntegrationDescription{
			Integration: typesv2.Integration{
				ApiGatewayManaged:                      integration.ApiGatewayManaged,
				ConnectionId:                           integration.ConnectionId,
				ConnectionType:                         integration.ConnectionType,
				ContentHandlingStrategy:                integration.ContentHandlingStrategy,
				CredentialsArn:                         integration.CredentialsArn,
				Description:                            integration.Description,
				IntegrationId:                          integration.IntegrationId,
				IntegrationMethod:                      integration.IntegrationMethod,
				IntegrationResponseSelectionExpression: integration.IntegrationResponseSelectionExpression,
				IntegrationSubtype:                     integration.IntegrationSubtype,
				IntegrationType:                        integration.IntegrationType,
				IntegrationUri:                         integration.IntegrationUri,
				PassthroughBehavior:                    integration.PassthroughBehavior,
				PayloadFormatVersion:                   integration.PayloadFormatVersion,
				RequestParameters:                      integration.RequestParameters,
				RequestTemplates:                       integration.RequestTemplates,
				ResponseParameters:                     integration.ResponseParameters,
				TemplateSelectionExpression:            integration.TemplateSelectionExpression,
				TimeoutInMillis:                        integration.TimeoutInMillis,
				TlsConfig:                              integration.TlsConfig,
			},
			ApiId: *api.ApiId,
		},
	})

	return values, nil
}
