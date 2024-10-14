# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>db_cluster_identifier</td><td>The friendly name to identify the DB Cluster.</td></tr>
	<tr><td>arn</td><td>The Amazon Resource Name (ARN) for the DB Cluster.</td></tr>
	<tr><td>status</td><td>Specifies the status of this DB Cluster.</td></tr>
	<tr><td>resource_id</td><td>The AWS Region-unique, immutable identifier for the DB cluster.</td></tr>
	<tr><td>create_time</td><td>Specifies the time when the DB cluster was created.</td></tr>
	<tr><td>activity_stream_kinesis_stream_name</td><td>The name of the Amazon Kinesis data stream used for the database activity stream.</td></tr>
	<tr><td>activity_stream_kms_key_id</td><td>The AWS KMS key identifier used for encrypting messages in the database activity stream.</td></tr>
	<tr><td>activity_stream_mode</td><td>The mode of the database activity stream.</td></tr>
	<tr><td>activity_stream_status</td><td>The status of the database activity stream.</td></tr>
	<tr><td>allocated_storage</td><td>Specifies the allocated storage size in gibibytes (GiB).</td></tr>
	<tr><td>backtrack_consumed_change_records</td><td>The number of change records stored for Backtrack.</td></tr>
	<tr><td>backtrack_window</td><td>The target backtrack window, in seconds.</td></tr>
	<tr><td>backup_retention_period</td><td>Specifies the number of days for which automatic DB snapshots are retained.</td></tr>
	<tr><td>capacity</td><td>The current capacity of an Aurora Serverless DB cluster.</td></tr>
	<tr><td>character_set_name</td><td>Specifies the name of the character set that this cluster is associated with.</td></tr>
	<tr><td>clone_group_id</td><td>Identifies the clone group to which the DB cluster is associated.</td></tr>
	<tr><td>copy_tags_to_snapshot</td><td>Specifies whether tags are copied from the DB cluster to snapshots of the DB cluster, or not.</td></tr>
	<tr><td>cross_account_clone</td><td>Specifies whether the DB cluster is a clone of a DB cluster owned by a different AWS account, or not.</td></tr>
	<tr><td>database_name</td><td>Contains the name of the initial database of this DB cluster that was provided at create time.</td></tr>
	<tr><td>db_cluster_parameter_group</td><td>Specifies the name of the DB cluster parameter group for the DB cluster.</td></tr>
	<tr><td>db_subnet_group</td><td>Specifies information on the subnet group associated with the DB cluster.</td></tr>
	<tr><td>deletion_protection</td><td>Specifies whether the DB cluster has deletion protection enabled, or not.</td></tr>
	<tr><td>earliest_backtrack_time</td><td>The earliest time to which a DB cluster can be backtracked.</td></tr>
	<tr><td>earliest_restorable_time</td><td>The earliest time to which a database can be restored with point-in-time restore.</td></tr>
	<tr><td>endpoint</td><td>Specifies the connection endpoint for the primary instance of the DB cluster.</td></tr>
	<tr><td>engine</td><td>The name of the database engine to be used for this DB cluster.</td></tr>
	<tr><td>engine_mode</td><td>The DB engine mode of the DB cluster.</td></tr>
	<tr><td>engine_version</td><td>Indicates the database engine version.</td></tr>
	<tr><td>global_write_forwarding_requested</td><td>Specifies whether you have requested to enable write forwarding for a secondary cluster in an Aurora global database, or not.</td></tr>
	<tr><td>global_write_forwarding_status</td><td>Specifies whether a secondary cluster in an Aurora global database has write forwarding enabled, or not.</td></tr>
	<tr><td>hosted_zone_id</td><td>Specifies the ID that Amazon Route 53 assigns when you create a hosted zone.</td></tr>
	<tr><td>http_endpoint_enabled</td><td>Specifies whether the HTTP endpoint for an Aurora Serverless DB cluster is enabled, or not.</td></tr>
	<tr><td>iam_database_authentication_enabled</td><td>Specifies whether the the mapping of AWS IAM accounts to database accounts is enabled, or not.</td></tr>
	<tr><td>kms_key_id</td><td>The AWS KMS key identifier for the encrypted DB cluster.</td></tr>
	<tr><td>latest_restorable_time</td><td>Specifies the latest time to which a database can be restored with point-in-time restore.</td></tr>
	<tr><td>master_user_name</td><td>Contains the master username for the DB cluster.</td></tr>
	<tr><td>multi_az</td><td>Specifies whether the DB cluster has instances in multiple Availability Zones, or not.</td></tr>
	<tr><td>percent_progress</td><td>Specifies the progress of the operation as a percentage.</td></tr>
	<tr><td>port</td><td>Specifies the port that the database engine is listening on.</td></tr>
	<tr><td>preferred_backup_window</td><td>Specifies the daily time range during which automated backups are created.</td></tr>
	<tr><td>preferred_maintenance_window</td><td>Specifies the weekly time range during which system maintenance can occur</td></tr>
	<tr><td>reader_endpoint</td><td>The reader endpoint for the DB cluster.</td></tr>
	<tr><td>storage_encrypted</td><td>Specifies whether the DB cluster is encrypted, or not.</td></tr>
	<tr><td>associated_roles</td><td>A list of AWS IAM roles that are associated with the DB cluster.</td></tr>
	<tr><td>availability_zones</td><td>A list of Availability Zones (AZs) where instances in the DB cluster can be created.</td></tr>
	<tr><td>custom_endpoints</td><td>A list of all custom endpoints associated with the cluster.</td></tr>
	<tr><td>members</td><td>A list of instances that make up the DB cluster.</td></tr>
	<tr><td>option_group_memberships</td><td>A list of option group memberships for this DB cluster.</td></tr>
	<tr><td>domain_memberships</td><td>A list of Active Directory Domain membership records associated with the DB cluster.</td></tr>
	<tr><td>enabled_cloudwatch_logs_exports</td><td>A list of log types that this DB cluster is configured to export to CloudWatch Logs.</td></tr>
	<tr><td>read_replica_identifiers</td><td>A list of identifiers of the read replicas associated with this DB cluster.</td></tr>
	<tr><td>vpc_security_groups</td><td>A list of VPC security groups that the DB cluster belongs to.</td></tr>
	<tr><td>tags_src</td><td>A list of tags attached to the DB Cluster.</td></tr>
	<tr><td>tags</td><td>A map of tags for the resource.</td></tr>
	<tr><td>title</td><td>Title of the resource.</td></tr>
	<tr><td>akas</td><td>Array of globally unique identifier strings (also known as) for the resource.</td></tr>
	<tr><td>partition</td><td>The AWS partition in which the resource is located (aws, aws-cn, or aws-us-gov).</td></tr>
	<tr><td>region</td><td>The AWS Region in which the resource is located.</td></tr>
	<tr><td>account_id</td><td>The AWS Account ID in which the resource is located.</td></tr>
	<tr><td>kaytu_account_id</td><td>The Kaytu Account ID in which the resource is located.</td></tr>
	<tr><td>kaytu_resource_id</td><td>The unique ID of the resource in opengovernance.</td></tr>
	<tr><td>kaytu_metadata</td><td>kaytu Metadata of the AWS resource.</td></tr>
</table>