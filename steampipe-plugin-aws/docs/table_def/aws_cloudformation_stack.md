# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>id</td><td>Unique identifier of the stack.</td></tr>
	<tr><td>name</td><td>The name associated with the stack.</td></tr>
	<tr><td>status</td><td>Current status of the stack.</td></tr>
	<tr><td>creation_time</td><td>The time at which the stack was created.</td></tr>
	<tr><td>disable_rollback</td><td>Boolean to enable or disable rollback on stack creation failures.</td></tr>
	<tr><td>enable_termination_protection</td><td>Specifies whether termination protection is enabled for the stack.</td></tr>
	<tr><td>last_updated_time</td><td>The time the stack was last updated. This field will only be returned if the stack has been updated at least once.</td></tr>
	<tr><td>parent_id</td><td>ID of the direct parent of this stack.</td></tr>
	<tr><td>role_arn</td><td>The Amazon Resource Name (ARN) of an AWS Identity and Access Management (IAM) role that is associated with the stack.</td></tr>
	<tr><td>root_id</td><td>ID of the top-level stack to which the nested stack ultimately belongs.</td></tr>
	<tr><td>description</td><td>A user-defined description associated with the stack.</td></tr>
	<tr><td>notification_arns</td><td>SNS topic ARNs to which stack related events are published.</td></tr>
	<tr><td>outputs</td><td>A list of output structures.</td></tr>
	<tr><td>rollback_configuration</td><td>The rollback triggers for AWS CloudFormation to monitor during stack creation and updating operations, and for the specified monitoring period afterwards.</td></tr>
	<tr><td>capabilities</td><td>The capabilities allowed in the stack.</td></tr>
	<tr><td>stack_drift_status</td><td>Status of the stack&#39;s actual configuration compared to its expected template configuration.</td></tr>
	<tr><td>parameters</td><td>A list of Parameter structures.</td></tr>
	<tr><td>template_body</td><td>Structure containing the template body.</td></tr>
	<tr><td>template_body_json</td><td>Structure containing the template body. Parsed into json object for better readability.</td></tr>
	<tr><td>resources</td><td>A list of Stack resource structures.</td></tr>
	<tr><td>tags_src</td><td>A list of tags associated with stack.</td></tr>
	<tr><td>tags</td><td>A map of tags for the resource.</td></tr>
	<tr><td>title</td><td>Title of the resource.</td></tr>
	<tr><td>akas</td><td>Array of globally unique identifier strings (also known as) for the resource.</td></tr>
	<tr><td>partition</td><td>The AWS partition in which the resource is located (aws, aws-cn, or aws-us-gov).</td></tr>
	<tr><td>region</td><td>The AWS Region in which the resource is located.</td></tr>
	<tr><td>account_id</td><td>The AWS Account ID in which the resource is located.</td></tr>
	<tr><td>kaytu_account_id</td><td>The Kaytu Account ID in which the resource is located.</td></tr>
	<tr><td>kaytu_resource_id</td><td>The unique ID of the resource in Kaytu.</td></tr>
	<tr><td>kaytu_metadata</td><td>kaytu Metadata of the AWS resource.</td></tr>
</table>