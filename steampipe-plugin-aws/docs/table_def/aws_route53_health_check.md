# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>id</td><td>The identifier that Amazon Route 53 assigned to the health check.</td></tr>
	<tr><td>caller_reference</td><td>A unique string that you specified when you created the health check.</td></tr>
	<tr><td>health_check_version</td><td>The version of the health check.</td></tr>
	<tr><td>linked_service_principal</td><td>If the health check was created by another service, the service that created the resource.</td></tr>
	<tr><td>linked_service_description</td><td>If the health check was created by another service, an configurationtional description that can be provided by the other service.</td></tr>
	<tr><td>cloud_watch_alarm_configuration</td><td>A complex type that contains information about the CloudWatch alarm that Amazon Route 53 is monitoring for this health check.</td></tr>
	<tr><td>health_check_config</td><td>A complex type that contains detailed information about one health check.</td></tr>
	<tr><td>health_check_status</td><td>A list that contains one HealthCheckObservation element for each Amazon Route 53 health checker that is reporting a status about the health check endpoint.</td></tr>
	<tr><td>tags_src</td><td>A map of tags for the resource.</td></tr>
	<tr><td>title</td><td>Title of the resource.</td></tr>
	<tr><td>tags</td><td>A map of tags for the resource.</td></tr>
	<tr><td>akas</td><td>Array of globally unique identifier strings (also known as) for the resource.</td></tr>
	<tr><td>partition</td><td>The AWS partition in which the resource is located (aws, aws-cn, or aws-us-gov).</td></tr>
	<tr><td>region</td><td>The AWS Region in which the resource is located.</td></tr>
	<tr><td>account_id</td><td>The AWS Account ID in which the resource is located.</td></tr>
	<tr><td>kaytu_account_id</td><td>The Kaytu Account ID in which the resource is located.</td></tr>
	<tr><td>kaytu_resource_id</td><td>The unique ID of the resource in Kaytu.</td></tr>
	<tr><td>kaytu_metadata</td><td>Kaytu Metadata of the AWS resource.</td></tr>
</table>