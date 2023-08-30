package aws

import (
	"context"

	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"
	"strings"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsRDSDBEventSubscription(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_rds_db_event_subscription",
		Description: "AWS RDS DB Event Subscription",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("cust_subscription_id"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"SubscriptionNotFound"}),
			},
			Hydrate: kaytu.GetRDSDBEventSubscription,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListRDSDBEventSubscription,
		},

		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "cust_subscription_id",
				Description: "The RDS event notification subscription Id.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.EventSubscription.CustSubscriptionId")},
			{
				Name:        "customer_aws_id",
				Description: "The AWS customer account associated with the RDS event notification subscription.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.EventSubscription.CustomerAwsId")},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) for the event subscription.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.EventSubscription.EventSubscriptionArn"),
			},
			{
				Name:        "status",
				Description: "The status of the RDS event notification subscription, it can be one of the following: creating | modifying | deleting | active | no-permission | topic-not-exist.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.EventSubscription.Status")},
			{
				Name:        "enabled",
				Description: "A Boolean value indicating if the subscription is enabled. True indicates the subscription is enabled.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.EventSubscription.Enabled")},
			{
				Name:        "sns_topic_arn",
				Description: "The topic ARN of the RDS event notification subscription.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.EventSubscription.SnsTopicArn")},
			{
				Name:        "source_type",
				Description: "The source type for the RDS event notification subscription.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.EventSubscription.SourceType")},
			{
				Name:        "subscription_creation_time",
				Description: "The time the RDS event notification subscription was created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.From(convertStringToRFC3339Timestamp),
			},
			{
				Name:        "event_categories_list",
				Description: "A list of event categories for the RDS event notification subscription.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.EventSubscription.EventCategoriesList")},
			{
				Name:        "source_ids_list",
				Description: "A list of source IDs for the RDS event notification subscription.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.EventSubscription.SourceIdsList")},
			// Steampipe standard columns

			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.EventSubscription.CustSubscriptionId")},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.EventSubscription.EventSubscriptionArn").Transform(transform.EnsureStringArray),
			},
		}),
	}
}

//// LIST FUNCTION

//// HYDRATE FUNCTIONS

//// TRANSFORM FUNCTION

func convertStringToRFC3339Timestamp(_ context.Context, d *transform.TransformData) (interface{}, error) {
	data := d.HydrateItem.(kaytu.RDSDBEventSubscription).Description.EventSubscription
	parsedTime := strings.Replace(*data.SubscriptionCreationTime, " ", "T", 1)

	return parsedTime, nil
}
