# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>name</td><td>The name of the instance fleet.</td></tr>
	<tr><td>id</td><td>The identifier of the instance fleet.</td></tr>
	<tr><td>arn</td><td>The Amazon Resource Name (ARN) specifying the instance fleet.</td></tr>
	<tr><td>cluster_id</td><td>The unique identifier for the cluster.</td></tr>
	<tr><td>instance_fleet_type</td><td>The type of the instance fleet. Valid values are MASTER, CORE or TASK.</td></tr>
	<tr><td>state</td><td>The current state of the instance fleet.</td></tr>
	<tr><td>provisioned_on_demand_capacity</td><td>The number of On-Demand units that have been provisioned for the instance fleet to fulfill TargetOnDemandCapacity.</td></tr>
	<tr><td>provisioned_spot_capacity</td><td>The number of Spot units that have been provisioned for this instance fleet to fulfill TargetSpotCapacity.</td></tr>
	<tr><td>target_on_demand_capacity</td><td>The target capacity of On-Demand units for the instance fleet, which determines how many On-Demand Instances to provision.</td></tr>
	<tr><td>target_spot_capacity</td><td>The target capacity of Spot units for the instance fleet, which determines how many Spot Instances to provision.</td></tr>
	<tr><td>instance_type_specifications</td><td>An array of specifications for the instance types that comprise an instance fleet.</td></tr>
	<tr><td>launch_specifications</td><td>Describes the launch specification for an instance fleet.</td></tr>
	<tr><td>state_change_reason</td><td>Provides status change reason details for the instance fleet.</td></tr>
	<tr><td>status_timeline</td><td>Provides historical timestamps for the instance fleet, including the time of creation, the time it became ready to run jobs, and the time of termination.</td></tr>
	<tr><td>title</td><td>Title of the resource.</td></tr>
	<tr><td>akas</td><td>Array of globally unique identifier strings (also known as) for the resource.</td></tr>
	<tr><td>partition</td><td>The AWS partition in which the resource is located (aws, aws-cn, or aws-us-gov).</td></tr>
	<tr><td>region</td><td>The AWS Region in which the resource is located.</td></tr>
	<tr><td>account_id</td><td>The AWS Account ID in which the resource is located.</td></tr>
	<tr><td>kaytu_account_id</td><td>The Kaytu Account ID in which the resource is located.</td></tr>
	<tr><td>kaytu_resource_id</td><td>The unique ID of the resource in Kaytu.</td></tr>
	<tr><td>kaytu_metadata</td><td>kaytu Metadata of the AWS resource.</td></tr>
</table>