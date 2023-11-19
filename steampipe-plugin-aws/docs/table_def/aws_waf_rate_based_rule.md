# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>name</td><td>The name for the rule.</td></tr>
	<tr><td>rule_id</td><td>The ID of the Rule.</td></tr>
	<tr><td>metric_name</td><td>The name or description for the metrics for a RateBasedRule.</td></tr>
	<tr><td>rate_key</td><td>The field that AWS WAF uses to determine if requests are likely arriving from single source and thus subject to rate monitoring.</td></tr>
	<tr><td>rate_limit</td><td>The maximum number of requests, which have an identical value in the field specified by the RateKey, allowed in a five-minute period.</td></tr>
	<tr><td>predicates</td><td>The Predicates object contains one Predicate element for each ByteMatchSet, IPSet or SqlInjectionMatchSet object that you want to include in a RateBasedRule.</td></tr>
	<tr><td>tags_src</td><td>A list of tags assigned to the Rule.</td></tr>
	<tr><td>title</td><td>Title of the resource.</td></tr>
	<tr><td>tags</td><td>A map of tags for the resource.</td></tr>
	<tr><td>akas</td><td>Array of globally unique identifier strings (also known as) for the resource.</td></tr>
	<tr><td>partition</td><td>The AWS partition in which the resource is located (aws, aws-cn, or aws-us-gov).</td></tr>
	<tr><td>region</td><td>The AWS Region in which the resource is located.</td></tr>
	<tr><td>account_id</td><td>The AWS Account ID in which the resource is located.</td></tr>
	<tr><td>kaytu_account_id</td><td>The Kaytu Account ID in which the resource is located.</td></tr>
	<tr><td>kaytu_resource_id</td><td>The unique ID of the resource in Kaytu.</td></tr>
	<tr><td>kaytu_metadata</td><td>Kaytu Metadata of the AWS resource.</td></tr>
</table>