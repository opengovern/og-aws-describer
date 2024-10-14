# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>name</td><td>The table name.</td></tr>
	<tr><td>catalog_id</td><td>The ID of the Data Catalog in which the table resides.</td></tr>
	<tr><td>create_time</td><td>The time when the table definition was created in the data catalog.</td></tr>
	<tr><td>description</td><td>A description of the table.</td></tr>
	<tr><td>created_by</td><td>The person or entity who created the table.</td></tr>
	<tr><td>database_name</td><td>The name of the database where the table metadata resides.</td></tr>
	<tr><td>is_registered_with_lake_formation</td><td>Indicates whether the table has been registered with lake formation.</td></tr>
	<tr><td>last_access_time</td><td>The last time that the table was accessed. This is usually taken from HDFS, and might not be reliable.</td></tr>
	<tr><td>last_analyzed_time</td><td>The last time that column statistics were computed for this table.</td></tr>
	<tr><td>owner</td><td>The owner of the table.</td></tr>
	<tr><td>retention</td><td>The retention time for this table.</td></tr>
	<tr><td>table_type</td><td>The type of this table (EXTERNAL_TABLE, VIRTUAL_VIEW, etc.).</td></tr>
	<tr><td>update_time</td><td>The last time that the table was updated.</td></tr>
	<tr><td>view_expanded_text</td><td>If the table is a view, the expanded text of the view otherwise null.</td></tr>
	<tr><td>view_original_text</td><td>If the table is a view, the original text of the view otherwise null.</td></tr>
	<tr><td>parameters</td><td>These key-value pairs define properties associated with the table.</td></tr>
	<tr><td>partition_keys</td><td>A list of columns by which the table is partitioned.</td></tr>
	<tr><td>storage_descriptor</td><td>A storage descriptor containing information about the physical storage of this table.</td></tr>
	<tr><td>target_table</td><td>A TableIdentifier structure that describes a target table for resource linking.</td></tr>
	<tr><td>title</td><td>Title of the resource.</td></tr>
	<tr><td>akas</td><td>Array of globally unique identifier strings (also known as) for the resource.</td></tr>
	<tr><td>partition</td><td>The AWS partition in which the resource is located (aws, aws-cn, or aws-us-gov).</td></tr>
	<tr><td>region</td><td>The AWS Region in which the resource is located.</td></tr>
	<tr><td>account_id</td><td>The AWS Account ID in which the resource is located.</td></tr>
	<tr><td>kaytu_account_id</td><td>The Kaytu Account ID in which the resource is located.</td></tr>
	<tr><td>kaytu_resource_id</td><td>The unique ID of the resource in opengovernance.</td></tr>
	<tr><td>kaytu_metadata</td><td>kaytu Metadata of the AWS resource.</td></tr>
</table>