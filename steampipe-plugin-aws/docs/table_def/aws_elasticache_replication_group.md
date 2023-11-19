# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>replication_group_id</td><td>The identifier for the replication group.</td></tr>
	<tr><td>arn</td><td>The ARN (Amazon Resource Name) of the replication group.</td></tr>
	<tr><td>description</td><td>The user supplied description of the replication group.</td></tr>
	<tr><td>at_rest_encryption_enabled</td><td>A flag that enables encryption at-rest when set to true.</td></tr>
	<tr><td>kms_key_id</td><td>The ID of the KMS key used to encrypt the disk in the cluster.</td></tr>
	<tr><td>auth_token_enabled</td><td>A flag that enables using an AuthToken (password) when issuing Redis commands.</td></tr>
	<tr><td>auth_token_last_modified_date</td><td>The date when the auth token was last modified.</td></tr>
	<tr><td>automatic_failover</td><td>Indicates the status of automatic failover for this Redis replication group.</td></tr>
	<tr><td>cache_node_type</td><td>The name of the compute and memory capacity node type for each node in the replication group.</td></tr>
	<tr><td>cluster_enabled</td><td>A flag indicating whether or not this replication group is cluster enabled.</td></tr>
	<tr><td>multi_az</td><td>A flag indicating if you have Multi-AZ enabled to enhance fault tolerance.</td></tr>
	<tr><td>snapshot_retention_limit</td><td>The number of days for which ElastiCache retains automatic cluster snapshots before deleting them.</td></tr>
	<tr><td>snapshot_window</td><td>The daily time range (in UTC) during which ElastiCache begins taking a daily snapshot of your node group (shard).</td></tr>
	<tr><td>snapshotting_cluster_id</td><td>The cluster ID that is used as the daily snapshot source for the replication group.</td></tr>
	<tr><td>status</td><td>The current state of this replication group - creating, available, modifying, deleting, create-failed, snapshotting.</td></tr>
	<tr><td>transit_encryption_enabled</td><td>A flag that enables in-transit encryption when set to true.</td></tr>
	<tr><td>configuration_endpoint</td><td>The configuration endpoint for this replication group.</td></tr>
	<tr><td>global_replication_group_info</td><td>The name of the Global Datastore and role of this replication group in the Global Datastore.</td></tr>
	<tr><td>member_clusters</td><td>The names of all the cache clusters that are part of this replication group.</td></tr>
	<tr><td>member_clusters_outpost_arns</td><td>The outpost ARNs of the replication group's member clusters.</td></tr>
	<tr><td>node_groups</td><td>A list of node groups in this replication group.</td></tr>
	<tr><td>pending_modified_values</td><td>A group of settings to be applied to the replication group, either immediately or during the next maintenance window.</td></tr>
	<tr><td>user_group_ids</td><td>The list of user group IDs that have access to the replication group.</td></tr>
	<tr><td>title</td><td>Title of the resource.</td></tr>
	<tr><td>akas</td><td>Array of globally unique identifier strings (also known as) for the resource.</td></tr>
	<tr><td>partition</td><td>The AWS partition in which the resource is located (aws, aws-cn, or aws-us-gov).</td></tr>
	<tr><td>region</td><td>The AWS Region in which the resource is located.</td></tr>
	<tr><td>account_id</td><td>The AWS Account ID in which the resource is located.</td></tr>
	<tr><td>kaytu_account_id</td><td>The Kaytu Account ID in which the resource is located.</td></tr>
	<tr><td>kaytu_resource_id</td><td>The unique ID of the resource in Kaytu.</td></tr>
	<tr><td>kaytu_metadata</td><td>kaytu Metadata of the AWS resource.</td></tr>
</table>