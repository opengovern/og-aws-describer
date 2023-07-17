package describer

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/securityhub"
	"github.com/aws/aws-sdk-go-v2/service/securityhub/types"
	"github.com/kaytu-io/kaytu-aws-describer/aws/model"
)

func SecurityHubHub(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	client := securityhub.NewFromConfig(cfg)
	out, err := client.DescribeHub(ctx, &securityhub.DescribeHubInput{})
	if err != nil {
		if isErr(err, "InvalidAccessException") {
			return nil, nil
		}
		return nil, err
	}

	var values []Resource

	resource, err := securityHubHubHandle(ctx, cfg, out)
	if err != nil {
		return nil, err
	}

	if stream != nil {
		if err := (*stream)(resource); err != nil {
			return nil, err
		}
	} else {
		values = append(values, resource)
	}

	return values, nil
}
func securityHubHubHandle(ctx context.Context, cfg aws.Config, out *securityhub.DescribeHubOutput) (Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := securityhub.NewFromConfig(cfg)

	tags, err := client.ListTagsForResource(ctx, &securityhub.ListTagsForResourceInput{ResourceArn: out.HubArn})
	if err != nil {
		if isErr(err, "InvalidAccessException") {
			return Resource{}, nil
		}
		return Resource{}, err
	}

	resource := Resource{
		Region: describeCtx.KaytuRegion,
		ARN:    *out.HubArn,
		Name:   nameFromArn(*out.HubArn),
		Description: model.SecurityHubHubDescription{
			Hub:  out,
			Tags: tags.Tags,
		},
	}
	return resource, nil
}
func GetSecurityHubHub(ctx context.Context, cfg aws.Config, fields map[string]string) ([]Resource, error) {
	arn := fields["arn"]
	var value []Resource
	client := securityhub.NewFromConfig(cfg)

	out, err := client.DescribeHub(ctx, &securityhub.DescribeHubInput{
		HubArn: &arn,
	})
	if err != nil {
		if isErr(err, "DescribeHubNotFound") || isErr(err, "InvalidParameterValue") {
			return nil, nil
		}
		return nil, err
	}

	resource, err := securityHubHubHandle(ctx, cfg, out)
	if err != nil {
		return nil, err
	}
	value = append(value, resource)
	return value, nil
}

func SecurityHubActionTarget(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := securityhub.NewFromConfig(cfg)

	var values []Resource
	paginator := securityhub.NewDescribeActionTargetsPaginator(client, &securityhub.DescribeActionTargetsInput{})
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			if isErr(err, "InvalidAccessException") {
				return nil, nil
			}
			return nil, err
		}

		for _, actionTarget := range page.ActionTargets {
			resource := Resource{
				Region: describeCtx.KaytuRegion,
				ARN:    *actionTarget.ActionTargetArn,
				Name:   *actionTarget.Name,
				Description: model.SecurityHubActionTargetDescription{
					ActionTarget: actionTarget,
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

func SecurityHubFinding(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := securityhub.NewFromConfig(cfg)

	var values []Resource
	paginator := securityhub.NewGetFindingsPaginator(client, &securityhub.GetFindingsInput{})
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			if isErr(err, "InvalidAccessException") {
				return nil, nil
			}
			return nil, err
		}

		for _, finding := range page.Findings {
			resource := Resource{
				Region: describeCtx.KaytuRegion,
				ID:     *finding.Id,
				Name:   *finding.Title,
				Description: model.SecurityHubFindingDescription{
					Finding: finding,
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

func SecurityHubFindingAggregator(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := securityhub.NewFromConfig(cfg)

	var values []Resource
	paginator := securityhub.NewListFindingAggregatorsPaginator(client, &securityhub.ListFindingAggregatorsInput{})
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			if isErr(err, "InvalidAccessException") {
				return nil, nil
			}
			return nil, err
		}

		for _, findingAggregatorSummary := range page.FindingAggregators {
			findingAggregator, err := client.GetFindingAggregator(ctx, &securityhub.GetFindingAggregatorInput{
				FindingAggregatorArn: findingAggregatorSummary.FindingAggregatorArn,
			})
			if err != nil {
				if isErr(err, "InvalidAccessException") {
					return nil, nil
				}
				return nil, err
			}
			resource := Resource{
				Region: describeCtx.KaytuRegion,
				ARN:    *findingAggregator.FindingAggregatorArn,
				Description: model.SecurityHubFindingAggregatorDescription{
					FindingAggregator: *findingAggregator,
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

func SecurityHubInsight(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := securityhub.NewFromConfig(cfg)

	var values []Resource
	paginator := securityhub.NewGetInsightsPaginator(client, &securityhub.GetInsightsInput{})
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			if isErr(err, "InvalidAccessException") {
				return nil, nil
			}
			return nil, err
		}

		for _, insight := range page.Insights {
			resource := Resource{
				Region: describeCtx.KaytuRegion,
				ARN:    *insight.InsightArn,
				Name:   *insight.Name,
				Description: model.SecurityHubInsightDescription{
					Insight: insight,
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

func SecurityHubMember(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := securityhub.NewFromConfig(cfg)

	var values []Resource
	paginator := securityhub.NewListMembersPaginator(client, &securityhub.ListMembersInput{})
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			if isErr(err, "InvalidAccessException") || isErr(err, "InvalidInputException") {
				return nil, nil
			}
			return nil, err
		}

		for _, member := range page.Members {
			resource := Resource{
				Region: describeCtx.KaytuRegion,
				Name:   *member.AccountId,
				Description: model.SecurityHubMemberDescription{
					Member: member,
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

func SecurityHubProduct(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := securityhub.NewFromConfig(cfg)

	var values []Resource
	paginator := securityhub.NewDescribeProductsPaginator(client, &securityhub.DescribeProductsInput{})
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			if isErr(err, "InvalidAccessException") {
				return nil, nil
			}
			return nil, err
		}

		for _, product := range page.Products {
			resource := Resource{
				Region: describeCtx.KaytuRegion,
				Name:   *product.ProductName,
				ARN:    *product.ProductArn,
				Description: model.SecurityHubProductDescription{
					Product: product,
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

func SecurityHubStandardsControl(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := securityhub.NewFromConfig(cfg)

	var values []Resource

	subPaginator := securityhub.NewGetEnabledStandardsPaginator(client, &securityhub.GetEnabledStandardsInput{})
	for subPaginator.HasMorePages() {
		subPage, err := subPaginator.NextPage(ctx)
		if err != nil {
			if isErr(err, "InvalidAccessException") {
				return nil, nil
			}
			return nil, err
		}

		for _, standard := range subPage.StandardsSubscriptions {
			paginator := securityhub.NewDescribeStandardsControlsPaginator(client, &securityhub.DescribeStandardsControlsInput{
				StandardsSubscriptionArn: standard.StandardsArn,
			})
			for paginator.HasMorePages() {
				page, err := paginator.NextPage(ctx)
				if err != nil {
					if isErr(err, "InvalidAccessException") || isErr(err, "InvalidInputException") {
						return nil, nil
					}
					return nil, err
				}

				for _, standardsControl := range page.Controls {
					resource := Resource{
						Region: describeCtx.KaytuRegion,
						ID:     *standardsControl.ControlId,
						Name:   *standardsControl.Title,
						ARN:    *standardsControl.StandardsControlArn,
						Description: model.SecurityHubStandardsControlDescription{
							StandardsControl: standardsControl,
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

func SecurityHubStandardsSubscription(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := securityhub.NewFromConfig(cfg)

	var values []Resource

	standardsPaginator := securityhub.NewDescribeStandardsPaginator(client, &securityhub.DescribeStandardsInput{})
	standards := make(map[string]types.Standard)
	for standardsPaginator.HasMorePages() {
		standardsPage, err := standardsPaginator.NextPage(ctx)
		if err != nil {
			if isErr(err, "InvalidAccessException") {
				return nil, nil
			}
			return nil, err
		}
		for _, standard := range standardsPage.Standards {
			standards[*standard.StandardsArn] = standard
		}
	}

	paginator := securityhub.NewGetEnabledStandardsPaginator(client, &securityhub.GetEnabledStandardsInput{})
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			if isErr(err, "InvalidAccessException") {
				return nil, nil
			}
			return nil, err
		}

		for _, standardSub := range page.StandardsSubscriptions {
			standard, _ := standards[*standardSub.StandardsArn]
			resource := Resource{
				Region: describeCtx.KaytuRegion,
				ARN:    *standardSub.StandardsSubscriptionArn,
				Description: model.SecurityHubStandardsSubscriptionDescription{
					Standard:              standard,
					StandardsSubscription: standardSub,
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
