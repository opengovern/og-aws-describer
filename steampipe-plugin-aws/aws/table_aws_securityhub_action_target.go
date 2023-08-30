package aws

import (
	"context"

	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsSecurityHubActionTarget(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_securityhub_action_target",
		Description: "AWS Security Hub Action Target",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("arn"),
			Hydrate:    kaytu.GetSecurityHubActionTarget,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"InvalidInputException"}),
			},
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListSecurityHubActionTarget,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"ResourceNotFoundException"}),
			},
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the action target.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ActionTarget.Name")},
			{
				Name:        "arn",
				Description: "The ARN for the target action.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ActionTarget.ActionTargetArn"),
			},
			{
				Name:        "description",
				Description: "The description of the target action.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ActionTarget.Description")},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ActionTarget.Name")},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ActionTarget.ActionTargetArn").Transform(transform.EnsureStringArray),
			},
		}),
	}
}
