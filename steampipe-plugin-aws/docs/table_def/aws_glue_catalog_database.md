<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>name</td><td>The name of the database. For Hive compatibility, this is folded to lowercase when it is stored.</td></tr>
	<tr><td>catalog_id</td><td>The ID of the Data Catalog in which the database resides.</td></tr>
	<tr><td>create_time</td><td>The time at which the metadata database was created in the catalog.</td></tr>
	<tr><td>description</td><td>A description of the database.</td></tr>
	<tr><td>location_uri</td><td>The location of the database (for example, an HDFS path).</td></tr>
	<tr><td>create_table_default_permissions</td><td>Creates a set of default permissions on the table for principals.</td></tr>
	<tr><td>parameters</td><td>These key-value pairs define parameters and properties of the database.</td></tr>
	<tr><td>target_database</td><td>A DatabaseIdentifier structure that describes a target database for resource linking.</td></tr>
	<tr><td>title</td><td>Title of the resource.</td></tr>
	<tr><td>akas</td><td>Array of globally unique identifier strings (also known as) for the resource.</td></tr>
	<tr><td>partition</td><td>The AWS partition in which the resource is located (aws, aws-cn, or aws-us-gov).</td></tr>
	<tr><td>region</td><td>The AWS Region in which the resource is located.</td></tr>
	<tr><td>account_id</td><td>The AWS Account ID in which the resource is located.</td></tr>
	<tr><td>kaytu_account_id</td><td>The Kaytu Account ID in which the resource is located.</td></tr>
	<tr><td>kaytu_resource_id</td><td>The unique ID of the resource in Kaytu.</td></tr>
	<tr><td>kaytu_metadata</td><td>kaytu Metadata of the AWS resource.</td></tr>
</table>