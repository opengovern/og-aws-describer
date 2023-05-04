package describer

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/amplify"
	"github.com/kaytu-io/kaytu-aws-describer/aws/model"
)

func AmplifyApp(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	//
	client := amplify.NewFromConfig(cfg)

	var values []Resource
	err := PaginateRetrieveAll(func(prevToken *string) (nextToken *string, err error) {
		output, err := client.ListApps(ctx, &amplify.ListAppsInput{
			MaxResults: 100,
			NextToken:  prevToken,
		})
		if err != nil {
			return nil, err
		}

		for _, item := range output.Apps {
			resource := Resource{
				Region: describeCtx.KaytuRegion,
				Name:   *item.Name,
				ARN:    *item.AppArn,
				ID:     *item.AppId,
				Description: model.AmplifyAppDescription{
					App: item,
				},
			}
			if stream != nil {
				m := *stream
				err := m(resource)
				if err != nil {
					return nil, err
				}
			} else {
				values = append(values, resource)
			}
		}
		return output.NextToken, nil
	})
	if err != nil {
		return nil, err
	}

	return values, nil
}
