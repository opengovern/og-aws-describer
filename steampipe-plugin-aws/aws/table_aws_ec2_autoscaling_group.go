package aws

import (
	"context"
	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAwsEc2ASG(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_ec2_autoscaling_group",
		Description: "AWS EC2 Autoscaling Group",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("name"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"ValidationError"}),
			},
			Hydrate: kaytu.GetAutoScalingGroup,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListAutoScalingGroup,
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the Auto Scaling group.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AutoScalingGroup.AutoScalingGroupName"),
			},
			{
				Name:        "autoscaling_group_arn",
				Description: "The Amazon Resource Name (ARN) of the Auto Scaling group.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AutoScalingGroup.AutoScalingGroupARN"),
			},
			{
				Name:        "status",
				Description: "The current state of the group when the DeleteAutoScalingGroup operation is in progress.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AutoScalingGroup.Status"),
			},
			{
				Name:        "created_time",
				Description: "The date and time group was created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.AutoScalingGroup.CreatedTime"),
			},
			{
				Name:        "new_instances_protected_from_scale_in",
				Description: "Indicates whether newly launched instances are protected from termination by Amazon EC2 Auto Scaling when scaling in.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.AutoScalingGroup.NewInstancesProtectedFromScaleIn"),
			},
			{
				Name:        "launch_configuration_name",
				Description: "The name of the associated launch configuration.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AutoScalingGroup.LaunchConfigurationName"),
			},
			{
				Name:        "default_cooldown",
				Description: "The duration of the default cooldown period, in seconds.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.AutoScalingGroup.DefaultCooldown"),
			},
			{
				Name:        "desired_capacity",
				Description: "The desired size of the group.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.AutoScalingGroup.DesiredCapacity"),
			},
			{
				Name:        "max_instance_lifetime",
				Description: "The maximum amount of time, in seconds, that an instance can be in service.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.AutoScalingGroup.MaxInstanceLifetime"),
			},
			{
				Name:        "max_size",
				Description: "The maximum size of the group.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.AutoScalingGroup.MaxSize"),
			},
			{
				Name:        "min_size",
				Description: "The minimum size of the group.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.AutoScalingGroup.MinSize"),
			},
			{
				Name:        "health_check_grace_period",
				Description: "The amount of time, in seconds, that Amazon EC2 Auto Scaling waits before checking the health status of an EC2 instance that has come into service.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.AutoScalingGroup.HealthCheckGracePeriod"),
			},
			{
				Name:        "health_check_type",
				Description: "The service to use for the health checks. The valid values are EC2 and ELB. If you configure an Auto Scaling group to use ELB health checks, it considers the instance unhealthy if it fails either the EC2 status checks or the load balancer health checks.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AutoScalingGroup.HealthCheckType"),
			},
			{
				Name:        "placement_group",
				Description: "The name of the placement group into which to launch your instances, if any.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AutoScalingGroup.PlacementGroup"),
			},
			{
				Name:        "service_linked_role_arn",
				Description: "The Amazon Resource Name (ARN) of the service-linked role that the Auto Scaling group uses to call other AWS services on your behalf.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AutoScalingGroup.ServiceLinkedRoleARN"),
			},
			{
				Name:        "vpc_zone_identifier",
				Description: "One or more subnet IDs, if applicable, separated by commas.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AutoScalingGroup.VPCZoneIdentifier"),
			},
			{
				Name:        "launch_template_name",
				Description: "The launch template name for the group.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AutoScalingGroup.LaunchTemplate.LaunchTemplateName"),
			},
			{
				Name:        "launch_template_id",
				Description: "The ID of the launch template.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AutoScalingGroup.LaunchTemplate.LaunchTemplateId"),
			},
			{
				Name:        "launch_template_version",
				Description: "The version number, $Latest, or $Default.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AutoScalingGroup.LaunchTemplate.Version"),
			},
			{
				Name:        "on_demand_allocation_strategy",
				Description: "Indicates how to allocate instance types to fulfill On-Demand capacity. The only valid value is prioritized, which is also the default value. This strategy uses the order of instance types in the overrides to define the launch priority of each instance type.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AutoScalingGroup.MixedInstancesPolicy.InstancesDistribution.OnDemandAllocationStrategy"),
			},
			{
				Name:        "on_demand_base_capacity",
				Description: "The minimum amount of the Auto Scaling group's capacity that must be fulfilled by On-Demand Instances. This base portion is provisioned first as group scales. Defaults to 0 if not specified.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.AutoScalingGroup.MixedInstancesPolicy.InstancesDistribution.OnDemandBaseCapacity"),
			},
			{
				Name:        "on_demand_percentage_above_base_capacity",
				Description: "Controls the percentages of On-Demand Instances and Spot Instances for your additional capacity beyond OnDemandBaseCapacity. Expressed as a number (for example, 20 specifies 20% On-Demand Instances, 80% Spot Instances). Defaults to 100 if not specified. If set to 100, only On-Demand Instances are provisioned.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.AutoScalingGroup.MixedInstancesPolicy.InstancesDistribution.OnDemandPercentageAboveBaseCapacity"),
			},
			{
				Name:        "spot_allocation_strategy",
				Description: "Indicates how to allocate instances across Spot Instance pools. If the allocation strategy is lowest-price, the Auto Scaling group launches instances using the Spot pools with the lowest price, and evenly allocates your instances across the number of Spot pools that you specify. If the allocation strategy is capacity-optimized, the Auto Scaling group launches instances using Spot pools that are optimally chosen based on the available Spot capacity. Defaults to lowest-price if not specified.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AutoScalingGroup.MixedInstancesPolicy.InstancesDistribution.SpotAllocationStrategy"),
			},
			{
				Name:        "spot_instance_pools",
				Description: "The number of Spot Instance pools across which to allocate your Spot Instances.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.AutoScalingGroup.MixedInstancesPolicy.InstancesDistribution.SpotInstancePools"),
			},
			{
				Name:        "spot_max_price",
				Description: "The maximum price per unit hour that user is willing to pay for a Spot Instance. If the value of this parameter is blank (which is the default), the maximum Spot price is set at the On-Demand price.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AutoScalingGroup.MixedInstancesPolicy.InstancesDistribution.SpotMaxPrice"),
			},
			{
				Name:        "mixed_instances_policy_launch_template_name",
				Description: "The ID of the launch template for mixed instances policy.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AutoScalingGroup.MixedInstancesPolicy.LaunchTemplate.LaunchTemplateSpecification.LaunchTemplateName"),
			},
			{
				Name:        "mixed_instances_policy_launch_template_id",
				Description: "The name of the launch template for mixed instances policy.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AutoScalingGroup.MixedInstancesPolicy.LaunchTemplate.LaunchTemplateSpecification.LaunchTemplateId"),
			},
			{
				Name:        "mixed_instances_policy_launch_template_version",
				Description: "The version of the launch template for mixed instances policy.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AutoScalingGroup.MixedInstancesPolicy.LaunchTemplate.LaunchTemplateSpecification.Version"),
			},
			{
				Name:        "mixed_instances_policy_launch_template_overrides",
				Description: "Any parameters that is specified in the list override the same parameters in the launch template.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.AutoScalingGroup.MixedInstancesPolicy.LaunchTemplate.Overrides"),
			},
			{
				Name:        "availability_zones",
				Description: "One or more Availability Zones for the group.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.AutoScalingGroup.AvailabilityZones"),
			},
			{
				Name:        "load_balancer_names",
				Description: "One or more load balancers associated with the group.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.AutoScalingGroup.LoadBalancerNames"),
			},
			{
				Name:        "target_group_arns",
				Description: "The Amazon Resource Names (ARN) of the target groups for your load balancer.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.AutoScalingGroup.TargetGroupARNs"),
			},
			{
				Name:        "instances",
				Description: "The EC2 instances associated with the group.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.AutoScalingGroup.Instances"),
			},
			{
				Name:        "enabled_metrics",
				Description: "The metrics enabled for the group.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.AutoScalingGroup.EnabledMetrics"),
			},
			{
				Name:        "policies",
				Description: "A set of scaling policies for the specified Auto Scaling group.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Policies"),
			},
			{
				Name:        "termination_policies",
				Description: "The termination policies for the group.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.AutoScalingGroup.TerminationPolicies"),
			},
			{
				Name:        "suspended_processes",
				Description: "The suspended processes associated with the group.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.AutoScalingGroup.SuspendedProcesses"),
			},
			{
				Name:        "tags_src",
				Description: "A list of tags assigned to the Auto Scaling Group.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.AutoScalingGroup.Tags"),
			},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(getASGTurbotTags),
			},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AutoScalingGroup.AutoScalingGroupName"),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.AutoScalingGroup.AutoScalingGroupARN").Transform(arnToAkas),
			},
		}),
	}
}

//// LIST FUNCTION

//// HYDRATE FUNCTIONS

//// UTILITY FUNCTIONS

//// TRANSFORM FUNCTIONS

func getASGTurbotTags(_ context.Context, d *transform.TransformData) (interface{}, error) {
	asg := d.HydrateItem.(kaytu.AutoScalingGroup).Description.AutoScalingGroup
	var turbotTagsMap map[string]string
	if asg.Tags == nil {
		return nil, nil
	}

	turbotTagsMap = map[string]string{}
	for _, i := range asg.Tags {
		turbotTagsMap[*i.Key] = *i.Value
	}

	return &turbotTagsMap, nil
}
