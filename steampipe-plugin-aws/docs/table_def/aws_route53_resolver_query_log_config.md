# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>id</td><td>The ID for the query logging configuration.</td></tr>
	<tr><td>name</td><td>The name of the query logging configuration.</td></tr>
	<tr><td>arn</td><td>The ARN (Amazon Resource Name) for the query logging configuration.</td></tr>
	<tr><td>creation_time</td><td>The date and time that the query logging configuration was created, in Unix time format and Coordinated Universal Time (UTC).</td></tr>
	<tr><td>status</td><td>The status of the specified query logging configuration. Valid values include CREATING|CREATED|DELETING|FAILED.</td></tr>
	<tr><td>association_count</td><td>The number of VPCs that are associated with the query logging configuration.</td></tr>
	<tr><td>ip_address_count</td><td>The number of IP addresses that you have associated with the Resolver endpoint.</td></tr>
	<tr><td>creator_request_id</td><td>A unique string that identifies the request that created the query logging configuration.</td></tr>
	<tr><td>destination_arn</td><td>The ARN of the resource that you want Resolver to send query logs: an Amazon S3 bucket, a CloudWatch Logs log group, or a Kinesis Data Firehose delivery stream.</td></tr>
	<tr><td>owner_id</td><td>The Amazon Web Services account ID for the account that created the query logging configuration.</td></tr>
	<tr><td>share_status</td><td>An indication of whether the query logging configuration is shared with other Amazon Web Services accounts, or was shared with the current account by another Amazon Web Services account. Sharing is configured through Resource Access Manager (RAM).</td></tr>
	<tr><td>title</td><td>Title of the resource.</td></tr>
	<tr><td>akas</td><td>Array of globally unique identifier strings (also known as) for the resource.</td></tr>
	<tr><td>partition</td><td>The AWS partition in which the resource is located (aws, aws-cn, or aws-us-gov).</td></tr>
	<tr><td>region</td><td>The AWS Region in which the resource is located.</td></tr>
	<tr><td>account_id</td><td>The AWS Account ID in which the resource is located.</td></tr>
</table>