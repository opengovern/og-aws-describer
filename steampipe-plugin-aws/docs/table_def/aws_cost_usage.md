# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>dimension_1</td><td>Valid values are AZ, INSTANCE_TYPE, LINKED_ACCOUNT, OPERATION, PURCHASE_TYPE, SERVICE, USAGE_TYPE, PLATFORM, TENANCY, RECORD_TYPE, LEGAL_ENTITY_NAME, DEPLOYMENT_OPTION, DATABASE_ENGINE, CACHE_ENGINE, INSTANCE_TYPE_FAMILY, REGION, BILLING_ENTITY, RESERVATION_ID, SAVINGS_PLANS_TYPE, SAVINGS_PLAN_ARN, OPERATING_SYSTEM</td></tr>
	<tr><td>dimension_2</td><td>Valid values are AZ, INSTANCE_TYPE, LINKED_ACCOUNT, OPERATION, PURCHASE_TYPE, SERVICE, USAGE_TYPE, PLATFORM, TENANCY, RECORD_TYPE, LEGAL_ENTITY_NAME, DEPLOYMENT_OPTION, DATABASE_ENGINE, CACHE_ENGINE, INSTANCE_TYPE_FAMILY, REGION, BILLING_ENTITY, RESERVATION_ID, SAVINGS_PLANS_TYPE, SAVINGS_PLAN_ARN, OPERATING_SYSTEM</td></tr>
	<tr><td>granularity</td><td></td></tr>
	<tr><td>search_start_time</td><td></td></tr>
	<tr><td>search_end_time</td><td></td></tr>
	<tr><td>dimension_type_1</td><td></td></tr>
	<tr><td>dimension_type_2</td><td></td></tr>
	<tr><td>period_start</td><td>Start timestamp for this cost metric.</td></tr>
	<tr><td>period_end</td><td>End timestamp for this cost metric.</td></tr>
	<tr><td>estimated</td><td>Whether the result is estimated.</td></tr>
	<tr><td>blended_cost_amount</td><td>This cost metric reflects the average cost of usage across the consolidated billing family. If you use the consolidated billing feature in AWS Organizations, you can view costs using blended rates.</td></tr>
	<tr><td>blended_cost_unit</td><td>Unit type for blended costs.</td></tr>
	<tr><td>unblended_cost_amount</td><td>Unblended costs represent your usage costs on the day they are charged to you. In finance terms, they represent your costs on a cash basis of accounting.</td></tr>
	<tr><td>unblended_cost_unit</td><td>Unit type for unblended costs.</td></tr>
	<tr><td>net_unblended_cost_amount</td><td>This cost metric reflects the unblended cost after discounts.</td></tr>
	<tr><td>net_unblended_cost_unit</td><td>Unit type for net unblended costs.</td></tr>
	<tr><td>amortized_cost_amount</td><td>This cost metric reflects the effective cost of the upfront and monthly reservation fees spread across the billing period. By default, Cost Explorer shows the fees for Reserved Instances as a spike on the day that you&#39;re charged, but if you choose to show costs as amortized costs, the costs are amortized over the billing period. This means that the costs are broken out into the effective daily rate. AWS estimates your amortized costs by combining your unblended costs with the amortized portion of your upfront and recurring reservation fees.</td></tr>
	<tr><td>amortized_cost_unit</td><td>Unit type for amortized costs.</td></tr>
	<tr><td>net_amortized_cost_amount</td><td>This cost metric amortizes the upfront and monthly reservation fees while including discounts such as RI volume discounts.</td></tr>
	<tr><td>net_amortized_cost_unit</td><td>Unit type for net amortized costs.</td></tr>
	<tr><td>usage_quantity_amount</td><td>The amount of usage that you incurred. NOTE: If you return the UsageQuantity metric, the service aggregates all usage numbers without taking into account the units. For example, if you aggregate usageQuantity across all of Amazon EC2, the results aren&#39;t meaningful because Amazon EC2 compute hours and data transfer are measured in different units (for example, hours vs. GB).</td></tr>
	<tr><td>usage_quantity_unit</td><td>Unit type for usage quantity.</td></tr>
	<tr><td>normalized_usage_amount</td><td>The amount of usage that you incurred, in normalized units, for size-flexible RIs. The NormalizedUsageAmount is equal to UsageAmount multiplied by NormalizationFactor.</td></tr>
	<tr><td>normalized_usage_unit</td><td>Unit type for normalized usage.</td></tr>
	<tr><td>partition</td><td>The AWS partition in which the resource is located (aws, aws-cn, or aws-us-gov).</td></tr>
	<tr><td>region</td><td>The AWS Region in which the resource is located.</td></tr>
	<tr><td>account_id</td><td>The AWS Account ID in which the resource is located.</td></tr>
</table>