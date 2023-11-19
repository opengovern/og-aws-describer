<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>rule_group_name</td><td>The descriptive name of the rule group.</td></tr>
	<tr><td>arn</td><td>The Amazon Resource Name (ARN) of the rule group.</td></tr>
	<tr><td>capacity</td><td>The maximum operating resources that this rule group can use. Rule group capacity is fixed at creation. When you update a rule group, you are limited to this capacity. When you reference a rule group from a firewall policy, Network Firewall reserves this capacity for the rule group.</td></tr>
	<tr><td>consumed_capacity</td><td>The number of capacity units currently consumed by the rule group rules.</td></tr>
	<tr><td>description</td><td>A description of the rule group.</td></tr>
	<tr><td>number_of_associations</td><td>The number of firewall policies that use this rule group.</td></tr>
	<tr><td>rule_group_id</td><td>The unique identifier for the rule group.</td></tr>
	<tr><td>rule_group_status</td><td>Detailed information about the current status of a rule group.</td></tr>
	<tr><td>rule_variables</td><td>Settings that are available for use in the rules in the rule group. You can only use these for stateful rule groups.</td></tr>
	<tr><td>rules_source</td><td>The stateful rules or stateless rules for the rule group.</td></tr>
	<tr><td>stateful_rule_options</td><td>Additional options governing how Network Firewall handles the rule group. You can only use these for stateful rule groups.</td></tr>
	<tr><td>type</td><td>Indicates whether the rule group is stateless or stateful. If the rule group is stateless, it contains stateless rules. If it is stateful, it contains stateful rules.</td></tr>
	<tr><td>tags_src</td><td>A list of tags assigned to the resource.</td></tr>
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