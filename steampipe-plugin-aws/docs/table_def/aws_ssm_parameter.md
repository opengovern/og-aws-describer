# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>name</td><td>The parameter name.</td></tr>
	<tr><td>type</td><td>The type of parameter. Valid parameter types include the following: String, StringList, and SecureString.</td></tr>
	<tr><td>value</td><td>The value of parameter.</td></tr>
	<tr><td>arn</td><td>The Amazon Resource Name (ARN) of the parameter.</td></tr>
	<tr><td>data_type</td><td>The data type of the parameter, such as text or aws:ec2:image. The default is text.</td></tr>
	<tr><td>key_id</td><td>The ID of the query key used for this parameter.</td></tr>
	<tr><td>last_modified_date</td><td>Date the parameter was last changed or updated.</td></tr>
	<tr><td>last_modified_user</td><td>Amazon Resource Name (ARN) of the AWS user who last changed the parameter.</td></tr>
	<tr><td>version</td><td>The parameter version.</td></tr>
	<tr><td>selector</td><td>Either the version number or the label used to retrieve the parameter value.</td></tr>
	<tr><td>source_result</td><td>SourceResult is the raw result or response from the source. Applies to parameters that reference information in other AWS services.</td></tr>
	<tr><td>tier</td><td>The parameter tier.</td></tr>
	<tr><td>policies</td><td>A list of policies associated with a parameter. Parameter policies help you manage a growing set of parameters by enabling you to assign specific criteria to a parameter such as an expiration date or time to live. Parameter policies are especially helpful in forcing you to update or delete passwords and configuration data stored in Parameter Store.</td></tr>
	<tr><td>tags_src</td><td>A list of tags assigned to the parameter.</td></tr>
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