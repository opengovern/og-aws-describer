<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>query_id</td><td>The ID of the query.</td></tr>
	<tr><td>event_data_store_arn</td><td>The ID of the event data store.</td></tr>
	<tr><td>creation_time</td><td>The creation time of the query.</td></tr>
	<tr><td>delivery_s3_uri</td><td>The URI for the S3 bucket where CloudTrail delivered query results, if applicable.</td></tr>
	<tr><td>delivery_status</td><td>The delivery status.</td></tr>
	<tr><td>error_message</td><td>The error message returned if a query failed.</td></tr>
	<tr><td>query_status</td><td>The status of a query. Values for QueryStatus include QUEUED, RUNNING, FINISHED, FAILED, TIMED_OUT, or CANCELLED.</td></tr>
	<tr><td>bytes_scanned</td><td>Gets metadata about a query, including the number of events that were matched, the total number of events scanned, the query run time in milliseconds, and the query's creation time.</td></tr>
	<tr><td>events_matched</td><td>The number of events that matched a query.</td></tr>
	<tr><td>events_scanned</td><td>The number of events that the query scanned in the event data store.</td></tr>
	<tr><td>execution_time_in_millis</td><td>The query's run time, in milliseconds.</td></tr>
	<tr><td>query_string</td><td>The SQL code of a query.</td></tr>
	<tr><td>title</td><td>Title of the resource.</td></tr>
	<tr><td>partition</td><td>The AWS partition in which the resource is located (aws, aws-cn, or aws-us-gov).</td></tr>
	<tr><td>region</td><td>The AWS Region in which the resource is located.</td></tr>
	<tr><td>account_id</td><td>The AWS Account ID in which the resource is located.</td></tr>
	<tr><td>kaytu_account_id</td><td>The Kaytu Account ID in which the resource is located.</td></tr>
	<tr><td>kaytu_resource_id</td><td>The unique ID of the resource in Kaytu.</td></tr>
	<tr><td>kaytu_metadata</td><td>kaytu Metadata of the AWS resource.</td></tr>
</table>