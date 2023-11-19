# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>instance_id</td><td>The ID of the managed node the high-level patch compliance information was collected for.</td></tr>
	<tr><td>baseline_id</td><td>The ID of the patch baseline used to patch the managed node.</td></tr>
	<tr><td>operation</td><td>The type of patching operation that was performed.</td></tr>
	<tr><td>operation_end_time</td><td>The time the most recent patching operation completed on the managed node.</td></tr>
	<tr><td>operation_start_time</td><td>The time the most recent patching operation was started on the managed node.</td></tr>
	<tr><td>patch_group</td><td>The name of the patch group the managed node belongs to.</td></tr>
	<tr><td>critical_non_compliant_count</td><td>The number of patches per node that are specified as Critical for compliance reporting in the patch baseline aren&#39;t installed. These patches might be missing, have failed installation, were rejected, or were installed but awaiting a required managed node reboot. The status of these managed nodes is NON_COMPLIANT.</td></tr>
	<tr><td>failed_count</td><td>The number of patches from the patch baseline that were attempted to be installed during the last patching operation, but failed to install.</td></tr>
	<tr><td>installed_count</td><td>The number of patches from the patch baseline that are installed on the managed node.</td></tr>
	<tr><td>installed_other_count</td><td>The number of patches not specified in the patch baseline that are installed on the managed node.</td></tr>
	<tr><td>installed_pending_reboot_count</td><td>The number of patches installed by Patch Manager since the last time the managed node was rebooted.</td></tr>
	<tr><td>installed_rejected_count</td><td>The number of patches installed on a managed node that are specified in a RejectedPatches list. Patches with a status of InstalledRejected were typically installed before they were added to a RejectedPatches list.</td></tr>
	<tr><td>last_no_reboot_install_operation_time</td><td>The time of the last attempt to patch the managed node with NoReboot specified as the reboot option.</td></tr>
	<tr><td>missing_count</td><td>The number of patches from the patch baseline that are applicable for the managed node but aren&#39;t currently installed.</td></tr>
	<tr><td>not_applicable_count</td><td>The number of patches from the patch baseline that aren&#39;t applicable for the managed node and therefore aren&#39;t installed on the node. This number may be truncated if the list of patch names is very large. The number of patches beyond this limit are reported in UnreportedNotApplicableCount.</td></tr>
	<tr><td>other_non_compliant_count</td><td>The number of patches per node that are specified as other than Critical or Security but aren&#39;t compliant with the patch baseline. The status of these managed nodes is NON_COMPLIANT.</td></tr>
	<tr><td>owner_information</td><td>Placeholder information. This field will always be empty in the current release of the service.</td></tr>
	<tr><td>reboot_option</td><td>Indicates the reboot option specified in the patch baseline. Reboot options apply to Install operations only. Reboots aren&#39;t attempted for Patch Manager Scan operations.</td></tr>
	<tr><td>security_non_compliant_count</td><td>The number of patches per node that are specified as Security in a patch advisory aren&#39;t installed. These patches might be missing, have failed installation, were rejected, or were installed but awaiting a required managed node reboot. The status of these managed nodes is NON_COMPLIANT.</td></tr>
	<tr><td>snapshot_id</td><td>The ID of the patch baseline snapshot used during the patching operation when this compliance data was collected.</td></tr>
	<tr><td>unreported_not_applicable_count</td><td>The number of patches beyond the supported limit of NotApplicableCount that aren&#39;t reported by name to Inventory. Inventory is a capability of Amazon Web Services Systems Manager.</td></tr>
	<tr><td>title</td><td>Title of the resource.</td></tr>
	<tr><td>partition</td><td>The AWS partition in which the resource is located (aws, aws-cn, or aws-us-gov).</td></tr>
	<tr><td>region</td><td>The AWS Region in which the resource is located.</td></tr>
	<tr><td>account_id</td><td>The AWS Account ID in which the resource is located.</td></tr>
</table>