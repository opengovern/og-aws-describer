# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>cluster_arn</td><td>The Amazon Resource Name (ARN) that identifies the cluster.</td></tr>
	<tr><td>cluster_name</td><td>A user-generated string that you use to identify your cluster.</td></tr>
	<tr><td>active_services_count</td><td>The number of services that are running on the cluster in an ACTIVE state.</td></tr>
	<tr><td>pending_tasks_count</td><td>The number of tasks in the cluster that are in the PENDING state.</td></tr>
	<tr><td>registered_container_instances_count</td><td>The number of container instances registered into the cluster. This includes container instances in both ACTIVE and DRAINING status.</td></tr>
	<tr><td>running_tasks_count</td><td>The number of tasks in the cluster that are in the RUNNING state.</td></tr>
	<tr><td>status</td><td>The status of the cluster.</td></tr>
	<tr><td>attachments_status</td><td>The status of the capacity providers associated with the cluster.</td></tr>
	<tr><td>attachments</td><td>The resources attached to a cluster. When using a capacity provider with a cluster, the Auto Scaling plan that is created will be returned as a cluster attachment.</td></tr>
	<tr><td>capacity_providers</td><td>The capacity providers associated with the cluster.</td></tr>
	<tr><td>default_capacity_provider_strategy</td><td>The default capacity provider strategy for the cluster.</td></tr>
	<tr><td>settings</td><td>The settings for the cluster. This parameter indicates whether CloudWatch Container Insights is enabled or disabled for a cluster.</td></tr>
	<tr><td>statistics</td><td>Additional information about your clusters that are separated by launch type.</td></tr>
	<tr><td>tags_src</td><td>A list of tags attached to the cluster.</td></tr>
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