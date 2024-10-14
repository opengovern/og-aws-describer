# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>name</td><td>The auto-generated name for the assessment run.</td></tr>
	<tr><td>arn</td><td>The ARN of the assessment run.</td></tr>
	<tr><td>assessment_template_arn</td><td>The ARN of the assessment template that is associated with the assessment run.</td></tr>
	<tr><td>state</td><td>The state of the assessment run.</td></tr>
	<tr><td>completed_at</td><td>The assessment run completion time that corresponds to the rules packages evaluation completion time or failure.</td></tr>
	<tr><td>created_at</td><td>The time when StartAssessmentRun was called.</td></tr>
	<tr><td>data_collected</td><td>Boolean value (true or false) that specifies whether the process of collecting data from the agents is completed.</td></tr>
	<tr><td>duration_in_seconds</td><td>The duration of the assessment run.</td></tr>
	<tr><td>started_at</td><td>The time when StartAssessmentRun was called.</td></tr>
	<tr><td>state_changed_at</td><td>The last time when the assessment run&#39;s state changed.</td></tr>
	<tr><td>finding_counts</td><td>Provides a total count of generated findings per severity.</td></tr>
	<tr><td>notifications</td><td>A list of notifications for the event subscriptions.</td></tr>
	<tr><td>rules_package_arns</td><td>The rules packages selected for the assessment run.</td></tr>
	<tr><td>state_changes</td><td>A list of the assessment run state changes.</td></tr>
	<tr><td>user_attributes_for_findings</td><td>The user-defined attributes that are assigned to every generated finding.</td></tr>
	<tr><td>title</td><td>Title of the resource.</td></tr>
	<tr><td>akas</td><td>Array of globally unique identifier strings (also known as) for the resource.</td></tr>
	<tr><td>partition</td><td>The AWS partition in which the resource is located (aws, aws-cn, or aws-us-gov).</td></tr>
	<tr><td>region</td><td>The AWS Region in which the resource is located.</td></tr>
	<tr><td>account_id</td><td>The AWS Account ID in which the resource is located.</td></tr>
	<tr><td>kaytu_account_id</td><td>The Kaytu Account ID in which the resource is located.</td></tr>
	<tr><td>kaytu_resource_id</td><td>The unique ID of the resource in opengovernance.</td></tr>
	<tr><td>kaytu_metadata</td><td>kaytu Metadata of the AWS resource.</td></tr>
</table>