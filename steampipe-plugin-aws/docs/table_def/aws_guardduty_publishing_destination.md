# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>destination_id</td><td>The ID of the publishing destination.</td></tr>
	<tr><td>arn</td><td>The Amazon Resource Name (ARN) for the publishing destination.</td></tr>
	<tr><td>detector_id</td><td>The ID of the detector.</td></tr>
	<tr><td>destination_arn</td><td>The ARN of the resource to publish to.</td></tr>
	<tr><td>destination_type</td><td>The type of publishing destination. Currently, only Amazon S3 buckets are supported.</td></tr>
	<tr><td>kms_key_arn</td><td>The ARN of the KMS key to use for encryption.</td></tr>
	<tr><td>publishing_failure_start_timestamp</td><td>The time, in epoch millisecond format, at which GuardDuty was first unable to publish findings to the destination.</td></tr>
	<tr><td>status</td><td>The status of the publishing destination.</td></tr>
	<tr><td>title</td><td>Title of the resource.</td></tr>
	<tr><td>akas</td><td>Array of globally unique identifier strings (also known as) for the resource.</td></tr>
	<tr><td>partition</td><td>The AWS partition in which the resource is located (aws, aws-cn, or aws-us-gov).</td></tr>
	<tr><td>region</td><td>The AWS Region in which the resource is located.</td></tr>
	<tr><td>account_id</td><td>The AWS Account ID in which the resource is located.</td></tr>
	<tr><td>kaytu_account_id</td><td>The Kaytu Account ID in which the resource is located.</td></tr>
	<tr><td>kaytu_resource_id</td><td>The unique ID of the resource in Kaytu.</td></tr>
	<tr><td>kaytu_metadata</td><td>kaytu Metadata of the AWS resource.</td></tr>
</table>