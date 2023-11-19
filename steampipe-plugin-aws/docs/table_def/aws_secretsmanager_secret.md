<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>name</td><td>The friendly name of the secret.</td></tr>
	<tr><td>arn</td><td>The Amazon Resource Name (ARN) of the secret.</td></tr>
	<tr><td>created_date</td><td>The date and time when a secret was created.</td></tr>
	<tr><td>description</td><td>The user-provided description of the secret.</td></tr>
	<tr><td>kms_key_id</td><td>The ARN or alias of the AWS KMS customer master key (CMK) used to encrypt the SecretString and SecretBinary fields in each version of the secret.</td></tr>
	<tr><td>deleted_date</td><td>The date and time the deletion of the secret occurred.</td></tr>
	<tr><td>last_accessed_date</td><td>The last date that this secret was accessed.</td></tr>
	<tr><td>last_changed_date</td><td>The last date and time that this secret was modified in any way.</td></tr>
	<tr><td>last_rotated_date</td><td>The most recent date and time that the Secrets Manager rotation process was successfully completed.</td></tr>
	<tr><td>owning_service</td><td>Returns the name of the service that created the secret.</td></tr>
	<tr><td>primary_region</td><td>The Region where Secrets Manager originated the secret.</td></tr>
	<tr><td>policy</td><td>A JSON-formatted string that describes the permissions that are associated with the attached secret.</td></tr>
	<tr><td>policy_std</td><td>Contains the permissions that are associated with the attached secret in a canonical form for easier searching.</td></tr>
	<tr><td>replication_status</td><td>Describes a list of replication status objects as InProgress, Failed or InSync.</td></tr>
	<tr><td>rotation_enabled</td><td>Indicates whether automatic, scheduled rotation is enabled for this secret.</td></tr>
	<tr><td>rotation_lambda_arn</td><td>The ARN of an AWS Lambda function invoked by Secrets Manager to rotate and expire the secret either automatically per the schedule or manually by a call to RotateSecret.</td></tr>
	<tr><td>rotation_rules</td><td>A structure that defines the rotation configuration for the secret.</td></tr>
	<tr><td>secret_versions_to_stages</td><td>A list of all of the currently assigned SecretVersionStage staging labels and the SecretVersionId attached to each one.</td></tr>
	<tr><td>tags_src</td><td>The list of user-defined tags associated with the secret.</td></tr>
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