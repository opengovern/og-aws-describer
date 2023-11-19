# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>name</td><td>The name of the endpoint configuration.</td></tr>
	<tr><td>arn</td><td>The Amazon Resource Name (ARN) of the endpoint configuration.</td></tr>
	<tr><td>creation_time</td><td>A timestamp that shows when the endpoint configuration was created.</td></tr>
	<tr><td>kms_key_id</td><td>AWS KMS key ID Amazon SageMaker uses to encrypt data when storing it on the ML storage volume attached to the instance.</td></tr>
	<tr><td>data_capture_config</td><td>Specifies the parameters to capture input/output of Sagemaker models endpoints.</td></tr>
	<tr><td>production_variants</td><td>An array of ProductionVariant objects, one for each model that you want to host at this endpoint.</td></tr>
	<tr><td>tags_src</td><td>The list of tags for the endpoint configuration.</td></tr>
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