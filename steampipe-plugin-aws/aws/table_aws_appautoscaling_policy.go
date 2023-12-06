package aws

import (
	"context"
	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsAppAutoScalingPolicy(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_appautoscaling_policy",
		Description: "AWS Application Auto Scaling Policy",
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListApplicationAutoScalingPolicy,
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:    "service_namespace",
					Require: plugin.Required,
				},
				{
					Name:    "resource_id",
					Require: plugin.Optional,
				},
				{
					Name:    "policy_name",
					Require: plugin.Optional,
				},
			},
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "policy_arn",
				Description: "The Amazon Resource Name (ARN) of the appautoscaling policy.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ScalablePolicy.PolicyARN"),
			},
			{
				Name:        "policy_name",
				Description: "The name of the scaling policy.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ScalablePolicy.PolicyARN"),
			},
			{
				Name:        "service_namespace",
				Description: "The namespace of the AWS service that provides the resource, or a custom-resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ScalablePolicy.PolicyName"),
			},
			{
				Name:        "resource_id",
				Description: "The identifier of the resource associated with the scaling policy.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ScalablePolicy.ResourceId"),
			},
			{
				Name:        "scalable_dimension",
				Description: "The scalable dimension associated with the scaling policy. This string consists of the service namespace, resource type, and scaling property.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ScalablePolicy.ScalableDimension"),
			},
			{
				Name:        "policy_type",
				Description: "The policy type. Currently supported values are TargetTrackingScaling and StepScaling",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ScalablePolicy.PolicyType"),
			},
			{
				Name:        "target_tracking_scaling_policy_configuration",
				Description: "The target tracking scaling policy configuration (if policy type is TargetTrackingScaling).",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ScalablePolicy.TargetTrackingScalingPolicyConfiguration"),
			},
			{
				Name:        "step_scaling_policy_configuration",
				Description: "The step tracking scaling policy configuration (if policy type is StepScaling).",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ScalablePolicy.StepScalingPolicyConfiguration"),
			},
			{
				Name:        "alarms",
				Description: "The CloudWatch alarms associated with the scaling policy.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ScalablePolicy.Alarms"),
			},
			{
				Name:        "creation_time",
				Description: "The Unix timestamp for when the scaling policy was created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.ScalablePolicy.CreationTime"),
			},

			// Standard columns for all tables
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ScalablePolicy.ResourceId"),
			},
		}),
	}
}
