package aws

import (
	"context"

	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINATION

func tableAwsMGNApplication(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_mgn_application",
		Description: "AWS MGN Application",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListMgnApplication,
			IgnoreConfig: &plugin.IgnoreConfig{
				// UninitializedAccountException - This error comes up when the service is not enabled for the account.
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"UninitializedAccountException"}),
			},
			KeyColumns: []*plugin.KeyColumn{
				{Name: "application_id", Require: plugin.Optional},
				{Name: "wave_id", Require: plugin.Optional},
				{Name: "is_archived", Require: plugin.Optional, Operators: []string{"=", "<>"}},
			},
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "Application name.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Application.Name")},
			{
				Name:        "application_id",
				Description: "Application ID.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Application.ApplicationID")},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the application.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Application.Arn")},
			{
				Name:        "creation_date_time",
				Description: "Application creation dateTime.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Application.CreationDateTime")},
			{
				Name:        "description",
				Description: "Application description.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Application.Description")},
			{
				Name:        "is_archived",
				Description: "Application archival status.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Application.IsArchived")},
			{
				Name:        "last_modified_date_time",
				Description: "Application last modified dateTime.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Application.LastModifiedDateTime")},
			{
				Name:        "wave_id",
				Description: "Application wave ID.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Application.WaveID")},
			{
				Name:        "application_aggregated_status",
				Description: "Application aggregated status.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Application.ApplicationAggregatedStatus")},

			// Steampipe standard columns
			{
				Name:        "tags",
				Description: "A list of tags attached to the application.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Application.Tags")},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Application.Name")},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Application.Arn").Transform(transform.EnsureStringArray),
			},
		}),
	}
}
