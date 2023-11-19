<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>stage_name</td><td>The name of the stage</td></tr>
	<tr><td>api_id</td><td>The id of the api which contains this stage</td></tr>
	<tr><td>api_gateway_managed</td><td>Specifies whether a stage is managed by API Gateway</td></tr>
	<tr><td>auto_deploy</td><td>Specifies whether updates to an API automatically trigger a new deployment</td></tr>
	<tr><td>client_certificate_id</td><td>The identifier of a client certificate for a Stage. Supported only for WebSocket APIs</td></tr>
	<tr><td>created_date</td><td>The timestamp when the stage was created</td></tr>
	<tr><td>deployment_id</td><td>The identifier of the Deployment that the Stage is associated with</td></tr>
	<tr><td>default_route_data_trace_enabled</td><td>Specifies whether (true) or not (false) data trace logging is enabled for this route. This property affects the log entries pushed to Amazon CloudWatch Logs. Supported only for WebSocket APIs</td></tr>
	<tr><td>default_route_detailed_metrics_enabled</td><td>Specifies whether detailed metrics are enabled</td></tr>
	<tr><td>default_route_logging_level</td><td>Specifies the logging level for this route: INFO, ERROR, or OFF. This property affects the log entries pushed to Amazon CloudWatch Logs. Supported only for WebSocket APIs</td></tr>
	<tr><td>default_route_throttling_burst_limit</td><td>Throttling burst limit for default route settings</td></tr>
	<tr><td>default_route_throttling_rate_limit</td><td>Throttling rate limit for default route settings</td></tr>
	<tr><td>last_deployment_status_message</td><td>Describes the status of the last deployment of a stage. Supported only for stages with autoDeploy enabled</td></tr>
	<tr><td>last_updated_date</td><td>The timestamp when the stage was last updated</td></tr>
	<tr><td>description</td><td>The stage's description</td></tr>
	<tr><td>access_log_settings</td><td>Access log settings of the stage.</td></tr>
	<tr><td>stage_variables</td><td>A map that defines the stage variables for a stage resource</td></tr>
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