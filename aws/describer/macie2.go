package describer

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/macie2"
	"github.com/kaytu-io/kaytu-aws-describer/aws/model"
)

func Macie2ClassificationJob(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := macie2.NewFromConfig(cfg)

	var values []Resource
	err := PaginateRetrieveAll(func(prevToken *string) (nextToken *string, err error) {
		classificationJobs, err := client.ListClassificationJobs(ctx, &macie2.ListClassificationJobsInput{
			NextToken: prevToken,
		})
		if err != nil {
			return nil, err
		}

		for _, jobSummary := range classificationJobs.Items {
			job, err := client.DescribeClassificationJob(ctx, &macie2.DescribeClassificationJobInput{
				JobId: jobSummary.JobId,
			})
			if err != nil {
				return nil, err
			}
			resource := Resource{
				Region: describeCtx.KaytuRegion,
				ARN:    *job.JobArn,
				Name:   *job.Name,
				ID:     *job.JobId,
				Description: model.Macie2ClassificationJobDescription{
					ClassificationJob: *job,
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

		return classificationJobs.NextToken, nil
	})
	if err != nil {
		return nil, err
	}

	return values, nil
}
