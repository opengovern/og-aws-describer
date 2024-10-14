# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>name</td><td>The name of the notebook instance.</td></tr>
	<tr><td>arn</td><td>The Amazon Resource Name (ARN) of the notebook instance.</td></tr>
	<tr><td>creation_time</td><td>A timestamp that shows when the notebook instance was created.</td></tr>
	<tr><td>default_code_repository</td><td>The Git repository associated with the notebook instance as its default code repository.</td></tr>
	<tr><td>direct_internet_access</td><td>Describes whether Amazon SageMaker provides internet access to the notebook instance.</td></tr>
	<tr><td>failure_reason</td><td>If status is Failed, the reason it failed.</td></tr>
	<tr><td>instance_type</td><td>The type of ML compute instance that the notebook instance is running on.</td></tr>
	<tr><td>kms_key_id</td><td>The AWS KMS key ID Amazon SageMaker uses to encrypt data when storing it on the ML storage volume attached to the instance.</td></tr>
	<tr><td>last_modified_time</td><td>A timestamp that shows when the notebook instance was last modified.</td></tr>
	<tr><td>network_interface_id</td><td>The network interface IDs that Amazon SageMaker created at the time of creating the instance.</td></tr>
	<tr><td>notebook_instance_lifecycle_config_name</td><td>The name of a notebook instance lifecycle configuration associated with this notebook instance.</td></tr>
	<tr><td>notebook_instance_status</td><td>The status of the notebook instance.</td></tr>
	<tr><td>role_arn</td><td>The Amazon Resource Name (ARN) of the IAM role associated with the instance.</td></tr>
	<tr><td>root_access</td><td>Whether root access is enabled or disabled for users of the notebook instance.Lifecycle configurations need root access to be able to set up a notebook instance</td></tr>
	<tr><td>subnet_id</td><td>The ID of the VPC subnet.</td></tr>
	<tr><td>url</td><td>The URL that you use to connect to the Jupyter notebook that is running in your notebook instance.</td></tr>
	<tr><td>volume_size_in_gb</td><td>The size, in GB, of the ML storage volume attached to the notebook instance.</td></tr>
	<tr><td>accelerator_types</td><td>The list of the Elastic Inference (EI) instance types associated with this notebook instance.</td></tr>
	<tr><td>additional_code_repositories</td><td>An array of up to three Git repositories associated with the notebook instance.</td></tr>
	<tr><td>security_groups</td><td>The IDs of the VPC security groups.</td></tr>
	<tr><td>tags_src</td><td>The list of tags for the notebook instance.</td></tr>
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