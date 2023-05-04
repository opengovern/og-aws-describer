package describer

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/opensearch"
	"github.com/kaytu-io/kaytu-aws-describer/aws/model"
)

func OpenSearchDomain(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := opensearch.NewFromConfig(cfg)

	domainNamesOutput, err := client.ListDomainNames(ctx, &opensearch.ListDomainNamesInput{})
	if err != nil {
		return nil, err
	}
	domainNames := make([]string, 0, len(domainNamesOutput.DomainNames))
	for _, domainName := range domainNamesOutput.DomainNames {
		domainNames = append(domainNames, *domainName.DomainName)
	}

	domains, err := client.DescribeDomains(ctx, &opensearch.DescribeDomainsInput{
		DomainNames: domainNames,
	})
	if err != nil {
		return nil, err
	}

	var values []Resource
	for _, domain := range domains.DomainStatusList {
		tags, err := client.ListTags(ctx, &opensearch.ListTagsInput{
			ARN: domain.ARN,
		})
		if err != nil {
			return nil, err
		}

		resource := Resource{
			Region: describeCtx.KaytuRegion,
			ARN:    *domain.ARN,
			Name:   *domain.DomainName,
			Description: model.OpenSearchDomainDescription{
				Domain: domain,
				Tags:   tags.TagList,
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
