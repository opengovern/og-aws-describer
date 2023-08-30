package aws

import (
	"context"

	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsElastiCacheParameterGroup(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_elasticache_parameter_group",
		Description: "AWS ElastiCache Parameter Group",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("cache_parameter_group_name"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"CacheParameterGroupNotFound", "InvalidParameterValueException"}),
			},
			Hydrate: kaytu.GetElastiCacheParameterGroup,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListElastiCacheParameterGroup,
		},

		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "cache_parameter_group_name",
				Description: "The name of the cache parameter group.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ParameterGroup.CacheParameterGroupName")},
			{
				Name:        "arn",
				Description: "The ARN (Amazon Resource Name) of the cache parameter group.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("ARN").Transform(arnToAkas),
			},
			{
				Name:        "description",
				Description: "The description for the cache parameter group.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ParameterGroup.Description")},
			{
				Name:        "cache_parameter_group_family",
				Description: "The name of the cache parameter group family that this cache parameter group is compatible with.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ParameterGroup.CacheParameterGroupFamily")},
			{
				Name:        "is_global",
				Description: "Indicates whether the parameter group is associated with a Global Datastore.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.ParameterGroup.IsGlobal")},

			// Standard columns
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ParameterGroup.CacheParameterGroupName")},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("ARN").Transform(arnToAkas),
			},
		}),
	}
}
