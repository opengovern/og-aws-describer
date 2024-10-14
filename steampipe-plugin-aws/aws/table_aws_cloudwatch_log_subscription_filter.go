package aws

import (
	"context"

	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAwsCloudwatchLogSubscriptionFilter(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_cloudwatch_log_subscription_filter",
		Description: "AWS CloudWatch Log Subscription Filter",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"name", "log_group_name"}),
			Hydrate:    opengovernance.GetCloudWatchLogSubscriptionFilter,
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListCloudWatchLogSubscriptionFilter,
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:    "name",
					Require: plugin.Optional,
				},
				{
					Name:    "log_group_name",
					Require: plugin.Optional,
				},
			},
		},

		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the subscription filter.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.SubscriptionFilter.FilterName")},
			{
				Name:        "log_group_name",
				Description: "The name of the log group.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.LogGroupName")},
			{
				Name:        "creation_time",
				Description: "The creation time of the subscription filter.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.SubscriptionFilter.CreationTime").Transform(convertTimestamp),
			},
			{
				Name:        "destination_arn",
				Description: "The Amazon Resource Name (ARN) of the destination.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.SubscriptionFilter.DestinationArn")},
			{
				Name:        "distribution",
				Description: "The method used to distribute log data to the destination.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.SubscriptionFilter.Distribution")},
			{
				Name:        "filter_pattern",
				Description: "A symbolic description of how CloudWatch Logs should interpret the data in each log event.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.SubscriptionFilter.FilterPattern")},
			{
				Name:        "role_arn",
				Description: "The role associated to the filter.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.SubscriptionFilter.RoleArn")},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.SubscriptionFilter.FilterName")},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("ARN").Transform(arnToAkas),
			},
		}),
	}
}
