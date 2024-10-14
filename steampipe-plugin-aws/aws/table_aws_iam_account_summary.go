package aws

import (
	"context"
	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

// error
// check this that is right or not .
func tableAwsIamAccountSummary(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:             "aws_iam_account_summary",
		Description:      "AWS IAM Account Summary",
		DefaultTransform: transform.FromGo(),
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListIAMAccountSummary,
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "access_keys_per_user_quota",
				Description: "Specifies the allowed quota of access keys per user.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.AccountSummary.AccessKeysPerUserQuota"),
			},
			{
				Name:        "account_access_keys_present",
				Description: "Specifies the number of account level access keys present.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.AccountSummary.AccountAccessKeysPresent"),
			},
			{
				Name:        "account_mfa_enabled",
				Description: "Specifies whether MFA is enabled for the account.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.AccountSummary.AccountMFAEnabled").Transform(mfaEnabledToBool),
			},
			{
				Name:        "account_signing_certificates_present",
				Description: "Specifies the number of account signing certificates present.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.AccountSummary.AccountSigningCertificatesPresent"),
			},
			{
				Name:        "assume_role_policy_size_quota",
				Description: "Specifies the allowed assume role policy size.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.AccountSummary.AssumeRolePolicySizeQuota"),
			},
			{
				Name:        "attached_policies_per_group_quota",
				Description: "Specifies the allowed attached policies per group.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.AccountSummary.AttachedPoliciesPerGroupQuota"),
			},
			{
				Name:        "attached_policies_per_role_quota",
				Description: "Specifies the allowed attached policies per role.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.AccountSummary.AttachedPoliciesPerRoleQuota"),
			},
			{
				Name:        "attached_policies_per_user_quota",
				Description: "Specifies the allowed attached policies per user.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.AccountSummary.AttachedPoliciesPerUserQuota"),
			},
			{
				Name:        "global_endpoint_token_version",
				Description: "Specifies the token version of the global endpoint.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.AccountSummary.GlobalEndpointTokenVersion"),
			},
			{
				Name:        "group_policy_size_quota",
				Description: "Specifies the allowed group policy size.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.AccountSummary.GroupPolicySizeQuota"),
			},
			{
				Name:        "groups",
				Description: "Specifies the number of groups.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.AccountSummary.Groups"),
			},
			{
				Name:        "groups_per_user_quota",
				Description: "Specifies the allowed number of groups.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.AccountSummary.GroupsPerUserQuota"),
			},
			{
				Name:        "groups_quota",
				Description: "Specifies the allowed number of groups.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.AccountSummary.GroupsQuota"),
			},
			{
				Name:        "instance_profiles",
				Description: "Specifies the number of groups.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.AccountSummary.InstanceProfiles"),
			},
			{
				Name:        "instance_profiles_quota",
				Description: "Specifies the allowed number of groups.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.AccountSummary.InstanceProfilesQuota"),
			},
			{
				Name:        "mfa_devices",
				Description: "Specifies the number of MFA devices.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.AccountSummary.MFADevices"),
			},
			{
				Name:        "mfa_devices_in_use",
				Description: "Specifies the number of MFA devices in use.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.AccountSummary.MFADevicesInUse"),
			},
			{
				Name:        "policies",
				Description: "Specifies the number of policies.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.AccountSummary.Policies"),
			},
			{
				Name:        "policies_quota",
				Description: "Specifies the allowed number of policies.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.AccountSummary.PoliciesQuota"),
			},
			{
				Name:        "policy_size_quota",
				Description: "Specifies the allowed size of policies.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.AccountSummary.PolicySizeQuota"),
			},
			{
				Name:        "policy_versions_in_use",
				Description: "Specifies the number of policy versions in use.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.AccountSummary.PolicyVersionsInUse"),
			},
			{
				Name:        "policy_versions_in_use_quota",
				Description: "Specifies the allowed number of policy versions.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.AccountSummary.PolicyVersionsInUseQuota"),
			},
			{
				Name:        "providers",
				Description: "Specifies the number of providers.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.AccountSummary.Providers"),
			},
			{
				Name:        "role_policy_size_quota",
				Description: "Specifies the allowed role policy size.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.AccountSummary.RolePolicySizeQuota"),
			},
			{
				Name:        "roles",
				Description: "Specifies the number of roles.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.AccountSummary.Roles"),
			},
			{
				Name:        "roles_quota",
				Description: "Specifies the allowed number of roles.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.AccountSummary.RolesQuota"),
			},
			{
				Name:        "server_certificates",
				Description: "Specifies the number of server certificates.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.AccountSummary.ServerCertificates"),
			},
			{
				Name:        "server_certificates_quota",
				Description: "Specifies the allowed number of server certificates.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.AccountSummary.ServerCertificatesQuota"),
			},
			{
				Name:        "signing_certificates_per_user_quota",
				Description: "Specifies the allowed number of signing certificates per user.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.AccountSummary.SigningCertificatesPerUserQuota"),
			}, {
				Name:        "user_policy_size_quota",
				Description: "Specifies the allowed user policy size.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.AccountSummary.UserPolicySizeQuota"),
			}, {
				Name:        "users",
				Description: "Specifies the number of users.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.AccountSummary.Users"),
			}, {
				Name:        "users_quota",
				Description: "Specifies the allowed number of users.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.AccountSummary.UsersQuota"),
			}, {
				Name:        "versions_per_policy_quota",
				Description: "Specifies the allowed number of versions per policy.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.AccountSummary.VersionsPerPolicyQuota"),
			},
		}),
	}
}

//// LIST FUNCTION

//// TRANFROM FUNCTION

func mfaEnabledToBool(_ context.Context, d *transform.TransformData) (interface{}, error) {
	return d.Value.(int32) == 1, nil
}
