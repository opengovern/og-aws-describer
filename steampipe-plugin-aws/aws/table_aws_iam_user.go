package aws

import (
	"context"

	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsIamUser(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_iam_user",
		Description: "AWS IAM User",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AnyColumn([]string{"name", "arn"}),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"ValidationError", "NoSuchEntity", "InvalidParameter"}),
			},
			Hydrate: kaytu.GetIAMUser,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListIAMUser,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "path", Require: plugin.Optional},
			},
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The friendly name identifying the user.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.User.UserName"),
			},
			{
				Name:        "user_id",
				Description: "The stable and unique string identifying the user.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.User.UserId"),
			},
			{
				Name:        "path",
				Description: "The path to the user.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.User.Path"),
			},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) that identifies the user.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.User.Arn"),
			},
			{
				Name:        "create_date",
				Description: "The date and time, when the user was created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.User.CreateDate"),
			},
			{
				Name:        "password_last_used",
				Description: "The date and time, when the user's password was last used to sign in to an AWS website.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.User.PasswordLastUsed"),
			},
			{
				Name:        "permissions_boundary_arn",
				Description: "The ARN of the policy used to set the permissions boundary for the user.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.User.PermissionsBoundary.PermissionsBoundaryArn"),
			},
			{
				Name: "permissions_boundary_type",
				Description: "The permissions boundary usage type that indicates what type of IAM resource " +
					"is used as the permissions boundary for an entity. This data type can only have a value of Policy.",
				Type:      proto.ColumnType_STRING,
				Transform: transform.FromField("Description.User.PermissionsBoundary.PermissionsBoundaryType"),
			},
			{
				Name:        "mfa_enabled",
				Description: "The MFA status of the user.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.From(userMfaStatus),
			},
			{
				Name:        "login_profile",
				Description: "Contains the user name and password create date for a user.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.LoginProfile"),
			},
			{
				Name:        "mfa_devices",
				Description: "A list of MFA devices attached to the user.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.MFADevices"),
			},
			{
				Name:        "groups",
				Description: "A list of groups attached to the user.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Groups"),
			},
			{
				Name:        "inline_policies",
				Description: "A list of policy documents that are embedded as inline policies for the user.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.InlinePolicies"),
			},
			{
				Name:        "inline_policies_std",
				Description: "Inline policies in canonical form for the user.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.InlinePolicies").Transform(inlinePoliciesToStd),
			},
			{
				Name:        "attached_policy_arns",
				Description: "A list of managed policies attached to the user.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.AttachedPolicyArns"),
			},
			{
				Name:        "tags_src",
				Description: "A list of tags that are attached to the user.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.User.Tags"),
			},
			// Standard columns for all tables
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(getAwsIamTurboTags),
			},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.User.UserName"),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.User.Arn").Transform(arnToAkas),
			},
		}),
	}
}

//// LIST FUNCTION

//// HYDRATE FUNCTIONS

//// TRANSFORM FUNCTION

func getAwsIamTurboTags(_ context.Context, d *transform.TransformData) (interface{}, error) {
	tags := d.HydrateItem.(kaytu.IAMUser).Description.User.Tags

	var turbotTags map[string]string
	if tags != nil {
		turbotTags = map[string]string{}
		for _, t := range tags {
			turbotTags[*t.Key] = *t.Value
		}
	}

	return turbotTags, nil
}

func userMfaStatus(_ context.Context, d *transform.TransformData) (interface{}, error) {
	data := d.HydrateItem.(kaytu.IAMUser).Description
	if data.MFADevices != nil && len(data.MFADevices) > 0 {
		return true, nil
	}

	return false, nil
}
