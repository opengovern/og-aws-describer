# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>id</td><td>The ID of the instance.</td></tr>
	<tr><td>service_id</td><td>The ID of the service.</td></tr>
	<tr><td>ec2_instance_id</td><td>The Amazon EC2 instance ID for the instance. When the AWS_EC2_INSTANCE_ID attribute is specified, then the AWS_INSTANCE_IPV4 attribute contains the primary private IPv4 address.</td></tr>
	<tr><td>alias_dns_name</td><td>For an alias record that routes traffic to an Elastic Load Balancing load balancer, the DNS name that&#39;s associated with the load balancer.</td></tr>
	<tr><td>instance_cname</td><td>A CNAME record, the domain name that Route 53 returns in response to DNS queries (for example, example.com ).</td></tr>
	<tr><td>init_health_status</td><td>If the service configuration includes HealthCheckCustomConfig, you can optionally use AWS_INIT_HEALTH_STATUS to specify the initial status of the custom health check, HEALTHY or UNHEALTHY. If you don&#39;t specify a value for AWS_INIT_HEALTH_STATUS, the initial status is HEALTHY.</td></tr>
	<tr><td>instance_ipv4</td><td>For an A record, the IPv4 address that Route 53 returns in response to DNS queries.</td></tr>
	<tr><td>instance_ipv6</td><td>For an AAAA record, the IPv6 address that Route 53 returns in response to DNS queries.</td></tr>
	<tr><td>instance_port</td><td>For an SRV record, the value that Route 53 returns for the port. In addition, if the service includes HealthCheckConfig, the port on the endpoint that Route 53 sends requests to.</td></tr>
	<tr><td>attributes</td><td>Attributes of the instance.</td></tr>
	<tr><td>title</td><td>Title of the resource.</td></tr>
	<tr><td>partition</td><td>The AWS partition in which the resource is located (aws, aws-cn, or aws-us-gov).</td></tr>
	<tr><td>region</td><td>The AWS Region in which the resource is located.</td></tr>
	<tr><td>account_id</td><td>The AWS Account ID in which the resource is located.</td></tr>
</table>