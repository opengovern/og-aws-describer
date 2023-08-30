package aws

import (
	"context"

	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsEc2ReservedInstance(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_ec2_reserved_instance",
		Description: "AWS EC2 Reserved Instance",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("reserved_instance_id"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"InvalidParameterValue", "InvalidInstanceID.Unavailable", "InvalidInstanceID.Malformed"}),
			},
			Hydrate: kaytu.GetEC2ReservedInstances,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListEC2ReservedInstances,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "availability_zone", Require: plugin.Optional},
				{Name: "duration", Require: plugin.Optional},
				{Name: "end_time", Require: plugin.Optional},
				{Name: "fixed_price", Require: plugin.Optional},
				{Name: "instance_type", Require: plugin.Optional},
				{Name: "scope", Require: plugin.Optional},
				{Name: "product_description", Require: plugin.Optional},
				{Name: "start_time", Require: plugin.Optional},
				{Name: "instance_state", Require: plugin.Optional},
				{Name: "usage_price", Require: plugin.Optional},
				{Name: "offering_class", Require: plugin.Optional},
				{Name: "offering_type", Require: plugin.Optional},
			},
		},

		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "reserved_instance_id",
				Description: "The ID of the Reserved instance.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.ReservedInstances.ReservedInstancesId")},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) specifying the instance.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("ARN"),
			},
			{
				Name:        "instance_type",
				Description: "The instance type on which the Reserved Instance can be used.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ReservedInstances.InstanceType")},
			{
				Name:        "instance_state",
				Description: "The state of the Reserved Instance purchase.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.ReservedInstances.State")},
			{
				Name:        "availability_zone",
				Description: "The Availability Zone in which the Reserved Instance can be used.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ReservedInstances.AvailabilityZone")},
			{
				Name:        "currency_code",
				Description: "The currency of the Reserved Instance. It's specified using ISO 4217 standard currency codes. At this time, the only supported currency is USD.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ReservedInstances.CurrencyCode")},
			{
				Name:        "duration",
				Description: "The duration of the Reserved Instance, in seconds.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.ReservedInstances.Duration")},
			{
				Name:        "end_time",
				Description: "The time when the Reserved Instance expires.",
				Type:        proto.ColumnType_TIMESTAMP,

				Transform: transform.FromField("Description.ReservedInstances.End")},
			{
				Name:        "fixed_price",
				Description: "The purchase price of the Reserved Instance.",
				Type:        proto.ColumnType_DOUBLE,
				Transform:   transform.FromField("Description.ReservedInstances.FixedPrice")},
			{
				Name:        "instance_count",
				Description: "The number of reservations purchased.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.ReservedInstances.InstanceCount")},
			{
				Name:        "instance_tenancy",
				Description: "The tenancy of the instance.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ReservedInstances.InstanceTenancy")},
			{
				Name:        "offering_class",
				Description: "The offering class of the Reserved Instance.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ReservedInstances.OfferingClass")},
			{
				Name:        "offering_type",
				Description: "The Reserved Instance offering type.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ReservedInstances.OfferingType")},
			{
				Name:        "product_description",
				Description: "The Reserved Instance product platform description.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ReservedInstances.ProductDescription")},
			{
				Name:        "scope",
				Description: "The scope of the Reserved Instance.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ReservedInstances.Scope")},
			{
				Name:        "start_time",
				Description: "The date and time the Reserved Instance started.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.ReservedInstances.Start")},
			{
				Name:        "usage_price",
				Description: "The usage price of the Reserved Instance, per hour.",
				Type:        proto.ColumnType_DOUBLE,
				Transform:   transform.FromField("Description.ReservedInstances.UsagePrice")},
			{
				Name:        "reserved_instances_modifications",
				Description: "The Reserved Instance modification information.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ModificationDetails")},
			{
				Name:        "tags_src",
				Description: "A list of tags assigned to the reserved instance.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ReservedInstances.Tags")},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ReservedInstances.ReservedInstancesId")},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(getEc2ReservedInstanceTurbotTags),
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

//// TRANSFORM FUNCTION

func getEc2ReservedInstanceTurbotTags(_ context.Context, d *transform.TransformData) (interface{}, error) {
	instance := d.HydrateItem.(kaytu.EC2ReservedInstances).Description.ReservedInstances
	return ec2V2TagsToMap(instance.Tags)
}
