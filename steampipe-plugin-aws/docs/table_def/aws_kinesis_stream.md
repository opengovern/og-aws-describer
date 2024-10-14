# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>stream_name</td><td>The name of the stream being described.</td></tr>
	<tr><td>stream_arn</td><td>The Amazon Resource Name (ARN) for the stream being described.</td></tr>
	<tr><td>stream_status</td><td>The current status of the stream being described.</td></tr>
	<tr><td>stream_creation_timestamp</td><td>The approximate time that the stream was created.</td></tr>
	<tr><td>encryption_type</td><td>The server-side encryption type used on the stream.</td></tr>
	<tr><td>key_id</td><td>The GUID for the customer-managed AWS KMS key to use for encryption.</td></tr>
	<tr><td>retention_period_hours</td><td>The current retention period, in hours.</td></tr>
	<tr><td>consumer_count</td><td>The number of enhanced fan-out consumers registered with the stream.</td></tr>
	<tr><td>open_shard_count</td><td>The number of open shards in the stream.</td></tr>
	<tr><td>has_more_shards</td><td>If set to true, more shards in the stream are available to describe.</td></tr>
	<tr><td>shards</td><td>The shards that comprise the stream.</td></tr>
	<tr><td>enhanced_monitoring</td><td>Represents the current enhanced monitoring settings of the stream.</td></tr>
	<tr><td>tags_src</td><td>A list of tags associated with the stream.</td></tr>
	<tr><td>title</td><td>Title of the resource.</td></tr>
	<tr><td>tags</td><td>A map of tags for the resource.</td></tr>
	<tr><td>akas</td><td>Array of globally unique identifier strings (also known as) for the resource.</td></tr>
	<tr><td>partition</td><td>The AWS partition in which the resource is located (aws, aws-cn, or aws-us-gov).</td></tr>
	<tr><td>region</td><td>The AWS Region in which the resource is located.</td></tr>
	<tr><td>account_id</td><td>The AWS Account ID in which the resource is located.</td></tr>
	<tr><td>kaytu_account_id</td><td>The Kaytu Account ID in which the resource is located.</td></tr>
	<tr><td>kaytu_resource_id</td><td>The unique ID of the resource in opengovernance.</td></tr>
	<tr><td>kaytu_metadata</td><td>kaytu Metadata of the AWS resource.</td></tr>
</table>