# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>name</td><td>The API&#39;s name</td></tr>
	<tr><td>api_id</td><td>The API&#39;s identifier. This identifier is unique across all of APIs in API Gateway</td></tr>
	<tr><td>version</td><td>A version identifier for the API</td></tr>
	<tr><td>api_key_source</td><td>The source of the API key for metering requests according to a usage plan</td></tr>
	<tr><td>created_date</td><td>The timestamp when the API was created</td></tr>
	<tr><td>description</td><td>The API&#39;s description</td></tr>
	<tr><td>minimum_compression_size</td><td>A nullable integer that is used to enable compression (with non-negative between 0 and 10485760 (10M) bytes, inclusive) or disable compression (with a null value) on an API. When compression is enabled, compression or decompression is not applied on the payload if the payload size is smaller than this value</td></tr>
	<tr><td>policy</td><td>A stringified JSON policy document that applies to this RestApi regardless of the caller and Method configuration</td></tr>
	<tr><td>policy_std</td><td>Contains the policy in a canonical form for easier searching.</td></tr>
	<tr><td>binary_media_types</td><td>The list of binary media types supported by the RestApi. By default, the RestApi supports only UTF-8-encoded text payloads</td></tr>
	<tr><td>endpoint_configuration_types</td><td>The endpoint configuration of this RestApi showing the endpoint types of the API</td></tr>
	<tr><td>endpoint_configuration_vpc_endpoint_ids</td><td>The endpoint configuration of this RestApi showing the endpoint types of the API</td></tr>
	<tr><td>warnings</td><td>The warning messages reported when failonwarnings is turned on during API import</td></tr>
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