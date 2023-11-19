<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>service_namespace</td><td>The namespace of the AWS service that provides the resource, or a custom-resource.</td></tr>
	<tr><td>resource_id</td><td>The identifier of the resource associated with the scalable target.</td></tr>
	<tr><td>scalable_dimension</td><td>The scalable dimension associated with the scalable target. This string consists of the service namespace, resource type, and scaling property.</td></tr>
	<tr><td>creation_time</td><td>The Unix timestamp for when the scalable target was created.</td></tr>
	<tr><td>min_capacity</td><td>The minimum value to scale to in response to a scale-in activity.</td></tr>
	<tr><td>max_capacity</td><td>The maximum value to scale to in response to a scale-out activity.</td></tr>
	<tr><td>role_arn</td><td>The ARN of an IAM role that allows Application Auto Scaling to modify the scalable target on your behalf.</td></tr>
	<tr><td>suspended_state</td><td>Specifies whether the scaling activities for a scalable target are in a suspended state.</td></tr>
	<tr><td>title</td><td>Title of the resource.</td></tr>
	<tr><td>partition</td><td>The AWS partition in which the resource is located (aws, aws-cn, or aws-us-gov).</td></tr>
	<tr><td>region</td><td>The AWS Region in which the resource is located.</td></tr>
	<tr><td>account_id</td><td>The AWS Account ID in which the resource is located.</td></tr>
	<tr><td>kaytu_account_id</td><td>The Kaytu Account ID in which the resource is located.</td></tr>
	<tr><td>kaytu_resource_id</td><td>The unique ID of the resource in Kaytu.</td></tr>
	<tr><td>kaytu_metadata</td><td>kaytu Metadata of the AWS resource.</td></tr>
</table>