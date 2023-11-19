# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>delivery_stream_name</td><td>The name of the delivery stream.</td></tr>
	<tr><td>arn</td><td>The Amazon Resource Name (ARN) of the delivery stream.</td></tr>
	<tr><td>delivery_stream_status</td><td>The server-side encryption type used on the stream.</td></tr>
	<tr><td>delivery_stream_type</td><td>The delivery stream type.</td></tr>
	<tr><td>version_id</td><td>The version id of the stream. Each time the destination is updated for a delivery stream, the version ID is changed, and the current version ID is required when updating the destination</td></tr>
	<tr><td>create_timestamp</td><td>The date and time that the delivery stream was created.</td></tr>
	<tr><td>has_more_destinations</td><td>Indicates whether there are more destinations available to list.</td></tr>
	<tr><td>last_update_timestamp</td><td>The date and time that the delivery stream was last updated.</td></tr>
	<tr><td>delivery_stream_encryption_configuration</td><td>Indicates the server-side encryption (SSE) status for the delivery stream.</td></tr>
	<tr><td>destinations</td><td>The destinations for the stream.</td></tr>
	<tr><td>failure_description</td><td>Provides details in case one of the following operations fails due to an error related to KMS: CreateDeliveryStream, DeleteDeliveryStream, StartDeliveryStreamEncryption,StopDeliveryStreamEncryption.</td></tr>
	<tr><td>source</td><td>If the DeliveryStreamType parameter is KinesisStreamAsSource, a SourceDescription object describing the source Kinesis data stream.</td></tr>
	<tr><td>tags_src</td><td>A list of tags associated with the delivery stream.</td></tr>
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