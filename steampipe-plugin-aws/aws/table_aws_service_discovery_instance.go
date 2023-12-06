package aws

import (
	"context"
	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsServiceDiscoveryInstance(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_service_discovery_instance",
		Description: "AWS Service Discovery Instance",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"service_id", "id"}),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"InstanceNotFound"}),
			},
			Hydrate: kaytu.GetServiceDiscoveryInstance,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListServiceDiscoveryInstance,
			KeyColumns: plugin.KeyColumnSlice{
				{
					Name:    "service_id",
					Require: plugin.Optional,
				},
			},
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "id",
				Description: "The ID of the instance.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Instance.Id"),
			},
			{
				Name:        "service_id",
				Description: "The ID of the service.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ServiceId"),
			},
			{
				Name:        "ec2_instance_id",
				Description: "The Amazon EC2 instance ID for the instance. When the AWS_EC2_INSTANCE_ID attribute is specified, then the AWS_INSTANCE_IPV4 attribute contains the primary private IPv4 address.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Instance.Attributes.AWS_EC2_INSTANCE_ID"),
			},
			{
				Name:        "alias_dns_name",
				Description: "For an alias record that routes traffic to an Elastic Load Balancing load balancer, the DNS name that's associated with the load balancer.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Instance.Attributes.AWS_ALIAS_DNS_NAME"),
			},
			{
				Name:        "instance_cname",
				Description: "A CNAME record, the domain name that Route 53 returns in response to DNS queries (for example, example.com ).",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Instance.Attributes.AWS_INSTANCE_CNAME"),
			},
			{
				Name:        "init_health_status",
				Description: "If the service configuration includes HealthCheckCustomConfig, you can optionally use AWS_INIT_HEALTH_STATUS to specify the initial status of the custom health check, HEALTHY or UNHEALTHY. If you don't specify a value for AWS_INIT_HEALTH_STATUS, the initial status is HEALTHY.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Instance.Attributes.AWS_INIT_HEALTH_STATUS"),
			},
			{
				Name:        "instance_ipv4",
				Description: "For an A record, the IPv4 address that Route 53 returns in response to DNS queries.",
				Type:        proto.ColumnType_IPADDR,
				Transform:   transform.FromField("Description.Instance.Attributes.AWS_INSTANCE_IPV4"),
			},
			{
				Name:        "instance_ipv6",
				Description: "For an AAAA record, the IPv6 address that Route 53 returns in response to DNS queries.",
				Type:        proto.ColumnType_IPADDR,
				Transform:   transform.FromField("Description.Instance.Attributes.AWS_INSTANCE_IPV6"),
			},
			{
				Name:        "instance_port",
				Description: "For an SRV record, the value that Route 53 returns for the port. In addition, if the service includes HealthCheckConfig, the port on the endpoint that Route 53 sends requests to.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Instance.Attributes.AWS_INSTANCE_PORT"),
			},
			{
				Name:        "attributes",
				Description: "Attributes of the instance.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Instance.Attributes")},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Instance.Id"),
			},
		}),
	}
}
