# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>vault_name</td><td>The name of the vault.</td></tr>
	<tr><td>vault_arn</td><td>The Amazon Resource Name (ARN) of the vault.</td></tr>
	<tr><td>creation_date</td><td>The Universal Coordinated Time (UTC) date when the vault was created.</td></tr>
	<tr><td>last_inventory_date</td><td>The Universal Coordinated Time (UTC) date when Amazon S3 Glacier completed the last vault inventory.</td></tr>
	<tr><td>number_of_archives</td><td>The number of archives in the vault as of the last inventory date.</td></tr>
	<tr><td>size_in_bytes</td><td>Total size, in bytes, of the archives in the vault as of the last inventory date.</td></tr>
	<tr><td>policy</td><td>Contains the returned vault access policy as a JSON string.</td></tr>
	<tr><td>policy_std</td><td>Contains the policy in a canonical form for easier searching.</td></tr>
	<tr><td>vault_lock_policy</td><td>The vault lock policy.</td></tr>
	<tr><td>vault_lock_policy_std</td><td>Contains the policy in a canonical form for easier searching.</td></tr>
	<tr><td>tags_src</td><td>A list of tags associated with the vault.</td></tr>
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