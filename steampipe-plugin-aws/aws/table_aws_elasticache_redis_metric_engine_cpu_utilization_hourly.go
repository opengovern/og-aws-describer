package aws

import (
	"context"

	"github.com/aws/aws-sdk-go/service/elasticache"
	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

// TABLE DEFINITION
func tableAwsElasticacheRedisEngineCPUUtilizationHourly(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_elasticache_redis_metric_engine_cpu_utilization_hourly",
		Description: "AWS Elasticache Redis EngineCPUUtilization metric (Hourly)",
		List: &plugin.ListConfig{
			ParentHydrate: kaytu.ListElastiCacheCluster,
			Hydrate:       listElastiCacheMetricEngineCPUUtilizationHourly,
		},

		Columns: awsKaytuRegionalColumns(cwMetricColumns(
			[]*plugin.Column{
				{
					Name:        "cache_cluster_id",
					Description: "The cache cluster id.",
					Type:        proto.ColumnType_STRING,
					Transform:   transform.FromField("DimensionValue"),
				},
			})),
	}
}

func listElastiCacheMetricEngineCPUUtilizationHourly(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	cacheClusterConfiguration := h.Item.(*elasticache.CacheCluster)
	return listCWMetricStatistics(ctx, d, "Hourly", "AWS/ElastiCache", "EngineCPUUtilization", "CacheClusterId", *cacheClusterConfiguration.CacheClusterId)
}