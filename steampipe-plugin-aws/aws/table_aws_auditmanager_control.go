package aws

import (
	"context"

	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsAuditManagerControl(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_auditmanager_control",
		Description: "AWS Audit Manager Control",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    kaytu.GetAuditManagerControl,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"ResourceNotFoundException", "ValidationException", "InvalidParameter"}),
			},
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListAuditManagerControl,
		},

		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the specified control.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Control.Name")},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the specified control.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Control.Arn")},
			{
				Name:        "id",
				Description: "An unique identifier for the specified control.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Control.Id")},
			{
				Name:        "type",
				Description: "The type of control, such as custom or standard.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Control.Type")},
			{
				Name:        "created_at",
				Description: "Specifies when the control was created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Control.CreatedAt")},
			{
				Name:        "created_by",
				Description: "The IAM user or role that created the control.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Control.CreatedBy")},
			{
				Name:        "action_plan_title",
				Description: "The title of the action plan for remediating the control.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Control.ActionPlanTitle")},
			{
				Name:        "action_plan_instructions",
				Description: "The recommended actions to carry out if the control is not fulfilled.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Control.ActionPlanInstructions")},
			{
				Name:        "control_sources",
				Description: "The data source that determines from where AWS Audit Manager collects evidence for the control.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Control.ControlSources")},
			{
				Name:        "description",
				Description: "The description of the specified control.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Control.Description")},
			{
				Name:        "last_updated_at",
				Description: "Specifies when the control was most recently updated.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Control.LastUpdatedAt")},
			{
				Name:        "last_updated_by",
				Description: "The IAM user or role that most recently updated the control.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Control.LastUpdatedBy")},
			{
				Name:        "testing_information",
				Description: "The steps to follow to determine if the control has been satisfied.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Control.TestingInformation")},
			{
				Name:        "control_mapping_sources",
				Description: "The data mapping sources for the specified control.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Control.ControlMappingSources")},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Control.Name")},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Control.Tags")},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Control.Arn").Transform(arnToAkas),
			},
		}),
	}
}
