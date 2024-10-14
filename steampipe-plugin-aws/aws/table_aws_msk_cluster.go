package aws

import (
	"context"
	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsMSKCluster(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_msk_cluster",
		Description: "AWS Managed Streaming for Apache Kafka",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("arn"),
			Hydrate:    opengovernance.GetKafkaCluster,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"NotFoundException"}),
			},
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListKafkaCluster,
		},

		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) that uniquely identifies the Cluster.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Cluster.ClusterArn"),
			},
			{
				Name:        "cluster_name",
				Description: "The name of the cluster.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Cluster.ClusterName")},
			{
				Name:        "active_operation_arn",
				Description: "Arn of active cluster operation.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Cluster.ActiveOperationArn")},
			{
				Name:        "cluster_type",
				Description: "The type of the cluster.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Cluster.ClusterType")},
			{
				Name:        "creation_time",
				Description: "The time when the cluster was created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Cluster.CreationTime")},
			{
				Name:        "current_version",
				Description: "The current version of the MSK cluster.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Cluster.CurrentVersion")},
			{
				Name:        "state",
				Description: "Settings for open monitoring using Prometheus.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Cluster.State")},
			{
				Name:        "cluster_configuration",
				Description: "Description of this MSK configuration.",
				Type:        proto.ColumnType_JSON,

				Transform: transform.FromField("Description.Configuration")},
			{
				Name:        "cluster_operation",
				Description: "Description of this MSK operation.",
				Type:        proto.ColumnType_JSON,

				Transform: transform.FromField("Description.ClusterOperationInfo")},
			{
				Name:        "provisioned",
				Description: "Information about the provisioned cluster.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Cluster.Provisioned")},
			{
				Name:        "state_info",
				Description: "State Info for the Amazon MSK cluster.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Cluster.StateInfo")},
			// Standard columns

			{
				Name:        "tags",
				Description: "A list of tags attached to the Cluster.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Cluster.Tags")},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Cluster.ClusterName")},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,

				Transform: transform.FromField("Description.Cluster.ClusterArn").Transform(transform.EnsureStringArray),
			},
		}),
	}
}
