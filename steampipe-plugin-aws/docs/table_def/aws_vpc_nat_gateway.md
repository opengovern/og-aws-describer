# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>nat_gateway_id</td><td>The ID of the NAT gateway.</td></tr>
	<tr><td>arn</td><td>The Amazon Resource Name (ARN) specifying the NAT gateway.</td></tr>
	<tr><td>nat_gateway_addresses</td><td>Information about the IP addresses and network interface associated with the NAT gateway.</td></tr>
	<tr><td>state</td><td>The current state of the NAT gateway (pending | failed | available | deleting | deleted).</td></tr>
	<tr><td>create_time</td><td>The date and time the NAT gateway was created.</td></tr>
	<tr><td>vpc_id</td><td>The ID of the VPC in which the NAT gateway is located.</td></tr>
	<tr><td>subnet_id</td><td>The ID of the subnet in which the NAT gateway is located.</td></tr>
	<tr><td>delete_time</td><td>The date and time the NAT gateway was deleted, if applicable.</td></tr>
	<tr><td>failure_code</td><td>If the NAT gateway could not be created, specifies the error code for the failure. (InsufficientFreeAddressesInSubnet | Gateway.NotAttached | InvalidAllocationID.NotFound | Resource.AlreadyAssociated | InternalError | InvalidSubnetID.NotFound).</td></tr>
	<tr><td>failure_message</td><td>If the NAT gateway could not be created, specifies the error message for the failure.</td></tr>
	<tr><td>provisioned_bandwidth</td><td>Reserved. If you need to sustain traffic greater than the documented limits (https://docs.aws.amazon.com/vpc/latest/userguide/vpc-nat-gateway.html).</td></tr>
	<tr><td>tags_src</td><td>A list of tags that are attached to NAT gateway.</td></tr>
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