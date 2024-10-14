package aws

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsEc2SpotPrice(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_ec2_spot_price",
		Description: "AWS EC2 Spot Price History",
		List: &plugin.ListConfig{
			//Hydrate: opengovernance.ListEC2SpotPrice,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"InvalidParameterValue"}),
			},
			KeyColumns: []*plugin.KeyColumn{
				{Name: "availability_zone", Require: plugin.Optional},
				{Name: "instance_type", Require: plugin.Optional},
				{Name: "product_description", Require: plugin.Optional},
				{Name: "start_time", Require: plugin.Optional},
				{Name: "end_time", Require: plugin.Optional},
			},
		},

		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{Name: "availability_zone", Description: "The Availability Zone.", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.SpotPrice.AvailabilityZone")},
			{Name: "instance_type", Description: "The instance type.", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.SpotPrice.InstanceType")},
			{Name: "product_description", Description: "A general description of the AMI.", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.SpotPrice.ProductDescription")},
			{Name: "spot_price", Description: "The maximum price per unit hour that you are willing to pay for a Spot Instance.", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.SpotPrice.SpotPrice")},
			{Name: "create_timestamp", Description: "The time stamp of the Spot price history.", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("Description.SpotPrice.Timestamp")},
			{Name: "start_time", Description: "The date and time, up to the past 90 days, from which to start retrieving the price history data.", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromQual("start_time")},
			{Name: "end_time", Description: "The date and time, up to the current date, from which to stop retrieving the price history data.", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromQual("end_time")},
		}),
	}
}
