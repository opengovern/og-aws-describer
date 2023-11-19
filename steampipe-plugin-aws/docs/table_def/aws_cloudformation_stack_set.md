# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>stack_set_id</td><td>The ID of the stack set.</td></tr>
	<tr><td>stack_set_name</td><td>The name of the stack set.</td></tr>
	<tr><td>arn</td><td>The Amazon Resource Name (ARN) of the stack set</td></tr>
	<tr><td>status</td><td>The status of the stack set.</td></tr>
	<tr><td>description</td><td>A description of the stack set that you specify when the stack set is created or updated.</td></tr>
	<tr><td>drift_status</td><td>Status of the stack set's actual configuration compared to its expected template and parameter configuration. A stack set is considered to have drifted if one or more of its stack instances have drifted from their expected template and parameter configuration.</td></tr>
	<tr><td>last_drift_check_timestamp</td><td>Most recent time when CloudFormation performed a drift detection operation on the stack set.</td></tr>
	<tr><td>permission_model</td><td>Describes how the IAM roles required for stack set operations are created.</td></tr>
	<tr><td>tags</td><td>A map of tags for the resource.</td></tr>
	<tr><td>akas</td><td>Array of globally unique identifier strings (also known as) for the resource.</td></tr>
	<tr><td>partition</td><td>The AWS partition in which the resource is located (aws, aws-cn, or aws-us-gov).</td></tr>
	<tr><td>region</td><td>The AWS Region in which the resource is located.</td></tr>
	<tr><td>account_id</td><td>The AWS Account ID in which the resource is located.</td></tr>
	<tr><td>kaytu_account_id</td><td>The Kaytu Account ID in which the resource is located.</td></tr>
	<tr><td>kaytu_resource_id</td><td>The unique ID of the resource in Kaytu.</td></tr>
	<tr><td>kaytu_metadata</td><td>kaytu Metadata of the AWS resource.</td></tr>
</table>