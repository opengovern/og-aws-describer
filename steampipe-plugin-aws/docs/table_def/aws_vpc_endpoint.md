# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>vpc_endpoint_id</td><td>The ID of the VPC endpoint.</td></tr>
	<tr><td>service_name</td><td>The name of the service to which the endpoint is associated.</td></tr>
	<tr><td>owner_id</td><td>The ID of the AWS account that owns the VPC endpoint.</td></tr>
	<tr><td>vpc_id</td><td>The ID of the VPC to which the endpoint is associated.</td></tr>
	<tr><td>state</td><td>The state of the VPC endpoint.</td></tr>
	<tr><td>private_dns_enabled</td><td>Indicates whether the VPC is associated with a private hosted zone.</td></tr>
	<tr><td>requester_managed</td><td>Indicates whether the VPC endpoint is being managed by its service.</td></tr>
	<tr><td>policy</td><td>The policy document associated with the endpoint, if applicable.</td></tr>
	<tr><td>policy_std</td><td>Contains the policy in a canonical form for easier searching.</td></tr>
	<tr><td>subnet_ids</td><td>One or more subnets in which the endpoint is located.</td></tr>
	<tr><td>route_table_ids</td><td>One or more route tables associated with the endpoint.</td></tr>
	<tr><td>groups</td><td>Information about the security groups that are associated with the network interface.</td></tr>
	<tr><td>network_interface_ids</td><td>One or more network interfaces for the endpoint.</td></tr>
	<tr><td>dns_entries</td><td>The DNS entries for the endpoint.</td></tr>
	<tr><td>creation_timestamp</td><td>The date and time that the VPC endpoint was created.</td></tr>
	<tr><td>tags_src</td><td>A list of tags assigned to the VPC endpoint.</td></tr>
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