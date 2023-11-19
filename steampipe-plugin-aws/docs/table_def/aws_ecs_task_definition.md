<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>task_definition_arn</td><td>The Amazon Resource Name (ARN) that identifies the task definition.</td></tr>
	<tr><td>cpu</td><td>The number of cpu units used by the task.</td></tr>
	<tr><td>status</td><td>The status of the task definition.</td></tr>
	<tr><td>execution_role_arn</td><td>The Amazon Resource Name (ARN) of the task execution role that grants the Amazon ECS container agent permission to make AWS API calls on your behalf.</td></tr>
	<tr><td>family</td><td>The name of a family that this task definition is registered to.</td></tr>
	<tr><td>ipc_mode</td><td>The IPC resource namespace to use for the containers in the task.</td></tr>
	<tr><td>memory</td><td>The amount (in MiB) of memory used by the task.</td></tr>
	<tr><td>network_mode</td><td>The Docker networking mode to use for the containers in the task.</td></tr>
	<tr><td>pid_mode</td><td>The process namespace to use for the containers in the task.</td></tr>
	<tr><td>revision</td><td>The revision of the task in a particular family.</td></tr>
	<tr><td>task_role_arn</td><td>The short name or full Amazon Resource Name (ARN) of the AWS Identity and Access Management (IAM) role that grants containers in the task permission to call AWS APIs on your behalf.</td></tr>
	<tr><td>registered_at</td><td>The Unix timestamp for when the task definition was registered.</td></tr>
	<tr><td>registered_by</td><td>The principal that registered the task definition.</td></tr>
	<tr><td>container_definitions</td><td>A list of container definitions in JSON format that describe the different containers that make up your task.</td></tr>
	<tr><td>compatibilities</td><td>The launch type to use with your task.</td></tr>
	<tr><td>inference_accelerators</td><td>The Elastic Inference accelerator associated with the task.</td></tr>
	<tr><td>placement_constraints</td><td>An array of placement constraint objects to use for tasks.</td></tr>
	<tr><td>proxy_configuration</td><td>The configuration details for the App Mesh proxy.</td></tr>
	<tr><td>requires_attributes</td><td>The container instance attributes required by your task.</td></tr>
	<tr><td>requires_compatibilities</td><td>The launch type the task requires. If no value is specified, it will default to EC2. Valid values include EC2 and FARGATE.</td></tr>
	<tr><td>volumes</td><td>The list of volume definitions for the task.</td></tr>
	<tr><td>tags_src</td><td>A list of tags associated with task.</td></tr>
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