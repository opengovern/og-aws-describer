package aws

import (
	"context"
	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsAuditManagerEvidence(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_auditmanager_evidence",
		Description: "AWS Audit Manager Evidence",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"id", "evidence_folder_id", "assessment_id", "control_set_id"}),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"ResourceNotFoundException", "InvalidParameter"}),
			},
			Hydrate: opengovernance.GetAuditManagerEvidence,
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListAuditManagerEvidence,
		},

		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "id",
				Description: "The identifier for the evidence.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Evidence.Id"),
			},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) specifying the evidence.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("ARN").Transform(arnToAkas),
			},
			{
				Name:        "assessment_id",
				Description: "An unique identifier for the assessment.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AssessmentID")},
			{
				Name:        "control_set_id",
				Description: "The identifier for the control set.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ControlSetID")},
			{
				Name:        "evidence_folder_id",
				Description: "The identifier for the folder in which the evidence is stored.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Evidence.EvidenceFolderId")},
			{
				Name:        "assessment_report_selection",
				Description: "Specifies whether the evidence is included in the assessment report.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Evidence.AssessmentReportSelection")},
			{
				Name:        "aws_account_id",
				Description: "The identifier for the specified AWS account.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Evidence.AwsAccountId")},
			{
				Name:        "aws_organization",
				Description: "The AWS account from which the evidence is collected, and its AWS organization path.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Evidence.AwsOrganization")},
			{
				Name:        "compliance_check",
				Description: "The evaluation status for evidence that falls under the compliance check category.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Evidence.ComplianceCheck")},
			{
				Name:        "data_source",
				Description: "The data source from which the specified evidence was collected.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Evidence.DataSource")},
			{
				Name:        "event_name",
				Description: "The name of the specified evidence event.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Evidence.EventName")},
			{
				Name:        "event_source",
				Description: "The AWS service from which the evidence is collected.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Evidence.EventSource")},
			{
				Name:        "evidence_aws_account_id",
				Description: "The identifier for the specified AWS account.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Evidence.EvidenceAwsAccountId")},
			{
				Name:        "evidence_by_type",
				Description: "The type of automated evidence.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Evidence.EvidenceByType")},
			{
				Name:        "iam_id",
				Description: "The unique identifier for the IAM user or role associated with the evidence.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Evidence.IamId")},
			{
				Name:        "time",
				Description: "The timestamp that represents when the evidence was collected.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Evidence.Time")},
			{
				Name:        "attributes",
				Description: "The names and values used by the evidence event",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Evidence.Attributes")},
			{
				Name:        "resources_included",
				Description: "The list of resources assessed to generate the evidence.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Evidence.ResourcesIncluded")},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Evidence.Id")},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("ARN").Transform(transform.EnsureStringArray),
			},
		}),
	}
}
