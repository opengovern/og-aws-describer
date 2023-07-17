package describer

import (
	"context"
	"math"

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

	var values []Resource
	// OpenSearch API only allows 5 domains per request
	for i := 0; i < len(domainNames); i = i + 5 {
		domains, err := client.DescribeDomains(ctx, &opensearch.DescribeDomainsInput{
			DomainNames: domainNames[i:int(math.Min(float64(i+5), float64(len(domainNames))))],
		})
		if err != nil {
			return nil, err
		}

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
	}

	return values, nil
}
