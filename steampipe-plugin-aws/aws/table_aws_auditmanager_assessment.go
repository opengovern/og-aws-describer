package aws

import (
	"context"

	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsAuditManagerAssessment(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_auditmanager_assessment",
		Description: "AWS Audit Manager Assessment",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"ResourceNotFoundException", "ValidationException", "InvalidParameter"}),
			},
			Hydrate: opengovernance.GetAuditManagerAssessment,
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListAuditManagerAssessment,
		},

		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the assessment.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Assessment.Metadata.Name")},
			{
				Name:        "id",
				Description: "An unique identifier for the assessment.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Assessment.Metadata.Id"),
			},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the assessment.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("ARN")},
			{
				Name:        "status",
				Description: "The current status of the assessment.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Assessment.Metadata.Status")},
			{
				Name:        "compliance_type",
				Description: "The name of the compliance standard related to the assessment.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Assessment.Metadata.ComplianceType")},
			{
				Name:        "assessment_report_destination",
				Description: "The destination of the assessment report.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Assessment.Metadata.AssessmentReportsDestination.Destination")},
			{
				Name:        "assessment_report_destination_type",
				Description: "The destination type, such as Amazon S3.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Assessment.Metadata.AssessmentReportsDestination.DestinationType")},
			{
				Name:        "creation_time",
				Description: "Specifies when the assessment was created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Assessment.Metadata.CreationTime")},
			{
				Name:        "description",
				Description: "The description of the assessment.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Assessment.Metadata.Description")},
			{
				Name:        "last_updated",
				Description: "The time of the most recent update.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Assessment.Metadata.LastUpdated")},
			{
				Name:        "aws_account",
				Description: "The AWS account associated with the assessment.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Assessment.AwsAccount")},
			{
				Name:        "delegations",
				Description: "The delegations associated with the assessment.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Assessment.Metadata.Delegations")},
			{
				Name:        "framework",
				Description: "The framework from which the assessment was created.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Assessment.Framework")},
			{
				Name:        "scope",
				Description: "The wrapper of AWS accounts and services in scope for the assessment.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Assessment.Metadata.Scope")},
			{
				Name:        "roles",
				Description: "The roles associated with the assessment.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Assessment.Metadata.Roles")},

			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Assessment.Metadata.Name")},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Assessment.Tags")},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("ARN").Transform(arnToAkas),
			},
		}),
	}
}
