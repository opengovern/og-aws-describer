<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>db_cluster_identifier</td><td>Contains a user-supplied DB cluster identifier. This identifier is the unique key that identifies a DB cluster.</td></tr>
	<tr><td>arn</td><td>The Amazon Resource Name (ARN) for the DB cluster.</td></tr>
	<tr><td>status</td><td>Specifies the current state of this DB cluster.</td></tr>
	<tr><td>cluster_create_time</td><td>Specifies the time when the DB cluster was created, in Universal Coordinated Time (UTC).</td></tr>
	<tr><td>allocated_storage</td><td>AllocatedStorage always returns 1, because Neptune DB cluster storage size is not fixed, but instead automatically adjusts as needed.</td></tr>
	<tr><td>automatic_restart_time</td><td>Time at which the DB cluster will be automatically restarted.</td></tr>
	<tr><td>backup_retention_period</td><td>Specifies the number of days for which automatic DB snapshots are retained.</td></tr>
	<tr><td>clone_group_id</td><td>Identifies the clone group to which the DB cluster is associated.</td></tr>
	<tr><td>copy_tags_to_snapshot</td><td>If set to true, tags are copied to any snapshot of the DB cluster that is created.</td></tr>
	<tr><td>cross_account_clone</td><td>If set to true, the DB cluster can be cloned across accounts.</td></tr>
	<tr><td>db_cluster_parameter_group</td><td>Specifies the name of the DB cluster parameter group for the DB cluster.</td></tr>
	<tr><td>db_subnet_group</td><td>Specifies information on the subnet group associated with the DB cluster.</td></tr>
	<tr><td>database_name</td><td>Contains the name of the initial database of this DB cluster that was provided.</td></tr>
	<tr><td>db_cluster_resource_id</td><td>The Amazon Region-unique, immutable identifier for the DB cluster.</td></tr>
	<tr><td>deletion_protection</td><td>Indicates whether or not the DB cluster has deletion protection enabled.</td></tr>
	<tr><td>earliest_restorable_time</td><td>Specifies the earliest time to which a database can be restored with point-in-time restore.</td></tr>
	<tr><td>endpoint</td><td>Specifies the connection endpoint for the primary instance of the DB cluster.</td></tr>
	<tr><td>engine</td><td>Provides the name of the database engine to be used for this DB cluster.</td></tr>
	<tr><td>engine_version</td><td>Indicates the database engine version.</td></tr>
	<tr><td>hosted_zone_id</td><td>Specifies the ID that Amazon Route 53 assigns when you create a hosted zone.</td></tr>
	<tr><td>iam_database_authentication_enabled</td><td>True if mapping of Amazon Identity and Access Management (IAM) accounts to database accounts is enabled, and otherwise false.</td></tr>
	<tr><td>kms_key_id</td><td>If StorageEncrypted is true, the Amazon KMS key identifier for the encrypted DB cluster.</td></tr>
	<tr><td>latest_restorable_time</td><td>Specifies the latest time to which a database can be restored with point-in-time restore.</td></tr>
	<tr><td>multi_az</td><td>Specifies whether the DB cluster has instances in multiple Availability Zones.</td></tr>
	<tr><td>percent_progress</td><td>Specifies the progress of the operation as a percentage.</td></tr>
	<tr><td>port</td><td>Specifies the port that the database engine is listening on.</td></tr>
	<tr><td>preferred_backup_window</td><td>Specifies the daily time range during which automated backups are created if automated backups are enabled, as determined by the BackupRetentionPeriod.</td></tr>
	<tr><td>preferred_maintenance_window</td><td>Specifies the weekly time range during which system maintenance can occur, in Universal Coordinated Time (UTC).</td></tr>
	<tr><td>reader_endpoint</td><td>The reader endpoint for the DB cluster.</td></tr>
	<tr><td>storage_encrypted</td><td>Specifies whether the DB cluster is encrypted.</td></tr>
	<tr><td>associated_roles</td><td>Provides a list of the Amazon Identity and Access Management (IAM) roles.</td></tr>
	<tr><td>availability_zones</td><td>Provides the list of EC2 Availability Zones that instances in the DB cluster can be created in.</td></tr>
	<tr><td>db_cluster_members</td><td>Provides the list of instances that make up the DB cluster.</td></tr>
	<tr><td>enabled_cloudwatch_logs_exports</td><td>A list of log types that this DB cluster is configured to export to CloudWatch Logs.</td></tr>
	<tr><td>read_replica_identifiers</td><td>Contains one or more identifiers of the Read Replicas associated with this DB cluster.</td></tr>
	<tr><td>vpc_security_groups</td><td>Provides a list of VPC security groups that the DB cluster belongs to.</td></tr>
	<tr><td>tags_src</td><td>A list of tags assigned to the resource.</td></tr>
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