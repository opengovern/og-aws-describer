package aws

import (
	"context"

	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsSecurityLakeSubscriber(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_securitylake_subscriber",
		Description: "AWS Security Lake Subscriber",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("subscription_id"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"ResourceNotFoundException"}),
			},
			Hydrate: kaytu.GetSecurityLakeSubscriber,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListSecurityLakeSubscriber,
		},

		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "subscriber_name",
				Description: "The name of your Amazon Security Lake subscriber account.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Subscriber.SubscriberName")},
			{
				Name:        "subscription_id",
				Description: "The subscription ID of the Amazon Security Lake subscriber account.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Subscriber.SubscriptionId")},
			{
				Name:        "created_at",
				Description: "The date and time when the subscription was created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Subscriber.CreatedAt")},
			{
				Name:        "subscription_status",
				Description: "Subscription status of the Amazon Security Lake subscriber account.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Subscriber.SubscriptionStatus")},
			{
				Name:        "updated_at",
				Description: "The date and time when the subscription was updated.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Subscriber.UpdatedAt")},
			{
				Name:        "external_id",
				Description: "The external ID of the subscriber.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Subscriber.ExternalId")},
			{
				Name:        "role_arn",
				Description: "The Amazon Resource Name (ARN) specifying the role of the subscriber.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Subscriber.RoleArn")},
			{
				Name:        "s3_bucket_arn",
				Description: "The Amazon Resource Name (ARN) for the Amazon S3 bucket.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Subscriber.S3BucketArn")},
			{
				Name:        "sns_arn",
				Description: "The Amazon Resource Name (ARN) for the Amazon Simple Notification Service.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Subscriber.SnsArn")},
			{
				Name:        "subscriber_description",
				Description: "The subscriber descriptions for a subscriber account.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Subscriber.SubscriberDescription")},
			{
				Name:        "subscription_endpoint",
				Description: "The subscription endpoint to which exception messages are posted.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Subscriber.SubscriptionEndpoint")},
			{
				Name:        "subscription_protocol",
				Description: "The subscription protocol to which exception messages are posted.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Subscriber.SubscriptionProtocol")},
			{
				Name:        "access_types",
				Description: "You can choose to notify subscribers of new objects with an Amazon Simple Queue Service (Amazon SQS) queue or through messaging to an HTTPS endpoint provided by the subscriber. Subscribers can consume data by directly querying Lake Formation tables in your S3 bucket via services like Amazon Athena. This subscription type is defined as LAKEFORMATION.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Subscriber.AccessTypes")},
			{
				Name:        "source_types",
				Description: "Amazon Security Lake supports logs and events collection for the natively-supported Amazon Web Services services.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Subscriber.SourceTypes")},

			// Standard columns
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Subscriber.SubscriberName")},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Subscriber.ProductArn").Transform(arnToAkas),
			},
		}),
	}
}
