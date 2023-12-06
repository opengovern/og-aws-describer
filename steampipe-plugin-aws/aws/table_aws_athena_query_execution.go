package aws

import (
	"context"
	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAwsAthenaQueryExecution(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_athena_query_execution",
		Description: "AWS Athena Query Execution",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    kaytu.GetAthenaQueryExecution,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListAthenaQueryExecution,
			KeyColumns: plugin.KeyColumnSlice{
				{
					Name:    "workgroup",
					Require: plugin.Optional,
				},
			},
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "id",
				Description: "The unique identifier for each query execution.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.QueryExecution.QueryExecutionId"),
			},
			{
				Name:        "workgroup",
				Description: "The name of the workgroup in which the query ran.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.QueryExecution.WorkGroup"),
			},
			{
				Name:        "catalog",
				Description: "The name of the data catalog used in the query execution.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.QueryExecution.QueryExecutionContext.Catalog"),
			},
			{
				Name:        "database",
				Description: "The name of the data database used in the query execution.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.QueryExecution.QueryExecutionContext.Database"),
			},
			{
				Name:        "query",
				Description: "The SQL query statements which the query execution ran.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.QueryExecution.Query"),
			},
			{
				Name:        "effective_engine_version",
				Description: "The engine version on which the query runs.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.QueryExecution.EngineVersion.EffectiveEngineVersion"),
			},
			{
				Name:        "selected_engine_version",
				Description: "The engine version requested by the users.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.QueryExecution.EngineVersion.SelectedEngineVersion"),
			},
			{
				Name:        "execution_parameters",
				Description: "A list of values for the parameters in a query.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.QueryExecution.ExecutionParameters"),
			},
			{
				Name:        "statement_type",
				Description: "The type of query statement that was run.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.QueryExecution.StatementType"),
			},
			{
				Name:        "substatement_type",
				Description: "The kind of query statement that was run.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.QueryExecution.SubstatementType"),
			},
			{
				Name:        "state",
				Description: "The state of query execution.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.QueryExecution.Status.State"),
			},
			{
				Name:        "state_change_reason",
				Description: "Further detail about the status of the query.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.QueryExecution.Status.StateChangeReason"),
			},
			{
				Name:        "submission_date_time",
				Description: "The date and time that the query was submitted.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.QueryExecution.Status.SubmissionDateTime"),
			},
			{
				Name:        "completion_date_time",
				Description: "The date and time that the query completed.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.QueryExecution.Status.CompletionDateTime"),
			},
			{
				Name:        "error_message",
				Description: "Contains a short description of the error that occurred.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.QueryExecution.Status.AthenaError.ErrorMessage"),
			},
			{
				Name:        "error_type",
				Description: "An integer value that provides specific information about an Athena query error.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.QueryExecution.Status.AthenaError.ErrorType"),
			},
			{
				Name:        "error_category",
				Description: "An integer value that specifies the category of a query failure error.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.QueryExecution.Status.AthenaError.ErrorCategory"),
			},
			{
				Name:        "retryable",
				Description: "True if the query might succeed if resubmitted.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.QueryExecution.Status.AthenaError.Retryable"),
			},
			{
				Name:        "data_manifest_location",
				Description: "The location and file name of a data manifest file.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.QueryExecution.Statistics.DataManifestLocation"),
			},
			{
				Name:        "data_scanned_in_bytes",
				Description: "The number of bytes in the data that was queried.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.QueryExecution.Statistics.DataScannedInBytes"),
			},
			{
				Name:        "engine_execution_time_in_millis",
				Description: "The number of milliseconds that the query took to execute.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.QueryExecution.Statistics.EngineExecutionTimeInMillis"),
			},
			{
				Name:        "query_planning_time_in_millis",
				Description: "The number of milliseconds that Athena took to plan the query processing flow.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.QueryExecution.Statistics.QueryPlanningTimeInMillis"),
			},
			{
				Name:        "query_queue_time_in_millis",
				Description: "The number of milliseconds that the query was in your query queue waiting for resources.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.QueryExecution.Statistics.QueryQueueTimeInMillis"),
			},
			{
				Name:        "service_processing_time_in_millis",
				Description: "The number of milliseconds that Athena took to finalize and publish the query results after the query engine finished running the query.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.QueryExecution.Statistics.ServiceProcessingTimeInMillis"),
			},
			{
				Name:        "total_execution_time_in_millis",
				Description: "The number of milliseconds that Athena took to run the query.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.QueryExecution.Statistics.TotalExecutionTimeInMillis"),
			},
			{
				Name:        "reused_previous_result",
				Description: "True if a previous query result was reused; false if the result was generated.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.QueryExecution.Statistics.ResultReuseInformation.ReusedPreviousResult"),
			},
			{
				Name:        "s3_acl_option",
				Description: "The Amazon S3 canned ACL that Athena should specify when storing query results.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.QueryExecution.ResultConfiguration.AclConfiguration.S3AclOption"),
			},
			{
				Name:        "encryption_option",
				Description: "Indicates whether Amazon S3 server-side encryption with Amazon S3-managed keys (SSE_S3), server-side encryption with KMS-managed keys (SSE_KMS), or client-side encryption with KMS-managed keys (CSE_KMS) is used.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.QueryExecution.ResultConfiguration.EncryptionConfiguration.EncryptionOption"),
			},
			{
				Name:        "kms_key",
				Description: "For SSE_KMS and CSE_KMS, this is the KMS key ARN or ID.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.QueryExecution.ResultConfiguration.EncryptionConfiguration.KmsKey"),
			},
			{
				Name:        "expected_bucket_owner",
				Description: "The Amazon Web Services account ID that you expect to be the owner of the Amazon S3 bucket specified by ResultConfiguration$OutputLocation.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.QueryExecution.ResultConfiguration.ExpectedBucketOwner"),
			},
			{
				Name:        "output_location",
				Description: "The location in Amazon S3 where your query results are stored.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.QueryExecution.ResultConfiguration.OutputLocation"),
			},
			{
				Name:        "result_reuse_by_age_enabled",
				Description: "True if previous query results can be reused when the query is run.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.QueryExecution.ResultReuseConfiguration.ResultReuseByAgeConfiguration.Enabled"),
			},
			{
				Name:        "result_reuse_by_age_mag_age_in_minutes",
				Description: "Specifies, in minutes, the maximum age of a previous query result that Athena should consider for reuse. The default is 60.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.QueryExecution.ResultReuseConfiguration.ResultReuseByAgeConfiguration.MaxAgeInMinutes"),
			},
		}),
	}
}
