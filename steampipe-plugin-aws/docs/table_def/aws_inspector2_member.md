# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>member_account_id</td><td>The AWS Account ID in which the resource is located.</td></tr>
	<tr><td>delegated_admin_account_id</td><td>The Amazon Web Services account ID of the Amazon Inspector delegated administrator for this member account.</td></tr>
	<tr><td>only_associated</td><td>Specifies whether to list only currently associated members if True or to list all members within the organization if False.</td></tr>
	<tr><td>relationship_status</td><td>The status of the member account. Valid values are: CREATED | INVITED | DISABLED | ENABLED | REMOVED | RESIGNED | DELETED | EMAIL_VERIFICATION_IN_PROGRESS | EMAIL_VERIFICATION_FAILED | REGION_DISABLED | ACCOUNT_SUSPENDED | CANNOT_CREATE_DETECTOR_IN_ORG_MASTER.</td></tr>
	<tr><td>updated_at</td><td>A timestamp showing when the status of this member was last updated.</td></tr>
	<tr><td>title</td><td>Title of the resource.</td></tr>
	<tr><td>partition</td><td>The AWS partition in which the resource is located (aws, aws-cn, or aws-us-gov).</td></tr>
	<tr><td>region</td><td>The AWS Region in which the resource is located.</td></tr>
	<tr><td>account_id</td><td>The AWS Account ID in which the resource is located.</td></tr>
</table>