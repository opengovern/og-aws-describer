package describer

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/codeartifact"
	"github.com/kaytu-io/kaytu-aws-describer/aws/model"
)

func CodeArtifactRepository(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	client := codeartifact.NewFromConfig(cfg)
	paginator := codeartifact.NewListRepositoriesPaginator(client, &codeartifact.ListRepositoriesInput{})

	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, v := range page.Repositories {
			tags, err := client.ListTagsForResource(ctx, &codeartifact.ListTagsForResourceInput{
				ResourceArn: v.Arn,
			})
			if err != nil {
				return nil, err
			}

			resource := Resource{
				ARN:  *v.Arn,
				Name: *v.Name,
				Description: model.CodeArtifactRepositoryDescription{
					Repository: v,
					Tags:       tags.Tags,
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

func CodeArtifactDomain(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	client := codeartifact.NewFromConfig(cfg)
	paginator := codeartifact.NewListDomainsPaginator(client, &codeartifact.ListDomainsInput{})

	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, v := range page.Domains {
			tags, err := client.ListTagsForResource(ctx, &codeartifact.ListTagsForResourceInput{
				ResourceArn: v.Arn,
			})
			if err != nil {
				return nil, err
			}

			domain, err := client.DescribeDomain(ctx, &codeartifact.DescribeDomainInput{
				Domain:      v.Name,
				DomainOwner: v.Owner,
			})
			if err != nil {
				return nil, err
			}

			policy, err := client.GetDomainPermissionsPolicy(ctx, &codeartifact.GetDomainPermissionsPolicyInput{
				Domain:      v.Name,
				DomainOwner: v.Owner,
			})
			if err != nil {
				return nil, err
			}

			resource := Resource{
				ARN:  *v.Arn,
				Name: *v.Name,
				Description: model.CodeArtifactDomainDescription{
					Domain: *domain.Domain,
					Policy: *policy.Policy,
					Tags:   tags.Tags,
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
