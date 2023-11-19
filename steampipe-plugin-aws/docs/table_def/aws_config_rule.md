<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>name</td><td>The name that you assign to the AWS Config rule.</td></tr>
	<tr><td>rule_id</td><td>The ID of the AWS Config rule.</td></tr>
	<tr><td>arn</td><td>The Amazon Resource Name (ARN) of the AWS Config rule.</td></tr>
	<tr><td>rule_state</td><td>It indicate the evaluation status for the AWS Config rule.</td></tr>
	<tr><td>created_by</td><td>Service principal name of the service that created the rule.</td></tr>
	<tr><td>description</td><td>The description that you provide for the AWS Config rule.</td></tr>
	<tr><td>maximum_execution_frequency</td><td>The maximum frequency with which AWS Config runs evaluations for a rule.</td></tr>
	<tr><td>compliance_by_config_rule</td><td>The compliance information of the config rule.</td></tr>
	<tr><td>evaluation_modes</td><td>The modes the Config rule can be evaluated in. The valid values are distinct objects. By default, the value is Detective evaluation mode only.</td></tr>
	<tr><td>input_parameters</td><td>A string, in JSON format, that is passed to the AWS Config rule Lambda function.</td></tr>
	<tr><td>scope</td><td>Defines which resources can trigger an evaluation for the rule. The scope can include one or more resource types, a combination of one resource type and one resource ID, or a combination of a tag key and value. Specify a scope to constrain the resources that can trigger an evaluation for the rule. If you do not specify a scope, evaluations are triggered when any resource in the recording group changes.</td></tr>
	<tr><td>source</td><td>Provides the rule owner (AWS or customer), the rule identifier, and the notifications that cause the function to evaluate your AWS resources.</td></tr>
	<tr><td>tags_src</td><td>A list of tags assigned to the rule.</td></tr>
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