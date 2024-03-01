package aws

import (
	"context"

	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsAvailabilityZone(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_availability_zone",
		Description: "AWS Availability Zone",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"name", "region_name"}),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"InvalidParameterValue"}),
			},
			Hydrate: kaytu.GetEC2AvailabilityZone,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListEC2AvailabilityZone,
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:    "name",
					Require: plugin.Optional,
				},
				{
					Name:    "zone_id",
					Require: plugin.Optional,
				},
			},
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the Availability Zone, Local Zone, or Wavelength Zone",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.AvailabilityZone.ZoneName")},
			{
				Name:        "zone_id",
				Description: "The ID of the Availability Zone, Local Zone, or Wavelength Zone.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AvailabilityZone.ZoneId")},
			{
				Name:        "zone_type",
				Description: "The type of zone. The valid values are availability-zone, local-zone, and wavelength-zone.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AvailabilityZone.ZoneType")},
			{
				Name:        "opt_in_status",
				Description: "For Availability Zones, this parameter always has the value of opt-in-not-required. For Local Zones and Wavelength Zones, this parameter is the opt-in status. The possible values are opted-in, and not-opted-in.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AvailabilityZone.OptInStatus")},
			{
				Name:        "group_name",
				Description: "For Availability Zones, this parameter has the same value as the Region name. For Local Zones, the name of the associated group, for example us-west-2-lax-1. For Wavelength Zones, the name of the associated group, for example us-east-1-wl1-bos-wlz-1.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AvailabilityZone.GroupName")},
			{
				Name:        "region_name",
				Description: "The name of the Region.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AvailabilityZone.RegionName")},
			{
				Name:        "parent_zone_name",
				Description: "The name of the zone that handles some of the Local Zone or Wavelength Zone control plane operations, such as API calls.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AvailabilityZone.ParentZoneName")},
			{
				Name:        "parent_zone_id",
				Description: "The ID of the zone that handles some of the Local Zone or Wavelength Zone control plane operations, such as API calls",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AvailabilityZone.ParentZoneId")},
			{
				Name:        "messages",
				Description: "Any messages about the Availability Zone, Local Zone, or Wavelength Zone.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AvailabilityZone.Messages")},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AvailabilityZone.ZoneName")},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("ARN").Transform(arnToAkas)},
		}),
	}
}
