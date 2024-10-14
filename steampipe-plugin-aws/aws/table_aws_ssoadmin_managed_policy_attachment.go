package aws

import (
	"context"
	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAwsSsoAdminManagedPolicyAttachment(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_ssoadmin_managed_policy_attachment",
		Description: "AWS SSO Managed Policy Attachment",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListSSOAdminPolicyAttachment,
		},

		Columns: awsRegionalColumns([]*plugin.Column{
			{
				Name:        "permission_set_arn",
				Description: "The ARN of the permission set.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.PermissionSetArn"),
			},
			{
				Name:        "name",
				Description: "The name of the IAM managed policy.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AttachedManagedPolicy.Name"),
			},
			{
				Name:        "instance_arn",
				Description: "The Amazon Resource Name (ARN) of the SSO Instance under which the operation will be executed.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.InstanceArn"),
			},
			{
				Name:        "managed_policy_arn",
				Description: "The ARN of the IAM managed policy.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AttachedManagedPolicy.Arn"),
			},
			// Standard columns for all tables
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AttachedManagedPolicy.Name"),
			},
		}),
	}
}
