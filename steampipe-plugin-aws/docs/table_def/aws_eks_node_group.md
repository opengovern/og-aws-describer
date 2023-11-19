# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>nodegroup_name</td><td>The name associated with an Amazon EKS managed node group.</td></tr>
	<tr><td>arn</td><td>The Amazon Resource Name (ARN) associated with the managed node group.</td></tr>
	<tr><td>cluster_name</td><td>The name of the cluster that the managed node group resides in.</td></tr>
	<tr><td>created_at</td><td>The Unix epoch timestamp in seconds for when the managed node group was created.</td></tr>
	<tr><td>status</td><td>The current status of the managed node group.</td></tr>
	<tr><td>ami_type</td><td>The AMI type that was specified in the node group configuration.</td></tr>
	<tr><td>capacity_type</td><td>The capacity type of your managed node group.</td></tr>
	<tr><td>disk_size</td><td>The disk size in the node group configuration.</td></tr>
	<tr><td>modified_at</td><td>The Unix epoch timestamp in seconds for when the managed node group was last modified.</td></tr>
	<tr><td>node_role</td><td>The IAM role associated with your node group.</td></tr>
	<tr><td>release_version</td><td>If the node group was deployed using a launch template with a custom AMI, then this is the AMI ID that was specified in the launch template. For node groups that weren&#39;t deployed using a launch template, this is the version of the Amazon EKS optimized AMI that the node group was deployed with.</td></tr>
	<tr><td>version</td><td>The Kubernetes version of the managed node group.</td></tr>
	<tr><td>health</td><td>The health status of the node group.</td></tr>
	<tr><td>instance_types</td><td>The instance type that is associated with the node group. If the node group was deployed with a launch template, then this is null.</td></tr>
	<tr><td>labels</td><td>The Kubernetes labels applied to the nodes in the node group.</td></tr>
	<tr><td>launch_template</td><td>If a launch template was used to create the node group, then this is the launch template that was used.</td></tr>
	<tr><td>remote_access</td><td>The remote access configuration that is associated with the node group. If the node group was deployed with a launch template, then this is null.</td></tr>
	<tr><td>resources</td><td>The resources associated with the node group, such as Auto Scaling groups and security groups for remote access.</td></tr>
	<tr><td>scaling_config</td><td>The scaling configuration details for the Auto Scaling group that is associated with your node group.</td></tr>
	<tr><td>subnets</td><td>The subnets that were specified for the Auto Scaling group that is associated with your node group.</td></tr>
	<tr><td>tags</td><td>A map of tags for the resource.</td></tr>
	<tr><td>taints</td><td>The Kubernetes taints to be applied to the nodes in the node group when they are created.</td></tr>
	<tr><td>update_config</td><td>The node group update configuration.</td></tr>
	<tr><td>title</td><td>Title of the resource.</td></tr>
	<tr><td>akas</td><td>Array of globally unique identifier strings (also known as) for the resource.</td></tr>
	<tr><td>partition</td><td>The AWS partition in which the resource is located (aws, aws-cn, or aws-us-gov).</td></tr>
	<tr><td>region</td><td>The AWS Region in which the resource is located.</td></tr>
	<tr><td>account_id</td><td>The AWS Account ID in which the resource is located.</td></tr>
	<tr><td>kaytu_account_id</td><td>The Kaytu Account ID in which the resource is located.</td></tr>
	<tr><td>kaytu_resource_id</td><td>The unique ID of the resource in Kaytu.</td></tr>
	<tr><td>kaytu_metadata</td><td>kaytu Metadata of the AWS resource.</td></tr>
</table>