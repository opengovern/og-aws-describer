package aws

import (
	"context"

	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsElastiCacheSubnetGroup(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_elasticache_subnet_group",
		Description: "AWS ElastiCache Subnet Group",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("cache_subnet_group_name"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"CacheSubnetGroupNotFoundFault"}),
			},
			Hydrate: kaytu.GetElastiCacheSubnetGroup,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListElastiCacheSubnetGroup,
		},

		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "cache_subnet_group_name",
				Description: "The name of the cache subnet group.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.SubnetGroup.CacheSubnetGroupName")},
			{
				Name:        "arn",
				Description: "The ARN (Amazon Resource Name) of the cache subnet group.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.SubnetGroup.ARN"),
			},
			{
				Name:        "cache_subnet_group_description",
				Description: "The description of the cache subnet group.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.SubnetGroup.CacheSubnetGroupDescription")},
			{
				Name:        "vpc_id",
				Description: "The Amazon Virtual Private Cloud identifier (VPC ID) of the cache subnet group.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.SubnetGroup.VpcId")},
			{
				Name:        "subnets",
				Description: "A list of subnets associated with the cache subnet group.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.SubnetGroup.Subnets")},

			// Standard columns
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.SubnetGroup.CacheSubnetGroupName")},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.SubnetGroup.ARN").Transform(arnToAkas),
			},
		}),
	}
}
