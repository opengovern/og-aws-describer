package describer

import (
	"context"
	types2 "github.com/aws/aws-sdk-go-v2/service/docdb/types"
	_ "google.golang.org/genproto/googleapis/bigtable/admin/cluster/v1"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/docdb"
	"github.com/kaytu-io/kaytu-aws-describer/aws/model"
)

func DocDBCluster(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	client := docdb.NewFromConfig(cfg)
	paginator := docdb.NewDescribeDBClustersPaginator(client, &docdb.DescribeDBClustersInput{})

	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, cluster := range page.DBClusters {
			resource, err := DocDBClusterHandle(ctx, cfg, cluster)
			if err != nil {
				return nil, err
			}
			emptyResource := Resource{}
			if err == nil && resource == emptyResource {
				continue
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
func DocDBClusterHandle(ctx context.Context, cfg aws.Config, cluster types2.DBCluster) (Resource, error) {
	client := docdb.NewFromConfig(cfg)
	describeCtx := GetDescribeContext(ctx)

	tags, err := client.ListTagsForResource(ctx, &docdb.ListTagsForResourceInput{
		ResourceName: cluster.DBClusterArn,
	})
	if err != nil {
		if isErr(err, "ListTagsForResourceNotFound") || isErr(err, "InvalidParameterValue") {
			return Resource{}, nil
		}
		return Resource{}, err
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
	return resource, nil
}
func GetDocDBCluster(ctx context.Context, cfg aws.Config, fields map[string]string) ([]Resource, error) {
	client := docdb.NewFromConfig(cfg)
	dbClusterIdentifier := fields["identifier"]

	out, err := client.DescribeDBClusters(ctx, &docdb.DescribeDBClustersInput{
		DBClusterIdentifier: &dbClusterIdentifier,
	})
	if err != nil {
		return nil, err
	}

	var values []Resource
	for _, cluster := range out.DBClusters {

		resource, err := DocDBClusterHandle(ctx, cfg, cluster)
		if err != nil {
			return nil, err
		}
		emptyResource := Resource{}
		if err == nil && resource == emptyResource {
			return nil, nil
		}

		values = append(values, resource)
	}
	return values, nil
}
