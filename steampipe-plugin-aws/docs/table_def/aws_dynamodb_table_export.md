# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>arn</td><td>The Amazon Resource Name (ARN) of the export.</td></tr>
	<tr><td>export_status</td><td>Export can be in one of the following states: IN_PROGRESS, COMPLETED, or FAILED.</td></tr>
	<tr><td>billed_size_bytes</td><td>The billable size of the table export.</td></tr>
	<tr><td>client_token</td><td>The client token that was provided for the export task. A client token makes calls to ExportTableToPointInTimeInput idempotent, meaning that multiple identical calls have the same effect as one single call.</td></tr>
	<tr><td>end_time</td><td>The time at which the export task completed.</td></tr>
	<tr><td>export_format</td><td>The format of the exported data. Valid values for ExportFormat are DYNAMODB_JSON or ION.</td></tr>
	<tr><td>export_manifest</td><td>The name of the manifest file for the export task.</td></tr>
	<tr><td>export_time</td><td>Point in time from which table data was exported.</td></tr>
	<tr><td>failure_code</td><td>Status code for the result of the failed export.</td></tr>
	<tr><td>failure_message</td><td>Export failure reason description.</td></tr>
	<tr><td>item_count</td><td>The number of items exported.</td></tr>
	<tr><td>s3_bucket</td><td>The name of the Amazon S3 bucket containing the export.</td></tr>
	<tr><td>s3_bucket_owner</td><td>The ID of the Amazon Web Services account that owns the bucket containing the export.</td></tr>
	<tr><td>s3_prefix</td><td>The Amazon S3 bucket prefix used as the file name and path of the exported snapshot.</td></tr>
	<tr><td>s3_sse_algorithm</td><td>Type of encryption used on the bucket where export data is stored.</td></tr>
	<tr><td>s3_sse_kms_key_id</td><td>The ID of the KMS managed key used to encrypt the S3 bucket where export data is stored (if applicable).</td></tr>
	<tr><td>start_time</td><td>The time at which the export task began.</td></tr>
	<tr><td>table_arn</td><td>The Amazon Resource Name (ARN) of the table that was exported.</td></tr>
	<tr><td>table_id</td><td>Unique ID of the table that was exported.</td></tr>
	<tr><td>akas</td><td>Array of globally unique identifier strings (also known as) for the resource.</td></tr>
	<tr><td>partition</td><td>The AWS partition in which the resource is located (aws, aws-cn, or aws-us-gov).</td></tr>
	<tr><td>region</td><td>The AWS Region in which the resource is located.</td></tr>
	<tr><td>account_id</td><td>The AWS Account ID in which the resource is located.</td></tr>
	<tr><td>kaytu_account_id</td><td>The Kaytu Account ID in which the resource is located.</td></tr>
	<tr><td>kaytu_resource_id</td><td>The unique ID of the resource in Kaytu.</td></tr>
	<tr><td>kaytu_metadata</td><td>kaytu Metadata of the AWS resource.</td></tr>
</table>