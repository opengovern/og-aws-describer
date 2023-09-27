package aws

import (
	"context"
	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsAppAutoScalingTarget(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_appautoscaling_target",
		Description: "AWS Application Auto Scaling Target",
		Get: &plugin.GetConfig{
			//KeyColumns: plugin.AllColumns([]string{"service_namespace", "resource_id"}),
			Hydrate: kaytu.GetApplicationAutoScalingTarget,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListApplicationAutoScalingTarget,
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
					Name:    "scalable_dimension",
					Require: plugin.Optional,
				},
			},
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "service_namespace",
				Description: "The namespace of the AWS service that provides the resource, or a custom-resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ScalableTarget.ServiceNamespace"),
			},
			{
				Name:        "resource_id",
				Description: "The identifier of the resource associated with the scalable target.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ScalableTarget.ResourceId"),
			},
			{
				Name:        "scalable_dimension",
				Description: "The scalable dimension associated with the scalable target. This string consists of the service namespace, resource type, and scaling property.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ScalableTarget.ScalableDimension"),
			},
			{
				Name:        "creation_time",
				Description: "The Unix timestamp for when the scalable target was created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.ScalableTarget.CreationTime"),
			},
			{
				Name:        "min_capacity",
				Description: "The minimum value to scale to in response to a scale-in activity.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.ScalableTarget.MinCapacity"),
			},
			{
				Name:        "max_capacity",
				Description: "The maximum value to scale to in response to a scale-out activity.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.ScalableTarget.MaxCapacity"),
			},
			{
				Name:        "role_arn",
				Description: "The ARN of an IAM role that allows Application Auto Scaling to modify the scalable target on your behalf.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ScalableTarget.RoleARN"),
			},
			{
				Name:        "suspended_state",
				Description: "Specifies whether the scaling activities for a scalable target are in a suspended state.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ScalableTarget.SuspendedState"),
			},

			// Standard columns for all tables
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ScalableTarget.ResourceId"),
			},
		}),
	}
}

//// LIST FUNCTION

//// HYDRATE FUNCTIONS
