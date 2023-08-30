package aws

import (
	"context"

	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAwsPipes(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_pipes_pipe",
		Description: "AWS Pipes Pipe",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("name"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"NotFoundException"}),
			},
			Hydrate: kaytu.GetPipesPipe,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListPipesPipe,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "current_state", Require: plugin.Optional},
				{Name: "desired_state", Require: plugin.Optional},
				{Name: "source_prefix", Require: plugin.Optional},
				{Name: "target_prefix", Require: plugin.Optional},
			},
		},

		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the pipe.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Pipe.Name")},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the pipe.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Pipe.Arn")},
			{
				Name:        "creation_time",
				Description: "The time the pipe was created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Pipe.CreationTime")},
			{
				Name:        "current_state",
				Description: "The state the pipe is in.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Pipe.CurrentState")},
			{
				Name:        "description",
				Description: "A description of the pipe.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.PipeOutput.Description")},
			{
				Name:        "desired_state",
				Description: "The state the pipe should be in.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Pipe.DesiredState")},
			{
				Name:        "enrichment",
				Description: "The ARN of the enrichment resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Pipe.Enrichment")},
			{
				Name:        "last_modified_time",
				Description: "When the pipe was last updated.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Pipe.LastModifiedTime")},
			{
				Name:        "role_arn",
				Description: "The ARN of the role that allows the pipe to send data to the target.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.PipeOutput.RoleArn")},
			{
				Name:        "source",
				Description: "The ARN of the source resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Pipe.Source")},
			{
				Name:        "source_prefix",
				Description: "The prefix matching the pipe source.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromQual("source_prefix"),
			},
			{
				Name:        "state_reason",
				Description: "The reason the pipe is in its current state.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Pipe.StateReason")},
			{
				Name:        "target",
				Description: "The ARN of the target resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Pipe.Target")},
			{
				Name:        "target_prefix",
				Description: "The prefix matching the pipe target.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromQual("target_prefix"),
			},
			{
				Name:        "enrichment_parameters",
				Description: "The parameters required to set up enrichment on your pipe.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.PipeOutput.EnrichmentParameters")},
			{
				Name:        "target_parameters",
				Description: "The parameters required to set up a target for your pipe.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.PipeOutput.TargetParameters")},

			// Standard columns for all tables

			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Pipe.Name")},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.PipeOutput.Tags")},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Pipe.Arn").Transform(transform.EnsureStringArray),

				//// LIST FUNCTION
			},
		}),
	}
}

// Get client

// Unsupported region, return no data

// Limiting the results

// Default to the maximum allowed

// API doesn't support aws-go-sdk-v2 paginator as of date

// Context may get cancelled due to manual cancellation or if the limit has been reached

//// HYDRATE FUNCTIONS

// Create Session

// Unsupported region, return no data

// Build the params

// Get call
