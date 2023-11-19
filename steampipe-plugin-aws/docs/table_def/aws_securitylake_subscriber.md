# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>subscriber_name</td><td>The name of your Amazon Security Lake subscriber account.</td></tr>
	<tr><td>subscription_id</td><td>The subscription ID of the Amazon Security Lake subscriber account.</td></tr>
	<tr><td>created_at</td><td>The date and time when the subscription was created.</td></tr>
	<tr><td>subscription_status</td><td>Subscription status of the Amazon Security Lake subscriber account.</td></tr>
	<tr><td>updated_at</td><td>The date and time when the subscription was updated.</td></tr>
	<tr><td>external_id</td><td>The external ID of the subscriber.</td></tr>
	<tr><td>role_arn</td><td>The Amazon Resource Name (ARN) specifying the role of the subscriber.</td></tr>
	<tr><td>s3_bucket_arn</td><td>The Amazon Resource Name (ARN) for the Amazon S3 bucket.</td></tr>
	<tr><td>sns_arn</td><td>The Amazon Resource Name (ARN) for the Amazon Simple Notification Service.</td></tr>
	<tr><td>subscriber_description</td><td>The subscriber descriptions for a subscriber account.</td></tr>
	<tr><td>subscription_endpoint</td><td>The subscription endpoint to which exception messages are posted.</td></tr>
	<tr><td>subscription_protocol</td><td>The subscription protocol to which exception messages are posted.</td></tr>
	<tr><td>access_types</td><td>You can choose to notify subscribers of new objects with an Amazon Simple Queue Service (Amazon SQS) queue or through messaging to an HTTPS endpoint provided by the subscriber. Subscribers can consume data by directly querying Lake Formation tables in your S3 bucket via services like Amazon Athena. This subscription type is defined as LAKEFORMATION.</td></tr>
	<tr><td>source_types</td><td>Amazon Security Lake supports logs and events collection for the natively-supported Amazon Web Services services.</td></tr>
	<tr><td>title</td><td>Title of the resource.</td></tr>
	<tr><td>akas</td><td>Array of globally unique identifier strings (also known as) for the resource.</td></tr>
	<tr><td>partition</td><td>The AWS partition in which the resource is located (aws, aws-cn, or aws-us-gov).</td></tr>
	<tr><td>region</td><td>The AWS Region in which the resource is located.</td></tr>
	<tr><td>account_id</td><td>The AWS Account ID in which the resource is located.</td></tr>
	<tr><td>kaytu_account_id</td><td>The Kaytu Account ID in which the resource is located.</td></tr>
	<tr><td>kaytu_resource_id</td><td>The unique ID of the resource in Kaytu.</td></tr>
	<tr><td>kaytu_metadata</td><td>kaytu Metadata of the AWS resource.</td></tr>
</table>