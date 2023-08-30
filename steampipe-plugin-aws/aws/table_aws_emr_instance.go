package aws

import (
	"context"

	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsEmrInstance(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_emr_instance",
		Description: "AWS EMR Instance",
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListEMRInstance,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "cluster_id", Require: plugin.Optional},
				{Name: "instance_fleet_id", Require: plugin.Optional},
				{Name: "instance_group_id", Require: plugin.Optional},
			},
		},

		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "id",
				Description: "The unique identifier for the instance in Amazon EMR.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Instance.Id"),
			},
			{
				Name:        "cluster_id",
				Description: "The unique identifier for the cluster.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ClusterID")},
			{
				Name:        "ec2_instance_id",
				Description: "The unique identifier of the instance in Amazon EC2.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Instance.Ec2InstanceId")},
			{
				Name:        "state",
				Description: "The current state of the instance.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Instance.Status.State")},
			{
				Name:        "instance_fleet_id",
				Description: "The unique identifier of the instance fleet to which an EC2 instance belongs.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Instance.InstanceFleetId")},
			{
				Name:        "instance_group_id",
				Description: "The identifier of the instance group to which this instance belongs.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Instance.InstanceGroupId")},
			{
				Name:        "instance_type",
				Description: "The EC2 instance type, for example m3.xlarge.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Instance.InstanceType")},
			{
				Name:        "market",
				Description: "The instance purchasing option. Valid values are ON_DEMAND or SPOT.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Instance.Market")},
			{
				Name:        "private_dns_name",
				Description: "The private DNS name of the instance.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Instance.PrivateDnsName")},
			{
				Name:        "private_ip_address",
				Description: "The private IP address of the instance.",
				Type:        proto.ColumnType_IPADDR,
				Transform:   transform.FromField("Description.Instance.PrivateIpAddress")},
			{
				Name:        "public_dns_name",
				Description: "The public DNS name of the instance.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Instance.PublicDnsName")},
			{
				Name:        "public_ip_address",
				Description: "The public IP address of the instance.",
				Type:        proto.ColumnType_IPADDR,
				Transform:   transform.FromField("Description.Instance.PublicIpAddress")},
			{
				Name:        "ebs_volumes",
				Description: "The list of Amazon EBS volumes that are attached to this instance.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Instance.EbsVolumes")},
			{
				Name:        "state_change_reason",
				Description: "The status change reason details for the instance.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Instance.Status.StateChangeReason")},
			{
				Name:        "status_timeline",
				Description: "The timeline of the instance status over time.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Instance.Status.Timeline")},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Instance.Id")},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("ARN").Transform(arnToAkas)},
		}),
	}
}
