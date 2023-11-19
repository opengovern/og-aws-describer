# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>id</td><td>The unique identifier for each query execution.</td></tr>
	<tr><td>workgroup</td><td>The name of the workgroup in which the query ran.</td></tr>
	<tr><td>catalog</td><td>The name of the data catalog used in the query execution.</td></tr>
	<tr><td>database</td><td>The name of the data database used in the query execution.</td></tr>
	<tr><td>query</td><td>The SQL query statements which the query execution ran.</td></tr>
	<tr><td>effective_engine_version</td><td>The engine version on which the query runs.</td></tr>
	<tr><td>selected_engine_version</td><td>The engine version requested by the users.</td></tr>
	<tr><td>execution_parameters</td><td>A list of values for the parameters in a query.</td></tr>
	<tr><td>statement_type</td><td>The type of query statement that was run.</td></tr>
	<tr><td>substatement_type</td><td>The kind of query statement that was run.</td></tr>
	<tr><td>state</td><td>The state of query execution.</td></tr>
	<tr><td>state_change_reason</td><td>Further detail about the status of the query.</td></tr>
	<tr><td>submission_date_time</td><td>The date and time that the query was submitted.</td></tr>
	<tr><td>completion_date_time</td><td>The date and time that the query completed.</td></tr>
	<tr><td>error_message</td><td>Contains a short description of the error that occurred.</td></tr>
	<tr><td>error_type</td><td>An integer value that provides specific information about an Athena query error.</td></tr>
	<tr><td>error_category</td><td>An integer value that specifies the category of a query failure error.</td></tr>
	<tr><td>retryable</td><td>True if the query might succeed if resubmitted.</td></tr>
	<tr><td>data_manifest_location</td><td>The location and file name of a data manifest file.</td></tr>
	<tr><td>data_scanned_in_bytes</td><td>The number of bytes in the data that was queried.</td></tr>
	<tr><td>engine_execution_time_in_millis</td><td>The number of milliseconds that the query took to execute.</td></tr>
	<tr><td>query_planning_time_in_millis</td><td>The number of milliseconds that Athena took to plan the query processing flow.</td></tr>
	<tr><td>query_queue_time_in_millis</td><td>The number of milliseconds that the query was in your query queue waiting for resources.</td></tr>
	<tr><td>service_processing_time_in_millis</td><td>The number of milliseconds that Athena took to finalize and publish the query results after the query engine finished running the query.</td></tr>
	<tr><td>total_execution_time_in_millis</td><td>The number of milliseconds that Athena took to run the query.</td></tr>
	<tr><td>reused_previous_result</td><td>True if a previous query result was reused; false if the result was generated.</td></tr>
	<tr><td>s3_acl_option</td><td>The Amazon S3 canned ACL that Athena should specify when storing query results.</td></tr>
	<tr><td>encryption_option</td><td>Indicates whether Amazon S3 server-side encryption with Amazon S3-managed keys (SSE_S3), server-side encryption with KMS-managed keys (SSE_KMS), or client-side encryption with KMS-managed keys (CSE_KMS) is used.</td></tr>
	<tr><td>kms_key</td><td>For SSE_KMS and CSE_KMS, this is the KMS key ARN or ID.</td></tr>
	<tr><td>expected_bucket_owner</td><td>The Amazon Web Services account ID that you expect to be the owner of the Amazon S3 bucket specified by ResultConfiguration$OutputLocation.</td></tr>
	<tr><td>output_location</td><td>The location in Amazon S3 where your query results are stored.</td></tr>
	<tr><td>result_reuse_by_age_enabled</td><td>True if previous query results can be reused when the query is run.</td></tr>
	<tr><td>result_reuse_by_age_mag_age_in_minutes</td><td>Specifies, in minutes, the maximum age of a previous query result that Athena should consider for reuse. The default is 60.</td></tr>
	<tr><td>partition</td><td>The AWS partition in which the resource is located (aws, aws-cn, or aws-us-gov).</td></tr>
	<tr><td>region</td><td>The AWS Region in which the resource is located.</td></tr>
	<tr><td>account_id</td><td>The AWS Account ID in which the resource is located.</td></tr>
</table>