package aws

import (
	"context"

	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAwsEc2CapacityReservationFleet(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_ec2_capacity_reservation_fleet",
		Description: "AWS Ec2 CapacityReservationFleet",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("capacity_reservation_fleet_id"),
			Hydrate:    opengovernance.GetEC2CapacityReservationFleet,
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListEC2CapacityReservationFleet,
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "capacity_reservation_fleet_id",
				Description: "The id of the capacity reservation fleet.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.CapacityReservationFleet.CapacityReservationFleetId")},
			{
				Name:        "capacity_reservation_fleet_arn",
				Description: "The Amazon Resource Name (ARN) of the capacity reservation fleet",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.CapacityReservationFleet.CapacityReservationFleetArn")},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.CapacityReservationFleet.CapacityReservationFleetId")},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(getEc2CapacityReservationFleetTurbotTags),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.CapacityReservationFleet.CapacityReservationFleetArn").Transform(arnToAkas),
			},
		}),
	}
}

//// TRANSFORM FUNCTIONS

func getEc2CapacityReservationFleetTurbotTags(_ context.Context, d *transform.TransformData) (interface{}, error) {
	tags := d.HydrateItem.(opengovernance.EC2CapacityReservationFleet).Description.CapacityReservationFleet.Tags
	return ec2V2TagsToMap(tags)
}
