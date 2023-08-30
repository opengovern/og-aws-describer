package aws

import (
	"context"

	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAwsHealthEvent(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_health_event",
		Description: "AWS Health Event",
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListHealthEvent,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "arn", Require: plugin.Optional},
				{Name: "availability_zone", Require: plugin.Optional},
				{Name: "end_time", Require: plugin.Optional},
				{Name: "event_type_category", Require: plugin.Optional},
				{Name: "event_type_code", Require: plugin.Optional},
				{Name: "last_updated_time", Require: plugin.Optional},
				{Name: "service", Require: plugin.Optional},
				{Name: "start_time", Require: plugin.Optional},
				{Name: "status_code", Require: plugin.Optional},
			},
		},

		Columns: awsKaytuColumns([]*plugin.Column{
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the HealthEvent.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Event.Arn")},
			{
				Name:        "availability_zone",
				Description: "The Amazon Web Services Availability Zone of the event.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Event.AvailabilityZone")},
			{
				Name:        "start_time",
				Description: "The date and time that the event began.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Event.StartTime")},
			{
				Name:        "end_time",
				Description: "The date and time that the event ended.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Event.EndTime")},
			{
				Name:        "event_scope_code",
				Description: "This parameter specifies if the Health event is a public Amazon Web Services service event or an account-specific event.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Event.EventScopeCode")},
			{
				Name:        "event_type_category",
				Description: "A list of event type category codes. Possible values are issue, accountNotification, or scheduledChange.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Event.EventTypeCategory")},
			{
				Name:        "event_type_code",
				Description: "The unique identifier for the event type.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Event.EventTypeCode")},
			{
				Name:        "last_updated_time",
				Description: "The most recent date and time that the event was updated.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Event.LastUpdatedTime")},
			{
				Name:        "service",
				Description: "The Amazon Web Services service that is affected by the event. For example, EC2, RDS.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Event.Service")},
			{
				Name:        "status_code",
				Description: "The most recent status of the event. Possible values are open, closed, and upcoming.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Event.StatusCode")},

			// Steampipe standard columns
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Event.Arn").Transform(transform.EnsureStringArray),
			},
		}),
	}
}
