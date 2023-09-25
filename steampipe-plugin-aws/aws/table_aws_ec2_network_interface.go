package aws

import (
	"context"
	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsEc2NetworkInterface(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_ec2_network_interface",
		Description: "AWS EC2 Network Interface",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("network_interface_id"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"InvalidNetworkInterfaceID.NotFound", "InvalidNetworkInterfaceID.Unavailable", "InvalidNetworkInterfaceID.Malformed"}),
			},
			Hydrate: kaytu.GetEC2NetworkInterface,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListEC2NetworkInterface,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "association_id", Require: plugin.Optional},
				{Name: "association_allocation_id", Require: plugin.Optional},
				{Name: "association_ip_owner_id", Require: plugin.Optional},
				{Name: "association_public_ip", Require: plugin.Optional},
				{Name: "association_public_dns_name", Require: plugin.Optional},
				{Name: "attachment_id", Require: plugin.Optional},
				{Name: "attachment_time", Require: plugin.Optional},
				{Name: "delete_on_instance_termination", Require: plugin.Optional, Operators: []string{"=", "<>"}},
				{Name: "attached_instance_id", Require: plugin.Optional},
				{Name: "attached_instance_owner_id", Require: plugin.Optional},
				{Name: "attachment_status", Require: plugin.Optional},
				{Name: "availability_zone", Require: plugin.Optional},
				{Name: "description", Require: plugin.Optional},
				{Name: "mac_address", Require: plugin.Optional},
				{Name: "owner_id", Require: plugin.Optional},
				{Name: "private_ip_address", Require: plugin.Optional},
				{Name: "private_dns_name", Require: plugin.Optional},
				{Name: "requester_id", Require: plugin.Optional},
				{Name: "requester_managed", Require: plugin.Optional, Operators: []string{"=", "<>"}},
				{Name: "source_dest_check", Require: plugin.Optional, Operators: []string{"=", "<>"}},
				{Name: "status", Require: plugin.Optional},
			},
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "network_interface_id",
				Description: "The ID of the network interface.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.NetworkInterface.NetworkInterfaceId")},
			{
				Name:        "status",
				Description: "The status of the network interface.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.NetworkInterface.Status")},
			{
				Name:        "interface_type",
				Description: "The type of network interface.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.NetworkInterface.InterfaceType")},
			{
				Name:        "description",
				Description: "A description.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.NetworkInterface.Description")},
			{
				Name:        "availability_zone",
				Description: "The Availability Zone.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.NetworkInterface.AvailabilityZone")},
			{
				Name:        "owner_id",
				Description: "The AWS account ID of the owner of the network interface.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.NetworkInterface.OwnerId")},
			{
				Name:        "association_allocation_id",
				Description: "Allocation id for the association. Association can be an Elastic IP address (IPv4 only), or a Carrier IP address.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.NetworkInterface.Association.AllocationId")},
			{
				Name:        "association_carrier_ip",
				Description: "The carrier IP address associated with the network interface.",
				Type:        proto.ColumnType_IPADDR,
				Transform:   transform.FromField("Description.NetworkInterface.Association.CarrierIp")},
			{
				Name:        "association_customer_owned_ip",
				Description: "The customer-owned IP address associated with the network interface.",
				Type:        proto.ColumnType_IPADDR,
				Transform:   transform.FromField("Description.NetworkInterface.Association.CustomerOwnedIp")},
			{
				Name:        "association_id",
				Description: "The association ID.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.NetworkInterface.Association.AssociationId")},
			{
				Name:        "association_ip_owner_id",
				Description: "The ID of the Elastic IP address owner.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.NetworkInterface.Association.IpOwnerId")},
			{
				Name:        "association_public_dns_name",
				Description: "The public DNS name of the association.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.NetworkInterface.Association.PublicDnsName")},
			{
				Name:        "association_public_ip",
				Description: "The address of the Elastic IP address bound to the network interface.",
				Type:        proto.ColumnType_IPADDR,
				Transform:   transform.FromField("Description.NetworkInterface.Association.PublicIp")},
			{
				Name:        "attached_instance_id",
				Description: "The ID of the attached instance.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.NetworkInterface.Attachment.InstanceId")},
			{
				Name:        "attached_instance_owner_id",
				Description: "The AWS account ID of the owner of the attached instance.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.NetworkInterface.Attachment.InstanceOwnerId")},
			{
				Name:        "attachment_id",
				Description: "The ID of the network interface attachment.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.NetworkInterface.Attachment.AttachmentId")},
			{
				Name:        "attachment_status",
				Description: "The attachment state.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.NetworkInterface.Attachment.Status")},
			{
				Name:        "attachment_time",
				Description: "The timestamp indicating when the attachment initiated.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.NetworkInterface.Attachment.AttachTime")},
			{
				Name:        "delete_on_instance_termination",
				Description: "Indicates whether the network interface is deleted when the instance is terminated.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.NetworkInterface.Attachment.DeleteOnTermination")},
			{
				Name:        "device_index",
				Description: "The device index of the network interface attachment on the instance.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.NetworkInterface.Attachment.DeviceIndex")},

			{
				Name:        "mac_address",
				Description: "The MAC address of the interface.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.NetworkInterface.MacAddress")},
			{
				Name:        "outpost_arn",
				Description: "The Amazon Resource Name (ARN) of the Outpost, if applicable.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.NetworkInterface.OutpostArn")},
			{
				Name:        "private_dns_name",
				Description: "The private DNS name",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.NetworkInterface.PrivateDnsName")},
			{
				Name:        "private_ip_address",
				Description: "The IPv4 address of the network interface within the subnet.",
				Type:        proto.ColumnType_IPADDR,
				Transform:   transform.FromField("Description.NetworkInterface.PrivateIpAddress")},
			{
				Name:        "requester_id",
				Description: "The ID of the entity that launched the instance on your behalf (for example, AWS Management Console or Auto Scaling).",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.NetworkInterface.RequesterId")},
			{
				Name:        "requester_managed",
				Description: "Indicates whether the network interface is being managed by AWS.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.NetworkInterface.RequesterManaged")},
			{
				Name:        "source_dest_check",
				Description: "Indicates whether traffic to or from the instance is validated.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.NetworkInterface.SourceDestCheck")},
			{
				Name:        "subnet_id",
				Description: "The ID of the subnet.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.NetworkInterface.SubnetId"),
			},
			{
				Name:        "vpc_id",
				Description: "The ID of the VPC.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.NetworkInterface.VpcId"),
			},
			{
				Name:        "groups",
				Description: "Any security groups for the network interface.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.NetworkInterface.Groups")},
			{
				Name:        "ipv6_addresses",
				Description: "The IPv6 addresses associated with the network interface.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.NetworkInterface.Ipv6Addresses")},
			{
				Name:        "private_ip_addresses",
				Description: "The IPv4 address of the network interface within the subnet.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.NetworkInterface.PrivateIpAddresses")},
			{
				Name:        "tags_src",
				Description: "A list of tags that are attached to the network interface.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.NetworkInterface.TagSet")},
			// Standard columns for all tables

			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.NetworkInterface.NetworkInterfaceId")},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(getEc2NetworkInterfaceTurbotTags),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Hydrate:     getAwsEc2NetworkInterfaceAkas,
				Transform:   transform.FromValue(),
			},
		}),
	}
}

//// LIST FUNCTION

//// HYDRATE FUNCTIONS

func getAwsEc2NetworkInterfaceAkas(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	networkInterface := h.Item.(kaytu.EC2NetworkInterface).Description.NetworkInterface
	metadata := h.Item.(kaytu.EC2NetworkInterface).Metadata

	// Get data for turbot defined properties
	akas := []string{"arn:" + metadata.Partition + ":ec2:" + metadata.Region + ":" + metadata.AccountID + ":network-interface/" + *networkInterface.NetworkInterfaceId}

	return akas, nil
}

//// TRANSFORM FUNCTIONS

func getEc2NetworkInterfaceTurbotTags(_ context.Context, d *transform.TransformData) (interface{}, error) {
	data := d.HydrateItem.(kaytu.EC2NetworkInterface).Description.NetworkInterface

	// Get resource tags
	var turbotTags map[string]string
	if data.TagSet != nil {
		turbotTags = map[string]string{}
		for _, i := range data.TagSet {
			turbotTags[*i.Key] = *i.Value
		}
	}
	return turbotTags, nil
}
