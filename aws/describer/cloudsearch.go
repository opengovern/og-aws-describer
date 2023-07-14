package describer

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudsearch"
	"github.com/kaytu-io/kaytu-aws-describer/aws/model"
)

func CloudSearchDomain(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := cloudsearch.NewFromConfig(cfg)
	var values []Resource

	output, err := client.ListDomainNames(ctx, &cloudsearch.ListDomainNamesInput{})
	if err != nil {
		return nil, err
	}

	var domainList []string
	for domainName := range output.DomainNames {
		domainList = append(domainList, domainName)
	}

	domains, err := client.DescribeDomains(ctx, &cloudsearch.DescribeDomainsInput{
		DomainNames: domainList,
	})
	if err != nil {
		return nil, err
	}

	for _, domain := range domains.DomainStatusList {
		resource := Resource{
			Region: describeCtx.KaytuRegion,
			ARN:    *domain.ARN,
			Name:   *domain.DomainName,
			ID:     *domain.DomainId,
			Description: model.CloudSearchDomainDescription{
				DomainStatus: domain,
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
func CloudSearchDomainHandle(ctx context.Context) Resource {
	describeCtx := GetDescribeContext(ctx)
	resource := Resource{
		Region: describeCtx.KaytuRegion,
		ARN:    *domain.ARN,
		Name:   *domain.DomainName,
		ID:     *domain.DomainId,
		Description: model.CloudSearchDomainDescription{
			DomainStatus: domain,
		},
	}
}
func GetCloudSearchDomain(ctx context.Context, cfg aws.Config, fields map[string]string) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	domainList := fields["domainList"]
	client := cloudsearch.NewFromConfig(cfg)

	var values []Resource
	domains, err := client.DescribeDomains(ctx, &cloudsearch.DescribeDomainsInput{
		DomainNames: []string{domainList},
	})
	if err != nil {
		return nil, err
	}

	for _, domain := range domains.DomainStatusList {
		values = append(values, Resource{
			Region: describeCtx.KaytuRegion,
			ARN:    *domain.ARN,
			Name:   *domain.DomainName,
			ID:     *domain.DomainId,
			Description: model.CloudSearchDomainDescription{
				DomainStatus: domain,
			},
		})
	}
	return values, nil
}
