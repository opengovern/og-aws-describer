# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>finding_account_id</td><td>The Amazon Web Services account ID associated with the finding.</td></tr>
	<tr><td>description</td><td>The description of the finding.</td></tr>
	<tr><td>exploit_available</td><td>If a finding discovered in your environment has an exploit available. Valid values are: YES | NO.</td></tr>
	<tr><td>arn</td><td>The Amazon Resource Number (ARN) of the finding.</td></tr>
	<tr><td>status</td><td>The status of the finding. Valid values are: ACTIVE | SUPPRESSED | CLOSED.</td></tr>
	<tr><td>type</td><td>The type of the finding. Valid values are: NETWORK_REACHABILITY | PACKAGE_VULNERABILITY.</td></tr>
	<tr><td>first_observed_at</td><td>The date and time that the finding was first observed.</td></tr>
	<tr><td>fix_available</td><td>Details on whether a fix is available through a version update. Valid values are: YES | NO | PARTIAL.</td></tr>
	<tr><td>inspector_score</td><td>The Amazon Inspector score given to the finding.</td></tr>
	<tr><td>resource_id</td><td>The ID of the resource.</td></tr>
	<tr><td>resource_type</td><td>The resource type supported by AWS.</td></tr>
	<tr><td>component_type</td><td>The component type.</td></tr>
	<tr><td>component_id</td><td>The component ID of the resource.</td></tr>
	<tr><td>ec2_instance_image_id</td><td>The Amazon EC2 instance image ID.</td></tr>
	<tr><td>ec2_instance_subnet_id</td><td>The Amazon EC2 instance subnet ID.</td></tr>
	<tr><td>ec2_instance_vpc_id</td><td>The Amazon EC2 instance VPC ID.</td></tr>
	<tr><td>ecr_image_architecture</td><td>The Amazon ECR image architecture.</td></tr>
	<tr><td>ecr_image_hash</td><td>The Amazon ECR image hash.</td></tr>
	<tr><td>ecr_image_pushed_at</td><td>The Amazon ECR image push date and time.</td></tr>
	<tr><td>ecr_image_registry</td><td>The Amazon ECR registry.</td></tr>
	<tr><td>ecr_image_repository_name</td><td>The name of the Amazon ECR repository.</td></tr>
	<tr><td>ecr_image_tags</td><td>The tags attached to the Amazon ECR container image.</td></tr>
	<tr><td>lambda_function_execution_role_arn</td><td>The AWS Lambda function execution role ARN.</td></tr>
	<tr><td>lambda_function_last_modified_at</td><td>The AWS Lambda functions the date and time that a user last updated the configuration.</td></tr>
	<tr><td>lambda_function_layers</td><td>The AWS Lambda function layer.</td></tr>
	<tr><td>lambda_function_name</td><td>The AWS Lambda function name.</td></tr>
	<tr><td>lambda_function_runtime</td><td>The AWS Lambda function runtime environment.</td></tr>
	<tr><td>network_protocol</td><td>The ingress source addresse.</td></tr>
	<tr><td>related_vulnerabilitie</td><td>The related vulnerabilitie.</td></tr>
	<tr><td>last_observed_at</td><td>The date and time that the finding was last observed.</td></tr>
	<tr><td>remediation_recommendation_text</td><td>The recommended course of action to remediate the finding.</td></tr>
	<tr><td>remediation_recommendation_url</td><td>The URL address to the CVE remediation recommendations.</td></tr>
	<tr><td>severity</td><td>The severity of the finding. Valid values are: INFORMATIONAL | LOW | MEDIUM | HIGH | CRITICAL | UNTRIAGED.</td></tr>
	<tr><td>updated_at</td><td>The date and time the finding was last updated at.</td></tr>
	<tr><td>source</td><td>The source of the vulnerability information.</td></tr>
	<tr><td>source_url</td><td>A URL to the source of the vulnerability information.</td></tr>
	<tr><td>vendor_created_at</td><td>The date and time that this vulnerability was first added to the vendor’s database.</td></tr>
	<tr><td>vendor_severity</td><td>The severity the vendor has given to this vulnerability type.</td></tr>
	<tr><td>vendor_updated_at</td><td>The date and time the vendor last updated this vulnerability in their database.</td></tr>
	<tr><td>vulnerability_id</td><td>The ID given to this vulnerability.</td></tr>
	<tr><td>exploitability_details</td><td>The details of an exploit available for a finding discovered in your environment.</td></tr>
	<tr><td>inspector_score_details</td><td>An object that contains details of the Amazon Inspector score.</td></tr>
	<tr><td>network_reachability_details</td><td>An object that contains the details of a network reachability finding.</td></tr>
	<tr><td>package_vulnerability_details</td><td>An object that contains the details of a package vulnerability finding.</td></tr>
	<tr><td>cvss</td><td>An object that contains details about the CVSS score of a finding.</td></tr>
	<tr><td>reference_urls</td><td>One or more URLs that contain details about this vulnerability type.</td></tr>
	<tr><td>related_vulnerabilities</td><td>One or more vulnerabilities related to the one identified in this finding.</td></tr>
	<tr><td>vulnerable_package</td><td>The package impacted by this vulnerability.</td></tr>
	<tr><td>vulnerable_packages</td><td>The packages impacted by this vulnerability.</td></tr>
	<tr><td>resources</td><td>Contains information on the resources involved in a finding.</td></tr>
	<tr><td>resource_tags</td><td>Details on the resource tags used to filter findings.</td></tr>
	<tr><td>title</td><td>The title of the finding.</td></tr>
	<tr><td>akas</td><td>Array of globally unique identifier strings (also known as) for the resource.</td></tr>
	<tr><td>partition</td><td>The AWS partition in which the resource is located (aws, aws-cn, or aws-us-gov).</td></tr>
	<tr><td>region</td><td>The AWS Region in which the resource is located.</td></tr>
	<tr><td>account_id</td><td>The AWS Account ID in which the resource is located.</td></tr>
</table>