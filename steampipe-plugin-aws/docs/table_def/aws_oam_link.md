# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>id</td><td>The random ID string that Amazon Web Service generates as part of the link ARN.</td></tr>
	<tr><td>arn</td><td>The ARN of the link.</td></tr>
	<tr><td>sink_arn</td><td>The ARN of the sink that this link is attached to.</td></tr>
	<tr><td>label</td><td>The label that was assigned to this link at creation, with the variables resolved to their actual values.</td></tr>
	<tr><td>label_template</td><td>The exact label template that was specified when the link was created, with the template variables not resolved.</td></tr>
	<tr><td>resource_types</td><td>The resource types supported by this link.</td></tr>
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