package aws

import (
	"context"
	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsIamGroup(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_iam_group",
		Description: "AWS IAM Group",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AnyColumn([]string{"name", "arn"}),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"ValidationError", "NoSuchEntity", "InvalidParameter"}),
			},
			Hydrate: kaytu.GetIAMGroup,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListIAMGroup,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "path", Require: plugin.Optional},
			},
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The friendly name that identifies the group.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Group.GroupName"),
			},
			{
				Name:        "group_id",
				Description: "The stable and unique string identifying the group.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Group.GroupId"),
			},
			{
				Name:        "path",
				Description: "The path to the group.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Group.Path"),
			},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) specifying the group.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Group.Arn"),
			},
			{
				Name:        "create_date",
				Description: "The date and time, when the group was created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Group.CreateDate"),
			},
			{
				Name:        "inline_policies",
				Description: "A list of policy documents that are embedded as inline policies for the group.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.InlinePolicies"),
			},
			{
				Name:        "inline_policies_std",
				Description: "Inline policies in canonical form for the group.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.InlinePolicies").Transform(inlinePoliciesToStd),
			},
			{
				Name:        "attached_policy_arns",
				Description: "A list of managed policies attached to the group.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.AttachedPolicyArns"),
			},
			{
				Name:        "users",
				Description: "A list of users in the group.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Users"),
			},
			/// Standard columns for all tables
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Group.GroupName"),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Group.Arn").Transform(arnToAkas),
			},
		}),
	}
}

//// LIST FUNCTION

//// HYDRATE FUNCTIONS
