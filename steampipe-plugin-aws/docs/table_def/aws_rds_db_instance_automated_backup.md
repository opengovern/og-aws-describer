# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>db_instance_identifier</td><td>The friendly name to identify the DB Instance.</td></tr>
	<tr><td>arn</td><td>The Amazon Resource Name (ARN) for the replicated automated backups.</td></tr>
	<tr><td>db_instance_arn</td><td>The Amazon Resource Name (ARN) for the automated backups.</td></tr>
	<tr><td>status</td><td>Specifies the current state of this database.</td></tr>
	<tr><td>allocated_storage</td><td>Specifies the allocated storage size in gibibytes (GiB).</td></tr>
	<tr><td>availability_zone</td><td>The Availability Zone that the automated backup was created in.</td></tr>
	<tr><td>backup_retention_period</td><td>The retention period for the automated backups.</td></tr>
	<tr><td>backup_target</td><td>Specifies where automated backups are stored: Amazon Web Services Outposts or the Amazon Web Services Region.</td></tr>
	<tr><td>dbi_resource_id</td><td>The identifier for the source DB instance, which can&#39;t be changed and which is unique to an Amazon Web Services Region.</td></tr>
	<tr><td>encrypted</td><td>Specifies whether the automated backup is encrypted.</td></tr>
	<tr><td>engine</td><td>The name of the database engine for this automated backup.</td></tr>
	<tr><td>engine_version</td><td>The version of the database engine for the automated backup.</td></tr>
	<tr><td>iam_database_authentication_enabled</td><td>True if mapping of Amazon Web Services Identity and Access Management (IAM) accounts to database accounts is enabled, and otherwise false.</td></tr>
	<tr><td>instance_create_time</td><td>True if mapping of Amazon Web Services Identity and Access Management (IAM) accounts to database accounts is enabled, and otherwise false.</td></tr>
	<tr><td>iops</td><td>True if mapping of Amazon Web Services Identity and Access Management (IAM) accounts to database accounts is enabled, and otherwise false.</td></tr>
	<tr><td>kms_key_id</td><td>The Amazon Web Services KMS key ID for an automated backup. The Amazon Web Services KMS key identifier is the key ARN, key ID, alias ARN, or alias name for the KMS key.</td></tr>
	<tr><td>license_model</td><td>The Amazon Web Services KMS key ID for an automated backup. The Amazon Web Services KMS key identifier is the key ARN, key ID, alias ARN, or alias name for the KMS key.</td></tr>
	<tr><td>master_username</td><td>The license model of an automated backup.</td></tr>
	<tr><td>option_group_name</td><td>The option group the automated backup is associated with. If omitted, the default option group for the engine specified is used.</td></tr>
	<tr><td>port</td><td>The port number that the automated backup used for connections. Default: Inherits from the source DB instance Valid Values: 1150-65535.</td></tr>
	<tr><td>storage_throughput</td><td>Specifies the storage throughput for the automated backup.</td></tr>
	<tr><td>storage_type</td><td>Specifies the storage type associated with the automated backup.</td></tr>
	<tr><td>tde_credential_arn</td><td>The ARN from the key store with which the automated backup is associated for TDE encryption.</td></tr>
	<tr><td>timezone</td><td>The time zone of the automated backup.</td></tr>
	<tr><td>vpc_id</td><td>Provides the VPC ID associated with the DB instance.</td></tr>
	<tr><td>db_instance_automated_backups_replications</td><td>The list of replications to different Amazon Web Services Regions associated with the automated backup.</td></tr>
	<tr><td>restore_window</td><td>Earliest and latest time an instance can be restored to.</td></tr>
	<tr><td>title</td><td>Title of the resource.</td></tr>
	<tr><td>akas</td><td>Array of globally unique identifier strings (also known as) for the resource.</td></tr>
	<tr><td>partition</td><td>The AWS partition in which the resource is located (aws, aws-cn, or aws-us-gov).</td></tr>
	<tr><td>region</td><td>The AWS Region in which the resource is located.</td></tr>
	<tr><td>account_id</td><td>The AWS Account ID in which the resource is located.</td></tr>
</table>