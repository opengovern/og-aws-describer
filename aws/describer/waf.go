package describer

import (
	"context"
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront"
	"github.com/aws/aws-sdk-go-v2/service/waf"
	waftypes "github.com/aws/aws-sdk-go-v2/service/waf/types"
	"github.com/aws/aws-sdk-go-v2/service/wafregional"
	regionaltypes "github.com/aws/aws-sdk-go-v2/service/wafregional/types"
	"github.com/aws/aws-sdk-go-v2/service/wafv2"
	"github.com/aws/aws-sdk-go-v2/service/wafv2/types"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/kaytu-io/kaytu-aws-describer/aws/model"
)

func WAFv2IPSet(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := wafv2.NewFromConfig(cfg)

	scopes := []types.Scope{
		types.ScopeRegional,
	}
	if strings.EqualFold(cfg.Region, "us-east-1") {
		scopes = append(scopes, types.ScopeCloudfront)
	}

	var values []Resource
	for _, scope := range scopes {
		err := PaginateRetrieveAll(func(prevToken *string) (nextToken *string, err error) {
			output, err := client.ListIPSets(ctx, &wafv2.ListIPSetsInput{
				Scope:      scope,
				NextMarker: prevToken,
			})
			if err != nil {
				return nil, err
			}

			for _, v := range output.IPSets {
				resource := Resource{
					Region:      describeCtx.Region,
					ARN:         *v.ARN,
					Name:        *v.Name,
					Description: v,
				}
				if stream != nil {
					if err := (*stream)(resource); err != nil {
						return nil, err
					}
				} else {
					values = append(values, resource)
				}

			}
			return output.NextMarker, nil
		})
		if err != nil {
			return nil, err
		}
	}

	return values, nil
}

func WAFv2LoggingConfiguration(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := wafv2.NewFromConfig(cfg)

	scopes := []types.Scope{
		types.ScopeRegional,
	}
	if strings.EqualFold(cfg.Region, "us-east-1") {
		scopes = append(scopes, types.ScopeCloudfront)
	}

	var values []Resource
	for _, scope := range scopes {
		err := PaginateRetrieveAll(func(prevToken *string) (nextToken *string, err error) {
			output, err := client.ListLoggingConfigurations(ctx, &wafv2.ListLoggingConfigurationsInput{
				Scope:      scope,
				NextMarker: prevToken,
			})
			if err != nil {
				return nil, err
			}

			for _, v := range output.LoggingConfigurations {
				resource := Resource{
					Region:      describeCtx.Region,
					ARN:         *v.ResourceArn, // TODO: might not be the actual ARN
					Name:        *v.ResourceArn,
					Description: v,
				}
				if stream != nil {
					if err := (*stream)(resource); err != nil {
						return nil, err
					}
				} else {
					values = append(values, resource)
				}

			}
			return output.NextMarker, nil
		})
		if err != nil {
			return nil, err
		}
	}

	return values, nil
}

func WAFv2RegexPatternSet(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := wafv2.NewFromConfig(cfg)

	scopes := []types.Scope{
		types.ScopeRegional,
	}
	if strings.EqualFold(cfg.Region, "us-east-1") {
		scopes = append(scopes, types.ScopeCloudfront)
	}

	var values []Resource
	for _, scope := range scopes {
		err := PaginateRetrieveAll(func(prevToken *string) (nextToken *string, err error) {
			output, err := client.ListRegexPatternSets(ctx, &wafv2.ListRegexPatternSetsInput{
				Scope:      scope,
				NextMarker: prevToken,
			})
			if err != nil {
				return nil, err
			}

			for _, v := range output.RegexPatternSets {
				resource := Resource{
					Region:      describeCtx.Region,
					ARN:         *v.ARN,
					Name:        *v.Name,
					Description: v,
				}
				if stream != nil {
					if err := (*stream)(resource); err != nil {
						return nil, err
					}
				} else {
					values = append(values, resource)
				}

			}
			return output.NextMarker, nil
		})
		if err != nil {
			return nil, err
		}
	}

	return values, nil
}

func WAFv2RuleGroup(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := wafv2.NewFromConfig(cfg)

	scopes := []types.Scope{
		types.ScopeRegional,
	}
	if strings.EqualFold(cfg.Region, "us-east-1") {
		scopes = append(scopes, types.ScopeCloudfront)
	}

	var values []Resource
	for _, scope := range scopes {
		err := PaginateRetrieveAll(func(prevToken *string) (nextToken *string, err error) {
			output, err := client.ListRuleGroups(ctx, &wafv2.ListRuleGroupsInput{
				Scope:      scope,
				NextMarker: prevToken,
			})
			if err != nil {
				return nil, err
			}

			for _, v := range output.RuleGroups {
				resource := Resource{
					Region:      describeCtx.Region,
					ARN:         *v.ARN,
					Name:        *v.Name,
					Description: v,
				}
				if stream != nil {
					if err := (*stream)(resource); err != nil {
						return nil, err
					}
				} else {
					values = append(values, resource)
				}

			}
			return output.NextMarker, nil
		})
		if err != nil {
			return nil, err
		}
	}

	return values, nil
}

func WAFv2WebACL(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := wafv2.NewFromConfig(cfg)

	scopes := []types.Scope{
		types.ScopeRegional,
	}
	if strings.EqualFold(cfg.Region, "us-east-1") {
		scopes = append(scopes, types.ScopeCloudfront)
	}

	var values []Resource
	for _, scope := range scopes {
		acls, err := listWAFv2WebACLs(ctx, cfg, scope)
		if err != nil {
			return nil, err
		}

		for _, v := range acls {
			out, err := client.GetWebACL(ctx, &wafv2.GetWebACLInput{
				Id:    v.Id,
				Name:  v.Name,
				Scope: scope,
			})
			if err != nil {
				return nil, err
			}

			logC, err := client.GetLoggingConfiguration(ctx, &wafv2.GetLoggingConfigurationInput{
				ResourceArn: out.WebACL.ARN,
			})
			if err != nil {
				if isErr(err, "WAFNonexistentItemException") {
					logC = &wafv2.GetLoggingConfigurationOutput{}
					err = nil
				}
				if a, ok := err.(awserr.Error); ok {
					if a.Code() == "WAFNonexistentItemException" {
						logC = &wafv2.GetLoggingConfigurationOutput{}
						err = nil
					}
				}

				if err != nil {
					return nil, err
				}
			}

			tags, err := client.ListTagsForResource(ctx, &wafv2.ListTagsForResourceInput{
				ResourceARN: out.WebACL.ARN,
			})
			if err != nil {
				return nil, err
			}

			resource := Resource{
				Region: describeCtx.Region,
				ARN:    *v.ARN,
				Name:   *v.Name,
				Description: model.WAFv2WebACLDescription{
					WebACL:               out.WebACL,
					Scope:                scope,
					LoggingConfiguration: logC.LoggingConfiguration,
					TagInfoForResource:   tags.TagInfoForResource,
					LockToken:            v.LockToken,
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

func listWAFv2WebACLs(ctx context.Context, cfg aws.Config, scope types.Scope) ([]types.WebACLSummary, error) {
	client := wafv2.NewFromConfig(cfg)

	var acls []types.WebACLSummary
	err := PaginateRetrieveAll(func(prevToken *string) (nextToken *string, err error) {
		output, err := client.ListWebACLs(ctx, &wafv2.ListWebACLsInput{
			Scope:      scope,
			NextMarker: prevToken,
		})
		if err != nil {
			return nil, err
		}

		acls = append(acls, output.WebACLs...)
		return output.NextMarker, nil
	})
	if err != nil {
		return nil, err
	}

	return acls, nil
}

// Returns ResourceArns that have a WebAcl Associated
func WAFv2WebACLAssociation(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	var values []Resource

	regionalACls, err := listWAFv2WebACLs(ctx, cfg, types.ScopeRegional)
	if err != nil {
		return nil, err
	}

	client := wafv2.NewFromConfig(cfg)
	for _, acl := range regionalACls {
		output, err := client.ListResourcesForWebACL(ctx, &wafv2.ListResourcesForWebACLInput{
			WebACLArn: acl.ARN,
		})
		if err != nil {
			return nil, err
		}

		resource := Resource{
			Region: describeCtx.Region,
			ID:     *acl.Id, // Unique per WebACL
			Name:   *acl.Name,
			Description: map[string]interface{}{
				"WebACLArn":    *acl.ARN,
				"ResourceArns": output.ResourceArns,
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

	if strings.EqualFold(cfg.Region, "us-east-1") {
		cloudFrontAcls, err := listWAFv2WebACLs(ctx, cfg, types.ScopeCloudfront)
		if err != nil {
			return nil, err
		}

		cfClient := cloudfront.NewFromConfig(cfg)
		for _, acl := range cloudFrontAcls {
			err := PaginateRetrieveAll(func(prevToken *string) (nextToken *string, err error) {
				output, err := cfClient.ListDistributionsByWebACLId(ctx, &cloudfront.ListDistributionsByWebACLIdInput{
					WebACLId: acl.Id,
					Marker:   prevToken,
				})
				if err != nil {
					return nil, err
				}

				resource := Resource{
					Region: describeCtx.Region,
					ID:     *acl.Id, // Unique per WebACL
					Name:   *acl.Name,
					Description: map[string]interface{}{
						"WebACLArn":     *acl.ARN,
						"Distributions": output.DistributionList.Items,
					},
				}
				if stream != nil {
					if err := (*stream)(resource); err != nil {
						return nil, err
					}
				} else {
					values = append(values, resource)
				}

				return output.DistributionList.NextMarker, nil
			})
			if err != nil {
				return nil, err
			}
		}
	}

	return values, nil
}

func WAFRegionalByteMatchSet(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := wafregional.NewFromConfig(cfg)

	var values []Resource
	err := PaginateRetrieveAll(func(prevToken *string) (nextToken *string, err error) {
		output, err := client.ListByteMatchSets(ctx, &wafregional.ListByteMatchSetsInput{
			NextMarker: prevToken,
		})
		if err != nil {
			return nil, err
		}

		for _, v := range output.ByteMatchSets {
			resource := Resource{
				Region:      describeCtx.Region,
				ID:          *v.ByteMatchSetId,
				Name:        *v.Name,
				Description: v,
			}
			if stream != nil {
				if err := (*stream)(resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, resource)
			}

		}
		return output.NextMarker, nil
	})
	if err != nil {
		return nil, err
	}

	return values, nil
}

func WAFRegionalGeoMatchSet(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := wafregional.NewFromConfig(cfg)

	var values []Resource
	err := PaginateRetrieveAll(func(prevToken *string) (nextToken *string, err error) {
		output, err := client.ListGeoMatchSets(ctx, &wafregional.ListGeoMatchSetsInput{
			NextMarker: prevToken,
		})
		if err != nil {
			return nil, err
		}

		for _, v := range output.GeoMatchSets {
			resource := Resource{
				Region:      describeCtx.Region,
				ID:          *v.GeoMatchSetId,
				Name:        *v.Name,
				Description: v,
			}
			if stream != nil {
				if err := (*stream)(resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, resource)
			}

		}
		return output.NextMarker, nil
	})
	if err != nil {
		return nil, err
	}

	return values, nil
}

func WAFRegionalIPSet(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := wafregional.NewFromConfig(cfg)

	var values []Resource
	err := PaginateRetrieveAll(func(prevToken *string) (nextToken *string, err error) {
		output, err := client.ListIPSets(ctx, &wafregional.ListIPSetsInput{
			NextMarker: prevToken,
		})
		if err != nil {
			return nil, err
		}

		for _, v := range output.IPSets {
			resource := Resource{
				Region:      describeCtx.Region,
				ID:          *v.IPSetId,
				Name:        *v.Name,
				Description: v,
			}
			if stream != nil {
				if err := (*stream)(resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, resource)
			}

		}
		return output.NextMarker, nil
	})
	if err != nil {
		return nil, err
	}

	return values, nil
}

func WAFRegionalRateBasedRule(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := wafregional.NewFromConfig(cfg)

	var values []Resource
	err := PaginateRetrieveAll(func(prevToken *string) (nextToken *string, err error) {
		output, err := client.ListRateBasedRules(ctx, &wafregional.ListRateBasedRulesInput{
			NextMarker: prevToken,
		})
		if err != nil {
			return nil, err
		}

		for _, v := range output.Rules {
			resource := Resource{
				Region:      describeCtx.Region,
				ID:          *v.RuleId,
				Name:        *v.Name,
				Description: v,
			}
			if stream != nil {
				if err := (*stream)(resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, resource)
			}

		}
		return output.NextMarker, nil
	})
	if err != nil {
		return nil, err
	}

	return values, nil
}

func WAFRegionalRegexPatternSet(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := wafregional.NewFromConfig(cfg)

	var values []Resource
	err := PaginateRetrieveAll(func(prevToken *string) (nextToken *string, err error) {
		output, err := client.ListRegexPatternSets(ctx, &wafregional.ListRegexPatternSetsInput{
			NextMarker: prevToken,
		})
		if err != nil {
			return nil, err
		}

		for _, v := range output.RegexPatternSets {
			resource := Resource{
				Region:      describeCtx.Region,
				ID:          *v.RegexPatternSetId,
				Name:        *v.Name,
				Description: v,
			}
			if stream != nil {
				if err := (*stream)(resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, resource)
			}

		}
		return output.NextMarker, nil
	})
	if err != nil {
		return nil, err
	}

	return values, nil
}

func WAFRegionalRule(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)

	client := wafregional.NewFromConfig(cfg)

	var values []Resource
	err := PaginateRetrieveAll(func(prevToken *string) (nextToken *string, err error) {
		output, err := client.ListRules(ctx, &wafregional.ListRulesInput{
			NextMarker: prevToken,
		})
		if err != nil {
			return nil, err
		}

		for _, v := range output.Rules {
			arn := fmt.Sprintf("arn:%s:waf-regional:%s:%s:rule/%s", describeCtx.Partition, describeCtx.Region, describeCtx.AccountID, *v.RuleId)

			tags, err := client.ListTagsForResource(ctx, &wafregional.ListTagsForResourceInput{
				ResourceARN: &arn,
			})
			if err != nil {
				return nil, err
			}

			resource := Resource{
				Region: describeCtx.Region,
				ARN:    arn,
				ID:     *v.RuleId,
				Name:   *v.Name,
				Description: model.WAFRegionalRuleDescription{
					Rule: v,
					Tags: tags.TagInfoForResource.TagList,
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
		return output.NextMarker, nil
	})
	if err != nil {
		return nil, err
	}

	return values, nil
}

func WAFRegionalSizeConstraintSet(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := wafregional.NewFromConfig(cfg)

	var values []Resource
	err := PaginateRetrieveAll(func(prevToken *string) (nextToken *string, err error) {
		output, err := client.ListSizeConstraintSets(ctx, &wafregional.ListSizeConstraintSetsInput{
			NextMarker: prevToken,
		})
		if err != nil {
			return nil, err
		}

		for _, v := range output.SizeConstraintSets {
			resource := Resource{
				Region:      describeCtx.Region,
				ID:          *v.SizeConstraintSetId,
				Name:        *v.Name,
				Description: v,
			}
			if stream != nil {
				if err := (*stream)(resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, resource)
			}

		}
		return output.NextMarker, nil
	})
	if err != nil {
		return nil, err
	}

	return values, nil
}

func WAFRegionalSqlInjectionMatchSet(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := wafregional.NewFromConfig(cfg)

	var values []Resource
	err := PaginateRetrieveAll(func(prevToken *string) (nextToken *string, err error) {
		output, err := client.ListSqlInjectionMatchSets(ctx, &wafregional.ListSqlInjectionMatchSetsInput{
			NextMarker: prevToken,
		})
		if err != nil {
			return nil, err
		}

		for _, v := range output.SqlInjectionMatchSets {
			resource := Resource{
				Region:      describeCtx.Region,
				ID:          *v.SqlInjectionMatchSetId,
				Name:        *v.Name,
				Description: v,
			}
			if stream != nil {
				if err := (*stream)(resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, resource)
			}

		}
		return output.NextMarker, nil
	})
	if err != nil {
		return nil, err
	}

	return values, nil
}

func WAFRegionalWebACL(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := wafregional.NewFromConfig(cfg)

	var values []Resource
	err := PaginateRetrieveAll(func(prevToken *string) (nextToken *string, err error) {
		output, err := client.ListWebACLs(ctx, &wafregional.ListWebACLsInput{
			NextMarker: prevToken,
		})
		if err != nil {
			return nil, err
		}

		for _, v := range output.WebACLs {
			resource := Resource{
				Region:      describeCtx.Region,
				ID:          *v.WebACLId,
				Name:        *v.Name,
				Description: v,
			}
			if stream != nil {
				if err := (*stream)(resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, resource)
			}

		}
		return output.NextMarker, nil
	})
	if err != nil {
		return nil, err
	}

	return values, nil
}

func WAFRegionalWebACLAssociation(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	acls, err := WAFRegionalWebACL(ctx, cfg, nil)
	if err != nil {
		return nil, err
	}

	client := wafregional.NewFromConfig(cfg)

	var values []Resource
	for _, a := range acls {
		acl := a.Description.(regionaltypes.WebACLSummary)
		output, err := client.ListResourcesForWebACL(ctx, &wafregional.ListResourcesForWebACLInput{
			WebACLId: acl.WebACLId,
		})
		if err != nil {
			return nil, err
		}

		resource := Resource{
			Region: describeCtx.Region,
			ID:     *acl.WebACLId, // Unique per WebACL
			Name:   *acl.Name,
			Description: map[string]interface{}{
				"WebACLId":     *acl.WebACLId,
				"ResourceArns": output.ResourceArns,
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

	return values, nil
}

func WAFRegionalXssMatchSet(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := wafregional.NewFromConfig(cfg)

	var values []Resource
	err := PaginateRetrieveAll(func(prevToken *string) (nextToken *string, err error) {
		output, err := client.ListXssMatchSets(ctx, &wafregional.ListXssMatchSetsInput{
			NextMarker: prevToken,
		})
		if err != nil {
			return nil, err
		}

		for _, v := range output.XssMatchSets {
			resource := Resource{
				Region:      describeCtx.Region,
				ID:          *v.XssMatchSetId,
				Name:        *v.Name,
				Description: v,
			}
			if stream != nil {
				if err := (*stream)(resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, resource)
			}

		}
		return output.NextMarker, nil
	})
	if err != nil {
		return nil, err
	}

	return values, nil
}

func WAFRule(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)

	client := waf.NewFromConfig(cfg)

	var values []Resource
	err := PaginateRetrieveAll(func(prevToken *string) (nextToken *string, err error) {
		output, err := client.ListRules(ctx, &waf.ListRulesInput{
			NextMarker: prevToken,
		})
		if err != nil {
			if !isErr(err, "WAFNonexistentItemException") {
				return nil, err
			}
			return nil, nil
		}

		for _, v := range output.Rules {
			rule, err := client.GetRule(ctx, &waf.GetRuleInput{
				RuleId: v.RuleId,
			})
			if err != nil {
				if !isErr(err, "WAFNonexistentItemException") {
					return nil, err
				}
				continue
			}

			arn := fmt.Sprintf("arn:%s:waf::%s:rule/%s", describeCtx.Partition, describeCtx.AccountID, *v.RuleId)

			tags, err := client.ListTagsForResource(ctx, &waf.ListTagsForResourceInput{
				ResourceARN: &arn,
			})
			if err != nil {
				if !isErr(err, "WAFNonexistentItemException") {
					return nil, err
				}
				tags = &waf.ListTagsForResourceOutput{
					TagInfoForResource: &waftypes.TagInfoForResource{},
				}
			}

			resource := Resource{
				Region: describeCtx.Region,
				ARN:    arn,
				Name:   *rule.Rule.Name,
				Description: model.WAFRuleDescription{
					Rule: *rule.Rule,
					Tags: tags.TagInfoForResource.TagList,
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
		return output.NextMarker, nil
	})
	if err != nil {
		return nil, err
	}

	return values, nil
}
