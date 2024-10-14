# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>name</td><td>The name of the access point. This is the value of the Name tag.</td></tr>
	<tr><td>access_point_id</td><td>The ID of the access point, assigned by Amazon EFS.</td></tr>
	<tr><td>access_point_arn</td><td>The unique Amazon Resource Name (ARN) associated with the access point.</td></tr>
	<tr><td>life_cycle_state</td><td>Identifies the lifecycle phase of the access point.</td></tr>
	<tr><td>file_system_id</td><td>The ID of the EFS file system that the access point applies to.</td></tr>
	<tr><td>client_token</td><td>The opaque string specified in the request to ensure idempotent creation.</td></tr>
	<tr><td>owner_id</td><td>Identified the AWS account that owns the access point resource.</td></tr>
	<tr><td>posix_user</td><td>The full POSIX identity, including the user ID, group ID, and secondary group IDs on the access point that is used for all file operations by NFS clients using the access point.</td></tr>
	<tr><td>root_directory</td><td>The directory on the Amazon EFS file system that the access point exposes as the root directory to NFS clients using the access point.</td></tr>
	<tr><td>tags_src</td><td>The tags associated with the access point, presented as an array of Tag objects.</td></tr>
	<tr><td>tags</td><td>A map of tags for the resource.</td></tr>
	<tr><td>title</td><td>Title of the resource.</td></tr>
	<tr><td>akas</td><td>Array of globally unique identifier strings (also known as) for the resource.</td></tr>
	<tr><td>partition</td><td>The AWS partition in which the resource is located (aws, aws-cn, or aws-us-gov).</td></tr>
	<tr><td>region</td><td>The AWS Region in which the resource is located.</td></tr>
	<tr><td>account_id</td><td>The AWS Account ID in which the resource is located.</td></tr>
	<tr><td>kaytu_account_id</td><td>The Kaytu Account ID in which the resource is located.</td></tr>
	<tr><td>kaytu_resource_id</td><td>The unique ID of the resource in opengovernance.</td></tr>
	<tr><td>kaytu_metadata</td><td>kaytu Metadata of the AWS resource.</td></tr>
</table>