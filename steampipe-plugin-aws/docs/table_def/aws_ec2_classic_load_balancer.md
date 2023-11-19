# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>name</td><td>The friendly name of the Load Balancer.</td></tr>
	<tr><td>arn</td><td>The Amazon Resource Name (ARN) specifying the classic load balancer.</td></tr>
	<tr><td>scheme</td><td>The load balancing scheme of load balancer.</td></tr>
	<tr><td>created_time</td><td>The date and time the load balancer was created.</td></tr>
	<tr><td>vpc_id</td><td>The ID of the VPC for the load balancer.</td></tr>
	<tr><td>access_log_emit_interval</td><td>The interval for publishing the access logs.</td></tr>
	<tr><td>access_log_enabled</td><td>Specifies whether access logs are enabled for the load balancer.</td></tr>
	<tr><td>access_log_s3_bucket_name</td><td>The name of the Amazon S3 bucket where the access logs are stored.</td></tr>
	<tr><td>access_log_s3_bucket_prefix</td><td>The logical hierarchy you created for your Amazon S3 bucket.</td></tr>
	<tr><td>canonical_hosted_zone_name</td><td>The name of the Amazon Route 53 hosted zone for the load balancer.</td></tr>
	<tr><td>canonical_hosted_zone_name_id</td><td>The ID of the Amazon Route 53 hosted zone for the load balancer.</td></tr>
	<tr><td>connection_draining_enabled</td><td>Specifies whether connection draining is enabled for the load balancer.</td></tr>
	<tr><td>connection_draining_timeout</td><td>The maximum time, in seconds, to keep the existing connections open before deregistering the instances.</td></tr>
	<tr><td>connection_settings_idle_timeout</td><td>The time, in seconds, that the connection is allowed to be idle (no data has been sent over the connection) before it is closed by the load balancer.</td></tr>
	<tr><td>cross_zone_load_balancing_enabled</td><td>Specifies whether cross-zone load balancing is enabled for the load balancer.</td></tr>
	<tr><td>dns_name</td><td>The DNS name of the load balancer.</td></tr>
	<tr><td>health_check_interval</td><td>The approximate interval, in seconds, between health checks of an individual instance.</td></tr>
	<tr><td>health_check_timeout</td><td>The amount of time, in seconds, during which no response means a failed health check.</td></tr>
	<tr><td>healthy_threshold</td><td>The number of consecutive health checks successes required before moving the instance to the Healthy state.</td></tr>
	<tr><td>health_check_target</td><td>The instance being checked. The protocol is either TCP, HTTP, HTTPS, or SSL. The range of valid ports is one (1) through 65535.</td></tr>
	<tr><td>source_security_group_name</td><td>The name of the security group.</td></tr>
	<tr><td>source_security_group_owner_alias</td><td>The owner of the security group.</td></tr>
	<tr><td>unhealthy_threshold</td><td>The number of consecutive health check failures required before moving the instance to the Unhealthy state.</td></tr>
	<tr><td>additional_attributes</td><td>A list of additional attributes.</td></tr>
	<tr><td>app_cookie_stickiness_policies</td><td>A list of the stickiness policies created using CreateAppCookieStickinessPolicy.</td></tr>
	<tr><td>availability_zones</td><td>A list of the Availability Zones for the load balancer.</td></tr>
	<tr><td>backend_server_descriptions</td><td>A list of information about your EC2 instances.</td></tr>
	<tr><td>instances</td><td>A list of the IDs of the instances for the load balancer.</td></tr>
	<tr><td>lb_cookie_stickiness_policies</td><td>A list of the stickiness policies created using CreateLBCookieStickinessPolicy.</td></tr>
	<tr><td>listener_descriptions</td><td>A list of the listeners for the load balancer</td></tr>
	<tr><td>other_policies</td><td>A list of policies other than the stickiness policies.</td></tr>
	<tr><td>security_groups</td><td>A list of the security groups for the load balancer.</td></tr>
	<tr><td>subnets</td><td>A list of the IDs of the subnets for the load balancer.</td></tr>
	<tr><td>tags_src</td><td>A list of tags attached to the load balancer.</td></tr>
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