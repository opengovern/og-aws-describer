package aws

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/route53resolver/types"

	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsRoute53ResolverRule(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_route53_resolver_rule",
		Description: "AWS Route53 Resolver Rule",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"ResourceNotFoundException"}),
			},
			Hydrate: kaytu.GetRoute53ResolverResolverRule,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListRoute53ResolverResolverRule,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "creator_request_id", Require: plugin.Optional},
				{Name: "domain_name", Require: plugin.Optional},
				{Name: "name", Require: plugin.Optional},
				{Name: "resolver_endpoint_id", Require: plugin.Optional},
				{Name: "status", Require: plugin.Optional},
			},
		},

		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The name for the Resolver rule, which you specified when you created the Resolver rule.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ResolverRole.Name")},
			{
				Name:        "id",
				Description: "The ID that Resolver assigned to the Resolver rule when you created it.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ResolverRole.Id")},
			{
				Name:        "arn",
				Description: "The ARN (Amazon Resource Name) for the Resolver rule specified by Id.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ResolverRole.Arn")},
			{
				Name:        "status",
				Description: "A code that specifies the current status of the Resolver rule.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ResolverRole.Status")},
			{
				Name:        "creator_request_id",
				Description: "A unique string that you specified when you created the Resolver rule. CreatorRequestId identifies the request and allows failed requests to be retried without the risk of executing the operation twice.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ResolverRole.CreatorRequestId")},
			{
				Name:        "domain_name",
				Description: "DNS queries for this domain name are forwarded to the IP addresses that are specified in TargetIps.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ResolverRole.DomainName")},
			{
				Name:        "owner_id",
				Description: "When a rule is shared with another AWS account, the account ID of the account that the rule is shared with.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ResolverRole.OwnerId")},
			{
				Name:        "resolver_endpoint_id",
				Description: "The ID of the endpoint that the rule is associated with.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ResolverRole.ResolverEndpointId")},
			{
				Name:        "rule_type",
				Description: "When you want to forward DNS queries for specified domain name to resolvers on your network, specify FORWARD.When you have a forwarding rule to forward DNS queries for a domain to your network and you want Resolver to process queries for a subdomain of that domain, specify SYSTEM.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ResolverRole.RuleType")},
			{
				Name:        "share_status",
				Description: "Indicates whether the rules is shared and, if so, whether the current account is sharing the rule with another account, or another account is sharing the rule with the current account.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ResolverRole.ShareStatus")},
			{
				Name:        "status_message",
				Description: "A detailed description of the status of a Resolver rule.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ResolverRole.StatusMessage")},
			{
				Name:        "creation_time",
				Description: "The date and time that the Resolver rule was created, in Unix time format and Coordinated Universal Time (UTC).",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ResolverRole.CreationTime")},
			{
				Name:        "modification_time",
				Description: "The date and time that the Resolver rule was last updated, in Unix time format and Coordinated Universal Time (UTC).",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ResolverRole.ModificationTime")},
			{
				Name:        "resolver_rule_associations",
				Description: "The associations that were created between Resolver rules and VPCs using the current AWS account, and that match the specified filters, if any.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.RuleAssociations.ResolverRuleAssociations")},
			{
				Name:        "target_ips",
				Description: "An array that contains the IP addresses and ports that an outbound endpoint forwards DNS queries to. Typically, these are the IP addresses of DNS resolvers on your network. Specify IPv4 addresses. IPv6 is not supported.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ResolverRole.TargetIps")},
			{
				Name:        "tags_src",
				Description: "A list of tags assigned to the Resolver Rule.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Tags")},

			// Standard columns for all tables
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ResolverRole.Name")},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Tags").Transform(route53resolverRuleTagListToTurbotTags),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ResolverRole.Arn")},

			//// LIST FUNCTION
		}),
	}
}

//// TRANSFORM FUNCTIONS

func route53resolverRuleTagListToTurbotTags(ctx context.Context, d *transform.TransformData) (interface{}, error) {

	if d.Value == nil {
		return nil, nil
	}
	tagList := d.Value.([]types.Tag)

	// Mapping the resource tags inside turbotTags
	var turbotTagsMap map[string]string
	if tagList != nil {
		turbotTagsMap = map[string]string{}
		for _, i := range tagList {
			turbotTagsMap[*i.Key] = *i.Value
		}
	} else {
		return nil, nil
	}

	return turbotTagsMap, nil
}

//// UTILITY FUNCTION
