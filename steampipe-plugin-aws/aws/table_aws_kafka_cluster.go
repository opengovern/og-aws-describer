package aws

import (
	"context"

	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAwsKafkaCluster(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_kafka_cluster",
		Description: "AWS Kafka Cluster",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("cluster_name"),
			Hydrate:    opengovernance.GetKafkaCluster,
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListKafkaCluster,
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "cluster_name",
				Description: "The name of the cluster.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Cluster.ClusterName")},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the cluster",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Cluster.ClusterArn")},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Cluster.ClusterName")},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Cluster.Tags"),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Cluster.ClusterArn").Transform(arnToAkas),
			},
		}),
	}
}
