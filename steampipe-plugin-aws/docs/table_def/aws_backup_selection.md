<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>selection_name</td><td>The display name of a resource selection document.</td></tr>
	<tr><td>selection_id</td><td>Uniquely identifies a request to assign a set of resources to a backup plan.</td></tr>
	<tr><td>backup_plan_id</td><td>An ID that uniquely identifies a backup plan.</td></tr>
	<tr><td>arn</td><td>The Amazon Resource Name (ARN) specifying the backup selection.</td></tr>
	<tr><td>creation_date</td><td>The date and time a resource backup plan is created.</td></tr>
	<tr><td>creator_request_id</td><td>An unique string that identifies the request and allows failed requests to be retried without the risk of running the operation twice.</td></tr>
	<tr><td>iam_role_arn</td><td>Specifies the IAM role Amazon Resource Name (ARN) to create the target recovery point.</td></tr>
	<tr><td>list_of_tags</td><td>An array of conditions used to specify a set of resources to assign to a backup plan.</td></tr>
	<tr><td>resources</td><td>Contains a list of BackupOptions for a resource type.</td></tr>
	<tr><td>title</td><td>Title of the resource.</td></tr>
	<tr><td>akas</td><td>Array of globally unique identifier strings (also known as) for the resource.</td></tr>
	<tr><td>partition</td><td>The AWS partition in which the resource is located (aws, aws-cn, or aws-us-gov).</td></tr>
	<tr><td>region</td><td>The AWS Region in which the resource is located.</td></tr>
	<tr><td>account_id</td><td>The AWS Account ID in which the resource is located.</td></tr>
	<tr><td>kaytu_account_id</td><td>The Kaytu Account ID in which the resource is located.</td></tr>
	<tr><td>kaytu_resource_id</td><td>The unique ID of the resource in Kaytu.</td></tr>
	<tr><td>kaytu_metadata</td><td>kaytu Metadata of the AWS resource.</td></tr>
</table>