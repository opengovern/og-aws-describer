package aws

import (
	"context"
	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsSSMPatchBaseline(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_ssm_patch_baseline",
		Description: "AWS SSM Patch Baseline",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("baseline_id"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"DoesNotExistException", "InvalidResourceId", "InvalidParameter", "ValidationException"}),
			},
			Hydrate: kaytu.GetSSMPatchBaseline,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListSSMPatchBaseline,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "name", Require: plugin.Optional},
				{Name: "operating_system", Require: plugin.Optional},
			},
		},

		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the patch baseline.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.PatchBaselineIdentity.BaselineName")},
			{
				Name:        "baseline_id",
				Description: "The ID of the retrieved patch baseline.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.PatchBaseline.BaselineId").Transform(lastPathElement),
			},
			{
				Name:        "description",
				Description: "A description of the patch baseline.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.PatchBaseline.Description")},
			{
				Name:        "operating_system",
				Description: "Returns the operating system specified for the patch baseline.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.PatchBaseline.OperatingSystem")},
			{
				Name:        "created_date",
				Description: "The date the patch baseline was created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.PatchBaseline.CreatedDate")},
			{
				Name:        "modified_date",
				Description: "The date the patch baseline was last modified.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.PatchBaseline.ModifiedDate")},
			{
				Name:        "approved_patches_compliance_level",
				Description: "Returns the specified compliance severity level for approved patches in the patch baseline.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.PatchBaseline.ApprovedPatchesComplianceLevel")},
			{
				Name:        "approved_patches_enable_non_security",
				Description: "Indicates whether the list of approved patches includes non-security updates that should be applied to the instances. The default value is 'false'. Applies to Linux instances only.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.PatchBaseline.ApprovedPatchesEnableNonSecurity")},
			{
				Name:        "approval_rules",
				Description: "A set of rules used to include patches in the baseline.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.PatchBaseline.ApprovalRules")},
			{
				Name:        "approved_patches",
				Description: "A list of explicitly approved patches for the baseline.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.PatchBaseline.ApprovedPatches")},
			{
				Name:        "global_filters",
				Description: "A set of global filters used to exclude patches from the baseline.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.PatchBaseline.GlobalFilters")},
			{
				Name:        "patch_groups",
				Description: "Patch groups included in the patch baseline.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.PatchBaseline.PatchGroups")},
			{
				Name:        "rejected_patches_action",
				Description: "The action specified to take on patches included in the RejectedPatches list. A patch can be allowed only if it is a dependency of another package, or blocked entirely along with packages that include it as a dependency.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.PatchBaseline.RejectedPatchesAction")},
			{
				Name:        "rejected_patches",
				Description: "A list of explicitly rejected patches for the baseline.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.PatchBaseline.RejectedPatches")},
			{
				Name:        "sources",
				Description: "Information about the patches to use to update the instances, including target operating systems and source repositories. Applies to Linux instances only.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.PatchBaseline.Sources")},
			{
				Name:        "tags_src",
				Description: "A list of tags assigned to the patch baseline.",
				Type:        proto.ColumnType_JSON,

				Transform: transform.FromField("Description.Tags")},

			// Standard columns for all tables

			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.PatchBaselineIdentity.BaselineName")},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,

				Transform: transform.FromField("Description.Tags").Transform(ssmTagListToTurbotTags),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ARN")},

			//error
			// check this in elastics data
			// TODO: The below mention field is coming from list call, but not from get call.
			// Need to check, if there is another way to fetch this value.

			// {
			// 	Name:        "default_baseline",
			// 	Description: "Whether this is the default baseline. Note that Systems Manager supports creating multiple default patch baselines. For example, you can create a default patch baseline for each operating system.",
			// 	Type:        proto.ColumnType_BOOL,
			// },
		}),
	}
}
