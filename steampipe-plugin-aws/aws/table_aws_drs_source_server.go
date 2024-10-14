package aws

import (
	"context"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

// Not Mapped
// Don't need this one
func tableAwsDRSSourceServer(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_drs_source_server",
		Description: "AWS DRS SourceServer",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("source_server_id"),
			//Hydrate:    opengovernance.GetDRSSourceServer,
		},
		List: &plugin.ListConfig{
			//Hydrate: opengovernance.ListDRSSourceServer,
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "source_server_id",
				Description: "The id of the source server.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.SourceServer.SourceServerID")},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the source server",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.SourceServer.Arn")},
			{
				Name:        "recovery_instance_id",
				Description: "The ID of the Recovery Instance associated with this Source Server.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.SourceServer.RecoveryInstanceId")},
			{
				Name:        "source_properties",
				Description: "The source properties of the Source Server.",
				Transform:   transform.FromField("Description.SourceServer.SourceProperties")},
			{
				Name:        "data_replication_info",
				Description: "The Data Replication Info of the Source Server.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.SourceServer.DataReplicationInfo")},
			{
				Name:        "last_launch_result",
				Description: "The status of the last recovery launch of this Source Server.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.SourceServer.LastLaunchResult")},
			{
				Name:        "life_cycle",
				Description: "The lifecycle information of this Source Server.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.SourceServer.LifeCycle")},
			{
				Name:        "replication_direction",
				Description: "Replication direction of the Source Server.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.SourceServer.ReplicationDirection")},
			{
				Name:        "reversed_direction_source_server_arn",
				Description: "For EC2-originated Source Servers which have been failed over and then failed back, this value will mean the ARN of the Source Server on the opposite replication direction.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.SourceServer.ReversedDirectionSourceServerArn")},
			{
				Name:        "source_cloud_properties",
				Description: "Source cloud properties of the Source Server.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.SourceServer.SourceCloudProperties")},
			{
				Name:        "staging_area",
				Description: "The staging area of the source server.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.SourceServer.StagingArea")},
			{
				Name:        "staging_account_id",
				Description: "The staging account ID that extended source servers belong to.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.SourceServer.StagingArea.StagingAccountID")},
			{
				Name:        "launch_configuration",
				Description: "The launch configuration settings of the source server.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.LaunchConfiguration"),
			},
			{
				Name:        "hardware_id",
				Description: "An ID that describes the hardware of the Source Server. This is either an EC2 instance id, a VMware uuid or a mac address.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromQual("hardware_id"),
			},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.SourceServer.SourceServerID")},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.SourceServer.Tags"),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.SourceServer.Arn").Transform(arnToAkas),
			},
		}),
	}
}
