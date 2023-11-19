# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>name</td><td>The name of the record.</td></tr>
	<tr><td>zone_id</td><td>The ID of the hosted zone to contain this record.</td></tr>
	<tr><td>type</td><td>The record type. Valid values are A, AAAA, CAA, CNAME, MX, NAPTR, NS, PTR, SOA, SPF, SRV and TXT.</td></tr>
	<tr><td>alias_target</td><td>Alias resource record sets only: Information about the AWS resource, such as a CloudFront distribution or an Amazon S3 bucket, that you want to route traffic to.</td></tr>
	<tr><td>failover</td><td>Failover resource record sets only: To configure failover, you add the Failover element to two resource record sets. For one resource record set, you specify PRIMARY as the value for Failover; for the other resource record set, you specify SECONDARY. In addition, you include the HealthCheckId element and specify the health check that you want Amazon Route 53 to perform for each resource record set.</td></tr>
	<tr><td>geo_location</td><td>Geolocation resource record sets only: A complex type that lets you control how Amazon Route 53 responds to DNS queries based on the geographic origin of the query. For example, if you want all queries from Africa to be routed to a web server with an IP address of 192.0.2.111, create a resource record set with a Type of A and a ContinentCode of AF.</td></tr>
	<tr><td>health_check_id</td><td>The health check the record should be associated with.</td></tr>
	<tr><td>multi_value_answer</td><td>Multivalue answer resource record sets only: To route traffic approximately randomly to multiple resources, such as web servers, create one multivalue answer record for each resource and specify true for MultiValueAnswer.</td></tr>
	<tr><td>latency_region</td><td>An AWS region from which to measure latency</td></tr>
	<tr><td>records</td><td>If the health check or hosted zone was created by another service, an optional description that can be provided by the other service.</td></tr>
	<tr><td>set_identifier</td><td>Unique identifier to differentiate records with routing policies from one another.</td></tr>
	<tr><td>ttl</td><td>The resource record cache time to live (TTL), in seconds.</td></tr>
	<tr><td>traffic_policy_instance_id</td><td>The ID of the traffic policy instance that Route 53 created this resource record set for.</td></tr>
	<tr><td>weight</td><td>Weighted resource record sets only: Among resource record sets that have the same combination of DNS name and type, a value that determines the proportion of DNS queries that Amazon Route 53 responds to using the current resource record set. Route 53 calculates the sum of the weights for the resource record sets that have the same combination of DNS name and type. Route 53 then responds to queries based on the ratio of a resource&#39;s weight to the total.</td></tr>
	<tr><td>title</td><td>Title of the resource.</td></tr>
	<tr><td>akas</td><td>Array of globally unique identifier strings (also known as) for the resource.</td></tr>
	<tr><td>partition</td><td>The AWS partition in which the resource is located (aws, aws-cn, or aws-us-gov).</td></tr>
	<tr><td>region</td><td>The AWS Region in which the resource is located.</td></tr>
	<tr><td>account_id</td><td>The AWS Account ID in which the resource is located.</td></tr>
	<tr><td>kaytu_account_id</td><td>The Kaytu Account ID in which the resource is located.</td></tr>
	<tr><td>kaytu_resource_id</td><td>The unique ID of the resource in Kaytu.</td></tr>
	<tr><td>kaytu_metadata</td><td>Kaytu Metadata of the AWS resource.</td></tr>
</table>