package aws

import (
	"context"

	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAwsBatchJob(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_batch_job",
		Description: "AWS Batch Job",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("job_name"),
			Hydrate:    kaytu.GetBatchJob,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListBatchJob,
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "id",
				Description: "The id of the job.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Job.JobId")},
			{
				Name:        "job_name",
				Description: "The name of the job.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Job.JobName")},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the job",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Job.JobArn")},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Job.JobName")},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Job.JobArn").Transform(arnToAkas),
			},
		}),
	}
}
