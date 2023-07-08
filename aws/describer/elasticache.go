package describer

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/elasticache"
	"github.com/aws/aws-sdk-go-v2/service/elasticache/types"
	"github.com/kaytu-io/kaytu-aws-describer/aws/model"
)

func ElastiCacheReplicationGroup(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	client := elasticache.NewFromConfig(cfg)
	paginator := elasticache.NewDescribeReplicationGroupsPaginator(client, &elasticache.DescribeReplicationGroupsInput{})

	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, item := range page.ReplicationGroups {
			resource := elastiCacheReplicationGroupHandel(ctx, item)
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
func elastiCacheReplicationGroupHandel(ctx context.Context, item types.ReplicationGroup) Resource {
	describeCtx := GetDescribeContext(ctx)
	resource := Resource{
		Region: describeCtx.KaytuRegion,
		ARN:    *item.ARN,
		Name:   *item.ARN,
		Description: model.ElastiCacheReplicationGroupDescription{
			ReplicationGroup: item,
		},
	}
	return resource
}
func GetElastiCacheReplicationGroup(ctx context.Context, cfg aws.Config, fields map[string]string) ([]Resource, error) {
	CacheReplicationId := fields["CacheReplicationId"]
	client := elasticache.NewFromConfig(cfg)
	out, err := client.DescribeReplicationGroups(ctx, &elasticache.DescribeReplicationGroupsInput{
		ReplicationGroupId: &CacheReplicationId,
	})
	if err != nil {
		if isErr(err, "CacheReplicationGroupNotFound") || isErr(err, "InvalidParameterValue") {
			return nil, nil
		}
		return nil, err
	}
	var value []Resource
	for _, v := range out.ReplicationGroups {
		resource := elastiCacheReplicationGroupHandel(ctx, v)
		value = append(value, resource)
	}
	return value, nil
}

func ElastiCacheCluster(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	client := elasticache.NewFromConfig(cfg)
	paginator := elasticache.NewDescribeCacheClustersPaginator(client, &elasticache.DescribeCacheClustersInput{})

	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			if isErr(err, "CacheClusterNotFound") || isErr(err, "InvalidParameterValue") {
				continue
			}
			return nil, err
		}

		for _, cluster := range page.CacheClusters {
			resource, err := elastiCacheClusterHandel(ctx, cluster, client)
			if err != nil {
				return nil, err
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
func elastiCacheClusterHandel(ctx context.Context, cluster types.CacheCluster, client *elasticache.Client) (Resource, error) {
	describeCtx := GetDescribeContext(ctx)

	tagsOutput, err := client.ListTagsForResource(ctx, &elasticache.ListTagsForResourceInput{
		ResourceName: cluster.ARN,
	})
	if err != nil {
		if !isErr(err, "CacheClusterNotFound") && !isErr(err, "InvalidParameterValue") {
			return Resource{}, err
		} else {
			tagsOutput = &elasticache.ListTagsForResourceOutput{}
		}
	}

	resource := Resource{
		Region: describeCtx.KaytuRegion,
		ARN:    *cluster.ARN,
		Name:   *cluster.ARN,
		Description: model.ElastiCacheClusterDescription{
			Cluster: cluster,
			TagList: tagsOutput.TagList,
		},
	}
	return resource, nil
}
func GetElastiCacheCluster(ctx context.Context, cfg aws.Config, fields map[string]string) ([]Resource, error) {
	clusterID := fields["id"]
	client := elasticache.NewFromConfig(cfg)
	out, err := client.DescribeCacheClusters(ctx, &elasticache.DescribeCacheClustersInput{
		CacheClusterId: &clusterID,
	})
	if err != nil {
		if isErr(err, "CacheClusterNotFound") || isErr(err, "InvalidParameterValue") {
			return nil, nil
		}
		return nil, err
	}

	var values []Resource
	for _, cluster := range out.CacheClusters {
		resource, err := elastiCacheClusterHandel(ctx, cluster, client)
		if err != nil {
			return nil, err
		}
		values = append(values, resource)
	}
	return values, nil
}

func ElastiCacheParameterGroup(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := elasticache.NewFromConfig(cfg)
	paginator := elasticache.NewDescribeCacheParameterGroupsPaginator(client, &elasticache.DescribeCacheParameterGroupsInput{})

	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, cacheParameterGroup := range page.CacheParameterGroups {
			resource := Resource{
				Region: describeCtx.KaytuRegion,
				ARN:    *cacheParameterGroup.ARN,
				Name:   *cacheParameterGroup.CacheParameterGroupName,
				Description: model.ElastiCacheParameterGroupDescription{
					ParameterGroup: cacheParameterGroup,
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

func ElastiCacheReservedCacheNode(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	client := elasticache.NewFromConfig(cfg)
	paginator := elasticache.NewDescribeReservedCacheNodesPaginator(client, &elasticache.DescribeReservedCacheNodesInput{})

	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, reservedCacheNode := range page.ReservedCacheNodes {
			resource := elastiCacheReservedCacheNodeHandel(ctx, reservedCacheNode)
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
func elastiCacheReservedCacheNodeHandel(ctx context.Context, reservedCacheNode types.ReservedCacheNode) Resource {
	describeCtx := GetDescribeContext(ctx)
	resource := Resource{
		Region: describeCtx.KaytuRegion,
		ARN:    *reservedCacheNode.ReservationARN,
		ID:     *reservedCacheNode.ReservedCacheNodeId,
		Description: model.ElastiCacheReservedCacheNodeDescription{
			ReservedCacheNode: reservedCacheNode,
		},
	}
	return resource
}
func GetElastiCacheReservedCacheNode(ctx context.Context, cfg aws.Config, fields map[string]string) ([]Resource, error) {
	id := fields["id"]
	client := elasticache.NewFromConfig(cfg)
	out, err := client.DescribeReservedCacheNodes(ctx, &elasticache.DescribeReservedCacheNodesInput{
		ReservedCacheNodeId: &id,
	})
	if err != nil {
		if isErr(err, "ReservedCacheNodeNotFound") || isErr(err, "InvalidParameterValue") {
			return nil, nil
		}
		return nil, err
	}

	var value []Resource
	for _, v := range out.ReservedCacheNodes {
		resource := elastiCacheReservedCacheNodeHandel(ctx, v)
		value = append(value, resource)
	}
	return value, nil
}

func ElastiCacheSubnetGroup(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := elasticache.NewFromConfig(cfg)
	paginator := elasticache.NewDescribeCacheSubnetGroupsPaginator(client, &elasticache.DescribeCacheSubnetGroupsInput{})

	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, cacheSubnetGroup := range page.CacheSubnetGroups {
			resource := Resource{
				Region: describeCtx.KaytuRegion,
				ARN:    *cacheSubnetGroup.ARN,
				Name:   *cacheSubnetGroup.CacheSubnetGroupName,
				Description: model.ElastiCacheSubnetGroupDescription{
					SubnetGroup: cacheSubnetGroup,
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
