# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>name</td><td>The name of the pipeline.</td></tr>
	<tr><td>arn</td><td>The Amazon Resource Name (ARN) of the pipeline.</td></tr>
	<tr><td>created_at</td><td>The date and time the pipeline was created.</td></tr>
	<tr><td>role_arn</td><td>The Amazon Resource Name (ARN) for AWS CodePipeline to use to either perform actions with no actionRoleArn, or to use to assume roles for actions with an actionRoleArn.</td></tr>
	<tr><td>updated_at</td><td>The date and time of the last update to the pipeline.</td></tr>
	<tr><td>version</td><td>The version number of the pipeline.</td></tr>
	<tr><td>encryption_key</td><td>The encryption key used to encrypt the data in the artifact store, such as an AWS Key Management Service (AWS KMS) key. If this is undefined, the default key for Amazon S3 is used.</td></tr>
	<tr><td>artifact_stores</td><td>A mapping of artifactStore objects and their corresponding AWS Regions. There must be an artifact store for the pipeline Region and for each cross-region action in the pipeline.</td></tr>
	<tr><td>stages</td><td>The stage in which to perform the action.</td></tr>
	<tr><td>tags_src</td><td>A list of tag key and value pairs associated with this pipeline.</td></tr>
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