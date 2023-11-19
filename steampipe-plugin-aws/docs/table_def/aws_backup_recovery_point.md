<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>backup_vault_name</td><td>The name of a logical container where backups are stored.</td></tr>
	<tr><td>recovery_point_arn</td><td>An ARN that uniquely identifies a recovery point.</td></tr>
	<tr><td>resource_type</td><td>The type of Amazon Web Services resource to save as a recovery point.</td></tr>
	<tr><td>status</td><td>A status code specifying the state of the recovery point.</td></tr>
	<tr><td>backup_size_in_bytes</td><td>The size, in bytes, of a backup.</td></tr>
	<tr><td>backup_vault_arn</td><td>An ARN that uniquely identifies a backup vault.</td></tr>
	<tr><td>creation_date</td><td>The date and time that a recovery point is created.</td></tr>
	<tr><td>completion_date</td><td>The date and time that a job to create a recovery point is completed.</td></tr>
	<tr><td>encryption_key_arn</td><td>The server-side encryption key used to protect your backups.</td></tr>
	<tr><td>iam_role_arn</td><td>Specifies the IAM role ARN used to create the target recovery point.</td></tr>
	<tr><td>is_encrypted</td><td>A Boolean value that is returned as TRUE if the specified recovery point is encrypted, or FALSE if the recovery point is not encrypted.</td></tr>
	<tr><td>last_restore_time</td><td>The date and time that a recovery point was last restored.</td></tr>
	<tr><td>resource_arn</td><td>An ARN that uniquely identifies a saved resource.</td></tr>
	<tr><td>source_backup_vault_arn</td><td>An Amazon Resource Name (ARN) that uniquely identifies the source vault where the resource was originally backed up in.</td></tr>
	<tr><td>status_message</td><td>A status message explaining the reason for the recovery point deletion failure.</td></tr>
	<tr><td>storage_class</td><td>Specifies the storage class of the recovery point. Valid values are WARM or COLD.</td></tr>
	<tr><td>calculated_lifecycle</td><td>An object containing DeleteAt and MoveToColdStorageAt timestamps.</td></tr>
	<tr><td>created_by</td><td>Contains identifying information about the creation of a recovery point, including the BackupPlanArn, BackupPlanId, BackupPlanVersion, and BackupRuleId of the backup plan used to create it.</td></tr>
	<tr><td>lifecycle</td><td>The lifecycle defines when a protected resource is transitioned to cold storage and when it expires.</td></tr>
	<tr><td>akas</td><td>Array of globally unique identifier strings (also known as) for the resource.</td></tr>
	<tr><td>partition</td><td>The AWS partition in which the resource is located (aws, aws-cn, or aws-us-gov).</td></tr>
	<tr><td>region</td><td>The AWS Region in which the resource is located.</td></tr>
	<tr><td>account_id</td><td>The AWS Account ID in which the resource is located.</td></tr>
	<tr><td>kaytu_account_id</td><td>The Kaytu Account ID in which the resource is located.</td></tr>
	<tr><td>kaytu_resource_id</td><td>The unique ID of the resource in Kaytu.</td></tr>
	<tr><td>kaytu_metadata</td><td>kaytu Metadata of the AWS resource.</td></tr>
</table>