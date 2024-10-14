package aws

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/inspector/types"

	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsInspectorAssessmentTemplate(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_inspector_assessment_template",
		Description: "AWS Inspector Assessment Template",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("arn"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{}),
			},
			Hydrate: opengovernance.GetInspectorAssessmentTemplate,
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListInspectorAssessmentTemplate,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "name", Require: plugin.Optional},
				{Name: "assessment_target_arn", Require: plugin.Optional},
			},
		},

		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the assessment template.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AssessmentTemplate.Name")},
			{
				Name:        "arn",
				Description: "The ARN of the assessment template.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AssessmentTemplate.Arn")},
			{
				Name:        "assessment_run_count",
				Description: "The number of existing assessment runs associated with this assessment template.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.AssessmentTemplate.AssessmentRunCount")},
			{
				Name:        "assessment_target_arn",
				Description: "The ARN of the assessment target that corresponds to this assessment template.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AssessmentTemplate.AssessmentTargetArn")},
			{
				Name:        "created_at",
				Description: "The time at which the assessment template is created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.AssessmentTemplate.CreatedAt")},
			{
				Name:        "duration_in_seconds",
				Description: "The duration in seconds specified for this assessment template.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.AssessmentTemplate.DurationInSeconds")},
			{
				Name:        "last_assessment_run_arn",
				Description: "The Amazon Resource Name (ARN) of the most recent assessment run associated with this assessment template.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AssessmentTemplate.LastAssessmentRunArn")},
			{
				Name:        "rules_package_arns",
				Description: "The rules packages that are specified for this assessment template.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.AssessmentTemplate.RulesPackageArns")},
			{
				Name:        "user_attributes_for_findings",
				Description: "The user-defined attributes that are assigned to every generated finding from the assessment run that uses this assessment template.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.AssessmentTemplate.UserAttributesForFindings")},
			{
				Name:        "tags_src",
				Description: "A list of tags associated with the Assessment Template.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Tags")},
			{
				Name:        "event_subscriptions",
				Description: "A list of event subscriptions associated with the Assessment Template.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.EventSubscriptions")},
			// Standard columns for all tables
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AssessmentTemplate.Name")},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Tags").Transform(inspectorTagListToTurbotTags),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.AssessmentTemplate.Arn").Transform(arnToAkas),
			},
		}),
	}
}

func inspectorTagListToTurbotTags(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	tagList := d.Value.([]types.Tag)

	if tagList == nil {
		return nil, nil
	}

	// Mapping the resource tags inside turbotTags
	var turbotTagsMap map[string]string
	if tagList != nil {
		turbotTagsMap = map[string]string{}
		for _, i := range tagList {
			turbotTagsMap[*i.Key] = *i.Value
		}
	}

	return turbotTagsMap, nil
}
