# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>name</td><td>The name that you assigned to the Resolver endpoint when you submitted a CreateResolverEndpoint.</td></tr>
	<tr><td>id</td><td>The ID of the Resolver endpoint.</td></tr>
	<tr><td>arn</td><td>The ARN (Amazon Resource Name) for the Resolver endpoint.</td></tr>
	<tr><td>creation_time</td><td>The date and time that the endpoint was created, in Unix time format and Coordinated Universal Time (UTC).</td></tr>
	<tr><td>creator_request_id</td><td>A unique string that identifies the request that created the Resolver endpoint.The CreatorRequestId allows failed requests to be retried without the risk of executing the operation twice.</td></tr>
	<tr><td>direction</td><td>Indicates whether the Resolver endpoint allows inbound or outbound DNS queries.</td></tr>
	<tr><td>host_vpc_id</td><td>The ID of the VPC that you want to create the Resolver endpoint in.</td></tr>
	<tr><td>ip_address_count</td><td>The number of IP addresses that the Resolver endpoint can use for DNS queries.</td></tr>
	<tr><td>modification_time</td><td>The date and time that the endpoint was last modified, in Unix time format and Coordinated Universal Time (UTC).</td></tr>
	<tr><td>status</td><td>A code that specifies the current status of the Resolver endpoint.</td></tr>
	<tr><td>status_message</td><td>A detailed description of the status of the Resolver endpoint.</td></tr>
	<tr><td>ip_addresses</td><td>Information about the IP addresses in your VPC that DNS queries originate from (for outbound endpoints) or that you forward DNS queries to (for inbound endpoints).</td></tr>
	<tr><td>security_group_ids</td><td>The ID of one or more security groups that control access to this VPC.</td></tr>
	<tr><td>tags_src</td><td>A list of tags assigned to the Resolver endpoint.</td></tr>
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