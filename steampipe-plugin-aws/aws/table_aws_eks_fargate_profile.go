package aws

import (
	"context"

	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsEksFargateProfile(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_eks_fargate_profile",
		Description: "AWS Elastic Kubernetes Service Fargate Profile",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"cluster_name", "fargate_profile_name"}),
			Hydrate:    kaytu.GetEKSFargateProfile,
		},
		List: &plugin.ListConfig{
			ParentHydrate: kaytu.ListEKSCluster,
			Hydrate:       kaytu.ListEKSFargateProfile,
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:    "cluster_name",
					Require: plugin.Optional,
				},
			},
		},

		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "fargate_profile_name",
				Description: "The name of the Fargate profile.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.FargateProfile.FargateProfileName")},
			{
				Name:        "cluster_name",
				Description: "The name of the Amazon EKS cluster that the Fargate profile belongs to.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.FargateProfile.ClusterName")},
			{
				Name:        "fargate_profile_arn",
				Description: "The full Amazon Resource Name (ARN) of the Fargate profile.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.FargateProfile.FargateProfileArn")},
			{
				Name:        "created_at",
				Description: "The Unix epoch timestamp in seconds for when the Fargate profile was created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.FargateProfile.CreatedAt")},
			{
				Name:        "pod_execution_role_arn",
				Description: "The Amazon Resource Name (ARN) of the pod execution role to use for pods that match the selectors in the Fargate profile.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.FargateProfile.PodExecutionRoleArn")},
			{
				Name:        "status",
				Description: "The current status of the Fargate profile.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.FargateProfile.Status")},
			{
				Name:        "selectors",
				Description: "The selectors to match for pods to use this Fargate profile.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.FargateProfile.Selectors")},
			{
				Name:        "subnets",
				Description: "The subnets used by the Fargate profile.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.FargateProfile.Subnets")},
			{
				Name:        "tags",
				Description: "A list of tags assigned to the Fargate profile.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.FargateProfile.Tags")},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.FargateProfile.FargateProfileName")},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.FargateProfile.FargateProfileArn").Transform(transform.EnsureStringArray),
			},
		}),
	}
}
