package describer

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/firehose"
	"github.com/kaytu-io/kaytu-aws-describer/aws/model"
)

func FirehoseDeliveryStream(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := firehose.NewFromConfig(cfg)

	var values []Resource

	err := PaginateRetrieveAll(func(prevToken *string) (lastName *string, err error) {
		deliveryStreams, err := client.ListDeliveryStreams(ctx, &firehose.ListDeliveryStreamsInput{
			ExclusiveStartDeliveryStreamName: prevToken,
		})
		if err != nil {
			return nil, err
		}
		for _, deliveryStreamName := range deliveryStreams.DeliveryStreamNames {
			lastName = &deliveryStreamName
			deliveryStream, err := client.DescribeDeliveryStream(ctx, &firehose.DescribeDeliveryStreamInput{
				DeliveryStreamName: &deliveryStreamName,
			})
			if err != nil {
				return nil, err
			}
			tags, err := client.ListTagsForDeliveryStream(ctx, &firehose.ListTagsForDeliveryStreamInput{
				DeliveryStreamName: &deliveryStreamName,
			})
			if err != nil {
				return nil, err
			}
			resource := Resource{
				Region: describeCtx.KaytuRegion,
				ARN:    *deliveryStream.DeliveryStreamDescription.DeliveryStreamARN,
				Name:   *deliveryStream.DeliveryStreamDescription.DeliveryStreamName,
				Description: model.FirehoseDeliveryStreamDescription{
					DeliveryStream: *deliveryStream.DeliveryStreamDescription,
					Tags:           tags.Tags,
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
		if deliveryStreams.HasMoreDeliveryStreams == nil || !*deliveryStreams.HasMoreDeliveryStreams {
			return nil, nil
		}
		return lastName, nil
	})
	if err != nil {
		return nil, err
	}

	return values, nil
}
