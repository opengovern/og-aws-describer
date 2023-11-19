# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>db_instance_identifier</td><td>The friendly name to identify the DB Instance.</td></tr>
	<tr><td>arn</td><td>The Amazon Resource Name (ARN) for the DB Instance.</td></tr>
	<tr><td>db_cluster_identifier</td><td>The friendly name to identify the DB cluster, that the DB instance is a member of.</td></tr>
	<tr><td>status</td><td>Specifies the current state of this database.</td></tr>
	<tr><td>class</td><td>Contains the name of the compute and memory capacity class of the DB instance.</td></tr>
	<tr><td>resource_id</td><td>The AWS Region-unique, immutable identifier for the DB instance.</td></tr>
	<tr><td>allocated_storage</td><td>Specifies the allocated storage size specified in gibibytes(GiB).</td></tr>
	<tr><td>auto_minor_version_upgrade</td><td>Specifies whether minor version patches are applied automatically, or not.</td></tr>
	<tr><td>availability_zone</td><td>Specifies the name of the Availability Zone the DB instance is located in.</td></tr>
	<tr><td>backup_retention_period</td><td>Specifies the number of days for which automatic DB snapshots are retained.</td></tr>
	<tr><td>ca_certificate_identifier</td><td>The identifier of the CA certificate for this DB instance.</td></tr>
	<tr><td>character_set_name</td><td>Specifies the name of the character set that this instance is associated with.</td></tr>
	<tr><td>copy_tags_to_snapshot</td><td>Specifies whether tags are copied from the DB instance to snapshots of the DB instance, or not.</td></tr>
	<tr><td>customer_owned_ip_enabled</td><td>Specifies whether a customer-owned IP address (CoIP) is enabled for an RDS on Outposts DB instance, or not.</td></tr>
	<tr><td>port</td><td>Specifies the port that the DB instance listens on.</td></tr>
	<tr><td>db_name</td><td>Contains the name of the initial database of this instance that was provided at create time.</td></tr>
	<tr><td>db_subnet_group_arn</td><td>The Amazon Resource Name (ARN) for the DB subnet group.</td></tr>
	<tr><td>db_subnet_group_description</td><td>Provides the description of the DB subnet group.</td></tr>
	<tr><td>db_subnet_group_name</td><td>The name of the DB subnet group.</td></tr>
	<tr><td>db_subnet_group_status</td><td>Provides the status of the DB subnet group.</td></tr>
	<tr><td>deletion_protection</td><td>Specifies whether the DB instance has deletion protection enabled, or not.</td></tr>
	<tr><td>endpoint_address</td><td>Specifies the DNS address of the DB instance.</td></tr>
	<tr><td>endpoint_hosted_zone_id</td><td>Specifies the ID that Amazon Route 53 assigns when you create a hosted zone.</td></tr>
	<tr><td>endpoint_port</td><td>Specifies the port that the database engine is listening on.</td></tr>
	<tr><td>engine</td><td>The name of the database engine to be used for this DB instance.</td></tr>
	<tr><td>engine_version</td><td>Indicates the database engine version.</td></tr>
	<tr><td>enhanced_monitoring_resource_arn</td><td>The ARN of the Amazon CloudWatch Logs log stream that receives the Enhanced Monitoring metrics data for the DB instance.</td></tr>
	<tr><td>iam_database_authentication_enabled</td><td>Specifies whether the the mapping of AWS IAM accounts to database accounts is enabled, or not.</td></tr>
	<tr><td>create_time</td><td>Provides the date and time the DB instance was created.</td></tr>
	<tr><td>iops</td><td>Specifies the Provisioned IOPS (I/O operations per second) value.</td></tr>
	<tr><td>kms_key_id</td><td>The AWS KMS key identifier for the encrypted DB instance.</td></tr>
	<tr><td>latest_restorable_time</td><td>Specifies the latest time to which a database can be restored with point-in-time restore.</td></tr>
	<tr><td>license_model</td><td>License model information for this DB instance.</td></tr>
	<tr><td>master_user_name</td><td>Contains the master username for the DB instance.</td></tr>
	<tr><td>max_allocated_storage</td><td>The upper limit to which Amazon RDS can automatically scale the storage of the DB instance.</td></tr>
	<tr><td>monitoring_interval</td><td>The interval, in seconds, between points when Enhanced Monitoring metrics are collected for the DB instance.</td></tr>
	<tr><td>monitoring_role_arn</td><td>The ARN for the IAM role that permits RDS to send Enhanced Monitoring metrics to Amazon CloudWatch Logs.</td></tr>
	<tr><td>multi_az</td><td>Specifies if the DB instance is a Multi-AZ deployment.</td></tr>
	<tr><td>nchar_character_set_name</td><td>The name of the NCHAR character set for the Oracle DB instance.</td></tr>
	<tr><td>performance_insights_enabled</td><td>Specifies whether Performance Insights is enabled for the DB instance, or not.</td></tr>
	<tr><td>performance_insights_kms_key_id</td><td>The AWS KMS key identifier for encryption of Performance Insights data.</td></tr>
	<tr><td>performance_insights_retention_period</td><td>The amount of time, in days, to retain Performance Insights data.</td></tr>
	<tr><td>preferred_backup_window</td><td>Specifies the daily time range during which automated backups are created.</td></tr>
	<tr><td>preferred_maintenance_window</td><td>Specifies the weekly time range during which system maintenance can occur.</td></tr>
	<tr><td>promotion_tier</td><td>Specifies the order in which an Aurora Replica is promoted to the primary instance after a failure of the existing primary instance.</td></tr>
	<tr><td>publicly_accessible</td><td>Specifies the accessibility options for the DB instance.</td></tr>
	<tr><td>read_replica_source_db_instance_identifier</td><td>Contains the identifier of the source DB instance if this DB instance is a read replica.</td></tr>
	<tr><td>replica_mode</td><td>The mode of an Oracle read replica.</td></tr>
	<tr><td>secondary_availability_zone</td><td>Specifies the name of the secondary Availability Zone for a DB instance with multi-AZ support.</td></tr>
	<tr><td>storage_encrypted</td><td>Specifies whether the DB instance is encrypted, or not.</td></tr>
	<tr><td>storage_type</td><td>Specifies the storage type associated with DB instance.</td></tr>
	<tr><td>tde_credential_arn</td><td> The ARN from the key store with which the instance is associated for TDE encryption.</td></tr>
	<tr><td>timezone</td><td>The time zone of the DB instance.</td></tr>
	<tr><td>vpc_id</td><td>Provides the VpcId of the DB subnet group.</td></tr>
	<tr><td>associated_roles</td><td>A list of AWS IAM roles that are associated with the DB instance.</td></tr>
	<tr><td>certificate</td><td>The CA certificate associated with the DB instance.</td></tr>
	<tr><td>db_parameter_groups</td><td>A list of DB parameter groups applied to this DB instance.</td></tr>
	<tr><td>db_security_groups</td><td>A list of DB security group associated with the DB instance.</td></tr>
	<tr><td>domain_memberships</td><td>A list of Active Directory Domain membership records associated with the DB instance.</td></tr>
	<tr><td>enabled_cloudwatch_logs_exports</td><td>A list of log types that this DB instance is configured to export to CloudWatch Logs.</td></tr>
	<tr><td>option_group_memberships</td><td>A list of option group memberships for this DB instance</td></tr>
	<tr><td>pending_maintenance_actions</td><td>A list that provides details about the pending maintenance actions for the resource.</td></tr>
	<tr><td>processor_features</td><td>The number of CPU cores and the number of threads per core for the DB instance class of the DB instance.</td></tr>
	<tr><td>read_replica_db_cluster_identifiers</td><td>A list of identifiers of Aurora DB clusters to which the RDS DB instance is replicated as a read replica.</td></tr>
	<tr><td>read_replica_db_instance_identifiers</td><td>A list of identifiers of the read replicas associated with this DB instance.</td></tr>
	<tr><td>status_infos</td><td>The status of a read replica.</td></tr>
	<tr><td>subnets</td><td>A list of subnet elements.</td></tr>
	<tr><td>vpc_security_groups</td><td>A list of VPC security group elements that the DB instance belongs to.</td></tr>
	<tr><td>tags_src</td><td>A list of tags attached to the DB Instance.</td></tr>
	<tr><td>tags</td><td>A map of tags for the resource.</td></tr>
	<tr><td>title</td><td>Title of the resource.</td></tr>
	<tr><td>akas</td><td>Array of globally unique identifier strings (also known as) for the resource.</td></tr>
	<tr><td>partition</td><td>The AWS partition in which the resource is located (aws, aws-cn, or aws-us-gov).</td></tr>
	<tr><td>region</td><td>The AWS Region in which the resource is located.</td></tr>
	<tr><td>account_id</td><td>The AWS Account ID in which the resource is located.</td></tr>
	<tr><td>kaytu_account_id</td><td>The Kaytu Account ID in which the resource is located.</td></tr>
	<tr><td>kaytu_resource_id</td><td>The unique ID of the resource in Kaytu.</td></tr>
	<tr><td>kaytu_metadata</td><td>kaytu Metadata of the AWS resource.</td></tr>
</table>