package aws

import (
	"context"
	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsServiceDiscoveryService(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_service_discovery_service",
		Description: "AWS Service Discovery Service",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"ServiceNotFound"}),
			},
			Hydrate: kaytu.GetServiceDiscoveryService,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListServiceDiscoveryService,
			KeyColumns: plugin.KeyColumnSlice{
				{
					Name:    "namespace_id",
					Require: plugin.Optional,
				},
			},
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the service.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Service.Name")},
			{
				Name:        "id",
				Description: "The ID that Cloud Map assigned to the service when you created it.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Service.Id")},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) that Cloud Map assigns to the service when you create it.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Service.Arn")},
			{
				Name:        "create_date",
				Description: "The date and time that the service was created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Service.CreateDate")},
			{
				Name:        "description",
				Description: "A description for the service.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Service.Description")},
			// The value for the namespace_id column will be populated the service type is DNS_HTTP.
			{
				Name:        "namespace_id",
				Description: "The ID of the namespace.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Service.DnsConfig.NamespaceId")},
			{
				Name:        "instance_count",
				Description: "The number of instances that are currently associated with the service.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Service.DnsConfig.InstanceCount")},
			{
				Name:        "type",
				Description: "Describes the systems that can be used to discover the service instances. DNS_HTTP The service instances can be discovered using either DNS queries or the DiscoverInstances API operation. HTTP The service instances can only be discovered using the DiscoverInstances API operation. DNS Reserved.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Service.Type")},
			{
				Name:        "routing_policy",
				Description: "The routing policy that you want to apply to all Route 53 DNS records that Cloud Map creates when you register an instance and specify this service.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Service.DnsConfig.RoutingPolicy")},
			{
				Name:        "dns_records",
				Description: "An array that contains one DnsRecord object for each Route 53 DNS record that you want Cloud Map to create when you register an instance.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Service.DnsConfig.DnsRecords")},
			{
				Name:        "health_check_config",
				Description: "Public DNS and HTTP namespaces only. Settings for an optional health check. If you specify settings for a health check, Cloud Map associates the health check with the records that you specify in DnsConfig.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Service.HealthCheckConfig")},
			{
				Name:        "health_check_custom_config",
				Description: "Information about an optional custom health check.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Service.HealthCheckCustomConfig")},
			{
				Name:        "tags_src",
				Description: "Information about the tags associated with the namespace.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Tags")},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Service.Name")},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Tags")},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Service.Arn").Transform(transform.EnsureStringArray)},
		}),
	}
}
