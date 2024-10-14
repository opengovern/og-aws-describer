# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>name</td><td>The name of the GlueJob.</td></tr>
	<tr><td>arn</td><td>The Amazon Resource Name (ARN) of the GlueJob.</td></tr>
	<tr><td>allocated_capacity</td><td>[DEPRECATED] This column has been deprecated and will be removed in a future release, use max_capacity instead. The number of Glue data processing units (DPUs) that can be allocated when this job runs.</td></tr>
	<tr><td>created_on</td><td>The time and date that this job definition was created.</td></tr>
	<tr><td>description</td><td>A description of the job.</td></tr>
	<tr><td>glue_version</td><td>Glue version determines the versions of Apache Spark and Python that Glue supports.</td></tr>
	<tr><td>last_modified_on</td><td>The last point in time when this job definition was modified.</td></tr>
	<tr><td>log_uri</td><td>This field is reserved for future use.</td></tr>
	<tr><td>max_capacity</td><td>The number of Glue data processing units (DPUs) that can be allocated when this job runs.</td></tr>
	<tr><td>max_retries</td><td>The maximum number of times to retry this job after a JobRun fails.</td></tr>
	<tr><td>number_of_workers</td><td>The number of workers of a defined workerType that are allocated when a job runs.</td></tr>
	<tr><td>role</td><td>The name or Amazon Resource Name (ARN) of the IAM role associated with this job.</td></tr>
	<tr><td>security_configuration</td><td>The name of the SecurityConfiguration structure to be used with this job.</td></tr>
	<tr><td>timeout</td><td>The job timeout in minutes.</td></tr>
	<tr><td>worker_type</td><td>The type of predefined worker that is allocated when a job runs. Accepts a value of Standard, G.1X, or G.2X.</td></tr>
	<tr><td>command</td><td>The JobCommand that runs this job.</td></tr>
	<tr><td>connections</td><td>The connections used for this job.</td></tr>
	<tr><td>default_arguments</td><td>The default arguments for this job, specified as name-value pairs.</td></tr>
	<tr><td>execution_property</td><td>An ExecutionProperty specifying the maximum number of concurrent runs allowed for this job.</td></tr>
	<tr><td>job_bookmark</td><td>Defines a point that a job can resume processing.</td></tr>
	<tr><td>non_overridable_arguments</td><td>Non-overridable arguments for this job, specified as name-value pairs.</td></tr>
	<tr><td>notification_property</td><td>Specifies configuration properties of a job notification.</td></tr>
	<tr><td>title</td><td>Title of the resource.</td></tr>
	<tr><td>akas</td><td>Array of globally unique identifier strings (also known as) for the resource.</td></tr>
	<tr><td>partition</td><td>The AWS partition in which the resource is located (aws, aws-cn, or aws-us-gov).</td></tr>
	<tr><td>region</td><td>The AWS Region in which the resource is located.</td></tr>
	<tr><td>account_id</td><td>The AWS Account ID in which the resource is located.</td></tr>
	<tr><td>kaytu_account_id</td><td>The Kaytu Account ID in which the resource is located.</td></tr>
	<tr><td>kaytu_resource_id</td><td>The unique ID of the resource in opengovernance.</td></tr>
	<tr><td>kaytu_metadata</td><td>kaytu Metadata of the AWS resource.</td></tr>
</table>