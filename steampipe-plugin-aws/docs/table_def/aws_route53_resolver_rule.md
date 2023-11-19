# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>name</td><td>The name for the Resolver rule, which you specified when you created the Resolver rule.</td></tr>
	<tr><td>id</td><td>The ID that Resolver assigned to the Resolver rule when you created it.</td></tr>
	<tr><td>arn</td><td>The ARN (Amazon Resource Name) for the Resolver rule specified by Id.</td></tr>
	<tr><td>status</td><td>A code that specifies the current status of the Resolver rule.</td></tr>
	<tr><td>creator_request_id</td><td>A unique string that you specified when you created the Resolver rule. CreatorRequestId identifies the request and allows failed requests to be retried without the risk of executing the operation twice.</td></tr>
	<tr><td>domain_name</td><td>DNS queries for this domain name are forwarded to the IP addresses that are specified in TargetIps.</td></tr>
	<tr><td>owner_id</td><td>When a rule is shared with another AWS account, the account ID of the account that the rule is shared with.</td></tr>
	<tr><td>resolver_endpoint_id</td><td>The ID of the endpoint that the rule is associated with.</td></tr>
	<tr><td>rule_type</td><td>When you want to forward DNS queries for specified domain name to resolvers on your network, specify FORWARD.When you have a forwarding rule to forward DNS queries for a domain to your network and you want Resolver to process queries for a subdomain of that domain, specify SYSTEM.</td></tr>
	<tr><td>share_status</td><td>Indicates whether the rules is shared and, if so, whether the current account is sharing the rule with another account, or another account is sharing the rule with the current account.</td></tr>
	<tr><td>status_message</td><td>A detailed description of the status of a Resolver rule.</td></tr>
	<tr><td>creation_time</td><td>The date and time that the Resolver rule was created, in Unix time format and Coordinated Universal Time (UTC).</td></tr>
	<tr><td>modification_time</td><td>The date and time that the Resolver rule was last updated, in Unix time format and Coordinated Universal Time (UTC).</td></tr>
	<tr><td>resolver_rule_associations</td><td>The associations that were created between Resolver rules and VPCs using the current AWS account, and that match the specified filters, if any.</td></tr>
	<tr><td>target_ips</td><td>An array that contains the IP addresses and ports that an outbound endpoint forwards DNS queries to. Typically, these are the IP addresses of DNS resolvers on your network. Specify IPv4 addresses. IPv6 is not supported.</td></tr>
	<tr><td>tags_src</td><td>A list of tags assigned to the Resolver Rule.</td></tr>
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