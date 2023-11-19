<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>target_group_name</td><td>The name of the target group.</td></tr>
	<tr><td>target_group_arn</td><td>The Amazon Resource Name (ARN) of the target group.</td></tr>
	<tr><td>target_type</td><td>The type of target that is specified when registering targets with this target group. The possible values are instance (register targets by instance ID), ip (register targets by IP address), or lambda (register a single Lambda function as a target).</td></tr>
	<tr><td>load_balancer_arns</td><td>The Amazon Resource Names (ARN) of the load balancers that route traffic to this target group.</td></tr>
	<tr><td>port</td><td>The port on which the targets are listening. Not used if the target is a Lambda function.</td></tr>
	<tr><td>vpc_id</td><td>The ID of the VPC for the target.</td></tr>
	<tr><td>protocol</td><td>The protocol to use for routing traffic to the target.</td></tr>
	<tr><td>matcher_http_code</td><td>The HTTP codes to use when checking for a successful response from a target.</td></tr>
	<tr><td>matcher_grpc_code</td><td>The gRPC codes to use when checking for a successful response from a target.</td></tr>
	<tr><td>healthy_threshold_count</td><td>The number of consecutive health checks successes required before considering an unhealthy target healthy.</td></tr>
	<tr><td>unhealthy_threshold_count</td><td>The number of consecutive health checks successes required before considering an unhealthy target healthy.</td></tr>
	<tr><td>health_check_enabled</td><td>Indicates whether health checks are enabled.</td></tr>
	<tr><td>health_check_interval_seconds</td><td>The approximate amount of time, in seconds, between health checks of an individual target.</td></tr>
	<tr><td>health_check_path</td><td>The destination for health checks on the target.</td></tr>
	<tr><td>health_check_port</td><td>The port to use to connect with the target.</td></tr>
	<tr><td>health_check_protocol</td><td>The protocol to use to connect with the target. The GENEVE, TLS, UDP, and TCP_UDP protocols are not supported for health checks.</td></tr>
	<tr><td>health_check_timeout_seconds</td><td>The amount of time, in seconds, during which no response means a failed health check.</td></tr>
	<tr><td>target_health_descriptions</td><td>Contains information about the health of the target.</td></tr>
	<tr><td>tags_src</td><td>A list of tags associated with target group.</td></tr>
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