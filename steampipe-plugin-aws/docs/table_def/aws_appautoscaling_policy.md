# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>policy_arn</td><td>The Amazon Resource Name (ARN) of the appautoscaling policy.</td></tr>
	<tr><td>policy_name</td><td>The name of the scaling policy.</td></tr>
	<tr><td>service_namespace</td><td>The namespace of the AWS service that provides the resource, or a custom-resource.</td></tr>
	<tr><td>resource_id</td><td>The identifier of the resource associated with the scaling policy.</td></tr>
	<tr><td>scalable_dimension</td><td>The scalable dimension associated with the scaling policy. This string consists of the service namespace, resource type, and scaling property.</td></tr>
	<tr><td>policy_type</td><td>The policy type. Currently supported values are TargetTrackingScaling and StepScaling</td></tr>
	<tr><td>target_tracking_scaling_policy_configuration</td><td>The target tracking scaling policy configuration (if policy type is TargetTrackingScaling).</td></tr>
	<tr><td>step_scaling_policy_configuration</td><td>The step tracking scaling policy configuration (if policy type is StepScaling).</td></tr>
	<tr><td>alarms</td><td>The CloudWatch alarms associated with the scaling policy.</td></tr>
	<tr><td>creation_time</td><td>The Unix timestamp for when the scaling policy was created.</td></tr>
	<tr><td>title</td><td>Title of the resource.</td></tr>
	<tr><td>partition</td><td>The AWS partition in which the resource is located (aws, aws-cn, or aws-us-gov).</td></tr>
	<tr><td>region</td><td>The AWS Region in which the resource is located.</td></tr>
	<tr><td>account_id</td><td>The AWS Account ID in which the resource is located.</td></tr>
</table>