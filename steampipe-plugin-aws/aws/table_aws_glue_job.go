package aws

import (
	"context"

	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsGlueJob(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_glue_job",
		Description: "AWS Glue Job",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("name"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"EntityNotFoundException"}),
			},
			Hydrate: opengovernance.GetGlueJob,
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListGlueJob,
		},
		DefaultIgnoreConfig: &plugin.IgnoreConfig{
			ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"EntityNotFoundException"}),
		},

		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the GlueJob.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Job.Name")},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the GlueJob.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("ARN").Transform(arnToAkas),
			},
			{
				Name:        "allocated_capacity",
				Description: "[DEPRECATED] This column has been deprecated and will be removed in a future release, use max_capacity instead. The number of Glue data processing units (DPUs) that can be allocated when this job runs.",
				Type:        proto.ColumnType_DOUBLE,
				Transform:   transform.FromField("Description.Job.AllocatedCapacity")},
			{
				Name:        "created_on",
				Description: "The time and date that this job definition was created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Job.CreatedOn")},
			{
				Name:        "description",
				Description: "A description of the job.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Job.Description")},
			{
				Name:        "glue_version",
				Description: "Glue version determines the versions of Apache Spark and Python that Glue supports.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Job.GlueVersion")},
			{
				Name:        "last_modified_on",
				Description: "The last point in time when this job definition was modified.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Job.LastModifiedOn")},
			{
				Name:        "log_uri",
				Description: "This field is reserved for future use.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Job.LogUri")},
			{
				Name:        "max_capacity",
				Description: "The number of Glue data processing units (DPUs) that can be allocated when this job runs.",
				Type:        proto.ColumnType_DOUBLE,
				Transform:   transform.FromField("Description.Job.MaxCapacity")},
			{
				Name:        "max_retries",
				Description: "The maximum number of times to retry this job after a JobRun fails.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Job.MaxRetries")},
			{
				Name:        "number_of_workers",
				Description: "The number of workers of a defined workerType that are allocated when a job runs.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Job.NumberOfWorkers")},
			{
				Name:        "role",
				Description: "The name or Amazon Resource Name (ARN) of the IAM role associated with this job.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Job.Role")},
			{
				Name:        "security_configuration",
				Description: "The name of the SecurityConfiguration structure to be used with this job.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Job.SecurityConfiguration")},
			{
				Name:        "timeout",
				Description: "The job timeout in minutes.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Job.Timeout")},
			{
				Name:        "worker_type",
				Description: "The type of predefined worker that is allocated when a job runs. Accepts a value of Standard, G.1X, or G.2X.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Job.WorkerType")},
			{
				Name:        "command",
				Description: "The JobCommand that runs this job.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Job.Command")},
			{
				Name:        "connections",
				Description: "The connections used for this job.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Job.Connections")},
			{
				Name:        "default_arguments",
				Description: "The default arguments for this job, specified as name-value pairs.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Job.DefaultArguments")},
			{
				Name:        "execution_property",
				Description: "An ExecutionProperty specifying the maximum number of concurrent runs allowed for this job.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Job.ExecutionProperty")},
			{
				Name:        "job_bookmark",
				Description: "Defines a point that a job can resume processing.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Bookmark"),
			},
			{
				Name:        "non_overridable_arguments",
				Description: "Non-overridable arguments for this job, specified as name-value pairs.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Job.NonOverridableArguments")},
			{
				Name:        "notification_property",
				Description: "Specifies configuration properties of a job notification.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Job.NotificationProperty")},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Job.Name")},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("ARN").Transform(transform.EnsureStringArray),
			},
		}),
	}
}
