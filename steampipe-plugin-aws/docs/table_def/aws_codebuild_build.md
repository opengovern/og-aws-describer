# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>arn</td><td>The ARN of the build.</td></tr>
	<tr><td>id</td><td>The unique identifier of the  build.</td></tr>
	<tr><td>build_batch_arn</td><td>The ARN of the batch build that this build is a member of, if applicable.</td></tr>
	<tr><td>build_complete</td><td>Indicates if the build is complete.</td></tr>
	<tr><td>build_number</td><td>The number of the build.</td></tr>
	<tr><td>build_status</td><td>The status of the build.</td></tr>
	<tr><td>current_phase</td><td>The current build phase.</td></tr>
	<tr><td>encryption_key</td><td>The Key Management Service customer master key (CMK) to be used for encrypting the build output artifacts.</td></tr>
	<tr><td>end_time</td><td>The date and time that the build process ended, expressed in Unix time format.</td></tr>
	<tr><td>initiator</td><td>The entity that started the build.</td></tr>
	<tr><td>project_name</td><td>The name of the build project.</td></tr>
	<tr><td>queued_timeout_in_minutes</td><td>Specifies the amount of time, in minutes, that a build is allowed to be queued before it times out.</td></tr>
	<tr><td>source_version</td><td>The identifier of the version of the source code to be built.</td></tr>
	<tr><td>start_time</td><td>The date and time that the build started.</td></tr>
	<tr><td>timeout_in_minutes</td><td>How long, in minutes, for CodeBuild to wait before timing out this build if it does not get marked as completed.</td></tr>
	<tr><td>resolved_source_version</td><td>The identifier of the resolved version of this build&#39;s source code.</td></tr>
	<tr><td>artifacts</td><td>A BuildArtifacts object the defines the build artifacts for this build.</td></tr>
	<tr><td>cache</td><td>Information about the cache for the build.</td></tr>
	<tr><td>debug_session</td><td>Contains information about the debug session for this build.</td></tr>
	<tr><td>environment</td><td>Information about the build environment for this build project.</td></tr>
	<tr><td>exported_environment_variables</td><td>A list of exported environment variables for this build.</td></tr>
	<tr><td>file_system_locations</td><td>An array of ProjectFileSystemLocation objects for a CodeBuild build project.</td></tr>
	<tr><td>logs</td><td>Information about the build&#39;s logs in CloudWatch Logs.</td></tr>
	<tr><td>network_interfaces</td><td>Describes a network interface.</td></tr>
	<tr><td>phases</td><td>Information about all previous build phases that are complete and information about any current build phase that is not yet complete.</td></tr>
	<tr><td>report_arns</td><td>An array of the ARNs associated with this build&#39;s reports.</td></tr>
	<tr><td>secondary_artifacts</td><td>An array of BuildArtifacts objects the define the build artifacts for this build.</td></tr>
	<tr><td>secondary_source_versions</td><td>An array of ProjectSourceVersion objects.</td></tr>
	<tr><td>secondary_sources</td><td>An array of ProjectSource objects that define the sources for the build.</td></tr>
	<tr><td>source</td><td>Information about the build input source code for the build project.</td></tr>
	<tr><td>vpc_config</td><td>Information about the VPC configuration that CodeBuild accesses.</td></tr>
	<tr><td>title</td><td>Title of the resource.</td></tr>
	<tr><td>akas</td><td>Array of globally unique identifier strings (also known as) for the resource.</td></tr>
	<tr><td>partition</td><td>The AWS partition in which the resource is located (aws, aws-cn, or aws-us-gov).</td></tr>
	<tr><td>region</td><td>The AWS Region in which the resource is located.</td></tr>
	<tr><td>account_id</td><td>The AWS Account ID in which the resource is located.</td></tr>
</table>