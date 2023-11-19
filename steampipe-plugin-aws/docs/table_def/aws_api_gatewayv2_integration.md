<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>integration_id</td><td>Represents the identifier of an integration.</td></tr>
	<tr><td>api_id</td><td>Represents the identifier of an API.</td></tr>
	<tr><td>arn</td><td>The Amazon Resource Name (ARN) specifying the integration.</td></tr>
	<tr><td>description</td><td>Represents the description of an integration.</td></tr>
	<tr><td>integration_method</td><td>Specifies the integration's HTTP method type.</td></tr>
	<tr><td>integration_type</td><td>Represents an API method integration type.</td></tr>
	<tr><td>integration_uri</td><td>A string representation of a URI with a length between [1-2048]. For a Lambda integration, specify the URI of a Lambda function. For an HTTP integration, specify a fully-qualified URL.</td></tr>
	<tr><td>api_gateway_managed</td><td>Specifies whether an integration is managed by API Gateway. If you created an API using using quick create, the resulting integration is managed by API Gateway. You can update a managed integration, but you can't delete it.</td></tr>
	<tr><td>connection_id</td><td>The ID of the VPC link for a private integration. Supported only for HTTP APIs.</td></tr>
	<tr><td>connection_type</td><td>Represents a connection type.</td></tr>
	<tr><td>content_handling_strategy</td><td>Specifies how to handle response payload content type conversions. Supported only for WebSocket APIs.</td></tr>
	<tr><td>credentials_arn</td><td>Specifies the credentials required for the integration, if any. For AWS integrations, three options are available. To specify an IAM Role for API Gateway to assume, use the role's Amazon Resource Name (ARN).</td></tr>
	<tr><td>integration_response_selection_expression</td><td>An expression used to extract information at runtime. See Selection Expressions(https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-apikey-selection-expressions for more information.</td></tr>
	<tr><td>integration_subtype</td><td>A string with a length between [1-128].</td></tr>
	<tr><td>passthrough_behavior</td><td>Represents passthrough behavior for an integration response. Supported only for WebSocket APIs.</td></tr>
	<tr><td>payload_format_version</td><td>Specifies the format of the payload sent to an integration. Required for HTTP APIs.</td></tr>
	<tr><td>template_selection_expression</td><td>The template selection expression for the integration. Supported only for WebSocket APIs.</td></tr>
	<tr><td>timeout_in_millis</td><td>Indicates custom timeout between 50 and 29,000 milliseconds for WebSocket APIs and between 50 and 30,000 milliseconds for HTTP APIs. The default timeout is 29 seconds for WebSocket APIs and 30 seconds for HTTP APIs.</td></tr>
	<tr><td>request_parameters</td><td>For HTTP API itegrations, without a specified integrationSubtype request parameters are a key-value map specifying how to transform HTTP requests before sending them to backend integrations. The key should follow the pattern <action>:<header|querystring|path>.<location>. The action can be append, overwrite or remove. For values, you can provide static values, or map request data, stage variables, or context variables that are evaluated at runtime. To learn more, see Transforming API requests and responses (https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-parameter-mapping.html).</td></tr>
	<tr><td>request_templates</td><td>Represents a map of Velocity templates that are applied on the request payload based on the value of the Content-Type header sent by the client. The content type value is the key in this map, and the template (as a String) is the value. Supported only for WebSocket APIs.</td></tr>
	<tr><td>response_parameters</td><td>API requests and responses (https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-parameter-mapping.html).</td></tr>
	<tr><td>tls_config</td><td>The TLS configuration for a private integration. If you specify a TLS configuration, private integration traffic uses the HTTPS protocol. Supported only for HTTP APIs.</td></tr>
	<tr><td>title</td><td>Title of the resource.</td></tr>
	<tr><td>akas</td><td>Array of globally unique identifier strings (also known as) for the resource.</td></tr>
	<tr><td>partition</td><td>The AWS partition in which the resource is located (aws, aws-cn, or aws-us-gov).</td></tr>
	<tr><td>region</td><td>The AWS Region in which the resource is located.</td></tr>
	<tr><td>account_id</td><td>The AWS Account ID in which the resource is located.</td></tr>
	<tr><td>kaytu_account_id</td><td>The Kaytu Account ID in which the resource is located.</td></tr>
	<tr><td>kaytu_resource_id</td><td>The unique ID of the resource in Kaytu.</td></tr>
	<tr><td>kaytu_metadata</td><td>kaytu Metadata of the AWS resource.</td></tr>
</table>