package describer

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/sfn"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/kaytu-io/kaytu-aws-describer/aws/model"
)

func StepFunctionsStateMachine(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)

	client := sfn.NewFromConfig(cfg)
	paginator := sfn.NewListStateMachinesPaginator(client, &sfn.ListStateMachinesInput{})

	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, v := range page.StateMachines {
			data, err := client.DescribeStateMachine(ctx, &sfn.DescribeStateMachineInput{
				StateMachineArn: v.StateMachineArn,
			})
			if err != nil {
				return nil, err
			}

			var name string
			if data.Name != nil {
				name = *data.Name
			}

			tags, err := client.ListTagsForResource(ctx, &sfn.ListTagsForResourceInput{
				ResourceArn: v.StateMachineArn,
			})
			if err != nil {
				return nil, err
			}

			resource := Resource{
				Region: describeCtx.KaytuRegion,
				ARN:    *v.StateMachineArn,
				Name:   name,
				Description: model.StepFunctionsStateMachineDescription{
					StateMachineItem: v,
					StateMachine:     data,
					Tags:             tags.Tags,
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
