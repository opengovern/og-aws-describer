package aws

import (
	"context"

	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsInspectorAssessmentRun(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_inspector_assessment_run",
		Description: "AWS Inspector Assessment Run",
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListInspectorAssessmentRun,
			KeyColumns: plugin.KeyColumnSlice{
				{Name: "assessment_template_arn", Require: plugin.Optional},
				{Name: "name", Require: plugin.Optional},
				{Name: "state", Require: plugin.Optional},
			},
		},

		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The auto-generated name for the assessment run.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AssessmentRun.Name")},
			{
				Name:        "arn",
				Description: "The ARN of the assessment run.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AssessmentRun.Arn")},
			{
				Name:        "assessment_template_arn",
				Description: "The ARN of the assessment template that is associated with the assessment run.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AssessmentRun.AssessmentTemplateArn")},
			{
				Name:        "state",
				Description: "The state of the assessment run.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AssessmentRun.State")},
			{
				Name:        "completed_at",
				Description: "The assessment run completion time that corresponds to the rules packages evaluation completion time or failure.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.AssessmentRun.CompletedAt")},
			{
				Name:        "created_at",
				Description: "The time when StartAssessmentRun was called.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.AssessmentRun.CreatedAt")},
			{
				Name:        "data_collected",
				Description: "Boolean value (true or false) that specifies whether the process of collecting data from the agents is completed.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.AssessmentRun.DataCollected")},
			{
				Name:        "duration_in_seconds",
				Description: "The duration of the assessment run.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.AssessmentRun.DurationInSeconds")},
			{
				Name:        "started_at",
				Description: "The time when StartAssessmentRun was called.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.AssessmentRun.StartedAt")},
			{
				Name:        "state_changed_at",
				Description: "The last time when the assessment run's state changed.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.AssessmentRun.StateChangedAt")},
			{
				Name:        "finding_counts",
				Description: "Provides a total count of generated findings per severity.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.AssessmentRun.FindingCounts")},
			{
				Name:        "notifications",
				Description: "A list of notifications for the event subscriptions.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.AssessmentRun.Notifications")},
			{
				Name:        "rules_package_arns",
				Description: "The rules packages selected for the assessment run.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.AssessmentRun.RulesPackageArns")},
			{
				Name:        "state_changes",
				Description: "A list of the assessment run state changes.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.AssessmentRun.StateChanges")},
			{
				Name:        "user_attributes_for_findings",
				Description: "The user-defined attributes that are assigned to every generated finding.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.AssessmentRun.UserAttributesForFindings")},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AssessmentRun.Name")},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.AssessmentRun.Arn").Transform(transform.EnsureStringArray),
			},
		}),
	}
}
