package aws

import (
	"context"

	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAwsAppstreamFleet(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_appstream_fleet",
		Description: "AWS AppStream Fleet",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("name"),
			Hydrate:    opengovernance.GetAppStreamFleet,
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListAppStreamFleet,
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the fleet.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Fleet.Name"),
			},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) for the fleet.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Fleet.Arn"),
			},
			{
				Name:        "instance_type",
				Description: "The instance type to use when launching fleet instances.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Fleet.InstanceType"),
			},
			{
				Name:        "state",
				Description: "The current state for the fleet.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Fleet.State"),
			},
			{
				Name:        "created_time",
				Description: "The time the fleet was created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Fleet.CreatedTime"),
			},
			{
				Name:        "description",
				Description: "The description to display.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Fleet.Description"),
			},
			{
				Name:        "display_name",
				Description: "The fleet name to display.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Fleet.DisplayName"),
			},
			{
				Name:        "disconnect_timeout_in_seconds",
				Description: "The amount of time that a streaming session remains active after users disconnect. If they try to reconnect to the streaming session after a disconnection or network interruption within this time interval, they are connected to their previous session. Otherwise, they are connected to a new session with a new streaming instance. Specify a value between 60 and 360000.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Fleet.DisconnectTimeoutInSeconds"),
			},
			{
				Name:        "directory_name",
				Description: "The fully qualified name of the directory (for example, corp.example.com).",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Fleet.DomainJoinInfo.DirectoryName"),
			},
			{
				Name:        "organizational_unit_distinguished_name",
				Description: "The distinguished name of the organizational unit for computer accounts.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Fleet.DomainJoinInfo.OrganizationalUnitDistinguishedName"),
			},
			{
				Name:        "enable_default_internet_access",
				Description: "Indicates whether default internet access is enabled for the fleet.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Fleet.EnableDefaultInternetAccess"),
			},
			{
				Name:        "fleet_type",
				Description: "The fleet type. ALWAYS_ON Provides users with instant-on access to their apps. You are charged for all running instances in your fleet, even if no users are streaming apps. ON_DEMAND Provide users with access to applications after they connect, which takes one to two minutes. You are charged for instance streaming when users are connected and a small hourly fee for instances that are not streaming apps.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Fleet.FleetType"),
			},
			{
				Name:        "iam_role_arn",
				Description: "The ARN of the IAM role that is applied to the fleet.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Fleet.IamRoleArn"),
			},
			{
				Name:        "idle_disconnect_timeout_in_seconds",
				Description: "The amount of time that users can be idle (inactive) before they are disconnected from their streaming session and the DisconnectTimeoutInSeconds time interval begins.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Fleet.IdleDisconnectTimeoutInSeconds"),
			},
			{
				Name:        "image_arn",
				Description: "The ARN for the public, private, or shared image.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Fleet.ImageArn"),
			},
			{
				Name:        "image_name",
				Description: "The name of the image used to create the fleet.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Fleet.ImageName"),
			},
			{
				Name:        "max_concurrent_sessions",
				Description: "The maximum number of concurrent sessions for the fleet.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Fleet.MaxConcurrentSessions"),
			},
			{
				Name:        "max_user_duration_in_seconds",
				Description: "The maximum amount of time that a streaming session can remain active, in seconds. If users are still connected to a streaming instance five minutes before this limit is reached, they are prompted to save any open documents before being disconnected.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Fleet.MaxUserDurationInSeconds"),
			},
			{
				Name:        "platform",
				Description: "The platform of the fleet.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Fleet.Platform"),
			},
			{
				Name:        "stream_view",
				Description: "The AppStream 2.0 view that is displayed to your users when they stream from the fleet. When APP is specified, only the windows of applications opened by users display. When DESKTOP is specified, the standard desktop that is provided by the operating system displays. The default value is APP.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Fleet.StreamView"),
			},
			{
				Name:        "compute_capacity_status",
				Description: "The capacity status for the fleet.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Fleet.ComputeCapacityStatus"),
			},
			{
				Name:        "fleet_errors",
				Description: "The fleet errors.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Fleet.FleetErrors"),
			},
			{
				Name:        "session_script_s3_location",
				Description: "The S3 location of the session scripts configuration zip file. This only applies to Elastic fleets.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Fleet.SessionScriptS3Location"),
			},
			{
				Name:        "usb_device_filter_strings",
				Description: "The USB device filter strings associated with the fleet.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Fleet.UsbDeviceFilterStrings"),
			},
			{
				Name:        "vpc_config",
				Description: "The VPC configuration for the fleet.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Fleet.VpcConfig"),
			},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Fleet.DisplayName"),
			},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Tags"),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Fleet.Arn").Transform(transform.EnsureStringArray),
			},
		}),
	}
}
