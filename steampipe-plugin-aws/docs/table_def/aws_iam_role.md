# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>name</td><td>The friendly name that identifies the role.</td></tr>
	<tr><td>arn</td><td>The Amazon Resource Name (ARN) specifying the role.</td></tr>
	<tr><td>role_id</td><td>The stable and unique string identifying the role.</td></tr>
	<tr><td>create_date</td><td>The date and time when the role was created.</td></tr>
	<tr><td>description</td><td>A user-provided description of the role.</td></tr>
	<tr><td>instance_profile_arns</td><td>A list of instance profiles associated with the role.</td></tr>
	<tr><td>max_session_duration</td><td>The maximum session duration (in seconds) for the specified role. Anyone who uses the AWS CLI, or API to assume the role can specify the duration using the optional DurationSeconds API parameter or duration-seconds CLI parameter.</td></tr>
	<tr><td>path</td><td>The path to the role.</td></tr>
	<tr><td>permissions_boundary_arn</td><td>The ARN of the policy used to set the permissions boundary for the role.</td></tr>
	<tr><td>permissions_boundary_type</td><td>The permissions boundary usage type that indicates what type of IAM resource is used as the permissions boundary for an entity. This data type can only have a value of Policy.</td></tr>
	<tr><td>role_last_used_date</td><td>Contains information about the last time that an IAM role was used. Activity is only reported for the trailing 400 days. This period can be shorter if your Region began supporting these features within the last year. The role might have been used more than 400 days ago.</td></tr>
	<tr><td>role_last_used_region</td><td>Contains the region in which the IAM role was used.</td></tr>
	<tr><td>tags_src</td><td>A list of tags that are attached to the role.</td></tr>
	<tr><td>inline_policies</td><td>A list of policy documents that are embedded as inline policies for the role..</td></tr>
	<tr><td>inline_policies_std</td><td>Inline policies in canonical form for the role.</td></tr>
	<tr><td>attached_policy_arns</td><td>A list of managed policies attached to the role.</td></tr>
	<tr><td>assume_role_policy</td><td>The policy that grants an entity permission to assume the role.</td></tr>
	<tr><td>assume_role_policy_std</td><td>Contains the assume role policy in a canonical form for easier searching.</td></tr>
	<tr><td>title</td><td>Title of the resource.</td></tr>
	<tr><td>tags</td><td>A map of tags for the resource.</td></tr>
	<tr><td>akas</td><td>Array of globally unique identifier strings (also known as) for the resource.</td></tr>
	<tr><td>partition</td><td>The AWS partition in which the resource is located (aws, aws-cn, or aws-us-gov).</td></tr>
	<tr><td>region</td><td>The AWS Region in which the resource is located.</td></tr>
	<tr><td>account_id</td><td>The AWS Account ID in which the resource is located.</td></tr>
	<tr><td>kaytu_account_id</td><td>The Kaytu Account ID in which the resource is located.</td></tr>
	<tr><td>kaytu_resource_id</td><td>The unique ID of the resource in opengovernance.</td></tr>
	<tr><td>kaytu_metadata</td><td>kaytu Metadata of the AWS resource.</td></tr>
</table>