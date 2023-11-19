# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>name</td><td>The name of the build project.</td></tr>
	<tr><td>arn</td><td>The Amazon Resource Name (ARN) of the build project.</td></tr>
	<tr><td>description</td><td>A description that makes the build project easy to identify.</td></tr>
	<tr><td>concurrent_build_limit</td><td>The maximum number of concurrent builds that are allowed for this project.</td></tr>
	<tr><td>created</td><td>When the build project was created, expressed in Unix time format.</td></tr>
	<tr><td>last_modified</td><td>When the build project&#39;s settings were last modified, expressed in Unix time format.</td></tr>
	<tr><td>encryption_key</td><td>The AWS Key Management Service (AWS KMS) customer master key (CMK) to be.</td></tr>
	<tr><td>queued_timeout_in_minutes</td><td>The number of minutes a build is allowed to be queued before it times out.</td></tr>
	<tr><td>service_role</td><td>The ARN of the AWS Identity and Access Management (IAM) role that enables AWS CodeBuild to interact with dependent AWS services on behalf of the AWS account.</td></tr>
	<tr><td>source_version</td><td>A version of the build input to be built for this project.</td></tr>
	<tr><td>timeout_in_minutes</td><td>How long, in minutes, from 5 to 480 (8 hours), for AWS CodeBuild to wait before timing out any related build that did not get marked as completed.</td></tr>
	<tr><td>artifacts</td><td>Information about the build output artifacts for the build project.</td></tr>
	<tr><td>badge</td><td>Information about the build badge for the build project.</td></tr>
	<tr><td>build_batch_config</td><td>A ProjectBuildBatchConfig object that defines the batch build options for the project.</td></tr>
	<tr><td>cache</td><td>Information about the cache for the build project.</td></tr>
	<tr><td>environment</td><td>Information about the build environment for this build project.</td></tr>
	<tr><td>file_system_locations</td><td>An array of ProjectFileSystemLocation objects for a CodeBuild build project.</td></tr>
	<tr><td>logs_config</td><td>Information about logs for the build project. A project can create logs in Amazon CloudWatch Logs, an S3 bucket or both.</td></tr>
	<tr><td>project_visibility</td><td>Visibility of the build project.</td></tr>
	<tr><td>secondary_artifacts</td><td>An array of ProjectArtifacts objects.</td></tr>
	<tr><td>secondary_source_versions</td><td>An array of ProjectSource objects.</td></tr>
	<tr><td>secondary_sources</td><td>An array of ProjectSource objects.</td></tr>
	<tr><td>source</td><td>Information about the build input source code for this build project.</td></tr>
	<tr><td>vpc_config</td><td>Information about the VPC configuration that AWS CodeBuild accesses.</td></tr>
	<tr><td>webhook</td><td> Information about a webhook that connects repository events to a build project in AWS CodeBuild.</td></tr>
	<tr><td>tags_src</td><td>A list of tag key and value pairs associated with this build project.</td></tr>
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