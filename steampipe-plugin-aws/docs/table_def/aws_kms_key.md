<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>id</td><td>Unique identifier of the key.</td></tr>
	<tr><td>arn</td><td>ARN of the key.</td></tr>
	<tr><td>key_manager</td><td>The manager of the CMK. CMKs in your AWS account are either customer managed or AWS managed.</td></tr>
	<tr><td>creation_date</td><td>The date and time when the CMK was created.</td></tr>
	<tr><td>aws_account_id</td><td>The twelve-digit account ID of the AWS account that owns the CMK.</td></tr>
	<tr><td>enabled</td><td>Specifies whether the CMK is enabled. When KeyState is Enabled this value is true, otherwise it is false.</td></tr>
	<tr><td>customer_master_key_spec</td><td>Describes the type of key material in the CMK.</td></tr>
	<tr><td>description</td><td>The description of the CMK.</td></tr>
	<tr><td>deletion_date</td><td>The date and time after which AWS KMS deletes the CMK.</td></tr>
	<tr><td>key_state</td><td>The current status of the CMK. For more information about how key state affects the use of a CMK, see [Key state: Effect on your CMK](https://docs.aws.amazon.com/kms/latest/developerguide/key-state.html).</td></tr>
	<tr><td>key_usage</td><td>The [cryptographic operations](https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#cryptographic-operations) for which you can use the CMK.</td></tr>
	<tr><td>origin</td><td>The source of the CMK's key material. When this value is AWS_KMS, AWS KMS created the key material. When this value is EXTERNAL, the key material was imported from your existing key management infrastructure or the CMK lacks key material. When this value is AWS_CLOUDHSM, the key material was created in the AWS CloudHSM cluster associated with a custom key store.</td></tr>
	<tr><td>valid_to</td><td>The time at which the imported key material expires.</td></tr>
	<tr><td>aliases</td><td>A list of aliases for the key.</td></tr>
	<tr><td>key_rotation_enabled</td><td>A Boolean value that specifies whether key rotation is enabled.</td></tr>
	<tr><td>policy</td><td>A key policy document in JSON format.</td></tr>
	<tr><td>policy_std</td><td>Contains the policy in a canonical form for easier searching.</td></tr>
	<tr><td>tags_src</td><td>A list of tags attached to key.</td></tr>
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