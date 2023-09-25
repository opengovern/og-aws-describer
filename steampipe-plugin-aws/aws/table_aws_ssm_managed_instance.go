package aws

import (
	"context"
	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsSSMManagedInstance(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_ssm_managed_instance",
		Description: "AWS SSM Managed Instance",
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListSSMManagedInstance,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "instance_id", Require: plugin.Optional},
				{Name: "agent_version", Require: plugin.Optional},
				{Name: "ping_status", Require: plugin.Optional},
				{Name: "platform_type", Require: plugin.Optional},
				{Name: "activation_id", Require: plugin.Optional},
				{Name: "resource_type", Require: plugin.Optional},
				{Name: "association_status", Require: plugin.Optional},
			},
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The name assigned to an on-premises server or virtual machine (VM) when it is activated as a Systems Manager managed instance.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.InstanceInformation.Name")},
			{
				Name:        "instance_id",
				Description: "The ID of the instance.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.InstanceInformation.InstanceId")},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) specifying the instance.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("ARN"),
			},
			{
				Name:        "resource_type",
				Description: "The type of instance. Instances are either EC2 instances or managed instances.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.InstanceInformation.ResourceType")},
			{
				Name:        "ping_status",
				Description: "Connection status of SSM Agent.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.InstanceInformation.PingStatus")},
			{
				Name:        "activation_id",
				Description: "The activation ID created by Systems Manager when the server or VM was registered.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.InstanceInformation.ActivationId")},
			{
				Name:        "agent_version",
				Description: "The version of SSM Agent running on your Linux instance.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.InstanceInformation.AgentVersion")},
			{
				Name:        "association_status",
				Description: "The status of the association.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.InstanceInformation.AssociationStatus")},
			{
				Name:        "computer_name",
				Description: "The fully qualified host name of the managed instance.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.InstanceInformation.ComputerName")},
			{
				Name:        "ip_address",
				Description: "The IP address of the managed instance.",
				Type:        proto.ColumnType_IPADDR,
				Transform:   transform.FromField("Description.InstanceInformation.IPAddress")},
			{
				Name:        "is_latest_version",
				Description: "Indicates whether the latest version of SSM Agent is running on your Linux Managed Instance.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.InstanceInformation.IsLatestVersion")},
			{
				Name:        "last_association_execution_date",
				Description: "The date the association was last run.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.InstanceInformation.LastAssociationExecutionDate")},
			{
				Name:        "last_ping_date_time",
				Description: "The date and time when the agent last pinged the Systems Manager service.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.InstanceInformation.LastPingDateTime")},
			{
				Name:        "last_successful_association_execution_date",
				Description: "The last date the association was successfully run.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.InstanceInformation.LastSuccessfulAssociationExecutionDate")},
			{
				Name:        "platform_name",
				Description: "The name of the operating system platform running on your instance.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.InstanceInformation.PlatformName")},
			{
				Name:        "platform_type",
				Description: "The operating system platform type.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.InstanceInformation.PlatformType")},
			{
				Name:        "platform_version",
				Description: "The version of the OS platform running on your instance.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.InstanceInformation.PlatformVersion")},
			{
				Name:        "registration_date",
				Description: "The date the server or VM was registered with AWS as a managed instance.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.InstanceInformation.RegistrationDate")},
			{
				Name:        "association_overview",
				Description: "Information about the association.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.InstanceInformation.AssociationOverview")},
			// Steampipe Standard columns

			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.InstanceInformation.InstanceId")},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("ARN").Transform(arnToAkas),
			},
		}),
	}
}

//// LIST FUNCTION
