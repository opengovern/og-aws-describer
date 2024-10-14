package aws

import (
	"context"

	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsGlobalAcceleratorAccelerator(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_globalaccelerator_accelerator",
		Description: "AWS Global Accelerator Accelerator",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("arn"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"EntityNotFoundException"}),
			},
			Hydrate: opengovernance.GetGlobalAcceleratorAccelerator,
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListGlobalAcceleratorAccelerator,
		},
		Columns: awsKaytuColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the accelerator.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Accelerator.Name")},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the accelerator.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Accelerator.AcceleratorArn"),
			},
			{
				Name:        "created_time",
				Description: "The date and time that the accelerator was created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Accelerator.CreatedTime")},
			{
				Name:        "dns_name",
				Description: "The Domain Name System (DNS) name that Global Accelerator creates that points to your accelerator's static IP addresses.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Accelerator.DnsName")},
			{
				Name:        "enabled",
				Description: "Indicates whether the accelerator is enabled.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Accelerator.Enabled")},
			{
				Name:        "ip_address_type",
				Description: "The value for the address type must be IPv4.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Accelerator.IpAddressType")},
			{
				Name:        "ip_sets",
				Description: "The static IP addresses that Global Accelerator associates with the accelerator.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Accelerator.IpSets")},
			{
				Name:        "last_modified_time",
				Description: "The date and time that the accelerator was last modified.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Accelerator.LastModifiedTime")},
			{
				Name:        "status",
				Description: "Describes the deployment status of the accelerator.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Accelerator.Status")},
			{
				Name:        "tags_src",
				Description: "A list of tags associated with the accelerator.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Tags")},
			{
				Name:        "accelerator_attributes",
				Description: "Attributes of the accelerator.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.AcceleratorAttributes")},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Accelerator.Name")},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(globalacceleratorAcceleratorTurbotTags),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Accelerator.AcceleratorArn").Transform(arnToAkas),
			},
		}),
	}
}

//// TRANSFORM FUNCTIONS

func globalacceleratorAcceleratorTurbotTags(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	tags := d.HydrateItem.(opengovernance.GlobalAcceleratorAccelerator).Description.Tags

	// Mapping the resource tags inside turbotTags
	var turbotTagsMap map[string]string
	if tags != nil {
		turbotTagsMap = map[string]string{}
		for _, i := range tags {
			turbotTagsMap[*i.Key] = *i.Value
		}
	}

	return turbotTagsMap, nil
}
