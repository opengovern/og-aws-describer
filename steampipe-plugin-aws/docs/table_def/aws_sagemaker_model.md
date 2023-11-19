# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>name</td><td>The name of the model.</td></tr>
	<tr><td>arn</td><td>The Amazon Resource Name (ARN) of the model.</td></tr>
	<tr><td>creation_time</td><td>A timestamp that indicates when the model was created.</td></tr>
	<tr><td>enable_network_isolation</td><td>If True, no inbound or outbound network calls can be made to or from the model container.</td></tr>
	<tr><td>execution_role_arn</td><td>The Amazon Resource Name (ARN) of the IAM role that you specified for the model.</td></tr>
	<tr><td>containers</td><td>The containers in the inference pipeline.</td></tr>
	<tr><td>inference_execution_config</td><td>Specifies details of how containers in a multi-container endpoint are called.</td></tr>
	<tr><td>primary_container</td><td>The location of the primary inference code, associated artifacts, and custom environment map that the inference code uses when it is deployed in production.</td></tr>
	<tr><td>vpc_config</td><td>A VpcConfig object that specifies the VPC that this model has access to.</td></tr>
	<tr><td>tags_src</td><td>The list of tags for the model.</td></tr>
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