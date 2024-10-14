package aws

import (
	"context"

	"github.com/kaytu-io/og-aws-describer/pkg/opengovernance-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsAppConfigApplication(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_appconfig_application",
		Description: "AWS AppConfig Application",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    opengovernance.GetAppConfigApplication,
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListAppConfigApplication,
		},

		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "id",
				Description: "The application ID.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Application.Id")},
			{
				Name:        "name",
				Description: "The application name.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Application.Name")},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) that identifies the application.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("ARN").Transform(arnToAkas),
			},
			{
				Name:        "description",
				Description: "The description of the application.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Application.Description")},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Application.Name")},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Tags")},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("ARN").Transform(transform.EnsureStringArray),
			},
		}),
	}
}
