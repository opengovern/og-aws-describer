# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>subnet_id</td><td>Contains the unique ID to specify a subnet.</td></tr>
	<tr><td>subnet_arn</td><td>Contains the Amazon Resource Name (ARN) of the subnet.</td></tr>
	<tr><td>vpc_id</td><td>ID of the VPC, the subnet is in.</td></tr>
	<tr><td>cidr_block</td><td>Contains the IPv4 CIDR block assigned to the subnet.</td></tr>
	<tr><td>state</td><td>Current state of the subnet.</td></tr>
	<tr><td>owner_id</td><td>Contains the AWS account that own the subnet.</td></tr>
	<tr><td>assign_ipv6_address_on_creation</td><td>Indicates whether a network interface created in this subnet (including a network interface created by RunInstances) receives an IPv6 address.</td></tr>
	<tr><td>available_ip_address_count</td><td>The number of unused private IPv4 addresses in the subnet. The IPv4 addresses for any stopped instances are considered unavailable.</td></tr>
	<tr><td>availability_zone</td><td>The Availability Zone of the subnet.</td></tr>
	<tr><td>availability_zone_id</td><td>The AZ ID of the subnet.</td></tr>
	<tr><td>customer_owned_ipv4_pool</td><td>The customer-owned IPv4 address pool associated with the subnet.</td></tr>
	<tr><td>default_for_az</td><td>Indicates whether this is the default subnet for the Availability Zone.</td></tr>
	<tr><td>map_customer_owned_ip_on_launch</td><td>Indicates whether a network interface created in this subnet (including a network interface created by RunInstances) receives a customer-owned IPv4 address.</td></tr>
	<tr><td>map_public_ip_on_launch</td><td>Indicates whether instances launched in this subnet receive a public IPv4 address.</td></tr>
	<tr><td>outpost_arn</td><td>The Amazon Resource Name (ARN) of the Outpost. Available only if subnet is on an outpost.</td></tr>
	<tr><td>ipv6_cidr_block_association_set</td><td>A list of IPv6 CIDR blocks associated with the subnet.</td></tr>
	<tr><td>tags_src</td><td>A list of tags that are attached to the subnet.</td></tr>
	<tr><td>tags</td><td>A map of tags for the resource.</td></tr>
	<tr><td>title</td><td>Title of the resource.</td></tr>
	<tr><td>akas</td><td>Array of globally unique identifier strings (also known as) for the resource.</td></tr>
	<tr><td>partition</td><td>The AWS partition in which the resource is located (aws, aws-cn, or aws-us-gov).</td></tr>
	<tr><td>region</td><td>The AWS Region in which the resource is located.</td></tr>
	<tr><td>account_id</td><td>The AWS Account ID in which the resource is located.</td></tr>
	<tr><td>kaytu_account_id</td><td>The Kaytu Account ID in which the resource is located.</td></tr>
	<tr><td>kaytu_resource_id</td><td>The unique ID of the resource in opengovernance.</td></tr>
	<tr><td>kaytu_metadata</td><td>kaytu Metadata of the AWS resource.</td></tr>
</table>