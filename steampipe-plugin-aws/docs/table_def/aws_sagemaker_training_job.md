# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>name</td><td>The name of the training job.</td></tr>
	<tr><td>arn</td><td>The Amazon Resource Name (ARN) of the training job.</td></tr>
	<tr><td>training_job_status</td><td>The status of the training job.</td></tr>
	<tr><td>auto_ml_job_arn</td><td>The Amazon Resource Name (ARN) of an AutoML job.</td></tr>
	<tr><td>billable_time_in_seconds</td><td>The billable time in seconds. Billable time refers to the absolute wall-clock time.</td></tr>
	<tr><td>creation_time</td><td>A timestamp that shows when the training job was created.</td></tr>
	<tr><td>enable_managed_spot_training</td><td>A Boolean indicating whether managed spot training is enabled or not.</td></tr>
	<tr><td>enable_network_isolation</td><td>Specifies enable network isolation for training jobs.</td></tr>
	<tr><td>enable_inter_container_traffic_encryption</td><td>To encrypt all communications between ML compute instances in distributed training, choose True.</td></tr>
	<tr><td>failure_reason</td><td>If the training job failed, the reason it failed.</td></tr>
	<tr><td>labeling_job_arn</td><td>The Amazon Resource Name (ARN) of the Amazon SageMaker Ground Truth labeling job that created the transform or training job.</td></tr>
	<tr><td>last_modified_time</td><td>Timestamp when the training job was last modified.</td></tr>
	<tr><td>profiling_status</td><td>Profiling status of a training job.</td></tr>
	<tr><td>role_arn</td><td>The AWS Identity and Access Management (IAM) role configured for the training job.</td></tr>
	<tr><td>secondary_status</td><td>Provides detailed information about the state of the training job.</td></tr>
	<tr><td>training_end_time</td><td>A timestamp that shows when the training job ended.</td></tr>
	<tr><td>training_start_time</td><td>Indicates the time when the training job starts on training instances.</td></tr>
	<tr><td>training_time_in_seconds</td><td>The training time in seconds.</td></tr>
	<tr><td>tuning_job_arn</td><td>The Amazon Resource Name (ARN) of the associated hyperparameter tuning job if the training job was launched by a hyperparameter tuning job.</td></tr>
	<tr><td>algorithm_specification</td><td>Information about the algorithm used for training, and algorithm metadata.</td></tr>
	<tr><td>checkpoint_config</td><td>Contains information about the output location for managed spot training checkpoint data.</td></tr>
	<tr><td>debug_hook_config</td><td>Configuration information for the Debugger hook parameters, metric and tensor collections, and storage paths.</td></tr>
	<tr><td>debug_rule_configurations</td><td>Configuration information for Debugger rules for debugging output tensors.</td></tr>
	<tr><td>debug_rule_evaluation_statuses</td><td>Evaluation status of Debugger rules for debugging on a training job.</td></tr>
	<tr><td>environment</td><td>The environment variables to set in the Docker container.</td></tr>
	<tr><td>experiment_config</td><td>Associates a SageMaker job as a trial component with an experiment and trial.</td></tr>
	<tr><td>final_metric_data_list</td><td>A collection of MetricData objects that specify the names, values, and dates and times that the training algorithm emitted to Amazon CloudWatch.</td></tr>
	<tr><td>hyper_parameters</td><td>Algorithm-specific parameters.</td></tr>
	<tr><td>input_data_config</td><td>An array of Channel objects that describes each data input channel.</td></tr>
	<tr><td>model_artifacts</td><td>Information about the Amazon S3 location that is configured for storing model artifacts.</td></tr>
	<tr><td>output_data_config</td><td>The S3 path where model artifacts that you configured when creating the job are stored.</td></tr>
	<tr><td>profiler_config</td><td>Configuration information for Debugger system monitoring,framework profiling and storage paths.</td></tr>
	<tr><td>profiler_rule_configurations</td><td>Configuration information for Debugger rules for profiling system and framework metrics.</td></tr>
	<tr><td>profiler_rule_evaluation_statuses</td><td>Evaluation status of Debugger rules for profiling on a training job.</td></tr>
	<tr><td>resource_config</td><td>Resources, including ML compute instances and ML storage volumes, that are configured for model training.</td></tr>
	<tr><td>secondary_status_transitions</td><td>A history of all of the secondary statuses that the training job has transitioned through.</td></tr>
	<tr><td>stopping_condition</td><td>Specifies a limit to how long a model training job can run.</td></tr>
	<tr><td>tensor_board_output_config</td><td>Configuration of storage locations for the Debugger TensorBoard output data.</td></tr>
	<tr><td>vpc_config</td><td>A VpcConfig object that specifies the VPC that this training job has access to.</td></tr>
	<tr><td>tags_src</td><td>A list of tags assigned to the training job.</td></tr>
	<tr><td>title</td><td>Title of the resource.</td></tr>
	<tr><td>tags</td><td>A map of tags for the resource.</td></tr>
	<tr><td>akas</td><td>Array of globally unique identifier strings (also known as) for the resource.</td></tr>
	<tr><td>partition</td><td>The AWS partition in which the resource is located (aws, aws-cn, or aws-us-gov).</td></tr>
	<tr><td>region</td><td>The AWS Region in which the resource is located.</td></tr>
	<tr><td>account_id</td><td>The AWS Account ID in which the resource is located.</td></tr>
	<tr><td>kaytu_account_id</td><td>The Kaytu Account ID in which the resource is located.</td></tr>
	<tr><td>kaytu_resource_id</td><td>The unique ID of the resource in opengovernance.</td></tr>
	<tr><td>kaytu_metadata</td><td>kaytu Metadata of the AWS resource.</td></tr>
</table>