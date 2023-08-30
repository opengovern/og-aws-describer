package aws

import (
	"context"

	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsEc2TargetGroup(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_ec2_target_group",
		Description: "AWS EC2 Target Group",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("target_group_arn"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"LoadBalancerNotFound", "TargetGroupNotFound", "ValidationError"}),
			},
			Hydrate: kaytu.GetElasticLoadBalancingV2TargetGroup,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListElasticLoadBalancingV2TargetGroup,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"TargetGroupNotFound", "ValidationError"}),
			},
			KeyColumns: []*plugin.KeyColumn{
				{Name: "target_group_name", Require: plugin.Optional},
			},
		},

		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "target_group_name",
				Description: "The name of the target group.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.TargetGroup.TargetGroupName")},
			{
				Name:        "target_group_arn",
				Description: "The Amazon Resource Name (ARN) of the target group.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.TargetGroup.TargetGroupArn")},
			{
				Name:        "target_type",
				Description: "The type of target that is specified when registering targets with this target group. The possible values are instance (register targets by instance ID), ip (register targets by IP address), or lambda (register a single Lambda function as a target).",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.TargetGroup.TargetType")},
			{
				Name:        "load_balancer_arns",
				Description: "The Amazon Resource Names (ARN) of the load balancers that route traffic to this target group.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.TargetGroup.LoadBalancerArns")},
			{
				Name:        "port",
				Description: "The port on which the targets are listening. Not used if the target is a Lambda function.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.TargetGroup.Port")},
			{
				Name:        "vpc_id",
				Description: "The ID of the VPC for the target.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.TargetGroup.VpcId")},
			{
				Name:        "protocol",
				Description: "The protocol to use for routing traffic to the target.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.TargetGroup.Protocol")},
			{
				Name:        "matcher_http_code",
				Description: "The HTTP codes to use when checking for a successful response from a target.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.TargetGroup.Matcher.HttpCode")},
			{
				Name:        "matcher_grpc_code",
				Description: "The gRPC codes to use when checking for a successful response from a target.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.TargetGroup.Matcher.GrpcCode")},
			{
				Name:        "healthy_threshold_count",
				Description: "The number of consecutive health checks successes required before considering an unhealthy target healthy.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.TargetGroup.HealthyThresholdCount")},
			{
				Name:        "unhealthy_threshold_count",
				Description: "The number of consecutive health checks successes required before considering an unhealthy target healthy.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.TargetGroup.UnhealthyThresholdCount")},
			{
				Name:        "health_check_enabled",
				Description: "Indicates whether health checks are enabled.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.TargetGroup.HealthCheckEnabled")},
			{
				Name:        "health_check_interval_seconds",
				Description: "The approximate amount of time, in seconds, between health checks of an individual target.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.TargetGroup.HealthCheckIntervalSeconds")},
			{
				Name:        "health_check_path",
				Description: "The destination for health checks on the target.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.TargetGroup.HealthCheckPath")},
			{
				Name:        "health_check_port",
				Description: "The port to use to connect with the target.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.TargetGroup.HealthCheckPort")},
			{
				Name:        "health_check_protocol",
				Description: "The protocol to use to connect with the target. The GENEVE, TLS, UDP, and TCP_UDP protocols are not supported for health checks.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.TargetGroup.HealthCheckProtocol")},
			{
				Name:        "health_check_timeout_seconds",
				Description: "The amount of time, in seconds, during which no response means a failed health check.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.TargetGroup.HealthCheckTimeoutSeconds")},
			{
				Name:        "target_health_descriptions",
				Description: "Contains information about the health of the target.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Health")},
			{
				Name:        "tags_src",
				Description: "A list of tags associated with target group.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Tags")},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.TargetGroup.TargetGroupName")},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(targetGroupTagsToTurbotTags),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.TargetGroup.TargetGroupArn").Transform(arnToAkas),
			},
		}),
	}
}

//// TRANSFORM FUNCTIONS

func targetGroupTagsToTurbotTags(_ context.Context, d *transform.TransformData) (interface{}, error) {
	data := d.HydrateItem.(kaytu.ElasticLoadBalancingV2TargetGroup).Description
	var turbotTagsMap map[string]string
	if data.Tags != nil {
		turbotTagsMap = map[string]string{}
		for _, i := range data.Tags {
			turbotTagsMap[*i.Key] = *i.Value
		}
	}

	return turbotTagsMap, nil
}
