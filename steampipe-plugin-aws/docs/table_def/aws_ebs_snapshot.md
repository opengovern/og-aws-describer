# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>snapshot_id</td><td>The ID of the snapshot. Each snapshot receives a unique identifier when it is created.</td></tr>
	<tr><td>arn</td><td>The Amazon Resource Name (ARN) specifying the snapshot.</td></tr>
	<tr><td>state</td><td>The snapshot state.</td></tr>
	<tr><td>volume_size</td><td>The size of the volume, in GiB.</td></tr>
	<tr><td>volume_id</td><td>The ID of the volume that was used to create the snapshot. Snapshots created by the CopySnapshot action have an arbitrary volume ID that should not be used for any purpose.</td></tr>
	<tr><td>encrypted</td><td>Indicates whether the snapshot is encrypted.</td></tr>
	<tr><td>start_time</td><td>The time stamp when the snapshot was initiated.</td></tr>
	<tr><td>description</td><td>The description for the snapshot.</td></tr>
	<tr><td>kms_key_id</td><td>The Amazon Resource Name (ARN) of the AWS Key Management Service (AWS KMS) customer master key (CMK) that was used to protect the volume encryption key for the parent volume.</td></tr>
	<tr><td>data_encryption_key_id</td><td>The data encryption key identifier for the snapshot. This value is a unique identifier that corresponds to the data encryption key that was used to encrypt the original volume or snapshot copy. Because data encryption keys are inherited by volumes created from snapshots, and vice versa, if snapshots share the same data encryption key identifier, then they belong to the same volume/snapshot lineage.</td></tr>
	<tr><td>progress</td><td>The progress of the snapshot, as a percentage.</td></tr>
	<tr><td>state_message</td><td>Encrypted Amazon EBS snapshots are copied asynchronously. If a snapshot copy operation fails this field displays error state details to help you diagnose why the error occurred.</td></tr>
	<tr><td>owner_alias</td><td>The AWS owner alias, from an Amazon-maintained list (amazon). This is not the user-configured AWS account alias set using the IAM console.</td></tr>
	<tr><td>owner_id</td><td>The AWS account ID of the EBS snapshot owner.</td></tr>
	<tr><td>create_volume_permissions</td><td>The users and groups that have the permissions for creating volumes from the snapshot.</td></tr>
	<tr><td>tags_src</td><td>A list of tags assigned to the snapshot.</td></tr>
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