package aws

import (
	"context"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

// Not Mapped
// Don't need this one
func tableAwsDRSJob(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_drs_job",
		Description: "AWS DRS Job",
		List: &plugin.ListConfig{
			KeyColumns: []*plugin.KeyColumn{
				{Name: "job_id", Require: plugin.Optional},
				{Name: "creation_date_time", Require: plugin.Optional, Operators: []string{">", ">=", "<", "<=", "="}},
				{Name: "end_date_time", Require: plugin.Optional, Operators: []string{">", ">=", "<", "<=", "="}},
			},
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"UninitializedAccountException", "BadRequestException"}),
			},
			//Hydrate: opengovernance.ListDRSJob,
		},

		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "job_id",
				Description: "The ID of the job.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Job.JobID")},
			{
				Name:        "arn",
				Description: "The ARN of a Job.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Job.Arn")},
			{
				Name:        "initiated_by",
				Description: "A string representing who initiated the Job.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Job.InitiatedBy")},
			{
				Name:        "status",
				Description: "The status of the Job.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Job.Status")},
			{
				Name:        "type",
				Description: "The type of the Job.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Job.Type")},
			{
				Name:        "creation_date_time",
				Description: "The date and time of when the Job was created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Job.CreationDateTime")},
			{
				Name:        "end_date_time",
				Description: "The date and time of when the Job ended.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Job.EndDateTime")},
			{
				Name:        "participating_servers",
				Description: "A list of servers that the Job is acting upon.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Job.ParticipatingServers")},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Job.JobID")},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Job.Tags")},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Job.Arn").Transform(arnToAkas),
			},
		}),
	}
}
