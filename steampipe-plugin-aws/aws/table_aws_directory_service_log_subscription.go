package aws

import (
	"context"
	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsDirectoryServiceLogSubscription(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_directory_service_log_subscription",
		Description: "AWS Directory Service Log Subscription",
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListDirectoryServiceLogSubscription,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"EntityDoesNotExistException"}),
			},
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:    "directory_id",
					Require: plugin.Optional,
				},
			},
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "directory_id",
				Description: "Identifier (ID) of the directory that you want to associate with the log subscription.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.LogSubscription.DirectoryId"),
			},
			{
				Name:        "log_group_name",
				Description: "The name of the log group.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.LogSubscription.LogGroupName"),
			},
			{
				Name:        "subscription_created_date_time",
				Description: "The date and time that the log subscription was created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.LogSubscription.SubscriptionCreatedDateTime"),
			},

			// Steampipe Standard columns
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.LogSubscription.LogGroupName"),
			},
		}),
	}
}
