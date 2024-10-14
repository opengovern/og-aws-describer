# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>endpoint_name</td><td>The name of the DevEndpoint.</td></tr>
	<tr><td>arn</td><td>The Amazon Resource Name (ARN) of the DevEndpoint.</td></tr>
	<tr><td>status</td><td>The current status of this DevEndpoint.</td></tr>
	<tr><td>availability_zone</td><td>The AWS Availability Zone where this DevEndpoint is located.</td></tr>
	<tr><td>created_timestamp</td><td>The point in time at which this DevEndpoint was created.</td></tr>
	<tr><td>extra_jars_s3_path</td><td>The path to one or more Java .jar files in an S3 bucket that should be loaded in your DevEndpoint.</td></tr>
	<tr><td>extra_python_libs_s3_path</td><td>The paths to one or more Python libraries in an Amazon S3 bucket that should be loaded in your DevEndpoint. Multiple values must be complete paths separated by a comma.</td></tr>
	<tr><td>failure_reason</td><td>The reason for a current failure in this DevEndpoint.</td></tr>
	<tr><td>glue_version</td><td>Glue version determines the versions of Apache Spark and Python that Glue supports.</td></tr>
	<tr><td>last_modified_timestamp</td><td>The point in time at which this DevEndpoint was last modified.</td></tr>
	<tr><td>last_update_status</td><td>The status of the last update.</td></tr>
	<tr><td>number_of_nodes</td><td>The number of Glue Data Processing Units (DPUs) allocated to this DevEndpoint.</td></tr>
	<tr><td>number_of_workers</td><td>The number of workers of a defined workerType that are allocated to the development endpoint.</td></tr>
	<tr><td>private_address</td><td>A private IP address to access the DevEndpoint within a VPC if the DevEndpoint is created within one.</td></tr>
	<tr><td>public_address</td><td>The public IP address used by this DevEndpoint. The PublicAddress field is present only when you create a non-virtual private cloud (VPC) DevEndpoint.</td></tr>
	<tr><td>public_key</td><td>The public key to be used by this DevEndpoint for authentication.</td></tr>
	<tr><td>role_arn</td><td>The Amazon Resource Name (ARN) of the IAM role used in this DevEndpoint.</td></tr>
	<tr><td>security_configuration</td><td>The name of the SecurityConfiguration structure to be used with this DevEndpoint.</td></tr>
	<tr><td>subnet_id</td><td>The subnet ID for this DevEndpoint.</td></tr>
	<tr><td>vpc_id</td><td>The ID of the virtual private cloud (VPC) used by this DevEndpoint.</td></tr>
	<tr><td>worker_type</td><td>The type of predefined worker that is allocated to the development endpoint. Accepts a value of Standard, G.1X, or G.2X.</td></tr>
	<tr><td>yarn_endpoint_address</td><td>The YARN endpoint address used by this DevEndpoint.</td></tr>
	<tr><td>zeppelin_remote_spark_interpreter_port</td><td>The Apache Zeppelin port for the remote Apache Spark interpreter.</td></tr>
	<tr><td>public_keys</td><td>A list of public keys to be used by the DevEndpoints for authentication.</td></tr>
	<tr><td>security_group_ids</td><td>A list of security group identifiers used in this DevEndpoint.</td></tr>
	<tr><td>title</td><td>Title of the resource.</td></tr>
	<tr><td>akas</td><td>Array of globally unique identifier strings (also known as) for the resource.</td></tr>
	<tr><td>partition</td><td>The AWS partition in which the resource is located (aws, aws-cn, or aws-us-gov).</td></tr>
	<tr><td>region</td><td>The AWS Region in which the resource is located.</td></tr>
	<tr><td>account_id</td><td>The AWS Account ID in which the resource is located.</td></tr>
	<tr><td>kaytu_account_id</td><td>The Kaytu Account ID in which the resource is located.</td></tr>
	<tr><td>kaytu_resource_id</td><td>The unique ID of the resource in opengovernance.</td></tr>
	<tr><td>kaytu_metadata</td><td>kaytu Metadata of the AWS resource.</td></tr>
</table>