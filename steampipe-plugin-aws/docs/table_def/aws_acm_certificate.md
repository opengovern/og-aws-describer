# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>certificate_arn</td><td>Amazon Resource Name (ARN) of the certificate. This is of the form: arn:aws:acm:region:123456789012:certificate/12345678-1234-1234-1234-123456789012</td></tr>
	<tr><td>certificate</td><td>The ACM-issued certificate corresponding to the ARN specified as input</td></tr>
	<tr><td>certificate_chain</td><td>The ACM-issued certificate corresponding to the ARN specified as input</td></tr>
	<tr><td>domain_name</td><td>Fully qualified domain name (FQDN), such as www.example.com or example.com, for the certificate</td></tr>
	<tr><td>certificate_transparency_logging_preference</td><td>Indicates whether to opt in to or out of certificate transparency logging. Certificates that are not logged typically generate a browser error. Transparency makes it possible for you to detect SSL/TLS certificates that have been mistakenly or maliciously issued for your domain.</td></tr>
	<tr><td>created_at</td><td>The time at which the certificate was requested. This value exists only when the certificate type is AMAZON_ISSUED</td></tr>
	<tr><td>subject</td><td>The name of the entity that is associated with the public key contained in the certificate</td></tr>
	<tr><td>imported_at</td><td>The name of the certificate authority that issued and signed the certificate</td></tr>
	<tr><td>issuer</td><td>The name of the certificate authority that issued and signed the certificate</td></tr>
	<tr><td>signature_algorithm</td><td>The algorithm that was used to sign the certificate</td></tr>
	<tr><td>extended_key_usages</td><td>Specify one or more ExtendedKeyUsage extension values.</td></tr>
	<tr><td>failure_reason</td><td>The reason the certificate request failed. This value exists only when the certificate status is FAILED</td></tr>
	<tr><td>issued_at</td><td>A list of ARNs for the AWS resources that are using the certificate. A certificate can be used by multiple AWS resources</td></tr>
	<tr><td>status</td><td>The status of the certificate</td></tr>
	<tr><td>key_algorithm</td><td>The algorithm that was used to generate the public-private key pair</td></tr>
	<tr><td>not_after</td><td>The time after which the certificate is not valid</td></tr>
	<tr><td>not_before</td><td>The time before which the certificate is not valid</td></tr>
	<tr><td>renewal_eligibility</td><td>Specifies whether the certificate is eligible for renewal.</td></tr>
	<tr><td>revocation_reason</td><td>The reason the certificate was revoked. This value exists only when the certificate status is REVOKED</td></tr>
	<tr><td>revoked_at</td><td>The time at which the certificate was revoked. This value exists only when the certificate status is REVOKED</td></tr>
	<tr><td>serial</td><td>The serial number of the certificate</td></tr>
	<tr><td>type</td><td>The source of the certificate. For certificates provided by ACM, this value is AMAZON_ISSUED.</td></tr>
	<tr><td>domain_validation_options</td><td>Contains information about the initial validation of each domain name that occurs as a result of the RequestCertificate request. This field exists only when the certificate type is AMAZON_ISSUED</td></tr>
	<tr><td>in_use_by</td><td>A list of ARNs for the AWS resources that are using the certificate</td></tr>
	<tr><td>subject_alternative_names</td><td>One or more domain names (subject alternative names) included in the certificate. This list contains the domain names that are bound to the public key that is contained in the certificate. The subject alternative names include the canonical domain name (CN) of the certificate and additional domain names that can be used to connect to the website</td></tr>
	<tr><td>tags_src</td><td>A list of tags associated with certificate</td></tr>
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