package aws

import (
	"context"
	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsWellArchitectedWorkload(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_wellarchitected_workload",
		Description: "AWS Well-Architected Workload",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("workload_id"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"ResourceNotFoundException"}),
			},
			Hydrate: kaytu.GetWellArchitectedWorkload,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListWellArchitectedWorkload,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "workload_name", Require: plugin.Optional},
			},
		},

		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "workload_name",
				Description: "The name of the workload.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Workload.WorkloadName")},
			{
				Name:        "workload_arn",
				Description: "The ARN for the workload.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Workload.WorkloadArn")},
			{
				Name:        "workload_id",
				Description: "The ID assigned to the workload.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Workload.WorkloadId")},
			{
				Name:        "architectural_design",
				Description: "The URL of the architectural design for the workload.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Workload.ArchitecturalDesign")},
			{
				Name:        "description",
				Description: "The description for the workload.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Workload.Description")},
			{
				Name:        "environment",
				Description: "The environment for the workload.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Workload.Environment")},
			{
				Name:        "improvement_status",
				Description: "The improvement status for a workload.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Workload.ImprovementStatus")},
			{
				Name:        "industry",
				Description: "The industry for the workload.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Workload.Industry")},
			{
				Name:        "industry_type",
				Description: "The industry type for the workload.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Workload.IndustryType")},
			{
				Name:        "is_review_owner_update_acknowledged",
				Description: "Flag indicating whether the workload owner has acknowledged that the review owner field is required.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Workload.IsReviewOwnerUpdateAcknowledged")},
			{
				Name:        "notes",
				Description: "The notes associated with the workload.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Workload.Notes")},
			{
				Name:        "owner",
				Description: "An AWS account ID.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Workload.Owner")},
			{
				Name:        "review_owner",
				Description: "The review owner of the workload.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Workload.ReviewOwner")},
			{
				Name:        "review_restriction_date",
				Description: "The date and time recorded.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Workload.ReviewRestrictionDate")},
			{
				Name:        "share_invitation_id",
				Description: "The ID assigned to the share invitation.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Workload.ShareInvitationId")},
			{
				Name:        "updated_at",
				Description: "The date and time recorded.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Workload.UpdatedAt")},
			{
				Name:        "account_ids",
				Description: "The list of AWS account IDs associated with the workload.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Workload.AccountIds")},
			{
				Name:        "aws_regions",
				Description: "The list of AWS Regions associated with the workload, for example, us-east-2, or ca-central-1.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Workload.AwsRegions")},
			{
				Name:        "lenses",
				Description: "The list of lenses associated with the workload. Each lens is identified by its LensSummary$LensAlias.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Workload.Lenses")},
			{
				Name:        "non_aws_regions",
				Description: "The list of non-AWS Regions associated with the workload.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Workload.NonAwsRegions")},
			{
				Name:        "pillar_priorities",
				Description: "The priorities of the pillars, which are used to order items in the improvement plan. ",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Workload.PillarPriorities")},
			{
				Name:        "risk_counts",
				Description: "A map from risk names to the count of how questions have that rating.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Workload.RiskCounts")},

			// Standard columns

			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Workload.WorkloadName")},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Workload.Tags")},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,

				Transform: transform.FromField("Description.Workload.WorkloadArn").Transform(arnToAkas),
			},
		}),
	}
}
