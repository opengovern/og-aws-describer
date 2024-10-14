# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>log_group_name</td><td>The name of the log group to which this event belongs.</td></tr>
	<tr><td>log_stream_name</td><td>The name of the log stream to which this event belongs.</td></tr>
	<tr><td>event_id</td><td>The ID of the event.</td></tr>
	<tr><td>timestamp</td><td>The time when the event occurred.</td></tr>
	<tr><td>ingestion_time</td><td>The time when the event was ingested.</td></tr>
	<tr><td>filter</td><td>Filter pattern for the search.</td></tr>
	<tr><td>message</td><td>The data contained in the log event.</td></tr>
	<tr><td>message_json</td><td>The data contained in the log event in json format. Only if data is valid json string.</td></tr>
	<tr><td>partition</td><td>The AWS partition in which the resource is located (aws, aws-cn, or aws-us-gov).</td></tr>
	<tr><td>region</td><td>The AWS Region in which the resource is located.</td></tr>
	<tr><td>account_id</td><td>The AWS Account ID in which the resource is located.</td></tr>
	<tr><td>kaytu_account_id</td><td>The Kaytu Account ID in which the resource is located.</td></tr>
	<tr><td>kaytu_resource_id</td><td>The unique ID of the resource in opengovernance.</td></tr>
	<tr><td>kaytu_metadata</td><td>kaytu Metadata of the AWS resource.</td></tr>
</table>