<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>name</td><td>Name of the file system provided by the user.</td></tr>
	<tr><td>file_system_id</td><td>The ID of the file system, assigned by Amazon EFS.</td></tr>
	<tr><td>arn</td><td>The Amazon Resource Name (ARN) for the EFS file system.</td></tr>
	<tr><td>owner_id</td><td>The AWS account that created the file system.</td></tr>
	<tr><td>creation_token</td><td>The opaque string specified in the request.</td></tr>
	<tr><td>creation_time</td><td>The time that the file system was created.</td></tr>
	<tr><td>automatic_backups</td><td>Automatic backups use a default backup plan with the AWS Backup recommended settings for automatic backups.</td></tr>
	<tr><td>life_cycle_state</td><td>The lifecycle phase of the file system.</td></tr>
	<tr><td>number_of_mount_targets</td><td>The current number of mount targets that the file system has.</td></tr>
	<tr><td>performance_mode</td><td>The performance mode of the file system.</td></tr>
	<tr><td>encrypted</td><td>A Boolean value that, if true, indicates that the file system is encrypted.</td></tr>
	<tr><td>kms_key_id</td><td>The ID of an AWS Key Management Service (AWS KMS) customer master key (CMK) that was used to protect the encrypted file system.</td></tr>
	<tr><td>throughput_mode</td><td>The throughput mode for a file system.</td></tr>
	<tr><td>provisioned_throughput_in_mibps</td><td>The throughput, measured in MiB/s, that you want to provision for a file system.</td></tr>
	<tr><td>size_in_bytes</td><td>The latest known metered size (in bytes) of data stored in the file system.</td></tr>
	<tr><td>policy</td><td>The JSON formatted FileSystemPolicy for the EFS file system.</td></tr>
	<tr><td>policy_std</td><td>Contains the policy in a canonical form for easier searching.</td></tr>
	<tr><td>tags_src</td><td>A list of tags associated with Filesystem.</td></tr>
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