# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>name</td><td>The friendly name that identifies the iam policy.</td></tr>
	<tr><td>policy_id</td><td>The stable and unique string identifying the policy.</td></tr>
	<tr><td>path</td><td>The path to the policy.</td></tr>
	<tr><td>arn</td><td>The Amazon Resource Name (ARN) specifying the iam policy.</td></tr>
	<tr><td>is_aws_managed</td><td>Specifies whether the policy is AWS Managed or Customer Managed. If true policy is aws managed otherwise customer managed.</td></tr>
	<tr><td>is_attachable</td><td>Specifies whether the policy can be attached to an IAM user, group, or role.</td></tr>
	<tr><td>create_date</td><td>The date and time, when the policy was created.</td></tr>
	<tr><td>update_date</td><td>The date and time, when the policy was last updated.</td></tr>
	<tr><td>attachment_count</td><td>The number of entities (users, groups, and roles) that the policy is attached to.</td></tr>
	<tr><td>is_attached</td><td>Specifies whether the policy is attached to at least one IAM user, group, or role.</td></tr>
	<tr><td>default_version_id</td><td>The identifier for the version of the policy that is set as the default version.</td></tr>
	<tr><td>permissions_boundary_usage_count</td><td>The number of entities (users and roles) for which the policy is used to set the permissions boundary.</td></tr>
	<tr><td>policy</td><td>Contains the details about the policy.</td></tr>
	<tr><td>policy_std</td><td>Contains the policy in a canonical form for easier searching.</td></tr>
	<tr><td>tags_src</td><td>A list of tags attached with the IAM policy.</td></tr>
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