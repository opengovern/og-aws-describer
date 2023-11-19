# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>allocation_id</td><td>Contains the ID representing the allocation of the address for use with EC2-VPC.</td></tr>
	<tr><td>arn</td><td>The Amazon Resource Name (ARN) specifying the VPC EIP.</td></tr>
	<tr><td>public_ip</td><td>Contains the Elastic IP address.</td></tr>
	<tr><td>public_ipv4_pool</td><td>The ID of an address pool.</td></tr>
	<tr><td>domain</td><td>Indicates whether Elastic IP address is for use with instances in EC2-Classic(standard) or instances in a VPC (vpc).</td></tr>
	<tr><td>association_id</td><td>Contains the ID representing the association of the address with an instance in a VPC.</td></tr>
	<tr><td>carrier_ip</td><td>The carrier IP address associated. This option is only available for network interfaces which reside in a subnet in a Wavelength Zone (for example an EC2 instance).</td></tr>
	<tr><td>customer_owned_ip</td><td>The customer-owned IP address.</td></tr>
	<tr><td>customer_owned_ipv4_pool</td><td>The ID of the customer-owned address pool.</td></tr>
	<tr><td>instance_id</td><td>Contains the ID of the instance that the address is associated with.</td></tr>
	<tr><td>network_border_group</td><td>The name of the unique set of Availability Zones, Local Zones, or Wavelength Zones from which AWS advertises IP addresses.</td></tr>
	<tr><td>network_interface_id</td><td>The ID of the network interface.</td></tr>
	<tr><td>network_interface_owner_id</td><td>The ID of the AWS account that owns the network interface.</td></tr>
	<tr><td>private_ip_address</td><td>The private IP address associated with the Elastic IP address.</td></tr>
	<tr><td>tags_src</td><td>A list of tags that are attached to the vpc.</td></tr>
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