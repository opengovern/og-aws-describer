package aws

import (
	"context"
	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAwsUserEffectiveAccess(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_user_effective_access",
		Description: "AWS User Effective Access",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListUserEffectiveAccess,
		},

		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "id",
				Description: "Contains ID to identify a role assignment uniquely.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("ID")},
			{
				Name:        "target_account_id",
				Description: "The identifier of the AWS account from which to list the assignments.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AccountAssignment.AccountId"),
			},
			{
				Name:        "instance_arn",
				Description: "The Amazon Resource Name (ARN) of the SSO Instance under which the operation will be executed.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Instance.InstanceArn"),
			},
			{
				Name:        "permission_set_arn",
				Description: "The ARN of the permission set from which to list assignments.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AccountAssignment.PermissionSetArn"),
			},
			{
				Name:        "user_id",
				Description: "An identifier for an object in IAM Identity Center.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.UserId"),
			},
			{
				Name:        "user_name",
				Description: "The user name.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.User.UserName"),
			},
		}),
	}
}
