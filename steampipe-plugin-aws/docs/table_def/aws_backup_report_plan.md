# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>arn</td><td>An Amazon Resource Name (ARN) that uniquely identifies a resource.</td></tr>
	<tr><td>report_plan_name</td><td>The unique name of the report plan.</td></tr>
	<tr><td>description</td><td>An optional description of the report plan with a maximum 1,024 characters.</td></tr>
	<tr><td>creation_time</td><td>The date and time that a report plan is created, in Unix format and Coordinated Universal Time (UTC).</td></tr>
	<tr><td>deployment_status</td><td>The deployment status of a report plan. The statuses are CREATE_IN_PROGRESS, UPDATE_IN_PROGRESS, DELETE_IN_PROGRESS, and COMPLETED.</td></tr>
	<tr><td>last_attempted_execution_time</td><td>The date and time that a report job associated with this report plan last attempted to run, in Unix format and Coordinated Universal Time (UTC).</td></tr>
	<tr><td>last_successful_execution_time</td><td>The date and time that a report job associated with this report plan last successfully ran, in Unix format and Coordinated Universal Time (UTC).</td></tr>
	<tr><td>report_delivery_channel</td><td>Contains information about where and how to deliver your reports, specifically your Amazon S3 bucket name, S3 key prefix, and the formats of your reports.</td></tr>
	<tr><td>report_setting</td><td>Identifies the report template for the report. Reports are built using a report template.</td></tr>
	<tr><td>title</td><td>Title of the resource.</td></tr>
	<tr><td>akas</td><td>Array of globally unique identifier strings (also known as) for the resource.</td></tr>
	<tr><td>partition</td><td>The AWS partition in which the resource is located (aws, aws-cn, or aws-us-gov).</td></tr>
	<tr><td>region</td><td>The AWS Region in which the resource is located.</td></tr>
	<tr><td>account_id</td><td>The AWS Account ID in which the resource is located.</td></tr>
	<tr><td>kaytu_account_id</td><td>The Kaytu Account ID in which the resource is located.</td></tr>
	<tr><td>kaytu_resource_id</td><td>The unique ID of the resource in opengovernance.</td></tr>
	<tr><td>kaytu_metadata</td><td>kaytu Metadata of the AWS resource.</td></tr>
</table>