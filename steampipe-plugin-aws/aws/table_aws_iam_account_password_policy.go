package aws

import (
	"context"

	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//type awsIamPolicySimulatorResult struct {
//	PrincipalArn string
//	Action       string
//	ResourceArn  string
//	Decision     *string
//	Result       *iam.EvaluationResult
//}

func tableAwsIamAccountPasswordPolicy(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_iam_account_password_policy",
		Description: "AWS IAM Account Password Policy",
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListIAMAccountPasswordPolicy,
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "allow_users_to_change_password",
				Description: "Specifies whether IAM users are allowed to change their own password.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.PasswordPolicy.AllowUsersToChangePassword"),
			},
			{
				Name:        "expire_passwords",
				Description: "Indicates whether passwords in the account expire. Returns true if MaxPasswordAge contains a value greater than 0. Returns false if MaxPasswordAge is 0 or not present.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.PasswordPolicy.ExpirePasswords"),
			},
			{
				Name:        "hard_expiry",
				Description: "Specifies whether IAM users are prevented from setting a new password after.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.PasswordPolicy.HardExpiry"),
			},
			{
				Name:        "max_password_age",
				Description: "The number of days that an IAM user password is valid.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.PasswordPolicy.MaxPasswordAge"),
			},
			{
				Name:        "minimum_password_length",
				Description: "Minimum length to require for IAM user passwords.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.PasswordPolicy.MinimumPasswordLength"),
			},
			{
				Name:        "password_reuse_prevention",
				Description: "Specifies the number of previous passwords that IAM users are prevented from reusing.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.PasswordPolicy.PasswordReusePrevention"),
			},
			{
				Name:        "require_lowercase_characters",
				Description: "Specifies whether to require lowercase characters for IAM user passwords.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.PasswordPolicy.RequireLowercaseCharacters"),
			},
			{
				Name:        "require_numbers",
				Description: "Specifies whether to require numbers for IAM user passwords.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.PasswordPolicy.RequireNumbers"),
			},
			{
				Name:        "require_symbols",
				Description: "Specifies whether to require symbols for IAM user passwords.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.PasswordPolicy.RequireSymbols"),
			},
			{
				Name:        "require_uppercase_characters",
				Description: "Specifies whether to require uppercase characters for IAM user passwords.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.PasswordPolicy.RequireUppercaseCharacters"),
			},
		}),
	}
}

//// LIST FUNCTION
