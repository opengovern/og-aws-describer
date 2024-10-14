package aws

import (
	"context"

	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAwsOAMSink(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_oam_sink",
		Description: "AWS OAM Sink",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("arn"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"InvalidParameterValue", "ResourceNotFoundException"}),
			},
			Hydrate: opengovernance.GetOAMSink,
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListOAMSink,
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the sink.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Sink.Name")},
			{
				Name:        "id",
				Description: "The random ID string that Amazon Web Service generates as part of the sink ARN.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Sink.Id")},
			{
				Name:        "arn",
				Description: "The ARN of the sink.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Sink.Arn")},
			/// Standard columns for all tables
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Sink.Name")},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Tags")},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Sink.Arn").Transform(transform.EnsureStringArray),
			},
		}),
	}
}
