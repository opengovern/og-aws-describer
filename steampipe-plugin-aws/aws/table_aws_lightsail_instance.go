package aws

import (
	"context"

	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"

	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsLightsailInstance(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_lightsail_instance",
		Description: "AWS Lightsail Instance",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("name"),
			Hydrate:    opengovernance.GetLightsailInstance,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"InvalidResourceName", "DoesNotExist"}),
			},
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListLightsailInstance,
		},

		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the instance.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Instance.Name")},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) specifying the instance.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Instance.Arn")},
			{
				Name:        "blueprint_id",
				Description: "The blueprint ID (e.g., os_amlinux_2016_03).",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Instance.BlueprintId")},
			{
				Name:        "blueprint_name",
				Description: "The friendly name of the blueprint (e.g., Amazon Linux).",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Instance.BlueprintName")},
			{
				Name:        "bundle_id",
				Description: "The bundle for the instance (e.g., micro_1_0).",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Instance.BundleId")},
			{
				Name:        "created_at",
				Description: "The timestamp when the instance was created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Instance.CreatedAt")},
			{
				Name:        "hardware",
				Description: "The size of the vCPU and the amount of RAM for the instance.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Instance.Hardware")},
			{
				Name:        "ip_address_type",
				Description: "The IP address type of the instance.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Instance.IpAddressType")},
			{
				Name:        "ip_v6_addresses",
				Description: "The IPv6 addresses of the instance.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Instance.Ipv6Addresses")},
			{
				Name:        "is_static_ip",
				Description: "A Boolean value indicating whether this instance has a static IP assigned to it.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Instance.IsStaticIp")},
			{
				Name:        "availability_zone",
				Description: "The Availability Zone where the instance is located.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Instance.Location.AvailabilityZone")},
			{
				Name:        "metadata_options",
				Description: "The metadata options for the Amazon Lightsail instance.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Instance.MetadataOptions")},
			{
				Name:        "networking",
				Description: "Information about the public ports and monthly data transfer rates for the instance.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Instance.Networking")},
			{
				Name:        "private_ip_address",
				Description: "The private IP address of the instance.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Instance.PrivateIpAddress")},
			{
				Name:        "public_ip_address",
				Description: "The public IP address of the instance.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Instance.PublicIpAddress")},
			{
				Name:        "resource_type",
				Description: "The type of resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Instance.ResourceType")},
			{
				Name:        "ssh_key_name",
				Description: "The name of the SSH key being used to connect to the instance.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Instance.SshKeyName")},
			{
				Name:        "state_code",
				Description: "The status code for the instance.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Instance.State.Code")},
			{
				Name:        "state_name",
				Description: "The status of the instance.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Instance.State.Name")},
			{
				Name:        "support_code",
				Description: "The support code.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Instance.SupportCode")},
			{
				Name:        "username",
				Description: "The user name for connecting to the instance.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Instance.Username")},
			{
				Name:        "tags_src",
				Description: "A list of tags assigned to the instance.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Instance.Tags")},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Instance.Name")},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Instance.Tags").Transform(getLightsailInstanceTurbotTags),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Instance.Arn").Transform(arnToAkas),
			},
		}),
	}
}

func getLightsailInstanceTurbotTags(_ context.Context, d *transform.TransformData) (interface{}, error) {
	tags := d.Value.([]types.Tag)
	var turbotTagsMap map[string]string
	if tags == nil {
		return nil, nil
	}

	turbotTagsMap = map[string]string{}
	for _, i := range tags {
		turbotTagsMap[*i.Key] = *i.Value
	}

	return &turbotTagsMap, nil
}
