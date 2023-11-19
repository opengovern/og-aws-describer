<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>name</td><td>The name of the alarm.</td></tr>
	<tr><td>arn</td><td>The Amazon Resource Name (ARN) of the alarm.</td></tr>
	<tr><td>state_value</td><td>The state value for the alarm.</td></tr>
	<tr><td>actions_enabled</td><td>Indicates whether actions should be executed during any changes to the alarm state.</td></tr>
	<tr><td>alarm_configuration_updated_timestamp</td><td>The time stamp of the last update to the alarm configuration.</td></tr>
	<tr><td>alarm_description</td><td>The description of the alarm.</td></tr>
	<tr><td>comparison_operator</td><td>The arithmetic operation to use when comparing the specified statistic and threshold. The specified statistic value is used as the first operand.</td></tr>
	<tr><td>datapoints_to_alarm</td><td>The number of data points that must be breaching to trigger the alarm.</td></tr>
	<tr><td>evaluate_low_sample_count_percentile</td><td>Used only for alarms based on percentiles.</td></tr>
	<tr><td>evaluation_periods</td><td>The number of periods over which data is compared to the specified threshold.</td></tr>
	<tr><td>extended_statistic</td><td>The percentile statistic for the metric associated with the alarm. Specify a value between p0.0 and p100.</td></tr>
	<tr><td>metric_name</td><td>The name of the metric associated with the alarm, if this is an alarm based on a single metric.</td></tr>
	<tr><td>namespace</td><td>The namespace of the metric associated with the alarm.</td></tr>
	<tr><td>period</td><td>The period, in seconds, over which the statistic is applied.</td></tr>
	<tr><td>state_reason</td><td>An explanation for the alarm state, in text format.</td></tr>
	<tr><td>state_reason_data</td><td>An explanation for the alarm state, in JSON format.</td></tr>
	<tr><td>state_updated_timestamp</td><td>The time stamp of the last update to the alarm state.</td></tr>
	<tr><td>statistic</td><td>The statistic for the metric associated with the alarm, other than percentile.</td></tr>
	<tr><td>threshold</td><td>The value to compare with the specified statistic.</td></tr>
	<tr><td>threshold_metric_id</td><td>In an alarm based on an anomaly detection model, this is the ID of the ANOMALY_DETECTION_BAND function used as the threshold for the alarm.</td></tr>
	<tr><td>treat_missing_data</td><td>Sets how this alarm is to handle missing data points. If this parameter is omitted, the default behavior of missing is used.</td></tr>
	<tr><td>unit</td><td>The unit of the metric associated with the alarm.</td></tr>
	<tr><td>alarm_actions</td><td>The actions to execute when this alarm transitions to the ALARM state from any other state. Each action is specified as an Amazon Resource Name (ARN).</td></tr>
	<tr><td>dimensions</td><td>The dimensions for the metric associated with the alarm.</td></tr>
	<tr><td>insufficient_data_actions</td><td>The actions to execute when this alarm transitions to the INSUFFICIENT_DATA state from any other state. Each action is specified as an Amazon Resource Name (ARN).</td></tr>
	<tr><td>metrics</td><td>An array of MetricDataQuery structures, used in an alarm based on a metric math expression.</td></tr>
	<tr><td>ok_actions</td><td>The actions to execute when this alarm transitions to the OK state from any other state. Each action is specified as an Amazon Resource Name (ARN).</td></tr>
	<tr><td>tags_src</td><td>The list of tag keys and values associated with alarm.</td></tr>
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