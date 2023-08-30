package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/inspector/types"

	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

type ExclusionInfo = struct {
	types.Exclusion
	AssessmentRunArn string
}

//// TABLE DEFINITION

func tableAwsInspectorExclusion(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_inspector_exclusion",
		Description: "AWS Inspector Exclusion",
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListInspectorExclusion,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "assessment_run_arn", Require: plugin.Optional},
			},
		},

		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "arn",
				Description: "The ARN that specifies the exclusion.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Exclusion.Arn")},
			{
				Name:        "assessment_run_arn",
				Description: "The ARN that specifies the assessment run, the exclusion belongs to.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AssessmentRunArn")},
			{
				Name:        "attributes",
				Description: "The system-defined attributes for the exclusion.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Exclusion.Attributes")},
			{
				Name:        "description",
				Description: "The description of the exclusion.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Exclusion.Description")},
			{
				Name:        "recommendation",
				Description: "The recommendation for the exclusion.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Exclusion.Recommendation")},
			{
				Name:        "scopes",
				Description: "The AWS resources for which the exclusion pertains.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Exclusion.Scopes")},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Exclusion.Title")},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Exclusion.Arn").Transform(transform.EnsureStringArray),
			},
		}),
	}
}
