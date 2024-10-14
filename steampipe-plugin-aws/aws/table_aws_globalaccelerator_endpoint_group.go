package aws

import (
	"context"

	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsGlobalAcceleratorEndpointGroup(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_globalaccelerator_endpoint_group",
		Description: "AWS Global Accelerator Endpoint Group",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("arn"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"EntityNotFoundException"}),
			},
			Hydrate: opengovernance.GetGlobalAcceleratorEndpointGroup,
		},
		List: &plugin.ListConfig{
			KeyColumns: []*plugin.KeyColumn{
				{Name: "listener_arn", Require: plugin.Optional},
			},
			Hydrate: opengovernance.ListGlobalAcceleratorEndpointGroup,
		},
		Columns: awsKaytuColumns([]*plugin.Column{
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the endpoint group.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.EndpointGroup.EndpointGroupArn"),
			},
			{
				Name:        "listener_arn",
				Description: "The Amazon Resource Name (ARN) of parent listener.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ListenerArn")},
			{
				Name:        "endpoint_descriptions",
				Description: "The list of endpoint objects.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.EndpointGroup.EndpointDescriptions")},
			{
				Name:        "endpoint_group_region",
				Description: "The AWS Region where the endpoint group is located.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.EndpointGroup.EndpointGroupRegion")},
			{
				Name:        "health_check_interval_seconds",
				Description: "The time—10 seconds or 30 seconds—between health checks for each endpoint.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.EndpointGroup.HealthCheckIntervalSeconds")},
			{
				Name:        "health_check_path",
				Description: "If the protocol is HTTP/S, then this value provides the ping path that Global Accelerator uses for the destination on the endpoints for health checks.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.EndpointGroup.HealthCheckPath")},
			{
				Name:        "health_check_port",
				Description: "The port that Global Accelerator uses to perform health checks on endpoints that are part of this endpoint group.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.EndpointGroup.HealthCheckPort")},
			{
				Name:        "health_check_protocol",
				Description: "The protocol that Global Accelerator uses to perform health checks on endpoints that are part of this endpoint group.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.EndpointGroup.HealthCheckProtocol")},
			{
				Name:        "port_overrides",
				Description: "Overrides for destination ports used to route traffic to an endpoint.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.EndpointGroup.PortOverrides")},
			{
				Name:        "threshold_count",
				Description: "The number of consecutive health checks required to set the state of a healthy endpoint to unhealthy, or to set an unhealthy endpoint to healthy.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.EndpointGroup.ThresholdCount")},
			{
				Name:        "traffic_dial_percentage",
				Description: "The percentage of traffic to send to an AWS Region.",
				Type:        proto.ColumnType_DOUBLE,
				Transform:   transform.FromField("Description.EndpointGroup.TrafficDialPercentage")},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.EndpointGroup.EndpointGroupArn").Transform(arnToTitle),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.EndpointGroup.EndpointGroupArn").Transform(arnToAkas),
			},
		}),
	}
}
