# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>name</td><td>The name of the domain. For public hosted zones, this is the name that is registered with your DNS registrar.</td></tr>
	<tr><td>id</td><td>The ID that Amazon Route 53 assigned to the hosted zone when it was created.</td></tr>
	<tr><td>caller_reference</td><td>The value that you specified for CallerReference when you created the hosted zone.</td></tr>
	<tr><td>comment</td><td>A comment for the zone.</td></tr>
	<tr><td>private_zone</td><td>If true, the zone is Private hosted Zone, otherwise it is public.</td></tr>
	<tr><td>linked_service_principal</td><td>If the health check or hosted zone was created by another service, the service that created the resource.</td></tr>
	<tr><td>linked_service_description</td><td>If the health check or hosted zone was created by another service, an optional description that can be provided by the other service.</td></tr>
	<tr><td>resource_record_set_count</td><td>The number of resource record sets in the hosted zone.</td></tr>
	<tr><td>query_logging_configs</td><td>A list of configuration for DNS query logging that is associated with the current AWS account.</td></tr>
	<tr><td>dnssec_key_signing_keys</td><td>The key-signing keys (KSKs) in AWS account.</td></tr>
	<tr><td>dnssec_status</td><td>The status of DNSSEC.</td></tr>
	<tr><td>tags_src</td><td>A map of tags for the resource.</td></tr>
	<tr><td>vpcs</td><td>The list of VPCs that are authorized to be associated with the specified hosted zone.</td></tr>
	<tr><td>title</td><td>Title of the resource.</td></tr>
	<tr><td>tags</td><td>A map of tags for the resource.</td></tr>
	<tr><td>akas</td><td>Array of globally unique identifier strings (also known as) for the resource.</td></tr>
	<tr><td>partition</td><td>The AWS partition in which the resource is located (aws, aws-cn, or aws-us-gov).</td></tr>
	<tr><td>region</td><td>The AWS Region in which the resource is located.</td></tr>
	<tr><td>account_id</td><td>The AWS Account ID in which the resource is located.</td></tr>
	<tr><td>kaytu_account_id</td><td>The Kaytu Account ID in which the resource is located.</td></tr>
	<tr><td>kaytu_resource_id</td><td>The unique ID of the resource in opengovernance.</td></tr>
	<tr><td>kaytu_metadata</td><td>Kaytu Metadata of the AWS resource.</td></tr>
</table>