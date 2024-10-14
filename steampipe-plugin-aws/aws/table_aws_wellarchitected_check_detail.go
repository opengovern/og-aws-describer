package aws

import (
	"context"
	"errors"
	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/wellarchitected"
	"github.com/aws/aws-sdk-go-v2/service/wellarchitected/types"
	"github.com/aws/smithy-go"

	"github.com/turbot/go-kit/helpers"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsWellArchitectedCheckDetail(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_wellarchitected_check_detail",
		Description: "AWS Well-Architected Check Detail",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListWellArchitectedCheckDetail,
			// IgnoreConfig: &plugin.IgnoreConfig{
			// 	ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"ResourceNotFoundException", "ValidationException"}),
			// },
			KeyColumns: []*plugin.KeyColumn{
				{Name: "workload_id", Require: plugin.Optional},
				{Name: "choice_id", Require: plugin.Optional},
				{Name: "lens_arn", Require: plugin.Optional},
				{Name: "pillar_id", Require: plugin.Optional},
				{Name: "question_id", Require: plugin.Optional},
			},
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "id",
				Description: "Trusted Advisor check ID.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.CheckDetail.Id"),
			},
			{
				Name:        "choice_id",
				Description: "The ID of a choice.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.CheckDetail.ChoiceId"),
			},
			{
				Name:        "description",
				Description: "Trusted Advisor check description.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.CheckDetail.Description"),
			},
			{
				Name:        "flagged_resources",
				Description: "Count of flagged resources associated to the check.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.CheckDetail.FlaggedResources"),
			},
			{
				Name:        "lens_arn",
				Description: "Well-Architected Lens ARN associated to the check.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.CheckDetail.LensArn"),
			},
			{
				Name:        "name",
				Description: "Trusted Advisor check name.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.CheckDetail.Name"),
			},
			{
				Name:        "owner_account_id",
				Description: "An Amazon Web Services account ID.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.CheckDetail.AccountId"),
			},
			{
				Name:        "pillar_id",
				Description: "The ID used to identify a pillar, for example, security. A pillar is identified by its PillarReviewSummary$PillarId.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.CheckDetail.PillarId"),
			},
			{
				Name:        "provider",
				Description: "Provider of the check related to the best practice.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.CheckDetail.Provider"),
			},
			{
				Name:        "question_id",
				Description: "The ID of the question.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.CheckDetail.QuestionId"),
			},
			{
				Name:        "reason",
				Description: "Reason associated to the check.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.CheckDetail.Reason"),
			},
			{
				Name:        "status",
				Description: "Status associated to the check.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.CheckDetail.Status"),
			},
			{
				Name:        "updated_at",
				Description: "The date and time recorded.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.CheckDetail.UpdatedAt"),
			},
			{
				Name:        "workload_id",
				Description: "The ID of the workload.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.WorkloadId"),
			},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.CheckDetail.Name"),
			},
		}),
	}
}

type CheckDetailInfo struct {
	types.CheckDetail
	WorkloadId *string
}

//// LIST FUNCTION

func listWellArchitectedCheckDetails(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {

	answerList, err := getAnswerDetailsForWorkload(ctx, d, h)
	if err != nil {
		plugin.Logger(ctx).Error("aws_wellarchitected_check_detail.getAnswerDetailsForWorkload", "error", err)
		return nil, err
	}

	if len(answerList) == 0 {
		return nil, nil
	}

	// Create session
	svc, err := WellArchitectedClient(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("aws_wellarchitected_check_detail.listWellArchitectedCheckDetails", "client_error", err)
		return nil, err
	}

	// Unsupported region, return no data
	if svc == nil {
		return nil, nil
	}

	for _, answer := range answerList {
		// List call
		// Context can be cancelled due to manual cancellation or the limit has been hit
		if d.EqualsQualString("pillar_id") != "" && d.EqualsQualString("pillar_id") != *answer.PillarId {
			continue
		}
		if d.EqualsQualString("question_id") != "" && d.EqualsQualString("question_id") != *answer.QuestionId {
			continue
		}
		_, err := fetchWellArchitectedCheckDetails(ctx, d, svc, answer)
		if err != nil {
			plugin.Logger(ctx).Error("aws_wellarchitected_check_detail.listWellArchitectedCheckDetails", "api_error", err)
			return nil, err
		}
	}

	return nil, nil
}

func fetchWellArchitectedCheckDetails(ctx context.Context, d *plugin.QueryData, svc *wellarchitected.Client, answer AnswerInfo) (interface{}, error) {

	// Limiting the results
	maxLimit := int32(50)
	if d.QueryContext.Limit != nil {
		limit := int32(*d.QueryContext.Limit)
		if limit < maxLimit {
			maxLimit = limit
		}
	}

	for _, choice := range answer.Choices {
		if d.EqualsQualString("choice_id") != "" && d.EqualsQualString("choice_id") != *choice.ChoiceId {
			continue
		}

		input := &wellarchitected.ListCheckDetailsInput{
			MaxResults: aws.Int32(maxLimit),
			LensArn:    aws.String(*answer.LensArn),
			PillarId:   aws.String(*answer.PillarId),
			QuestionId: aws.String(*answer.QuestionId),
			WorkloadId: aws.String(*answer.WorkloadId),
			ChoiceId:   aws.String(*choice.ChoiceId),
		}

		paginator := wellarchitected.NewListCheckDetailsPaginator(svc, input, func(o *wellarchitected.ListCheckDetailsPaginatorOptions) {
			o.Limit = maxLimit
			o.StopOnDuplicateToken = true
		})

		for paginator.HasMorePages() {
			output, err := paginator.NextPage(ctx)
			if err != nil {
				var ae smithy.APIError
				if errors.As(err, &ae) {
					if helpers.StringSliceContains([]string{"ResourceNotFoundException"}, ae.ErrorCode()) {
						return nil, nil
					}
				}

				return nil, err
			}

			for _, item := range output.CheckDetails {
				d.StreamListItem(ctx, CheckDetailInfo{item, answer.WorkloadId})

				if d.RowsRemaining(ctx) == 0 {
					return nil, nil
				}
			}
		}
	}
	return nil, nil
}

func getAnswerDetailsForWorkload(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) ([]AnswerInfo, error) {
	workload := h.Item.(types.WorkloadSummary)
	workloadId := workload.WorkloadId

	// Validate - User inputs must not be blank and return nil if doesn't match the hydrated workload id
	if d.EqualsQualString("workload_id") != "" && d.EqualsQualString("workload_id") != *workloadId {
		return nil, nil
	}

	var answerList []AnswerInfo

	// Create session
	svc, err := WellArchitectedClient(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("aws_wellarchitected_check_detail.getAnswerDetailsForWorkload", "client_error", err)
		return nil, err
	}

	// Unsupported region, return no data
	if svc == nil {
		return nil, nil
	}

	for _, lensAlias := range workload.Lenses {
		lensArn := lensAlias
		if d.EqualsQualString("lens_arn") != "" && d.EqualsQualString("lens_arn") != lensArn {
			continue
		}

		input := &wellarchitected.ListAnswersInput{
			MaxResults: aws.Int32(50),
			LensAlias:  aws.String(lensArn),
			WorkloadId: aws.String(*workloadId),
		}

		paginator := wellarchitected.NewListAnswersPaginator(svc, input, func(o *wellarchitected.ListAnswersPaginatorOptions) {
			o.Limit = int32(50)
			o.StopOnDuplicateToken = true
		})

		// List call
		for paginator.HasMorePages() {
			output, err := paginator.NextPage(ctx)
			if err != nil {
				var ae smithy.APIError
				if errors.As(err, &ae) {
					if helpers.StringSliceContains([]string{"ResourceNotFoundException"}, ae.ErrorCode()) {
						return nil, nil
					}
				}

				plugin.Logger(ctx).Error("aws_wellarchitected_check_detail.getAnswerDetailsForWorkload", "api_error", err)
				return nil, err
			}

			for _, item := range output.AnswerSummaries {

				answer := types.Answer{
					Choices:       item.Choices,
					IsApplicable:  item.IsApplicable,
					PillarId:      item.PillarId,
					QuestionId:    item.QuestionId,
					QuestionTitle: item.QuestionTitle,
					Reason:        item.Reason,
					Risk:          item.Risk,
				}

				answerList = append(answerList, AnswerInfo{answer, output.LensAlias, output.LensArn, output.MilestoneNumber, output.WorkloadId})
			}
		}
	}
	return answerList, nil
}
