# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>name</td><td>The name of the instance group.</td></tr>
	<tr><td>id</td><td>The identifier of the instance group.</td></tr>
	<tr><td>arn</td><td>The Amazon Resource Name (ARN) specifying the instance group.</td></tr>
	<tr><td>cluster_id</td><td>The unique identifier for the cluster.</td></tr>
	<tr><td>instance_group_type</td><td>The type of the instance group. Valid values are MASTER, CORE or TASK.</td></tr>
	<tr><td>instance_type</td><td>The EC2 instance type for all instances in the instance group.</td></tr>
	<tr><td>state</td><td>The current state of the instance group.</td></tr>
	<tr><td>bid_price</td><td>The maximum price you are willing to pay for Spot Instances. If specified, indicates that the instance group uses Spot Instances.</td></tr>
	<tr><td>configurations_version</td><td>The version number of the requested configuration specification for this instance group.</td></tr>
	<tr><td>ebs_optimized</td><td>Indicates whether the instance group is EBS-optimized, or not.  An Amazon EBS-optimized instance uses an optimized configuration stack and provides additional, dedicated capacity for Amazon EBS I/O.</td></tr>
	<tr><td>last_successfully_applied_configurations_version</td><td>The version number of a configuration specification that was successfully applied for an instance group last time.</td></tr>
	<tr><td>market</td><td>The marketplace to provision instances for this group. Valid values are ON_DEMAND or SPOT.</td></tr>
	<tr><td>requested_instance_count</td><td>The target number of instances for the instance group.</td></tr>
	<tr><td>running_instance_count</td><td>The number of instances currently running in this instance group.</td></tr>
	<tr><td>autoscaling_policy</td><td>An automatic scaling policy for a core instance group or task instance group in an Amazon EMR cluster.</td></tr>
	<tr><td>configurations</td><td>A list of configurations supplied for an EMR cluster instance group. Only availbale for Amazon EMR releases 4.x or later.</td></tr>
	<tr><td>ebs_block_devices</td><td>The EBS block devices that are mapped to this instance group.</td></tr>
	<tr><td>last_successfully_applied_configurations</td><td>A list of configurations that were successfully applied for an instance group last time.</td></tr>
	<tr><td>shrink_policy</td><td>Policy for customizing shrink operations.</td></tr>
	<tr><td>state_change_reason</td><td>The status change reason details for the instance group.</td></tr>
	<tr><td>status_timeline</td><td>The timeline of the instance group status over time.</td></tr>
	<tr><td>title</td><td>Title of the resource.</td></tr>
	<tr><td>akas</td><td>Array of globally unique identifier strings (also known as) for the resource.</td></tr>
	<tr><td>partition</td><td>The AWS partition in which the resource is located (aws, aws-cn, or aws-us-gov).</td></tr>
	<tr><td>region</td><td>The AWS Region in which the resource is located.</td></tr>
	<tr><td>account_id</td><td>The AWS Account ID in which the resource is located.</td></tr>
	<tr><td>kaytu_account_id</td><td>The Kaytu Account ID in which the resource is located.</td></tr>
	<tr><td>kaytu_resource_id</td><td>The unique ID of the resource in Kaytu.</td></tr>
	<tr><td>kaytu_metadata</td><td>kaytu Metadata of the AWS resource.</td></tr>
</table>