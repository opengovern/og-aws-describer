# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>name</td><td>The friendly name identifying the user.</td></tr>
	<tr><td>user_id</td><td>The stable and unique string identifying the user.</td></tr>
	<tr><td>path</td><td>The path to the user.</td></tr>
	<tr><td>arn</td><td>The Amazon Resource Name (ARN) that identifies the user.</td></tr>
	<tr><td>create_date</td><td>The date and time, when the user was created.</td></tr>
	<tr><td>password_last_used</td><td>The date and time, when the user's password was last used to sign in to an AWS website.</td></tr>
	<tr><td>permissions_boundary_arn</td><td>The ARN of the policy used to set the permissions boundary for the user.</td></tr>
	<tr><td>permissions_boundary_type</td><td>The permissions boundary usage type that indicates what type of IAM resource is used as the permissions boundary for an entity. This data type can only have a value of Policy.</td></tr>
	<tr><td>mfa_enabled</td><td>The MFA status of the user.</td></tr>
	<tr><td>login_profile</td><td>Contains the user name and password create date for a user.</td></tr>
	<tr><td>mfa_devices</td><td>A list of MFA devices attached to the user.</td></tr>
	<tr><td>groups</td><td>A list of groups attached to the user.</td></tr>
	<tr><td>inline_policies</td><td>A list of policy documents that are embedded as inline policies for the user.</td></tr>
	<tr><td>inline_policies_std</td><td>Inline policies in canonical form for the user.</td></tr>
	<tr><td>attached_policy_arns</td><td>A list of managed policies attached to the user.</td></tr>
	<tr><td>tags_src</td><td>A list of tags that are attached to the user.</td></tr>
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