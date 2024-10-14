# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>topic_arn</td><td>Amazon Resource Name (ARN) of the Topic.</td></tr>
	<tr><td>display_name</td><td>The human-readable name used in the From field for notifications to email and email-json endpoints.</td></tr>
	<tr><td>application_failure_feedback_role_arn</td><td>IAM role for failed deliveries of notification messages sent to topics with platform application endpoint.</td></tr>
	<tr><td>application_success_feedback_role_arn</td><td>IAM role for successful deliveries of notification messages sent to topics with platform application endpoint.</td></tr>
	<tr><td>application_success_feedback_sample_rate</td><td>Sample rate for successful deliveries of notification messages sent to topics with platform application endpoint.</td></tr>
	<tr><td>firehose_failure_feedback_role_arn</td><td>IAM role for failed deliveries of notification messages sent to topics with kinesis data firehose endpoint.</td></tr>
	<tr><td>firehose_success_feedback_role_arn</td><td>IAM role for successful deliveries of notification messages sent to topics with kinesis data firehose endpoint.</td></tr>
	<tr><td>firehose_success_feedback_sample_rate</td><td>Sample rate for successful deliveries of notification messages sent to topics with firehose endpoint.</td></tr>
	<tr><td>http_failure_feedback_role_arn</td><td>IAM role for failed deliveries of notification messages sent to topics with http endpoint.</td></tr>
	<tr><td>http_success_feedback_role_arn</td><td>IAM role for successful deliveries of notification messages sent to topics with http endpoint.</td></tr>
	<tr><td>http_success_feedback_sample_rate</td><td>Sample rate for successful deliveries of notification messages sent to topics with http endpoint.</td></tr>
	<tr><td>lambda_failure_feedback_role_arn</td><td>IAM role for failed deliveries of notification messages sent to topics with lambda endpoint.</td></tr>
	<tr><td>lambda_success_feedback_role_arn</td><td>IAM role for successful deliveries of notification messages sent to topics with lambda endpoint.</td></tr>
	<tr><td>lambda_success_feedback_sample_rate</td><td>Sample rate for successful deliveries of notification messages sent to topics with lambda endpoint.</td></tr>
	<tr><td>owner</td><td>The AWS account ID of the topic&#39;s owner.</td></tr>
	<tr><td>sqs_failure_feedback_role_arn</td><td>IAM role for failed deliveries of notification messages sent to topics with sqs endpoint.</td></tr>
	<tr><td>sqs_success_feedback_role_arn</td><td>IAM role for successful deliveries of notification messages sent to topics with sqs endpoint.</td></tr>
	<tr><td>sqs_success_feedback_sample_rate</td><td>Sample rate for successful deliveries of notification messages sent to topics with sqs endpoint.</td></tr>
	<tr><td>subscriptions_confirmed</td><td>The number of confirmed subscriptions for the topic.</td></tr>
	<tr><td>subscriptions_deleted</td><td>The number of deleted subscriptions for the topic.</td></tr>
	<tr><td>subscriptions_pending</td><td>The number of subscriptions pending confirmation for the topic.</td></tr>
	<tr><td>kms_master_key_id</td><td>The ID of an AWS-managed customer master key (CMK) for Amazon SNS or a custom CMK.</td></tr>
	<tr><td>tags_src</td><td>The list of tags associated with the topic.</td></tr>
	<tr><td>policy</td><td>The topic&#39;s access control policy (i.e. Resource IAM Policy).</td></tr>
	<tr><td>policy_std</td><td>Contains the policy in a canonical form for easier searching.</td></tr>
	<tr><td>delivery_policy</td><td>The JSON object of the topic&#39;s delivery policy.</td></tr>
	<tr><td>effective_delivery_policy</td><td>The effective delivery policy, taking system defaults into account.</td></tr>
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