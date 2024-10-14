# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>workload_name</td><td>The name of the workload.</td></tr>
	<tr><td>workload_arn</td><td>The ARN for the workload.</td></tr>
	<tr><td>workload_id</td><td>The ID assigned to the workload.</td></tr>
	<tr><td>include_shared_resources</td><td>Set to true to have shared resources included in the report.</td></tr>
	<tr><td>base64_string</td><td>The Base64-encoded string representation of a lens review report. This data can be used to create a PDF file. Only returned by GetConsolidatedReport when PDF format is requested.</td></tr>
	<tr><td>lenses_applied_count</td><td>The total number of lenses applied to the workload.</td></tr>
	<tr><td>metric_type</td><td>The metric type of a metric in the consolidated report. Currently only WORKLOAD metric types are supported.</td></tr>
	<tr><td>updated_at</td><td>The date and time when the consolidated report was updated.</td></tr>
	<tr><td>lenses</td><td>The metrics for the lenses in the workload.</td></tr>
	<tr><td>risk_counts</td><td>A map from risk names to the count of how many questions have that rating.</td></tr>
	<tr><td>partition</td><td>The AWS partition in which the resource is located (aws, aws-cn, or aws-us-gov).</td></tr>
	<tr><td>region</td><td>The AWS Region in which the resource is located.</td></tr>
	<tr><td>account_id</td><td>The AWS Account ID in which the resource is located.</td></tr>
	<tr><td>kaytu_account_id</td><td>The Kaytu Account ID in which the resource is located.</td></tr>
	<tr><td>kaytu_resource_id</td><td>The unique ID of the resource in opengovernance.</td></tr>
	<tr><td>kaytu_metadata</td><td>kaytu Metadata of the AWS resource.</td></tr>
</table>