package aws

import (
	"context"

	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAwsElasticLoadBalancingV2TargetGroup(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_ec2_load_balancer_target_group",
		Description: "AWS ElasticLoadBalancingV2 TargetGroup",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("target_group_name"),
			Hydrate:    kaytu.GetElasticLoadBalancingV2TargetGroup,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListElasticLoadBalancingV2TargetGroup,
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "target_group_name",
				Description: "The name of the target group.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.TargetGroup.TargetGroupName")},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the target group",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.TargetGroup.TargetGroupArn")},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.TargetGroup.TargetGroupName")},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.TargetGroup.TargetGroupArn").Transform(arnToAkas),
			},
		}),
	}
}
