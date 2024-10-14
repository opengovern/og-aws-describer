package aws

import (
	"context"
	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"
	"regexp"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/turbot/go-kit/helpers"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsEc2Instance(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_ec2_instance",
		Description: "AWS EC2 Instance",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("instance_id"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"InvalidInstanceID.NotFound", "InvalidInstanceID.Unavailable", "InvalidInstanceID.Malformed"}),
			},
			Hydrate: opengovernance.GetEC2Instance,
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListEC2Instance,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "hypervisor", Require: plugin.Optional},
				{Name: "iam_instance_profile_arn", Require: plugin.Optional},
				{Name: "image_id", Require: plugin.Optional},
				{Name: "instance_lifecycle", Require: plugin.Optional},
				{Name: "instance_state", Require: plugin.Optional},
				{Name: "instance_type", Require: plugin.Optional},
				{Name: "monitoring_state", Require: plugin.Optional},
				{Name: "outpost_arn", Require: plugin.Optional},
				{Name: "placement_availability_zone", Require: plugin.Optional},
				{Name: "placement_group_name", Require: plugin.Optional},
				{Name: "public_dns_name", Require: plugin.Optional},
				{Name: "ram_disk_id", Require: plugin.Optional},
				{Name: "root_device_name", Require: plugin.Optional},
				{Name: "root_device_type", Require: plugin.Optional},
				{Name: "subnet_id", Require: plugin.Optional},
				{Name: "placement_tenancy", Require: plugin.Optional},
				{Name: "virtualization_type", Require: plugin.Optional},
				{Name: "vpc_id", Require: plugin.Optional},
			},
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "instance_id",
				Description: "The ID of the instance.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Instance.InstanceId"),
			},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) specifying the instance.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("ARN"),
			},
			{
				Name:        "instance_type",
				Description: "The instance type.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Instance.InstanceType"),
			},
			{
				Name:        "instance_state",
				Description: "The state of the instance (pending | running | shutting-down | terminated | stopping | stopped).",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Instance.State.Name"),
			},
			{
				Name:        "monitoring_state",
				Description: "Indicates whether detailed monitoring is enabled (disabled | enabled).",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Instance.Monitoring.State"),
			},
			{
				Name:        "disable_api_termination",
				Default:     false,
				Description: "If the value is true, instance can't be terminated through the Amazon EC2 console, CLI, or API.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Attributes.DisableApiTermination"),
			},
			{
				Name:        "ami_launch_index",
				Description: "The AMI launch index, which can be used to find this instance in the launch group.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Instance.AmiLaunchIndex"),
			},
			{
				Name:        "architecture",
				Description: "The architecture of the image.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Instance.Architecture"),
			},
			{
				Name:        "boot_mode",
				Description: "The boot mode of the instance.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Instance.BootMode"),
			},
			{
				Name:        "capacity_reservation_id",
				Description: "The ID of the Capacity Reservation.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Instance.CapacityReservationId"),
			},
			{
				Name:        "capacity_reservation_specification",
				Description: "Information about the Capacity Reservation targeting option.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Instance.CapacityReservationSpecification"),
			},
			{
				Name:        "client_token",
				Description: "The idempotency token you provided when you launched the instance, if applicable.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Instance.ClientToken"),
			},
			{
				Name:        "cpu_options_core_count",
				Description: "The number of CPU cores for the instance.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Instance.CpuOptions.CoreCount"),
			},
			{
				Name:        "cpu_options_threads_per_core",
				Description: "The number of threads per CPU core.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Instance.CpuOptions.ThreadsPerCore"),
			},
			{
				Name:        "ebs_optimized",
				Description: "Indicates whether the instance is optimized for Amazon EBS I/O. This optimization provides dedicated throughput to Amazon EBS and an optimized configuration stack to provide optimal I/O performance. This optimization isn't available with all instance types.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Instance.EbsOptimized"),
			},
			{
				Name:        "ena_support",
				Description: "Specifies whether enhanced networking with ENA is enabled.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Instance.EnaSupport"),
			},
			{
				Name:        "hypervisor",
				Description: "The hypervisor type of the instance. The value xen is used for both Xen and Nitro hypervisors.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Instance.Hypervisor"),
			},
			{
				Name:        "iam_instance_profile_arn",
				Description: "The Amazon Resource Name (ARN) of IAM instance profile associated with the instance, if applicable.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Instance.IamInstanceProfile.Arn"),
			},
			{
				Name:        "iam_instance_profile_id",
				Description: "The ID of the instance profile associated with the instance, if applicable.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Instance.IamInstanceProfile.Id"),
			},
			{
				Name:        "image_id",
				Description: "The ID of the AMI used to launch the instance.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Instance.ImageId"),
			},
			{
				Name:        "instance_initiated_shutdown_behavior",
				Description: "Indicates whether an instance stops or terminates when you initiate shutdown from the instance (using the operating system command for system shutdown).",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Attributes.InstanceInitiatedShutdownBehavior"),
			},
			{
				Name:        "instance_lifecycle",
				Description: "Indicates whether this is a spot instance or a scheduled instance.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Instance.InstanceLifecycle"),
			},
			{
				Name:        "kernel_id",
				Description: "The kernel ID",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Instance.KernelId"),
			},
			{
				Name:        "key_name",
				Description: "The name of the key pair, if this instance was launched with an associated key pair.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Instance.KeyName"),
			},
			{
				Name:        "launch_time",
				Description: "The time the instance was launched.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Instance.LaunchTime"),
			},
			{
				Name:        "outpost_arn",
				Description: "The Amazon Resource Name (ARN) of the Outpost, if applicable.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Instance.OutpostArn"),
			},
			{
				Name:        "placement_affinity",
				Description: "The affinity setting for the instance on the Dedicated Host.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Instance.Placement.Affinity"),
			},
			{
				Name:        "placement_availability_zone",
				Description: "The Availability Zone of the instance.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Instance.Placement.AvailabilityZone"),
			},
			{
				Name:        "placement_group_id",
				Description: "The ID of the placement group that the instance is in.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Instance.Placement.GroupId"),
			},
			{
				Name:        "placement_group_name",
				Description: "The name of the placement group the instance is in.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Instance.Placement.GroupName"),
			},
			{
				Name:        "placement_host_id",
				Description: "The ID of the Dedicated Host on which the instance resides.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Instance.Placement.HostId"),
			},
			{
				Name:        "placement_host_resource_group_arn",
				Description: "The ARN of the host resource group in which to launch the instances.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Instance.Placement.HostResourceGroupArn"),
			},
			{
				Name:        "placement_partition_number",
				Description: "The ARN of the host resource group in which to launch the instances.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Instance.Placement.PartitionNumber"),
			},
			{
				Name:        "placement_tenancy",
				Description: "The tenancy of the instance (if the instance is running in a VPC). An instance with a tenancy of dedicated runs on single-tenant hardware.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Instance.Placement.Tenancy"),
			},
			{
				Name:        "platform",
				Description: "The value is 'Windows' for Windows instances; otherwise blank.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Instance.Platform"),
			},
			{
				Name:        "platform_details",
				Description: "The platform details value for the instance.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Instance.PlatformDetails"),
			},
			{
				Name:        "private_ip_address",
				Description: "The private IPv4 address assigned to the instance.",
				Type:        proto.ColumnType_IPADDR,
				Transform:   transform.FromField("Description.Instance.PrivateIpAddress"),
			},
			{
				Name:        "private_dns_name",
				Description: "The private DNS hostname name assigned to the instance. This DNS hostname can only be used inside the Amazon EC2 network. This name is not available until the instance enters the running state.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Instance.PrivateDnsName"),
			},
			{
				Name:        "public_dns_name",
				Description: "The public DNS name assigned to the instance. This name is not available until the instance enters the running state.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Instance.PublicDnsName"),
			},
			{
				Name:        "public_ip_address",
				Description: "The public IPv4 address, or the Carrier IP address assigned to the instance, if applicable.",
				Type:        proto.ColumnType_IPADDR,
				Transform:   transform.FromField("Description.Instance.PublicIpAddress"),
			},
			{
				Name:        "ram_disk_id",
				Description: "The RAM disk ID.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Instance.RamdiskId"),
			},
			{
				Name:        "root_device_name",
				Description: "The device name of the root device volume (for example, /dev/sda1).",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Instance.RootDeviceName"),
			},
			{
				Name:        "root_device_type",
				Description: "The root device type used by the AMI. The AMI can use an EBS volume or an instance store volume.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Instance.RootDeviceType"),
			},
			{
				Name:        "source_dest_check",
				Description: "Specifies whether to enable an instance launched in a VPC to perform NAT. This controls whether source/destination checking is enabled on the instance.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Instance.SourceDestCheck"),
			},
			{
				Name:        "spot_instance_request_id",
				Description: "If the request is a Spot Instance request, the ID of the request.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Instance.SpotInstanceRequestId"),
			},
			{
				Name:        "sriov_net_support",
				Description: "Indicates whether enhanced networking with the Intel 82599 Virtual Function interface is enabled.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Instance.SriovNetSupport"),
			},
			{
				Name:        "state_code",
				Description: "The reason code for the state change.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Instance.State.Code"),
			},
			{
				Name:        "state_transition_reason",
				Description: "The reason for the most recent state transition.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Instance.StateTransitionReason"),
			},
			{
				Name:        "state_transition_time",
				Description: "The date and time, the instance state was last modified.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.From(ec2InstanceStateChangeTime),
			},
			{
				Name:        "subnet_id",
				Description: "The ID of the subnet in which the instance is running.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Instance.SubnetId"),
			},
			{
				Name:        "tpm_support",
				Description: "If the instance is configured for NitroTPM support, the value is v2.0.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Instance.TpmSupport"),
			},
			{
				Name:        "usage_operation",
				Description: "The usage operation value for the instance.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Instance.UsageOperation"),
			},
			{
				Name:        "usage_operation_update_time",
				Description: "The time that the usage operation was last updated.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Instance.UsageOperationUpdateTime"),
			},
			{
				Name:        "user_data",
				Description: "The user data of the instance.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.LaunchTemplateData.UserData").Transform(base64DecodedData),
			},
			{
				Name:        "virtualization_type",
				Description: "The virtualization type of the instance.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Instance.VirtualizationType"),
			},
			{
				Name:        "vpc_id",
				Description: "The ID of the VPC in which the instance is running.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Instance.VpcId"),
			},
			{
				Name:        "block_device_mappings",
				Description: "Block device mapping entries for the instance.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Instance.BlockDeviceMappings"),
			},
			{
				Name:        "elastic_gpu_associations",
				Description: "The Elastic GPU associated with the instance.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Instance.ElasticGpuAssociations"),
			},
			{
				Name:        "elastic_inference_accelerator_associations",
				Description: "The elastic inference accelerator associated with the instance.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Instance.ElasticInferenceAcceleratorAssociations"),
			},
			{
				Name:        "enclave_options",
				Description: "Indicates whether the instance is enabled for Amazon Web Services Nitro Enclaves.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Instance.EnclaveOptions"),
			},
			{
				Name:        "hibernation_options",
				Description: "Indicates whether the instance is enabled for hibernation.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Instance.HibernationOptions"),
			},
			{
				Name:        "launch_template_data",
				Description: "The configuration data of the specified instance.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.LaunchTemplateData"),
			},
			{
				Name:        "licenses",
				Description: "The license configurations for the instance.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Instance.Licenses"),
			},
			{
				Name:        "maintenance_options",
				Description: "The metadata options for the instance.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Instance.MaintenanceOptions"),
			},
			{
				Name:        "metadata_options",
				Description: "The metadata options for the instance.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Instance.MetadataOptions"),
			},
			{
				Name:        "network_interfaces",
				Description: "The network interfaces for the instance.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Instance.NetworkInterfaces"),
			},
			{
				Name:        "private_dns_name_options",
				Description: "The options for the instance hostname.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Instance.PrivateDnsNameOptions"),
			},
			{
				Name:        "product_codes",
				Description: "The product codes attached to this instance, if applicable.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Instance.ProductCodes"),
			},
			{
				Name:        "security_groups",
				Description: "The security groups for the instance.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Instance.SecurityGroups"),
			},
			{
				Name:        "instance_status",
				Description: "The status of an instance. Instance status includes scheduled events, status checks and instance state information.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Attributes.InstanceStatus"),
			},
			{
				Name:        "tags_src",
				Description: "A list of tags assigned to the instance.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Instance.Tags"),
			},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.From(getEc2InstanceTurbotTitle),
			},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(getEc2InstanceTurbotTags),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("ARN").Transform(transform.EnsureStringArray),
			},
		}),
	}
}

//// TRANSFORM FUNCTIONS

func getEc2InstanceTurbotTags(_ context.Context, d *transform.TransformData) (interface{}, error) {
	desc := d.HydrateItem.(opengovernance.EC2Instance).Description
	instance := desc.Instance
	return ec2V2TagsToMap(instance.Tags)
}

func getEc2InstanceTurbotTitle(_ context.Context, d *transform.TransformData) (interface{}, error) {
	desc := d.HydrateItem.(opengovernance.EC2Instance).Description
	data := desc.Instance
	title := data.InstanceId
	if data.Tags != nil {
		for _, i := range data.Tags {
			if *i.Key == "Name" {
				title = i.Value
			}
		}
	}
	return title, nil
}

func ec2InstanceStateChangeTime(_ context.Context, d *transform.TransformData) (interface{}, error) {
	desc := d.HydrateItem.(opengovernance.EC2Instance).Description
	data := desc.Instance

	if *desc.Instance.StateTransitionReason != "" {
		if helpers.StringSliceContains([]string{"shutting-down", "stopped", "stopping", "terminated"}, string(data.State.Name)) {
			// User initiated (2019-09-12 16:38:34 GMT)
			regexExp := regexp.MustCompile(`\((.*?) *\)`)
			stateTransitionTime := regexExp.FindStringSubmatch(*data.StateTransitionReason)
			if len(stateTransitionTime) >= 1 {
				stateTransitionTimeInUTC := strings.Replace(strings.Replace(stateTransitionTime[1], " ", "T", 1), " GMT", "Z", 1)
				return stateTransitionTimeInUTC, nil
			}
		}
	}
	return data.LaunchTime, nil
}

//// UTILITY FUNCTIONS

func getListValuesString(listValue *proto.QualValueList) []string {
	values := make([]string, 0, len(listValue.Values))
	for _, value := range listValue.Values {
		values = append(values, value.GetStringValue())
	}
	return values
}

func getListValues(listValue *proto.QualValueList) []*string {
	values := make([]*string, 0)
	if listValue != nil {
		for _, value := range listValue.Values {
			if value.GetStringValue() != "" {
				values = append(values, aws.String(value.GetStringValue()))
			}
		}
	}
	return values
}
