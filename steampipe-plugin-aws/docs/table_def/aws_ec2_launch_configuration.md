# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>name</td><td>The name of the launch configuration.</td></tr>
	<tr><td>launch_configuration_arn</td><td>The Amazon Resource Name (ARN) of the launch configuration.</td></tr>
	<tr><td>created_time</td><td>The creation date and time for the launch configuration.</td></tr>
	<tr><td>image_id</td><td>The ID of the Amazon Machine Image (AMI) to use to launch EC2 instances.</td></tr>
	<tr><td>instance_type</td><td>The instance type for the instances.</td></tr>
	<tr><td>associate_public_ip_address</td><td>For Auto Scaling groups that are running in a VPC, specifies whether to assign a public IP address to the group's instances.</td></tr>
	<tr><td>kernel_id</td><td>The ID of the kernel associated with the AMI.</td></tr>
	<tr><td>key_name</td><td>The name of the key pair to be associated with instances.</td></tr>
	<tr><td>ramdisk_id</td><td>The ID of the RAM disk associated with the AMI.</td></tr>
	<tr><td>ebs_optimized</td><td>Specifies whether the launch configuration is optimized for EBS I/O (true) or not (false).</td></tr>
	<tr><td>classic_link_vpc_id</td><td>The ID of a ClassicLink-enabled VPC to link EC2-Classic instances to.</td></tr>
	<tr><td>spot_price</td><td>The maximum hourly price to be paid for any Spot Instance launched to fulfill the request. Spot Instances are launched when the price you specified exceeds the current Spot price.</td></tr>
	<tr><td>user_data</td><td>The Base64-encoded user data to make available to the launched EC2 instances.</td></tr>
	<tr><td>placement_tenancy</td><td>The tenancy of the instance, either default or dedicated. An instance with dedicated tenancy runs on isolated, single-tenant hardware and can only be launched into a VPC.</td></tr>
	<tr><td>iam_instance_profile</td><td>The name or the Amazon Resource Name (ARN) of the instance profile associated with the IAM role for the instance.</td></tr>
	<tr><td>instance_monitoring_enabled</td><td>Describes whether detailed monitoring is enabled for the Auto Scaling instances.</td></tr>
	<tr><td>metadata_options_http_endpoint</td><td>This parameter enables or disables the HTTP metadata endpoint on instances. If the parameter is not specified, the default state is enabled.</td></tr>
	<tr><td>metadata_options_put_response_hop_limit</td><td>The desired HTTP PUT response hop limit for instance metadata requests. The larger the number, the further instance metadata requests can travel.</td></tr>
	<tr><td>metadata_options_http_tokens</td><td>The state of token usage for your instance metadata requests. If the parameter is not specified in the request, the default state is optional.</td></tr>
	<tr><td>block_device_mappings</td><td>A block device mapping, which specifies the block devices for the instance.</td></tr>
	<tr><td>classic_link_vpc_security_groups</td><td>The IDs of one or more security groups for the VPC specified in ClassicLinkVPCId.</td></tr>
	<tr><td>security_groups</td><td>A list that contains the security groups to assign to the instances in the Auto Scaling group.</td></tr>
	<tr><td>title</td><td>Title of the resource.</td></tr>
	<tr><td>akas</td><td>Array of globally unique identifier strings (also known as) for the resource.</td></tr>
	<tr><td>partition</td><td>The AWS partition in which the resource is located (aws, aws-cn, or aws-us-gov).</td></tr>
	<tr><td>region</td><td>The AWS Region in which the resource is located.</td></tr>
	<tr><td>account_id</td><td>The AWS Account ID in which the resource is located.</td></tr>
	<tr><td>kaytu_account_id</td><td>The Kaytu Account ID in which the resource is located.</td></tr>
	<tr><td>kaytu_resource_id</td><td>The unique ID of the resource in Kaytu.</td></tr>
	<tr><td>kaytu_metadata</td><td>kaytu Metadata of the AWS resource.</td></tr>
</table>