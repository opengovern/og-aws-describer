# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>name</td><td>The name of the container.</td></tr>
	<tr><td>arn</td><td>The Amazon Resource Name (ARN) of the container.</td></tr>
	<tr><td>status</td><td>The status of container creation or deletion. The status is one of the following: &#39;CREATING&#39;, &#39;ACTIVE&#39;, or &#39;DELETING&#39;.</td></tr>
	<tr><td>access_logging_enabled</td><td>The state of access logging on the container. This value is false by default, indicating that AWS Elemental MediaStore does not send access logs to Amazon CloudWatch Logs. When you enable access logging on the container, MediaStore changes this value to true, indicating that the service delivers access logs for objects stored in that container to CloudWatch Logs.</td></tr>
	<tr><td>creation_time</td><td>The Unix timestamp that the container was created.</td></tr>
	<tr><td>endpoint</td><td>The DNS endpoint of the container.</td></tr>
	<tr><td>policy</td><td>The contents of the access policy.</td></tr>
	<tr><td>policy_std</td><td>Contains the contents of the access policy in a canonical form for easier searching.</td></tr>
	<tr><td>tags_src</td><td>A list of tags associated with the container</td></tr>
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