package aws

import (
	"context"

	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsEksAddon(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_eks_addon",
		Description: "AWS EKS Addon",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"addon_name", "cluster_name"}),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"ResourceNotFoundException", "InvalidParameterException", "InvalidParameter"}),
			},
			Hydrate: opengovernance.GetEKSAddon,
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListEKSAddon,
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "addon_name",
				Description: "The name of the add-on.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Addon.AddonName"),
			},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the add-on.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Addon.AddonArn"),
			},
			{
				Name:        "cluster_name",
				Description: "The name of the cluster.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Addon.ClusterName"),
			},
			{
				Name:        "addon_version",
				Description: "The version of the add-on.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Addon.AddonVersion"),
			},
			{
				Name:        "status",
				Description: "The status of the add-on.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Addon.Status"),
			},
			{
				Name:        "created_at",
				Description: "The date and time that the add-on was created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Addon.CreatedAt"),
			},
			{
				Name:        "modified_at",
				Description: "The date and time that the add-on was last modified.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Addon.ModifiedAt"),
			},
			{
				Name:        "service_account_role_arn",
				Description: "The Amazon Resource Name (ARN) of the IAM role that is bound to the Kubernetes service account used by the add-on.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Addon.ServiceAccountRoleArn"),
			},
			{
				Name:        "health_issues",
				Description: "An object that represents the add-on's health issues.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Addon.Health.Issues"),
			},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Addon.AddonName"),
			},
			{
				Name:        "tags",
				Description: "The metadata that you apply to the cluster to assist with categorization and organization.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Addon.Tags"),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Addon.AddonArn").Transform(transform.EnsureStringArray),
			},
		}),
	}
}

//// LIST FUNCTION

//// HYDRATE FUNCTIONS
