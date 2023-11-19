<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>name</td><td>The name of the Web ACL. You cannot change the name of a Web ACL after you create it.</td></tr>
	<tr><td>arn</td><td>The Amazon Resource Name (ARN) of the entity.</td></tr>
	<tr><td>associated_resources</td><td>The array of Amazon Resource Names (ARNs) of the associated resources.</td></tr>
	<tr><td>id</td><td>The unique identifier for the Web ACL.</td></tr>
	<tr><td>scope</td><td>Specifies the scope of the Web ACL. Possibles values are: 'REGIONAL' and 'CLOUDFRONT'.</td></tr>
	<tr><td>description</td><td>A description of the Web ACL that helps with identification.</td></tr>
	<tr><td>capacity</td><td>The Web ACL capacity units(WCUs) currently being used by this resource.</td></tr>
	<tr><td>lock_token</td><td>A token used for optimistic locking.</td></tr>
	<tr><td>managed_by_firewall_manager</td><td>Indicates whether this web ACL is managed by AWS Firewall Manager.</td></tr>
	<tr><td>default_action</td><td>The action to perform if none of the Rules contained in the Web ACL match.</td></tr>
	<tr><td>logging_configuration</td><td>The logging configuration for the specified web ACL.</td></tr>
	<tr><td>pre_process_firewall_manager_rule_groups</td><td>The first set of rules for AWS WAF to process in the web ACL.</td></tr>
	<tr><td>post_process_firewall_manager_rule_groups</td><td>The last set of rules for AWS WAF to process in the web ACL.</td></tr>
	<tr><td>rules</td><td>The Rule statements used to identify the web requests that you want to allow, block, or count.</td></tr>
	<tr><td>visibility_config</td><td>Defines and enables Amazon CloudWatch metrics and web request sample collection.</td></tr>
	<tr><td>tags_src</td><td>A list of tags associated with the resource.</td></tr>
	<tr><td>title</td><td>Title of the resource.</td></tr>
	<tr><td>tags</td><td>A map of tags for the resource.</td></tr>
	<tr><td>akas</td><td>Array of globally unique identifier strings (also known as) for the resource.</td></tr>
	<tr><td>region</td><td>The AWS Region in which the resource is located.</td></tr>
	<tr><td>partition</td><td>The AWS partition in which the resource is located (aws, aws-cn, or aws-us-gov).</td></tr>
	<tr><td>account_id</td><td>The AWS Account ID in which the resource is located.</td></tr>
	<tr><td>kaytu_account_id</td><td>The Kaytu Account ID in which the resource is located.</td></tr>
	<tr><td>kaytu_resource_id</td><td>The unique ID of the resource in Kaytu.</td></tr>
	<tr><td>kaytu_metadata</td><td>kaytu Metadata of the AWS resource.</td></tr>
</table>