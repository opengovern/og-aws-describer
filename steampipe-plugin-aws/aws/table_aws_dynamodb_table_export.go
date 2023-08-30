package aws

import (
	"context"

	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAwsDynamoDBTableExport(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_dynamodb_table_export",
		Description: "AWS DynamoDB Table Export",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("arn"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"ResourceNotFoundException", "ExportNotFoundException"}),
			},
			Hydrate: kaytu.GetDynamoDbTableExport,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListDynamoDbTableExport,
		},

		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the export.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Export.ExportArn"),
			},
			{
				Name:        "export_status",
				Description: "Export can be in one of the following states: IN_PROGRESS, COMPLETED, or FAILED.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Export.ExportStatus")},
			{
				Name:        "billed_size_bytes",
				Description: "The billable size of the table export.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Export.BilledSizeBytes")},
			{
				Name:        "client_token",
				Description: "The client token that was provided for the export task. A client token makes calls to ExportTableToPointInTimeInput idempotent, meaning that multiple identical calls have the same effect as one single call.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Export.ClientToken")},
			{
				Name:        "end_time",
				Description: "The time at which the export task completed.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Export.EndTime")},
			{
				Name:        "export_format",
				Description: "The format of the exported data. Valid values for ExportFormat are DYNAMODB_JSON or ION.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Export.ExportFormat")},
			{
				Name:        "export_manifest",
				Description: "The name of the manifest file for the export task.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Export.ExportManifest")},
			{
				Name:        "export_time",
				Description: "Point in time from which table data was exported.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Export.ExportTime")},
			{
				Name:        "failure_code",
				Description: "Status code for the result of the failed export.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Export.FailureCode")},
			{
				Name:        "failure_message",
				Description: "Export failure reason description.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Export.FailureMessage")},
			{
				Name:        "item_count",
				Description: "The number of items exported.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Export.ItemCount")},
			{
				Name:        "s3_bucket",
				Description: "The name of the Amazon S3 bucket containing the export.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Export.S3Bucket")},
			{
				Name:        "s3_bucket_owner",
				Description: "The ID of the Amazon Web Services account that owns the bucket containing the export.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Export.S3BucketOwner")},
			{
				Name:        "s3_prefix",
				Description: "The Amazon S3 bucket prefix used as the file name and path of the exported snapshot.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Export.S3Prefix")},
			{
				Name:        "s3_sse_algorithm",
				Description: "Type of encryption used on the bucket where export data is stored.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Export.S3SseAlgorithm")},
			{
				Name:        "s3_sse_kms_key_id",
				Description: "The ID of the KMS managed key used to encrypt the S3 bucket where export data is stored (if applicable).",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Export.S3SseKmsKeyId")},
			{
				Name:        "start_time",
				Description: "The time at which the export task began.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Export.StartTime")},
			{
				Name:        "table_arn",
				Description: "The Amazon Resource Name (ARN) of the table that was exported.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Export.TableArn")},
			{
				Name:        "table_id",
				Description: "Unique ID of the table that was exported.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Export.TableId")},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Export.ExportArn").Transform(transform.EnsureStringArray),
			},
		}),
	}
}
