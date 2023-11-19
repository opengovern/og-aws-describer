# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>name</td><td>The name of the service.</td></tr>
	<tr><td>id</td><td>The ID that Cloud Map assigned to the service when you created it.</td></tr>
	<tr><td>arn</td><td>The Amazon Resource Name (ARN) that Cloud Map assigns to the service when you create it.</td></tr>
	<tr><td>create_date</td><td>The date and time that the service was created.</td></tr>
	<tr><td>description</td><td>A description for the service.</td></tr>
	<tr><td>namespace_id</td><td>The ID of the namespace.</td></tr>
	<tr><td>instance_count</td><td>The number of instances that are currently associated with the service.</td></tr>
	<tr><td>type</td><td>Describes the systems that can be used to discover the service instances. DNS_HTTP The service instances can be discovered using either DNS queries or the DiscoverInstances API operation. HTTP The service instances can only be discovered using the DiscoverInstances API operation. DNS Reserved.</td></tr>
	<tr><td>routing_policy</td><td>The routing policy that you want to apply to all Route 53 DNS records that Cloud Map creates when you register an instance and specify this service.</td></tr>
	<tr><td>dns_records</td><td>An array that contains one DnsRecord object for each Route 53 DNS record that you want Cloud Map to create when you register an instance.</td></tr>
	<tr><td>health_check_config</td><td>Public DNS and HTTP namespaces only. Settings for an optional health check. If you specify settings for a health check, Cloud Map associates the health check with the records that you specify in DnsConfig.</td></tr>
	<tr><td>health_check_custom_config</td><td>Information about an optional custom health check.</td></tr>
	<tr><td>tags_src</td><td>Information about the tags associated with the namespace.</td></tr>
	<tr><td>title</td><td>Title of the resource.</td></tr>
	<tr><td>tags</td><td>A map of tags for the resource.</td></tr>
	<tr><td>akas</td><td>Array of globally unique identifier strings (also known as) for the resource.</td></tr>
	<tr><td>partition</td><td>The AWS partition in which the resource is located (aws, aws-cn, or aws-us-gov).</td></tr>
	<tr><td>region</td><td>The AWS Region in which the resource is located.</td></tr>
	<tr><td>account_id</td><td>The AWS Account ID in which the resource is located.</td></tr>
</table>