# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>name</td><td>The user friendly name of the bucket.</td></tr>
	<tr><td>arn</td><td>The ARN of the AWS S3 Bucket.</td></tr>
	<tr><td>creation_date</td><td>The date and tiem when bucket was created.</td></tr>
	<tr><td>bucket_policy_is_public</td><td>The policy status for an Amazon S3 bucket, indicating whether the bucket is public.</td></tr>
	<tr><td>versioning_enabled</td><td>The versioning state of a bucket.</td></tr>
	<tr><td>versioning_mfa_delete</td><td>The MFA Delete status of the versioning state.</td></tr>
	<tr><td>block_public_acls</td><td>Specifies whether Amazon S3 should block public access control lists (ACLs) for this bucket and objects in this bucket.</td></tr>
	<tr><td>block_public_policy</td><td>Specifies whether Amazon S3 should block public bucket policies for this bucket. If TRUE it causes Amazon S3 to reject calls to PUT Bucket policy if the specified bucket policy allows public access.</td></tr>
	<tr><td>ignore_public_acls</td><td>Specifies whether Amazon S3 should ignore public ACLs for this bucket and objects in this bucket. Setting this element to TRUE causes Amazon S3 to ignore all public ACLs on this bucket and objects in this bucket.</td></tr>
	<tr><td>restrict_public_buckets</td><td>Specifies whether Amazon S3 should restrict public bucket policies for this bucket. Setting this element to TRUE restricts access to this bucket to only AWS service principals and authorized users within this account if the bucket has a public policy.</td></tr>
	<tr><td>event_notification_configuration</td><td>A container for specifying the notification configuration of the bucket. If this element is empty, notifications are turned off for the bucket.</td></tr>
	<tr><td>server_side_encryption_configuration</td><td>The default encryption configuration for an Amazon S3 bucket.</td></tr>
	<tr><td>acl</td><td>The access control list (ACL) of a bucket.</td></tr>
	<tr><td>lifecycle_rules</td><td>The lifecycle configuration information of the bucket.</td></tr>
	<tr><td>logging</td><td>The logging status of a bucket and the permissions users have to view and modify that status.</td></tr>
	<tr><td>object_lock_configuration</td><td>The specified bucket's object lock configuration.</td></tr>
	<tr><td>object_ownership_controls</td><td>The Ownership Controls for an Amazon S3 bucket.</td></tr>
	<tr><td>policy</td><td>The resource IAM access document for the bucket.</td></tr>
	<tr><td>policy_std</td><td>Contains the policy in a canonical form for easier searching.</td></tr>
	<tr><td>replication</td><td>The replication configuration of a bucket.</td></tr>
	<tr><td>website_configuration</td><td>The website configuration information of the bucket.</td></tr>
	<tr><td>tags_src</td><td>A list of tags assigned to bucket.</td></tr>
	<tr><td>tags</td><td>A map of tags for the resource.</td></tr>
	<tr><td>title</td><td>Title of the resource.</td></tr>
	<tr><td>akas</td><td>Array of globally unique identifier strings (also known as) for the resource.</td></tr>
	<tr><td>region</td><td>The AWS Region in which the resource is located.</td></tr>
	<tr><td>partition</td><td>The AWS partition in which the resource is located (aws, aws-cn, or aws-us-gov).</td></tr>
	<tr><td>account_id</td><td>The AWS Account ID in which the resource is located.</td></tr>
	<tr><td>kaytu_account_id</td><td>The Kaytu Account ID in which the resource is located.</td></tr>
	<tr><td>kaytu_resource_id</td><td>The unique ID of the resource in Kaytu.</td></tr>
	<tr><td>kaytu_metadata</td><td>kaytu Metadata of the AWS resource.</td></tr>
</table>