package aws

import (
	"context"

	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsGuardDutyFinding(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_guardduty_finding",
		Description: "AWS GuardDuty Finding",
		List: &plugin.ListConfig{
			ParentHydrate: opengovernance.ListGuardDutyFinding,
			Hydrate:       opengovernance.ListGuardDutyFinding,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "detector_id", Require: plugin.Optional},
				{Name: "id", Require: plugin.Optional, Operators: []string{"=", "<>"}},
				{Name: "type", Require: plugin.Optional, Operators: []string{"=", "<>"}},
			},
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The title of the finding.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Finding.Title")},
			{
				Name:        "id",
				Description: "The ID of the finding.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Finding.Id")},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) specifying the finding.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Finding.Arn")},
			{
				Name:        "detector_id",
				Description: "The ID of the detector.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Finding.Service.DetectorId")},
			{
				Name:        "severity",
				Description: "The severity of the finding.",
				Type:        proto.ColumnType_DOUBLE,
				Transform:   transform.FromField("Description.Finding.Severity")},
			{
				Name:        "created_at",
				Description: "The time and date when the finding was created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Finding.CreatedAt")},
			{
				Name:        "confidence",
				Description: "The confidence score for the finding.",
				Type:        proto.ColumnType_DOUBLE,
				Transform:   transform.FromField("Description.Finding.Confidence")},
			{
				Name:        "description",
				Description: "The description of the finding.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Finding.Description")},
			{
				Name:        "schema_version",
				Description: "The version of the schema used for the finding.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Finding.SchemaVersion")},
			{
				Name:        "type",
				Description: "The type of finding.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Finding.Type")},
			{
				Name:        "updated_at",
				Description: "The time and date when the finding was last updated.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Finding.UpdatedAt")},
			{
				Name:        "resource",
				Description: "Contains information about the AWS resource associated with the activity that prompted GuardDuty to generate a finding.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Finding.Resource")},
			{
				Name:        "service",
				Description: "Contains additional information about the generated finding.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Finding.Service")},
			// Steampipe standard columns

			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Finding.Title")},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Finding.Arn").Transform(arnToAkas),
			},
		}),
	}
}

//// LIST FUNCTION
