package aws

import (
	"context"

	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsRDSDBSubnetGroup(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_rds_db_subnet_group",
		Description: "AWS RDS DB Subnet Group",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("name"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"DBSubnetGroupNotFoundFault"}),
			},
			Hydrate: opengovernance.GetRDSDBSubnetGroup,
		},

		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListRDSDBSubnetGroup,
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The friendly name to identify the DB subnet group.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DBSubnetGroup.DBSubnetGroupName")},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) for the DB subnet group.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DBSubnetGroup.DBSubnetGroupArn"),
			},
			{
				Name:        "description",
				Description: "Provides the description of the DB subnet group.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DBSubnetGroup.DBSubnetGroupDescription")},
			{
				Name:        "status",
				Description: "Provides the status of the DB subnet group.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DBSubnetGroup.SubnetGroupStatus")},
			{
				Name:        "vpc_id",
				Description: "Provides the VpcId of the DB subnet group.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DBSubnetGroup.VpcId")},
			{
				Name:        "subnets",
				Description: "A list of Subnet elements.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DBSubnetGroup.Subnets")},
			{
				Name:        "tags_src",
				Description: "A list of tags attached to the DB subnet group.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Tags.TagList")},

			// Standard columns
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(getRDSDBSubnetGroupTurbotTags),
			},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DBSubnetGroup.DBSubnetGroupName")},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform: transform.

					//// LIST FUNCTION
					FromField("Description.DBSubnetGroup.DBSubnetGroupArn").Transform(arnToAkas),
			},
		}),
	}
}

//// TRANSFORM FUNCTIONS

func getRDSDBSubnetGroupTurbotTags(_ context.Context, d *transform.TransformData) (interface{}, error) {
	dbSubnetGroup := d.HydrateItem.(opengovernance.RDSDBSubnetGroup).Description.Tags
	if dbSubnetGroup == nil {
		return nil, nil
	}
	if dbSubnetGroup.TagList != nil {
		turbotTagsMap := map[string]string{}
		for _, i := range dbSubnetGroup.TagList {
			turbotTagsMap[*i.Key] = *i.Value
		}
		return turbotTagsMap, nil
	}
	return nil, nil
}
