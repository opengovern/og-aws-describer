<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>name</td><td>The name of the Auto Scaling group.</td></tr>
	<tr><td>autoscaling_group_arn</td><td>The Amazon Resource Name (ARN) of the Auto Scaling group.</td></tr>
	<tr><td>status</td><td>The current state of the group when the DeleteAutoScalingGroup operation is in progress.</td></tr>
	<tr><td>created_time</td><td>The date and time group was created.</td></tr>
	<tr><td>new_instances_protected_from_scale_in</td><td>Indicates whether newly launched instances are protected from termination by Amazon EC2 Auto Scaling when scaling in.</td></tr>
	<tr><td>launch_configuration_name</td><td>The name of the associated launch configuration.</td></tr>
	<tr><td>default_cooldown</td><td>The duration of the default cooldown period, in seconds.</td></tr>
	<tr><td>desired_capacity</td><td>The desired size of the group.</td></tr>
	<tr><td>max_instance_lifetime</td><td>The maximum amount of time, in seconds, that an instance can be in service.</td></tr>
	<tr><td>max_size</td><td>The maximum size of the group.</td></tr>
	<tr><td>min_size</td><td>The minimum size of the group.</td></tr>
	<tr><td>health_check_grace_period</td><td>The amount of time, in seconds, that Amazon EC2 Auto Scaling waits before checking the health status of an EC2 instance that has come into service.</td></tr>
	<tr><td>health_check_type</td><td>The service to use for the health checks. The valid values are EC2 and ELB. If you configure an Auto Scaling group to use ELB health checks, it considers the instance unhealthy if it fails either the EC2 status checks or the load balancer health checks.</td></tr>
	<tr><td>placement_group</td><td>The name of the placement group into which to launch your instances, if any.</td></tr>
	<tr><td>service_linked_role_arn</td><td>The Amazon Resource Name (ARN) of the service-linked role that the Auto Scaling group uses to call other AWS services on your behalf.</td></tr>
	<tr><td>vpc_zone_identifier</td><td>One or more subnet IDs, if applicable, separated by commas.</td></tr>
	<tr><td>launch_template_name</td><td>The launch template name for the group.</td></tr>
	<tr><td>launch_template_id</td><td>The ID of the launch template.</td></tr>
	<tr><td>launch_template_version</td><td>The version number, $Latest, or $Default.</td></tr>
	<tr><td>on_demand_allocation_strategy</td><td>Indicates how to allocate instance types to fulfill On-Demand capacity. The only valid value is prioritized, which is also the default value. This strategy uses the order of instance types in the overrides to define the launch priority of each instance type.</td></tr>
	<tr><td>on_demand_base_capacity</td><td>The minimum amount of the Auto Scaling group's capacity that must be fulfilled by On-Demand Instances. This base portion is provisioned first as group scales. Defaults to 0 if not specified.</td></tr>
	<tr><td>on_demand_percentage_above_base_capacity</td><td>Controls the percentages of On-Demand Instances and Spot Instances for your additional capacity beyond OnDemandBaseCapacity. Expressed as a number (for example, 20 specifies 20% On-Demand Instances, 80% Spot Instances). Defaults to 100 if not specified. If set to 100, only On-Demand Instances are provisioned.</td></tr>
	<tr><td>spot_allocation_strategy</td><td>Indicates how to allocate instances across Spot Instance pools. If the allocation strategy is lowest-price, the Auto Scaling group launches instances using the Spot pools with the lowest price, and evenly allocates your instances across the number of Spot pools that you specify. If the allocation strategy is capacity-optimized, the Auto Scaling group launches instances using Spot pools that are optimally chosen based on the available Spot capacity. Defaults to lowest-price if not specified.</td></tr>
	<tr><td>spot_instance_pools</td><td>The number of Spot Instance pools across which to allocate your Spot Instances.</td></tr>
	<tr><td>spot_max_price</td><td>The maximum price per unit hour that user is willing to pay for a Spot Instance. If the value of this parameter is blank (which is the default), the maximum Spot price is set at the On-Demand price.</td></tr>
	<tr><td>mixed_instances_policy_launch_template_name</td><td>The ID of the launch template for mixed instances policy.</td></tr>
	<tr><td>mixed_instances_policy_launch_template_id</td><td>The name of the launch template for mixed instances policy.</td></tr>
	<tr><td>mixed_instances_policy_launch_template_version</td><td>The version of the launch template for mixed instances policy.</td></tr>
	<tr><td>mixed_instances_policy_launch_template_overrides</td><td>Any parameters that is specified in the list override the same parameters in the launch template.</td></tr>
	<tr><td>availability_zones</td><td>One or more Availability Zones for the group.</td></tr>
	<tr><td>load_balancer_names</td><td>One or more load balancers associated with the group.</td></tr>
	<tr><td>target_group_arns</td><td>The Amazon Resource Names (ARN) of the target groups for your load balancer.</td></tr>
	<tr><td>instances</td><td>The EC2 instances associated with the group.</td></tr>
	<tr><td>enabled_metrics</td><td>The metrics enabled for the group.</td></tr>
	<tr><td>policies</td><td>A set of scaling policies for the specified Auto Scaling group.</td></tr>
	<tr><td>termination_policies</td><td>The termination policies for the group.</td></tr>
	<tr><td>suspended_processes</td><td>The suspended processes associated with the group.</td></tr>
	<tr><td>tags_src</td><td>A list of tags assigned to the Auto Scaling Group.</td></tr>
	<tr><td>tags</td><td>A map of tags for the resource.</td></tr>
	<tr><td>title</td><td>Title of the resource.</td></tr>
	<tr><td>akas</td><td>Array of globally unique identifier strings (also known as) for the resource.</td></tr>
	<tr><td>partition</td><td>The AWS partition in which the resource is located (aws, aws-cn, or aws-us-gov).</td></tr>
	<tr><td>region</td><td>The AWS Region in which the resource is located.</td></tr>
	<tr><td>account_id</td><td>The AWS Account ID in which the resource is located.</td></tr>
	<tr><td>kaytu_account_id</td><td>The Kaytu Account ID in which the resource is located.</td></tr>
	<tr><td>kaytu_resource_id</td><td>The unique ID of the resource in Kaytu.</td></tr>
	<tr><td>kaytu_metadata</td><td>kaytu Metadata of the AWS resource.</td></tr>
</table>