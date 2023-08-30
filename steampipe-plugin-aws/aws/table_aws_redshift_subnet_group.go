package aws

import (
	"context"

	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsRedshiftSubnetGroup(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_redshift_subnet_group",
		Description: "AWS Redshift Subnet Group",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("cluster_subnet_group_name"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"ClusterSubnetGroupNotFoundFault"}),
			},
			Hydrate: kaytu.GetRedshiftSubnetGroup,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListRedshiftSubnetGroup,
		},

		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "cluster_subnet_group_name",
				Description: "The name of the cluster subnet group.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ClusterSubnetGroup.ClusterSubnetGroupName")},
			{
				Name:        "subnet_group_status",
				Description: "The status of the cluster subnet group.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ClusterSubnetGroup.SubnetGroupStatus")},
			{
				Name:        "description",
				Description: "The description of the cluster subnet group.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ClusterSubnetGroup.Description")},
			{
				Name:        "vpc_id",
				Description: "The VPC ID of the cluster subnet group.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ClusterSubnetGroup.VpcId")},
			{
				Name:        "subnets",
				Description: "A list of the VPC Subnet elements.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ClusterSubnetGroup.Subnets")},
			{
				Name:        "tags_src",
				Description: "A list of tags attached to the subnet group.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ClusterSubnetGroup.Tags")},

			// Standard columns
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ClusterSubnetGroup.ClusterSubnetGroupName")},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(redshiftSubnetGroupTurbotTags),
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

func redshiftSubnetGroupTurbotTags(_ context.Context, d *transform.TransformData) (interface{}, error) {
	clusterSubnetGroup := d.HydrateItem.(kaytu.RedshiftSubnetGroup).Description.ClusterSubnetGroup

	// Get the resource tags
	var turbotTagsMap map[string]string
	if len(clusterSubnetGroup.Tags) > 0 {
		turbotTagsMap = map[string]string{}
		for _, i := range clusterSubnetGroup.Tags {
			turbotTagsMap[*i.Key] = *i.Value
		}
	}
	return turbotTagsMap, nil
}
