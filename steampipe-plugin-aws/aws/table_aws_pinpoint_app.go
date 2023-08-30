package aws

import (
	"context"

	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsPinpointApp(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_pinpoint_app",
		Description: "AWS Pinpoint App",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"NotFoundException"}),
			},
			Hydrate: kaytu.GetPinPointApp,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListPinPointApp,
		},

		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "id",
				Description: "The unique identifier for the application.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.App.Id")},
			{
				Name:        "name",
				Description: "The display name of the application.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.App.Name")},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the application.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.App.Arn")},
			{
				Name:        "last_modified_date",
				Description: "The date and time, in ISO 8601 format, when the application's settings were last modified.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Settings.LastModifiedDate")},
			{
				Name:        "campaign_hook",
				Description: "The settings for the AWS Lambda function to invoke by default as a code hook for campaigns in the application.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Settings.CampaignHook")},
			{
				Name:        "limits",
				Description: "The default sending limits for campaigns in the application.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Settings.Limits")},
			{
				Name:        "quiet_time",
				Description: "The default quiet time for campaigns in the application.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Settings.QuietTime")},

			// Steampipe standard columns

			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.App.Tags")},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.App.Name")},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.App.Arn").Transform(transform.EnsureStringArray),
			},
		}),
	}
}
