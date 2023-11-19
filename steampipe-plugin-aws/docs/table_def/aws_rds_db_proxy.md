<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>db_proxy_name</td><td>The identifier for the proxy.</td></tr>
	<tr><td>db_proxy_arn</td><td>The Amazon Resource Name (ARN) for the proxy</td></tr>
	<tr><td>created_date</td><td>The date and time when the proxy was first created.</td></tr>
	<tr><td>status</td><td>The current status of this proxy.</td></tr>
	<tr><td>debug_logging</td><td>Whether the proxy includes detailed information about SQL statements in its logs.</td></tr>
	<tr><td>endpoint</td><td>The endpoint that you can use to connect to the DB proxy.</td></tr>
	<tr><td>engine_family</td><td>The kinds of databases that the proxy can connect to.</td></tr>
	<tr><td>idle_client_timeout</td><td>The number of seconds a connection to the proxy can have no activity before the proxy drops the client connection.</td></tr>
	<tr><td>require_tls</td><td>Indicates whether Transport Layer Security (TLS) encryption is required for connections to the proxy.</td></tr>
	<tr><td>role_arn</td><td>The Amazon Resource Name (ARN) for the IAM role that the proxy uses to access Amazon Secrets Manager.</td></tr>
	<tr><td>updated_date</td><td>The date and time when the proxy was last updated.</td></tr>
	<tr><td>vpc_id</td><td>Provides the VPC ID of the DB proxy.</td></tr>
	<tr><td>auth</td><td>One or more data structures specifying the authorization mechanism to connect to the associated RDS DB instance or Aurora DB cluster.</td></tr>
	<tr><td>vpc_security_group_ids</td><td>Provides a list of VPC security groups that the proxy belongs to.</td></tr>
	<tr><td>vpc_subnet_ids</td><td>The EC2 subnet IDs for the proxy.</td></tr>
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