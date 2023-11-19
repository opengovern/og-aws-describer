<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>name</td><td>The name of a logical container where backups are stored.</td></tr>
	<tr><td>arn</td><td>An Amazon Resource Name (ARN) that uniquely identifies a backup vault.</td></tr>
	<tr><td>creation_date</td><td>The date and time a resource backup is created.</td></tr>
	<tr><td>creator_request_id</td><td>An unique string that identifies the request and allows failed requests to be retried without the risk of running the operation twice.</td></tr>
	<tr><td>encryption_key_arn</td><td>The server-side encryption key that is used to protect your backups.</td></tr>
	<tr><td>number_of_recovery_points</td><td>The number of recovery points that are stored in a backup vault.</td></tr>
	<tr><td>sns_topic_arn</td><td>An ARN that uniquely identifies an Amazon Simple Notification Service.</td></tr>
	<tr><td>policy</td><td>The backup vault access policy document in JSON format.</td></tr>
	<tr><td>policy_std</td><td>Contains the backup vault access policy document in a canonical form for easier searching.</td></tr>
	<tr><td>backup_vault_events</td><td>An array of events that indicate the status of jobs to back up resources to the backup vault.</td></tr>
	<tr><td>title</td><td>Title of the resource.</td></tr>
	<tr><td>akas</td><td>Array of globally unique identifier strings (also known as) for the resource.</td></tr>
	<tr><td>partition</td><td>The AWS partition in which the resource is located (aws, aws-cn, or aws-us-gov).</td></tr>
	<tr><td>region</td><td>The AWS Region in which the resource is located.</td></tr>
	<tr><td>account_id</td><td>The AWS Account ID in which the resource is located.</td></tr>
	<tr><td>kaytu_account_id</td><td>The Kaytu Account ID in which the resource is located.</td></tr>
	<tr><td>kaytu_resource_id</td><td>The unique ID of the resource in Kaytu.</td></tr>
	<tr><td>kaytu_metadata</td><td>kaytu Metadata of the AWS resource.</td></tr>
</table>