package aws

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAwsSSMMaintenanceWindow(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_ssm_maintenance_window",
		Description: "AWS SSM Maintenance Window",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("window_id"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"DoesNotExistException"}),
			},
			Hydrate: opengovernance.GetSSMMaintenanceWindow,
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListSSMMaintenanceWindow,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "name", Require: plugin.Optional},
				{Name: "enabled", Require: plugin.Optional, Operators: []string{"=", "<>"}},
			},
		},

		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the Maintenance Window.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.MaintenanceWindow.Name")},
			{
				Name:        "window_id",
				Description: "The ID of the Maintenance Window.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.MaintenanceWindowIdentity.WindowId")},
			{
				Name:        "enabled",
				Description: "Indicates whether the Maintenance Window is enabled.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.MaintenanceWindow.Enabled")},
			{
				Name:        "allow_unassociated_targets",
				Description: "Indicates whether targets must be registered with the Maintenance Window before tasks can be defined for those targets.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.MaintenanceWindow.AllowUnassociatedTargets")},
			{
				Name:        "description",
				Description: "A description of the Maintenance Window.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.MaintenanceWindow.Description")},
			{
				Name:        "tags_src",
				Description: "A list of tags assigned to the Maintenance Window",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Tags")},
			{
				Name:        "duration",
				Description: "The duration of the Maintenance Window in hours.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.MaintenanceWindow.Duration")},
			{
				Name:        "cutoff",
				Description: "The number of hours before the end of the Maintenance Window that Systems Manager stops scheduling new tasks for execution.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.MaintenanceWindow.Cutoff")},
			{
				Name:        "schedule",
				Description: "The schedule of the Maintenance Window in the form of a cron or rate expression.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.MaintenanceWindow.Schedule")},
			{
				Name:        "schedule_offset",
				Description: "The number of days to wait to run a Maintenance Window after the scheduled CRON expression date and time.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.MaintenanceWindow.ScheduleOffset")},
			{
				Name:        "targets",
				Description: "The targets of Maintenance Window.",
				Type:        proto.ColumnType_JSON,

				Transform: transform.FromField("Description.Targets")},
			{
				Name:        "tasks",
				Description: "The Tasks of Maintenance Window.",
				Type:        proto.ColumnType_JSON,

				Transform: transform.FromField("Description.Tasks")},
			{
				Name:        "created_date",
				Description: "The date the maintenance window was created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.MaintenanceWindow.CreatedDate")},
			{
				Name:        "end_date",
				Description: "The date and time, in ISO-8601 Extended format, for when the maintenance window is scheduled to become inactive. The maintenance window will not run after this specified time.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.MaintenanceWindow.EndDate")},
			{
				Name:        "schedule_timezone",
				Description: "The schedule of the maintenance window in the form of a cron or rate expression.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.MaintenanceWindow.ScheduleTimezone")},
			{
				Name:        "start_date",
				Description: "The date and time, in ISO-8601 Extended format, for when the maintenance window is scheduled to become active.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.MaintenanceWindow.StartDate")},
			{
				Name:        "modified_date",
				Description: "The date the Maintenance Window was last modified.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.MaintenanceWindow.ModifiedDate")},
			{
				Name:        "next_execution_time",
				Description: "The next time the maintenance window will actually run, taking into account any specified times for the Maintenance Window to become active or inactive.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.MaintenanceWindow.NextExecutionTime")},

			/// Standard columns for all tables

			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.MaintenanceWindow.Name")},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,

				Transform: transform.FromField("Description.Tags").Transform(ssmMaintenanceWindowTagListToTurbotTags),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,

				Transform: transform.FromField("Description.ARN")},
		}),
	}
}

/// TRANSFORM FUNCTIONS

func ssmMaintenanceWindowTagListToTurbotTags(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	plugin.Logger(ctx).Trace("ssmMaintenanceWindowTagListToTurbotTags")
	tagList := d.Value.([]types.Tag)

	// Mapping the resource tags inside turbotTags
	var turbotTagsMap map[string]string
	if tagList != nil {
		turbotTagsMap = map[string]string{}
		for _, i := range tagList {
			turbotTagsMap[*i.Key] = *i.Value
		}
	} else {
		return nil, nil
	}

	return turbotTagsMap, nil
}
