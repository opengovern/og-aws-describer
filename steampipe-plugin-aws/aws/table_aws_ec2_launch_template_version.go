package aws

import (
	"context"
	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsEc2LaunchTemplateVersion(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_ec2_launch_template_version",
		Description: "AWS EC2 Launch Template Version",
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListEC2LaunchTemplateVersion,
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:    "launch_template_id",
					Require: plugin.Optional,
				},
				{
					Name:    "launch_template_name",
					Require: plugin.Optional,
				},
				{
					Name:    "default_version",
					Require: plugin.Optional,
				},
				{
					Name:    "ebs_optimized",
					Require: plugin.Optional,
				},
				{
					Name:    "image_id",
					Require: plugin.Optional,
				},
				{
					Name:    "instance_type",
					Require: plugin.Optional,
				},
				{
					Name:    "kernel_id",
					Require: plugin.Optional,
				},
				{
					Name:    "ram_disk_id",
					Require: plugin.Optional,
				},
			},
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "launch_template_name",
				Description: "The name of the launch template.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.LaunchTemplateVersion.LaunchTemplateName"),
			},
			{
				Name:        "launch_template_id",
				Description: "The ID of the launch template.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.LaunchTemplateVersion.LaunchTemplateId"),
			},
			{
				Name:        "create_time",
				Description: "The time the version was created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.LaunchTemplateVersion.CreateTime"),
			},
			{
				Name:        "created_by",
				Description: "The principal that created the version.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.LaunchTemplateVersion.CreatedBy"),
			},
			{
				Name:        "default_version",
				Description: "Indicates whether the version is the default version.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.LaunchTemplateVersion.DefaultVersion"),
			},
			{
				Name:        "disable_api_stop",
				Description: "Indicates whether the instance is enabled for stop protection.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.LaunchTemplateVersion.LaunchTemplateData.DisableApiStop"),
			},
			{
				Name:        "disable_api_termination",
				Description: "If set to true, indicates that the instance cannot be terminated using the Amazon EC2 console, command line tool, or API.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.LaunchTemplateVersion.LaunchTemplateData.DisableApiTermination"),
			},
			{
				Name:        "ebs_optimized",
				Description: "Indicates whether the instance is optimized for Amazon EBS I/O.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.LaunchTemplateVersion.LaunchTemplateData.EbsOptimized"),
			},
			{
				Name:        "image_id",
				Description: "The ID of the AMI or a Systems Manager parameter.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.LaunchTemplateVersion.LaunchTemplateData.ImageId"),
			},
			{
				Name:        "instance_type",
				Description: "The instance type.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.LaunchTemplateVersion.LaunchTemplateData.InstanceType"),
			},
			{
				Name:        "kernel_id",
				Description: "The ID of the kernel, if applicable.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.LaunchTemplateVersion.LaunchTemplateData.KernelId"),
			},
			{
				Name:        "key_name",
				Description: "The name of the key pair.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.LaunchTemplateVersion.LaunchTemplateData.KeyName"),
			},
			{
				Name:        "ram_disk_id",
				Description: "The ID of the RAM disk, if applicable.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.LaunchTemplateVersion.LaunchTemplateData.RamDiskId"),
			},
			{
				Name:        "security_groups",
				Description: "The security group names.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.LaunchTemplateVersion.LaunchTemplateData.SecurityGroups"),
			},
			{
				Name:        "security_group_ids",
				Description: "The security group IDs.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.LaunchTemplateVersion.LaunchTemplateData.SecurityGroupIds"),
			},
			{
				Name:        "version_description",
				Description: "The description for the version.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.LaunchTemplateVersion.VersionDescription"),
			},
			{
				Name:        "version_number",
				Description: "The version number.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.LaunchTemplateVersion.VersionNumber"),
			},
			{
				Name:        "user_data",
				Description: "The user data of the launch template.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.LaunchTemplateVersion.LaunchTemplateData.UserData").Transform(base64DecodedData),
			},
			{
				Name:        "launch_template_data",
				Description: "Information about the launch template.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.LaunchTemplateVersion.LaunchTemplateData"),
			},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.LaunchTemplateVersion.LaunchTemplateName"),
			},
		}),
	}
}
