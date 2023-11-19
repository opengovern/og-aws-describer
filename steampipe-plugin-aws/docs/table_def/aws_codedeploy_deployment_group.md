# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>arn</td><td>The Amazon Resource Name (ARN) of the deployment group</td></tr>
	<tr><td>application_name</td><td>The application name.</td></tr>
	<tr><td>deployment_config_name</td><td>The deployment configuration name.</td></tr>
	<tr><td>deployment_group_id</td><td>The deployment group ID.</td></tr>
	<tr><td>deployment_group_name</td><td>The name of the deployment group.</td></tr>
	<tr><td>service_role_arn</td><td>A service role Amazon Resource Name (ARN) that grants CodeDeploy permission to make calls to Amazon Web Services services on your behalf.</td></tr>
	<tr><td>alarm_configuration</td><td>A list of alarms associated with the deployment group.</td></tr>
	<tr><td>auto_rollback_configuration</td><td>Information about the automatic rollback configuration associated with the deployment group.</td></tr>
	<tr><td>auto_scaling_groups</td><td>A list of associated Auto Scaling groups.</td></tr>
	<tr><td>blue_green_deployment_configuration</td><td>Information about blue/green deployment options for a deployment group.</td></tr>
	<tr><td>compute_platform</td><td>The destination platform type for the deployment (Lambda, Server, or ECS).</td></tr>
	<tr><td>deployment_style</td><td>Information about the type of deployment, either in-place or blue/green, you want to run and whether to route deployment traffic behind a load balancer.</td></tr>
	<tr><td>ec2_tag_filters</td><td>The Amazon EC2 tags on which to filter. The deployment group includes EC2 instances with any of the specified tags.</td></tr>
	<tr><td>ec2_tag_set</td><td>Information about groups of tags applied to an Amazon EC2 instance.</td></tr>
	<tr><td>ecs_services</td><td>The target Amazon ECS services in the deployment group.</td></tr>
	<tr><td>last_attempted_deployment</td><td>Information about the most recent attempted deployment to the deployment group.</td></tr>
	<tr><td>last_successful_deployment</td><td>Information about the most recent successful deployment to the deployment group.</td></tr>
	<tr><td>load_balancer_info</td><td>Information about the load balancer to use in a deployment.</td></tr>
	<tr><td>on_premises_instance_tag_filters</td><td>The on-premises instance tags on which to filter.</td></tr>
	<tr><td>on_premises_tag_set</td><td>Information about groups of tags applied to an on-premises instance.</td></tr>
	<tr><td>outdated_instances_strategy</td><td>Indicates what happens when new Amazon EC2 instances are launched mid-deployment and do not receive the deployed application revision.</td></tr>
	<tr><td>tags_src</td><td>A list of tags associated with deployment group.</td></tr>
	<tr><td>target_revision</td><td>Information about the deployment group&#39;s target revision, including type and location.</td></tr>
	<tr><td>trigger_configurations</td><td>Information about triggers associated with the deployment group.</td></tr>
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