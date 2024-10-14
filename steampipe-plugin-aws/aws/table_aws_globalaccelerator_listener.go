package aws

import (
	"context"

	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsGlobalAcceleratorListener(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_globalaccelerator_listener",
		Description: "AWS Global Accelerator Listener",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("arn"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"EntityNotFoundException"}),
			},
			Hydrate: opengovernance.GetGlobalAcceleratorListener,
		},
		List: &plugin.ListConfig{
			KeyColumns: []*plugin.KeyColumn{
				{Name: "accelerator_arn", Require: plugin.Optional},
			},
			Hydrate: opengovernance.ListGlobalAcceleratorListener,
		},
		Columns: awsKaytuColumns([]*plugin.Column{
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the listener.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Listener.ListenerArn"),
			},
			{
				Name:        "accelerator_arn",
				Description: "The Amazon Resource Name (ARN) of parent accelerator.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AcceleratorArn")},
			{
				Name:        "client_affinity",
				Description: "Client affinity setting for the listener.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Listener.ClientAffinity")},
			{
				Name:        "port_ranges",
				Description: "The list of port ranges for the connections from clients to the accelerator.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Listener.PortRanges")},
			{
				Name:        "protocol",
				Description: "The protocol for the connections from clients to the accelerator.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Listener.Protocol")},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Listener.ListenerArn").Transform(arnToTitle),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Listener.ListenerArn").Transform(arnToAkas),
			},
		}),
	}
}
