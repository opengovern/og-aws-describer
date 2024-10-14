# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>cluster_identifier</td><td>The unique identifier of the cluster.</td></tr>
	<tr><td>arn</td><td>The Amazon Resource Name (ARN) specifying the cluster.</td></tr>
	<tr><td>cluster_namespace_arn</td><td>The namespace Amazon Resource Name (ARN) of the cluster.</td></tr>
	<tr><td>allow_version_upgrade</td><td>A boolean value that, if true, indicates that major version upgrades will be applied automatically to the cluster during the maintenance window.</td></tr>
	<tr><td>automated_snapshot_retention_period</td><td>The number of days that automatic cluster snapshots are retained.</td></tr>
	<tr><td>availability_zone</td><td>The name of the Availability Zone in which the cluster is located.</td></tr>
	<tr><td>availability_zone_relocation_status</td><td>Describes the status of the Availability Zone relocation operation.</td></tr>
	<tr><td>cluster_availability_status</td><td>The availability status of the cluster for queries.</td></tr>
	<tr><td>cluster_create_time</td><td>The date and time that the cluster was created.</td></tr>
	<tr><td>cluster_nodes</td><td>The nodes in the cluster.</td></tr>
	<tr><td>cluster_parameter_groups</td><td>The list of cluster parameter groups that are associated with this cluster. Each parameter group in the list is returned with its status.</td></tr>
	<tr><td>cluster_public_key</td><td>The public key for the cluster.</td></tr>
	<tr><td>cluster_revision_number</td><td>The specific revision number of the database in the cluster.</td></tr>
	<tr><td>cluster_security_groups</td><td>A list of cluster security group that are associated with the cluster. Each security group is represented by an element that contains ClusterSecurityGroup.Name and ClusterSecurityGroup.Status subelements. Cluster security groups are used when the cluster is not created in an Amazon Virtual Private Cloud (VPC). Clusters that are created in a VPC use VPC security groups, which are listed by the VpcSecurityGroups parameter.</td></tr>
	<tr><td>cluster_snapshot_copy_status</td><td>A value that returns the destination region and retention period that are configured for cross-region snapshot copy.</td></tr>
	<tr><td>cluster_status</td><td>The current state of the cluster.</td></tr>
	<tr><td>cluster_subnet_group_name</td><td>The name of the subnet group that is associated with the cluster. This parameter is valid only when the cluster is in a VPC.</td></tr>
	<tr><td>cluster_version</td><td>The version ID of the Amazon Redshift engine that is running on the cluster.</td></tr>
	<tr><td>data_transfer_progress</td><td>Describes the status of a cluster while it is in the process of resizing with an incremental resize.</td></tr>
	<tr><td>db_name</td><td>The name of the initial database that was created when the cluster was created. This same name is returned for the life of the cluster. If an initial database was not specified, a database named devdev was created by default.</td></tr>
	<tr><td>deferred_maintenance_windows</td><td>Describes a group of DeferredMaintenanceWindow objects.</td></tr>
	<tr><td>elastic_ip_status</td><td>The status of the elastic IP (EIP) address.</td></tr>
	<tr><td>elastic_resize_number_of_node_options</td><td>The number of nodes that you can resize the cluster to with the elastic resize method.</td></tr>
	<tr><td>encrypted</td><td>A boolean value that, if true, indicates that data in the cluster is encrypted at rest.</td></tr>
	<tr><td>endpoint</td><td>The connection endpoint.</td></tr>
	<tr><td>enhanced_vpc_routing</td><td>An option that specifies whether to create the cluster with enhanced VPC routing enabled. To create a cluster that uses enhanced VPC routing, the cluster must be in a VPC. If this option is true, enhanced VPC routing is enabled.</td></tr>
	<tr><td>expected_next_snapshot_schedule_time</td><td>The date and time when the next snapshot is expected to be taken for clusters with a valid snapshot schedule and backups enabled.</td></tr>
	<tr><td>expected_next_snapshot_schedule_time_status</td><td>The status of next expected snapshot for clusters having a valid snapshot schedule and backups enabled.</td></tr>
	<tr><td>hsm_status</td><td>A value that reports whether the Amazon Redshift cluster has finished applying any hardware security module (HSM) settings changes specified in a modify cluster command.</td></tr>
	<tr><td>iam_roles</td><td>A list of AWS Identity and Access Management (IAM) roles that can be used by the cluster to access other AWS services.</td></tr>
	<tr><td>kms_key_id</td><td>The AWS Key Management Service (AWS KMS) key ID of the encryption key used to encrypt data in the cluster.</td></tr>
	<tr><td>maintenance_track_name</td><td>The name of the maintenance track for the cluster.</td></tr>
	<tr><td>manual_snapshot_retention_period</td><td>The default number of days to retain a manual snapshot. If the value is -1, the snapshot is retained indefinitely. This setting doesn&#39;t change the retention period of existing snapshots. The value must be either -1 or an integer between 1 and 3,653.</td></tr>
	<tr><td>master_username</td><td>The master user name for the cluster. This name is used to connect to the database that is specified in the DBName parameter.</td></tr>
	<tr><td>modify_status</td><td>The status of a modify operation, if any, initiated for the cluster.</td></tr>
	<tr><td>next_maintenance_window_start_time</td><td>The date and time in UTC when system maintenance can begin.</td></tr>
	<tr><td>node_type</td><td>The node type for the nodes in the cluster.</td></tr>
	<tr><td>number_of_nodes</td><td>The number of compute nodes in the cluster.</td></tr>
	<tr><td>pending_actions</td><td>Cluster operations that are waiting to be started.</td></tr>
	<tr><td>pending_modified_values</td><td>A value that, if present, indicates that changes to the cluster are pending. Specific pending changes are identified by subelements.</td></tr>
	<tr><td>preferred_maintenance_window</td><td>The weekly time range, in Universal Coordinated Time (UTC), during which system maintenance can occur.</td></tr>
	<tr><td>publicly_accessible</td><td>A boolean value that, if true, indicates that the cluster can be accessed from a public network.</td></tr>
	<tr><td>resize_info</td><td>Describes a resize operation.</td></tr>
	<tr><td>restore_status</td><td>A value that describes the status of a cluster restore action. This parameter returns null if the cluster was not created by restoring a snapshot.</td></tr>
	<tr><td>snapshot_schedule_identifier</td><td>A unique identifier for the cluster snapshot schedule.</td></tr>
	<tr><td>snapshot_schedule_state</td><td>The current state of the cluster snapshot schedule.</td></tr>
	<tr><td>vpc_id</td><td>The identifier of the VPC the cluster is in, if the cluster is in a VPC.</td></tr>
	<tr><td>vpc_security_groups</td><td>A list of Amazon Virtual Private Cloud (Amazon VPC) security groups that are associated with the cluster. This parameter is returned only if the cluster is in a VPC.</td></tr>
	<tr><td>logging_status</td><td>Describes the status of logging for a cluster.</td></tr>
	<tr><td>scheduled_actions</td><td>A list of scheduled actions for specified cluster.</td></tr>
	<tr><td>tags_src</td><td>The list of tags for the cluster.</td></tr>
	<tr><td>title</td><td>Title of the resource.</td></tr>
	<tr><td>tags</td><td>A map of tags for the resource.</td></tr>
	<tr><td>akas</td><td>Array of globally unique identifier strings (also known as) for the resource.</td></tr>
	<tr><td>partition</td><td>The AWS partition in which the resource is located (aws, aws-cn, or aws-us-gov).</td></tr>
	<tr><td>region</td><td>The AWS Region in which the resource is located.</td></tr>
	<tr><td>account_id</td><td>The AWS Account ID in which the resource is located.</td></tr>
	<tr><td>kaytu_account_id</td><td>The Kaytu Account ID in which the resource is located.</td></tr>
	<tr><td>kaytu_resource_id</td><td>The unique ID of the resource in opengovernance.</td></tr>
	<tr><td>kaytu_metadata</td><td>kaytu Metadata of the AWS resource.</td></tr>
</table>