package aws

import (
	"context"
	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsInstanceAvailability(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_ec2_instance_availability",
		Description: "AWS EC2 Instance Availability",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListEC2InstanceAvailability,
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:    "instance_type",
					Require: plugin.Optional,
				},
			},
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "instance_type",
				Description: "The instance type. For more information, see [ Instance Types ](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-types.html) in the Amazon Elastic Compute Cloud User Guide.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.InstanceAvailability.InstanceType")},
			{
				Name:        "location",
				Description: "The identifier for the location. This depends on the location type. For example, if the location type is region, the location is the Region code (for example, us-east-2.)",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.InstanceAvailability.Location")},
			{
				Name:        "location_type",
				Description: "The type of location.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.InstanceAvailability.LocationType")},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.InstanceAvailability.InstanceType")},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("ARN").Transform(arnToAkas)},
		}),
	}
}
