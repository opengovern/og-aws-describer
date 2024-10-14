package aws

import (
	"context"
	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsSSMManagedInstanceCompliance(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_ssm_managed_instance_compliance",
		Description: "AWS SSM Managed Instance Compliance",
		List: &plugin.ListConfig{
			//KeyColumns: plugin.SingleColumn("resource_id"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"InvalidResourceId", "ValidationException"}),
			},
			Hydrate: opengovernance.ListSSMManagedInstanceCompliance,
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "id",
				Description: "An ID for the compliance item.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ComplianceItem.Id")},
			{
				Name:        "name",
				Description: "A title for the compliance item.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ComplianceItem.Title")},
			{
				Name:        "resource_id",
				Description: "An ID for the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ComplianceItem.ResourceId")},
			{
				Name:        "status",
				Description: "The status of the compliance item.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ComplianceItem.Status")},
			{
				Name:        "compliance_type",
				Description: "The compliance type.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ComplianceItem.ComplianceType")},
			{
				Name:        "resource_type",
				Description: "The type of resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ComplianceItem.ResourceType")},
			{
				Name:        "severity",
				Description: "The severity of the compliance status.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ComplianceItem.Severity")},
			{
				Name:        "details",
				Description: "A key-value combination details for the compliance item.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ComplianceItem.Details")},
			{
				Name:        "execution_summary",
				Description: "A summary for the compliance item.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ComplianceItem.ExecutionSummary")},
			// Steampipe standard columns

			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.From(ssmManagedInstanceComplianceTitle),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Hydrate:     getSSMManagedInstanceComplianceAkas,
				Transform:   transform.FromValue(),
			},
		}),
	}
}

//// LIST FUNCTION

//// HYDRATE FUNCTIONS

func getSSMManagedInstanceComplianceAkas(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("getSSMInstanceComplianceAkas")
	data := h.Item.(opengovernance.SSMManagedInstanceCompliance).Description.ComplianceItem
	metadata := h.Item.(opengovernance.SSMManagedInstanceCompliance).Metadata
	region := d.EqualsQualString(matrixKeyRegion)

	akas := []string{"arn:" + metadata.Partition + ":ssm:" + region + ":" + metadata.AccountID + ":managed-instance/" + *data.ResourceId + "/compliance-item/" + *data.Id + ":" + *data.ComplianceType}

	return akas, nil
}

//// TRANSFORM FUNCTIONS

func ssmManagedInstanceComplianceTitle(_ context.Context, d *transform.TransformData) (interface{}, error) {
	data := d.HydrateItem.(opengovernance.SSMManagedInstanceCompliance).Description.ComplianceItem
	title := ""
	if data.Id != nil {
		title = *data.Id
	}
	if data.Title != nil && len(*data.Title) > 0 {
		title = *data.Title
	}

	return title, nil
}
