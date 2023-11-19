# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>name</td><td>The name of the Maintenance Window.</td></tr>
	<tr><td>window_id</td><td>The ID of the Maintenance Window.</td></tr>
	<tr><td>enabled</td><td>Indicates whether the Maintenance Window is enabled.</td></tr>
	<tr><td>allow_unassociated_targets</td><td>Indicates whether targets must be registered with the Maintenance Window before tasks can be defined for those targets.</td></tr>
	<tr><td>description</td><td>A description of the Maintenance Window.</td></tr>
	<tr><td>tags_src</td><td>A list of tags assigned to the Maintenance Window</td></tr>
	<tr><td>duration</td><td>The duration of the Maintenance Window in hours.</td></tr>
	<tr><td>cutoff</td><td>The number of hours before the end of the Maintenance Window that Systems Manager stops scheduling new tasks for execution.</td></tr>
	<tr><td>schedule</td><td>The schedule of the Maintenance Window in the form of a cron or rate expression.</td></tr>
	<tr><td>schedule_offset</td><td>The number of days to wait to run a Maintenance Window after the scheduled CRON expression date and time.</td></tr>
	<tr><td>targets</td><td>The targets of Maintenance Window.</td></tr>
	<tr><td>tasks</td><td>The Tasks of Maintenance Window.</td></tr>
	<tr><td>created_date</td><td>The date the maintenance window was created.</td></tr>
	<tr><td>end_date</td><td>The date and time, in ISO-8601 Extended format, for when the maintenance window is scheduled to become inactive. The maintenance window will not run after this specified time.</td></tr>
	<tr><td>schedule_timezone</td><td>The schedule of the maintenance window in the form of a cron or rate expression.</td></tr>
	<tr><td>start_date</td><td>The date and time, in ISO-8601 Extended format, for when the maintenance window is scheduled to become active.</td></tr>
	<tr><td>modified_date</td><td>The date the Maintenance Window was last modified.</td></tr>
	<tr><td>next_execution_time</td><td>The next time the maintenance window will actually run, taking into account any specified times for the Maintenance Window to become active or inactive.</td></tr>
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