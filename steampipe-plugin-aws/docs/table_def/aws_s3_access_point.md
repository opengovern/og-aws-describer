# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>name</td><td>Specifies the name of the access point.</td></tr>
	<tr><td>access_point_arn</td><td>Amazon Resource Name (ARN) of the access point.</td></tr>
	<tr><td>bucket_name</td><td>The name of the bucket associated with this access point.</td></tr>
	<tr><td>access_point_policy_is_public</td><td>Indicates whether this access point policy is public, or not.</td></tr>
	<tr><td>block_public_acls</td><td>Specifies whether Amazon S3 should block public access control lists (ACLs) for buckets in this account.</td></tr>
	<tr><td>block_public_policy</td><td>Specifies whether Amazon S3 should block public bucket policies for buckets in this account.</td></tr>
	<tr><td>ignore_public_acls</td><td>Specifies whether Amazon S3 should ignore public ACLs for buckets in this account.</td></tr>
	<tr><td>restrict_public_buckets</td><td>Specifies whether Amazon S3 should restrict public bucket policies for buckets in this account.</td></tr>
	<tr><td>creation_date</td><td>The date and time when the specified access point was created.</td></tr>
	<tr><td>network_origin</td><td>Indicates whether this access point allows access from the public internet. If VpcConfiguration is specified for this access point, then NetworkOrigin is VPC, and the access point doesn't allow access from the public internet.</td></tr>
	<tr><td>vpc_id</td><td>Specifies the VPC ID from which the access point will only allow connections.</td></tr>
	<tr><td>policy</td><td>The access point policy associated with the specified access point.</td></tr>
	<tr><td>policy_std</td><td>Contains the policy in a canonical form for easier searching.</td></tr>
	<tr><td>title</td><td>Title of the resource.</td></tr>
	<tr><td>akas</td><td>Array of globally unique identifier strings (also known as) for the resource.</td></tr>
	<tr><td>partition</td><td>The AWS partition in which the resource is located (aws, aws-cn, or aws-us-gov).</td></tr>
	<tr><td>region</td><td>The AWS Region in which the resource is located.</td></tr>
	<tr><td>account_id</td><td>The AWS Account ID in which the resource is located.</td></tr>
	<tr><td>kaytu_account_id</td><td>The Kaytu Account ID in which the resource is located.</td></tr>
	<tr><td>kaytu_resource_id</td><td>The unique ID of the resource in Kaytu.</td></tr>
	<tr><td>kaytu_metadata</td><td>kaytu Metadata of the AWS resource.</td></tr>
</table>