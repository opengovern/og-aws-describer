<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>queue_url</td><td>The URL of the Amazon SQS queue.</td></tr>
	<tr><td>queue_arn</td><td>The Amazon resource name (ARN) of the queue.</td></tr>
	<tr><td>fifo_queue</td><td>Returns true if the queue is FIFO.</td></tr>
	<tr><td>fifo_throughput_limit</td><td>Specifies whether the FIFO queue throughput quota applies to the entire queue or per message group.</td></tr>
	<tr><td>delay_seconds</td><td>The default delay on the queue in seconds.</td></tr>
	<tr><td>max_message_size</td><td>The limit of how many bytes a message can contain before Amazon SQS rejects it.</td></tr>
	<tr><td>message_retention_seconds</td><td>The length of time, in seconds, for which Amazon SQS retains a message.</td></tr>
	<tr><td>receive_wait_time_seconds</td><td>The length of time, in seconds, for which the ReceiveMessage action waits for a message to arrive.</td></tr>
	<tr><td>sqs_managed_sse_enabled</td><td>Returns true if the queue is using SSE-SQS encryption with SQS-owned encryption keys.</td></tr>
	<tr><td>visibility_timeout_seconds</td><td>The visibility timeout for the queue in seconds.</td></tr>
	<tr><td>policy</td><td>The resource IAM policy of the queue.</td></tr>
	<tr><td>policy_std</td><td>Contains the policy in a canonical form for easier searching.</td></tr>
	<tr><td>redrive_policy</td><td>The string that includes the parameters for the dead-letter queue functionality of the source queue as a JSON object.</td></tr>
	<tr><td>content_based_deduplication</td><td>Mentions whether content-based deduplication is enabled for the queue.</td></tr>
	<tr><td>kms_master_key_id</td><td>the ID of an AWS-managed customer master key (CMK) for Amazon SQS or a custom CMK.</td></tr>
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