package aws

import (
	"context"
	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAwsBatchJobQueue(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_batchjob_queue",
		Description: "AWS BatchJob Queue",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("arn"), //TODO: change this to the primary key columns in model.go
			Hydrate:    opengovernance.GetBatchJobQueue,
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListBatchJobQueue,
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "id",
				Description: "The id of the queue.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Queue.Id")},
			{
				Name:        "name",
				Description: "The name of the queue.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Queue.Name")},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the queue",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Queue.ARN")}, // or generate it below
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Queue.Name")},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Queue.Tags"), // probably needs a transform function
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Queue.ARN").Transform(arnToAkas), // or generate it below (keep the Transform(arnToTurbotAkas) or use Transform(transform.EnsureStringArray))
			},
		}),
	}
}
