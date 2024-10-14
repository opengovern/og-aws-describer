package aws

import (
	"context"
	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsRDSDBRecommendation(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_rds_db_recommendation",
		Description: "AWS RDS DB Recommendation",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListRDSDBRecommendation,
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "recommendation_id",
				Description: "The unique identifier for the recommendation.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DBRecommendation.RecommendationId"),
			},
			{
				Name:        "resource_arn",
				Description: "The Amazon Resource Name (ARN) of the RDS resource associated with the recommendation.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DBRecommendation.ResourceArn"),
			},
			{
				Name:        "category",
				Description: "The category of the recommendation.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DBRecommendation.Category"),
			},
			{
				Name:        "status",
				Description: "The current status of the recommendation.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DBRecommendation.Status"),
			},
			{
				Name:        "created_time",
				Description: "The time when the recommendation was created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.DBRecommendation.CreatedTime"),
			},
			{
				Name:        "updated_time",
				Description: "The time when the recommendation was last updated.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.DBRecommendation.UpdatedTime").Transform(transform.NullIfZeroValue),
			},
			{
				Name:        "description",
				Description: "A detailed description of the recommendation.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DBRecommendation.Description"),
			},
			{
				Name:        "detection",
				Description: "A short description of the issue identified for this recommendation.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DBRecommendation.Detection"),
			},
			{
				Name:        "impact",
				Description: "A short description that explains the possible impact of an issue.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DBRecommendation.Impact"),
			},
			{
				Name:        "reason",
				Description: "The reason why this recommendation was created.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DBRecommendation.Reason"),
			},
			{
				Name:        "recommendation",
				Description: "A short description of the recommendation to resolve an issue.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DBRecommendation.Recommendation"),
			},
			{
				Name:        "severity",
				Description: "The severity level of the recommendation.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DBRecommendation.Severity"),
			},
			{
				Name:        "source",
				Description: "The AWS service that generated the recommendations.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DBRecommendation.Source"),
			},
			{
				Name:        "type_detection",
				Description: "A short description of the recommendation type.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DBRecommendation.TypeDetection"),
			},
			{
				Name:        "type_id",
				Description: "A value that indicates the type of recommendation.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DBRecommendation.TypeId"),
			},
			{
				Name:        "type_recommendation",
				Description: "A short description that summarizes the recommendation to fix all the issues of the recommendation type.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DBRecommendation.TypeRecommendation"),
			},
			{
				Name:        "additional_info",
				Description: "Additional information about the recommendation.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DBRecommendation.AdditionalInfo"),
			},

			// JSON fields
			{
				Name:        "issue_details",
				Description: "Details of the issue that caused the recommendation.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DBRecommendation.IssueDetails"),
			},
			{
				Name:        "links",
				Description: "A link to documentation that provides additional information about the recommendation.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DBRecommendation.Links"),
			},
			{
				Name:        "recommended_actions",
				Description: "A list of recommended actions.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DBRecommendation.RecommendedActions"),
			},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DBRecommendation.RecommendationId"),
			},
		}),
	}
}
