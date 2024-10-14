# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>instance_id</td><td>The ID of the instance.</td></tr>
	<tr><td>arn</td><td>The Amazon Resource Name (ARN) specifying the instance.</td></tr>
	<tr><td>instance_type</td><td>The instance type.</td></tr>
	<tr><td>instance_state</td><td>The state of the instance (pending | running | shutting-down | terminated | stopping | stopped).</td></tr>
	<tr><td>monitoring_state</td><td>Indicates whether detailed monitoring is enabled (disabled | enabled).</td></tr>
	<tr><td>disable_api_termination</td><td>If the value is true, instance can&#39;t be terminated through the Amazon EC2 console, CLI, or API.</td></tr>
	<tr><td>ami_launch_index</td><td>The AMI launch index, which can be used to find this instance in the launch group.</td></tr>
	<tr><td>architecture</td><td>The architecture of the image.</td></tr>
	<tr><td>boot_mode</td><td>The boot mode of the instance.</td></tr>
	<tr><td>capacity_reservation_id</td><td>The ID of the Capacity Reservation.</td></tr>
	<tr><td>capacity_reservation_specification</td><td>Information about the Capacity Reservation targeting option.</td></tr>
	<tr><td>client_token</td><td>The idempotency token you provided when you launched the instance, if applicable.</td></tr>
	<tr><td>cpu_options_core_count</td><td>The number of CPU cores for the instance.</td></tr>
	<tr><td>cpu_options_threads_per_core</td><td>The number of threads per CPU core.</td></tr>
	<tr><td>ebs_optimized</td><td>Indicates whether the instance is optimized for Amazon EBS I/O. This optimization provides dedicated throughput to Amazon EBS and an optimized configuration stack to provide optimal I/O performance. This optimization isn&#39;t available with all instance types.</td></tr>
	<tr><td>ena_support</td><td>Specifies whether enhanced networking with ENA is enabled.</td></tr>
	<tr><td>hypervisor</td><td>The hypervisor type of the instance. The value xen is used for both Xen and Nitro hypervisors.</td></tr>
	<tr><td>iam_instance_profile_arn</td><td>The Amazon Resource Name (ARN) of IAM instance profile associated with the instance, if applicable.</td></tr>
	<tr><td>iam_instance_profile_id</td><td>The ID of the instance profile associated with the instance, if applicable.</td></tr>
	<tr><td>image_id</td><td>The ID of the AMI used to launch the instance.</td></tr>
	<tr><td>instance_initiated_shutdown_behavior</td><td>Indicates whether an instance stops or terminates when you initiate shutdown from the instance (using the operating system command for system shutdown).</td></tr>
	<tr><td>instance_lifecycle</td><td>Indicates whether this is a spot instance or a scheduled instance.</td></tr>
	<tr><td>kernel_id</td><td>The kernel ID</td></tr>
	<tr><td>key_name</td><td>The name of the key pair, if this instance was launched with an associated key pair.</td></tr>
	<tr><td>launch_time</td><td>The time the instance was launched.</td></tr>
	<tr><td>outpost_arn</td><td>The Amazon Resource Name (ARN) of the Outpost, if applicable.</td></tr>
	<tr><td>placement_affinity</td><td>The affinity setting for the instance on the Dedicated Host.</td></tr>
	<tr><td>placement_availability_zone</td><td>The Availability Zone of the instance.</td></tr>
	<tr><td>placement_group_id</td><td>The ID of the placement group that the instance is in.</td></tr>
	<tr><td>placement_group_name</td><td>The name of the placement group the instance is in.</td></tr>
	<tr><td>placement_host_id</td><td>The ID of the Dedicated Host on which the instance resides.</td></tr>
	<tr><td>placement_host_resource_group_arn</td><td>The ARN of the host resource group in which to launch the instances.</td></tr>
	<tr><td>placement_partition_number</td><td>The ARN of the host resource group in which to launch the instances.</td></tr>
	<tr><td>placement_tenancy</td><td>The tenancy of the instance (if the instance is running in a VPC). An instance with a tenancy of dedicated runs on single-tenant hardware.</td></tr>
	<tr><td>platform</td><td>The value is &#39;Windows&#39; for Windows instances; otherwise blank.</td></tr>
	<tr><td>platform_details</td><td>The platform details value for the instance.</td></tr>
	<tr><td>private_ip_address</td><td>The private IPv4 address assigned to the instance.</td></tr>
	<tr><td>private_dns_name</td><td>The private DNS hostname name assigned to the instance. This DNS hostname can only be used inside the Amazon EC2 network. This name is not available until the instance enters the running state.</td></tr>
	<tr><td>public_dns_name</td><td>The public DNS name assigned to the instance. This name is not available until the instance enters the running state.</td></tr>
	<tr><td>public_ip_address</td><td>The public IPv4 address, or the Carrier IP address assigned to the instance, if applicable.</td></tr>
	<tr><td>ram_disk_id</td><td>The RAM disk ID.</td></tr>
	<tr><td>root_device_name</td><td>The device name of the root device volume (for example, /dev/sda1).</td></tr>
	<tr><td>root_device_type</td><td>The root device type used by the AMI. The AMI can use an EBS volume or an instance store volume.</td></tr>
	<tr><td>source_dest_check</td><td>Specifies whether to enable an instance launched in a VPC to perform NAT. This controls whether source/destination checking is enabled on the instance.</td></tr>
	<tr><td>spot_instance_request_id</td><td>If the request is a Spot Instance request, the ID of the request.</td></tr>
	<tr><td>sriov_net_support</td><td>Indicates whether enhanced networking with the Intel 82599 Virtual Function interface is enabled.</td></tr>
	<tr><td>state_code</td><td>The reason code for the state change.</td></tr>
	<tr><td>state_transition_reason</td><td>The reason for the most recent state transition.</td></tr>
	<tr><td>state_transition_time</td><td>The date and time, the instance state was last modified.</td></tr>
	<tr><td>subnet_id</td><td>The ID of the subnet in which the instance is running.</td></tr>
	<tr><td>tpm_support</td><td>If the instance is configured for NitroTPM support, the value is v2.0.</td></tr>
	<tr><td>usage_operation</td><td>The usage operation value for the instance.</td></tr>
	<tr><td>usage_operation_update_time</td><td>The time that the usage operation was last updated.</td></tr>
	<tr><td>user_data</td><td>The user data of the instance.</td></tr>
	<tr><td>virtualization_type</td><td>The virtualization type of the instance.</td></tr>
	<tr><td>vpc_id</td><td>The ID of the VPC in which the instance is running.</td></tr>
	<tr><td>block_device_mappings</td><td>Block device mapping entries for the instance.</td></tr>
	<tr><td>elastic_gpu_associations</td><td>The Elastic GPU associated with the instance.</td></tr>
	<tr><td>elastic_inference_accelerator_associations</td><td>The elastic inference accelerator associated with the instance.</td></tr>
	<tr><td>enclave_options</td><td>Indicates whether the instance is enabled for Amazon Web Services Nitro Enclaves.</td></tr>
	<tr><td>hibernation_options</td><td>Indicates whether the instance is enabled for hibernation.</td></tr>
	<tr><td>launch_template_data</td><td>The configuration data of the specified instance.</td></tr>
	<tr><td>licenses</td><td>The license configurations for the instance.</td></tr>
	<tr><td>maintenance_options</td><td>The metadata options for the instance.</td></tr>
	<tr><td>metadata_options</td><td>The metadata options for the instance.</td></tr>
	<tr><td>network_interfaces</td><td>The network interfaces for the instance.</td></tr>
	<tr><td>private_dns_name_options</td><td>The options for the instance hostname.</td></tr>
	<tr><td>product_codes</td><td>The product codes attached to this instance, if applicable.</td></tr>
	<tr><td>security_groups</td><td>The security groups for the instance.</td></tr>
	<tr><td>instance_status</td><td>The status of an instance. Instance status includes scheduled events, status checks and instance state information.</td></tr>
	<tr><td>tags_src</td><td>A list of tags assigned to the instance.</td></tr>
	<tr><td>title</td><td>Title of the resource.</td></tr>
	<tr><td>tags</td><td>A map of tags for the resource.</td></tr>
	<tr><td>akas</td><td>Array of globally unique identifier strings (also known as) for the resource.</td></tr>
	<tr><td>partition</td><td>The AWS partition in which the resource is located (aws, aws-cn, or aws-us-gov).</td></tr>
	<tr><td>region</td><td>The AWS Region in which the resource is located.</td></tr>
	<tr><td>account_id</td><td>The AWS Account ID in which the resource is located.</td></tr>
	<tr><td>kaytu_account_id</td><td>The Kaytu Account ID in which the resource is located.</td></tr>
	<tr><td>kaytu_resource_id</td><td>The unique ID of the resource in opengovernance.</td></tr>
	<tr><td>kaytu_metadata</td><td>kaytu Metadata of the AWS resource.</td></tr>
</table>