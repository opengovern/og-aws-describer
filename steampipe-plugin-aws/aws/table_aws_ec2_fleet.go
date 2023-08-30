package aws

import (
	"context"
	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAwsEc2Fleet(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_ec2_fleet",
		Description: "AWS Ec2 Fleet",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("fleet_id"),
			Hydrate:    kaytu.GetEC2Fleet,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListEC2Fleet,
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "fleet_id",
				Description: "The id of the fleet.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Fleet.FleetId")},
			{
				Name:        "fleet_arn",
				Description: "The Amazon Resource Name (ARN) of the fleet",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("ARN")},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Fleet.FleetId")},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(getEc2FleetTurbotTags), // probably needs a transform function like Transform({{service}}FleetTagListToTurbotTags)
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

func getEc2FleetTurbotTags(_ context.Context, d *transform.TransformData) (interface{}, error) {
	tags := d.HydrateItem.(kaytu.EC2Fleet).Description.Fleet.Tags
	return ec2V2TagsToMap(tags)
}
