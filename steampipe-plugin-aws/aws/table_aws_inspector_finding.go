package aws

import (
	"context"

	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsInspectorFinding(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_inspector_finding",
		Description: "AWS Inspector Finding",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("arn"),
			Hydrate:    opengovernance.GetInspectorFinding,
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListInspectorFinding,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"InvalidInputException", "NoSuchEntity", "InvalidParameter"}),
			},
			KeyColumns: plugin.KeyColumnSlice{
				{Name: "agent_id", Require: plugin.Optional, Operators: []string{"="}},
				{Name: "auto_scaling_group", Require: plugin.Optional, Operators: []string{"="}},
				{Name: "severity", Require: plugin.Optional, Operators: []string{"="}},
			},
		},

		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "id",
				Description: "The ID of the finding.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Finding.Id"),
			},
			{
				Name:        "arn",
				Description: "The ARN that specifies the finding.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Finding.Arn"),
			},
			{
				Name:        "agent_id",
				Description: "The ID of the agent that is installed on the EC2 instance where the finding is generated.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Finding.AssetAttributes.AgentId")},
			{
				Name:        "asset_type",
				Description: "The type of the host from which the finding is generated.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Finding.AssetType")},
			{
				Name:        "auto_scaling_group",
				Description: "The Auto Scaling group of the EC2 instance where the finding is generated.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Finding.AssetAttributes.AutoScalingGroup")},
			{
				Name:        "confidence",
				Description: "This data element is currently not used.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Finding.Confidence")},
			{
				Name:        "created_at",
				Description: "The time when the finding was generated.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Finding.CreatedAt")},
			{
				Name:        "updated_at",
				Description: "The time when AddAttributesToFindings is called.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Finding.UpdatedAt")},
			{
				Name:        "description",
				Description: "The description of the finding.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Finding.Description")},
			{
				Name:        "indicator_of_compromise",
				Description: "This data element is currently not used.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Finding.IndicatorOfCompromise")},
			{
				Name:        "numeric_severity",
				Description: "The numeric value of the finding severity.",
				Type:        proto.ColumnType_DOUBLE,
				Transform:   transform.FromField("Description.Finding.NumericSeverity")},
			{
				Name:        "recommendation",
				Description: "The recommendation for the finding.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Finding.Recommendation")},
			{
				Name:        "schema_version",
				Description: "The schema version of this data type.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Finding.SchemaVersion")},
			{
				Name:        "service",
				Description: "The data element is set to 'Inspector'.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Finding.Service")},
			{
				Name:        "severity",
				Description: "The finding severity. Values can be set to High, Medium, Low, and Informational.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Finding.Severity")},

			// Json columns
			{
				Name:        "asset_attributes",
				Description: "A collection of attributes of the host from which the finding is generated.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Finding.AssetAttributes")},
			{
				Name:        "attributes",
				Description: "The system-defined attributes for the finding.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Finding.Attributes")},
			{
				Name:        "failed_items",
				Description: "Attributes details that cannot be described. An error code is provided for each failed item.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.FailedItems")},
			{
				Name:        "service_attributes",
				Description: "This data type is used in the Finding data type.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Finding.ServiceAttributes")},
			{
				Name:        "user_attributes",
				Description: "The user-defined attributes that are assigned to the finding.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Finding.UserAttributes")},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: "The name of the finding.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Finding.Title")},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Finding.Arn").Transform(transform.EnsureStringArray),
			},
		}),
	}
}
