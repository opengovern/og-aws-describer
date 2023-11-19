<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>name</td><td>The name of the simulation.</td></tr>
	<tr><td>arn</td><td>The Amazon Resource Name (ARN) of the simulation.</td></tr>
	<tr><td>creation_time</td><td>The time when the simulation was created, expressed as the number of seconds and milliseconds in UTC since the Unix epoch (0:0:0.000, January 1, 1970).</td></tr>
	<tr><td>status</td><td>The current status of the simulation.</td></tr>
	<tr><td>execution_id</td><td>A universally unique identifier (UUID) for this simulation.</td></tr>
	<tr><td>maximum_duration</td><td>The maximum running time of the simulation, specified as a number of months (m or M), hours (h or H), or days (d or D).</td></tr>
	<tr><td>role_arn</td><td>The Amazon Resource Name (ARN) of the Identity and Access Management (IAM) role that the simulation assumes to perform actions.</td></tr>
	<tr><td>schema_error</td><td>An error message that SimSpace Weaver returns only if there is a problem with the simulation schema.</td></tr>
	<tr><td>live_simulation_state</td><td>A collection of additional state information, such as domain and clock configuration.</td></tr>
	<tr><td>logging_configuration</td><td>Settings that control how SimSpace Weaver handles your simulation log data.</td></tr>
	<tr><td>schema_s3_location</td><td>The location of the simulation schema in Amazon Simple Storage Service (Amazon S3).</td></tr>
	<tr><td>target_status</td><td>The desired status of the simulation.</td></tr>
	<tr><td>title</td><td>Title of the resource.</td></tr>
	<tr><td>tags</td><td>A map of tags for the resource.</td></tr>
	<tr><td>akas</td><td>Array of globally unique identifier strings (also known as) for the resource.</td></tr>
	<tr><td>partition</td><td>The AWS partition in which the resource is located (aws, aws-cn, or aws-us-gov).</td></tr>
	<tr><td>region</td><td>The AWS Region in which the resource is located.</td></tr>
	<tr><td>account_id</td><td>The AWS Account ID in which the resource is located.</td></tr>
	<tr><td>kaytu_account_id</td><td>The Kaytu Account ID in which the resource is located.</td></tr>
	<tr><td>kaytu_resource_id</td><td>The unique ID of the resource in Kaytu.</td></tr>
	<tr><td>kaytu_metadata</td><td>kaytu Metadata of the AWS resource.</td></tr>
</table>