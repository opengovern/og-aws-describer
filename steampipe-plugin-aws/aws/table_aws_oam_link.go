package aws

import (
	"context"

	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAwsOAMLink(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_oam_link",
		Description: "AWS OAM Link",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("arn"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"InvalidParameterException", "ResourceNotFoundException"}),
			},
			Hydrate: kaytu.GetOAMLink,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListOAMLink,
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "id",
				Description: "The random ID string that Amazon Web Service generates as part of the link ARN.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Link.Id")},
			{
				Name:        "arn",
				Description: "The ARN of the link.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Link.Arn")},
			{
				Name:        "sink_arn",
				Description: "The ARN of the sink that this link is attached to.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Link.SinkArn")},
			{
				Name:        "label",
				Description: "The label that was assigned to this link at creation, with the variables resolved to their actual values.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Link.Label")},
			{
				Name:        "label_template",
				Description: "The exact label template that was specified when the link was created, with the template variables not resolved.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Link.LabelTemplate")},
			{
				Name:        "resource_types",
				Description: "The resource types supported by this link.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Link.ResourceTypes")},
			/// Standard columns for all tables

			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Link.Label")},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Link.Tags")},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Link.Tags").Transform(transform.EnsureStringArray),
			},
		}),
	}
}
