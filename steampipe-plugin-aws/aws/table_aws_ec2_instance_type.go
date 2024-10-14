package aws

import (
	"context"

	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsInstanceType(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_ec2_instance_type",
		Description: "AWS EC2 Instance Type",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("instance_type"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"InvalidInstanceType"}),
			},
			Hydrate: opengovernance.GetEC2InstanceType,
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListEC2InstanceType,
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "instance_type",
				Description: "The instance type. For more information, see [ Instance Types ](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-types.html) in the Amazon Elastic Compute Cloud User Guide.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.InstanceType.InstanceType")},
			{
				Name:        "auto_recovery_supported",
				Description: "Indicates whether auto recovery is supported.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.InstanceType.AutoRecoverySupported")},
			{
				Name:        "bare_metal",
				Description: "Indicates whether the instance is a bare metal instance type.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.InstanceType.BareMetal")},
			{
				Name:        "burstable_performance_supported",
				Description: "Indicates whether the instance type is a burstable performance instance type.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.InstanceType.BurstablePerformanceSupported")},
			{
				Name:        "current_generation",
				Description: "Indicates whether the instance type is current generation.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.InstanceType.CurrentGeneration")},
			{
				Name:        "dedicated_hosts_supported",
				Description: "Indicates whether Dedicated Hosts are supported on the instance type.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.InstanceType.DedicatedHostsSupported")},
			{
				Name:        "free_tier_eligible",
				Description: "Indicates whether the instance type is eligible for the free tier.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.InstanceType.FreeTierEligible")},
			{
				Name:        "hibernation_supported",
				Description: "Indicates whether On-Demand hibernation is supported.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.InstanceType.HibernationSupported")},
			{
				Name:        "hypervisor",
				Description: "The hypervisor for the instance type.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.InstanceType.Hypervisor")},
			{
				Name:        "instance_storage_supported",
				Description: "Describes the instance storage for the instance type.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.InstanceType.InstanceStorageSupported")},
			{
				Name:        "ebs_info",
				Description: "Describes the Amazon EBS settings for the instance type.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.InstanceType.EbsInfo")},
			{
				Name:        "memory_info",
				Description: "Describes the memory for the instance type.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.InstanceType.MemoryInfo")},
			{
				Name:        "network_info",
				Description: "Describes the network settings for the instance type.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.InstanceType.NetworkInfo")},
			{
				Name:        "placement_group_info",
				Description: "Describes the placement group settings for the instance type.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.InstanceType.PlacementGroupInfo")},
			{
				Name:        "processor_info",
				Description: "Describes the processor.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.InstanceType.ProcessorInfo")},
			{
				Name:        "supported_root_device_types",
				Description: "The supported root device types.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.InstanceType.SupportedRootDeviceTypes")},
			{
				Name:        "supported_usage_classes",
				Description: "Indicates whether the instance type is offered for spot or On-Demand.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.InstanceType.SupportedUsageClasses")},
			{
				Name:        "supported_virtualization_types",
				Description: "The supported virtualization types.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.InstanceType.SupportedVirtualizationTypes")},
			{
				Name:        "v_cpu_info",
				Description: "Describes the vCPU configurations for the instance type.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.InstanceType.VCpuInfo")},
			{
				Name:        "gpu_info",
				Description: "Describes the GPU accelerator settings for the instance type.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.InstanceType.GpuInfo")},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.InstanceType.InstanceType")},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("ARN").Transform(arnToAkas)},
		}),
	}
}
