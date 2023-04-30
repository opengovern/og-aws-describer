package describer

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/appconfig"
	"github.com/kaytu-io/kaytu-aws-describer/aws/model"
)

func AppConfigApplication(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := appconfig.NewFromConfig(cfg)
	paginator := appconfig.NewListApplicationsPaginator(client, &appconfig.ListApplicationsInput{})

	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, application := range page.Items {
			arn := fmt.Sprintf("arn:%s:appconfig:%s:%s:application/%s", describeCtx.Partition, describeCtx.Region, describeCtx.AccountID, *application.Id)

			tags, err := client.ListTagsForResource(ctx, &appconfig.ListTagsForResourceInput{
				ResourceArn: aws.String(arn),
			})
			if err != nil {
				return nil, err
			}

			resource := Resource{
				Region: describeCtx.Region,
				ID:     *application.Id,
				Name:   *application.Name,
				ARN:    arn,
				Description: model.AppConfigApplicationDescription{
					Application: application,
					Tags:        tags.Tags,
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
