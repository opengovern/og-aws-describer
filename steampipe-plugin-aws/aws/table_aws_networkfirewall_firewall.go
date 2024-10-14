package aws

import (
	"context"

	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsNetworkFirewallFirewall(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_networkfirewall_firewall",
		Description: "AWS Network Firewall Firewall",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AnyColumn([]string{"arn", "name"}),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"ResourceNotFoundException", "InvalidRequestException", "ValidationException"}),
			},
			Hydrate: opengovernance.GetNetworkFirewallFirewall,
		},
		List: &plugin.ListConfig{
			KeyColumns: plugin.OptionalColumns([]string{"vpc_id"}),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"InvalidRequestException", "ValidationException"}),
			},
			Hydrate: opengovernance.ListNetworkFirewallFirewall,
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the firewall",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Firewall.FirewallArn")},
			{
				Name:        "name",
				Description: "The descriptive name of the firewall.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Firewall.FirewallName")},
			{
				Name:        "vpc_id",
				Description: "The unique identifier of the VPC where the firewall is in use.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Firewall.VpcId")},
			{
				Name:        "delete_protection",
				Description: "A flag indicating whether it is possible to delete the firewall.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Firewall.DeleteProtection")},
			{
				Name:        "description",
				Description: "A description of the firewall.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Firewall.Description")},
			{
				Name:        "id",
				Description: "The id of the firewall.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Firewall.FirewallId")},
			{
				Name:        "policy_arn",
				Description: "The Amazon Resource Name (ARN) of the firewall policy.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Firewall.FirewallPolicyArn")},
			{
				Name:        "policy_change_protection",
				Description: "A setting indicating whether the firewall is protected against a change to the firewall policy association.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Firewall.FirewallPolicyChangeProtection")},
			{
				Name:        "subnet_change_protection",
				Description: "A setting indicating whether the firewall is protected against changes to the subnet associations.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Firewall.SubnetChangeProtection")},
			{
				Name:        "encryption_configuration",
				Description: "A complex type that contains the Amazon Web Services KMS encryption configuration settings for the firewall.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Firewall.EncryptionConfiguration")},
			{
				Name:        "firewall_status",
				Description: "Detailed information about the current status of a Firewall.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.FirewallStatus")},
			{
				Name:        "subnet_mappings",
				Description: "The public subnets that Network Firewall is using for the firewall.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Firewall.SubnetMappings")},
			{
				Name:        "logging_configuration",
				Description: "Defines how Network Firewall performs logging for a Firewall.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.LoggingConfiguration"),
			},
			{
				Name:        "tags_src",
				Description: "A list of tags associated with the firewall",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Firewall.Tags")},
			// Steampipe standard columns
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Firewall.FirewallName")},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(getNetworkFirewallFirewallTurbotTags),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Firewall.FirewallArn").Transform(arnToAkas),
			},
		}),
	}
}

//// TRANSFORM FUNCTIONS

func getNetworkFirewallFirewallTurbotTags(_ context.Context, d *transform.TransformData) (interface{}, error) {
	tags := d.HydrateItem.(opengovernance.NetworkFirewallFirewall).Description.Firewall.Tags
	return networkFirewallV2TagsToMap(tags)
}
