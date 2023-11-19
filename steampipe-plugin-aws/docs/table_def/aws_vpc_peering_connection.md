# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>id</td><td>The ID of the VPC peering connection.</td></tr>
	<tr><td>status_code</td><td>The status of the VPC peering connection. Possible values include: &#39;pending-acceptance&#39;, &#39;failed&#39;, &#39;expired&#39;, &#39;provisioning&#39;, &#39;active&#39;, &#39;deleting&#39;, &#39;deleted&#39; or &#39;rejected&#39;.</td></tr>
	<tr><td>accepter_cidr_block</td><td>The IPv4 CIDR block for the accepter VPC.</td></tr>
	<tr><td>accepter_owner_id</td><td>The ID of the Amazon Web Services account that owns the accepter VPC.</td></tr>
	<tr><td>accepter_region</td><td>The Region in which the accepter VPC is located.</td></tr>
	<tr><td>accepter_vpc_id</td><td>The ID of the accepter VPC.</td></tr>
	<tr><td>expiration_time</td><td>The time that an unaccepted VPC peering connection will expire.</td></tr>
	<tr><td>requester_cidr_block</td><td>The IPv4 CIDR block for the requester VPC.</td></tr>
	<tr><td>requester_owner_id</td><td>The ID of the Amazon Web Services account that owns the requester VPC.</td></tr>
	<tr><td>requester_region</td><td>The Region in which the requester VPC is located.</td></tr>
	<tr><td>requester_vpc_id</td><td>The ID of the requester VPC.</td></tr>
	<tr><td>status_message</td><td>A message that provides more information about the status, if applicable.</td></tr>
	<tr><td>accepter_cidr_block_set</td><td>Information about the IPv4 CIDR blocks for the accepter VPC.</td></tr>
	<tr><td>accepter_ipv6_cidr_block_set</td><td>The IPv6 CIDR block for the accepter VPC.</td></tr>
	<tr><td>accepter_peering_options</td><td>Information about the VPC peering connection options for the accepter VPC.</td></tr>
	<tr><td>requester_cidr_block_set</td><td>Information about the IPv4 CIDR blocks for the requester VPC.</td></tr>
	<tr><td>requester_ipv6_cidr_block_set</td><td>The IPv6 CIDR block for the requester VPC.</td></tr>
	<tr><td>requester_peering_options</td><td>Information about the VPC peering connection options for the requester VPC.</td></tr>
	<tr><td>tags_src</td><td>The tags assigned to the resource.</td></tr>
	<tr><td>title</td><td>Title of the resource.</td></tr>
	<tr><td>tags</td><td>A map of tags for the resource.</td></tr>
	<tr><td>partition</td><td>The AWS partition in which the resource is located (aws, aws-cn, or aws-us-gov).</td></tr>
	<tr><td>region</td><td>The AWS Region in which the resource is located.</td></tr>
	<tr><td>account_id</td><td>The AWS Account ID in which the resource is located.</td></tr>
	<tr><td>kaytu_account_id</td><td>The Kaytu Account ID in which the resource is located.</td></tr>
	<tr><td>kaytu_resource_id</td><td>The unique ID of the resource in Kaytu.</td></tr>
	<tr><td>kaytu_metadata</td><td>kaytu Metadata of the AWS resource.</td></tr>
</table>