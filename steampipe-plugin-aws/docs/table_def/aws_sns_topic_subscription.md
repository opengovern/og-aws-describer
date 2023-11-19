<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>subscription_arn</td><td>Amazon Resource Name of the subscription.</td></tr>
	<tr><td>topic_arn</td><td>The topic ARN that the subscription is associated with.</td></tr>
	<tr><td>owner</td><td>The AWS account ID of the subscription's owner.</td></tr>
	<tr><td>protocol</td><td>The subscription's protocol.</td></tr>
	<tr><td>endpoint</td><td>The subscription's endpoint (format depends on the protocol).</td></tr>
	<tr><td>confirmation_was_authenticated</td><td>Reflects authentication status of the subscription.</td></tr>
	<tr><td>pending_confirmation</td><td>Reflects the confirmation status of the subscription. True if the subscription hasn't been confirmed.</td></tr>
	<tr><td>raw_message_delivery</td><td>true if raw message delivery is enabled for the subscription.</td></tr>
	<tr><td>delivery_policy</td><td>The JSON of the subscription's delivery policy.</td></tr>
	<tr><td>effective_delivery_policy</td><td>The JSON of the effective delivery policy that takes into account the topic delivery policy and account system defaults.</td></tr>
	<tr><td>redrive_policy</td><td>When specified, sends undeliverable messages to the specified Amazon SQS dead-letter queue. Messages that can't be delivered due to client errors (for example, when the subscribed endpoint is unreachable) or server errors (for example, when the service that powers the subscribed endpoint becomes unavailable) are held in the dead-letter queue for further analysis or reprocessing.</td></tr>
	<tr><td>filter_policy</td><td>The filter policy JSON that is assigned to the subscription.</td></tr>
	<tr><td>title</td><td>Title of the resource.</td></tr>
	<tr><td>akas</td><td>Array of globally unique identifier strings (also known as) for the resource.</td></tr>
	<tr><td>partition</td><td>The AWS partition in which the resource is located (aws, aws-cn, or aws-us-gov).</td></tr>
	<tr><td>region</td><td>The AWS Region in which the resource is located.</td></tr>
	<tr><td>account_id</td><td>The AWS Account ID in which the resource is located.</td></tr>
	<tr><td>kaytu_account_id</td><td>The Kaytu Account ID in which the resource is located.</td></tr>
	<tr><td>kaytu_resource_id</td><td>The unique ID of the resource in Kaytu.</td></tr>
	<tr><td>kaytu_metadata</td><td>kaytu Metadata of the AWS resource.</td></tr>
</table>