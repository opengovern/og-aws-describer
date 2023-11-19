# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>id</td><td>The identifier for the Distribution.</td></tr>
	<tr><td>arn</td><td>The ARN (Amazon Resource Name) for the distribution.</td></tr>
	<tr><td>status</td><td>The current status of the Distribution.</td></tr>
	<tr><td>caller_reference</td><td>A unique value that ensures that the request can&#39;t be replayed.</td></tr>
	<tr><td>comment</td><td>The comment originally specified when this Distribution was created.</td></tr>
	<tr><td>default_root_object</td><td>The object that you want CloudFront to request from your origin.</td></tr>
	<tr><td>domain_name</td><td>The domain name that corresponds to the Distribution.</td></tr>
	<tr><td>enabled</td><td>Whether the Distribution is enabled to accept user requests for content.</td></tr>
	<tr><td>e_tag</td><td>The current version of the configuration.</td></tr>
	<tr><td>http_version</td><td>Specify the maximum HTTP version that you want viewers to use to communicate with CloudFront. The default value for new web Distributions is http2. Viewers that don&#39;t support HTTP/2 will automatically use an earlier version.</td></tr>
	<tr><td>is_ipv6_enabled</td><td>Whether CloudFront responds to IPv6 DNS requests with an IPv6 address for your Distribution.</td></tr>
	<tr><td>in_progress_invalidation_batches</td><td>The number of invalidation batches currently in progress.</td></tr>
	<tr><td>last_modified_time</td><td>The date and time the Distribution was last modified.</td></tr>
	<tr><td>price_class</td><td>A complex type that contains information about price class for this streaming Distribution.</td></tr>
	<tr><td>web_acl_id</td><td>The Web ACL Id (if any) associated with the distribution.</td></tr>
	<tr><td>active_trusted_key_groups</td><td>CloudFront automatically adds this field to the response if you’ve configured a cache behavior in this distribution to serve private content using key groups.</td></tr>
	<tr><td>active_trusted_signers</td><td>A list of AWS accounts and the identifiers of active CloudFront key pairs in each account that CloudFront can use to verify the signatures of signed URLs and signed cookies.</td></tr>
	<tr><td>aliases</td><td>A complex type that contains information about CNAMEs (alternate domain names),if any, for this distribution.</td></tr>
	<tr><td>alias_icp_recordals</td><td>AWS services in China customers must file for an Internet Content Provider (ICP) recordal if they want to serve content publicly on an alternate domain name, also known as a CNAME, that they&#39;ve added to CloudFront. AliasICPRecordal provides the ICP recordal status for CNAMEs associated with distributions.</td></tr>
	<tr><td>cache_behaviors</td><td>The number of cache behaviors for this Distribution.</td></tr>
	<tr><td>custom_error_responses</td><td>A complex type that contains zero or more CustomErrorResponses elements.</td></tr>
	<tr><td>default_cache_behavior</td><td>A complex type that describes the default cache behavior if you don&#39;t specify a CacheBehavior element or if files don&#39;t match any of the values of PathPattern in CacheBehavior elements. You must create exactly one default cache behavior.</td></tr>
	<tr><td>logging</td><td>A complex type that controls whether access logs are written for the distribution.</td></tr>
	<tr><td>origins</td><td>A complex type that contains information about origins for this distribution.</td></tr>
	<tr><td>origin_groups</td><td>A complex type that contains information about origin groups for this distribution.</td></tr>
	<tr><td>restrictions</td><td>A complex type that identifies ways in which you want to restrict distribution of your content.</td></tr>
	<tr><td>tags_src</td><td>A list of tags assigned to the Maintenance Window</td></tr>
	<tr><td>viewer_certificate</td><td>A complex type that determines the distribution&#39;s SSL/TLS configuration for communicating with viewers.</td></tr>
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