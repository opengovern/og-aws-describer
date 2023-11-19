# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>id</td><td>The identifier for the authorizer resource</td></tr>
	<tr><td>name</td><td>The name of the authorizer</td></tr>
	<tr><td>rest_api_id</td><td>The id of the rest api</td></tr>
	<tr><td>auth_type</td><td>Optional customer-defined field, used in OpenAPI imports and exports without functional impact</td></tr>
	<tr><td>authorizer_credentials</td><td>Specifies the required credentials as an IAM role for API Gateway to invoke the authorizer</td></tr>
	<tr><td>authorizer_uri</td><td>Specifies the authorizer&#39;s Uniform Resource Identifier (URI). For TOKEN or REQUEST authorizers, this must be a well-formed Lambda function URI</td></tr>
	<tr><td>identity_validation_expression</td><td>A validation expression for the incoming identity token. For TOKEN authorizers, this value is a regular expression</td></tr>
	<tr><td>identity_source</td><td>The identity source for which authorization is requested</td></tr>
	<tr><td>provider_arns</td><td>A list of the Amazon Cognito user pool ARNs for the COGNITO_USER_POOLS authorizer</td></tr>
	<tr><td>title</td><td>Title of the resource.</td></tr>
	<tr><td>akas</td><td>Array of globally unique identifier strings (also known as) for the resource.</td></tr>
	<tr><td>partition</td><td>The AWS partition in which the resource is located (aws, aws-cn, or aws-us-gov).</td></tr>
	<tr><td>region</td><td>The AWS Region in which the resource is located.</td></tr>
	<tr><td>account_id</td><td>The AWS Account ID in which the resource is located.</td></tr>
	<tr><td>kaytu_account_id</td><td>The Kaytu Account ID in which the resource is located.</td></tr>
	<tr><td>kaytu_resource_id</td><td>The unique ID of the resource in Kaytu.</td></tr>
	<tr><td>kaytu_metadata</td><td>kaytu Metadata of the AWS resource.</td></tr>
</table>