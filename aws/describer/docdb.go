package describer

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/docdb"
	"github.com/kaytu-io/kaytu-aws-describer/aws/model"
)

func DocDBCluster(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	//
	client := docdb.NewFromConfig(cfg)
	paginator := docdb.NewDescribeDBClustersPaginator(client, &docdb.DescribeDBClustersInput{})

	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, cluster := range page.DBClusters {
			tags, err := client.ListTagsForResource(ctx, &docdb.ListTagsForResourceInput{
				ResourceName: cluster.DBClusterArn,
			})
			if err != nil {
				return nil, err
			}

			resource := Resource{
				Region: describeCtx.KaytuRegion,
				ID:     *cluster.DBClusterIdentifier,
				ARN:    *cluster.DBClusterArn,
				Description: model.DocDBClusterDescription{
					DBCluster: cluster,
					Tags:      tags.TagList,
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
