package describer

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/securitylake"
	"github.com/kaytu-io/kaytu-aws-describer/aws/model"
)

func SecurityLakeDataLake(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := securitylake.NewFromConfig(cfg)

	var values []Resource
	lakes, err := client.GetDatalake(ctx, &securitylake.GetDatalakeInput{})
	if err != nil {
		if isErr(err, "AccessDeniedException") {
			return nil, nil
		} else {
			return nil, err
		}
	}
	for lakeKey, lake := range lakes.Configurations {
		resource := Resource{
			Region: describeCtx.KaytuRegion,
			Name:   lakeKey,
			Description: model.SecurityLakeDataLakeDescription{
				DataLake: lake,
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

func SecurityLakeSubscriber(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := securitylake.NewFromConfig(cfg)

	var values []Resource
	paginator := securitylake.NewListSubscribersPaginator(client, &securitylake.ListSubscribersInput{})
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			if isErr(err, "AccessDeniedException") {
				continue
			}
			return nil, err
		}

		for _, subscriber := range page.Subscribers {
			resource := Resource{
				Region: describeCtx.KaytuRegion,
				Name:   *subscriber.SubscriberName,
				Description: model.SecurityLakeSubscriberDescription{
					Subscriber: subscriber,
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
