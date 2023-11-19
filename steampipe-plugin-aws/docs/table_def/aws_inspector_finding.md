# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>id</td><td>The ID of the finding.</td></tr>
	<tr><td>arn</td><td>The ARN that specifies the finding.</td></tr>
	<tr><td>agent_id</td><td>The ID of the agent that is installed on the EC2 instance where the finding is generated.</td></tr>
	<tr><td>asset_type</td><td>The type of the host from which the finding is generated.</td></tr>
	<tr><td>auto_scaling_group</td><td>The Auto Scaling group of the EC2 instance where the finding is generated.</td></tr>
	<tr><td>confidence</td><td>This data element is currently not used.</td></tr>
	<tr><td>created_at</td><td>The time when the finding was generated.</td></tr>
	<tr><td>updated_at</td><td>The time when AddAttributesToFindings is called.</td></tr>
	<tr><td>description</td><td>The description of the finding.</td></tr>
	<tr><td>indicator_of_compromise</td><td>This data element is currently not used.</td></tr>
	<tr><td>numeric_severity</td><td>The numeric value of the finding severity.</td></tr>
	<tr><td>recommendation</td><td>The recommendation for the finding.</td></tr>
	<tr><td>schema_version</td><td>The schema version of this data type.</td></tr>
	<tr><td>service</td><td>The data element is set to &#39;Inspector&#39;.</td></tr>
	<tr><td>severity</td><td>The finding severity. Values can be set to High, Medium, Low, and Informational.</td></tr>
	<tr><td>asset_attributes</td><td>A collection of attributes of the host from which the finding is generated.</td></tr>
	<tr><td>attributes</td><td>The system-defined attributes for the finding.</td></tr>
	<tr><td>failed_items</td><td>Attributes details that cannot be described. An error code is provided for each failed item.</td></tr>
	<tr><td>service_attributes</td><td>This data type is used in the Finding data type.</td></tr>
	<tr><td>user_attributes</td><td>The user-defined attributes that are assigned to the finding.</td></tr>
	<tr><td>title</td><td>The name of the finding.</td></tr>
	<tr><td>akas</td><td>Array of globally unique identifier strings (also known as) for the resource.</td></tr>
	<tr><td>partition</td><td>The AWS partition in which the resource is located (aws, aws-cn, or aws-us-gov).</td></tr>
	<tr><td>region</td><td>The AWS Region in which the resource is located.</td></tr>
	<tr><td>account_id</td><td>The AWS Account ID in which the resource is located.</td></tr>
	<tr><td>kaytu_account_id</td><td>The Kaytu Account ID in which the resource is located.</td></tr>
	<tr><td>kaytu_resource_id</td><td>The unique ID of the resource in Kaytu.</td></tr>
	<tr><td>kaytu_metadata</td><td>kaytu Metadata of the AWS resource.</td></tr>
</table>