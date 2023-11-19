<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>app_id</td><td>The unique ID of the Amplify app.</td></tr>
	<tr><td>arn</td><td>The Amazon Resource Name (ARN) of the Amplify app.</td></tr>
	<tr><td>name</td><td>The name for the Amplify app.</td></tr>
	<tr><td>description</td><td>The description for the Amplify app.</td></tr>
	<tr><td>create_time</td><td>Creates a date and time for the Amplify app.</td></tr>
	<tr><td>update_time</td><td>Updates the date and time for the Amplify app.</td></tr>
	<tr><td>basic_auth_credentials</td><td>The basic authorization credentials for branches for the Amplify app. You must base64-encode the authorization credentials and provide them in the format user:password.</td></tr>
	<tr><td>custom_headers</td><td>Describes the custom HTTP headers for the Amplify app.</td></tr>
	<tr><td>default_domain</td><td>The default domain for the Amplify app.</td></tr>
	<tr><td>enable_auto_branch_creation</td><td>Enables automated branch creation for the Amplify app.</td></tr>
	<tr><td>enable_basic_auth</td><td>Enables basic authorization for the Amplify app's branches.</td></tr>
	<tr><td>enable_branch_auto_build</td><td>Enables the auto-building of branches for the Amplify app.</td></tr>
	<tr><td>enable_branch_auto_deletion</td><td>Automatically disconnect a branch in the Amplify Console when you delete a branch from your Git repository.</td></tr>
	<tr><td>iam_service_role_arn</td><td>The AWS Identity and Access Management (IAM) service role for the Amazon Resource Name (ARN) of the Amplify app.</td></tr>
	<tr><td>platform</td><td>The platform for the Amplify app.</td></tr>
	<tr><td>repository</td><td>The Git repository for the Amplify app.</td></tr>
	<tr><td>repository_clone_method</td><td>The Amplify service uses this parameter to specify the authentication protocol to use to access the Git repository for an Amplify app. Amplify specifies TOKEN for a GitHub repository, SIGV4 for an AWS CodeCommit repository, and SSH for GitLab and Bitbucket repositories.</td></tr>
	<tr><td>auto_branch_creation_config</td><td>Describes the automated branch creation configuration for the Amplify app.</td></tr>
	<tr><td>auto_branch_creation_patterns</td><td>Describes the automated branch creation glob patterns for the Amplify app.</td></tr>
	<tr><td>build_spec</td><td>Describes the content of the build specification (build spec) for the Amplify app.</td></tr>
	<tr><td>custom_rules</td><td>Describes the custom redirect and rewrite rules for the Amplify app.</td></tr>
	<tr><td>environment_variables</td><td>The environment variables for the Amplify app.</td></tr>
	<tr><td>production_branch</td><td>Describes the information about a production branch of the Amplify app.</td></tr>
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