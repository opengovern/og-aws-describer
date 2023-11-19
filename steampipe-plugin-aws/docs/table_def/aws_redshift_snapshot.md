<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>snapshot_identifier</td><td>The unique identifier of the cluster.</td></tr>
	<tr><td>cluster_identifier</td><td>The identifier of the cluster for which the snapshot was taken.</td></tr>
	<tr><td>snapshot_type</td><td>The snapshot type.</td></tr>
	<tr><td>actual_incremental_backup_size_in_mega_bytes</td><td>The size of the incremental backup.</td></tr>
	<tr><td>availability_zone</td><td>The Availability Zone in which the cluster was created.</td></tr>
	<tr><td>backup_progress_in_mega-bytes</td><td>The number of megabytes that have been transferred to the snapshot backup.</td></tr>
	<tr><td>cluster_create_time</td><td>The time (UTC) when the cluster was originally created.</td></tr>
	<tr><td>cluster_version</td><td>The version ID of the Amazon Redshift engine that is running on the cluster.</td></tr>
	<tr><td>current_backup_rate_in_mega_bytes_per_second</td><td>The number of megabytes per second being transferred to the snapshot backup.</td></tr>
	<tr><td>db_name</td><td>The name of the database that was created when the cluster was created.</td></tr>
	<tr><td>elapsed_time_in_seconds</td><td>The amount of time an in-progress snapshot backup has been running, or the amount of time it took a completed backup to finish.</td></tr>
	<tr><td>encrypted</td><td>If true, the data in the snapshot is encrypted at rest.</td></tr>
	<tr><td>encrypted_with_hsm</td><td>A boolean that indicates whether the snapshot data is encrypted using the HSM keys of the source cluster.</td></tr>
	<tr><td>engine_full_version</td><td>The cluster version of the cluster used to create the snapshot.</td></tr>
	<tr><td>enhanced_vpc_routing</td><td>An option that specifies whether to create the cluster with enhanced VPC routing enabled.</td></tr>
	<tr><td>estimated_seconds_to_completion</td><td>The estimate of the time remaining before the snapshot backup will complete.</td></tr>
	<tr><td>kms_key_id</td><td>The AWS KMS key ID of the encryption key that was used to encrypt data in the cluster from which the snapshot was taken.</td></tr>
	<tr><td>maintenance_track_name</td><td>The name of the maintenance track for the snapshot.</td></tr>
	<tr><td>manual_snapshot_remaining_days</td><td>The number of days until a manual snapshot will pass its retention period.</td></tr>
	<tr><td>manual_snapshot_retention_period</td><td>The number of days that a manual snapshot is retained.</td></tr>
	<tr><td>master_username</td><td>The master user name for the cluster.</td></tr>
	<tr><td>node_type</td><td>The node type of the nodes in the cluster.</td></tr>
	<tr><td>number_of_nodes</td><td>The number of nodes in the cluster.</td></tr>
	<tr><td>owner_account</td><td>The AWS customer account used to create or copy the snapshot.</td></tr>
	<tr><td>port</td><td>The port that the cluster is listening on.</td></tr>
	<tr><td>snapshot_create_time</td><td>The time (in UTC format) when Amazon Redshift began the snapshot.</td></tr>
	<tr><td>snapshot_retention_start_time</td><td>A timestamp representing the start of the retention period for the snapshot.</td></tr>
	<tr><td>source_region</td><td>The source region from which the snapshot was copied.</td></tr>
	<tr><td>status</td><td>The snapshot status.</td></tr>
	<tr><td>total_backup_size_in_mega_bytes</td><td>The size of the complete set of backup data that would be used to restore the cluster.</td></tr>
	<tr><td>vpc_id</td><td>The VPC identifier of the cluster if the snapshot is from a cluster in a VPC.</td></tr>
	<tr><td>accounts_with_restore_access</td><td>A list of the AWS customer accounts authorized to restore the snapshot.</td></tr>
	<tr><td>restorable_node_types</td><td>The list of node types that this cluster snapshot is able to restore into.</td></tr>
	<tr><td>tags_src</td><td>The list of tags for the cluster.</td></tr>
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