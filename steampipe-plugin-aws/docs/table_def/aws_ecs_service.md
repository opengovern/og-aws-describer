# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>service_name</td><td>The name of the service.</td></tr>
	<tr><td>arn</td><td>The Amazon Resource Name (ARN) specifying the service.</td></tr>
	<tr><td>status</td><td>The status of the service. Valid values are: ACTIVE, DRAINING, or INACTIVE.</td></tr>
	<tr><td>cluster_arn</td><td>The Amazon Resource Name (ARN) of the cluster that hosts the service.</td></tr>
	<tr><td>task_definition</td><td>The task definition to use for tasks in the service.</td></tr>
	<tr><td>created_at</td><td>The date and time when the service was created.</td></tr>
	<tr><td>created_by</td><td>The principal that created the service.</td></tr>
	<tr><td>deployment_controller_type</td><td>The deployment controller type to use. Possible values are: ECS, CODE_DEPLOY, and EXTERNAL.</td></tr>
	<tr><td>desired_count</td><td>The desired number of instantiations of the task definition to keep running on the service.</td></tr>
	<tr><td>enable_ecs_managed_tags</td><td>Specifies whether to enable Amazon ECS managed tags for the tasks in the service.</td></tr>
	<tr><td>enable_execute_command</td><td>Indicates whether or not the execute command functionality is enabled for the service.</td></tr>
	<tr><td>health_check_grace_period_seconds</td><td>The period of time, in seconds, that the Amazon ECS service scheduler ignores unhealthy Elastic Load Balancing target health checks after a task has first started.</td></tr>
	<tr><td>launch_type</td><td>The launch type on which your service is running. If no value is specified, it will default to EC2.</td></tr>
	<tr><td>pending_count</td><td>The number of tasks in the cluster that are in the PENDING state.</td></tr>
	<tr><td>platform_family</td><td>The operating system that your tasks in the service run on.</td></tr>
	<tr><td>platform_version</td><td>The platform version on which to run your service.</td></tr>
	<tr><td>propagate_tags</td><td>Specifies whether to propagate the tags from the task definition or the service to the task. If no value is specified, the tags are not propagated.</td></tr>
	<tr><td>role_arn</td><td>The ARN of the IAM role associated with the service that allows the Amazon ECS container agent to register container instances with an Elastic Load Balancing load balancer.</td></tr>
	<tr><td>running_count</td><td>The number of tasks in the cluster that are in the RUNNING state.</td></tr>
	<tr><td>scheduling_strategy</td><td>The scheduling strategy to use for the service.</td></tr>
	<tr><td>capacity_provider_strategy</td><td>The capacity provider strategy associated with the service.</td></tr>
	<tr><td>deployment_configuration</td><td>Optional deployment parameters that control how many tasks run during the deployment and the ordering of stopping and starting tasks.</td></tr>
	<tr><td>deployments</td><td>The current state of deployments for the service.</td></tr>
	<tr><td>events</td><td>The event stream for your service. A maximum of 100 of the latest events are displayed.</td></tr>
	<tr><td>load_balancers</td><td>A list of Elastic Load Balancing load balancer objects, containing the load balancer name, the container name (as it appears in a container definition), and the container port to access from the load balancer.</td></tr>
	<tr><td>network_configuration</td><td>The VPC subnet and security group configuration for tasks that receive their own elastic network interface by using the awsvpc networking mode.</td></tr>
	<tr><td>placement_constraints</td><td>The placement constraints for the tasks in the service.</td></tr>
	<tr><td>placement_strategy</td><td>The placement strategy that determines how tasks for the service are placed.</td></tr>
	<tr><td>service_registries</td><td>The details of the service discovery registries to assign to this service.</td></tr>
	<tr><td>task_sets</td><td>Information about a set of Amazon ECS tasks in either an AWS CodeDeploy or an EXTERNAL deployment.</td></tr>
	<tr><td>tags_src</td><td>The metadata that you apply to the service to help you categorize and organize them.</td></tr>
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