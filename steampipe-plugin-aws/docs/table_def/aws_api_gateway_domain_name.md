# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>domain_name</td><td>The custom domain name as an API host name.</td></tr>
	<tr><td>certificate_arn</td><td>The reference to an AWS-managed certificate that will be used by edge-optimized endpoint for this domain name.</td></tr>
	<tr><td>certificate_name</td><td>The name of the certificate that will be used by edge-optimized endpoint for this domain name.</td></tr>
	<tr><td>certificate_upload_date</td><td>The timestamp when the certificate that was used by edge-optimized endpoint for this domain name was uploaded.</td></tr>
	<tr><td>distribution_domain_name</td><td>The domain name of the Amazon CloudFront distribution associated with this custom domain name for an edge-optimized endpoint.</td></tr>
	<tr><td>distribution_hosted_zone_id</td><td>The region-agnostic Amazon Route 53 Hosted Zone ID of the edge-optimized endpoint. The valid value is Z2FDTNDATAQYW2 for all the regions.</td></tr>
	<tr><td>domain_name_status</td><td>The status of the DomainName migration. The valid values are AVAILABLE and UPDATING. If the status is UPDATING, the domain cannot be modified further until the existing operation is complete.</td></tr>
	<tr><td>domain_name_status_message</td><td>An optional text message containing detailed information about status of the DomainName migration.</td></tr>
	<tr><td>ownership_verification_certificate_arn</td><td>The ARN of the public certificate issued by ACM to validate ownership of your custom domain. Only required when configuring mutual TLS and using an ACM imported or private CA certificate ARN as the regionalCertificateArn.</td></tr>
	<tr><td>regional_certificate_arn</td><td>The reference to an AWS-managed certificate that will be used for validating the regional domain name. AWS Certificate Manager is the only supported source.</td></tr>
	<tr><td>regional_certificate_name</td><td>The name of the certificate that will be used for validating the regional domain name.</td></tr>
	<tr><td>regional_domain_name</td><td>The domain name associated with the regional endpoint for this custom domain name. You set up this association by adding a DNS record that points the custom domain name to this regional domain name. The regional domain name is returned by API Gateway when you create a regional endpoint.</td></tr>
	<tr><td>regional_hosted_zone_id</td><td>The region-specific Amazon Route 53 Hosted Zone ID of the regional endpoint. For more information, see Set up a Regional Custom Domain Name and AWS Regions and Endpoints for API Gateway.</td></tr>
	<tr><td>security_policy</td><td>The Transport Layer Security (TLS) version + cipher suite for this DomainName. The valid values are TLS_1_0 and TLS_1_2.</td></tr>
	<tr><td>endpoint_configuration</td><td>The endpoint configuration of this DomainName showing the endpoint types of the domain name.</td></tr>
	<tr><td>mutual_tls_authentication</td><td>The mutual TLS authentication configuration for a custom domain name. If specified, API Gateway performs two-way authentication between the client and the server. Clients must present a trusted certificate to access your API.</td></tr>
	<tr><td>tags</td><td>A map of tags for the resource.</td></tr>
	<tr><td>title</td><td>Title of the resource.</td></tr>
	<tr><td>akas</td><td>Array of globally unique identifier strings (also known as) for the resource.</td></tr>
	<tr><td>partition</td><td>The AWS partition in which the resource is located (aws, aws-cn, or aws-us-gov).</td></tr>
	<tr><td>region</td><td>The AWS Region in which the resource is located.</td></tr>
	<tr><td>account_id</td><td>The AWS Account ID in which the resource is located.</td></tr>
</table>