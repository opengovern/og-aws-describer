package describer

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/networkfirewall"
	"github.com/kaytu-io/kaytu-aws-describer/aws/model"
)

func NetworkFirewallFirewall(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	client := networkfirewall.NewFromConfig(cfg)
	paginator := networkfirewall.NewListFirewallsPaginator(client, &networkfirewall.ListFirewallsInput{})

	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, v := range page.Firewalls {
			firewall, err := client.DescribeFirewall(ctx, &networkfirewall.DescribeFirewallInput{
				FirewallArn: v.FirewallArn,
			})
			if err != nil {
				return nil, err
			}

			resource := NetworkFirewallFirewallHandle(ctx, firewall, *v.FirewallName, *v.FirewallArn)
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
func NetworkFirewallFirewallHandle(ctx context.Context, firewall *networkfirewall.DescribeFirewallOutput, firewallName string, firewallArn string) Resource {
	describeCtx := GetDescribeContext(ctx)
	resource := Resource{
		Region: describeCtx.KaytuRegion,
		ARN:    firewallArn,
		Name:   firewallName,
		Description: model.NetworkFirewallFirewallDescription{
			Firewall: *firewall.Firewall,
		},
	}
	return resource
}
func GetNetworkFirewallFirewall(ctx context.Context, cfg aws.Config, fields map[string]string) ([]Resource, error) {
	var values []Resource
	firewallName := fields["firewallName"]
	firewallArn := fields["firewallArn"]
	client := networkfirewall.NewFromConfig(cfg)
	firewall, err := client.DescribeFirewall(ctx, &networkfirewall.DescribeFirewallInput{
		FirewallName: &firewallName,
		FirewallArn:  &firewallArn,
	})
	if err != nil {
		return nil, err
	}

	resource := NetworkFirewallFirewallHandle(ctx, firewall, firewallName, firewallArn)
	values = append(values, resource)
	return values, nil
}

func NetworkFirewallPolicy(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := networkfirewall.NewFromConfig(cfg)
	paginator := networkfirewall.NewListFirewallPoliciesPaginator(client, &networkfirewall.ListFirewallPoliciesInput{})

	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, v := range page.FirewallPolicies {
			if v.Arn == nil {
				continue
			}

			data, err := client.DescribeFirewallPolicy(ctx, &networkfirewall.DescribeFirewallPolicyInput{
				FirewallPolicyArn:  v.Arn,
				FirewallPolicyName: v.Name,
			})
			if err != nil {
				return nil, err
			}

			var name string
			if v.Name != nil {
				name = *v.Name
			} else {
				name = *v.Arn
			}
			resource := Resource{
				Region: describeCtx.KaytuRegion,
				ARN:    *v.Arn,
				Name:   name,
				Description: model.NetworkFirewallFirewallPolicyDescription{
					FirewallPolicy:         data.FirewallPolicy,
					FirewallPolicyResponse: data.FirewallPolicyResponse,
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

func NetworkFirewallRuleGroup(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := networkfirewall.NewFromConfig(cfg)
	paginator := networkfirewall.NewListRuleGroupsPaginator(client, &networkfirewall.ListRuleGroupsInput{})

	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, v := range page.RuleGroups {
			if v.Arn == nil {
				continue
			}

			data, err := client.DescribeRuleGroup(ctx, &networkfirewall.DescribeRuleGroupInput{
				RuleGroupArn:  v.Arn,
				RuleGroupName: v.Name,
			})
			if err != nil {
				return nil, err
			}

			var name string
			if v.Name != nil {
				name = *v.Name
			} else {
				name = *v.Arn
			}
			resource := Resource{
				Region: describeCtx.KaytuRegion,
				ARN:    *v.Arn,
				Name:   name,
				Description: model.NetworkFirewallRuleGroupDescription{
					RuleGroup:         data.RuleGroup,
					RuleGroupResponse: data.RuleGroupResponse,
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
