# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>name</td><td>The name of the API</td></tr>
	<tr><td>api_id</td><td>The API ID</td></tr>
	<tr><td>api_endpoint</td><td>The URI of the API, of the form {api-id}.execute-api.{region}.amazonaws.com</td></tr>
	<tr><td>protocol_type</td><td>The API protocol</td></tr>
	<tr><td>api_key_selection_expression</td><td>An API key selection expression. Supported only for WebSocket APIs</td></tr>
	<tr><td>disable_execute_api_endpoint</td><td>Specifies whether clients can invoke your API by using the default execute-api endpoint.</td></tr>
	<tr><td>route_selection_expression</td><td>The route selection expression for the API. For HTTP APIs, the routeSelectionExpression must be ${request.method} ${request.path}. If not provided, this will be the default for HTTP APIs</td></tr>
	<tr><td>created_date</td><td>The timestamp when the API was created</td></tr>
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