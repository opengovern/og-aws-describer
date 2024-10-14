package aws

import (
	"context"
	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsSecurityHub(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_securityhub_hub",
		Description: "AWS Security Hub",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("hub_arn"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"InvalidAccessException", "ResourceNotFoundException"}),
			},
			Hydrate: opengovernance.GetSecurityHubHub,
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListSecurityHubHub,
		},

		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "hub_arn",
				Description: "The ARN of the Hub resource that was retrieved.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Hub.HubArn")},
			{
				Name:        "administrator_account",
				Description: "Provides the details for the Security Hub administrator account for the current member account.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.AdministratorAccount"),
			},
			{
				Name:        "auto_enable_controls",
				Description: "Whether to automatically enable new controls when they are added to standards that are enabled.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Hub.AutoEnableControls")},
			{
				Name:        "subscribed_at",
				Description: "The date and time when Security Hub was enabled in the account.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Hub.SubscribedAt")},
			/// Standard columns
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Tags")},
			{
				Name:        "title",
				Description: "The title of hub. This is a constant value 'default'",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromConstant("default"),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("ARN").Transform(arnToAkas),
			},
		}),
	}
}

//// LIST FUNCTION

//// HYDRATE FUNCTIONS
