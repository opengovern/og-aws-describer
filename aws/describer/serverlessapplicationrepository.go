package describer

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/serverlessapplicationrepository"
	"github.com/kaytu-io/kaytu-aws-describer/aws/model"
)

func ServerlessApplicationRepositoryApplication(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := serverlessapplicationrepository.NewFromConfig(cfg)
	paginator := serverlessapplicationrepository.NewListApplicationsPaginator(client, &serverlessapplicationrepository.ListApplicationsInput{})

	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, applicationSummary := range page.Applications {
			application, err := client.GetApplication(ctx, &serverlessapplicationrepository.GetApplicationInput{
				ApplicationId: applicationSummary.ApplicationId,
			})
			if err != nil {
				return nil, err
			}

			policy, err := client.GetApplicationPolicy(ctx, &serverlessapplicationrepository.GetApplicationPolicyInput{
				ApplicationId: applicationSummary.ApplicationId,
			})
			if err != nil {
				policy = &serverlessapplicationrepository.GetApplicationPolicyOutput{}
			}

			resource := Resource{
				Region: describeCtx.KaytuRegion,
				ARN:    *application.ApplicationId,
				Name:   *application.ApplicationId,
				Description: model.ServerlessApplicationRepositoryApplicationDescription{
					Application: *application,
					Statements:  policy.Statements,
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
