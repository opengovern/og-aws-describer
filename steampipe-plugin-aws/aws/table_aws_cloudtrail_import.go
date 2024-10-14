package aws

import (
	"context"

	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsCloudtrailImport(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_cloudtrail_import",
		Description: "AWS CloudTrail Import",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("import_id"),
			Hydrate:    opengovernance.GetCloudTrailImport,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"UnsupportedOperationException", "ImportNotFoundException"}),
			},
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListCloudTrailImport,
			// For the location where the API operation is not supported, we receive UnsupportedOperationException.
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"UnsupportedOperationException"}),
			},
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:    "import_status",
					Require: plugin.Optional,
				},
			},
		},

		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "import_id",
				Description: "The ID of the import.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Import.ImportId")},
			{
				Name:        "created_timestamp",
				Description: "The timestamp of the import's creation.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Import.CreatedTimestamp")},
			{
				Name:        "import_status",
				Description: "The status of the import.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Import.ImportStatus")},
			{
				Name:        "end_event_time",
				Description: "Used with EndEventTime to bound a StartImport request, and limit imported trail events to only those events logged within a specified time period.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Import.EndEventTime")},
			{
				Name:        "start_event_time",
				Description: "Used with StartEventTime to bound a StartImport request, and limit imported trail events to only those events logged within a specified time period.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Import.StartEventTime")},
			{
				Name:        "updated_timestamp",
				Description: "The timestamp of the import's last update.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Import.UpdatedTimestamp")},
			{
				Name:        "destinations",
				Description: "The ARN of the destination event data store.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Import.Destinations")},
			{
				Name:        "import_source",
				Description: "The source S3 bucket.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Import.ImportSource")},
			{
				Name:        "import_statistics",
				Description: "Provides statistics for the import.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Import.ImportStatistics")},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Import.ImportId")},
		}),
	}
}
