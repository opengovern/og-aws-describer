package aws

import (
	"context"
	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"

	servicediscoveryv1 "github.com/aws/aws-sdk-go/service/servicediscovery"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsServiceDiscoveryNamespace(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_service_discovery_namespace",
		Description: "AWS Service Discovery Namespace",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"NamespaceNotFound"}),
			},
			Hydrate: kaytu.GetServiceDiscoveryNamespace,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListServiceDiscoveryNamespace,
			KeyColumns: plugin.KeyColumnSlice{
				{
					Name:    "type",
					Require: plugin.Optional,
				},
				{
					Name:    "name",
					Require: plugin.Optional,
				},
			},
		},
		GetMatrixItemFunc: SupportedRegionMatrix(servicediscoveryv1.EndpointsID),
		Columns: awsRegionalColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the namespace.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Namespace.Name")},
			{
				Name:        "id",
				Description: "The ID of the namespace.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Namespace.Id")},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) that Cloud Map assigns to the namespace when you create it.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Namespace.Arn")},
			{
				Name:        "create_date",
				Description: "The date and time that the namespace was created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Namespace.CreateDate")},
			{
				Name:        "description",
				Description: "A description for the namespace.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Namespace.Description")},
			{
				Name:        "service_count",
				Description: "The number of services that were created using the namespace.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Namespace.ServiceCount")},
			{
				Name:        "type",
				Description: "The type of the namespace, either public or private.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Namespace.Type")},
			{
				Name:        "dns_properties",
				Description: "A complex type that contains the ID for the Route 53 hosted zone that Cloud Map creates when you create a namespace.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Namespace.Properties.DnsProperties")},
			{
				Name:        "http_properties",
				Description: "A complex type that contains the name of an HTTP namespace.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Namespace.Properties.HttpProperties")},
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
				Transform:   transform.FromField("Description.Namespace.Name")},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Tags")},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Namespace.Arn").Transform(transform.EnsureStringArray)},
		}),
	}
}
