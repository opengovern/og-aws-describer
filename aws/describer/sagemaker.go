package describer

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/kaytu-io/kaytu-aws-describer/aws/model"
)

func SageMakerEndpointConfiguration(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := sagemaker.NewFromConfig(cfg)
	paginator := sagemaker.NewListEndpointConfigsPaginator(client, &sagemaker.ListEndpointConfigsInput{})

	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, item := range page.EndpointConfigs {
			out, err := client.DescribeEndpointConfig(ctx, &sagemaker.DescribeEndpointConfigInput{
				EndpointConfigName: item.EndpointConfigName,
			})
			if err != nil {
				return nil, err
			}

			tags, err := client.ListTags(ctx, &sagemaker.ListTagsInput{
				ResourceArn: item.EndpointConfigArn,
			})
			if err != nil {
				return nil, err
			}

			resource := Resource{
				Region: describeCtx.KaytuRegion,
				ARN:    *out.EndpointConfigArn,
				Name:   *out.EndpointConfigName,
				Description: model.SageMakerEndpointConfigurationDescription{
					EndpointConfig: out,
					Tags:           tags.Tags,
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

func SageMakerApp(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := sagemaker.NewFromConfig(cfg)

	var values []Resource
	paginator := sagemaker.NewListDomainsPaginator(client, &sagemaker.ListDomainsInput{})
	for paginator.HasMorePages() {
		output, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, domain := range output.Domains {
			appPaginator := sagemaker.NewListAppsPaginator(client, &sagemaker.ListAppsInput{
				DomainIdEquals: domain.DomainId,
			})

			for appPaginator.HasMorePages() {
				appOutput, err := appPaginator.NextPage(ctx)
				if err != nil {
					return nil, err
				}

				for _, items := range appOutput.Apps {
					params := sagemaker.DescribeAppInput{
						AppName:         items.AppName,
						AppType:         items.AppType,
						DomainId:        domain.DomainId,
						UserProfileName: items.UserProfileName,
					}

					data, err := client.DescribeApp(ctx, &params)
					if err != nil {
						return nil, err
					}

					resource := Resource{
						Region: describeCtx.KaytuRegion,
						ARN:    *data.AppArn,
						Name:   *data.AppName,
						Description: model.SageMakerAppDescription{
							AppDetails:        items,
							DescribeAppOutput: data,
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
	}
	return values, nil
}

func SageMakerDomain(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := sagemaker.NewFromConfig(cfg)

	var values []Resource
	paginator := sagemaker.NewListDomainsPaginator(client, &sagemaker.ListDomainsInput{})
	for paginator.HasMorePages() {
		output, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, domain := range output.Domains {
			data, err := client.DescribeDomain(ctx, &sagemaker.DescribeDomainInput{
				DomainId: domain.DomainId,
			})
			if err != nil {
				return nil, err
			}
			params := &sagemaker.ListTagsInput{
				ResourceArn: domain.DomainArn,
			}

			pagesLeft := true
			var tags []types.Tag
			for pagesLeft {
				keyTags, err := client.ListTags(ctx, params)
				if err != nil {
					return nil, err
				}
				tags = append(tags, keyTags.Tags...)

				if keyTags.NextToken != nil {
					params.NextToken = keyTags.NextToken
				} else {
					pagesLeft = false
				}
			}

			resource := Resource{
				Region: describeCtx.KaytuRegion,
				ARN:    *domain.DomainArn,
				Name:   *domain.DomainName,
				Description: model.SageMakerDomainDescription{
					Domain:     data,
					DomainItem: domain,
					Tags:       tags,
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

func SageMakerNotebookInstance(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := sagemaker.NewFromConfig(cfg)
	paginator := sagemaker.NewListNotebookInstancesPaginator(client, &sagemaker.ListNotebookInstancesInput{})

	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, item := range page.NotebookInstances {
			out, err := client.DescribeNotebookInstance(ctx, &sagemaker.DescribeNotebookInstanceInput{
				NotebookInstanceName: item.NotebookInstanceName,
			})
			if err != nil {
				return nil, err
			}

			tags, err := client.ListTags(ctx, &sagemaker.ListTagsInput{
				ResourceArn: out.NotebookInstanceArn,
			})
			if err != nil {
				return nil, err
			}

			resource := Resource{
				Region: describeCtx.KaytuRegion,
				ARN:    *out.NotebookInstanceArn,
				Name:   *out.NotebookInstanceName,
				Description: model.SageMakerNotebookInstanceDescription{
					NotebookInstance: out,
					Tags:             tags.Tags,
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
