package aws

import (
	"context"
	ec2 "github.com/aws/aws-sdk-go-v2/service/ec2/types"

	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsEc2CapacityReservation(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_ec2_capacity_reservation",
		Description: "AWS EC2 Capacity Reservation",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("capacity_reservation_id"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"InvalidCapacityReservationId.NotFound", "InvalidCapacityReservationId.Unavailable", "InvalidCapacityReservationId.Malformed"}),
			},
			Hydrate: kaytu.GetEC2CapacityReservation,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListEC2CapacityReservation,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "instance_type", Require: plugin.Optional},
				{Name: "owner_id", Require: plugin.Optional},
				{Name: "availability_zone_id", Require: plugin.Optional},
				{Name: "availability_zone", Require: plugin.Optional},
				{Name: "instance_platform", Require: plugin.Optional},
				{Name: "tenancy", Require: plugin.Optional},
				{Name: "state", Require: plugin.Optional},
				{Name: "start_date", Require: plugin.Optional},
				{Name: "end_date", Require: plugin.Optional},
				{Name: "end_date_type", Require: plugin.Optional},
				{Name: "instance_match_criteria", Require: plugin.Optional},
			},
		},

		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "capacity_reservation_id",
				Description: "The ID of the capacity reservation.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.CapacityReservation.CapacityReservationId")},
			{
				Name:        "capacity_reservation_arn",
				Description: "The Amazon Resource Name (ARN) of the capacity reservation.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.CapacityReservation.CapacityReservationArn")},
			{
				Name:        "instance_type",
				Description: "The type of instance for which the capacity reservation reserves capacity.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.CapacityReservation.InstanceType")},
			{
				Name:        "state",
				Description: "The current state of the capacity reservation. A capacity reservation can be in one of the following states: 'active', 'expired', 'cancelled', 'pending', 'failed'.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.CapacityReservation.State")},
			{
				Name:        "availability_zone",
				Description: "The availability zone in which the capacity is reserved.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.CapacityReservation.AvailabilityZone")},
			{
				Name:        "availability_zone_id",
				Description: "The availability zone ID of the capacity reservation.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.CapacityReservation.AvailabilityZoneId")},
			{
				Name:        "available_instance_count",
				Description: "The remaining capacity. Indicates the number of instances that can be launched in the capacity reservation.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.CapacityReservation.AvailableInstanceCount")},
			{
				Name:        "create_date",
				Description: "The date and time at which the capacity reservation was created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.CapacityReservation.CreateDate")},
			{
				Name:        "ebs_optimized",
				Description: "Indicates whether the capacity reservation supports EBS-optimized instances.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.CapacityReservation.EbsOptimized")},
			{
				Name:        "end_date",
				Description: "The date and time at which the capacity reservation expires.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.CapacityReservation.EndDate")},
			{
				Name:        "end_date_type",
				Description: "Indicates the way in which the capacity reservation ends. A capacity reservation can have one of the following end types: 'unlimited', 'limited'.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.CapacityReservation.EndDateType")},
			{
				Name:        "ephemeral_storage",
				Description: "Indicates whether the capacity reservation supports instances with temporary, block-level storage.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.CapacityReservation.EphemeralStorage")},
			{
				Name:        "instance_match_criteria",
				Description: "Indicates the type of instance launches that the capacity reservation accepts. The options include: 'open', 'targeted'.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.CapacityReservation.InstanceMatchCriteria")},
			{
				Name:        "instance_platform",
				Description: "The type of operating system for which the capacity reservation reserves capacity.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.CapacityReservation.InstancePlatform")},
			{
				Name:        "owner_id",
				Description: "The ID of the AWS account that owns the capacity reservation.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.CapacityReservation.OwnerId")},
			{
				Name:        "start_date",
				Description: "The date and time at which the capacity reservation was started.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.CapacityReservation.StartDate")},
			{
				Name:        "tenancy",
				Description: "Indicates the tenancy of the capacity reservation. A capacity reservation can have one of the following tenancy settings: 'default', 'dedicated'.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.CapacityReservation.Tenancy")},
			{
				Name:        "total_instance_count",
				Description: "The total number of instances for which the capacity reservation reserves capacity.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.CapacityReservation.TotalInstanceCount")},
			{
				Name:        "capacity_allocations",
				Description: "Information about instance capacity usage.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "tag_src",
				Description: "Any tags assigned to the capacity reservation.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.CapacityReservation.Tags")},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.CapacityReservation.CapacityReservationId")},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.CapacityReservation.Tags").Transform(ec2CapacityReservationTagListToTurbotTags),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.CapacityReservation.CapacityReservationArn").Transform(transform.EnsureStringArray),
			},
		}),
	}
}

func ec2CapacityReservationTagListToTurbotTags(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	plugin.Logger(ctx).Trace("ec2CapacityReservationTagListToTurbotTags")
	tagList := d.Value.([]ec2.Tag)

	// Mapping the resource tags inside turbotTags
	var turbotTagsMap map[string]string
	if tagList != nil {
		turbotTagsMap = map[string]string{}
		for _, i := range tagList {
			turbotTagsMap[*i.Key] = *i.Value
		}
	}

	return turbotTagsMap, nil
}

//// UTILITY FUNCTION

// Build ec2 capacity reservation list call input filter
