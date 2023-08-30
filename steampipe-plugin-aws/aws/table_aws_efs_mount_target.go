package aws

import (
	"context"

	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsEfsMountTarget(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_efs_mount_target",
		Description: "AWS EFS Mount Target",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("mount_target_id"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"MountTargetNotFound", "InvalidParameter"}),
			},
			Hydrate: kaytu.GetEFSMountTarget,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListEFSMountTarget,
		},

		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "mount_target_id",
				Description: "The ID of the mount target.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.MountTarget.MountTargetId")},
			{
				Name:        "file_system_id",
				Description: "The ID of the file system for which the mount target is intended.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.MountTarget.FileSystemId")},
			{
				Name:        "life_cycle_state",
				Description: "Lifecycle state of the mount target.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.MountTarget.LifeCycleState")},
			{
				Name:        "availability_zone_id",
				Description: "The unique and consistent identifier of the Availability Zone that the mount target resides in.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.MountTarget.AvailabilityZoneId")},
			{
				Name:        "availability_zone_name",
				Description: "The name of the Availability Zone in which the mount target is located.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.MountTarget.AvailabilityZoneName")},
			{
				Name:        "ip_address",
				Description: "Address at which the file system can be mounted by using the mount target.",
				Type:        proto.ColumnType_IPADDR,
				Transform:   transform.FromField("Description.MountTarget.IpAddress")},
			{
				Name:        "network_interface_id",
				Description: "The ID of the network interface that Amazon EFS created when it created the mount target.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.MountTarget.NetworkInterfaceId")},
			{
				Name:        "owner_id",
				Description: "AWS account ID that owns the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.MountTarget.OwnerId")},
			{
				Name:        "subnet_id",
				Description: "The ID of the mount target's subnet.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.MountTarget.SubnetId")},
			{
				Name:        "vpc_id",
				Description: "The virtual private cloud (VPC) ID that the mount target is configured in.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.MountTarget.VpcId")},
			{
				Name:        "security_groups",
				Description: "Specifies the security groups currently in effect for a mount target.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.SecurityGroups")},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.MountTarget.MountTargetId")},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("ARN").Transform(transform.EnsureStringArray),
			},
		}),
	}
}
