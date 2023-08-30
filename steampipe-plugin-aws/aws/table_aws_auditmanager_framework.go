package aws

import (
	"context"
	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsAuditManagerFramework(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_auditmanager_framework",
		Description: "AWS Audit Manager Framework",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"id", "region"}),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"ResourceNotFoundException", "ValidationException", "InternalServerException"}),
			},
			Hydrate: kaytu.GetAuditManagerFramework,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListAuditManagerFramework,
		},

		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the specified framework.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Framework.Name")},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the framework.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Framework.Arn")},
			{
				Name:        "id",
				Description: "The unique identified for the specified framework.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Framework.Id")},
			{
				Name:        "type",
				Description: "The framework type, such as standard or custom.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Framework.Type")},
			{
				Name:        "created_at",
				Description: "Specifies when the framework was created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Framework.CreatedAt")},
			{
				Name:        "created_by",
				Description: "The IAM user or role that created the framework.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Framework.CreatedBy")},
			{
				Name:        "compliance_type",
				Description: "The compliance type that the new custom framework supports, such as CIS or HIPAA.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Framework.ComplianceType")},
			{
				Name:        "controls_count",
				Description: "The number of controls associated with the specified framework.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Framework.ControlSources")},
			{
				Name:        "control_sets_count",
				Description: "The number of control sets associated with the specified framework.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Framework.ControlSources")},
			{
				Name:        "control_sources",
				Description: "The sources from which AWS Audit Manager collects evidence for the control.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Framework.ControlSources")},
			{
				Name:        "description",
				Description: "The description of the specified framework.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Framework.Description")},
			{
				Name:        "last_updated_at",
				Description: "Specifies when the framework was most recently updated.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Framework.LastUpdatedAt")},
			{
				Name:        "last_updated_by",
				Description: "The IAM user or role that most recently updated the framework.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Framework.LastUpdatedBy")},
			{
				Name:        "logo",
				Description: "The logo associated with the framework.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Framework.Logo")},
			{
				Name:        "control_sets",
				Description: "The control sets associated with the framework.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Framework.ControlSets")},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Framework.Name")},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Framework.Tags")},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Framework.Arn").Transform(arnToAkas),
			},
		}),
	}
}
