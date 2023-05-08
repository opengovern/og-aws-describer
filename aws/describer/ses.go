package describer

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ses"
	"github.com/aws/aws-sdk-go-v2/service/sesv2"
	"github.com/kaytu-io/kaytu-aws-describer/aws/model"
)

func SESConfigurationSet(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)

	client := sesv2.NewFromConfig(cfg)
	paginator := sesv2.NewListConfigurationSetsPaginator(client, &sesv2.ListConfigurationSetsInput{})

	sesClient := ses.NewFromConfig(cfg)

	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, v := range page.ConfigurationSets {
			output, err := sesClient.DescribeConfigurationSet(ctx, &ses.DescribeConfigurationSetInput{ConfigurationSetName: aws.String(v)})
			if err != nil {
				return nil, err
			}

			arn := fmt.Sprintf("arn:%s:ses:%s:%s:configuration-set/%s", describeCtx.Partition, describeCtx.Region, describeCtx.AccountID, *output.ConfigurationSet.Name)

			resource := Resource{
				Region: describeCtx.KaytuRegion,
				ARN:    arn,
				Name:   *output.ConfigurationSet.Name,
				Description: model.SESConfigurationSetDescription{
					ConfigurationSet: *output.ConfigurationSet,
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

func SESIdentity(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)

	client := ses.NewFromConfig(cfg)
	clientv2 := sesv2.NewFromConfig(cfg)
	paginator := ses.NewListIdentitiesPaginator(client, &ses.ListIdentitiesInput{})

	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, v := range page.Identities {
			arn := fmt.Sprintf("arn:%s:ses:%s:%s:identity/%s", describeCtx.Partition, describeCtx.Region, describeCtx.AccountID, v)

			tags, err := clientv2.ListTagsForResource(ctx, &sesv2.ListTagsForResourceInput{
				ResourceArn: &arn,
			})
			if err != nil {
				return nil, err
			}

			identity, err := client.GetIdentityVerificationAttributes(ctx, &ses.GetIdentityVerificationAttributesInput{
				Identities: []string{v},
			})
			if err != nil {
				return nil, err
			}

			notif, err := client.GetIdentityNotificationAttributes(ctx, &ses.GetIdentityNotificationAttributesInput{
				Identities: []string{v},
			})
			if err != nil {
				return nil, err
			}

			resource := Resource{
				Region: describeCtx.KaytuRegion,
				ARN:    arn,
				Name:   v,
				Description: model.SESIdentityDescription{
					Identity:               v,
					VerificationAttributes: identity.VerificationAttributes[v],
					NotificationAttributes: notif.NotificationAttributes[v],
					Tags:                   tags.Tags,
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

func SESContactList(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := sesv2.NewFromConfig(cfg)
	paginator := sesv2.NewListContactListsPaginator(client, &sesv2.ListContactListsInput{})

	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, v := range page.ContactLists {
			resource := Resource{
				Region:      describeCtx.Region,
				ID:          *v.ContactListName,
				Name:        *v.ContactListName,
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
	}

	return values, nil
}

func SESReceiptFilter(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := ses.NewFromConfig(cfg)

	output, err := client.ListReceiptFilters(ctx, &ses.ListReceiptFiltersInput{})
	if err != nil {
		return nil, err
	}

	var values []Resource
	for _, v := range output.Filters {
		resource := Resource{
			Region:      describeCtx.Region,
			ID:          *v.Name,
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

	return values, nil
}

func SESReceiptRuleSet(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := ses.NewFromConfig(cfg)

	var values []Resource
	err := PaginateRetrieveAll(func(prevToken *string) (nextToken *string, err error) {
		output, err := client.ListReceiptRuleSets(ctx, &ses.ListReceiptRuleSetsInput{NextToken: prevToken})
		if err != nil {
			return nil, err
		}

		for _, v := range output.RuleSets {
			output, err := client.DescribeReceiptRuleSet(ctx, &ses.DescribeReceiptRuleSetInput{RuleSetName: v.Name})
			if err != nil {
				return nil, err
			}

			resource := Resource{
				Region:      describeCtx.Region,
				ID:          *output.Metadata.Name,
				Name:        *output.Metadata.Name,
				Description: output,
			}
			if stream != nil {
				if err := (*stream)(resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, resource)
			}

		}

		return output.NextToken, nil
	})
	if err != nil {
		return nil, err
	}

	return values, nil
}

func SESTemplate(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := ses.NewFromConfig(cfg)

	var values []Resource
	err := PaginateRetrieveAll(func(prevToken *string) (nextToken *string, err error) {
		output, err := client.ListTemplates(ctx, &ses.ListTemplatesInput{NextToken: prevToken})
		if err != nil {
			return nil, err
		}

		for _, v := range output.TemplatesMetadata {
			resource := Resource{
				Region:      describeCtx.Region,
				ID:          *v.Name,
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

		return output.NextToken, nil
	})
	if err != nil {
		return nil, err
	}

	return values, nil
}
