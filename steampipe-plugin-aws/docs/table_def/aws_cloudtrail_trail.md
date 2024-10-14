# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>name</td><td>The name of the trail.</td></tr>
	<tr><td>arn</td><td>The Amazon Resource Name (ARN) of the trail.</td></tr>
	<tr><td>home_region</td><td>The region in which the trail was created.</td></tr>
	<tr><td>is_multi_region_trail</td><td>Specifies whether the trail exists only in one region or exists in all regions.</td></tr>
	<tr><td>log_file_validation_enabled</td><td>Specifies whether log file validation is enabled, or not.</td></tr>
	<tr><td>is_logging</td><td>Specifies whether the CloudTrail is currently logging AWS API calls, or not.</td></tr>
	<tr><td>log_group_arn</td><td>Specifies an Amazon Resource Name (ARN), a unique identifier that represents the log group to which CloudTrail logs will be delivered.</td></tr>
	<tr><td>cloudwatch_logs_role_arn</td><td>Specifies the role for the CloudWatch Logs endpoint to assume to write to a user&#39;s log group.</td></tr>
	<tr><td>has_custom_event_selectors</td><td>Specifies whether the trail has custom event selectors, or not.</td></tr>
	<tr><td>has_insight_selectors</td><td>Specifies whether a trail has insight types specified in an InsightSelector list, or not.</td></tr>
	<tr><td>include_global_service_events</td><td>Specifies whether to include AWS API calls from AWS global services, or not.</td></tr>
	<tr><td>is_organization_trail</td><td>Specifies whether the trail is an organization trail, or not.</td></tr>
	<tr><td>kms_key_id</td><td>Specifies the KMS key ID that encrypts the logs delivered by CloudTrail.</td></tr>
	<tr><td>s3_bucket_name</td><td>Name of the Amazon S3 bucket into which CloudTrail delivers your trail files.</td></tr>
	<tr><td>s3_key_prefix</td><td>Specifies the Amazon S3 key prefix that comes after the name of the bucket you have designated for log file delivery.</td></tr>
	<tr><td>sns_topic_arn</td><td>Specifies the ARN of the Amazon SNS topic that CloudTrail uses to send notifications when log files are delivered.</td></tr>
	<tr><td>latest_cloudwatch_logs_delivery_error</td><td>Displays any CloudWatch Logs error that CloudTrail encountered when attempting to deliver logs to CloudWatch Logs.</td></tr>
	<tr><td>latest_cloudwatch_logs_delivery_time</td><td>Displays the most recent date and time when CloudTrail delivered logs to CloudWatch Logs.</td></tr>
	<tr><td>latest_delivery_error</td><td>Displays any Amazon S3 error that CloudTrail encountered when attempting to deliver log files to the designated bucket.</td></tr>
	<tr><td>latest_delivery_time</td><td>Specifies the date and time that CloudTrail last delivered log files to an account&#39;s Amazon S3 bucket.</td></tr>
	<tr><td>latest_digest_delivery_error</td><td>Displays any Amazon S3 error that CloudTrail encountered when attempting to deliver a digest file to the designated bucket.</td></tr>
	<tr><td>latest_digest_delivery_time</td><td>Specifies the date and time that CloudTrail last delivered a digest file to an account&#39;s Amazon S3 bucket.</td></tr>
	<tr><td>latest_notification_error</td><td>Displays any Amazon SNS error that CloudTrail encountered when attempting to send a notification.</td></tr>
	<tr><td>latest_notification_time</td><td>Specifies the date and time of the most recent Amazon SNS notification that CloudTrail has written a new log file to an account&#39;s Amazon S3 bucket.</td></tr>
	<tr><td>start_logging_time</td><td>Specifies the most recent date and time when CloudTrail started recording API calls for an AWS account.</td></tr>
	<tr><td>stop_logging_time</td><td>Specifies the most recent date and time when CloudTrail stopped recording API calls for an AWS account.</td></tr>
	<tr><td>advanced_event_selectors</td><td>Describes the advanced event selectors that are configured for the trail.</td></tr>
	<tr><td>event_selectors</td><td>Describes the event selectors that are configured for the trail.</td></tr>
	<tr><td>insight_selectors</td><td>A JSON string that contains the insight types you want to log on a trail.</td></tr>
	<tr><td>tags_src</td><td>A list of tags assigned to the trail.</td></tr>
	<tr><td>tags</td><td>A map of tags for the resource.</td></tr>
	<tr><td>akas</td><td>Array of globally unique identifier strings (also known as) for the resource.</td></tr>
	<tr><td>title</td><td>Title of the resource.</td></tr>
	<tr><td>partition</td><td>The AWS partition in which the resource is located (aws, aws-cn, or aws-us-gov).</td></tr>
	<tr><td>region</td><td>The AWS Region in which the resource is located.</td></tr>
	<tr><td>account_id</td><td>The AWS Account ID in which the resource is located.</td></tr>
	<tr><td>kaytu_account_id</td><td>The Kaytu Account ID in which the resource is located.</td></tr>
	<tr><td>kaytu_resource_id</td><td>The unique ID of the resource in opengovernance.</td></tr>
	<tr><td>kaytu_metadata</td><td>kaytu Metadata of the AWS resource.</td></tr>
</table>