package aws

import (
	"context"
	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsCodeDeployDeploymentConfig(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_codedeploy_deployment_config",
		Description: "AWS CodeDeploy Deployment Config",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("deployment_config_name"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"DeploymentConfigDoesNotExistException"}),
			},
			Hydrate: kaytu.GetCodeDeployDeploymentConfig,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListCodeDeployDeploymentConfig,
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) specifying the application.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("ARN").Transform(arnToAkas),
			},
			{
				Name:        "create_time",
				Description: "The time at which the deployment configuration was created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Config.CreateTime"),
			},
			{
				Name:        "deployment_config_id",
				Description: "The deployment configuration ID.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Config.DeploymentConfigId"),
			},
			{
				Name:        "deployment_config_name",
				Description: "The deployment configuration name.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Config.DeploymentConfigName"),
			},
			{
				Name:        "compute_platform",
				Description: "The destination platform type for the deployment (Lambda, Server, or ECS).",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Config.ComputePlatform"),
			},
			{
				Name:        "minimum_healthy_hosts",
				Description: "Information about the number or percentage of minimum healthy instances.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Config.MinimumHealthyHosts"),
			},
			{
				Name:        "traffic_routing_config",
				Description: "The configuration that specifies how the deployment traffic is routed. Used for deployments with a Lambda or Amazon ECS compute platform only.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Config.TrafficRoutingConfig"),
			},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Config.DeploymentConfigName"),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("ARN").Transform(transform.EnsureStringArray),
			},
		}),
	}
}
