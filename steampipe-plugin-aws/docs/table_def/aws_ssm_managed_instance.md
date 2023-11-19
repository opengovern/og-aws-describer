<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>name</td><td>The name assigned to an on-premises server or virtual machine (VM) when it is activated as a Systems Manager managed instance.</td></tr>
	<tr><td>instance_id</td><td>The ID of the instance.</td></tr>
	<tr><td>arn</td><td>The Amazon Resource Name (ARN) specifying the instance.</td></tr>
	<tr><td>resource_type</td><td>The type of instance. Instances are either EC2 instances or managed instances.</td></tr>
	<tr><td>ping_status</td><td>Connection status of SSM Agent.</td></tr>
	<tr><td>activation_id</td><td>The activation ID created by Systems Manager when the server or VM was registered.</td></tr>
	<tr><td>agent_version</td><td>The version of SSM Agent running on your Linux instance.</td></tr>
	<tr><td>association_status</td><td>The status of the association.</td></tr>
	<tr><td>computer_name</td><td>The fully qualified host name of the managed instance.</td></tr>
	<tr><td>ip_address</td><td>The IP address of the managed instance.</td></tr>
	<tr><td>is_latest_version</td><td>Indicates whether the latest version of SSM Agent is running on your Linux Managed Instance.</td></tr>
	<tr><td>last_association_execution_date</td><td>The date the association was last run.</td></tr>
	<tr><td>last_ping_date_time</td><td>The date and time when the agent last pinged the Systems Manager service.</td></tr>
	<tr><td>last_successful_association_execution_date</td><td>The last date the association was successfully run.</td></tr>
	<tr><td>platform_name</td><td>The name of the operating system platform running on your instance.</td></tr>
	<tr><td>platform_type</td><td>The operating system platform type.</td></tr>
	<tr><td>platform_version</td><td>The version of the OS platform running on your instance.</td></tr>
	<tr><td>registration_date</td><td>The date the server or VM was registered with AWS as a managed instance.</td></tr>
	<tr><td>association_overview</td><td>Information about the association.</td></tr>
	<tr><td>title</td><td>Title of the resource.</td></tr>
	<tr><td>akas</td><td>Array of globally unique identifier strings (also known as) for the resource.</td></tr>
	<tr><td>partition</td><td>The AWS partition in which the resource is located (aws, aws-cn, or aws-us-gov).</td></tr>
	<tr><td>region</td><td>The AWS Region in which the resource is located.</td></tr>
	<tr><td>account_id</td><td>The AWS Account ID in which the resource is located.</td></tr>
	<tr><td>kaytu_account_id</td><td>The Kaytu Account ID in which the resource is located.</td></tr>
	<tr><td>kaytu_resource_id</td><td>The unique ID of the resource in Kaytu.</td></tr>
	<tr><td>kaytu_metadata</td><td>kaytu Metadata of the AWS resource.</td></tr>
</table>