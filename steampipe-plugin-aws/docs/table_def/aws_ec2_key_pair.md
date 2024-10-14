# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>key_name</td><td>The name of the key pair</td></tr>
	<tr><td>key_pair_id</td><td>The ID of the key pair</td></tr>
	<tr><td>key_fingerprint</td><td>If key pair was created using CreateKeyPair, this is the SHA-1 digest of the DER encoded private key. If key pair was created using ImportKeyPair to provide AWS the public key, this is the MD5 public key fingerprint as specified in section 4 of RFC4716</td></tr>
	<tr><td>tags_src</td><td>A list of tags assigned to the key pair</td></tr>
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