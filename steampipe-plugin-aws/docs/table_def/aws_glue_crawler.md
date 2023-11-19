# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>name</td><td>The name of the crawler.</td></tr>
	<tr><td>arn</td><td>The ARN of the crawler.</td></tr>
	<tr><td>database_name</td><td>The name of the database in which the crawler's output is stored.</td></tr>
	<tr><td>state</td><td>Indicates whether the crawler is running or pending.</td></tr>
	<tr><td>role</td><td>The Amazon Resource Name (ARN) of an IAM role that's used to access customer resources, such as Amazon Simple Storage Service (Amazon S3) data.</td></tr>
	<tr><td>creation_time</td><td>The time that the crawler was created.</td></tr>
	<tr><td>description</td><td>A description of the crawler.</td></tr>
	<tr><td>crawl_elapsed_time</td><td>If the crawler is running, contains the total time elapsed since the last crawl began.</td></tr>
	<tr><td>crawler_lineage_settings</td><td>Specifies whether data lineage is enabled for the crawler.</td></tr>
	<tr><td>crawler_security_configuration</td><td>The name of the SecurityConfiguration structure to be used by this crawler.</td></tr>
	<tr><td>last_updated</td><td>The time that the crawler was last updated.</td></tr>
	<tr><td>recrawl_behavior</td><td>Specifies whether to crawl the entire dataset again or to crawl only folders that were added since the last crawler run. A value of CRAWL_EVERYTHING specifies crawling the entire dataset again. A value of CRAWL_NEW_FOLDERS_ONLY specifies crawling only folders that were added since the last crawler run. A value of CRAWL_EVENT_MODE specifies crawling only the changes identified by Amazon S3 events.</td></tr>
	<tr><td>table_prefix</td><td>The prefix added to the names of tables that are created.</td></tr>
	<tr><td>version</td><td>The version of the crawler.</td></tr>
	<tr><td>classifiers</td><td>A list of UTF-8 strings that specify the custom classifiers that are associated with the crawler.</td></tr>
	<tr><td>configuration</td><td>Crawler configuration information.</td></tr>
	<tr><td>last_crawl</td><td>The status of the last crawl, and potentially error information if an error occurred.</td></tr>
	<tr><td>schedule</td><td>For scheduled crawlers, the schedule when the crawler runs.</td></tr>
	<tr><td>schema_change_policy</td><td>The policy that specifies update and delete behaviors for the crawler.</td></tr>
	<tr><td>targets</td><td>A collection of targets to crawl.</td></tr>
	<tr><td>title</td><td>Title of the resource.</td></tr>
	<tr><td>akas</td><td>Array of globally unique identifier strings (also known as) for the resource.</td></tr>
	<tr><td>partition</td><td>The AWS partition in which the resource is located (aws, aws-cn, or aws-us-gov).</td></tr>
	<tr><td>region</td><td>The AWS Region in which the resource is located.</td></tr>
	<tr><td>account_id</td><td>The AWS Account ID in which the resource is located.</td></tr>
	<tr><td>kaytu_account_id</td><td>The Kaytu Account ID in which the resource is located.</td></tr>
	<tr><td>kaytu_resource_id</td><td>The unique ID of the resource in Kaytu.</td></tr>
	<tr><td>kaytu_metadata</td><td>kaytu Metadata of the AWS resource.</td></tr>
</table>