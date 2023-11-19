# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>name</td><td>The workgroup name.</td></tr>
	<tr><td>description</td><td>The workgroup description.</td></tr>
	<tr><td>creation_time</td><td>The date and time the workgroup was created.</td></tr>
	<tr><td>state</td><td>The state of the workgroup.</td></tr>
	<tr><td>additional_configuration</td><td>Specifies a user defined JSON string that is passed to the notebook engine.</td></tr>
	<tr><td>bytes_scanned_cutoff_per_query</td><td>The upper data usage limit (cutoff) for the amount of bytes a single query in a workgroup is allowed to scan.</td></tr>
	<tr><td>customer_content_kms_key</td><td>Specifies the KMS key that is used to encrypt the user&#39;s data stores in Athena.</td></tr>
	<tr><td>enforce_workgroup_configuration</td><td>If set to &#34;true&#34;, the settings for the workgroup override client-side settings.</td></tr>
	<tr><td>effective_engine_version</td><td>The engine version on which the query runs.</td></tr>
	<tr><td>selected_engine_version</td><td>The engine version requested by the user.</td></tr>
	<tr><td>execution_role</td><td>Role used in a notebook session for accessing the user&#39;s resources.</td></tr>
	<tr><td>publish_cloudwatch_metrics_enabled</td><td>Indicates that the Amazon CloudWatch metrics are enabled for the workgroup.</td></tr>
	<tr><td>requester_pays_enabled</td><td>If set to true, allows members assigned to a workgroup to reference Amazon S3 Requester Pays buckets in queries.</td></tr>
	<tr><td>s3_acl_option</td><td>The Amazon S3 canned ACL that Athena should specify when storing query results.</td></tr>
	<tr><td>encryption_option</td><td>Indicates whether Amazon S3 server-side encryption with Amazon S3-managed keys (SSE_S3), server-side encryption with KMS-managed keys (SSE_KMS), or client-side encryption with KMS-managed keys (CSE_KMS) is used.</td></tr>
	<tr><td>result_configuration_kms_key</td><td>For SSE_KMS and CSE_KMS, this is the KMS key ARN or ID.</td></tr>
	<tr><td>expected_bucket_owner</td><td>The Amazon Web Services account ID that you expect to be the owner of the Amazon S3 bucket specified by ResultConfiguration$OutputLocation.</td></tr>
	<tr><td>output_location</td><td>The location in Amazon S3 where your query results are stored.</td></tr>
	<tr><td>partition</td><td>The AWS partition in which the resource is located (aws, aws-cn, or aws-us-gov).</td></tr>
	<tr><td>region</td><td>The AWS Region in which the resource is located.</td></tr>
	<tr><td>account_id</td><td>The AWS Account ID in which the resource is located.</td></tr>
</table>