package aws

import (
	"context"

	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAwsEc2LaunchConfiguration(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_ec2_launch_configuration",
		Description: "AWS EC2 Launch Configuration",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("name"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"ResourceNotFoundException", "ValidationError"}),
			},
			Hydrate: opengovernance.GetAutoScalingLaunchConfiguration,
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListAutoScalingLaunchConfiguration,
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the launch configuration.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.LaunchConfiguration.LaunchConfigurationName")},
			{
				Name:        "launch_configuration_arn",
				Description: "The Amazon Resource Name (ARN) of the launch configuration.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.LaunchConfiguration.LaunchConfigurationARN")},
			{
				Name:        "created_time",
				Description: "The creation date and time for the launch configuration.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.LaunchConfiguration.CreatedTime")},
			{
				Name:        "image_id",
				Description: "The ID of the Amazon Machine Image (AMI) to use to launch EC2 instances.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.LaunchConfiguration.ImageId")},
			{
				Name:        "instance_type",
				Description: "The instance type for the instances.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.LaunchConfiguration.InstanceType")},
			{
				Name:        "associate_public_ip_address",
				Description: "For Auto Scaling groups that are running in a VPC, specifies whether to assign a public IP address to the group's instances.",
				Type:        proto.ColumnType_BOOL,
				Default:     false,
				Transform:   transform.FromField("Description.LaunchConfiguration.AssociatePublicIpAddress")},
			{
				Name:        "kernel_id",
				Description: "The ID of the kernel associated with the AMI.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.LaunchConfiguration.KernelId")},
			{
				Name:        "key_name",
				Description: "The name of the key pair to be associated with instances.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.LaunchConfiguration.KeyName")},
			{
				Name:        "ramdisk_id",
				Description: "The ID of the RAM disk associated with the AMI.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.LaunchConfiguration.RamdiskId"),
			},
			{
				Name:        "ebs_optimized",
				Description: "Specifies whether the launch configuration is optimized for EBS I/O (true) or not (false).",
				Type:        proto.ColumnType_BOOL,
				Default:     false,
				Transform:   transform.FromField("Description.LaunchConfiguration.EbsOptimized")},
			{
				Name:        "classic_link_vpc_id",
				Description: "The ID of a ClassicLink-enabled VPC to link EC2-Classic instances to.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.LaunchConfiguration.ClassicLinkVPCId")},
			{
				Name:        "spot_price",
				Description: "The maximum hourly price to be paid for any Spot Instance launched to fulfill the request. Spot Instances are launched when the price you specified exceeds the current Spot price.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.LaunchConfiguration.SpotPrice")},
			{
				Name:        "user_data",
				Description: "The Base64-encoded user data to make available to the launched EC2 instances.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.LaunchConfiguration.UserData").Transform(base64DecodedData),
			},
			{
				Name:        "placement_tenancy",
				Description: "The tenancy of the instance, either default or dedicated. An instance with dedicated tenancy runs on isolated, single-tenant hardware and can only be launched into a VPC.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.LaunchConfiguration.PlacementTenancy")},
			{
				Name:        "iam_instance_profile",
				Description: "The name or the Amazon Resource Name (ARN) of the instance profile associated with the IAM role for the instance.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.LaunchConfiguration.IamInstanceProfile")},
			{
				Name:        "instance_monitoring_enabled",
				Description: "Describes whether detailed monitoring is enabled for the Auto Scaling instances.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.LaunchConfiguration.InstanceMonitoring.Enabled"), Default: false,
			},
			{
				Name:        "metadata_options_http_endpoint",
				Description: "This parameter enables or disables the HTTP metadata endpoint on instances. If the parameter is not specified, the default state is enabled.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.LaunchConfiguration.MetadataOptions.HttpEndpoint")},
			{
				Name:        "metadata_options_put_response_hop_limit",
				Description: "The desired HTTP PUT response hop limit for instance metadata requests. The larger the number, the further instance metadata requests can travel.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.LaunchConfiguration.MetadataOptions.HttpPutResponseHopLimit")},
			{
				Name:        "metadata_options_http_tokens",
				Description: "The state of token usage for your instance metadata requests. If the parameter is not specified in the request, the default state is optional.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.LaunchConfiguration.MetadataOptions.HttpTokens")},
			{
				Name:        "block_device_mappings",
				Description: "A block device mapping, which specifies the block devices for the instance.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.LaunchConfiguration.BlockDeviceMappings")},
			{
				Name:        "classic_link_vpc_security_groups",
				Description: "The IDs of one or more security groups for the VPC specified in ClassicLinkVPCId.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.LaunchConfiguration.ClassicLinkVPCSecurityGroups")},
			{
				Name:        "security_groups",
				Description: "A list that contains the security groups to assign to the instances in the Auto Scaling group.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.LaunchConfiguration.SecurityGroups"),
			},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.LaunchConfiguration.LaunchConfigurationName")},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.LaunchConfiguration.LaunchConfigurationARN").Transform(arnToAkas),
			},
		}),
	}
}

//// LIST FUNCTION

//// HYDRATE FUNCTIONS
