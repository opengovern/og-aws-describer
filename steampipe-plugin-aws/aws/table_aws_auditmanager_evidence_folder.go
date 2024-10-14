package aws

import (
	"context"

	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsAuditManagerEvidenceFolder(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_auditmanager_evidence_folder",
		Description: "AWS Audit Manager Evidence Folder",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"id", "assessment_id", "control_set_id"}),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"ResourceNotFoundException", "InvalidParameter"}),
			},
			Hydrate: opengovernance.GetAuditManagerEvidenceFolder,
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListAuditManagerEvidenceFolder,
		},

		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the specified evidence folder.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.EvidenceFolder.Name")},
			{
				Name:        "id",
				Description: "The identifier for the folder in which evidence is stored.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.EvidenceFolder.Id")},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) specifying the evidence folder.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("ARN").Transform(arnToAkas),
			},
			{
				Name:        "assessment_id",
				Description: "The identifier for the specified assessment.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AssessmentID")},
			{
				Name:        "control_set_id",
				Description: "The identifier for the control set.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.EvidenceFolder.ControlSetId")},
			{
				Name:        "assessment_report_selection_count",
				Description: "The total count of evidence included in the assessment report.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.EvidenceFolder.AssessmentReportSelectionCount")},
			{
				Name:        "author",
				Description: "The name of the user who created the evidence folder.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.EvidenceFolder.Author")},
			{
				Name:        "control_id",
				Description: "The unique identifier for the specified control.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.EvidenceFolder.ControlId")},
			{
				Name:        "control_name",
				Description: "The name of the control.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.EvidenceFolder.ControlName")},
			{
				Name:        "data_source",
				Description: "The AWS service from which the evidence was collected.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.EvidenceFolder.DataSource")},
			{
				Name:        "date",
				Description: "The date when the first evidence was added to the evidence folder.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.EvidenceFolder.Date")},
			{
				Name:        "evidence_aws_service_source_count",
				Description: "The total number of AWS resources assessed to generate the evidence.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.EvidenceFolder.EvidenceAwsServiceSourceCount")},
			{
				Name:        "evidence_by_type_compliance_check_count",
				Description: "The number of evidence that falls under the compliance check category.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.EvidenceFolder.EvidenceByTypeComplianceCheckCount")},
			{
				Name:        "evidence_by_type_compliance_check_issues_count",
				Description: "The total number of issues that were reported directly from AWS Security Hub, AWS Config, or both.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.EvidenceFolder.EvidenceByTypeComplianceCheckIssuesCount")},
			{
				Name:        "evidence_by_type_configuration_data_count",
				Description: "The number of evidence that falls under the configuration data category.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.EvidenceFolder.EvidenceByTypeConfigurationDataCount")},
			{
				Name:        "evidence_by_type_manual_count",
				Description: "The number of evidence that falls under the manual category.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.EvidenceFolder.EvidenceByTypeManualCount")},
			{
				Name:        "evidence_by_type_user_activity_count",
				Description: "The number of evidence that falls under the user activity category.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.EvidenceFolder.EvidenceByTypeUserActivityCount")},
			{
				Name:        "evidence_resources_included_count",
				Description: "The amount of evidence included in the evidence folder.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.EvidenceFolder.EvidenceResourcesIncludedCount")},
			{
				Name:        "total_evidence",
				Description: "The total amount of evidence in the evidence folder.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.EvidenceFolder.TotalEvidence")},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.EvidenceFolder.Name")},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("ARN").Transform(transform.EnsureStringArray),
			},
		}),
	}
}
