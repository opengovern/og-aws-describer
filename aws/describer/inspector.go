package describer

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/inspector"
	"github.com/kaytu-io/kaytu-aws-describer/aws/model"
)

func InspectorAssessmentRun(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	client := inspector.NewFromConfig(cfg)
	paginator := inspector.NewListAssessmentRunsPaginator(client, &inspector.ListAssessmentRunsInput{})

	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		assessmentRuns, err := client.DescribeAssessmentRuns(ctx, &inspector.DescribeAssessmentRunsInput{
			AssessmentRunArns: page.AssessmentRunArns,
		})
		if err != nil {
			return nil, err
		}

		for _, assessmentRun := range assessmentRuns.AssessmentRuns {
			resource := Resource{
				Name: *assessmentRun.Name,
				ARN:  *assessmentRun.Arn,
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

		assessmentTargets, err := client.DescribeAssessmentTargets(ctx, &inspector.DescribeAssessmentTargetsInput{
			AssessmentTargetArns: page.AssessmentTargetArns,
		})
		if err != nil {
			return nil, err
		}

		for _, assessmentTarget := range assessmentTargets.AssessmentTargets {
			resource := Resource{
				Name: *assessmentTarget.Name,
				ARN:  *assessmentTarget.Arn,
				Description: model.InspectorAssessmentTargetDescription{
					AssessmentTarget: assessmentTarget,
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

func InspectorAssessmentTemplate(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	client := inspector.NewFromConfig(cfg)
	paginator := inspector.NewListAssessmentTemplatesPaginator(client, &inspector.ListAssessmentTemplatesInput{})

	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
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

			resource := Resource{
				Name: *assessmentTemplate.Name,
				ARN:  *assessmentTemplate.Arn,
				Description: model.InspectorAssessmentTemplateDescription{
					AssessmentTemplate: assessmentTemplate,
					EventSubscriptions: eventSubscriptions.Subscriptions,
					Tags:               tags.Tags,
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

		values = append(values, Resource{
			Name: *assessmentTemplate.Name,
			ARN:  *assessmentTemplate.Arn,
			Description: model.InspectorAssessmentTemplateDescription{
				AssessmentTemplate: assessmentTemplate,
				EventSubscriptions: eventSubscriptions.Subscriptions,
				Tags:               tags.Tags,
			},
		})
	}

	return values, nil
}

func InspectorExclusion(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
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
						Name: *exclusion.Title,
						ARN:  *exclusion.Arn,
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
	client := inspector.NewFromConfig(cfg)
	paginator := inspector.NewListFindingsPaginator(client, &inspector.ListFindingsInput{})

	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		findings, err := client.DescribeFindings(ctx, &inspector.DescribeFindingsInput{
			FindingArns: page.FindingArns,
		})
		if err != nil {
			return nil, err
		}

		for _, finding := range findings.Findings {
			resource := Resource{
				Name: *finding.Title,
				ID:   *finding.Id,
				ARN:  *finding.Arn,
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
