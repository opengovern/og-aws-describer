package describer

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/neptune"
	"github.com/kaytu-io/kaytu-aws-describer/aws/model"
)

func NeptuneDatabase(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := neptune.NewFromConfig(cfg)
	paginator := neptune.NewDescribeDBInstancesPaginator(client, &neptune.DescribeDBInstancesInput{})

	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, v := range page.DBInstances {
			tags, err := client.ListTagsForResource(ctx, &neptune.ListTagsForResourceInput{
				ResourceName: v.DBInstanceArn,
			})
			if err != nil {
				return nil, err
			}

			resource := Resource{
				Region: describeCtx.Region,
				ARN:    *v.DBInstanceArn,
				Name:   *v.DBClusterIdentifier,
				Description: model.NeptuneDatabaseDescription{
					Database: v,
					Tags:     tags.TagList,
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
