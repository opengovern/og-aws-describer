package aws

import (
	"context"

	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsSnsTopic(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_sns_topic",
		Description: "AWS SNS Topic",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("topic_arn"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"NotFound", "InvalidParameter"}),
			},
			Hydrate: opengovernance.GetSNSTopic,
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListSNSTopic,
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "topic_arn",
				Description: "Amazon Resource Name (ARN) of the Topic.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Attributes.TopicArn"),
			},
			{
				Name:        "display_name",
				Description: "The human-readable name used in the From field for notifications to email and email-json endpoints.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Attributes.DisplayName"),
			},
			{
				Name:        "application_failure_feedback_role_arn",
				Description: "IAM role for failed deliveries of notification messages sent to topics with platform application endpoint.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Attributes.ApplicationFailureFeedbackRoleArn"),
			},
			{
				Name:        "application_success_feedback_role_arn",
				Description: "IAM role for successful deliveries of notification messages sent to topics with platform application endpoint.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Attributes.ApplicationSuccessFeedbackRoleArn"),
			},
			{
				Name:        "application_success_feedback_sample_rate",
				Description: "Sample rate for successful deliveries of notification messages sent to topics with platform application endpoint.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Attributes.ApplicationSuccessFeedbackSampleRate"),
			},
			{
				Name:        "firehose_failure_feedback_role_arn",
				Description: "IAM role for failed deliveries of notification messages sent to topics with kinesis data firehose endpoint.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Attributes.FirehoseFailureFeedbackRoleArn"),
			},
			{
				Name:        "firehose_success_feedback_role_arn",
				Description: "IAM role for successful deliveries of notification messages sent to topics with kinesis data firehose endpoint.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Attributes.FirehoseSuccessFeedbackRoleArn"),
			},
			{
				Name:        "firehose_success_feedback_sample_rate",
				Description: "Sample rate for successful deliveries of notification messages sent to topics with firehose endpoint.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Attributes.FirehoseSuccessFeedbackSampleRate"),
			},
			{
				Name:        "http_failure_feedback_role_arn",
				Description: "IAM role for failed deliveries of notification messages sent to topics with http endpoint.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Attributes.HTTPFailureFeedbackRoleArn"),
			},
			{
				Name:        "http_success_feedback_role_arn",
				Description: "IAM role for successful deliveries of notification messages sent to topics with http endpoint.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Attributes.HTTPSuccessFeedbackRoleArn"),
			},
			{
				Name:        "http_success_feedback_sample_rate",
				Description: "Sample rate for successful deliveries of notification messages sent to topics with http endpoint.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Attributes.HTTPSuccessFeedbackSampleRate"),
			},
			{
				Name:        "lambda_failure_feedback_role_arn",
				Description: "IAM role for failed deliveries of notification messages sent to topics with lambda endpoint.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Attributes.LambdaFailureFeedbackRoleArn"),
			},
			{
				Name:        "lambda_success_feedback_role_arn",
				Description: "IAM role for successful deliveries of notification messages sent to topics with lambda endpoint.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Attributes.LambdaSuccessFeedbackRoleArn"),
			},
			{
				Name:        "lambda_success_feedback_sample_rate",
				Description: "Sample rate for successful deliveries of notification messages sent to topics with lambda endpoint.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Attributes.LambdaSuccessFeedbackSampleRate"),
			},
			{
				Name:        "owner",
				Description: "The AWS account ID of the topic's owner.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Attributes.Owner"),
			},
			{
				Name:        "sqs_failure_feedback_role_arn",
				Description: "IAM role for failed deliveries of notification messages sent to topics with sqs endpoint.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Attributes.SQSFailureFeedbackRoleArn"),
			},
			{
				Name:        "sqs_success_feedback_role_arn",
				Description: "IAM role for successful deliveries of notification messages sent to topics with sqs endpoint.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Attributes.SQSSuccessFeedbackRoleArn"),
			},
			{
				Name:        "sqs_success_feedback_sample_rate",
				Description: "Sample rate for successful deliveries of notification messages sent to topics with sqs endpoint.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Attributes.SQSSuccessFeedbackSampleRate"),
			},
			{
				Name:        "subscriptions_confirmed",
				Description: "The number of confirmed subscriptions for the topic.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Attributes.SubscriptionsConfirmed"),
			},
			{
				Name:        "subscriptions_deleted",
				Description: "The number of deleted subscriptions for the topic.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Attributes.SubscriptionsDeleted"),
			},
			{
				Name:        "subscriptions_pending",
				Description: "The number of subscriptions pending confirmation for the topic.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Attributes.SubscriptionsPending"),
			},
			{
				Name:        "kms_master_key_id",
				Description: "The ID of an AWS-managed customer master key (CMK) for Amazon SNS or a custom CMK.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Attributes.KmsMasterKeyId"),
			},
			{
				Name:        "tags_src",
				Description: "The list of tags associated with the topic.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Tags"),
			},
			{
				Name:        "policy",
				Description: "The topic's access control policy (i.e. Resource IAM Policy).",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Attributes.Policy").Transform(transform.UnmarshalYAML),
			},
			{
				Name:        "policy_std",
				Description: "Contains the policy in a canonical form for easier searching.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Attributes.Policy").Transform(unescape).Transform(policyToCanonical),
			},

			{
				Name:        "delivery_policy",
				Description: "The JSON object of the topic's delivery policy.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Attributes.DeliveryPolicy").Transform(transform.UnmarshalYAML),
			},
			{
				Name:        "effective_delivery_policy",
				Description: "The effective delivery policy, taking system defaults into account.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Attributes.EffectiveDeliveryPolicy").Transform(transform.UnmarshalYAML),
			},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Attributes.TopicArn").Transform(arnToTitle),
			},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(snsTopicTurbotTags),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Attributes.TopicArn").Transform(arnToAkas),
			},
		}),
	}
}

//// LIST FUNCTION

//// HYDRATE FUNCTIONS

//// TRANSFORM FUNCTIONS

func snsTopicTurbotTags(_ context.Context, d *transform.TransformData) (interface{}, error) {
	tags := d.HydrateItem.(opengovernance.SNSTopic).Description.Tags
	var turbotTagsMap map[string]string
	if tags != nil {
		turbotTagsMap = map[string]string{}
		for _, i := range tags {
			turbotTagsMap[*i.Key] = *i.Value
		}
	}
	return turbotTagsMap, nil
}
