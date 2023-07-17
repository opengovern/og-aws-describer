package describer

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/aws/aws-sdk-go-v2/service/lambda/types"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/smithy-go"
	"github.com/kaytu-io/kaytu-aws-describer/aws/model"
)

func LambdaFunction(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := lambda.NewFromConfig(cfg)
	paginator := lambda.NewListFunctionsPaginator(client, &lambda.ListFunctionsInput{
		FunctionVersion: types.FunctionVersionAll,
	})

	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, v := range page.Functions {
			policy, err := client.GetPolicy(ctx, &lambda.GetPolicyInput{
				FunctionName: v.FunctionName,
			})
			if err != nil {
				var ae smithy.APIError
				if errors.As(err, &ae) && ae.ErrorCode() == "ResourceNotFoundException" {
					policy = &lambda.GetPolicyOutput{}
					err = nil
				}

				if awsErr, ok := err.(awserr.Error); ok {
					log.Println("Describe Lambda Error:", awsErr.Code(), awsErr.Message())
					if awsErr.Code() == "ResourceNotFoundException" {
						policy = &lambda.GetPolicyOutput{}
						err = nil
					}
				}

				if err != nil {
					return nil, err
				}
			}

			function, err := client.GetFunction(ctx, &lambda.GetFunctionInput{
				FunctionName: v.FunctionName,
			})
			if err != nil {
				return nil, err
			}

			resource := Resource{
				Region: describeCtx.KaytuRegion,
				ARN:    *v.FunctionArn,
				Name:   *v.FunctionName,
				Description: model.LambdaFunctionDescription{
					Function: function,
					Policy:   policy,
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
	}

	return values, nil
}

func GetLambdaFunction(ctx context.Context, cfg aws.Config, fields map[string]string) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	functionName := fields["name"]
	client := lambda.NewFromConfig(cfg)
	out, err := client.GetFunction(ctx, &lambda.GetFunctionInput{
		FunctionName: &functionName,
		Qualifier:    nil,
	})
	if err != nil {
		return nil, err
	}
	v := out.Configuration

	var values []Resource
	policy, err := client.GetPolicy(ctx, &lambda.GetPolicyInput{
		FunctionName: v.FunctionName,
	})
	if err != nil {
		var ae smithy.APIError
		if errors.As(err, &ae) && ae.ErrorCode() == "ResourceNotFoundException" {
			policy = &lambda.GetPolicyOutput{}
			err = nil
		}

		if awsErr, ok := err.(awserr.Error); ok {
			log.Println("Describe Lambda Error:", awsErr.Code(), awsErr.Message())
			if awsErr.Code() == "ResourceNotFoundException" {
				policy = &lambda.GetPolicyOutput{}
				err = nil
			}
		}

		if err != nil {
			return nil, err
		}
	}

	function, err := client.GetFunction(ctx, &lambda.GetFunctionInput{
		FunctionName: v.FunctionName,
	})
	if err != nil {
		return nil, err
	}

	values = append(values, Resource{
		Region: describeCtx.KaytuRegion,
		ARN:    *v.FunctionArn,
		Name:   *v.FunctionName,
		Description: model.LambdaFunctionDescription{
			Function: function,
			Policy:   policy,
		},
	})

	return values, nil
}

func LambdaFunctionVersion(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := lambda.NewFromConfig(cfg)
	paginator := lambda.NewListFunctionsPaginator(client, &lambda.ListFunctionsInput{
		FunctionVersion: types.FunctionVersionAll,
	})

	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, v := range page.Functions {
			id := fmt.Sprintf("%s:%s", *v.FunctionArn, *v.Version)

			policy, err := client.GetPolicy(ctx, &lambda.GetPolicyInput{
				FunctionName: v.FunctionName,
				Qualifier:    v.Version,
			})
			if err != nil {
				if isErr(err, "ResourceNotFoundException") {
					policy = &lambda.GetPolicyOutput{}
				} else {
					return nil, err
				}
			}

			resource := Resource{
				Region: describeCtx.KaytuRegion,
				ID:     id,
				Description: model.LambdaFunctionVersionDescription{
					FunctionVersion: v,
					Policy:          policy,
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
	}
	return values, nil
}

func LambdaAlias(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	fns, err := LambdaFunction(ctx, cfg, nil)
	if err != nil {
		return nil, err
	}

	client := lambda.NewFromConfig(cfg)

	var values []Resource
	for _, f := range fns {
		fn := f.Description.(model.LambdaFunctionDescription).Function.Configuration
		paginator := lambda.NewListAliasesPaginator(client, &lambda.ListAliasesInput{
			FunctionName:    fn.FunctionName,
			FunctionVersion: fn.Version,
		})

		for paginator.HasMorePages() {
			page, err := paginator.NextPage(ctx)
			if err != nil {
				if isErr(err, "ResourceNotFoundException") {
					continue
				}
				return nil, err
			}

			for _, v := range page.Aliases {
				policy, err := client.GetPolicy(ctx, &lambda.GetPolicyInput{
					FunctionName: fn.FunctionName,
					Qualifier:    v.Name,
				})
				if err != nil {
					if isErr(err, "ResourceNotFoundException") {
						policy = &lambda.GetPolicyOutput{}
					} else {
						return nil, err
					}
				}

				urlConfig, err := client.GetFunctionUrlConfig(ctx, &lambda.GetFunctionUrlConfigInput{
					FunctionName: fn.FunctionName,
					Qualifier:    v.Name,
				})
				if err != nil {
					if isErr(err, "ResourceNotFoundException") {
						urlConfig = &lambda.GetFunctionUrlConfigOutput{}
					} else {
						return nil, err
					}
				}

				resource := Resource{
					Region: describeCtx.Region,
					ARN:    *v.AliasArn,
					Name:   *v.Name,
					Description: model.LambdaAliasDescription{
						FunctionName: *fn.FunctionName,
						Alias:        v,
						Policy:       policy,
						UrlConfig:    *urlConfig,
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
		}
	}

	return values, nil
}

func LambdaPermission(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	fns, err := LambdaFunction(ctx, cfg, nil)
	if err != nil {
		return nil, err
	}

	client := lambda.NewFromConfig(cfg)

	var values []Resource
	for _, f := range fns {
		fn := f.Description.(model.LambdaFunctionDescription).Function.Configuration
		v, err := client.GetPolicy(ctx, &lambda.GetPolicyInput{
			FunctionName: fn.FunctionArn,
		})
		if err != nil {
			var ae smithy.APIError
			if errors.As(err, &ae) && ae.ErrorCode() == "ResourceNotFoundException" {
				continue
			}

			return nil, err
		}

		resource := Resource{
			Region:      describeCtx.Region,
			ID:          CompositeID(*fn.FunctionArn, *v.Policy),
			Name:        *v.Policy,
			Description: v,
		}
		if stream != nil {
			if err := (*stream)(resource); err != nil {
				return nil, err
			}
		} else {
			values = append(values, resource)
		}
	}

	return values, nil
}

func LambdaEventInvokeConfig(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	fns, err := LambdaFunction(ctx, cfg, nil)
	if err != nil {
		return nil, err
	}

	client := lambda.NewFromConfig(cfg)

	var values []Resource
	for _, f := range fns {
		fn := f.Description.(model.LambdaFunctionDescription).Function.Configuration
		paginator := lambda.NewListFunctionEventInvokeConfigsPaginator(client, &lambda.ListFunctionEventInvokeConfigsInput{
			FunctionName: fn.FunctionName,
		})

		for paginator.HasMorePages() {
			page, err := paginator.NextPage(ctx)
			if err != nil {
				return nil, err
			}

			for _, v := range page.FunctionEventInvokeConfigs {
				resource := Resource{
					Region:      describeCtx.Region,
					ID:          *fn.FunctionName, // Invoke Config is unique per function
					Name:        *fn.FunctionName,
					Description: v,
				}
				if stream != nil {
					if err := (*stream)(resource); err != nil {
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

func LambdaCodeSigningConfig(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := lambda.NewFromConfig(cfg)
	paginator := lambda.NewListCodeSigningConfigsPaginator(client, &lambda.ListCodeSigningConfigsInput{})

	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, v := range page.CodeSigningConfigs {
			resource := Resource{
				Region:      describeCtx.Region,
				ARN:         *v.CodeSigningConfigArn,
				Name:        *v.CodeSigningConfigArn,
				Description: v,
			}
			if stream != nil {
				if err := (*stream)(resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, resource)
			}
		}
	}

	return values, nil
}

func LambdaEventSourceMapping(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := lambda.NewFromConfig(cfg)
	paginator := lambda.NewListEventSourceMappingsPaginator(client, &lambda.ListEventSourceMappingsInput{})

	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, v := range page.EventSourceMappings {
			resource := Resource{
				Region:      describeCtx.Region,
				ARN:         *v.EventSourceArn,
				Name:        *v.UUID,
				Description: v,
			}
			if stream != nil {
				if err := (*stream)(resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, resource)
			}
		}
	}

	return values, nil
}

func LambdaLayerVersion(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	layers, err := listLayers(ctx, cfg)
	if err != nil {
		return nil, err
	}

	client := lambda.NewFromConfig(cfg)

	var values []Resource
	for _, layer := range layers {
		paginator := lambda.NewListLayerVersionsPaginator(client, &lambda.ListLayerVersionsInput{
			LayerName: layer.LayerArn,
		})

		for paginator.HasMorePages() {
			page, err := paginator.NextPage(ctx)
			if err != nil {
				return nil, err
			}

			for _, v := range page.LayerVersions {
				layerVersion, err := client.GetLayerVersion(ctx, &lambda.GetLayerVersionInput{
					LayerName:     layer.LayerArn,
					VersionNumber: v.Version,
				})
				if err != nil {
					return nil, err
				}

				policy, err := client.GetLayerVersionPolicy(ctx, &lambda.GetLayerVersionPolicyInput{
					LayerName:     layer.LayerArn,
					VersionNumber: v.Version,
				})
				if err != nil {
					if isErr(err, "ResourceNotFoundException") {
						policy = &lambda.GetLayerVersionPolicyOutput{}
					} else {
						return nil, err
					}
				}

				resource := Resource{
					Region: describeCtx.Region,
					ARN:    *v.LayerVersionArn,
					Name:   *v.LayerVersionArn,
					Description: model.LambdaLayerVersionDescription{
						LayerName:    *layer.LayerName,
						LayerVersion: *layerVersion,
						Policy:       *policy,
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
		}
	}

	return values, nil
}

func LambdaLayer(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	layers, err := listLayers(ctx, cfg)
	if err != nil {
		return nil, err
	}

	var values []Resource
	for _, layer := range layers {
		resource := Resource{
			Region: describeCtx.Region,
			ARN:    *layer.LayerArn,
			Name:   *layer.LayerName,
			Description: model.LambdaLayerDescription{
				Layer: layer,
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

	return values, nil
}

func LambdaLayerVersionPermission(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	lvs, err := LambdaLayerVersion(ctx, cfg, nil)
	if err != nil {
		return nil, err
	}

	client := lambda.NewFromConfig(cfg)

	var values []Resource
	for _, lv := range lvs {
		arn := lv.Description.(model.LambdaLayerVersionDescription).LayerVersion.LayerVersionArn
		version := lv.Description.(model.LambdaLayerVersionDescription).LayerVersion.Version
		v, err := client.GetLayerVersionPolicy(ctx, &lambda.GetLayerVersionPolicyInput{
			LayerName:     arn,
			VersionNumber: version,
		})
		if err != nil {
			return nil, err
		}

		resource := Resource{
			Region:      describeCtx.Region,
			ID:          CompositeID(*arn, fmt.Sprintf("%d", version)),
			Name:        *arn,
			Description: v,
		}
		if stream != nil {
			if err := (*stream)(resource); err != nil {
				return nil, err
			}
		} else {
			values = append(values, resource)
		}
	}

	return values, nil
}

func listLayers(ctx context.Context, cfg aws.Config) ([]types.LayersListItem, error) {
	client := lambda.NewFromConfig(cfg)
	paginator := lambda.NewListLayersPaginator(client, &lambda.ListLayersInput{})

	var values []types.LayersListItem
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		values = append(values, page.Layers...)
	}

	return values, nil
}
