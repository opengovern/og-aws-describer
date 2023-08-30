package aws

import (
	"context"

	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsEksNodeGroup(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_eks_node_group",
		Description: "AWS EKS Node Group",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"nodegroup_name", "cluster_name"}),
			Hydrate:    kaytu.GetEKSNodegroup,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"ResourceNotFoundException", "InvalidParameterException", "InvalidParameter"}),
			},
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListEKSNodegroup,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "cluster_name", Require: plugin.Optional},
			},
		},

		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "nodegroup_name",
				Description: "The name associated with an Amazon EKS managed node group.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Nodegroup.NodegroupName")},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) associated with the managed node group.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("NodegroupArn"),
			},
			{
				Name:        "cluster_name",
				Description: "The name of the cluster that the managed node group resides in.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Nodegroup.ClusterName")},
			{
				Name:        "created_at",
				Description: "The Unix epoch timestamp in seconds for when the managed node group was created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Nodegroup.CreatedAt")},
			{
				Name:        "status",
				Description: "The current status of the managed node group.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Nodegroup.Status")},
			{
				Name:        "ami_type",
				Description: "The AMI type that was specified in the node group configuration.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Nodegroup.AmiType")},
			{
				Name:        "capacity_type",
				Description: "The capacity type of your managed node group.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Nodegroup.CapacityType")},
			{
				Name:        "disk_size",
				Description: "The disk size in the node group configuration.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Nodegroup.DiskSize")},
			{
				Name:        "modified_at",
				Description: "The Unix epoch timestamp in seconds for when the managed node group was last modified.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Nodegroup.ModifiedAt")},
			{
				Name:        "node_role",
				Description: "The IAM role associated with your node group.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Nodegroup.NodeRole")},
			{
				Name:        "release_version",
				Description: "If the node group was deployed using a launch template with a custom AMI, then this is the AMI ID that was specified in the launch template. For node groups that weren't deployed using a launch template, this is the version of the Amazon EKS optimized AMI that the node group was deployed with.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Nodegroup.ReleaseVersion")},
			{
				Name:        "version",
				Description: "The Kubernetes version of the managed node group.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Nodegroup.Version")},
			{
				Name:        "health",
				Description: "The health status of the node group.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Nodegroup.Health")},
			{
				Name:        "instance_types",
				Description: "The instance type that is associated with the node group. If the node group was deployed with a launch template, then this is null.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Nodegroup.InstanceTypes")},
			{
				Name:        "labels",
				Description: "The Kubernetes labels applied to the nodes in the node group.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Nodegroup.Labels")},
			{
				Name:        "launch_template",
				Description: "If a launch template was used to create the node group, then this is the launch template that was used.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Nodegroup.LaunchTemplate")},
			{
				Name:        "remote_access",
				Description: "The remote access configuration that is associated with the node group. If the node group was deployed with a launch template, then this is null.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Nodegroup.RemoteAccess")},
			{
				Name:        "resources",
				Description: "The resources associated with the node group, such as Auto Scaling groups and security groups for remote access.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Nodegroup.Resources")},
			{
				Name:        "scaling_config",
				Description: "The scaling configuration details for the Auto Scaling group that is associated with your node group.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Nodegroup.ScalingConfig")},
			{
				Name:        "subnets",
				Description: "The subnets that were specified for the Auto Scaling group that is associated with your node group.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Nodegroup.Subnets")},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Nodegroup.Tags")},
			{
				Name:        "taints",
				Description: "The Kubernetes taints to be applied to the nodes in the node group when they are created.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Nodegroup.Taints")},
			{
				Name:        "update_config",
				Description: "The node group update configuration.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Nodegroup.UpdateConfig")},

			// Steampipe Standard columns
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Nodegroup.NodegroupName")},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Nodegroup.NodegroupArn").Transform(transform.EnsureStringArray),
			},
		}),
	}
}
