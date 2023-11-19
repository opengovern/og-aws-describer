# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>id</td><td>The domain ID.</td></tr>
	<tr><td>name</td><td>The domain name.</td></tr>
	<tr><td>arn</td><td>The Amazon Resource Name (ARN) of the domain.</td></tr>
	<tr><td>creation_time</td><td>A timestamp that indicates when the domain was created.</td></tr>
	<tr><td>app_network_access_type</td><td>Specifies the VPC used for non-EFS traffic.</td></tr>
	<tr><td>app_security_group_management</td><td>The entity that creates and manages the required security groups for inter-app communication in VPCOnly mode.</td></tr>
	<tr><td>auth_mode</td><td>The domain's authentication mode.</td></tr>
	<tr><td>failure_reason</td><td>The domain's failure reason.</td></tr>
	<tr><td>home_efs_file_system_id</td><td>The ID of the Amazon Elastic File System (EFS) managed by this domain.</td></tr>
	<tr><td>kms_key_id</td><td>The Amazon Web Services KMS customer managed key used to encrypt the EFS volume attached to the domain.</td></tr>
	<tr><td>last_modified_time</td><td>The domain's last modified time.</td></tr>
	<tr><td>security_group_id_for_domain_boundary</td><td>The ID of the security group that authorizes traffic between the RSessionGateway apps and the RStudioServerPro app.</td></tr>
	<tr><td>single_sign_on_managed_application_instance_id</td><td>The SSO managed application instance ID.</td></tr>
	<tr><td>status</td><td>The domain's status.</td></tr>
	<tr><td>tags_src</td><td>The list of tags for the domain.</td></tr>
	<tr><td>url</td><td>The domain's URL.</td></tr>
	<tr><td>default_user_settings</td><td>Settings which are applied to UserProfiles in this domain if settings are not explicitly specified in a given UserProfile.</td></tr>
	<tr><td>domain_settings</td><td>A collection of domain settings.</td></tr>
	<tr><td>subnet_ids</td><td>The VPC subnets that studio uses for communication.</td></tr>
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