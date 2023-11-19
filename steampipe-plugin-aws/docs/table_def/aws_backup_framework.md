# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>framework_name</td><td>The unique name of a backup framework.</td></tr>
	<tr><td>arn</td><td>An Amazon Resource Name (ARN) that uniquely identifies a backup framework resource.</td></tr>
	<tr><td>framework_description</td><td>An optional description of the backup framework.</td></tr>
	<tr><td>deployment_status</td><td>The deployment status of a backup framework.</td></tr>
	<tr><td>creation_time</td><td>The date and time that a framework was created.</td></tr>
	<tr><td>number_of_controls</td><td>The number of controls contained by the framework.</td></tr>
	<tr><td>framework_status</td><td>The framework status based on recording statuses for resources governed by the framework (ACTIVE | PARTIALLY_ACTIVE | INACTIVE | UNAVAILABLE).</td></tr>
	<tr><td>framework_controls</td><td>A list of the controls that make up the framework. Each control in the list has a name, input parameters, and scope.</td></tr>
	<tr><td>tags</td><td>A map of tags for the resource.</td></tr>
	<tr><td>title</td><td>Title of the resource.</td></tr>
	<tr><td>akas</td><td>Array of globally unique identifier strings (also known as) for the resource.</td></tr>
	<tr><td>partition</td><td>The AWS partition in which the resource is located (aws, aws-cn, or aws-us-gov).</td></tr>
	<tr><td>region</td><td>The AWS Region in which the resource is located.</td></tr>
	<tr><td>account_id</td><td>The AWS Account ID in which the resource is located.</td></tr>
	<tr><td>kaytu_account_id</td><td>The Kaytu Account ID in which the resource is located.</td></tr>
	<tr><td>kaytu_resource_id</td><td>The unique ID of the resource in Kaytu.</td></tr>
	<tr><td>kaytu_metadata</td><td>kaytu Metadata of the AWS resource.</td></tr>
</table>