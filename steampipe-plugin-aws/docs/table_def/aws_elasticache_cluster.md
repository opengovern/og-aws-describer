# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>cache_cluster_id</td><td>An unique identifier for ElastiCache cluster.</td></tr>
	<tr><td>arn</td><td>The ARN (Amazon Resource Name) of the cache cluster.</td></tr>
	<tr><td>cache_node_type</td><td>The name of the compute and memory capacity node type for the cluster.</td></tr>
	<tr><td>cache_cluster_status</td><td>The current state of this cluster, one of the following values: available, creating, deleted, deleting, incompatible-network, modifying, rebooting cluster nodes, restore-failed, or snapshotting.</td></tr>
	<tr><td>at_rest_encryption_enabled</td><td>A flag that enables encryption at-rest when set to true.</td></tr>
	<tr><td>auth_token_enabled</td><td>A flag that enables using an AuthToken (password) when issuing Redis commands.</td></tr>
	<tr><td>auto_minor_version_upgrade</td><td>This parameter is currently disabled.</td></tr>
	<tr><td>cache_cluster_create_time</td><td>The date and time when the cluster was created.</td></tr>
	<tr><td>cache_subnet_group_name</td><td>The name of the cache subnet group associated with the cluster.</td></tr>
	<tr><td>client_download_landing_page</td><td>The URL of the web page where you can download the latest ElastiCache client library.</td></tr>
	<tr><td>configuration_endpoint</td><td>Represents a Memcached cluster endpoint which can be used by an application to connect to any node in the cluster.</td></tr>
	<tr><td>engine</td><td>The name of the cache engine (memcached or redis) to be used for this cluster.</td></tr>
	<tr><td>engine_version</td><td>The version of the cache engine that is used in this cluster.</td></tr>
	<tr><td>num_cache_nodes</td><td>The number of cache nodes in the cluster.</td></tr>
	<tr><td>preferred_availability_zone</td><td>The name of the Availability Zone in which the cluster is located or 'Multiple' if the cache nodes are located in different Availability Zones.</td></tr>
	<tr><td>preferred_maintenance_window</td><td>Specifies the weekly time range during which maintenance on the cluster is performed.</td></tr>
	<tr><td>replication_group_id</td><td>The replication group to which this cluster belongs.</td></tr>
	<tr><td>snapshot_retention_limit</td><td>The number of days for which ElastiCache retains automatic cluster snapshots before deleting them.</td></tr>
	<tr><td>snapshot_window</td><td>The daily time range (in UTC) during which ElastiCache begins taking a daily snapshot of your cluster.</td></tr>
	<tr><td>transit_encryption_enabled</td><td>A flag that enables in-transit encryption when set to true.</td></tr>
	<tr><td>cache_parameter_group</td><td>Status of the cache parameter group.</td></tr>
	<tr><td>notification_configuration</td><td>Describes a notification topic and its status.</td></tr>
	<tr><td>pending_modified_values</td><td>A group of settings that are applied to the cluster in the future, or that are currently being applied.</td></tr>
	<tr><td>security_groups</td><td>A list of VPC Security Groups associated with the cluster.</td></tr>
	<tr><td>tags_src</td><td>A list of tags associated with the cluster.</td></tr>
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