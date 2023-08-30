package aws

import (
	"context"

	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsCloudtrailEventDataStore(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_cloudtrail_event_data_store",
		Description: "AWS CloudTrail Event Data Store",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("arn"),
			Hydrate:    kaytu.GetCloudTrailEventDataStore,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListCloudTrailEventDataStore,
		},

		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the event data store.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.EventDataStore.Name")},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the event data store.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.EventDataStore.EventDataStoreArn"),
			},
			{
				Name:        "status",
				Description: "The status of an event data store.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.EventDataStore.Status")},
			{
				Name:        "created_timestamp",
				Description: "The timestamp of the event data store's creation.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.EventDataStore.CreatedTimestamp")},
			{
				Name:        "multi_region_enabled",
				Description: "Indicates whether the event data store includes events from all regions, or only from the region in which it was created.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.EventDataStore.MultiRegionEnabled")},
			{
				Name:        "organization_enabled",
				Description: "Indicates that an event data store is collecting logged events for an organization.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.EventDataStore.OrganizationEnabled")},
			{
				Name:        "retention_period",
				Description: "The retention period, in days.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.EventDataStore.RetentionPeriod")},
			{
				Name:        "termination_protection_enabled",
				Description: "Indicates whether the event data store is protected from termination.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.EventDataStore.TerminationProtectionEnabled")},
			{
				Name:        "updated_timestamp",
				Description: "The timestamp showing when an event data store was updated, if applicable. UpdatedTimestamp is always either the same or newer than the time shown in CreatedTimestamp.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.EventDataStore.UpdatedTimestamp")},
			{
				Name:        "advanced_event_selectors",
				Description: "The advanced event selectors that were used to select events for the data store.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.EventDataStore.AdvancedEventSelectors")},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.EventDataStore.EventDataStoreArn").Transform(transform.EnsureStringArray),
			},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.EventDataStore.Name")},
		}),
	}
}
