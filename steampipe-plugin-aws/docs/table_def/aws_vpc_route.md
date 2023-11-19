# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>route_table_id</td><td>The ID of the route table containing the route.</td></tr>
	<tr><td>state</td><td>The state of the route. The blackhole state indicates that the route&#39;s target isn&#39;t available (for example, the specified gateway isn&#39;t attached to the VPC, or the specified NAT instance has been terminated).</td></tr>
	<tr><td>destination_cidr_block</td><td>The IPv4 CIDR block used for the destination match.</td></tr>
	<tr><td>destination_ipv6_cidr_block</td><td>The IPv6 CIDR block used for the destination match.</td></tr>
	<tr><td>carrier_gateway_id</td><td>The ID of the carrier gateway.</td></tr>
	<tr><td>destination_prefix_list_id</td><td>The prefix of the AWS service.</td></tr>
	<tr><td>egress_only_internet_gateway_id</td><td>The ID of the egress-only internet gateway.</td></tr>
	<tr><td>gateway_id</td><td>The ID of a gateway attached to your VPC.</td></tr>
	<tr><td>instance_id</td><td>The ID of a NAT instance in your VPC.</td></tr>
	<tr><td>instance_owner_id</td><td>The AWS account ID of the owner of the instance.</td></tr>
	<tr><td>local_gateway_id</td><td>The ID of the local gateway.</td></tr>
	<tr><td>nat_gateway_id</td><td>The ID of a NAT gateway.</td></tr>
	<tr><td>network_interface_id</td><td>The ID of the network interface.</td></tr>
	<tr><td>transit_gateway_id</td><td>The ID of a transit gateway.</td></tr>
	<tr><td>vpc_peering_connection_id</td><td>The ID of a VPC peering connection.</td></tr>
	<tr><td>origin</td><td>Describes how the route was created. CreateRouteTable - The route was automatically created when the route table was created. CreateRoute - The route was manually added to the route table. EnableVgwRoutePropagation - The route was propagated by route propagation.</td></tr>
	<tr><td>title</td><td>Title of the resource.</td></tr>
	<tr><td>akas</td><td>Array of globally unique identifier strings (also known as) for the resource.</td></tr>
	<tr><td>partition</td><td>The AWS partition in which the resource is located (aws, aws-cn, or aws-us-gov).</td></tr>
	<tr><td>region</td><td>The AWS Region in which the resource is located.</td></tr>
	<tr><td>account_id</td><td>The AWS Account ID in which the resource is located.</td></tr>
	<tr><td>kaytu_account_id</td><td>The Kaytu Account ID in which the resource is located.</td></tr>
	<tr><td>kaytu_resource_id</td><td>The unique ID of the resource in Kaytu.</td></tr>
	<tr><td>kaytu_metadata</td><td>kaytu Metadata of the AWS resource.</td></tr>
</table>