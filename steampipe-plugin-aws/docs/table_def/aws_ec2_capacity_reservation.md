# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>capacity_reservation_id</td><td>The ID of the capacity reservation.</td></tr>
	<tr><td>capacity_reservation_arn</td><td>The Amazon Resource Name (ARN) of the capacity reservation.</td></tr>
	<tr><td>instance_type</td><td>The type of instance for which the capacity reservation reserves capacity.</td></tr>
	<tr><td>state</td><td>The current state of the capacity reservation. A capacity reservation can be in one of the following states: &#39;active&#39;, &#39;expired&#39;, &#39;cancelled&#39;, &#39;pending&#39;, &#39;failed&#39;.</td></tr>
	<tr><td>availability_zone</td><td>The availability zone in which the capacity is reserved.</td></tr>
	<tr><td>availability_zone_id</td><td>The availability zone ID of the capacity reservation.</td></tr>
	<tr><td>available_instance_count</td><td>The remaining capacity. Indicates the number of instances that can be launched in the capacity reservation.</td></tr>
	<tr><td>create_date</td><td>The date and time at which the capacity reservation was created.</td></tr>
	<tr><td>ebs_optimized</td><td>Indicates whether the capacity reservation supports EBS-optimized instances.</td></tr>
	<tr><td>end_date</td><td>The date and time at which the capacity reservation expires.</td></tr>
	<tr><td>end_date_type</td><td>Indicates the way in which the capacity reservation ends. A capacity reservation can have one of the following end types: &#39;unlimited&#39;, &#39;limited&#39;.</td></tr>
	<tr><td>ephemeral_storage</td><td>Indicates whether the capacity reservation supports instances with temporary, block-level storage.</td></tr>
	<tr><td>instance_match_criteria</td><td>Indicates the type of instance launches that the capacity reservation accepts. The options include: &#39;open&#39;, &#39;targeted&#39;.</td></tr>
	<tr><td>instance_platform</td><td>The type of operating system for which the capacity reservation reserves capacity.</td></tr>
	<tr><td>owner_id</td><td>The ID of the AWS account that owns the capacity reservation.</td></tr>
	<tr><td>start_date</td><td>The date and time at which the capacity reservation was started.</td></tr>
	<tr><td>tenancy</td><td>Indicates the tenancy of the capacity reservation. A capacity reservation can have one of the following tenancy settings: &#39;default&#39;, &#39;dedicated&#39;.</td></tr>
	<tr><td>total_instance_count</td><td>The total number of instances for which the capacity reservation reserves capacity.</td></tr>
	<tr><td>capacity_allocations</td><td>Information about instance capacity usage.</td></tr>
	<tr><td>tag_src</td><td>Any tags assigned to the capacity reservation.</td></tr>
	<tr><td>title</td><td>Title of the resource.</td></tr>
	<tr><td>tags</td><td>A map of tags for the resource.</td></tr>
	<tr><td>akas</td><td>Array of globally unique identifier strings (also known as) for the resource.</td></tr>
	<tr><td>partition</td><td>The AWS partition in which the resource is located (aws, aws-cn, or aws-us-gov).</td></tr>
	<tr><td>region</td><td>The AWS Region in which the resource is located.</td></tr>
	<tr><td>account_id</td><td>The AWS Account ID in which the resource is located.</td></tr>
	<tr><td>kaytu_account_id</td><td>The Kaytu Account ID in which the resource is located.</td></tr>
	<tr><td>kaytu_resource_id</td><td>The unique ID of the resource in Kaytu.</td></tr>
	<tr><td>kaytu_metadata</td><td>kaytu Metadata of the AWS resource.</td></tr>
</table>