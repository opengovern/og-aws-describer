<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>flow_log_id</td><td>The ID of the flow log.</td></tr>
	<tr><td>creation_time</td><td>The date and time the flow log was created.</td></tr>
	<tr><td>deliver_logs_error_message</td><td>Information about the error that occurred.</td></tr>
	<tr><td>deliver_logs_permission_arn</td><td>The ARN of the IAM role that posts logs to CloudWatch Logs.</td></tr>
	<tr><td>deliver_logs_status</td><td>The status of the logs delivery (SUCCESS | FAILED).</td></tr>
	<tr><td>flow_log_status</td><td>The status of the flow log (ACTIVE).</td></tr>
	<tr><td>log_group_name</td><td>The name of the flow log group.</td></tr>
	<tr><td>resource_id</td><td>The ID of the VPC, subnet, or network interface.</td></tr>
	<tr><td>traffic_type</td><td>The type of traffic. Valid values are: 'ACCEPT', 'REJECT',  'ALL'.</td></tr>
	<tr><td>log_destination_type</td><td>Specifies the type of destination to which the flow log data is published.</td></tr>
	<tr><td>log_destination</td><td>Specifies the destination to which the flow log data is published.</td></tr>
	<tr><td>bucket_name</td><td>The name of the destination bucket to which the flow log data is published.</td></tr>
	<tr><td>log_format</td><td>The format of the flow log record.</td></tr>
	<tr><td>max_aggregation_interval</td><td>The maximum interval of time, in seconds, during which a flow of packets is captured and aggregated into a flow log record.</td></tr>
	<tr><td>tags_src</td><td>A list of tags assigned to the VPC flowlog.</td></tr>
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