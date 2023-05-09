package describer

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/mediastore"
	"github.com/kaytu-io/kaytu-aws-describer/aws/model"
)

func MediaStoreContainer(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := mediastore.NewFromConfig(cfg)

	var values []Resource
	err := PaginateRetrieveAll(func(prevToken *string) (nextToken *string, err error) {
		containers, err := client.ListContainers(ctx, &mediastore.ListContainersInput{
			NextToken: prevToken,
		})
		if err != nil {
			return nil, err
		}

		for _, container := range containers.Containers {
			policy, err := client.GetContainerPolicy(ctx, &mediastore.GetContainerPolicyInput{
				ContainerName: container.Name,
			})
			if err != nil {
				policy = nil
			}

			tags, err := client.ListTagsForResource(ctx, &mediastore.ListTagsForResourceInput{
				Resource: container.ARN,
			})
			if err != nil {
				tags = &mediastore.ListTagsForResourceOutput{}
			}

			resource := Resource{
				Region: describeCtx.KaytuRegion,
				ARN:    *container.ARN,
				Name:   *container.Name,
				Description: model.MediaStoreContainerDescription{
					Container: container,
					Policy:    policy,
					Tags:      tags.Tags,
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

		return containers.NextToken, nil
	})
	if err != nil {
		return nil, err
	}

	return values, nil
}
