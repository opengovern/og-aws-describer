package aws

import (
	"context"

	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsEmrInstanceFleet(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_emr_instance_fleet",
		Description: "AWS EMR Instance Fleet",
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListEMRInstanceFleet,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"InvalidRequestException"}),
			},
		},

		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the instance fleet.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.InstanceFleet.Name")},
			{
				Name:        "id",
				Description: "The identifier of the instance fleet.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.InstanceFleet.Id")},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) specifying the instance fleet.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("ARN").Transform(arnToAkas),
			},
			{
				Name:        "cluster_id",
				Description: "The unique identifier for the cluster.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ClusterID")},
			{
				Name:        "instance_fleet_type",
				Description: "The type of the instance fleet. Valid values are MASTER, CORE or TASK.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.InstanceFleet.InstanceFleetType")},
			{
				Name:        "state",
				Description: "The current state of the instance fleet.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.InstanceFleet.Status.State")},
			{
				Name:        "provisioned_on_demand_capacity",
				Description: "The number of On-Demand units that have been provisioned for the instance fleet to fulfill TargetOnDemandCapacity.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.InstanceFleet.ProvisionedOnDemandCapacity")},
			{
				Name:        "provisioned_spot_capacity",
				Description: "The number of Spot units that have been provisioned for this instance fleet to fulfill TargetSpotCapacity.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.InstanceFleet.ProvisionedSpotCapacity")},
			{
				Name:        "target_on_demand_capacity",
				Description: "The target capacity of On-Demand units for the instance fleet, which determines how many On-Demand Instances to provision.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.InstanceFleet.TargetOnDemandCapacity")},
			{
				Name:        "target_spot_capacity",
				Description: "The target capacity of Spot units for the instance fleet, which determines how many Spot Instances to provision.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.InstanceFleet.TargetSpotCapacity")},
			{
				Name:        "instance_type_specifications",
				Description: "An array of specifications for the instance types that comprise an instance fleet.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.InstanceFleet.InstanceTypeSpecifications")},
			{
				Name:        "launch_specifications",
				Description: "Describes the launch specification for an instance fleet.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.InstanceFleet.LaunchSpecifications")},
			{
				Name:        "state_change_reason",
				Description: "Provides status change reason details for the instance fleet.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.InstanceFleet.Status.StateChangeReason")},
			{
				Name:        "status_timeline",
				Description: "Provides historical timestamps for the instance fleet, including the time of creation, the time it became ready to run jobs, and the time of termination.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.InstanceFleet.Status.Timeline")},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.From(emrInstanceFleetTitle),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("ARN").Transform(transform.EnsureStringArray),
			},
		}),
	}
}

//// TRANSFORM FUNCTIONS

func emrInstanceFleetTitle(_ context.Context, d *transform.TransformData) (interface{}, error) {
	data := d.HydrateItem.(kaytu.EMRInstanceFleet).Description.InstanceFleet

	if *data.Name == "" {
		return data.Id, nil
	}
	return data.Name, nil
}
