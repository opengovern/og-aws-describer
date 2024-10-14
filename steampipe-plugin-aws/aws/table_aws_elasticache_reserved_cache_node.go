package aws

import (
	"context"

	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsElastiCacheReservedCacheNode(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_elasticache_reserved_cache_node",
		Description: "AWS ElastiCache Reserved Cache Node",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("reserved_cache_node_id"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"ReservedCacheNodeNotFound"}),
			},
			Hydrate: opengovernance.GetElastiCacheReservedCacheNode,
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListElastiCacheReservedCacheNode,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "cache_node_type", Require: plugin.Optional},
				{Name: "duration", Require: plugin.Optional},
				{Name: "offering_type", Require: plugin.Optional},
				{Name: "reserved_cache_nodes_offering_id", Require: plugin.Optional},
			},
		},

		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "reserved_cache_node_id",
				Description: "The unique identifier for the reservation.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ReservedCacheNode.ReservedCacheNodeId")},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the reserved cache node.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ReservedCacheNode.ReservationARN"),
			},
			{
				Name:        "reserved_cache_nodes_offering_id",
				Description: "The offering identifier.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ReservedCacheNode.ReservedCacheNodesOfferingId")},
			{
				Name:        "state",
				Description: "The state of the reserved cache node.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ReservedCacheNode.State")},
			{
				Name:        "cache_node_type",
				Description: "The cache node type for the reserved cache nodes.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ReservedCacheNode.CacheNodeType")},
			{
				Name:        "cache_node_count",
				Description: "The number of cache nodes that have been reserved.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.ReservedCacheNode.CacheNodeCount")},
			{
				Name:        "duration",
				Description: "The duration of the reservation in seconds.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.ReservedCacheNode.Duration")},
			{
				Name:        "fixed_price",
				Description: "The fixed price charged for this reserved cache node.",
				Type:        proto.ColumnType_DOUBLE,
				Transform:   transform.FromField("Description.ReservedCacheNode.FixedPrice")},
			{
				Name:        "offering_type",
				Description: "The offering type of this reserved cache node.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ReservedCacheNode.OfferingType")},
			{
				Name:        "product_description",
				Description: "The description of the reserved cache node.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ReservedCacheNode.ProductDescription")},
			{
				Name:        "start_time",
				Description: "The time the reservation started.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.ReservedCacheNode.StartTime")},
			{
				Name:        "usage_price",
				Description: "The hourly price charged for this reserved cache node.",
				Type:        proto.ColumnType_DOUBLE,
				Transform:   transform.FromField("Description.ReservedCacheNode.UsagePrice")},
			{
				Name:        "recurring_charges",
				Description: "The recurring price charged to run this reserved cache node.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ReservedCacheNode.RecurringCharges")},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ReservedCacheNode.ReservedCacheNodeId")},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ReservedCacheNode.ReservationARN").Transform(transform.EnsureStringArray),
			},
		}),
	}
}
