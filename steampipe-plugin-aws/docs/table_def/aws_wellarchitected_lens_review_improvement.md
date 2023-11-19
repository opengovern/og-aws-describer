# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>lens_alias</td><td>The alias of the lens. For Amazon Web Services official lenses, this is either the lens alias, such as serverless, or the lens ARN, such as arn:aws:wellarchitected:us-east-1:123456789012:lens/my-lens. Each lens is identified by its LensSummary$LensAlias.</td></tr>
	<tr><td>lens_arn</td><td>The ARN for the lens.</td></tr>
	<tr><td>workload_id</td><td>The ID assigned to the workload.</td></tr>
	<tr><td>milestone_number</td><td>The milestone number. A workload can have a maximum of 100 milestones.</td></tr>
	<tr><td>improvement_plan_url</td><td>The improvement plan URL for a question. This value is only available if the question has been answered.</td></tr>
	<tr><td>pillar_id</td><td>The ID used to identify a pillar, for example, security. A pillar is identified by its PillarReviewSummary$PillarId.</td></tr>
	<tr><td>question_id</td><td>The ID of the question.</td></tr>
	<tr><td>question_title</td><td>The title of the question.</td></tr>
	<tr><td>risk</td><td>The risk for a given workload, lens review, pillar, or question.</td></tr>
	<tr><td>improvement_plans</td><td>The improvement plan details.</td></tr>
	<tr><td>title</td><td>Title of the resource.</td></tr>
	<tr><td>partition</td><td>The AWS partition in which the resource is located (aws, aws-cn, or aws-us-gov).</td></tr>
	<tr><td>region</td><td>The AWS Region in which the resource is located.</td></tr>
	<tr><td>account_id</td><td>The AWS Account ID in which the resource is located.</td></tr>
	<tr><td>kaytu_account_id</td><td>The Kaytu Account ID in which the resource is located.</td></tr>
	<tr><td>kaytu_resource_id</td><td>The unique ID of the resource in Kaytu.</td></tr>
	<tr><td>kaytu_metadata</td><td>kaytu Metadata of the AWS resource.</td></tr>
</table>