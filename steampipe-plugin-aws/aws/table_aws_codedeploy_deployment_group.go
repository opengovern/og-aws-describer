package aws

import (
	"context"
	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsCodeDeployDeploymentGroup(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_codedeploy_deployment_group",
		Description: "AWS CodeDeploy DeploymentGroup",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("deployment_group_name"),
			Hydrate:    kaytu.GetCodeDeployDeploymentGroup,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListCodeDeployDeploymentGroup,
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the deployment group",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("ARN")},
			{
				Name:        "application_name",
				Description: "The application name.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DeploymentGroup.ApplicationName")},
			{
				Name:        "deployment_config_name",
				Description: "The deployment configuration name.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DeploymentGroup.DeploymentConfigName")},
			{
				Name:        "deployment_group_id",
				Description: "The deployment group ID.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DeploymentGroup.DeploymentGroupId")},
			{
				Name:        "deployment_group_name",
				Description: "The name of the deployment group.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DeploymentGroup.DeploymentGroupName")},
			{
				Name:        "service_role_arn",
				Description: "A service role Amazon Resource Name (ARN) that grants CodeDeploy permission to make calls to Amazon Web Services services on your behalf.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DeploymentGroup.ServiceRoleArn")},
			{
				Name:        "alarm_configuration",
				Description: "A list of alarms associated with the deployment group.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DeploymentGroup.AlarmConfiguration")},
			{
				Name:        "auto_rollback_configuration",
				Description: "Information about the automatic rollback configuration associated with the deployment group.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DeploymentGroup.AutoRollbackConfiguration")},
			{
				Name:        "auto_scaling_groups",
				Description: "A list of associated Auto Scaling groups.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DeploymentGroup.AutoScalingGroups")},
			{
				Name:        "blue_green_deployment_configuration",
				Description: "Information about blue/green deployment options for a deployment group.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DeploymentGroup.BlueGreenDeploymentConfiguration")},
			{
				Name:        "compute_platform",
				Description: "The destination platform type for the deployment (Lambda, Server, or ECS).",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DeploymentGroup.ComputePlatform")},
			{
				Name:        "deployment_style",
				Description: "Information about the type of deployment, either in-place or blue/green, you want to run and whether to route deployment traffic behind a load balancer.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DeploymentGroup.DeploymentStyle")},
			{
				Name:        "ec2_tag_filters",
				Description: "The Amazon EC2 tags on which to filter. The deployment group includes EC2 instances with any of the specified tags.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DeploymentGroup.Ec2TagFilters")},
			{
				Name:        "ec2_tag_set",
				Description: "Information about groups of tags applied to an Amazon EC2 instance.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DeploymentGroup.Ec2TagSet")},
			{
				Name:        "ecs_services",
				Description: "The target Amazon ECS services in the deployment group.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DeploymentGroup.EcsServices")},
			{
				Name:        "last_attempted_deployment",
				Description: "Information about the most recent attempted deployment to the deployment group.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DeploymentGroup.LastAttemptedDeployment")},
			{
				Name:        "last_successful_deployment",
				Description: "Information about the most recent successful deployment to the deployment group.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DeploymentGroup.LastSuccessfulDeployment")},
			{
				Name:        "load_balancer_info",
				Description: "Information about the load balancer to use in a deployment.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DeploymentGroup.LoadBalancerInfo")},
			{
				Name:        "on_premises_instance_tag_filters",
				Description: "The on-premises instance tags on which to filter.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DeploymentGroup.OnPremisesInstanceTagFilters")},
			{
				Name:        "on_premises_tag_set",
				Description: "Information about groups of tags applied to an on-premises instance.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DeploymentGroup.OnPremisesTagSet")},
			{
				Name:        "outdated_instances_strategy",
				Description: "Indicates what happens when new Amazon EC2 instances are launched mid-deployment and do not receive the deployed application revision.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DeploymentGroup.OutdatedInstancesStrategy")},
			{
				Name:        "tags_src",
				Description: "A list of tags associated with deployment group.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Tags")},
			{
				Name:        "target_revision",
				Description: "Information about the deployment group's target revision, including type and location.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DeploymentGroup.TargetRevision")},
			{
				Name:        "trigger_configurations",
				Description: "Information about triggers associated with the deployment group.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DeploymentGroup.TriggerConfigurations")},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DeploymentGroup.DeploymentGroupName")},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(getCodeDeployDeploymentGroupTurbotTags),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("ARN").Transform(arnToAkas),
			},
		}),
	}
}

//// TRANSFORM FUNCTIONS

func getCodeDeployDeploymentGroupTurbotTags(_ context.Context, d *transform.TransformData) (interface{}, error) {
	tags := d.HydrateItem.(kaytu.CodeDeployDeploymentGroup).Description.Tags
	return codeDeployV2TagsToMap(tags)
}
