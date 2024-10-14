package aws

import (
	"context"

	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsCloudTrailQuery(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_cloudtrail_query",
		Description: "AWS CloudTrail Query",
		Get: &plugin.GetConfig{
			Hydrate:    opengovernance.GetCloudTrailQuery,
			KeyColumns: plugin.AllColumns([]string{"event_data_store_arn", "query_id"}),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"EventDataStoreNotFoundException", "QueryIdNotFoundException"}),
			},
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListCloudTrailQuery,
			KeyColumns: plugin.KeyColumnSlice{
				{
					Name:    "event_data_store_arn",
					Require: plugin.Optional,
				},
				{
					Name:    "query_status",
					Require: plugin.Optional,
				},
				{
					Name:      "creation_time",
					Require:   plugin.Optional,
					Operators: []string{"=", "<=", "<", ">", ">="},
				},
			},
		},

		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "query_id",
				Description: "The ID of the query.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Query.QueryId")},
			{
				Name:        "event_data_store_arn",
				Description: "The ID of the event data store.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.EventDataStoreARN")},
			{
				Name:        "creation_time",
				Description: "The creation time of the query.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Query.QueryStatistics.CreationTime")},
			{
				Name:        "delivery_s3_uri",
				Description: "The URI for the S3 bucket where CloudTrail delivered query results, if applicable.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Query.DeliveryS3Uri")},
			{
				Name:        "delivery_status",
				Description: "The delivery status.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Query.DeliveryStatus")},
			{
				Name:        "error_message",
				Description: "The error message returned if a query failed.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Query.ErrorMessage")},
			{
				Name:        "query_status",
				Description: "The status of a query. Values for QueryStatus include QUEUED, RUNNING, FINISHED, FAILED, TIMED_OUT, or CANCELLED.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Query.QueryStatus")},
			{
				Name:        "bytes_scanned",
				Description: "Gets metadata about a query, including the number of events that were matched, the total number of events scanned, the query run time in milliseconds, and the query's creation time.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Query.QueryStatistics.BytesScanned")},
			{
				Name:        "events_matched",
				Description: "The number of events that matched a query.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Query.QueryStatistics.EventsMatched")},
			{
				Name:        "events_scanned",
				Description: "The number of events that the query scanned in the event data store.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Query.QueryStatistics.EventsScanned")},
			{
				Name:        "execution_time_in_millis",
				Description: "The query's run time, in milliseconds.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Query.QueryStatistics.ExecutionTimeInMillis")},
			{
				Name:        "query_string",
				Description: "The SQL code of a query.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Query.QueryString")},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Query.QueryId")},
		}),
	}
}
