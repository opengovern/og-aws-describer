# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>arn</td><td>The Amazon Resource Name (ARN) of the endpoint group.</td></tr>
	<tr><td>listener_arn</td><td>The Amazon Resource Name (ARN) of parent listener.</td></tr>
	<tr><td>endpoint_descriptions</td><td>The list of endpoint objects.</td></tr>
	<tr><td>endpoint_group_region</td><td>The AWS Region where the endpoint group is located.</td></tr>
	<tr><td>health_check_interval_seconds</td><td>The time—10 seconds or 30 seconds—between health checks for each endpoint.</td></tr>
	<tr><td>health_check_path</td><td>If the protocol is HTTP/S, then this value provides the ping path that Global Accelerator uses for the destination on the endpoints for health checks.</td></tr>
	<tr><td>health_check_port</td><td>The port that Global Accelerator uses to perform health checks on endpoints that are part of this endpoint group.</td></tr>
	<tr><td>health_check_protocol</td><td>The protocol that Global Accelerator uses to perform health checks on endpoints that are part of this endpoint group.</td></tr>
	<tr><td>port_overrides</td><td>Overrides for destination ports used to route traffic to an endpoint.</td></tr>
	<tr><td>threshold_count</td><td>The number of consecutive health checks required to set the state of a healthy endpoint to unhealthy, or to set an unhealthy endpoint to healthy.</td></tr>
	<tr><td>traffic_dial_percentage</td><td>The percentage of traffic to send to an AWS Region.</td></tr>
	<tr><td>title</td><td>Title of the resource.</td></tr>
	<tr><td>akas</td><td>Array of globally unique identifier strings (also known as) for the resource.</td></tr>
	<tr><td>partition</td><td>The AWS partition in which the resource is located (aws, aws-cn, or aws-us-gov).</td></tr>
	<tr><td>region</td><td>The AWS Region in which the resource is located.</td></tr>
	<tr><td>account_id</td><td>The AWS Account ID in which the resource is located.</td></tr>
	<tr><td>kaytu_account_id</td><td>The Kaytu Account ID in which the resource is located.</td></tr>
	<tr><td>kaytu_resource_id</td><td>The unique ID of the resource in opengovernance.</td></tr>
	<tr><td>kaytu_metadata</td><td>Kaytu Metadata of the AWS resource.</td></tr>
</table>