# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>principal_arn</td><td>The ARN of the IAM resource (user, group, role, or managed policy) used to generate information about when the resource was last used in an attempt to access an AWS service.</td></tr>
	<tr><td>service_name</td><td>The name of the service in which access was attempted.</td></tr>
	<tr><td>service_namespace</td><td>The namespace of the service in which access was attempted.</td></tr>
	<tr><td>last_authenticated</td><td>The date and time when an authenticated entity most recently attempted to access the service. AWS does not report unauthenticated requests.</td></tr>
	<tr><td>last_authenticated_entity</td><td>The ARN of the authenticated entity (user or role) that last attempted to access the service. AWS does not report unauthenticated requests.</td></tr>
	<tr><td>last_authenticated_region</td><td>The Region from which the authenticated entity (user or role) last attempted to access the service. AWS does not report unauthenticated requests.</td></tr>
	<tr><td>total_authenticated_entities</td><td>The total number of authenticated principals (root user, IAM users, or IAM roles) that have attempted to access the service.</td></tr>
	<tr><td>tracked_actions_last_accessed</td><td>An array of objects that contains details about the most recent attempt to access a tracked action within the service.  Currently, only S3 supports action level tracking.</td></tr>
	<tr><td>partition</td><td>The AWS partition in which the resource is located (aws, aws-cn, or aws-us-gov).</td></tr>
	<tr><td>region</td><td>The AWS Region in which the resource is located.</td></tr>
	<tr><td>account_id</td><td>The AWS Account ID in which the resource is located.</td></tr>
</table>