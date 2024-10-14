package aws

import (
	"context"
	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsSSMManagedInstancePatchState(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_ssm_managed_instance_patch_state",
		Description: "AWS SSM Managed Instance Patch State",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListSSMManagedInstancePatchState,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "instance_id", Require: plugin.Optional},
			},
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "instance_id",
				Description: "The ID of the managed node the high-level patch compliance information was collected for.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.PatchState.InstanceId")},
			{
				Name:        "baseline_id",
				Description: "The ID of the patch baseline used to patch the managed node.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.PatchState.BaselineId")},
			{
				Name:        "operation",
				Description: "The type of patching operation that was performed.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.PatchState.Operation")},
			{
				Name:        "operation_end_time",
				Description: "The time the most recent patching operation completed on the managed node.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.PatchState.OperationEndTime")},
			{
				Name:        "operation_start_time",
				Description: "The time the most recent patching operation was started on the managed node.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.PatchState.OperationStartTime")},
			{
				Name:        "patch_group",
				Description: "The name of the patch group the managed node belongs to.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.PatchState.PatchGroup")},
			{
				Name:        "critical_non_compliant_count",
				Description: "The number of patches per node that are specified as Critical for compliance reporting in the patch baseline aren't installed. These patches might be missing, have failed installation, were rejected, or were installed but awaiting a required managed node reboot. The status of these managed nodes is NON_COMPLIANT.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.PatchState.CriticalNonCompliantCount")},
			{
				Name:        "failed_count",
				Description: "The number of patches from the patch baseline that were attempted to be installed during the last patching operation, but failed to install.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.PatchState.FailedCount")},
			{
				Name:        "installed_count",
				Description: "The number of patches from the patch baseline that are installed on the managed node.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.PatchState.InstalledCount")},
			{
				Name:        "installed_other_count",
				Description: "The number of patches not specified in the patch baseline that are installed on the managed node.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.PatchState.InstalledOtherCount")},
			{
				Name:        "installed_pending_reboot_count",
				Description: "The number of patches installed by Patch Manager since the last time the managed node was rebooted.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.PatchState.InstalledPendingRebootCount")},
			{
				Name:        "installed_rejected_count",
				Description: "The number of patches installed on a managed node that are specified in a RejectedPatches list. Patches with a status of InstalledRejected were typically installed before they were added to a RejectedPatches list.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.PatchState.InstalledRejectedCount")},
			{
				Name:        "last_no_reboot_install_operation_time",
				Description: "The time of the last attempt to patch the managed node with NoReboot specified as the reboot option.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.PatchState.LastNoRebootInstallOperationTime")},
			{
				Name:        "missing_count",
				Description: "The number of patches from the patch baseline that are applicable for the managed node but aren't currently installed.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.PatchState.MissingCount")},
			{
				Name:        "not_applicable_count",
				Description: "The number of patches from the patch baseline that aren't applicable for the managed node and therefore aren't installed on the node. This number may be truncated if the list of patch names is very large. The number of patches beyond this limit are reported in UnreportedNotApplicableCount.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.PatchState.NotApplicableCount")},
			{
				Name:        "other_non_compliant_count",
				Description: "The number of patches per node that are specified as other than Critical or Security but aren't compliant with the patch baseline. The status of these managed nodes is NON_COMPLIANT.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.PatchState.OtherNonCompliantCount")},
			{
				Name:        "owner_information",
				Description: "Placeholder information. This field will always be empty in the current release of the service.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.PatchState.OwnerInformation")},
			{
				Name:        "reboot_option",
				Description: "Indicates the reboot option specified in the patch baseline. Reboot options apply to Install operations only. Reboots aren't attempted for Patch Manager Scan operations.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.PatchState.RebootOption")},
			{
				Name:        "security_non_compliant_count",
				Description: "The number of patches per node that are specified as Security in a patch advisory aren't installed. These patches might be missing, have failed installation, were rejected, or were installed but awaiting a required managed node reboot. The status of these managed nodes is NON_COMPLIANT.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.PatchState.SecurityNonCompliantCount")},
			{
				Name:        "snapshot_id",
				Description: "The ID of the patch baseline snapshot used during the patching operation when this compliance data was collected.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.PatchState.SnapshotId")},
			{
				Name:        "unreported_not_applicable_count",
				Description: "The number of patches beyond the supported limit of NotApplicableCount that aren't reported by name to Inventory. Inventory is a capability of Amazon Web Services Systems Manager.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.PatchState.UnreportedNotApplicableCount")},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.PatchState.BaselineId")},
		}),
	}
}
