# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>verified_access_endpoint_id</td><td>The ID of the AWS verified access endpoint.</td></tr>
	<tr><td>verified_access_group_id</td><td>The ID of the AWS verified access group.</td></tr>
	<tr><td>verified_access_instance_id</td><td>The ID of the AWS verified access instance.</td></tr>
	<tr><td>creation_time</td><td>The creation time.</td></tr>
	<tr><td>status_code</td><td>The endpoint status code. Possible values are pending, active, updating, deleting or deleted.</td></tr>
	<tr><td>application_domain</td><td>The DNS name for users to reach your application.</td></tr>
	<tr><td>attachment_type</td><td>The type of attachment used to provide connectivity between the AWS verified access endpoint and the application.</td></tr>
	<tr><td>deletion_time</td><td>The deletion time.</td></tr>
	<tr><td>description</td><td>A description for the AWS verified access endpoint.</td></tr>
	<tr><td>device_validation_domain</td><td>Returned if endpoint has a device trust provider attached.</td></tr>
	<tr><td>domain_certificate_arn</td><td>The ARN of a public TLS/SSL certificate imported into or created with ACM.</td></tr>
	<tr><td>endpoint_domain</td><td>A DNS name that is generated for the endpoint..</td></tr>
	<tr><td>endpoint_type</td><td>The type of AWS verified access endpoint. Incoming application requests will be sent to an IP address, load balancer or a network interface depending on the endpoint type specified. Possible values are load-balancer or network-interface.</td></tr>
	<tr><td>last_updated_time</td><td>The last updated time.</td></tr>
	<tr><td>load_balancer_options</td><td>The load balancer details if creating the AWS verified access endpoint as load-balancertype.</td></tr>
	<tr><td>network_interface_options</td><td>The options for network-interface type endpoint.</td></tr>
	<tr><td>status</td><td>The endpoint status.</td></tr>
	<tr><td>tags_src</td><td>A map of tags for the resource.</td></tr>
	<tr><td>title</td><td>Title of the resource.</td></tr>
	<tr><td>tags</td><td>A map of tags for the resource.</td></tr>
	<tr><td>partition</td><td>The AWS partition in which the resource is located (aws, aws-cn, or aws-us-gov).</td></tr>
	<tr><td>region</td><td>The AWS Region in which the resource is located.</td></tr>
	<tr><td>account_id</td><td>The AWS Account ID in which the resource is located.</td></tr>
	<tr><td>kaytu_account_id</td><td>The Kaytu Account ID in which the resource is located.</td></tr>
	<tr><td>kaytu_resource_id</td><td>The unique ID of the resource in Kaytu.</td></tr>
	<tr><td>kaytu_metadata</td><td>kaytu Metadata of the AWS resource.</td></tr>
</table>