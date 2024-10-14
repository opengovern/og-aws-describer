# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>id</td><td>ID of the inventory result entity.</td></tr>
	<tr><td>type_name</td><td>The type of inventory item returned by the request.</td></tr>
	<tr><td>capture_time</td><td>The time that inventory information was collected for the managed node(s).</td></tr>
	<tr><td>schema_version</td><td>The inventory schema version used by the managed node(s).</td></tr>
	<tr><td>content</td><td>Contains all the inventory data of the item type. Results include attribute names and values.</td></tr>
	<tr><td>schema</td><td>The inventory item schema definition. Users can use this to compose inventory query filters.</td></tr>
	<tr><td>title</td><td>Title of the resource.</td></tr>
	<tr><td>partition</td><td>The AWS partition in which the resource is located (aws, aws-cn, or aws-us-gov).</td></tr>
	<tr><td>region</td><td>The AWS Region in which the resource is located.</td></tr>
	<tr><td>account_id</td><td>The AWS Account ID in which the resource is located.</td></tr>
	<tr><td>kaytu_account_id</td><td>The Kaytu Account ID in which the resource is located.</td></tr>
	<tr><td>kaytu_resource_id</td><td>The unique ID of the resource in opengovernance.</td></tr>
	<tr><td>kaytu_metadata</td><td>kaytu Metadata of the AWS resource.</td></tr>
</table>