# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>arn</td><td>The Amazon Resource Name (ARN) of the HealthEvent.</td></tr>
	<tr><td>availability_zone</td><td>The Amazon Web Services Availability Zone of the event.</td></tr>
	<tr><td>start_time</td><td>The date and time that the event began.</td></tr>
	<tr><td>end_time</td><td>The date and time that the event ended.</td></tr>
	<tr><td>event_scope_code</td><td>This parameter specifies if the Health event is a public Amazon Web Services service event or an account-specific event.</td></tr>
	<tr><td>event_type_category</td><td>A list of event type category codes. Possible values are issue, accountNotification, or scheduledChange.</td></tr>
	<tr><td>event_type_code</td><td>The unique identifier for the event type.</td></tr>
	<tr><td>last_updated_time</td><td>The most recent date and time that the event was updated.</td></tr>
	<tr><td>service</td><td>The Amazon Web Services service that is affected by the event. For example, EC2, RDS.</td></tr>
	<tr><td>status_code</td><td>The most recent status of the event. Possible values are open, closed, and upcoming.</td></tr>
	<tr><td>akas</td><td>Array of globally unique identifier strings (also known as) for the resource.</td></tr>
	<tr><td>partition</td><td>The AWS partition in which the resource is located (aws, aws-cn, or aws-us-gov).</td></tr>
	<tr><td>region</td><td>The AWS Region in which the resource is located.</td></tr>
	<tr><td>account_id</td><td>The AWS Account ID in which the resource is located.</td></tr>
	<tr><td>kaytu_account_id</td><td>The Kaytu Account ID in which the resource is located.</td></tr>
	<tr><td>kaytu_resource_id</td><td>The unique ID of the resource in Kaytu.</td></tr>
	<tr><td>kaytu_metadata</td><td>Kaytu Metadata of the AWS resource.</td></tr>
</table>