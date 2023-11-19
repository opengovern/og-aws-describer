# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>logical_resource_id</td><td>The logical name of the resource specified in the template.</td></tr>
	<tr><td>stack_name</td><td>The name associated with the stack.</td></tr>
	<tr><td>stack_id</td><td>Unique identifier of the stack.</td></tr>
	<tr><td>last_updated_timestamp</td><td>Time the status was updated.</td></tr>
	<tr><td>resource_status</td><td>Current status of the resource.</td></tr>
	<tr><td>resource_type</td><td>Type of resource.</td></tr>
	<tr><td>description</td><td>User defined description associated with the resource.</td></tr>
	<tr><td>physical_resource_id</td><td>The name or unique identifier that corresponds to a physical instance ID of a resource supported by CloudFormation.</td></tr>
	<tr><td>resource_status_reason</td><td>Success/failure message associated with the resource.</td></tr>
	<tr><td>drift_information</td><td>Information about whether the resource&#39;s actual configuration differs, or has drifted, from its expected configuration, as defined in the stack template and any values specified as template parameters. For more information, see Detecting Unregulated Configuration Changes to Stacks and Resources.</td></tr>
	<tr><td>module_info</td><td>Contains information about the module from which the resource was created, if the resource was created from a module included in the stack template.</td></tr>
	<tr><td>title</td><td>Title of the resource.</td></tr>
	<tr><td>partition</td><td>The AWS partition in which the resource is located (aws, aws-cn, or aws-us-gov).</td></tr>
	<tr><td>region</td><td>The AWS Region in which the resource is located.</td></tr>
	<tr><td>account_id</td><td>The AWS Account ID in which the resource is located.</td></tr>
</table>