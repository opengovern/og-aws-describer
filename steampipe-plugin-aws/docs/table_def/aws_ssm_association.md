<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>association_id</td><td>The ID created by the system when you create an association.</td></tr>
	<tr><td>association_name</td><td>The Name of association.</td></tr>
	<tr><td>arn</td><td>The Amazon Resource Name (ARN) specifying the association.</td></tr>
	<tr><td>document_name</td><td>The name of the Systems Manager document.</td></tr>
	<tr><td>date</td><td>The date when the association was made.</td></tr>
	<tr><td>compliance_severity</td><td>A cron expression that specifies a schedule when the association runs.</td></tr>
	<tr><td>apply_only_at_cron_interval</td><td>By default, when you create a new associations, the system runs it immediately after it is created and then according to the schedule you specified. Specify this option if you don't want an association to run immediately after you create it. This parameter is not supported for rate expressions.</td></tr>
	<tr><td>association_version</td><td>The association version.</td></tr>
	<tr><td>automation_target_parameter_name</td><td>Specify the target for the association. This target is required for associations that use an Automation document and target resources by using rate controls.</td></tr>
	<tr><td>document_version</td><td>The version of the document used in the association.</td></tr>
	<tr><td>instance_id</td><td>The ID of the instance.</td></tr>
	<tr><td>last_execution_date</td><td>The date on which the association was last run.</td></tr>
	<tr><td>last_successful_execution_date</td><td>The last date on which the association was successfully run.</td></tr>
	<tr><td>last_update_association_date</td><td>The date when the association was last updated.</td></tr>
	<tr><td>schedule_expression</td><td>A cron expression that specifies a schedule when the association runs.</td></tr>
	<tr><td>max_concurrency</td><td>The maximum number of targets allowed to run the association at the same time.</td></tr>
	<tr><td>max_errors</td><td>The number of errors that are allowed before the system stops sending requests to run the association on additional targets.</td></tr>
	<tr><td>sync_compliance</td><td>The mode for generating association compliance. You can specify AUTO or MANUAL.</td></tr>
	<tr><td>overview</td><td>Information about the association.</td></tr>
	<tr><td>output_location</td><td>An S3 bucket where you want to store the output details of the request.</td></tr>
	<tr><td>parameters</td><td>A description of the parameters for a document.</td></tr>
	<tr><td>status</td><td>The status of the association. Status can be: Pending, Success, or Failed.</td></tr>
	<tr><td>targets</td><td>A cron expression that specifies a schedule when the association runs.</td></tr>
	<tr><td>target_locations</td><td>The combination of AWS Regions and AWS accounts where you want to run the association.</td></tr>
	<tr><td>title</td><td>Title of the resource.</td></tr>
	<tr><td>akas</td><td>Array of globally unique identifier strings (also known as) for the resource.</td></tr>
	<tr><td>partition</td><td>The AWS partition in which the resource is located (aws, aws-cn, or aws-us-gov).</td></tr>
	<tr><td>region</td><td>The AWS Region in which the resource is located.</td></tr>
	<tr><td>account_id</td><td>The AWS Account ID in which the resource is located.</td></tr>
	<tr><td>kaytu_account_id</td><td>The Kaytu Account ID in which the resource is located.</td></tr>
	<tr><td>kaytu_resource_id</td><td>The unique ID of the resource in Kaytu.</td></tr>
	<tr><td>kaytu_metadata</td><td>kaytu Metadata of the AWS resource.</td></tr>
</table>