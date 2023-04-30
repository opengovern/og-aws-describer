package describer

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/mq"
	"github.com/kaytu-io/kaytu-aws-describer/aws/model"
)

func MQBroker(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := mq.NewFromConfig(cfg)
	paginator := mq.NewListBrokersPaginator(client, &mq.ListBrokersInput{})

	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, v := range page.BrokerSummaries {
			tags, err := client.ListTags(ctx, &mq.ListTagsInput{
				ResourceArn: v.BrokerArn,
			})
			if err != nil {
				return nil, err
			}

			resource := Resource{
				Region: describeCtx.Region,
				ARN:    *v.BrokerArn,
				Name:   *v.BrokerName,
				Description: model.MQBrokerDescription{
					Broker: v,
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
