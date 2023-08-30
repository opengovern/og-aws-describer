package aws

import (
	"context"

	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsInspectorAssessmentTarget(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_inspector_assessment_target",
		Description: "AWS Inspector Assessment Target",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("arn"),
			Hydrate:    kaytu.GetInspectorAssessmentTarget,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListInspectorAssessmentTarget,
		},

		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the Amazon Inspector assessment target.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AssessmentTarget.Name")},
			{
				Name:        "arn",
				Description: "The ARN that specifies the Amazon Inspector assessment target.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AssessmentTarget.Arn")},
			{
				Name:        "resource_group_arn",
				Description: "The ARN that specifies the resource group that is associated with the assessment target.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AssessmentTarget.ResourceGroupArn")},
			{
				Name:        "created_at",
				Description: "The time at which the assessment target is created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.AssessmentTarget.CreatedAt")},
			{
				Name:        "updated_at",
				Description: "The time at which UpdateAssessmentTarget is called.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.AssessmentTarget.UpdatedAt")},

			// Standard columns for all tables
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AssessmentTarget.Name")},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.AssessmentTarget.Arn").Transform(arnToAkas),
			},
		}),
	}
}
