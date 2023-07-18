package describer

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/inspector/types"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/inspector"
	"github.com/kaytu-io/kaytu-aws-describer/aws/model"
)

func InspectorAssessmentRun(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := inspector.NewFromConfig(cfg)
	paginator := inspector.NewListAssessmentRunsPaginator(client, &inspector.ListAssessmentRunsInput{})

	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		if len(page.AssessmentRunArns) == 0 {
			continue
		}

		assessmentRuns, err := client.DescribeAssessmentRuns(ctx, &inspector.DescribeAssessmentRunsInput{
			AssessmentRunArns: page.AssessmentRunArns,
		})
		if err != nil {
			return nil, err
		}

		for _, assessmentRun := range assessmentRuns.AssessmentRuns {
			resource := Resource{
				Region: describeCtx.KaytuRegion,
				Name:   *assessmentRun.Name,
				ARN:    *assessmentRun.Arn,
				Description: model.InspectorAssessmentRunDescription{
					AssessmentRun: assessmentRun,
				},
			}
			if stream != nil {
				if err := (*stream)(resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, resource)
			}
		}
	}

	return values, nil
}

func InspectorAssessmentTarget(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	client := inspector.NewFromConfig(cfg)
	paginator := inspector.NewListAssessmentTargetsPaginator(client, &inspector.ListAssessmentTargetsInput{})

	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		if len(page.AssessmentTargetArns) == 0 {
			continue
		}

		assessmentTargets, err := client.DescribeAssessmentTargets(ctx, &inspector.DescribeAssessmentTargetsInput{
			AssessmentTargetArns: page.AssessmentTargetArns,
		})
		if err != nil {
			return nil, err
		}

		for _, assessmentTarget := range assessmentTargets.AssessmentTargets {
			resource := inspectorAssessmentTargetHandel(ctx, assessmentTarget)
			if stream != nil {
				if err := (*stream)(resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, resource)
			}
		}
	}
	return values, nil
}
func inspectorAssessmentTargetHandel(ctx context.Context, assessmentTarget types.AssessmentTarget) Resource {
	describeCtx := GetDescribeContext(ctx)
	resource := Resource{
		Region: describeCtx.KaytuRegion,
		Name:   *assessmentTarget.Name,
		ARN:    *assessmentTarget.Arn,
		Description: model.InspectorAssessmentTargetDescription{
			AssessmentTarget: assessmentTarget,
		},
	}
	return resource
}
func GetInspectorAssessmentTarget(ctx context.Context, cfg aws.Config, fields map[string]string) ([]Resource, error) {
	AssessmentTargetArn := fields["arn"]
	var values []Resource
	client := inspector.NewFromConfig(cfg)

	describeAssessments, err := client.DescribeAssessmentTargets(ctx, &inspector.DescribeAssessmentTargetsInput{
		AssessmentTargetArns: []string{AssessmentTargetArn},
	})
	if err != nil {
		if isErr(err, "DescribeAssessmentTargetsNotFound") || isErr(err, "InvalidParameterValue") {
			return nil, nil
		}
		return nil, err
	}

	for _, assessmentTarget := range describeAssessments.AssessmentTargets {
		resource := inspectorAssessmentTargetHandel(ctx, assessmentTarget)
		values = append(values, resource)
	}
	return values, nil
}

func InspectorAssessmentTemplate(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	client := inspector.NewFromConfig(cfg)
	paginator := inspector.NewListAssessmentTemplatesPaginator(client, &inspector.ListAssessmentTemplatesInput{})

	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		if len(page.AssessmentTemplateArns) == 0 {
			continue
		}

		assessmentTemplates, err := client.DescribeAssessmentTemplates(ctx, &inspector.DescribeAssessmentTemplatesInput{
			AssessmentTemplateArns: page.AssessmentTemplateArns,
		})
		if err != nil {
			return nil, err
		}

		for _, assessmentTemplate := range assessmentTemplates.AssessmentTemplates {
			tags, err := client.ListTagsForResource(ctx, &inspector.ListTagsForResourceInput{
				ResourceArn: assessmentTemplate.Arn,
			})
			if err != nil {
				return nil, err
			}

			eventSubscriptions, err := client.ListEventSubscriptions(ctx, &inspector.ListEventSubscriptionsInput{
				ResourceArn: assessmentTemplate.Arn,
			})
			if err != nil {
				return nil, err
			}

			resource := inspectorAssessmentTemplateHandel(ctx, assessmentTemplate, eventSubscriptions, tags)
			if stream != nil {
				if err := (*stream)(resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, resource)
			}
		}
	}

	return values, nil
}
func inspectorAssessmentTemplateHandel(ctx context.Context, assessmentTemplate types.AssessmentTemplate, eventSubscriptions *inspector.ListEventSubscriptionsOutput, tags *inspector.ListTagsForResourceOutput) Resource {
	describeCtx := GetDescribeContext(ctx)
	resource := Resource{
		Region: describeCtx.KaytuRegion,
		Name:   *assessmentTemplate.Name,
		ARN:    *assessmentTemplate.Arn,
		Description: model.InspectorAssessmentTemplateDescription{
			AssessmentTemplate: assessmentTemplate,
			EventSubscriptions: eventSubscriptions.Subscriptions,
			Tags:               tags.Tags,
		},
	}
	return resource
}
func GetInspectorAssessmentTemplate(ctx context.Context, cfg aws.Config, fields map[string]string) ([]Resource, error) {
	arn := fields["arn"]
	client := inspector.NewFromConfig(cfg)

	var values []Resource
	assessmentTemplates, err := client.DescribeAssessmentTemplates(ctx, &inspector.DescribeAssessmentTemplatesInput{
		AssessmentTemplateArns: []string{arn},
	})
	if err != nil {
		return nil, err
	}

	for _, assessmentTemplate := range assessmentTemplates.AssessmentTemplates {
		tags, err := client.ListTagsForResource(ctx, &inspector.ListTagsForResourceInput{
			ResourceArn: assessmentTemplate.Arn,
		})
		if err != nil {
			return nil, err
		}

		eventSubscriptions, err := client.ListEventSubscriptions(ctx, &inspector.ListEventSubscriptionsInput{
			ResourceArn: assessmentTemplate.Arn,
		})
		if err != nil {
			return nil, err
		}

		resource := inspectorAssessmentTemplateHandel(ctx, assessmentTemplate, eventSubscriptions, tags)
		values = append(values, resource)
	}
	return values, nil
}

func InspectorExclusion(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := inspector.NewFromConfig(cfg)
	paginator := inspector.NewListAssessmentRunsPaginator(client, &inspector.ListAssessmentRunsInput{})

	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, assessmentRun := range page.AssessmentRunArns {
			exclusionsPaginator := inspector.NewListExclusionsPaginator(client, &inspector.ListExclusionsInput{
				AssessmentRunArn: &assessmentRun,
			})

			for exclusionsPaginator.HasMorePages() {
				exclusionPage, err := exclusionsPaginator.NextPage(ctx)
				if err != nil {
					return nil, err
				}

				exclusions, err := client.DescribeExclusions(ctx, &inspector.DescribeExclusionsInput{
					ExclusionArns: exclusionPage.ExclusionArns,
				})
				if err != nil {
					return nil, err
				}

				for _, exclusion := range exclusions.Exclusions {
					resource := Resource{
						Region: describeCtx.KaytuRegion,
						Name:   *exclusion.Title,
						ARN:    *exclusion.Arn,
						Description: model.InspectorExclusionDescription{
							Exclusion: exclusion,
						},
					}
					if stream != nil {
						if err := (*stream)(resource); err != nil {
							return nil, err
						}
					} else {
						values = append(values, resource)
					}
				}
			}
		}
	}
	return values, nil
}

func InspectorFinding(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := inspector.NewFromConfig(cfg)
	paginator := inspector.NewListFindingsPaginator(client, &inspector.ListFindingsInput{})

	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		if len(page.FindingArns) == 0 {
			continue
		}

		findings, err := client.DescribeFindings(ctx, &inspector.DescribeFindingsInput{
			FindingArns: page.FindingArns,
		})
		if err != nil {
			return nil, err
		}

		for _, finding := range findings.Findings {
			resource := Resource{
				Region: describeCtx.KaytuRegion,
				Name:   *finding.Title,
				ID:     *finding.Id,
				ARN:    *finding.Arn,
				Description: model.InspectorFindingDescription{
					Finding:     finding,
					FailedItems: findings.FailedItems,
				},
			}
			if stream != nil {
				if err := (*stream)(resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, resource)
			}
		}
	}

	return values, nil
}
