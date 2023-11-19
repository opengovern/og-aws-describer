# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>db_instance_identifier</td><td>Contains a user-provided database identifier. This identifier is the unique key that identifies an instance.</td></tr>
	<tr><td>db_instance_arn</td><td>The Amazon Resource Name (ARN) for the instance.</td></tr>
	<tr><td>db_cluster_identifier</td><td>Contains the name of the cluster that the instance is a member of if the instance is a member of a cluster.</td></tr>
	<tr><td>db_instance_status</td><td>Specifies the current state of this database.</td></tr>
	<tr><td>db_instance_class</td><td>Contains the name of the compute and memory capacity class of the instance.</td></tr>
	<tr><td>dbi_resource_id</td><td>The Amazon Web Services Region-unique, immutable identifier for the instance.</td></tr>
	<tr><td>availability_zone</td><td>Specifies the name of the availability zone the instance is located in.</td></tr>
	<tr><td>backup_retention_period</td><td>Specifies the number of days for which automatic snapshots are retained.</td></tr>
	<tr><td>ca_certificate_identifier</td><td>The identifier of the CA certificate for this DB instance.</td></tr>
	<tr><td>copy_tags_to_snapshot</td><td>Specifies whether tags are copied from the DB instance to snapshots of the DB instance, or not.</td></tr>
	<tr><td>db_subnet_group_arn</td><td>The Amazon Resource Name (ARN) for the DB subnet group.</td></tr>
	<tr><td>db_subnet_group_description</td><td>Provides the description of the DB subnet group.</td></tr>
	<tr><td>db_subnet_group_name</td><td>The name of the DB subnet group.</td></tr>
	<tr><td>db_subnet_group_status</td><td>Provides the status of the DB subnet group.</td></tr>
	<tr><td>endpoint_address</td><td>Specifies the DNS address of the instance.</td></tr>
	<tr><td>endpoint_hosted_zone_id</td><td>Specifies the ID that Amazon Route 53 assigns when you create a hosted zone.</td></tr>
	<tr><td>endpoint_port</td><td>Specifies the port that the database engine is listening on.</td></tr>
	<tr><td>engine</td><td>The name of the database engine to be used for this instance.</td></tr>
	<tr><td>engine_version</td><td>Indicates the database engine version.</td></tr>
	<tr><td>instance_create_time</td><td>Provides the date and time the instance was created.</td></tr>
	<tr><td>kms_key_id</td><td>If StorageEncrypted is true, the KMS key identifier for the encrypted instance.</td></tr>
	<tr><td>latest_restorable_time</td><td>Specifies the latest time to which a database can be restored with point-in-time restore.</td></tr>
	<tr><td>preferred_backup_window</td><td>Specifies the daily time range during which automated backups are created.</td></tr>
	<tr><td>preferred_maintenance_window</td><td>Specifies the weekly time range during which system maintenance can occur.</td></tr>
	<tr><td>promotion_tier</td><td>A value that specifies the order in which an Amazon DocumentDB replica is promoted to the primary instance after a failure of the existing primary instance.</td></tr>
	<tr><td>publicly_accessible</td><td>Specifies the accessibility options for the DB instance.</td></tr>
	<tr><td>storage_encrypted</td><td>Specifies whether or not the instance is encrypted.</td></tr>
	<tr><td>vpc_id</td><td>Provides the VpcId of the DB subnet group.</td></tr>
	<tr><td>enabled_cloudwatch_logs_exports</td><td>A list of log types that this instance is configured to export to CloudWatch Logs.</td></tr>
	<tr><td>pending_modified_values</td><td>Specifies that changes to the instance are pending.</td></tr>
	<tr><td>status_infos</td><td>The status of a read replica.</td></tr>
	<tr><td>subnets</td><td>A list of subnet elements.</td></tr>
	<tr><td>vpc_security_groups</td><td>A list of VPC security group elements that the instance belongs to.</td></tr>
	<tr><td>tags_src</td><td>A list of tags attached to the Instance.</td></tr>
	<tr><td>tags</td><td>A map of tags for the resource.</td></tr>
	<tr><td>title</td><td>Title of the resource.</td></tr>
	<tr><td>akas</td><td>Array of globally unique identifier strings (also known as) for the resource.</td></tr>
	<tr><td>partition</td><td>The AWS partition in which the resource is located (aws, aws-cn, or aws-us-gov).</td></tr>
	<tr><td>region</td><td>The AWS Region in which the resource is located.</td></tr>
	<tr><td>account_id</td><td>The AWS Account ID in which the resource is located.</td></tr>
</table>