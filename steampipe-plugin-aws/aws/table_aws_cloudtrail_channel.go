package aws

import (
	"context"

	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsCloudtrailChannel(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_cloudtrail_channel",
		Description: "AWS CloudTrail Channel",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("arn"),
			Hydrate:    opengovernance.GetCloudTrailChannel,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"ChannelNotFoundException"}),
			},
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListCloudTrailChannel,
		},

		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the cloudtrail channel.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Channel.Name")},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of a channel.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Channel.ChannelArn"),
			},
			{
				Name:        "apply_to_all_regions",
				Description: "Specifies whether the channel applies to a single region or to all regions.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Channel.SourceConfig.ApplyToAllRegions")},
			{
				Name:        "source",
				Description: "The event source for the cloudtrail channel.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Channel.Source")},
			{
				Name:        "advanced_event_selectors",
				Description: "The advanced event selectors that are configured for the channel.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Channel.SourceConfig.AdvancedEventSelectors")},
			{
				Name:        "destinations",
				Description: "The Amazon Web Services service that created the service-linked channel.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Channel.Destinations")},
			{
				Name:        "source_config",
				Description: "Configuration information about the channel.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Channel.SourceConfig")},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Channel.ChannelArn").Transform(transform.EnsureStringArray),
			},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Channel.Name")},
		}),
	}
}
