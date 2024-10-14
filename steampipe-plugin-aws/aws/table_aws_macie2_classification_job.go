package aws

import (
	"context"

	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsMacie2ClassificationJob(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_macie2_classification_job",
		Description: "AWS Macie2 Classification Job",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("job_id"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"ValidationException", "InvalidParameter"}),
			},
			Hydrate: opengovernance.GetMacie2ClassificationJob,
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListMacie2ClassificationJob,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "name", Require: plugin.Optional, Operators: []string{"=", "<>"}},
				{Name: "job_status", Require: plugin.Optional, Operators: []string{"=", "<>"}},
				{Name: "job_type", Require: plugin.Optional, Operators: []string{"=", "<>"}},
			},
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The custom name of the job.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ClassificationJob.Name")},
			{
				Name:        "job_id",
				Description: "The unique identifier for the job.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ClassificationJob.JobId")},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the job.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ClassificationJob.JobArn"),
			},
			{
				Name:        "job_status",
				Description: "The status of a classification job.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ClassificationJob.JobStatus")},
			{
				Name:        "job_type",
				Description: "The schedule for running a classification job.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ClassificationJob.JobType")},
			{
				Name:        "client_token",
				Description: "The token that was provided to ensure the idempotency of the request to create the job.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ClassificationJob.ClientToken")},
			{
				Name:        "created_at",
				Description: "The date and time, in UTC and extended ISO 8601 format, when the job was created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.ClassificationJob.CreatedAt")},
			{
				Name:        "last_run_time",
				Description: "This value indicates when the most recent run started.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.ClassificationJob.LastRunTime")},
			{
				Name:        "sampling_percentage",
				Description: "The sampling depth, as a percentage, that determines the percentage of eligible objects that the job analyzes.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.ClassificationJob.SamplingPercentage")},
			{
				Name:        "bucket_definitions",
				Description: "The namespace of the AWS service that provides the resource, or a custom-resource.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ClassificationJob.S3JobDefinition.BucketDefinitions")},
			{
				Name:        "custom_data_identifier_ids",
				Description: "The custom data identifiers that the job uses to analyze data.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ClassificationJob.CustomDataIdentifierIds")},
			{
				Name:        "last_run_error_status",
				Description: "Specifies whether any account- or bucket-level access errors occurred when a classification job ran.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ClassificationJob.LastRunErrorStatus")},
			{
				Name:        "s3_job_definition",
				Description: "Specifies which S3 buckets contain the objects that a classification job analyzes, and the scope of that analysis.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ClassificationJob.S3JobDefinition")},
			{
				Name:        "schedule_frequency",
				Description: "Specifies the recurrence pattern for running a classification job.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ClassificationJob.ScheduleFrequency")},
			{
				Name:        "statistics",
				Description: "Provides processing statistics for a classification job.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ClassificationJob.Statistics")},
			{
				Name:        "user_paused_details",
				Description: "Provides information about when a classification job was paused.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ClassificationJob.UserPausedDetails")},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ClassificationJob.Name")},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ClassificationJob.Tags")},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ClassificationJob.JobArn").Transform(transform.EnsureStringArray),
			},
		}),
	}
}
