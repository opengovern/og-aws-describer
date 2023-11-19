# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>service_name</td><td>The Amazon Resource Name (ARN) of the service.</td></tr>
	<tr><td>service_id</td><td>The ID of the endpoint service.</td></tr>
	<tr><td>owner</td><td>The AWS account ID of the service owner.</td></tr>
	<tr><td>acceptance_required</td><td>Indicates whether VPC endpoint connection requests to the service must be accepted by the service owner.</td></tr>
	<tr><td>manages_vpc_endpoints</td><td>Indicates whether the service manages its VPC endpoints. Management of the service VPC endpoints using the VPC endpoint API is restricted.</td></tr>
	<tr><td>private_dns_name</td><td>The private DNS name for the service.</td></tr>
	<tr><td>private_dns_name_verification_state</td><td>The verification state of the VPC endpoint service. Consumers of the endpoint service cannot use the private name when the state is not verified.</td></tr>
	<tr><td>vpc_endpoint_policy_supported</td><td>Indicates whether the service supports endpoint policies.</td></tr>
	<tr><td>availability_zones</td><td>The Availability Zones in which the service is available.</td></tr>
	<tr><td>base_endpoint_dns_names</td><td>The DNS names for the service.</td></tr>
	<tr><td>service_type</td><td>The type of service.</td></tr>
	<tr><td>vpc_endpoint_connections</td><td>Information about one or more VPC endpoint connections.</td></tr>
	<tr><td>vpc_endpoint_service_permissions</td><td>Information about one or more allowed principals.</td></tr>
	<tr><td>tags_src</td><td>A list of tags assigned to the service.</td></tr>
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