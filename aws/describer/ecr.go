package describer

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ecr"
	"github.com/aws/aws-sdk-go-v2/service/ecr/types"
	"github.com/aws/aws-sdk-go-v2/service/ecrpublic"
	public_types "github.com/aws/aws-sdk-go-v2/service/ecrpublic/types"
	"github.com/aws/smithy-go"
	"github.com/kaytu-io/kaytu-aws-describer/aws/model"
)

func ECRPublicRepository(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	// Only supported in US-EAST-1
	if !strings.EqualFold(cfg.Region, "us-east-1") {
		return []Resource{}, nil
	}

	client := ecrpublic.NewFromConfig(cfg)
	paginator := ecrpublic.NewDescribeRepositoriesPaginator(client, &ecrpublic.DescribeRepositoriesInput{})
	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			if isErr(err, "RepositoryNotFoundException") || isErr(err, "RepositoryPolicyNotFoundException") || isErr(err, "LifecyclePolicyNotFoundException") {
				continue
			}
			return nil, err
		}

		for _, v := range page.Repositories {
			var imageDetails []public_types.ImageDetail
			imagePaginator := ecrpublic.NewDescribeImagesPaginator(client, &ecrpublic.DescribeImagesInput{
				RepositoryName: v.RepositoryName,
			})
			for imagePaginator.HasMorePages() {
				imagePage, err := imagePaginator.NextPage(ctx)
				if err != nil {
					if isErr(err, "RepositoryNotFoundException") || isErr(err, "RepositoryPolicyNotFoundException") || isErr(err, "LifecyclePolicyNotFoundException") {
						continue
					}
					return nil, err
				}
				imageDetails = append(imageDetails, imagePage.ImageDetails...)
			}

			policyOutput, err := client.GetRepositoryPolicy(ctx, &ecrpublic.GetRepositoryPolicyInput{
				RepositoryName: v.RepositoryName,
			})
			if err != nil {
				if !isErr(err, "RepositoryNotFoundException") && !isErr(err, "RepositoryPolicyNotFoundException") && !isErr(err, "LifecyclePolicyNotFoundException") {
					return nil, err
				}
			}

			tagsOutput, err := client.ListTagsForResource(ctx, &ecrpublic.ListTagsForResourceInput{
				ResourceArn: v.RepositoryArn,
			})
			if err != nil {
				if !isErr(err, "RepositoryNotFoundException") && !isErr(err, "RepositoryPolicyNotFoundException") && !isErr(err, "LifecyclePolicyNotFoundException") {
					return nil, err
				} else {
					tagsOutput = &ecrpublic.ListTagsForResourceOutput{}
				}
			}

			resource := Resource{
				Region: describeCtx.KaytuRegion,
				ARN:    *v.RepositoryArn,
				Name:   *v.RepositoryName,
				Description: model.ECRPublicRepositoryDescription{
					PublicRepository: v,
					ImageDetails:     imageDetails,
					Policy:           policyOutput,
					Tags:             tagsOutput.Tags,
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

func ECRPublicRegistry(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	// Only supported in US-EAST-1
	if !strings.EqualFold(cfg.Region, "us-east-1") {
		return []Resource{}, nil
	}

	client := ecrpublic.NewFromConfig(cfg)
	paginator := ecrpublic.NewDescribeRegistriesPaginator(client, &ecrpublic.DescribeRegistriesInput{})
	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, v := range page.Registries {
			var tags []public_types.Tag
			tagsOutput, err := client.ListTagsForResource(ctx, &ecrpublic.ListTagsForResourceInput{
				ResourceArn: v.RegistryArn,
			})
			if err == nil {
				tags = tagsOutput.Tags
			}

			resource := Resource{
				Region: describeCtx.KaytuRegion,
				ARN:    *v.RegistryArn,
				Name:   *v.RegistryId,
				Description: model.ECRPublicRegistryDescription{
					PublicRegistry: v,
					Tags:           tags,
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

func ECRRepository(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := ecr.NewFromConfig(cfg)
	paginator := ecr.NewDescribeRepositoriesPaginator(client, &ecr.DescribeRepositoriesInput{})

	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			if isErr(err, "RepositoryNotFoundException") || isErr(err, "RepositoryPolicyNotFoundException") || isErr(err, "LifecyclePolicyNotFoundException") {
				continue
			}
			return nil, err
		}

		for _, v := range page.Repositories {
			lifeCyclePolicyOutput, err := client.GetLifecyclePolicy(ctx, &ecr.GetLifecyclePolicyInput{
				RepositoryName: v.RepositoryName,
			})
			if err != nil {
				if !isErr(err, "RepositoryNotFoundException") && !isErr(err, "RepositoryPolicyNotFoundException") && !isErr(err, "LifecyclePolicyNotFoundException") {
					return nil, err
				}
			}

			var imageDetails []types.ImageDetail
			imagePaginator := ecr.NewDescribeImagesPaginator(client, &ecr.DescribeImagesInput{
				RepositoryName: v.RepositoryName,
			})
			for imagePaginator.HasMorePages() {
				imagePage, err := imagePaginator.NextPage(ctx)
				if err != nil {
					if isErr(err, "RepositoryNotFoundException") || isErr(err, "RepositoryPolicyNotFoundException") || isErr(err, "LifecyclePolicyNotFoundException") {
						continue
					}
					return nil, err
				}
				imageDetails = append(imageDetails, imagePage.ImageDetails...)
			}

			policyOutput, err := client.GetRepositoryPolicy(ctx, &ecr.GetRepositoryPolicyInput{
				RepositoryName: v.RepositoryName,
			})
			if err != nil {
				if !isErr(err, "RepositoryNotFoundException") && !isErr(err, "RepositoryPolicyNotFoundException") && !isErr(err, "LifecyclePolicyNotFoundException") {
					return nil, err
				}
			}

			tagsOutput, err := client.ListTagsForResource(ctx, &ecr.ListTagsForResourceInput{
				ResourceArn: v.RepositoryArn,
			})
			if err != nil {
				if !isErr(err, "RepositoryNotFoundException") && !isErr(err, "RepositoryPolicyNotFoundException") && !isErr(err, "LifecyclePolicyNotFoundException") {
					return nil, err
				} else {
					tagsOutput = &ecr.ListTagsForResourceOutput{}
				}
			}

			resource := Resource{
				Region: describeCtx.KaytuRegion,
				ARN:    *v.RepositoryArn,
				Name:   *v.RepositoryName,
				Description: model.ECRRepositoryDescription{
					Repository:      v,
					LifecyclePolicy: lifeCyclePolicyOutput,
					ImageDetails:    imageDetails,
					Policy:          policyOutput,
					Tags:            tagsOutput.Tags,
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

func ECRRegistryPolicy(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := ecr.NewFromConfig(cfg)
	output, err := client.GetRegistryPolicy(ctx, &ecr.GetRegistryPolicyInput{})
	if err != nil {
		var ae smithy.APIError
		e := types.RegistryPolicyNotFoundException{}
		if errors.As(err, &ae) && ae.ErrorCode() == e.ErrorCode() {
			return []Resource{}, nil
		}
		return nil, err
	}

	var values []Resource
	resource := Resource{
		Region:      describeCtx.Region,
		ID:          *output.RegistryId,
		Name:        *output.RegistryId,
		Description: output,
	}
	if stream != nil {
		if err := (*stream)(resource); err != nil {
			return nil, err
		}
	} else {
		values = append(values, resource)
	}

	return values, nil
}

func ECRRegistry(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := ecr.NewFromConfig(cfg)
	output, err := client.DescribeRegistry(ctx, &ecr.DescribeRegistryInput{})
	if err != nil {
		return nil, err
	}

	var values []Resource
	resource := Resource{
		Region:      describeCtx.Region,
		ID:          *output.RegistryId,
		Name:        *output.RegistryId,
		Description: output,
	}
	if stream != nil {
		if err := (*stream)(resource); err != nil {
			return nil, err
		}
	} else {
		values = append(values, resource)
	}

	return values, nil
}

func ECRImage(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := ecr.NewFromConfig(cfg)
	repositoryPaginator := ecr.NewDescribeRepositoriesPaginator(client, &ecr.DescribeRepositoriesInput{})

	var values []Resource
	for repositoryPaginator.HasMorePages() {
		repositoryPage, err := repositoryPaginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, repository := range repositoryPage.Repositories {
			imagesPaginator := ecr.NewDescribeImagesPaginator(client, &ecr.DescribeImagesInput{
				RepositoryName: repository.RepositoryName,
				RegistryId:     repository.RegistryId,
			})
			if err != nil {
				return nil, err
			}

			for imagesPaginator.HasMorePages() {
				page, err := imagesPaginator.NextPage(ctx)
				if err != nil {
					return nil, err
				}
				for _, image := range page.ImageDetails {
					desc := model.ECRImageDescription{
						Image: image,
					}

					if len(image.ImageTags) > 0 {
						findingsPaginator := ecr.NewDescribeImageScanFindingsPaginator(client, &ecr.DescribeImageScanFindingsInput{
							RepositoryName: repository.RepositoryName,
							ImageId: &types.ImageIdentifier{
								ImageDigest: image.ImageDigest,
							},
						})

						// List call
						for findingsPaginator.HasMorePages() {
							output, err := findingsPaginator.NextPage(ctx)
							if err != nil {
								return nil, err
							}

							for _, finding := range output.ImageScanFindings.Findings {
								desc.ImageDigest = output.ImageId.ImageDigest
								desc.ImageScanFinding = finding
								desc.ImageScanStatus = *output.ImageScanStatus
								desc.ImageTag = output.ImageId.ImageTag
								if output.ImageScanFindings.ImageScanCompletedAt != nil {
									desc.ImageScanCompletedAt = output.ImageScanFindings.ImageScanCompletedAt
								}
								if output.ImageScanFindings.VulnerabilitySourceUpdatedAt != nil {
									desc.VulnerabilitySourceUpdatedAt = output.ImageScanFindings.VulnerabilitySourceUpdatedAt
								}
							}
						}
					}

					resource := Resource{
						Region:      describeCtx.KaytuRegion,
						Name:        fmt.Sprintf("%s:%s", *repository.RepositoryName, *image.ImageDigest),
						Description: desc,
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
