package aws

import (
	"context"
	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAwsEc2PlacementGroup(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_ec2_placement_group",
		Description: "AWS Ec2 Placement Group",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("group_name"),
			Hydrate:    kaytu.GetEC2PlacementGroup,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListEC2PlacementGroup,
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "group_id",
				Description: "The id of the placement group.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.PlacementGroup.GroupId")},
			{
				Name:        "group_name",
				Description: "The name of the placement group.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.PlacementGroup.GroupName")},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the placement group",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("ARN")},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.PlacementGroup.GroupName")},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(getEc2PlacementGroupTurbotTags)},
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

func getEc2PlacementGroupTurbotTags(_ context.Context, d *transform.TransformData) (interface{}, error) {
	tags := d.HydrateItem.(kaytu.EC2PlacementGroup).Description.PlacementGroup.Tags
	return ec2V2TagsToMap(tags)
}
