<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>name</td><td>Name of the backup.</td></tr>
	<tr><td>arn</td><td>Amazon Resource Name associated with the backup.</td></tr>
	<tr><td>table_name</td><td>Unique identifier for the table to which backup belongs.</td></tr>
	<tr><td>table_arn</td><td>Name of the table to which backup belongs.</td></tr>
	<tr><td>table_id</td><td>ARN associated with the table to which backup belongs.</td></tr>
	<tr><td>backup_status</td><td>Current status of the backup. Backup can be in one of the following states: CREATING, ACTIVE, DELETED.</td></tr>
	<tr><td>backup_type</td><td>Backup type (USER | SYSTEM | AWS_BACKUP).</td></tr>
	<tr><td>backup_creation_datetime</td><td>Time at which the backup was created.</td></tr>
	<tr><td>backup_expiry_datetime</td><td>Time at which the automatic on-demand backup created by DynamoDB will expire.</td></tr>
	<tr><td>backup_size_bytes</td><td>Size of the backup in bytes.</td></tr>
	<tr><td>title</td><td>Title of the resource.</td></tr>
	<tr><td>akas</td><td>Array of globally unique identifier strings (also known as) for the resource.</td></tr>
	<tr><td>partition</td><td>The AWS partition in which the resource is located (aws, aws-cn, or aws-us-gov).</td></tr>
	<tr><td>region</td><td>The AWS Region in which the resource is located.</td></tr>
	<tr><td>account_id</td><td>The AWS Account ID in which the resource is located.</td></tr>
	<tr><td>kaytu_account_id</td><td>The Kaytu Account ID in which the resource is located.</td></tr>
	<tr><td>kaytu_resource_id</td><td>The unique ID of the resource in Kaytu.</td></tr>
	<tr><td>kaytu_metadata</td><td>kaytu Metadata of the AWS resource.</td></tr>
</table>