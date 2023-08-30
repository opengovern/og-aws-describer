package aws

import (
	"context"

	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsIamPolicyAttachment(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_iam_policy_attachment",
		Description: "AWS IAM Policy Attachment",
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListIAMPolicyAttachment,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "is_attached", Require: plugin.Optional, Operators: []string{"<>", "="}},
			},
		},
		Columns: awsKaytuColumns([]*plugin.Column{
			{
				Name:        "policy_arn",
				Description: "The Amazon Resource Name (ARN) specifying the IAM policy.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.PolicyArn")},
			{
				Name:        "is_attached",
				Description: "Specifies whether the policy is attached to at least one IAM user, group, or role.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.PolicyAttachmentCount").Transform(attachementCountToBool),
			},
			{
				Name:        "policy_groups",
				Description: "A list of IAM groups that the policy is attached to.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.PolicyGroups")},
			{
				Name:        "policy_roles",
				Description: "A list of IAM roles that the policy is attached to.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.PolicyRoles")},
			{
				Name:        "policy_users",
				Description: "A list of IAM users that the policy is attached to.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.PolicyUsers")},
		}),
	}
}
