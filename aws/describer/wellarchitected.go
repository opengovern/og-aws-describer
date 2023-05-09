package describer

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/wellarchitected"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/kaytu-io/kaytu-aws-describer/aws/model"
)

func WellArchitectedWorkload(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := wellarchitected.NewFromConfig(cfg)
	paginator := wellarchitected.NewListWorkloadsPaginator(client, &wellarchitected.ListWorkloadsInput{})

	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, v := range page.WorkloadSummaries {
			op, err := client.GetWorkload(ctx, &wellarchitected.GetWorkloadInput{
				WorkloadId: v.WorkloadId,
			})
			if err != nil {
				return nil, err
			}

			resource := Resource{
				Region: describeCtx.KaytuRegion,
				ARN:    *v.WorkloadArn,
				Name:   *v.WorkloadName,
				Description: model.WellArchitectedWorkloadDescription{
					WorkloadSummary: v,
					Workload:        op.Workload,
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
