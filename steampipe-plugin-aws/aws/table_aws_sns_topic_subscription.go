package aws

import (
	"context"

	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"
	"github.com/turbot/go-kit/types"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsSnsTopicSubscription(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_sns_topic_subscription",
		Description: "AWS SNS Topic Subscription",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("subscription_arn"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"NotFound", "InvalidParameter"}),
			},
			Hydrate: opengovernance.GetSNSSubscription,
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListSNSSubscription,
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "subscription_arn",
				Description: "Amazon Resource Name of the subscription.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Subscription.SubscriptionArn").NullIfEqual("PendingConfirmation"),
			},
			{
				Name:        "topic_arn",
				Description: "The topic ARN that the subscription is associated with.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Subscription.TopicArn"),
			},
			{
				Name:        "owner",
				Description: "The AWS account ID of the subscription's owner.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Subscription.Owner"),
			},
			{
				Name:        "protocol",
				Description: "The subscription's protocol.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Subscription.Protocol"),
			},
			{
				Name:        "endpoint",
				Description: "The subscription's endpoint (format depends on the protocol).",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Subscription.Endpoint"),
			},
			{
				Name:        "confirmation_was_authenticated",
				Description: "Reflects authentication status of the subscription.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Attributes.ConfirmationWasAuthenticated"),
			},
			{
				Name:        "pending_confirmation",
				Description: "Reflects the confirmation status of the subscription. True if the subscription hasn't been confirmed.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Attributes.PendingConfirmation"),
			},
			{
				Name:        "raw_message_delivery",
				Description: "true if raw message delivery is enabled for the subscription.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Attributes.RawMessageDelivery"),
			},
			{
				Name:        "delivery_policy",
				Description: "The JSON of the subscription's delivery policy.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Attributes.DeliveryPolicy").Transform(transform.UnmarshalYAML),
			},
			{
				Name:        "effective_delivery_policy",
				Description: "The JSON of the effective delivery policy that takes into account the topic delivery policy and account system defaults.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Attributes.EffectiveDeliveryPolicy").Transform(transform.UnmarshalYAML),
			},
			{
				Name:        "redrive_policy",
				Description: "When specified, sends undeliverable messages to the specified Amazon SQS dead-letter queue. Messages that can't be delivered due to client errors (for example, when the subscribed endpoint is unreachable) or server errors (for example, when the service that powers the subscribed endpoint becomes unavailable) are held in the dead-letter queue for further analysis or reprocessing.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Attributes.RedrivePolicy").Transform(transform.UnmarshalYAML),
			},
			{
				Name:        "filter_policy",
				Description: "The filter policy JSON that is assigned to the subscription.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Attributes.FilterPolicy").Transform(transform.UnmarshalYAML),
			},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Subscription.SubscriptionArn").Transform(arnToTitle),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Subscription.SubscriptionArn").Transform(subscriptionArnToAkas),
			},
		}),
	}
}

//// LIST FUNCTION

//// HYDRATE FUNCTIONS

//// TRANSFORM FUNCTIONS

func subscriptionArnToAkas(_ context.Context, d *transform.TransformData) (interface{}, error) {
	arn := types.SafeString(d.Value)

	if arn == "PendingConfirmation" {
		return []string{}, nil
	}

	return []string{arn}, nil
}
