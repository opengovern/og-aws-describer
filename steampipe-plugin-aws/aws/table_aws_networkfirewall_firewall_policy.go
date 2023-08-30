package aws

import (
	"context"

	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsNetworkFirewallPolicy(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_networkfirewall_firewall_policy",
		Description: "AWS Network Firewall Policy",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AnyColumn([]string{"arn", "name"}),
			Hydrate:    kaytu.GetNetworkFirewallFirewallPolicy,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListNetworkFirewallFirewallPolicy,
		},

		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The descriptive name of the rule group.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.FirewallPolicyResponse.FirewallPolicyName")},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the rule group.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Arn", "FirewallPolicyResponse.FirewallPolicyArn"),
			},
			{
				Name:        "firewall_policy_id",
				Description: "The unique identifier for the firewall policy.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.FirewallPolicyResponse.FirewallPolicyId")},
			{
				Name:        "consumed_stateful_rule_capacity",
				Description: "The number of capacity units currently consumed by the policy's stateful rules.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.FirewallPolicyResponse.ConsumedStatefulRuleCapacity")},
			{
				Name:        "consumed_stateless_rule_capacity",
				Description: "The number of capacity units currently consumed by the policy's stateless rules.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.FirewallPolicyResponse.ConsumedStatelessRuleCapacity")},
			{
				Name:        "description",
				Description: "A description of the firewall policy.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.FirewallPolicyResponse.Description")},
			{
				Name:        "firewall_policy_status",
				Description: "The current status of the firewall policy.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.FirewallPolicyResponse.FirewallPolicyStatus")},
			{
				Name:        "last_modified_time",
				Description: "The last time that the firewall policy was changed.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.FirewallPolicyResponse.LastModifiedTime")},
			{
				Name:        "number_of_associations",
				Description: "The number of firewall policies that use this rule group.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.FirewallPolicyResponse.NumberOfAssociations")},

			{
				Name:        "encryption_configuration",
				Description: "A complex type that contains the Amazon Web Services KMS encryption configuration settings for your firewall policy.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.FirewallPolicyResponse.EncryptionConfiguration")},
			{
				Name:        "firewall_policy",
				Description: "The policy for the specified firewall policy.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.FirewallPolicy")},
			{
				Name:        "tags_src",
				Description: "A list of tags assigned to the resource.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.FirewallPolicyResponse.Tags")},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.FirewallPolicyResponse.FirewallPolicyName")},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.FirewallPolicyResponse.Tags").Transform(networkFirewallPolicyTurbotTags),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.FirewallPolicyResponse.Tags").Transform(transform.EnsureStringArray),
			},
		}),
	}
}

func networkFirewallPolicyTurbotTags(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	policy := d.HydrateItem.(kaytu.NetworkFirewallFirewallPolicy).Description

	// Mapping the resource tags inside turbotTags
	var turbotTagsMap map[string]string
	if policy.FirewallPolicyResponse.Tags != nil {
		turbotTagsMap = map[string]string{}
		for _, i := range policy.FirewallPolicyResponse.Tags {
			turbotTagsMap[*i.Key] = *i.Value
		}
	}

	return turbotTagsMap, nil
}
