package aws

import (
	"context"
	"strings"

	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsSecurityHubFinding(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_securityhub_finding",
		Description: "AWS Security Hub Finding",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    kaytu.GetSecurityHubFinding,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"InvalidAccessException"}),
			},
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListSecurityHubFinding,
			KeyColumns: plugin.KeyColumnSlice{
				{Name: "company_name", Require: plugin.Optional, Operators: []string{"=", "<>"}},
				{Name: "compliance_status", Require: plugin.Optional, Operators: []string{"=", "<>"}},
				{Name: "confidence", Require: plugin.Optional, Operators: []string{"=", ">=", "<="}},
				{Name: "criticality", Require: plugin.Optional, Operators: []string{"=", ">=", "<="}},
				{Name: "generator_id", Require: plugin.Optional, Operators: []string{"=", "<>"}},
				{Name: "product_arn", Require: plugin.Optional, Operators: []string{"=", "<>"}},
				{Name: "product_name", Require: plugin.Optional, Operators: []string{"=", "<>"}},
				{Name: "record_state", Require: plugin.Optional, Operators: []string{"=", "<>"}},
				{Name: "title", Require: plugin.Optional, Operators: []string{"=", "<>"}},
				{Name: "verification_state", Require: plugin.Optional, Operators: []string{"=", "<>"}},
				{Name: "workflow_state", Require: plugin.Optional, Operators: []string{"=", "<>"}},
				{Name: "workflow_status", Require: plugin.Optional, Operators: []string{"=", "<>"}},
				{Name: "source_account_id", Require: plugin.Optional, Operators: []string{"=", "<>"}},
			},
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"InvalidAccessException"}),
			},
		},

		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "id",
				Description: "The security findings provider-specific identifier for a finding.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Finding.Id")},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) for the finding.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Finding.Id"),
			},
			{
				Name:        "company_name",
				Description: "The name of the company for the product that generated the finding.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Finding.CompanyName")},
			{
				Name:        "confidence",
				Description: "A finding's confidence. Confidence is defined as the likelihood that a finding accurately identifies the behavior or issue that it was intended to identify.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Finding.Confidence")},
			{
				Name:        "created_at",
				Description: "Indicates when the security-findings provider created the potential security issue that a finding captured.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Finding.CreatedAt")},
			{
				Name:        "compliance_status",
				Description: "The result of a compliance standards check.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Finding.Compliance.Status")},
			{
				Name:        "updated_at",
				Description: "Indicates when the security-findings provider last updated the finding record.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Finding.UpdatedAt")},
			{
				Name:        "criticality",
				Description: "The level of importance assigned to the resources associated with the finding.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Finding.Criticality")},
			{
				Name:        "description",
				Description: "A finding's description.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Finding.Description")},
			{
				Name:        "first_observed_at",
				Description: "Indicates when the security-findings provider first observed the potential security issue that a finding captured.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Finding.FirstObservedAt")},
			{
				Name:        "generator_id",
				Description: "The identifier for the solution-specific component (a discrete unit of logic) that generated a finding.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Finding.GeneratorId")},
			{
				Name:        "last_observed_at",
				Description: "Indicates when the security-findings provider most recently observed the potential security issue that a finding captured.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Finding.LastObservedAt")},
			{
				Name:        "product_arn",
				Description: "The ARN generated by Security Hub that uniquely identifies a product that generates findings.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Finding.ProductArn")},
			{
				Name:        "product_name",
				Description: "The name of the product that generated the finding.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Finding.ProductName")},
			{
				Name:        "record_state",
				Description: "The record state of a finding.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Finding.RecordState")},
			{
				Name:        "schema_version",
				Description: "The schema version that a finding is formatted for.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Finding.SchemaVersion")},
			{
				Name:        "source_url",
				Description: "A URL that links to a page about the current finding in the security-findings provider's solution.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Finding.SourceUrl")},
			{
				Name:        "verification_state",
				Description: "Indicates the veracity of a finding.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Finding.VerificationState")},
			{
				Name:        "workflow_state",
				Description: "[DEPRECATED] This column has been deprecated and will be removed in a future release. The workflow state of a finding.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Finding.WorkflowState")},
			{
				Name:        "workflow_status",
				Description: "The workflow status of a finding. Possible values are NEW, NOTIFIED, SUPPRESSED, RESOLVED.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Finding.Workflow.Status")},
			{
				Name:        "standards_control_arn",
				Description: "The ARN of the security standard control.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.From(extractStandardControlArn),
			},
			{
				Name:        "action",
				Description: "Provides details about an action that affects or that was taken on a resource.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Finding.Action")},
			{
				Name:        "compliance",
				Description: "This data type is exclusive to findings that are generated as the result of a check run against a specific rule in a supported security standard, such as CIS Amazon Web Services Foundations.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Finding.Compliance")},
			{
				Name:        "finding_provider_fields",
				Description: "In a BatchImportFindings request, finding providers use FindingProviderFields to provide and update their own values for confidence, criticality, related findings, severity, and types.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Finding.FindingProviderFields")},
			{
				Name:        "malware",
				Description: "A list of malware related to a finding.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Finding.Malware")},
			{
				Name:        "network",
				Description: "The details of network-related information about a finding.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Finding.Network")},
			{
				Name:        "network_path",
				Description: "Provides information about a network path that is relevant to a finding. Each entry under NetworkPath represents a component of that path.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Finding.NetworkPath")},
			{
				Name:        "note",
				Description: "A user-defined note added to a finding.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Finding.Note")},
			{
				Name:        "patch_summary",
				Description: "Provides an overview of the patch compliance status for an instance against a selected compliance standard.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Finding.PatchSummary")},
			{
				Name:        "process",
				Description: "The details of process-related information about a finding.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Finding.Process")},
			{
				Name:        "product_fields",
				Description: "A data type where security-findings providers can include additional solution-specific details that aren't part of the defined AwsSecurityFinding format.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Finding.ProductFields")},
			{
				Name:        "related_findings",
				Description: "A list of related findings.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Finding.RelatedFindings")},
			{
				Name:        "remediation",
				Description: "A data type that describes the remediation options for a finding.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Finding.Remediation")},
			{
				Name:        "resources",
				Description: "A set of resource data types that describe the resources that the finding refers to.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Finding.Resources")},
			{
				Name:        "severity",
				Description: "A finding's severity.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Finding.Severity")},
			{
				Name:        "threat_intel_indicators",
				Description: "Threat intelligence details related to a finding.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Finding.ThreatIntelIndicators")},
			{
				Name:        "user_defined_fields",
				Description: "A list of name/value string pairs associated with the finding.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Finding.UserDefinedFields")},
			{
				Name:        "vulnerabilities",
				Description: "Provides a list of vulnerabilities associated with the findings.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Finding.Vulnerabilities")},
			{
				Name:        "source_account_id",
				Description: "The account id where the affected resource lives.",
				Type:        proto.ColumnType_STRING,
			},
			/// Steampipe standard columns
			{
				Name:        "title",
				Description: "A finding's title.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Finding.Title")},
		}),
	}
}

//// TRANSFORM FUNCTIONS

func extractStandardControlArn(_ context.Context, d *transform.TransformData) (interface{}, error) {
	findingArn := d.HydrateItem.(kaytu.SecurityHubFinding).Description.Finding.Id

	if strings.Contains(*findingArn, "arn:aws:securityhub") {
		standardControlArn := strings.Replace(strings.Split(*findingArn, "/finding")[0], "subscription", "control", 1)
		return standardControlArn, nil
	}
	return nil, nil
}
