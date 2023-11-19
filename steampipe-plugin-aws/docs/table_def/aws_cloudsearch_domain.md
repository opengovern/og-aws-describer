# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>domain_name</td><td>A string that represents the name of a domain.</td></tr>
	<tr><td>domain_id</td><td>An internally generated unique identifier for a domain.</td></tr>
	<tr><td>arn</td><td>The Amazon Resource Name (ARN) of the search domain.</td></tr>
	<tr><td>created</td><td>True if the search domain is created.</td></tr>
	<tr><td>deleted</td><td>True if the search domain has been deleted.</td></tr>
	<tr><td>processing</td><td>True if processing is being done to activate the current domain configuration.</td></tr>
	<tr><td>requires_index_documents</td><td>True if Index Documents need to be called to activate the current domain configuration.</td></tr>
	<tr><td>search_instance_count</td><td>The number of search instances that are available to process search requests.</td></tr>
	<tr><td>search_instance_type</td><td>The instance type that is being used to process search requests.</td></tr>
	<tr><td>search_partition_count</td><td>The number of partitions across which the search index is spread.</td></tr>
	<tr><td>doc_service</td><td>The service endpoint for updating documents in a search domain.</td></tr>
	<tr><td>limits</td><td>Limit details for a search domain.</td></tr>
	<tr><td>search_service</td><td>The service endpoint for requesting search results from a search domain.</td></tr>
	<tr><td>akas</td><td>Array of globally unique identifier strings (also known as) for the resource.</td></tr>
	<tr><td>title</td><td>Title of the resource.</td></tr>
	<tr><td>partition</td><td>The AWS partition in which the resource is located (aws, aws-cn, or aws-us-gov).</td></tr>
	<tr><td>region</td><td>The AWS Region in which the resource is located.</td></tr>
	<tr><td>account_id</td><td>The AWS Account ID in which the resource is located.</td></tr>
	<tr><td>kaytu_account_id</td><td>The Kaytu Account ID in which the resource is located.</td></tr>
	<tr><td>kaytu_resource_id</td><td>The unique ID of the resource in Kaytu.</td></tr>
	<tr><td>kaytu_metadata</td><td>kaytu Metadata of the AWS resource.</td></tr>
</table>