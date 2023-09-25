package aws

import (
	"context"
	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"

	"github.com/turbot/go-kit/types"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsSqsQueue(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_sqs_queue",
		Description: "AWS SQS Queue",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("queue_url"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"AWS.SimpleQueueService.NonExistentQueue"}),
			},
			Hydrate: kaytu.GetSQSQueue,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListSQSQueue,
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "queue_url",
				Description: "The URL of the Amazon SQS queue.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Attributes.QueueUrl"),
			},
			{
				Name:        "queue_arn",
				Description: "The Amazon resource name (ARN) of the queue.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Attributes.QueueArn"),
			},
			{
				Name:        "fifo_queue",
				Description: "Returns true if the queue is FIFO.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Attributes.FifoQueue"),
				Default:     false,
			},
			{
				Name:        "fifo_throughput_limit",
				Description: "Specifies whether the FIFO queue throughput quota applies to the entire queue or per message group.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Attributes.FifoThroughputLimit"),
				Default:     false,
			},
			{
				Name:        "delay_seconds",
				Description: "The default delay on the queue in seconds.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Attributes.DelaySeconds"),
			},
			{
				Name:        "max_message_size",
				Description: "The limit of how many bytes a message can contain before Amazon SQS rejects it.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Attributes.MaximumMessageSize"),
			},
			{
				Name:        "message_retention_seconds",
				Description: "The length of time, in seconds, for which Amazon SQS retains a message.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Attributes.MessageRetentionPeriod"),
			},
			{
				Name:        "receive_wait_time_seconds",
				Description: "The length of time, in seconds, for which the ReceiveMessage action waits for a message to arrive.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Attributes.ReceiveMessageWaitTimeSeconds"),
			},
			{
				Name:        "sqs_managed_sse_enabled",
				Description: "Returns true if the queue is using SSE-SQS encryption with SQS-owned encryption keys.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Attributes.SqsManagedSseEnabled"),
			},
			{
				Name:        "visibility_timeout_seconds",
				Description: "The visibility timeout for the queue in seconds.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Attributes.VisibilityTimeout"),
			},
			{
				Name:        "policy",
				Description: "The resource IAM policy of the queue.",
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
				Name:        "redrive_policy",
				Description: "The string that includes the parameters for the dead-letter queue functionality of the source queue as a JSON object.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Attributes.RedrivePolicy").Transform(transform.UnmarshalYAML),
			},
			{
				Name:        "content_based_deduplication",
				Description: "Mentions whether content-based deduplication is enabled for the queue.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Attributes.ContentBasedDeduplication"),
			},
			{
				Name:        "kms_master_key_id",
				Description: "the ID of an AWS-managed customer master key (CMK) for Amazon SQS or a custom CMK.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Attributes.KmsMasterKeyId"),
			},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Tags"),
			},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Attributes.QueueUrl").Transform(getAwsSqsQueueTitle),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Attributes.QueueArn").Transform(arnToAkas),
			},
		}),
	}
}

//// LIST FUNCTION

//// HYDRATE FUNCTIONS

//// TRANSFORM FUNCTION

func getAwsSqsQueueTitle(_ context.Context, d *transform.TransformData) (interface{}, error) {
	queueURL := types.SafeString(d.Value)

	queueName, err := extractNameFromSqsQueueURL(queueURL)
	if err != nil {
		return nil, err
	}

	return queueName, nil
}
