package aws

import (
	"context"

	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAwsElasticLoadBalancingV2Rule(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_ec2_load_balancer_listener_rule",
		Description: "AWS ElasticLoadBalancingV2 Rule",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("arn"),
			Hydrate:    opengovernance.GetElasticLoadBalancingV2Rule,
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListElasticLoadBalancingV2Rule,
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the rule",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Rule.RuleArn")},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Rule.RuleArn")},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Rule.RuleArn").Transform(arnToAkas),
			},
		}),
	}
}
