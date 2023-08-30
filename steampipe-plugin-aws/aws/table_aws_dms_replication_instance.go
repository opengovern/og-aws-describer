package aws

import (
	"context"

	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsDmsReplicationInstance(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_dms_replication_instance",
		Description: "AWS DMS Replication Instance",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("arn"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"InvalidParameterValueException", "ResourceNotFoundFault", "InvalidParameterCombinationException"}),
			},
			Hydrate: kaytu.GetDMSReplicationInstance,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListDMSReplicationInstance,
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:    "replication_instance_identifier",
					Require: plugin.Optional,
				},
				{
					Name:    "arn",
					Require: plugin.Optional,
				},
				{
					Name:    "replication_instance_class",
					Require: plugin.Optional,
				},
				{
					Name:    "engine_version",
					Require: plugin.Optional,
				},
			},
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "replication_instance_identifier",
				Description: "The identifier of the replication instance.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ReplicationInstance.ReplicationInstanceIdentifier")},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the replication instance.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ReplicationInstance.ReplicationInstanceArn")},
			{
				Name:        "replication_instance_class",
				Description: "The compute and memory capacity of the replication instance as defined for the specified replication instance class.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ReplicationInstance.ReplicationInstanceClass")},
			{
				Name:        "engine_version",
				Description: "The engine version number of the replication instance.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ReplicationInstance.EngineVersion")},
			{
				Name:        "publicly_accessible",
				Description: "Specifies the accessibility options for the replication instance.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.ReplicationInstance.PubliclyAccessible")},
			{
				Name:        "allocated_storage",
				Description: "The amount of storage (in gigabytes) that is allocated for the replication instance.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.ReplicationInstance.AllocatedStorage")},
			{
				Name:        "auto_minor_version_upgrade",
				Description: "Boolean value indicating if minor version upgrades will be automatically applied to the instance.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.ReplicationInstance.AutoMinorVersionUpgrade")},
			{
				Name:        "availability_zone",
				Description: "The Availability Zone for the instance.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ReplicationInstance.AvailabilityZone")},
			{
				Name:        "dns_name_servers",
				Description: "The DNS name servers supported for the replication instance to access your on-premise source or target database.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ReplicationInstance.DnsNameServers")},
			{
				Name:        "free_until",
				Description: "The expiration date of the free replication instance that is part of the Free DMS program.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.ReplicationInstance.FreeUntil")},
			{
				Name:        "instance_create_time",
				Description: "The time the replication instance was created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.ReplicationInstance.InstanceCreateTime")},
			{
				Name:        "kms_key_id",
				Description: "An AWS KMS key identifier that is used to encrypt the data on the replication instance.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ReplicationInstance.KmsKeyId")},
			{
				Name:        "multi_az",
				Description: "Specifies whether the replication instance is a Multi-AZ deployment.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.ReplicationInstance.MultiAZ")},
			{
				Name:        "preferred_maintenance_window",
				Description: "The maintenance window times for the replication instance.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ReplicationInstance.PreferredMaintenanceWindow")},
			{
				Name:        "replication_instance_private_ip_address",
				Description: "The private IP address of the replication instance.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ReplicationInstance.ReplicationInstancePrivateIpAddress")},
			{
				Name:        "replication_instance_public_ip_address",
				Description: "The public IP address of the replication instance.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ReplicationInstance.ReplicationInstancePublicIpAddress")},
			{
				Name:        "replication_instance_status",
				Description: "The status of the replication instance.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ReplicationInstance.ReplicationInstanceStatus")},
			{
				Name:        "secondary_availability_zone",
				Description: "The Availability Zone of the standby replication instance in a Multi-AZ deployment.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ReplicationInstance.SecondaryAvailabilityZone")},
			{
				Name:        "pending_modified_values",
				Description: "The pending modification values.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ReplicationInstance.PendingModifiedValues")},
			{
				Name:        "replication_instance_private_ip_addresses",
				Description: "One or more private IP addresses for the replication instance.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ReplicationInstance.ReplicationInstancePrivateIpAddresses")},
			{
				Name:        "replication_instance_public_ip_addresses",
				Description: "One or more public IP addresses for the replication instance.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ReplicationInstance.ReplicationInstancePublicIpAddresses")},
			{
				Name:        "replication_subnet_group",
				Description: "The subnet group for the replication instance.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ReplicationInstance.ReplicationSubnetGroup")},
			{
				Name:        "vpc_security_groups",
				Description: "The VPC security group for the instance.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ReplicationInstance.VpcSecurityGroups")},
			{
				Name:        "tags_src",
				Description: "A list of tags currently associated with the replication instance.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Tags"),
			},

			// Steampipe Standard columns
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ReplicationInstance.ReplicationInstanceIdentifier"),
			},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(dmsReplicationInstanceTagListToTagsMap),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ReplicationInstance.ReplicationInstanceArn").Transform(arnToAkas),
			},
		}),
	}
}

//// LIST FUNCTION

//// HYDRATE FUNCTIONS

//// TRANSFORM FUNCTIONS

func dmsReplicationInstanceTagListToTagsMap(_ context.Context, d *transform.TransformData) (interface{}, error) {
	data := d.HydrateItem.(kaytu.DMSReplicationInstance).Description

	// Mapping the resource tags inside turbotTags
	if data.Tags != nil {
		turbotTagsMap := map[string]string{}
		for _, i := range data.Tags {
			turbotTagsMap[*i.Key] = *i.Value
		}
		return turbotTagsMap, nil
	}
	return nil, nil
}
