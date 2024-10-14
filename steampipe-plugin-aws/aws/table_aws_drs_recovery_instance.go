package aws

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

// Not Mapped
// Don't need this one
func tableAwsDRSRecoveryInstance(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_drs_recovery_instance",
		Description: "AWS DRS RecoveryInstance",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("recovery_instance_id"),
			//Hydrate:    opengovernance.GetDRSRecoveryInstance,
		},
		List: &plugin.ListConfig{
			//Hydrate: opengovernance.ListDRSRecoveryInstance,
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "recovery_instance_id",
				Description: "The id of the recovery instance.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.RecoveryInstance.RecoveryInstanceID")},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the recovery instance",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.RecoveryInstance.Arn")},
			{
				Name:        "source_server_id",
				Description: "The source server ID that this recovery instance is associated with.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.RecoveryInstance.SourceServerID")},
			{
				Name:        "ec2_instance_id",
				Description: "The EC2 instance ID of the recovery instance.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Ec2InstanceID"),
			},
			{
				Name:        "ec2_instance_state",
				Description: "The state of the EC2 instance for this recovery instance.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "is_drill",
				Description: "Whether this recovery instance was created for a drill or for an actual recovery event.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "job_id",
				Description: "The ID of the Job that created the recovery instance.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("JobID"),
			},
			{
				Name:        "origin_environment",
				Description: "Environment (On Premises/AWS) of the instance that the recovery instance originated from.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "point_in_time_snapshot_date_time",
				Description: "The date and time of the Point in Time (PIT) snapshot that this recovery instance was launched from.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "data_replication_info",
				Description: "The Data Replication Info of the recovery instance.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "failback",
				Description: "An object representing failback related information of the recovery instance.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "recovery_instance_properties",
				Description: "Properties of the recovery instance machine.",
				Type:        proto.ColumnType_JSON,
			},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.RecoveryInstance.RecoveryInstanceID")},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.RecoveryInstance.Tags"),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.RecoveryInstance.Arn").Transform(arnToAkas),
			},
		}),
	}
}
