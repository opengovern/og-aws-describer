# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>name</td><td>The name of the function.</td></tr>
	<tr><td>arn</td><td>The function&#39;s Amazon Resource Name (ARN).</td></tr>
	<tr><td>code_sha_256</td><td>The SHA256 hash of the function&#39;s deployment package.</td></tr>
	<tr><td>code_size</td><td>The size of the function&#39;s deployment package, in bytes.</td></tr>
	<tr><td>dead_letter_config_target_arn</td><td>The Amazon Resource Name (ARN) of an Amazon SQS queue or Amazon SNS topic.</td></tr>
	<tr><td>description</td><td>The function&#39;s description.</td></tr>
	<tr><td>handler</td><td>The function that Lambda calls to begin executing your function.</td></tr>
	<tr><td>kms_key_arn</td><td>The KMS key that&#39;s used to encrypt the function&#39;s environment variables. This key is only returned if you&#39;ve configured a customer managed CMK.</td></tr>
	<tr><td>last_modified</td><td>The date and time that the function was last updated.</td></tr>
	<tr><td>timeout</td><td>The amount of time in seconds that Lambda allows a function to run before stopping it.</td></tr>
	<tr><td>version</td><td>The version of the Lambda function.</td></tr>
	<tr><td>package_type</td><td>The type of deployment package.</td></tr>
	<tr><td>master_arn</td><td>For Lambda@Edge functions, the ARN of the master function.</td></tr>
	<tr><td>memory_size</td><td>The memory that&#39;s allocated to the function.</td></tr>
	<tr><td>revision_id</td><td>The latest updated revision of the function or alias.</td></tr>
	<tr><td>role</td><td>The function&#39;s execution role.</td></tr>
	<tr><td>runtime</td><td>The runtime environment for the Lambda function.</td></tr>
	<tr><td>state</td><td>The current state of the function.</td></tr>
	<tr><td>state_reason</td><td>The reason for the function&#39;s current state.</td></tr>
	<tr><td>state_reason_code</td><td>The reason code for the function&#39;s current state.</td></tr>
	<tr><td>last_update_status</td><td>The status of the last update that was performed on the function.</td></tr>
	<tr><td>last_update_status_reason</td><td>The reason for the last update that was performed on the function.</td></tr>
	<tr><td>last_update_status_reason_code</td><td>The reason code for the last update that was performed on the function.</td></tr>
	<tr><td>reserved_concurrent_executions</td><td>The number of concurrent executions that are reserved for this function.</td></tr>
	<tr><td>vpc_id</td><td>The VPC ID that is attached to Lambda function.</td></tr>
	<tr><td>architectures</td><td>The instruction set architecture that the function supports. Architecture is a string array with one of the valid values.</td></tr>
	<tr><td>code</td><td>The deployment package of the function or version.</td></tr>
	<tr><td>environment_variables</td><td>The environment variables that are accessible from function code during execution.</td></tr>
	<tr><td>file_system_configs</td><td>Connection settings for an Amazon EFS file system.</td></tr>
	<tr><td>policy</td><td>The resource-based iam policy of Lambda function.</td></tr>
	<tr><td>policy_std</td><td>Contains the policy in a canonical form for easier searching.</td></tr>
	<tr><td>tracing_config</td><td>The function&#39;s X-Ray tracing configuration.</td></tr>
	<tr><td>snap_start</td><td>Set ApplyOn to PublishedVersions to create a snapshot of the initialized execution environment when you publish a function version.</td></tr>
	<tr><td>url_config</td><td>The function URL configuration details of the function.</td></tr>
	<tr><td>vpc_security_group_ids</td><td>A list of VPC security groups IDs attached to Lambda function.</td></tr>
	<tr><td>vpc_subnet_ids</td><td>A list of VPC subnet IDs attached to Lambda function.</td></tr>
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