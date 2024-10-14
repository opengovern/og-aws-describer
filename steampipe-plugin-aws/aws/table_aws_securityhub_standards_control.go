package aws

import (
	"context"

	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsSecurityHubStandardsControl(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_securityhub_standards_control",
		Description: "AWS Security Hub Standards Control",
		List: &plugin.ListConfig{
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"ResourceNotFoundException"}),
			},
			Hydrate: opengovernance.ListSecurityHubStandardsControl,
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "control_id",
				Description: "The identifier of the security standard control.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.StandardsControl.ControlId")},
			{
				Name:        "arn",
				Description: "The ARN of the security standard control.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.StandardsControl.StandardsControlArn"),
			},
			{
				Name:        "control_status",
				Description: "The current status of the security standard control. Indicates whether the control is enabled or disabled. Security Hub does not check against disabled controls.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.StandardsControl.ControlStatus")},
			{
				Name:        "severity_rating",
				Description: "The severity of findings generated from this security standard control.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.StandardsControl.SeverityRating")},
			{
				Name:        "control_status_updated_at",
				Description: "The date and time that the status of the security standard control was most recently updated.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.StandardsControl.ControlStatusUpdatedAt")},
			{
				Name:        "description",
				Description: "The longer description of the security standard control.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.StandardsControl.Description")},
			{
				Name:        "disabled_reason",
				Description: "The reason provided for the most recent change in status for the control.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.StandardsControl.DisabledReason")},
			{
				Name:        "remediation_url",
				Description: "A link to remediation information for the control in the Security Hub user documentation.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.StandardsControl.RemediationUrl")},

			// JSON columns
			{
				Name:        "related_requirements",
				Description: "The list of requirements that are related to this control.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.StandardsControl.RelatedRequirements")},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.StandardsControl.Title")},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.StandardsControl.StandardsControlArn").Transform(transform.EnsureStringArray),
			},
		}),
	}
}
