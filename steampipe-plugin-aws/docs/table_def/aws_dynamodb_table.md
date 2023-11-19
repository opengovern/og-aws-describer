<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>name</td><td>The name of the table.</td></tr>
	<tr><td>arn</td><td>The Amazon Resource Name (ARN) that uniquely identifies the table.</td></tr>
	<tr><td>table_id</td><td>Unique identifier for the table.</td></tr>
	<tr><td>creation_date_time</td><td>The date and time when the table was created.</td></tr>
	<tr><td>table_class</td><td>The table class of the specified table. Valid values are STANDARD and STANDARD_INFREQUENT_ACCESS.</td></tr>
	<tr><td>table_status</td><td>The current state of the table.</td></tr>
	<tr><td>billing_mode</td><td>Controls how AWS charges for read and write throughput and manage capacity.</td></tr>
	<tr><td>item_count</td><td>Number of items in the table. Note that this is an approximate value.</td></tr>
	<tr><td>global_table_version</td><td>Represents the version of global tables (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/GlobalTables.html) in use, if the table is replicated across AWS Regions.</td></tr>
	<tr><td>read_capacity</td><td>The maximum number of strongly consistent reads consumed per second before DynamoDB returns a ThrottlingException.</td></tr>
	<tr><td>write_capacity</td><td>The maximum number of writes consumed per second before DynamoDB returns a ThrottlingException.</td></tr>
	<tr><td>latest_stream_arn</td><td>The Amazon Resource Name (ARN) that uniquely identifies the latest stream for this table.</td></tr>
	<tr><td>latest_stream_label</td><td>A timestamp, in ISO 8601 format, for this stream.</td></tr>
	<tr><td>table_size_bytes</td><td>Size of the table in bytes. Note that this is an approximate value.</td></tr>
	<tr><td>archival_summary</td><td>Contains information about the table archive.</td></tr>
	<tr><td>attribute_definitions</td><td>An array of AttributeDefinition objects. Each of these objects describes one attribute in the table and index key schema.</td></tr>
	<tr><td>key_schema</td><td>The primary key structure for the table.</td></tr>
	<tr><td>sse_description</td><td>The description of the server-side encryption status on the specified table.</td></tr>
	<tr><td>continuous_backups_status</td><td>The continuous backups status of the table. ContinuousBackupsStatus can be one of the following states: ENABLED, DISABLED.</td></tr>
	<tr><td>streaming_destination</td><td>Provides information about the status of Kinesis streaming.</td></tr>
	<tr><td>point_in_time_recovery_description</td><td>The description of the point in time recovery settings applied to the table.</td></tr>
	<tr><td>tags_src</td><td>A list of tags assigned to the table.</td></tr>
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