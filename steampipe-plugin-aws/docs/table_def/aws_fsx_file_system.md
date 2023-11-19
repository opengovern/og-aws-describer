<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>file_system_id</td><td>The system-generated, unique 17-digit ID of the file system.</td></tr>
	<tr><td>arn</td><td>The Amazon Resource Name (ARN) for the EFS file system.</td></tr>
	<tr><td>file_system_type</td><td>The type of Amazon FSx file system, which can be LUSTRE, WINDOWS, or ONTAP.</td></tr>
	<tr><td>lifecycle</td><td>The lifecycle status of the file system, following are the possible values AVAILABLE, CREATING, DELETING, FAILED, MISCONFIGURED, UPDATING.</td></tr>
	<tr><td>creation_time</td><td>The time that the file system was created.</td></tr>
	<tr><td>dns_name</td><td>The DNS name for the file system.</td></tr>
	<tr><td>file_system_type_version</td><td>The version of your Amazon FSx for Lustre file system, either 2.10 or 2.12.</td></tr>
	<tr><td>kms_key_id</td><td>The ID of the Key Management Service (KMS) key used to encrypt the file system's.</td></tr>
	<tr><td>owner_id</td><td>The AWS account that created the file system.</td></tr>
	<tr><td>storage_capacity</td><td>The storage capacity of the file system in gibibytes (GiB).</td></tr>
	<tr><td>storage_type</td><td>The storage type of the file system.</td></tr>
	<tr><td>vpc_id</td><td>The ID of the primary VPC for the file system.</td></tr>
	<tr><td>administrative_actions</td><td>A list of administrative actions for the file system that are in process or waiting to be processed.</td></tr>
	<tr><td>failure_details</td><td>A structure providing details of any failures that occur when creating the file system has failed.</td></tr>
	<tr><td>lustre_configuration</td><td>The configuration for the Amazon FSx for Lustre file system.</td></tr>
	<tr><td>network_interface_ids</td><td>The IDs of the elastic network interface from which a specific file system is accessible.</td></tr>
	<tr><td>ontap_configuration</td><td>The configuration for this FSx for NetApp ONTAP file system.</td></tr>
	<tr><td>subnet_ids</td><td>Specifies the IDs of the subnets that the file system is accessible from.</td></tr>
	<tr><td>tags_src</td><td>A list of tags associated with Filesystem.</td></tr>
	<tr><td>windows_configuration</td><td>The configuration for this Microsoft Windows file system.</td></tr>
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