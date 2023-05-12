package describer

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go-v2/service/route53domains"
	"github.com/aws/smithy-go"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/route53"
	"github.com/aws/aws-sdk-go-v2/service/route53/types"
	"github.com/aws/aws-sdk-go-v2/service/route53resolver"
	resolvertypes "github.com/aws/aws-sdk-go-v2/service/route53resolver/types"
	"github.com/kaytu-io/kaytu-aws-describer/aws/model"
)

func Route53HealthCheck(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := route53.NewFromConfig(cfg)

	var values []Resource
	err := PaginateRetrieveAll(func(prevToken *string) (nextToken *string, err error) {
		output, err := client.ListHealthChecks(ctx, &route53.ListHealthChecksInput{Marker: prevToken})
		if err != nil {
			return nil, err
		}

		for _, v := range output.HealthChecks {
			item, err := client.GetHealthCheckStatus(ctx, &route53.GetHealthCheckStatusInput{
				HealthCheckId: v.Id,
			})
			if err != nil {
				var ae smithy.APIError
				if errors.As(err, &ae) {
					if ae.ErrorCode() == "InvalidInput" {
						item = nil
					} else {
						return nil, err
					}
				} else {
					return nil, err
				}
			}
			if item == nil {
				item = &route53.GetHealthCheckStatusOutput{}
			}

			resp, err := client.ListTagsForResource(ctx, &route53.ListTagsForResourceInput{
				ResourceId:   v.Id,
				ResourceType: "healthcheck",
			})
			if err != nil {
				return nil, err
			}

			resource := Resource{
				Region: describeCtx.Region,
				ID:     *v.Id,
				Description: model.Route53HealthCheckDescription{
					HealthCheck: v,
					Status:      item,
					Tags:        resp,
				},
			}
			if v.HealthCheckConfig != nil && v.HealthCheckConfig.FullyQualifiedDomainName != nil {
				resource.Name = *v.HealthCheckConfig.FullyQualifiedDomainName
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

func Route53HostedZone(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)

	client := route53.NewFromConfig(cfg)

	var values []Resource
	err := PaginateRetrieveAll(func(prevToken *string) (nextToken *string, err error) {
		output, err := client.ListHostedZones(ctx, &route53.ListHostedZonesInput{Marker: prevToken})
		if err != nil {
			if !isErr(err, "NoSuchHostedZone") {
				return nil, err
			}
			return nil, nil
		}

		for _, v := range output.HostedZones {
			id := strings.Split(*v.Id, "/")[2]
			arn := fmt.Sprintf("arn:%s:route53:::hostedzone/%s", describeCtx.Partition, id)

			queryLoggingConfigs, err := client.ListQueryLoggingConfigs(ctx, &route53.ListQueryLoggingConfigsInput{
				HostedZoneId: &id,
			})
			if err != nil {
				if !isErr(err, "NoSuchHostedZone") {
					return nil, err
				}
				queryLoggingConfigs = &route53.ListQueryLoggingConfigsOutput{}
			}

			dnsSec := &route53.GetDNSSECOutput{}
			if !v.Config.PrivateZone {
				dnsSec, err = client.GetDNSSEC(ctx, &route53.GetDNSSECInput{
					HostedZoneId: &id,
				})
				if err != nil {
					if !isErr(err, "NoSuchHostedZone") {
						return nil, err
					}
					dnsSec = &route53.GetDNSSECOutput{}
				}
			}

			tags, err := client.ListTagsForResource(ctx, &route53.ListTagsForResourceInput{
				ResourceId:   &id,
				ResourceType: types.TagResourceType("hostedzone"),
			})
			if err != nil {
				if !isErr(err, "NoSuchHostedZone") {
					return nil, err
				}
				tags = &route53.ListTagsForResourceOutput{
					ResourceTagSet: &types.ResourceTagSet{},
				}
			}

			resource := Resource{
				Region: describeCtx.KaytuRegion,
				ARN:    arn,
				Name:   *v.Name,
				Description: model.Route53HostedZoneDescription{
					ID:                  id,
					HostedZone:          v,
					QueryLoggingConfigs: queryLoggingConfigs.QueryLoggingConfigs,
					DNSSec:              *dnsSec,
					Tags:                tags.ResourceTagSet.Tags,
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

func GetRoute53HostedZone(ctx context.Context, cfg aws.Config, hostedZoneID string) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)

	client := route53.NewFromConfig(cfg)

	var values []Resource
	out, err := client.GetHostedZone(ctx, &route53.GetHostedZoneInput{Id: &hostedZoneID})
	if err != nil {
		return nil, err
	}

	v := out.HostedZone
	id := strings.Split(*v.Id, "/")[2]
	arn := fmt.Sprintf("arn:%s:route53:::hostedzone/%s", describeCtx.Partition, id)

	queryLoggingConfigs, err := client.ListQueryLoggingConfigs(ctx, &route53.ListQueryLoggingConfigsInput{
		HostedZoneId: &id,
	})
	if err != nil {
		if !isErr(err, "NoSuchHostedZone") {
			return nil, err
		}
		queryLoggingConfigs = &route53.ListQueryLoggingConfigsOutput{}
	}

	dnsSec := &route53.GetDNSSECOutput{}
	if !v.Config.PrivateZone {
		dnsSec, err = client.GetDNSSEC(ctx, &route53.GetDNSSECInput{
			HostedZoneId: &id,
		})
		if err != nil {
			if !isErr(err, "NoSuchHostedZone") {
				return nil, err
			}
			dnsSec = &route53.GetDNSSECOutput{}
		}
	}

	tags, err := client.ListTagsForResource(ctx, &route53.ListTagsForResourceInput{
		ResourceId:   &id,
		ResourceType: types.TagResourceType("hostedzone"),
	})
	if err != nil {
		if !isErr(err, "NoSuchHostedZone") {
			return nil, err
		}
		tags = &route53.ListTagsForResourceOutput{
			ResourceTagSet: &types.ResourceTagSet{},
		}
	}

	values = append(values, Resource{
		Region: describeCtx.KaytuRegion,
		ARN:    arn,
		Name:   *v.Name,
		Description: model.Route53HostedZoneDescription{
			ID:                  id,
			HostedZone:          *v,
			QueryLoggingConfigs: queryLoggingConfigs.QueryLoggingConfigs,
			DNSSec:              *dnsSec,
			Tags:                tags.ResourceTagSet.Tags,
		},
	})

	if err != nil {
		return nil, err
	}

	return values, nil
}

func Route53DNSSEC(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	zones, err := Route53HostedZone(ctx, cfg, nil)
	if err != nil {
		return nil, err
	}

	client := route53.NewFromConfig(cfg)

	var values []Resource
	for _, zone := range zones {
		id := zone.Description.(types.HostedZone).Id
		v, err := client.GetDNSSEC(ctx, &route53.GetDNSSECInput{
			HostedZoneId: id,
		})
		if err != nil {
			return nil, err
		}

		resource := Resource{
			Region:      describeCtx.Region,
			ID:          *id, // Unique per HostedZone
			Name:        *id,
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

func Route53RecordSet(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	zones, err := Route53HostedZone(ctx, cfg, nil)
	if err != nil {
		return nil, err
	}

	client := route53.NewFromConfig(cfg)

	var values []Resource
	for _, zone := range zones {
		id := zone.Description.(types.HostedZone).Id
		var prevType types.RRType
		err = PaginateRetrieveAll(func(prevName *string) (nextName *string, err error) {
			output, err := client.ListResourceRecordSets(ctx, &route53.ListResourceRecordSetsInput{
				HostedZoneId:    id,
				StartRecordName: prevName,
				StartRecordType: prevType,
			})
			if err != nil {
				return nil, err
			}

			for _, v := range output.ResourceRecordSets {
				resource := Resource{
					Region:      describeCtx.Region,
					ID:          CompositeID(*id, *v.Name),
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

			prevType = output.NextRecordType
			return output.NextRecordName, nil
		})
		if err != nil {
			return nil, err
		}
	}

	return values, nil
}

func Route53ResolverFirewallDomainList(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := route53resolver.NewFromConfig(cfg)
	paginator := route53resolver.NewListFirewallDomainListsPaginator(client, &route53resolver.ListFirewallDomainListsInput{})

	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, v := range page.FirewallDomainLists {
			resource := Resource{
				Region:      describeCtx.Region,
				ARN:         *v.Arn,
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
	}

	return values, nil
}

func Route53ResolverFirewallRuleGroup(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := route53resolver.NewFromConfig(cfg)
	paginator := route53resolver.NewListFirewallRuleGroupsPaginator(client, &route53resolver.ListFirewallRuleGroupsInput{})

	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, v := range page.FirewallRuleGroups {
			resource := Resource{
				Region:      describeCtx.Region,
				ARN:         *v.Arn,
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
	}

	return values, nil
}

func Route53ResolverFirewallRuleGroupAssociation(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := route53resolver.NewFromConfig(cfg)
	paginator := route53resolver.NewListFirewallRuleGroupAssociationsPaginator(client, &route53resolver.ListFirewallRuleGroupAssociationsInput{})

	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, v := range page.FirewallRuleGroupAssociations {
			resource := Resource{
				Region:      describeCtx.Region,
				ARN:         *v.Arn,
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
	}

	return values, nil
}

func Route53ResolverResolverDNSSECConfig(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	vpcs, err := EC2VPC(ctx, cfg, nil)
	if err != nil {
		return nil, err
	}

	client := route53resolver.NewFromConfig(cfg)

	var values []Resource
	for _, vpc := range vpcs {
		v, err := client.GetResolverDnssecConfig(ctx, &route53resolver.GetResolverDnssecConfigInput{
			ResourceId: vpc.Description.(model.EC2VpcDescription).Vpc.VpcId,
		})
		if err != nil {
			return nil, err
		}

		resource := Resource{
			Region:      describeCtx.Region,
			ID:          *v.ResolverDNSSECConfig.Id,
			Name:        *v.ResolverDNSSECConfig.Id,
			Description: v.ResolverDNSSECConfig,
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

func Route53ResolverResolverQueryLoggingConfig(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := route53resolver.NewFromConfig(cfg)
	paginator := route53resolver.NewListResolverQueryLogConfigsPaginator(client, &route53resolver.ListResolverQueryLogConfigsInput{})

	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, v := range page.ResolverQueryLogConfigs {
			resource := Resource{
				Region:      describeCtx.Region,
				ARN:         *v.Arn,
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
	}

	return values, nil
}

func Route53ResolverResolverQueryLoggingConfigAssociation(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := route53resolver.NewFromConfig(cfg)
	paginator := route53resolver.NewListResolverQueryLogConfigAssociationsPaginator(client, &route53resolver.ListResolverQueryLogConfigAssociationsInput{})

	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, v := range page.ResolverQueryLogConfigAssociations {
			resource := Resource{
				Region:      describeCtx.Region,
				ID:          *v.Id,
				Name:        *v.Id,
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

func Route53ResolverResolverRule(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := route53resolver.NewFromConfig(cfg)
	paginator := route53resolver.NewListResolverRulesPaginator(client, &route53resolver.ListResolverRulesInput{})

	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, v := range page.ResolverRules {
			defaultID := "rslvr-autodefined-rr-internet-resolver"

			var tags []resolvertypes.Tag
			if *v.Id != defaultID {
				tagsOut, err := client.ListTagsForResource(ctx, &route53resolver.ListTagsForResourceInput{
					ResourceArn: v.Arn,
				})
				if err != nil {
					return nil, err
				}
				tags = tagsOut.Tags
			}

			// Build the params
			params := &route53resolver.ListResolverRuleAssociationsInput{
				Filters: []resolvertypes.Filter{
					{
						Name: aws.String("ResolverRuleId"),
						Values: []string{
							*v.Id,
						},
					},
				},
			}

			ruleass, err := client.ListResolverRuleAssociations(ctx, params)
			if err != nil {
				return nil, err
			}

			resource := Resource{
				Region: describeCtx.Region,
				ARN:    *v.Arn,
				Name:   *v.Name,
				Description: model.Route53ResolverResolverRuleDescription{
					ResolverRole:     v,
					Tags:             tags,
					RuleAssociations: ruleass,
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

func Route53ResolverResolverEndpoint(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := route53resolver.NewFromConfig(cfg)
	paginator := route53resolver.NewListResolverEndpointsPaginator(client, &route53resolver.ListResolverEndpointsInput{})

	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, resolverEndpoint := range page.ResolverEndpoints {
			ipAddresesses, err := client.ListResolverEndpointIpAddresses(ctx, &route53resolver.ListResolverEndpointIpAddressesInput{
				ResolverEndpointId: resolverEndpoint.Id,
			})
			if err != nil {
				ipAddresesses = &route53resolver.ListResolverEndpointIpAddressesOutput{}
			}

			tags, err := client.ListTagsForResource(ctx, &route53resolver.ListTagsForResourceInput{
				ResourceArn: resolverEndpoint.Arn,
			})
			if err != nil {
				tags = &route53resolver.ListTagsForResourceOutput{}
			}

			resource := Resource{
				Region: describeCtx.KaytuRegion,
				ARN:    *resolverEndpoint.Arn,
				Name:   *resolverEndpoint.Name,
				ID:     *resolverEndpoint.Id,
				Description: model.Route53ResolverEndpointDescription{
					ResolverEndpoint: resolverEndpoint,
					IpAddresses:      ipAddresesses.IpAddresses,
					Tags:             tags.Tags,
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

func Route53ResolverResolverRuleAssociation(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := route53resolver.NewFromConfig(cfg)
	paginator := route53resolver.NewListResolverRuleAssociationsPaginator(client, &route53resolver.ListResolverRuleAssociationsInput{})

	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, v := range page.ResolverRuleAssociations {
			resource := Resource{
				Region:      describeCtx.Region,
				ID:          *v.Id,
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
	}

	return values, nil
}

func Route53Domain(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)

	client := route53domains.NewFromConfig(cfg)

	paginator := route53domains.NewListDomainsPaginator(client, &route53domains.ListDomainsInput{})
	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, v := range page.Domains {
			domain, err := client.GetDomainDetail(ctx, &route53domains.GetDomainDetailInput{
				DomainName: v.DomainName,
			})
			if err != nil {
				return nil, err
			}

			tags, err := client.ListTagsForDomain(ctx, &route53domains.ListTagsForDomainInput{
				DomainName: v.DomainName,
			})
			if err != nil {
				tags = &route53domains.ListTagsForDomainOutput{}
			}

			arn := fmt.Sprintf("arn:%s:route53domains:::domain/%s", describeCtx.Partition, *v.DomainName)
			resource := Resource{
				Region: describeCtx.Region,
				Name:   *domain.DomainName,
				ARN:    arn,
				Description: model.Route53DomainDescription{
					DomainSummary: v,
					Domain:        *domain,
					Tags:          tags.TagList,
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

func Route53Record(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)

	client := route53.NewFromConfig(cfg)
	paginator := route53.NewListHostedZonesPaginator(client, &route53.ListHostedZonesInput{})
	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, v := range page.HostedZones {
			err := PaginateRetrieveAll(func(prevToken *string) (nextToken *string, err error) {
				records, err := client.ListResourceRecordSets(ctx, &route53.ListResourceRecordSetsInput{
					HostedZoneId:    v.Id,
					StartRecordName: prevToken,
				})
				if err != nil {
					return nil, err
				}
				for _, record := range records.ResourceRecordSets {
					arn := fmt.Sprintf("arn:%s:route53:::hostedzone/%s/recordset/%s/%s", describeCtx.Partition, *v.Id, *record.Name, record.Type)
					resource := Resource{
						Region: describeCtx.Region,
						Name:   *record.Name,
						ARN:    arn,
						Description: model.Route53RecordDescription{
							ZoneID: *v.Id,
							Record: record,
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
				if records.IsTruncated {
					return records.NextRecordName, nil
				}
				return nil, nil
			})
			if err != nil {
				return nil, err
			}
		}
	}

	return values, nil
}

func Route53TrafficPolicy(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)

	client := route53.NewFromConfig(cfg)
	var values []Resource
	err := PaginateRetrieveAll(func(prevToken *string) (nextToken *string, err error) {
		policies, err := client.ListTrafficPolicies(ctx, &route53.ListTrafficPoliciesInput{
			TrafficPolicyIdMarker: prevToken,
		})
		if err != nil {
			return nil, err
		}
		for _, policySummary := range policies.TrafficPolicySummaries {
			policy, err := client.GetTrafficPolicy(ctx, &route53.GetTrafficPolicyInput{
				Id: policySummary.Id,
			})
			if err != nil {
				return nil, err
			}

			arn := fmt.Sprintf("arn:%s:route53::%s:trafficpolicy/%s/%s", describeCtx.Partition, describeCtx.AccountID, *policy.TrafficPolicy.Id, string(*policy.TrafficPolicy.Version))
			resource := Resource{
				Region: describeCtx.Region,
				Name:   *policy.TrafficPolicy.Name,
				ID:     *policy.TrafficPolicy.Id,
				ARN:    arn,
				Description: model.Route53TrafficPolicyDescription{
					TrafficPolicy: *policy.TrafficPolicy,
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
		if policies.IsTruncated {
			return policies.TrafficPolicyIdMarker, nil
		}
		return nil, nil
	})
	if err != nil {
		return nil, err
	}

	return values, nil
}

func Route53TrafficPolicyInstance(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)

	client := route53.NewFromConfig(cfg)
	var values []Resource
	err := PaginateRetrieveAll(func(prevToken *string) (nextToken *string, err error) {
		policies, err := client.ListTrafficPolicyInstances(ctx, &route53.ListTrafficPolicyInstancesInput{
			TrafficPolicyInstanceNameMarker: prevToken,
		})
		if err != nil {
			return nil, err
		}
		for _, policyInstance := range policies.TrafficPolicyInstances {
			arn := fmt.Sprintf("arn:%s:route53::%s:trafficpolicyinstance/%s", describeCtx.Partition, describeCtx.AccountID, *policyInstance.Id)
			resource := Resource{
				Region: describeCtx.Region,
				Name:   *policyInstance.Name,
				ID:     *policyInstance.Id,
				ARN:    arn,
				Description: model.Route53TrafficPolicyInstanceDescription{
					TrafficPolicyInstance: policyInstance,
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
		if policies.IsTruncated {
			return policies.TrafficPolicyInstanceNameMarker, nil
		}
		return nil, nil
	})
	if err != nil {
		return nil, err
	}

	return values, nil
}
