package aws

import (
	"context"
	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsVpcRouteTable(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_vpc_route_table",
		Description: "AWS VPC Route table",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("route_table_id"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"InvalidRouteTableID.NotFound", "InvalidRouteTableID.Malformed"}),
			},
			Hydrate: kaytu.GetEC2RouteTable,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListEC2RouteTable,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "owner_id", Require: plugin.Optional},
				{Name: "vpc_id", Require: plugin.Optional},
			},
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "route_table_id",
				Description: "Contains the ID of the route table.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.RouteTable.RouteTableId"),
			},
			{
				Name:        "vpc_id",
				Description: "The ID of the VPC.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.RouteTable.VpcId"),
			},
			{
				Name:        "owner_id",
				Description: "The ID of the AWS account that owns the route table.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.RouteTable.OwnerId"),
			},
			{
				Name:        "associations",
				Description: "Contains the associations between the route table and one or more subnets or a gateway.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.RouteTable.Associations"),
			},
			{
				Name:        "routes",
				Description: "A list of routes in the route table.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.RouteTable.Routes"),
			},
			{
				Name:        "propagating_vgws",
				Description: "A list of virtual private gateway (VGW) propagating routes.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.RouteTable.PropagatingVgws"),
			},
			{
				Name:        "tags_src",
				Description: "A list of tags that are attached to the route table.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.RouteTable.Tags"),
			},

			// Standard columns for all tables
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(getVpcRouteTableTurbotTags),
			},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.RouteTable.RouteTableId"),
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

//// LIST FUNCTION

//// HYDRATE FUNCTIONS

//// TRANSFORM FUNCTIONS

func getVpcRouteTableTurbotTags(_ context.Context, d *transform.TransformData) (interface{}, error) {
	routeTable := d.HydrateItem.(kaytu.EC2RouteTable).Description.RouteTable
	var turbotTagsMap map[string]string

	// Get the resource tags
	if routeTable.Tags != nil {
		turbotTagsMap = map[string]string{}
		for _, i := range routeTable.Tags {
			turbotTagsMap[*i.Key] = *i.Value
		}
	}

	return turbotTagsMap, nil
}
