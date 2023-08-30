package aws

import (
	"context"

	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsEksCluster(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_eks_cluster",
		Description: "AWS Elastic Kubernetes Service Cluster",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("name"),
			Hydrate:    kaytu.GetEKSCluster,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListEKSCluster,
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the cluster.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Cluster.Name"),
			},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the cluster.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Cluster.Arn"),
			},
			{
				Name:        "created_at",
				Description: "The Unix epoch timestamp in seconds for when the cluster was created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Cluster.CreatedAt"),
			},
			{
				Name:        "version",
				Description: "The Kubernetes server version for the cluster.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Cluster.Version"),
			},
			{
				Name:        "endpoint",
				Description: "The endpoint for your Kubernetes API server.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Cluster.Endpoint"),
			},
			{
				Name:        "role_arn",
				Description: "The Amazon Resource Name (ARN) of the IAM role that provides permissions for the Kubernetes control plane to make calls to AWS API operations on your behalf.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Cluster.RoleArn"),
			},
			{
				Name:        "encryption_config",
				Description: "The encryption configuration for the cluster.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Cluster.EncryptionConfig"),
			},
			{
				Name:        "resources_vpc_config",
				Description: "The VPC configuration used by the cluster control plane.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Cluster.ResourcesVpcConfig"),
			},
			{
				Name:        "kubernetes_network_config",
				Description: "The Kubernetes network configuration for the cluster.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Cluster.KubernetesNetworkConfig"),
			},
			{
				Name:        "logging",
				Description: "The logging configuration for the cluster.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Cluster.Logging"),
			},
			{
				Name:        "identity",
				Description: "The identity provider information for the cluster.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Cluster.Identity"),
			},
			{
				Name:        "status",
				Description: "The current status of the cluster.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Cluster.Status"),
			},
			{
				Name:        "certificate_authority",
				Description: "The certificate-authority-data for the cluster.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Cluster.CertificateAuthority"),
			},
			{
				Name:        "platform_version",
				Description: "The platform version of your Amazon EKS cluster.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Cluster.PlatformVersion"),
			},
			{
				Name:        "tags",
				Description: "A list of tags assigned to the table",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Cluster.Tags"),
			},

			// Standard columns for all tables
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Cluster.Name"),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Cluster.Arn").Transform(arnToAkas),
			},
		}),
	}
}

//// LIST FUNCTION

//// HYDRATE FUNCTIONS
