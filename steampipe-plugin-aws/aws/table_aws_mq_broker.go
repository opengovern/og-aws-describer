package aws

import (
	"context"
	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAwsMQBroker(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_mq_broker",
		Description: "AWS MQ Broker",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("broker_name"),
			Hydrate:    kaytu.GetMQBroker,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListMQBroker,
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "broker_id",
				Description: "The unique ID that Amazon MQ generates for the broker.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.BrokerDescription.BrokerId")},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the broker.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.BrokerDescription.BrokerArn"),
			},
			{
				Name:        "broker_state",
				Description: "The broker's status.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.BrokerDescription.BrokerState")},
			{
				Name:        "deployment_mode",
				Description: "The broker's deployment mode.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.BrokerDescription.DeploymentMode")},
			{
				Name:        "created",
				Description: "The time when the broker was created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.BrokerDescription.Created")},
			{
				Name:        "host_instance_type",
				Description: "The broker's instance type.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.BrokerDescription.HostInstanceType")},
			{
				Name:        "authentication_strategy",
				Description: "The authentication strategy used to secure the broker. The default is SIMPLE.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.BrokerDescription.AuthenticationStrategy"),
			},
			{
				Name:        "data_replication_mode",
				Description: "Describes whether this broker is a part of a data replication pair.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.BrokerDescription.DataReplicationMode"),
			},
			{
				Name:        "engine_type",
				Description: "The type of broker engine. Currently, Amazon MQ supports ACTIVEMQ and RABBITMQ.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.BrokerDescription.EngineType"),
			},
			{
				Name:        "engine_version",
				Description: "The broker engine's version.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.BrokerDescription.EngineVersion"),
			},
			{
				Name:        "pending_authentication_strategy",
				Description: "The authentication strategy that will be applied when the broker is rebooted. The default is SIMPLE.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.BrokerDescription.PendingAuthenticationStrategy"),
			},
			{
				Name:        "pending_data_replication_mode",
				Description: "Describes whether this broker will be a part of a data replication pair after reboot.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.BrokerDescription.PendingDataReplicationMode"),
			},
			{
				Name:        "pending_engine_version",
				Description: "The broker engine version to upgrade to.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.BrokerDescription.PendingEngineVersion"),
			},
			{
				Name:        "pending_host_instance_type",
				Description: "The broker's host instance type to upgrade to.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.BrokerDescription.PendingHostInstanceType"),
			},
			{
				Name:        "publicly_accessible",
				Description: "Enables connections from applications outside of the VPC that hosts the broker's subnets.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.BrokerDescription.PubliclyAccessible"),
			},
			{
				Name:        "storage_type",
				Description: "The broker's storage type.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.BrokerDescription.StorageType"),
			},
			{
				Name:        "auto_minor_version_upgrade",
				Description: "Enables automatic upgrades to new minor versions for brokers, as new versions are released and supported by Amazon MQ.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.BrokerDescription.AutoMinorVersionUpgrade"),
			},

			// JSON columns
			{
				Name:        "actions_required",
				Description: "Actions required for a broker.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.BrokerDescription.ActionsRequired"),
			},
			{
				Name:        "broker_instances",
				Description: "A list of information about allocated brokers.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.BrokerDescription.BrokerInstances"),
			},
			{
				Name:        "configurations",
				Description: "The list of all revisions for the specified configuration.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.BrokerDescription.Configurations"),
			},
			{
				Name:        "data_replication_metadata",
				Description: "The replication details of the data replication-enabled broker. Only returned if dataReplicationMode is set to CRDR.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.BrokerDescription.DataReplicationMetadata"),
			},
			{
				Name:        "encryption_options",
				Description: "Encryption options for the broker.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.BrokerDescription.EncryptionOptions"),
			},
			{
				Name:        "ldap_server_metadata",
				Description: "The metadata of the LDAP server used to authenticate and authorize connections to the broker.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.BrokerDescription.LdapServerMetadata"),
			},
			{
				Name:        "logs",
				Description: "The list of information about logs currently enabled and pending to be deployed for the specified broker.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.BrokerDescription.Logs"),
			},
			{
				Name:        "pending_ldap_server_metadata",
				Description: "The metadata of the LDAP server that will be used to authenticate and authorize connections to the broker after it is rebooted.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.BrokerDescription.PendingLdapServerMetadata"),
			},
			{
				Name:        "maintenance_window_start_time",
				Description: "The parameters that determine the WeeklyStartTime.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.BrokerDescription.MaintenanceWindowStartTime"),
			},
			{
				Name:        "pending_data_replication_metadata",
				Description: "The pending replication details of the data replication-enabled broker. Only returned if pendingDataReplicationMode is set to CRDR.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.BrokerDescription.PendingDataReplicationMetadata"),
			},
			{
				Name:        "pending_security_groups",
				Description: "The list of pending security groups to authorize connections to brokers.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.BrokerDescription.PendingSecurityGroups"),
			},
			{
				Name:        "security_groups",
				Description: "The list of rules (1 minimum, 125 maximum) that authorize connections to brokers.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.BrokerDescription.SecurityGroups"),
			},
			{
				Name:        "subnet_ids",
				Description: "The list of groups that define which subnets and IP ranges the broker can use from different Availability Zones.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.BrokerDescription.SubnetIds"),
			},
			{
				Name:        "users",
				Description: "The list of all broker usernames for the specified broker.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.BrokerDescription.Users"),
			},

			// Steampipe standard columns
			{
				Name:        "tags",
				Description: "A list of tags attached to the broker.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.BrokerDescription.Tags"),
			},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.BrokerDescription.BrokerName"),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.BrokerDescription.BrokerArn").Transform(transform.EnsureStringArray),
			},
		}),
	}
}
