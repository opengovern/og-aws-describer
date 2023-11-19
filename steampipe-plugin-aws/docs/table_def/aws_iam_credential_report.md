<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>user_name</td><td>The friendly name of the user.</td></tr>
	<tr><td>user_arn</td><td>The Amazon Resource Name (ARN) of the user.</td></tr>
	<tr><td>user_creation_time</td><td>The date and time when the user was created.</td></tr>
	<tr><td>generated_time</td><td>The date and time when the credential report was created, in ISO 8601 date-time format (http://www.iso.org/iso/iso8601).</td></tr>
	<tr><td>access_key_1_active</td><td>Does the user have an access key and is the access key's status Active.</td></tr>
	<tr><td>access_key_1_last_rotated</td><td>The date and time when the user's access key was created or last changed.</td></tr>
	<tr><td>access_key_1_last_used_date</td><td>The date and time when the user's access key was most recently used to sign an AWS API request. When an access key is used more than once in a 15-minute span, only the first use is recorded in this field.</td></tr>
	<tr><td>access_key_1_last_used_region</td><td>The AWS Region in which the access key was most recently used. When an access key is used more than once in a 15-minute span, only the first use is recorded in this field.</td></tr>
	<tr><td>access_key_1_last_used_service</td><td>The AWS service that was most recently accessed with the access key. The value in this field uses the service's namespace—for example, s3 for Amazon S3 and ec2 for Amazon EC2. When an access key is used more than once in a 15-minute span, only the first use is recorded in this field.</td></tr>
	<tr><td>access_key_2_active</td><td>Does the user have a second access key and is the access key's status Active.</td></tr>
	<tr><td>access_key_2_last_rotated</td><td>The date and time when the user's second access key was created or last changed.</td></tr>
	<tr><td>access_key_2_last_used_date</td><td>The date and time when the user's second access key was most recently used to sign an AWS API request. When an access key is used more than once in a 15-minute span, only the first use is recorded in this field.</td></tr>
	<tr><td>access_key_2_last_used_region</td><td>The AWS Region in which the user's second access key was most recently used. When an access key is used more than once in a 15-minute span, only the first use is recorded in this field.</td></tr>
	<tr><td>access_key_2_last_used_service</td><td>The AWS service that was most recently accessed with the user's second access key. The value in this field uses the service's namespace—for example, s3 for Amazon S3 and ec2 for Amazon EC2. When an access key is used more than once in a 15-minute span, only the first use is recorded in this field.</td></tr>
	<tr><td>cert_1_active</td><td>Does the user have an X.509 signing certificate and is that certificate's status Active.</td></tr>
	<tr><td>cert_1_last_rotated</td><td>The date and time when the user's signing certificate was created or last changed.</td></tr>
	<tr><td>cert_2_active</td><td>Does the user have a second X.509 signing certificate and is that certificate's status Active.</td></tr>
	<tr><td>cert_2_last_rotated</td><td>The date and time when the user's second signing certificate was created or last changed.</td></tr>
	<tr><td>mfa_active</td><td>Whether a multi-factor authentication (MFA) device has been enabled for the user.</td></tr>
	<tr><td>password_enabled</td><td>When the user has a password, this value is true. Otherwise it is false. The value for the AWS account root user is always false.</td></tr>
	<tr><td>password_last_changed</td><td>The date and time when the user's password was last set.</td></tr>
	<tr><td>password_last_used</td><td>The date and time when the AWS account root user or IAM user's password was last used to sign in to an AWS website.</td></tr>
	<tr><td>password_status</td><td>The status of an user password. Password status can be one of used, never_used and not_set.</td></tr>
	<tr><td>password_next_rotation</td><td>When the account has a password policy that requires password rotation, this field contains the date and time.</td></tr>
	<tr><td>partition</td><td>The AWS partition in which the resource is located (aws, aws-cn, or aws-us-gov).</td></tr>
	<tr><td>region</td><td>The AWS Region in which the resource is located.</td></tr>
	<tr><td>account_id</td><td>The AWS Account ID in which the resource is located.</td></tr>
	<tr><td>kaytu_account_id</td><td>The Kaytu Account ID in which the resource is located.</td></tr>
	<tr><td>kaytu_resource_id</td><td>The unique ID of the resource in Kaytu.</td></tr>
	<tr><td>kaytu_metadata</td><td>kaytu Metadata of the AWS resource.</td></tr>
</table>