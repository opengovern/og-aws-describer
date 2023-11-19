# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>name</td><td>The name of the Systems Manager document.</td></tr>
	<tr><td>account_ids</td><td>The account IDs that have permission to use this document.The ID can be either an AWS account or All.</td></tr>
	<tr><td>account_sharing_info_list</td><td>A list of AWS accounts where the current document is shared and the version shared with each account.</td></tr>
	<tr><td>arn</td><td>The Amazon Resource Name (ARN) of the document.</td></tr>
	<tr><td>approved_version</td><td>The version of the document currently approved for use in the organization.</td></tr>
	<tr><td>attachments_information</td><td>Details about the document attachments, including names, locations, sizes,and so on.</td></tr>
	<tr><td>author</td><td>The user in your organization who created the document.</td></tr>
	<tr><td>created_date</td><td>The date when the document was created.</td></tr>
	<tr><td>default_version</td><td>The default version.</td></tr>
	<tr><td>description</td><td>A description of the document.</td></tr>
	<tr><td>document_format</td><td>The document format, either JSON or YAML.</td></tr>
	<tr><td>document_type</td><td>The type of document.</td></tr>
	<tr><td>document_version</td><td>The document version.</td></tr>
	<tr><td>hash</td><td>The Sha256 or Sha1 hash created by the system when the document was created.</td></tr>
	<tr><td>hash_type</td><td>The hash type of the document.</td></tr>
	<tr><td>latest_version</td><td>The latest version of the document.</td></tr>
	<tr><td>owner_type</td><td>The AWS user account type to filter the documents. Possible values: Self, Amazon, Public, Private, ThirdParty, All, Default.</td></tr>
	<tr><td>owner</td><td>The AWS user account that created the document.</td></tr>
	<tr><td>parameters</td><td>A description of the parameters for a document.</td></tr>
	<tr><td>pending_review_version</td><td>The version of the document that is currently under review.</td></tr>
	<tr><td>platform_types</td><td>The operating system platform.</td></tr>
	<tr><td>requires</td><td>A list of SSM documents required by a document.</td></tr>
	<tr><td>review_information</td><td>Details about the review of a document.</td></tr>
	<tr><td>review_status</td><td>The current status of the review.</td></tr>
	<tr><td>schema_version</td><td>The schema version.</td></tr>
	<tr><td>sha1</td><td>The SHA1 hash of the document, which you can use for verification.</td></tr>
	<tr><td>status</td><td>The user in your organization who created the document.</td></tr>
	<tr><td>status_information</td><td>A message returned by AWS Systems Manager that explains the Status value.</td></tr>
	<tr><td>tags_src</td><td>A list of tags associated with document</td></tr>
	<tr><td>target_type</td><td>The target type which defines the kinds of resources the document can run on.</td></tr>
	<tr><td>version_name</td><td>The version of the artifact associated with the document.</td></tr>
	<tr><td>akas</td><td>Array of globally unique identifier strings (also known as) for the resource.</td></tr>
	<tr><td>tags</td><td>A map of tags for the resource.</td></tr>
	<tr><td>title</td><td>Title of the resource.</td></tr>
	<tr><td>partition</td><td>The AWS partition in which the resource is located (aws, aws-cn, or aws-us-gov).</td></tr>
	<tr><td>region</td><td>The AWS Region in which the resource is located.</td></tr>
	<tr><td>account_id</td><td>The AWS Account ID in which the resource is located.</td></tr>
	<tr><td>kaytu_account_id</td><td>The Kaytu Account ID in which the resource is located.</td></tr>
	<tr><td>kaytu_resource_id</td><td>The unique ID of the resource in Kaytu.</td></tr>
	<tr><td>kaytu_metadata</td><td>kaytu Metadata of the AWS resource.</td></tr>
</table>