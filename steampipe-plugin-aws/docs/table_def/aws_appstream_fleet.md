# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>name</td><td>The name of the fleet.</td></tr>
	<tr><td>arn</td><td>The Amazon Resource Name (ARN) for the fleet.</td></tr>
	<tr><td>instance_type</td><td>The instance type to use when launching fleet instances.</td></tr>
	<tr><td>state</td><td>The current state for the fleet.</td></tr>
	<tr><td>created_time</td><td>The time the fleet was created.</td></tr>
	<tr><td>description</td><td>The description to display.</td></tr>
	<tr><td>display_name</td><td>The fleet name to display.</td></tr>
	<tr><td>disconnect_timeout_in_seconds</td><td>The amount of time that a streaming session remains active after users disconnect. If they try to reconnect to the streaming session after a disconnection or network interruption within this time interval, they are connected to their previous session. Otherwise, they are connected to a new session with a new streaming instance. Specify a value between 60 and 360000.</td></tr>
	<tr><td>directory_name</td><td>The fully qualified name of the directory (for example, corp.example.com).</td></tr>
	<tr><td>organizational_unit_distinguished_name</td><td>The distinguished name of the organizational unit for computer accounts.</td></tr>
	<tr><td>enable_default_internet_access</td><td>Indicates whether default internet access is enabled for the fleet.</td></tr>
	<tr><td>fleet_type</td><td>The fleet type. ALWAYS_ON Provides users with instant-on access to their apps. You are charged for all running instances in your fleet, even if no users are streaming apps. ON_DEMAND Provide users with access to applications after they connect, which takes one to two minutes. You are charged for instance streaming when users are connected and a small hourly fee for instances that are not streaming apps.</td></tr>
	<tr><td>iam_role_arn</td><td>The ARN of the IAM role that is applied to the fleet.</td></tr>
	<tr><td>idle_disconnect_timeout_in_seconds</td><td>The amount of time that users can be idle (inactive) before they are disconnected from their streaming session and the DisconnectTimeoutInSeconds time interval begins.</td></tr>
	<tr><td>image_arn</td><td>The ARN for the public, private, or shared image.</td></tr>
	<tr><td>image_name</td><td>The name of the image used to create the fleet.</td></tr>
	<tr><td>max_concurrent_sessions</td><td>The maximum number of concurrent sessions for the fleet.</td></tr>
	<tr><td>max_user_duration_in_seconds</td><td>The maximum amount of time that a streaming session can remain active, in seconds. If users are still connected to a streaming instance five minutes before this limit is reached, they are prompted to save any open documents before being disconnected.</td></tr>
	<tr><td>platform</td><td>The platform of the fleet.</td></tr>
	<tr><td>stream_view</td><td>The AppStream 2.0 view that is displayed to your users when they stream from the fleet. When APP is specified, only the windows of applications opened by users display. When DESKTOP is specified, the standard desktop that is provided by the operating system displays. The default value is APP.</td></tr>
	<tr><td>compute_capacity_status</td><td>The capacity status for the fleet.</td></tr>
	<tr><td>fleet_errors</td><td>The fleet errors.</td></tr>
	<tr><td>session_script_s3_location</td><td>The S3 location of the session scripts configuration zip file. This only applies to Elastic fleets.</td></tr>
	<tr><td>usb_device_filter_strings</td><td>The USB device filter strings associated with the fleet.</td></tr>
	<tr><td>vpc_config</td><td>The VPC configuration for the fleet.</td></tr>
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