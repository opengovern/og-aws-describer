package describer

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/servicediscovery"
	"github.com/kaytu-io/kaytu-aws-describer/aws/model"
)

func ServiceDiscoveryService(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := servicediscovery.NewFromConfig(cfg)

	paginator := servicediscovery.NewListServicesPaginator(client, &servicediscovery.ListServicesInput{})
	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, item := range page.Services {
			tag, err := client.ListTagsForResource(ctx, &servicediscovery.ListTagsForResourceInput{
				ResourceARN: item.Arn,
			})
			if err != nil {
				return nil, err
			}

			resource := Resource{
				Region: describeCtx.Region,
				ID:     *item.Id,
				Description: model.ServiceDiscoveryServiceDescription{
					Service: item,
					Tags:    tag.Tags,
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

func ServiceDiscoveryNamespace(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := servicediscovery.NewFromConfig(cfg)

	paginator := servicediscovery.NewListNamespacesPaginator(client, &servicediscovery.ListNamespacesInput{})
	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, v := range page.Namespaces {
			tag, err := client.ListTagsForResource(ctx, &servicediscovery.ListTagsForResourceInput{
				ResourceARN: v.Arn,
			})
			if err != nil {
				return nil, err
			}

			resource := Resource{
				Region: describeCtx.KaytuRegion,
				ID:     *v.Id,
				Name:   *v.Name,
				Description: model.ServiceDiscoveryNamespace{
					Namespace: v,
					Tags:      tag.Tags,
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

func ServiceDiscoveryInstance(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := servicediscovery.NewFromConfig(cfg)

	services, err := client.ListServices(ctx, &servicediscovery.ListServicesInput{})
	if err != nil {
		return nil, err
	}
	var values []Resource
	for _, service := range services.Services {
		paginator := servicediscovery.NewListInstancesPaginator(client, &servicediscovery.ListInstancesInput{
			ServiceId: service.Id,
		})
		for paginator.HasMorePages() {
			page, err := paginator.NextPage(ctx)
			if err != nil {
				return nil, err
			}
			for _, v := range page.Instances {
				resource := Resource{
					Region: describeCtx.Region,
					ID:     *v.Id,
					Name:   *v.Id,
					Description: model.ServiceDiscoveryInstance{
						Instance:  v,
						ServiceId: service.Id,
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
