package describer

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/lightsail"
	"github.com/kaytu-io/kaytu-aws-describer/aws/model"
)

func LightsailInstance(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := lightsail.NewFromConfig(cfg)

	var values []Resource
	err := PaginateRetrieveAll(func(prevToken *string) (nextToken *string, err error) {
		instances, err := client.GetInstances(ctx, &lightsail.GetInstancesInput{
			PageToken: prevToken,
		})
		if err != nil {
			return nil, err
		}

		for _, instance := range instances.Instances {
			resource := Resource{
				Region: describeCtx.KaytuRegion,
				ARN:    *instance.Arn,
				Name:   *instance.Name,
				Description: model.LightsailInstanceDescription{
					Instance: instance,
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

		return instances.NextPageToken, nil
	})
	if err != nil {
		return nil, err
	}

	return values, nil
}
