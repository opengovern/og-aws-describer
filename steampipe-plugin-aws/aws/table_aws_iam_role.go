package aws

import (
	"context"

	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsIamRole(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_iam_role",
		Description: "AWS IAM Role",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AnyColumn([]string{"name", "arn"}),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"ValidationError", "NoSuchEntity", "InvalidParameter"}),
			},
			Hydrate: opengovernance.GetIAMRole,
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListIAMRole,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "path", Require: plugin.Optional},
			},
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			// "Key" Columns
			{
				Name:        "name",
				Description: "The friendly name that identifies the role.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Role.RoleName"),
			},
			{
				Name:        "arn",
				Type:        proto.ColumnType_STRING,
				Description: "The Amazon Resource Name (ARN) specifying the role.",
				Transform:   transform.FromField("Description.Role.Arn"),
			},
			{
				Name:        "role_id",
				Type:        proto.ColumnType_STRING,
				Description: "The stable and unique string identifying the role.",
				Transform:   transform.FromField("Description.Role.RoleId"),
			},
			// Other Columns
			{
				Name:        "create_date",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The date and time when the role was created.",
				Transform:   transform.FromField("Description.Role.CreateDate"),
			},
			{
				Name:        "description",
				Type:        proto.ColumnType_STRING,
				Description: "A user-provided description of the role.",
				Transform:   transform.FromField("Description.Role.Description"),
			},
			{
				Name:        "instance_profile_arns",
				Description: "A list of instance profiles associated with the role.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.InstanceProfileArns"),
			},
			{
				Name:        "max_session_duration",
				Description: "The maximum session duration (in seconds) for the specified role. Anyone who uses the AWS CLI, or API to assume the role can specify the duration using the optional DurationSeconds API parameter or duration-seconds CLI parameter.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Role.MaxSessionDuration"),
			},
			{
				Name:        "path",
				Description: "The path to the role.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Role.Path"),
			},
			{
				Name:        "permissions_boundary_arn",
				Description: "The ARN of the policy used to set the permissions boundary for the role.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Role.PermissionsBoundary.PermissionsBoundaryArn"),
			},
			{
				Name: "permissions_boundary_type",
				Description: "The permissions boundary usage type that indicates what type of IAM resource " +
					"is used as the permissions boundary for an entity. This data type can only have a value of Policy.",
				Type:      proto.ColumnType_STRING,
				Transform: transform.FromField("Description.Role.PermissionsBoundary.PermissionsBoundaryType"),
			},
			{
				Name: "role_last_used_date",
				Type: proto.ColumnType_TIMESTAMP,
				Description: "Contains information about the last time that an IAM role was used. " +
					"Activity is only reported for the trailing 400 days. This period can be " +
					"shorter if your Region began supporting these features within the last year. " +
					"The role might have been used more than 400 days ago.",
				Transform: transform.FromField("Description.Role.RoleLastUsed.LastUsedDate"),
			},
			{
				Name:        "role_last_used_region",
				Description: "Contains the region in which the IAM role was used.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Role.RoleLastUsed.Region"),
			},
			{
				Name:        "tags_src",
				Description: "A list of tags that are attached to the role.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Role.Tags"),
			},
			{
				Name:        "inline_policies",
				Description: "A list of policy documents that are embedded as inline policies for the role..",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.InlinePolicies"),
			},
			{
				Name:        "inline_policies_std",
				Description: "Inline policies in canonical form for the role.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.InlinePolicies").Transform(inlinePoliciesToStd),
			},
			{
				Name:        "attached_policy_arns",
				Description: "A list of managed policies attached to the role.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.AttachedPolicyArns"),
			},
			{
				Name:        "assume_role_policy",
				Description: "The policy that grants an entity permission to assume the role.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Role.AssumeRolePolicyDocument").Transform(transform.UnmarshalYAML),
			},
			{
				Name:        "assume_role_policy_std",
				Description: "Contains the assume role policy in a canonical form for easier searching.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Role.AssumeRolePolicyDocument").Transform(unescape).Transform(policyToCanonical),
			},

			// Standard columns for all tables
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Role.RoleName"),
			},

			// Standard columns for all tables
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(getIamRoleTurbotTags),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Role.Arn").Transform(arnToAkas),
			},
		}),
	}
}

//// LIST FUNCTION

//// HYDRATE FUNCTIONS

//// TRANSFORM FUNCTIONS

func getIamRoleTurbotTags(_ context.Context, d *transform.TransformData) (interface{}, error) {
	data := d.HydrateItem.(opengovernance.IAMRole).Description.Role
	var turbotTagsMap map[string]string

	if data.Tags != nil {
		turbotTagsMap = map[string]string{}
		for _, i := range data.Tags {
			turbotTagsMap[*i.Key] = *i.Value
		}
	}
	return turbotTagsMap, nil
}
