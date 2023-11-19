# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>db_cluster_identifier</td><td>Contains a user-supplied cluster identifier. This identifier is the unique key that identifies a cluster.</td></tr>
	<tr><td>arn</td><td>The Amazon Resource Name (ARN) for the Cluster.</td></tr>
	<tr><td>status</td><td>Specifies the current state of this cluster.</td></tr>
	<tr><td>cluster_create_time</td><td>Specifies the time when the cluster was created.</td></tr>
	<tr><td>backup_retention_period</td><td>Specifies the number of days for which automatic snapshots are retained.</td></tr>
	<tr><td>clone_group_id</td><td>Identifies the clone group to which the DB cluster is associated.</td></tr>
	<tr><td>db_cluster_parameter_group</td><td>Specifies the name of the cluster parameter group for the cluster.</td></tr>
	<tr><td>db_cluster_resource_id</td><td>The Region-unique, immutable identifier for the cluster.</td></tr>
	<tr><td>db_subnet_group</td><td>Specifies information on the subnet group associated with the cluster.</td></tr>
	<tr><td>deletion_protection</td><td>Specifies whether the cluster has deletion protection enabled, or not.</td></tr>
	<tr><td>earliest_restorable_time</td><td>The earliest time to which a database can be restored with point-in-time restore.</td></tr>
	<tr><td>endpoint</td><td>Specifies the connection endpoint for the primary instance of the DB cluster.</td></tr>
	<tr><td>engine</td><td>The name of the database engine to be used for this DB cluster.</td></tr>
	<tr><td>engine_version</td><td>Indicates the database engine version.</td></tr>
	<tr><td>hosted_zone_id</td><td>Specifies the ID that Amazon Route 53 assigns when you create a hosted zone.</td></tr>
	<tr><td>kms_key_id</td><td>The AWS KMS key identifier for the encrypted cluster.</td></tr>
	<tr><td>latest_restorable_time</td><td>Specifies the latest time to which a database can be restored with point-in-time restore.</td></tr>
	<tr><td>master_user_name</td><td>Contains the master username for the cluster.</td></tr>
	<tr><td>multi_az</td><td>Specifies whether the cluster has instances in multiple Availability Zones, or not.</td></tr>
	<tr><td>percent_progress</td><td>Specifies the progress of the operation as a percentage.</td></tr>
	<tr><td>port</td><td>Specifies the port that the database engine is listening on.</td></tr>
	<tr><td>preferred_backup_window</td><td>Specifies the daily time range during which automated backups are created.</td></tr>
	<tr><td>preferred_maintenance_window</td><td>Specifies the weekly time range during which system maintenance can occur</td></tr>
	<tr><td>reader_endpoint</td><td>The reader endpoint for the DB cluster.</td></tr>
	<tr><td>replication_source_identifier</td><td>Contains the identifier of the source cluster if this cluster is a secondary cluster.</td></tr>
	<tr><td>storage_encrypted</td><td>Specifies whether the cluster is encrypted, or not.</td></tr>
	<tr><td>associated_roles</td><td>A list of AWS IAM roles that are associated with the cluster.</td></tr>
	<tr><td>availability_zones</td><td>A list of Availability Zones (AZs) where instances in the cluster can be created.</td></tr>
	<tr><td>enabled_cloudwatch_logs_exports</td><td>A list of log types that this cluster is configured to export to Amazon CloudWatch Logs.</td></tr>
	<tr><td>members</td><td>A list of instances that make up the cluster.</td></tr>
	<tr><td>read_replica_identifiers</td><td>A list of identifiers of the read replicas associated with this cluster.</td></tr>
	<tr><td>vpc_security_groups</td><td>A list of VPC security groups that the DB cluster belongs to.</td></tr>
	<tr><td>tags_src</td><td>A list of tags attached to the Cluster.</td></tr>
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