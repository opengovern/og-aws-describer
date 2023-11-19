# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>name</td><td>The custom name of the job.</td></tr>
	<tr><td>job_id</td><td>The unique identifier for the job.</td></tr>
	<tr><td>arn</td><td>The Amazon Resource Name (ARN) of the job.</td></tr>
	<tr><td>job_status</td><td>The status of a classification job.</td></tr>
	<tr><td>job_type</td><td>The schedule for running a classification job.</td></tr>
	<tr><td>client_token</td><td>The token that was provided to ensure the idempotency of the request to create the job.</td></tr>
	<tr><td>created_at</td><td>The date and time, in UTC and extended ISO 8601 format, when the job was created.</td></tr>
	<tr><td>last_run_time</td><td>This value indicates when the most recent run started.</td></tr>
	<tr><td>sampling_percentage</td><td>The sampling depth, as a percentage, that determines the percentage of eligible objects that the job analyzes.</td></tr>
	<tr><td>bucket_definitions</td><td>The namespace of the AWS service that provides the resource, or a custom-resource.</td></tr>
	<tr><td>custom_data_identifier_ids</td><td>The custom data identifiers that the job uses to analyze data.</td></tr>
	<tr><td>last_run_error_status</td><td>Specifies whether any account- or bucket-level access errors occurred when a classification job ran.</td></tr>
	<tr><td>s3_job_definition</td><td>Specifies which S3 buckets contain the objects that a classification job analyzes, and the scope of that analysis.</td></tr>
	<tr><td>schedule_frequency</td><td>Specifies the recurrence pattern for running a classification job.</td></tr>
	<tr><td>statistics</td><td>Provides processing statistics for a classification job.</td></tr>
	<tr><td>user_paused_details</td><td>Provides information about when a classification job was paused.</td></tr>
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