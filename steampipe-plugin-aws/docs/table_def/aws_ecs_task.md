# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>task_arn</td><td>The Amazon Resource Name (ARN) of the task.</td></tr>
	<tr><td>container_instance_arn</td><td>The ARN of the container instances that host the task.</td></tr>
	<tr><td>cluster_name</td><td>A user-generated string that you use to identify your cluster.</td></tr>
	<tr><td>desired_status</td><td>The desired status of the task.</td></tr>
	<tr><td>launch_type</td><td>The infrastructure on which your task is running.</td></tr>
	<tr><td>availability_zone</td><td>The availability zone of the task.</td></tr>
	<tr><td>capacity_provider_name</td><td>The capacity provider associated with the task.</td></tr>
	<tr><td>cluster_arn</td><td>The ARN of the cluster that hosts the task.</td></tr>
	<tr><td>connectivity</td><td>The connectivity status of a task.</td></tr>
	<tr><td>connectivity_at</td><td>The Unix timestamp for when the task last went into CONNECTED status.</td></tr>
	<tr><td>cpu</td><td>The number of CPU units used by the task as expressed in a task definition.</td></tr>
	<tr><td>created_at</td><td>The Unix timestamp for when the task was created.</td></tr>
	<tr><td>enable_execute_command</td><td>Whether or not execute command functionality is enabled for this task. If true, this enables execute command functionality on all containers in the task.</td></tr>
	<tr><td>execution_stopped_at</td><td>The Unix timestamp for when the task execution stopped.</td></tr>
	<tr><td>group</td><td>The name of the task group associated with the task.</td></tr>
	<tr><td>health_status</td><td>The health status for the task, which is determined by the health of the essential containers in the task. If all essential containers in the task are reporting as HEALTHY, then the task status also reports as HEALTHY.</td></tr>
	<tr><td>last_status</td><td>The last known status of the task.</td></tr>
	<tr><td>memory</td><td>The amount of memory (in MiB) used by the task as expressed in a task definition.</td></tr>
	<tr><td>platform_version</td><td>The platform version on which your task is running.</td></tr>
	<tr><td>pull_started_at</td><td>The Unix timestamp for when the container image pull began.</td></tr>
	<tr><td>pull_stopped_at</td><td>The Unix timestamp for when the container image pull completed.</td></tr>
	<tr><td>service_name</td><td>The name of the service.</td></tr>
	<tr><td>started_at</td><td>The Unix timestamp for when the task started.</td></tr>
	<tr><td>started_by</td><td>The tag specified when a task is started.</td></tr>
	<tr><td>stop_code</td><td>The stop code indicating why a task was stopped.</td></tr>
	<tr><td>stopped_at</td><td>The Unix timestamp for when the task was stopped.</td></tr>
	<tr><td>stopped_reason</td><td>The reason that the task was stopped.</td></tr>
	<tr><td>stopping_at</td><td>The Unix timestamp for when the task stops.</td></tr>
	<tr><td>task_definition_arn</td><td>The ARN of the task definition that creates the task.</td></tr>
	<tr><td>version</td><td>The version counter for the task.</td></tr>
	<tr><td>attachments</td><td>The Elastic Network Adapter associated with the task if the task uses the awsvpc network mode.</td></tr>
	<tr><td>attributes</td><td>The attributes of the task.</td></tr>
	<tr><td>containers</td><td>The containers associated with the task.</td></tr>
	<tr><td>ephemeral_storage</td><td>The ephemeral storage settings for the task.</td></tr>
	<tr><td>inference_accelerators</td><td>The Elastic Inference accelerator associated with the task.</td></tr>
	<tr><td>overrides</td><td>One or more container overrides.</td></tr>
	<tr><td>protection</td><td>Protection status of task in an Amazon ECS service.</td></tr>
	<tr><td>tags_src</td><td>A list of tags associated with task.</td></tr>
	<tr><td>tags</td><td>A map of tags for the resource.</td></tr>
	<tr><td>akas</td><td>Array of globally unique identifier strings (also known as) for the resource.</td></tr>
	<tr><td>partition</td><td>The AWS partition in which the resource is located (aws, aws-cn, or aws-us-gov).</td></tr>
	<tr><td>region</td><td>The AWS Region in which the resource is located.</td></tr>
	<tr><td>account_id</td><td>The AWS Account ID in which the resource is located.</td></tr>
	<tr><td>kaytu_account_id</td><td>The Kaytu Account ID in which the resource is located.</td></tr>
	<tr><td>kaytu_resource_id</td><td>The unique ID of the resource in opengovernance.</td></tr>
	<tr><td>kaytu_metadata</td><td>kaytu Metadata of the AWS resource.</td></tr>
</table>