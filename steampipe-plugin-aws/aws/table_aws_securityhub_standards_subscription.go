package aws

import (
	"context"

	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsSecurityHubStandardsSubscription(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_securityhub_standards_subscription",
		Description: "AWS Security Hub Standards Subscription",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListSecurityHubStandardsSubscription,
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the standard.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Standard.Name")},
			{
				Name:        "standards_arn",
				Description: "The ARN of a standard.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Standard.StandardsArn")},
			{
				Name:        "description",
				Description: "The description of the standard.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Standard.Description")},
			{
				Name:        "enabled_by_default",
				Description: "Indicates whether the standard is enabled by default.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Standard.EnabledByDefault")},
			{
				Name:        "standards_status",
				Description: "The status of the standard subscription.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.StandardsSubscription.StandardsStatus")},
			{
				Name:        "standards_status_reason_code",
				Description: "The reason code that represents the reason for the current status of a standard subscription.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.StandardsSubscription.StandardsStatusReason.StatusReasonCode")},
			{
				Name:        "standards_subscription_arn",
				Description: "The ARN of a resource that represents your subscription to a supported standard.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.StandardsSubscription.StandardsSubscriptionArn")},

			// JSON columns
			{
				Name:        "standards_input",
				Description: "A key-value pair of input for the standard.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.StandardsSubscription.StandardsInput")},
			{
				Name:        "standards_managed_by",
				Description: "Provides details about the management of a security standard.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Standard.StandardsManagedBy")},

			// Standard columns for all tables
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Standard.Name")},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Standard.StandardsArn").Transform(arnToAkas),
			},
		}),
	}
}
