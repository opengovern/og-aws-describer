package describer

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	dms "github.com/aws/aws-sdk-go-v2/service/databasemigrationservice"
	"github.com/kaytu-io/kaytu-aws-describer/aws/model"
)

func DMSReplicationInstance(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := dms.NewFromConfig(cfg)

	paginator := dms.NewDescribeReplicationInstancesPaginator(client,
		&dms.DescribeReplicationInstancesInput{})

	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, item := range page.ReplicationInstances {
			tags, err := client.ListTagsForResource(ctx, &dms.ListTagsForResourceInput{
				ResourceArn: item.ReplicationInstanceArn,
			})
			if err != nil {
				return nil, err
			}

			resource := Resource{
				Region: describeCtx.KaytuRegion,
				ARN:    *item.ReplicationInstanceArn,
				Name:   *item.ReplicationInstanceIdentifier,
				Description: model.DMSReplicationInstanceDescription{
					ReplicationInstance: item,
					Tags:                tags.TagList,
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
