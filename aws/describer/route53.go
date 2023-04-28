package describer

import (
	"context"
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/route53"
	"github.com/aws/aws-sdk-go-v2/service/route53/types"
	"github.com/aws/aws-sdk-go-v2/service/route53resolver"
	"github.com/kaytu-io/kaytu-aws-describer/aws/model"
)

func Route53HealthCheck(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	client := route53.NewFromConfig(cfg)

	var values []Resource
	err := PaginateRetrieveAll(func(prevToken *string) (nextToken *string, err error) {
		output, err := client.ListHealthChecks(ctx, &route53.ListHealthChecksInput{Marker: prevToken})
		if err != nil {
			return nil, err
		}

		for _, v := range output.HealthChecks {
			resource := Resource{
				ID:          *v.Id,
				Name:        *v.HealthCheckConfig.FullyQualifiedDomainName,
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
				ARN:  arn,
				Name: *v.Name,
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
		ARN:  arn,
		Name: *v.Name,
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

func Route53ResolverResolverEndpoint(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	client := route53resolver.NewFromConfig(cfg)
	paginator := route53resolver.NewListResolverEndpointsPaginator(client, &route53resolver.ListResolverEndpointsInput{})

	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, v := range page.ResolverEndpoints {
			resource := Resource{
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

func Route53ResolverResolverQueryLoggingConfig(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
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
	client := route53resolver.NewFromConfig(cfg)
	paginator := route53resolver.NewListResolverRulesPaginator(client, &route53resolver.ListResolverRulesInput{})

	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, v := range page.ResolverRules {
			resource := Resource{
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

func Route53ResolverResolverRuleAssociation(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
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
