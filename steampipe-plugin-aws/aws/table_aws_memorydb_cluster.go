package aws

import (
	"context"

	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAwsMemoryDbCluster(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_memorydb_cluster",
		Description: "AWS MemoryDb Cluster",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("name"),
			Hydrate:    opengovernance.GetMemoryDbCluster,
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListMemoryDbCluster,
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the cluster.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Cluster.Name")},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the cluster",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Cluster.ARN")},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Cluster.Name")},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(getMemoryDbClusterTurbotTags),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Cluster.ARN").Transform(arnToAkas),
			},
		}),
	}
}

//// TRANSFORM FUNCTIONS

func getMemoryDbClusterTurbotTags(_ context.Context, d *transform.TransformData) (interface{}, error) {
	tags := d.HydrateItem.(opengovernance.MemoryDbCluster).Description.Tags
	return memoryDbV2TagsToMap(tags)
}
