package aws

import (
	"context"

	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAwsEksNodegroup(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_eks_nodegroup",
		Description: "AWS EKS Nodegroup",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("nodegroup_name"),
			Hydrate:    kaytu.GetEKSNodegroup,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListEKSNodegroup,
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "nodegroup_name",
				Description: "The name of the nodegroup.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Nodegroup.NodegroupName")},
			{
				Name:        "nodegroup_arn",
				Description: "The Amazon Resource Name (ARN) of the nodegroup",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Nodegroup.NodegroupArn")},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Nodegroup.NodegroupName")},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Nodegroup.Tags"),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Nodegroup.NodegroupArn").Transform(arnToAkas),
			},
		}),
	}
}
