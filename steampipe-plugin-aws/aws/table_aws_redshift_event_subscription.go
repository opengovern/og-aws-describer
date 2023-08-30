package aws

import (
	"context"
	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"
	"strings"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAwsRedshiftEventSubscription(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_redshift_event_subscription",
		Description: "AWS Redshift Event Subscription",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("cust_subscription_id"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"SubscriptionNotFound"}),
			},
			Hydrate: kaytu.GetRedshiftEventSubscription,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListRedshiftEventSubscription,
		},

		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "cust_subscription_id",
				Description: "The name of the Amazon Redshift event notification subscription.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.EventSubscription.CustSubscriptionId")},
			{
				Name:        "customer_aws_id",
				Description: "The AWS customer account associated with the Amazon Redshift event notification subscription.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.EventSubscription.CustomerAwsId")},
			{
				Name:        "enabled",
				Description: "A boolean value indicating whether the subscription is enabled or disabled",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.EventSubscription.Enabled")},
			{
				Name:        "severity",
				Description: "The event severity specified in the Amazon Redshift event notification subscription.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.EventSubscription.Severity")},
			{
				Name:        "sns_topic_arn",
				Description: "The Amazon Resource Name (ARN) of the Amazon SNS topic used by the event notification subscription.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.EventSubscription.SnsTopicArn")},
			{
				Name:        "source_type",
				Description: "The source type of the events returned by the Amazon Redshift event notification.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.EventSubscription.SourceType")},
			{
				Name:        "status",
				Description: "The status of the Amazon Redshift event notification subscription.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.EventSubscription.Status")},
			{
				Name:        "subscription_creation_time",
				Description: "The date and time the Amazon Redshift event notification subscription was created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.EventSubscription.SubscriptionCreationTime")},
			{
				Name:        "event_categories_list",
				Description: "The list of Amazon Redshift event categories specified in the event notification subscription.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.EventSubscription.EventCategoriesList")},
			{
				Name:        "source_ids_list",
				Description: "A list of the sources that publish events to the Amazon Redshift event notification subscription.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.EventSubscription.SourceIdsList")},
			{
				Name:        "tags_src",
				Description: "The list of tags for the event subscription.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.EventSubscription.Tags")},

			// Standard columns for all tables

			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.EventSubscription.CustSubscriptionId")},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.EventSubscription.Tags").Transform(redshiftEventSubListToTurbotTags),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Hydrate:     getRedshiftEventSubscriptionAkas,
				Transform:   transform.FromValue(),
			},
		}),
	}
}

func getRedshiftEventSubscriptionAkas(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	region := d.EqualsQualString(matrixKeyRegion)
	parameterData := h.Item.(kaytu.RedshiftEventSubscription)

	aka := "arn:" + parameterData.Metadata.Partition + ":redshift:" + region + ":" + parameterData.Metadata.AccountID + ":eventsubscription"

	if strings.HasPrefix(*parameterData.Description.EventSubscription.CustSubscriptionId, ":") {
		aka = aka + *parameterData.Description.EventSubscription.CustSubscriptionId
	} else {
		aka = aka + ":" + *parameterData.Description.EventSubscription.CustSubscriptionId
	}

	return []string{aka}, nil
}

func redshiftEventSubListToTurbotTags(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	tagList := d.HydrateItem.(kaytu.RedshiftEventSubscription)

	if len(tagList.Description.EventSubscription.Tags) > 0 {
		turbotTagsMap := map[string]string{}
		for _, i := range tagList.Description.EventSubscription.Tags {
			turbotTagsMap[*i.Key] = *i.Value
		}
		return turbotTagsMap, nil
	}
	return nil, nil
}
