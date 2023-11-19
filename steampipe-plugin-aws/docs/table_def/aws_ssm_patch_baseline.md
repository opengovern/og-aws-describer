# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>name</td><td>The name of the patch baseline.</td></tr>
	<tr><td>baseline_id</td><td>The ID of the retrieved patch baseline.</td></tr>
	<tr><td>description</td><td>A description of the patch baseline.</td></tr>
	<tr><td>operating_system</td><td>Returns the operating system specified for the patch baseline.</td></tr>
	<tr><td>created_date</td><td>The date the patch baseline was created.</td></tr>
	<tr><td>modified_date</td><td>The date the patch baseline was last modified.</td></tr>
	<tr><td>approved_patches_compliance_level</td><td>Returns the specified compliance severity level for approved patches in the patch baseline.</td></tr>
	<tr><td>approved_patches_enable_non_security</td><td>Indicates whether the list of approved patches includes non-security updates that should be applied to the instances. The default value is 'false'. Applies to Linux instances only.</td></tr>
	<tr><td>approval_rules</td><td>A set of rules used to include patches in the baseline.</td></tr>
	<tr><td>approved_patches</td><td>A list of explicitly approved patches for the baseline.</td></tr>
	<tr><td>global_filters</td><td>A set of global filters used to exclude patches from the baseline.</td></tr>
	<tr><td>patch_groups</td><td>Patch groups included in the patch baseline.</td></tr>
	<tr><td>rejected_patches_action</td><td>The action specified to take on patches included in the RejectedPatches list. A patch can be allowed only if it is a dependency of another package, or blocked entirely along with packages that include it as a dependency.</td></tr>
	<tr><td>rejected_patches</td><td>A list of explicitly rejected patches for the baseline.</td></tr>
	<tr><td>sources</td><td>Information about the patches to use to update the instances, including target operating systems and source repositories. Applies to Linux instances only.</td></tr>
	<tr><td>tags_src</td><td>A list of tags assigned to the patch baseline.</td></tr>
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