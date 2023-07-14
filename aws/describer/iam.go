package describer

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/aws/aws-sdk-go-v2/service/organizations"
	"github.com/gocarina/gocsv"
	"github.com/kaytu-io/kaytu-aws-describer/aws/model"
)

const (
	organizationsNotInUseException = "AWSOrganizationsNotInUseException"
)

func IAMAccount(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	orgClient := organizations.NewFromConfig(cfg)

	accountId, err := STSAccount(ctx, cfg)
	if err != nil {
		return nil, err
	}

	output, err := orgClient.DescribeOrganization(ctx, &organizations.DescribeOrganizationInput{})
	if err != nil {
		if isErr(err, organizationsNotInUseException) {
			output = &organizations.DescribeOrganizationOutput{}
		} else {
			return nil, err
		}
	}

	client := iam.NewFromConfig(cfg)
	paginator := iam.NewListAccountAliasesPaginator(client, &iam.ListAccountAliasesInput{})

	var aliases []string
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		aliases = append(aliases, page.AccountAliases...)
	}

	var values []Resource
	resource := Resource{
		Region: describeCtx.KaytuRegion,
		// No ID or ARN. Per Account Configuration
		Name: accountId,
		Description: model.IAMAccountDescription{
			Aliases:      aliases,
			Organization: output.Organization,
		},
	}
	if stream != nil {
		if err := (*stream)(resource); err != nil {
			return nil, err
		}
	} else {
		values = append(values, resource)
	}

	return values, nil
}

func IAMAccountSummary(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := iam.NewFromConfig(cfg)
	output, err := client.GetAccountSummary(ctx, &iam.GetAccountSummaryInput{})
	if err != nil {
		return nil, err
	}

	desc := model.IAMAccountSummaryDescription{
		AccountSummary: model.AccountSummary{
			AccountMFAEnabled:                 output.SummaryMap["AccountMFAEnabled"],
			AccessKeysPerUserQuota:            output.SummaryMap["AccessKeysPerUserQuota"],
			AccountAccessKeysPresent:          output.SummaryMap["AccountAccessKeysPresent"],
			AccountSigningCertificatesPresent: output.SummaryMap["AccountSigningCertificatesPresent"],
			AssumeRolePolicySizeQuota:         output.SummaryMap["AssumeRolePolicySizeQuota"],
			AttachedPoliciesPerGroupQuota:     output.SummaryMap["AttachedPoliciesPerGroupQuota"],
			AttachedPoliciesPerRoleQuota:      output.SummaryMap["AttachedPoliciesPerRoleQuota"],
			AttachedPoliciesPerUserQuota:      output.SummaryMap["AttachedPoliciesPerUserQuota"],
			GlobalEndpointTokenVersion:        output.SummaryMap["GlobalEndpointTokenVersion"],
			GroupPolicySizeQuota:              output.SummaryMap["GroupPolicySizeQuota"],
			Groups:                            output.SummaryMap["Groups"],
			GroupsPerUserQuota:                output.SummaryMap["GroupsPerUserQuota"],
			GroupsQuota:                       output.SummaryMap["GroupsQuota"],
			InstanceProfiles:                  output.SummaryMap["InstanceProfiles"],
			InstanceProfilesQuota:             output.SummaryMap["InstanceProfilesQuota"],
			MFADevices:                        output.SummaryMap["MFADevices"],
			MFADevicesInUse:                   output.SummaryMap["MFADevicesInUse"],
			Policies:                          output.SummaryMap["Policies"],
			PoliciesQuota:                     output.SummaryMap["PoliciesQuota"],
			PolicySizeQuota:                   output.SummaryMap["PolicySizeQuota"],
			PolicyVersionsInUse:               output.SummaryMap["PolicyVersionsInUse"],
			PolicyVersionsInUseQuota:          output.SummaryMap["PolicyVersionsInUseQuota"],
			Providers:                         output.SummaryMap["Providers"],
			RolePolicySizeQuota:               output.SummaryMap["RolePolicySizeQuota"],
			Roles:                             output.SummaryMap["Roles"],
			RolesQuota:                        output.SummaryMap["RolesQuota"],
			ServerCertificates:                output.SummaryMap["ServerCertificates"],
			ServerCertificatesQuota:           output.SummaryMap["ServerCertificatesQuota"],
			SigningCertificatesPerUserQuota:   output.SummaryMap["SigningCertificatesPerUserQuota"],
			UserPolicySizeQuota:               output.SummaryMap["UserPolicySizeQuota"],
			Users:                             output.SummaryMap["Users"],
			UsersQuota:                        output.SummaryMap["UsersQuota"],
			VersionsPerPolicyQuota:            output.SummaryMap["VersionsPerPolicyQuota"],
		},
	}

	accountId, err := STSAccount(ctx, cfg)
	if err != nil {
		return nil, err
	}

	var values []Resource
	resource := Resource{
		Region: describeCtx.KaytuRegion,
		// No ID or ARN. Per Account Configuration
		Name:        accountId + " Account Summary",
		Description: desc,
	}

	if stream != nil {
		if err := (*stream)(resource); err != nil {
			return nil, err
		}
	} else {
		values = append(values, resource)
	}

	return values, nil
}

func IAMAccountPasswordPolicy(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	var values []Resource
	resource, err := IAMAccountPasswordPolicyHandle(ctx, cfg)
	if err != nil {
		return nil, err
	}
	if stream != nil {
		if err := (*stream)(resource); err != nil {
			return nil, err
		}
	} else {
		values = append(values, resource)
	}
	return values, nil
}
func IAMAccountPasswordPolicyHandle(ctx context.Context, cfg aws.Config) (Resource, error) {
	client := iam.NewFromConfig(cfg)
	describeCtx := GetDescribeContext(ctx)
	output, err := client.GetAccountPasswordPolicy(ctx, &iam.GetAccountPasswordPolicyInput{})
	if err != nil {
		if !isErr(err, "NoSuchEntity") {
			return Resource{}, err
		}

		output = &iam.GetAccountPasswordPolicyOutput{}
	}

	if output.PasswordPolicy == nil {
		return Resource{}, nil
	}

	accountId, err := STSAccount(ctx, cfg)
	if err != nil {
		return Resource{}, err
	}

	resource := Resource{
		Region: describeCtx.KaytuRegion,
		// No ID or ARN. Per Account Configuration
		Name: accountId + " IAM Password Policy",
		Description: model.IAMAccountPasswordPolicyDescription{
			PasswordPolicy: *output.PasswordPolicy,
		},
	}
	return resource, nil
}
func GetIAMAccountPasswordPolicy(ctx context.Context, cfg aws.Config, fields map[string]string) ([]Resource, error) {
	var values []Resource
	resource, err := IAMAccountPasswordPolicyHandle(ctx, cfg)
	if err != nil {
		return nil, err
	}
	values = append(values, resource)
	return values, nil
}

func IAMAccessKey(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	client := iam.NewFromConfig(cfg)
	paginator := iam.NewListAccessKeysPaginator(client, &iam.ListAccessKeysInput{})

	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, v := range page.AccessKeyMetadata {
			resource := iAMAccessKeyHandel(ctx, v)
			if stream != nil {
				if err := (*stream)(resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, resource)
			}
		}
	}

	return values, nil
}
func iAMAccessKeyHandel(ctx context.Context, v types.AccessKeyMetadata) Resource {
	describeCtx := GetDescribeContext(ctx)
	arn := "arn:" + describeCtx.Partition + ":iam::" + describeCtx.AccountID + ":user/" + *v.UserName + "/accesskey/" + *v.AccessKeyId
	resource := Resource{
		Region: describeCtx.KaytuRegion,
		ARN:    arn,
		Name:   *v.UserName,
		Description: model.IAMAccessKeyDescription{
			AccessKey: v,
		},
	}
	return resource
}
func GetIAMAccessKey(ctx context.Context, cfg aws.Config, fields map[string]string) ([]Resource, error) {
	userName := fields["name"]
	var value []Resource
	client := iam.NewFromConfig(cfg)

	accessKeys, err := client.ListAccessKeys(ctx, &iam.ListAccessKeysInput{
		UserName: &userName,
	})
	if err != nil {
		if isErr(err, "ListAccessKeysNotFound") || isErr(err, "InvalidParameterValue") {
			return nil, nil
		}
		return nil, err
	}

	for _, v := range accessKeys.AccessKeyMetadata {
		resource := iAMAccessKeyHandel(ctx, v)
		value = append(value, resource)
	}
	return value, nil
}

func IAMCredentialReport(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := iam.NewFromConfig(cfg)
	output, err := client.GetCredentialReport(ctx, &iam.GetCredentialReportInput{})
	if err != nil {
		if isErr(err, (&types.CredentialReportNotPresentException{}).ErrorCode()) ||
			isErr(err, (&types.CredentialReportExpiredException{}).ErrorCode()) ||
			isErr(err, (&types.CredentialReportNotPresentException{}).ErrorCode()) {
			return nil, nil
		}
		return nil, err
	}

	reports := []model.CredentialReport{}
	if err := gocsv.UnmarshalString(string(output.Content), &reports); err != nil {
		return nil, err
	}

	var values []Resource
	for _, report := range reports {
		report.GeneratedTime = output.GeneratedTime
		resource := Resource{
			Region: describeCtx.KaytuRegion,
			ID:     report.UserName, // Unique report entry per user
			Name:   report.UserName + " Credential Report",
			Description: model.IAMCredentialReportDescription{
				CredentialReport: report,
			},
		}
		if stream != nil {
			if err := (*stream)(resource); err != nil {
				return nil, err
			}
		} else {
			values = append(values, resource)
		}

	}

	return values, nil
}

func IAMPolicy(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	client := iam.NewFromConfig(cfg)
	paginator := iam.NewListPoliciesPaginator(client, &iam.ListPoliciesInput{
		OnlyAttached: false,
		Scope:        types.PolicyScopeTypeAll,
	})

	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, v := range page.Policies {
			version, err := client.GetPolicyVersion(ctx, &iam.GetPolicyVersionInput{
				PolicyArn: v.Arn,
				VersionId: v.DefaultVersionId,
			})
			if err != nil {
				return nil, err
			}

			resource := iAMPolicyHandel(ctx, v, version)
			if stream != nil {
				if err := (*stream)(resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, resource)
			}
		}
	}
	return values, nil
}
func iAMPolicyHandel(ctx context.Context, v types.Policy, version *iam.GetPolicyVersionOutput) Resource {
	describeCtx := GetDescribeContext(ctx)
	resource := Resource{
		Region: describeCtx.KaytuRegion,
		ARN:    *v.Arn,
		Name:   *v.PolicyName,
		Description: model.IAMPolicyDescription{
			Policy:        v,
			PolicyVersion: *version.PolicyVersion,
		},
	}
	return resource
}
func GetIAMPolicy(ctx context.Context, cfg aws.Config, fields map[string]string) ([]Resource, error) {
	arn := fields["arn"]
	client := iam.NewFromConfig(cfg)
	out, err := client.GetPolicy(ctx, &iam.GetPolicyInput{PolicyArn: &arn})
	if err != nil {
		return nil, err
	}
	v := out.Policy

	var values []Resource
	version, err := client.GetPolicyVersion(ctx, &iam.GetPolicyVersionInput{
		PolicyArn: v.Arn,
		VersionId: v.DefaultVersionId,
	})
	if err != nil {
		return nil, err
	}

	resource := iAMPolicyHandel(ctx, *v, version)
	values = append(values, resource)

	return values, nil
}

func IAMGroup(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	client := iam.NewFromConfig(cfg)
	paginator := iam.NewListGroupsPaginator(client, &iam.ListGroupsInput{})

	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, v := range page.Groups {
			users, err := getGroupUsers(ctx, client, v.GroupName)
			if err != nil {
				return nil, err
			}

			policies, err := getGroupPolicies(ctx, client, v.GroupName)
			if err != nil {
				return nil, err
			}

			aPolicies, err := getGroupAttachedPolicyArns(ctx, client, v.GroupName)
			if err != nil {
				return nil, err
			}

			resource := iAMGroupHandel(ctx, v, aPolicies, policies, users)
			if stream != nil {
				if err := (*stream)(resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, resource)
			}
		}
	}

	return values, nil
}
func iAMGroupHandel(ctx context.Context, v types.Group, aPolicies []string, policies []model.InlinePolicy, users []types.User) Resource {
	describeCtx := GetDescribeContext(ctx)
	resource := Resource{
		Region: describeCtx.KaytuRegion,
		ARN:    *v.Arn,
		Name:   *v.GroupName,
		Description: model.IAMGroupDescription{
			Group:              v,
			Users:              users,
			InlinePolicies:     policies,
			AttachedPolicyArns: aPolicies,
		},
	}
	return resource
}
func GetIAMGroup(ctx context.Context, cfg aws.Config, fields map[string]string) ([]Resource, error) {
	var value []Resource
	groupName := fields["name"]
	client := iam.NewFromConfig(cfg)
	groupOut, err := client.GetGroup(ctx, &iam.GetGroupInput{
		GroupName: &groupName,
	})
	v := groupOut.Group
	if err != nil {
		if isErr(err, "GetGroupNotFound") || isErr(err, "InvalidParameterValue") {
			return nil, nil
		}
		return nil, err
	}

	users, err := getGroupUsers(ctx, client, v.GroupName)
	if err != nil {
		if isErr(err, "getGroupUsersNotFound") || isErr(err, "InvalidParameterValue") {
			return nil, nil
		}
		return nil, err
	}

	policies, err := getGroupPolicies(ctx, client, v.GroupName)
	if err != nil {
		if isErr(err, "getGroupPoliciesNotFound") || isErr(err, "InvalidParameterValue") {
			return nil, nil
		}
		return nil, err
	}

	aPolicies, err := getGroupAttachedPolicyArns(ctx, client, v.GroupName)
	if err != nil {
		if isErr(err, "getGroupAttachedPolicyArnsNotFound") || isErr(err, "InvalidParameterValue") {
			return nil, nil
		}
		return nil, err
	}

	resource := iAMGroupHandel(ctx, *v, aPolicies, policies, users)
	value = append(value, resource)
	return value, nil
}

func getGroupUsers(ctx context.Context, client *iam.Client, groupname *string) ([]types.User, error) {
	paginator := iam.NewGetGroupPaginator(client, &iam.GetGroupInput{
		GroupName: groupname,
	})

	var users []types.User
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		users = append(users, page.Users...)
	}

	return users, nil
}
func getGroupPolicies(ctx context.Context, client *iam.Client, groupname *string) ([]model.InlinePolicy, error) {
	paginator := iam.NewListGroupPoliciesPaginator(client, &iam.ListGroupPoliciesInput{
		GroupName: groupname,
	})

	var policies []model.InlinePolicy
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, p := range page.PolicyNames {
			output, err := client.GetGroupPolicy(ctx, &iam.GetGroupPolicyInput{
				PolicyName: aws.String(p),
				GroupName:  groupname,
			})
			if err != nil {
				return nil, err
			}

			policies = append(policies, model.InlinePolicy{
				PolicyName:     *output.PolicyName,
				PolicyDocument: *output.PolicyDocument,
			})
		}
	}

	return policies, nil
}
func getGroupAttachedPolicyArns(ctx context.Context, client *iam.Client, groupname *string) ([]string, error) {
	paginator := iam.NewListAttachedGroupPoliciesPaginator(client, &iam.ListAttachedGroupPoliciesInput{
		GroupName: groupname,
	})

	var arns []string
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, p := range page.AttachedPolicies {
			arns = append(arns, *p.PolicyArn)

		}
	}

	return arns, nil
}

func IAMInstanceProfile(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := iam.NewFromConfig(cfg)
	paginator := iam.NewListInstanceProfilesPaginator(client, &iam.ListInstanceProfilesInput{})

	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, v := range page.InstanceProfiles {
			resource := Resource{
				Region:      describeCtx.Region,
				ARN:         *v.Arn,
				Name:        *v.InstanceProfileName,
				Description: v,
			}
			if stream != nil {
				if err := (*stream)(resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, resource)
			}
		}
	}

	return values, nil
}

func IAMManagedPolicy(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := iam.NewFromConfig(cfg)
	paginator := iam.NewListPoliciesPaginator(client, &iam.ListPoliciesInput{
		OnlyAttached: true,
	})

	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, v := range page.Policies {
			resource := Resource{
				Region:      describeCtx.Region,
				ARN:         *v.Arn,
				Name:        *v.PolicyName,
				Description: v,
			}
			if stream != nil {
				if err := (*stream)(resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, resource)
			}
		}
	}

	return values, nil
}

func IAMOIDCProvider(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := iam.NewFromConfig(cfg)
	output, err := client.ListOpenIDConnectProviders(ctx, &iam.ListOpenIDConnectProvidersInput{})
	if err != nil {
		return nil, err
	}

	var values []Resource
	for _, v := range output.OpenIDConnectProviderList {
		resource := Resource{
			Region:      describeCtx.Region,
			ARN:         *v.Arn,
			Name:        *v.Arn,
			Description: v,
		}
		if stream != nil {
			if err := (*stream)(resource); err != nil {
				return nil, err
			}
		} else {
			values = append(values, resource)
		}
	}

	return values, nil
}

func IAMGroupPolicy(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	groups, err := IAMGroup(ctx, cfg, nil)
	if err != nil {
		return nil, err
	}

	client := iam.NewFromConfig(cfg)

	var values []Resource
	for _, g := range groups {
		group := g.Description.(model.IAMGroupDescription).Group
		err := PaginateRetrieveAll(func(prevToken *string) (nextToken *string, err error) {
			output, err := client.ListGroupPolicies(ctx, &iam.ListGroupPoliciesInput{
				GroupName: group.GroupName,
				Marker:    prevToken,
			})
			if err != nil {
				return nil, err
			}

			for _, policy := range output.PolicyNames {
				v, err := client.GetGroupPolicy(ctx, &iam.GetGroupPolicyInput{
					GroupName:  group.GroupName,
					PolicyName: aws.String(policy),
				})
				if err != nil {
					return nil, err
				}

				resource := Resource{
					Region:      describeCtx.Region,
					ID:          CompositeID(*v.GroupName, *v.PolicyName),
					Name:        *v.GroupName,
					Description: v,
				}
				if stream != nil {
					if err := (*stream)(resource); err != nil {
						return nil, err
					}
				} else {
					values = append(values, resource)
				}

			}

			return output.Marker, nil
		})
		if err != nil {
			return nil, err
		}
	}

	return values, nil
}

func IAMUserPolicy(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	users, err := IAMUser(ctx, cfg, nil)
	if err != nil {
		return nil, err
	}

	client := iam.NewFromConfig(cfg)

	var values []Resource
	for _, u := range users {
		user := u.Description.(model.IAMUserDescription).User
		err := PaginateRetrieveAll(func(prevToken *string) (nextToken *string, err error) {
			output, err := client.ListUserPolicies(ctx, &iam.ListUserPoliciesInput{
				UserName: user.UserName,
				Marker:   prevToken,
			})
			if err != nil {
				return nil, err
			}

			for _, policy := range output.PolicyNames {
				v, err := client.GetUserPolicy(ctx, &iam.GetUserPolicyInput{
					UserName:   user.UserName,
					PolicyName: aws.String(policy),
				})
				if err != nil {
					return nil, err
				}

				resource := Resource{
					Region:      describeCtx.Region,
					ID:          CompositeID(*v.UserName, *v.PolicyName),
					Name:        *v.UserName,
					Description: v,
				}
				if stream != nil {
					if err := (*stream)(resource); err != nil {
						return nil, err
					}
				} else {
					values = append(values, resource)
				}
			}

			return output.Marker, nil
		})
		if err != nil {
			return nil, err
		}
	}

	return values, nil
}

func IAMRolePolicy(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	roles, err := IAMRole(ctx, cfg, nil)
	if err != nil {
		return nil, err
	}

	client := iam.NewFromConfig(cfg)

	var values []Resource

	for _, r := range roles {
		role := r.Description.(model.IAMRoleDescription).Role
		err := PaginateRetrieveAll(func(prevToken *string) (nextToken *string, err error) {
			output, err := client.ListRolePolicies(ctx, &iam.ListRolePoliciesInput{
				RoleName: role.RoleName,
				Marker:   prevToken,
			})
			if err != nil {
				return nil, err
			}

			for _, policy := range output.PolicyNames {
				v, err := client.GetRolePolicy(ctx, &iam.GetRolePolicyInput{
					RoleName:   role.RoleName,
					PolicyName: aws.String(policy),
				})
				if err != nil {
					return nil, err
				}

				resource := Resource{
					Region:      describeCtx.Region,
					ID:          CompositeID(*v.RoleName, *v.PolicyName),
					Name:        *v.RoleName,
					Description: v,
				}
				if stream != nil {
					if err := (*stream)(resource); err != nil {
						return nil, err
					}
				} else {
					values = append(values, resource)
				}
			}

			return output.Marker, nil
		})
		if err != nil {
			return nil, err
		}
	}

	return values, nil
}

func IAMRole(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	client := iam.NewFromConfig(cfg)
	paginator := iam.NewListRolesPaginator(client, &iam.ListRolesInput{})

	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, v := range page.Roles {
			profiles, err := getRoleInstanceProfileArns(ctx, client, v.RoleName)
			if err != nil {
				return nil, err
			}

			policies, err := getRolePolicies(ctx, client, v.RoleName)
			if err != nil {
				return nil, err
			}

			aPolicies, err := getRoleAttachedPolicyArns(ctx, client, v.RoleName)
			if err != nil {
				return nil, err
			}

			resource := iAMRoleHandel(ctx, v, profiles, policies, aPolicies)
			if stream != nil {
				if err := (*stream)(resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, resource)
			}
		}
	}
	return values, nil
}
func iAMRoleHandel(ctx context.Context, v types.Role, profiles []string, policies []model.InlinePolicy, aPolicies []string) Resource {
	describeCtx := GetDescribeContext(ctx)
	resource := Resource{
		Region: describeCtx.KaytuRegion,
		ARN:    *v.Arn,
		Name:   *v.RoleName,
		Description: model.IAMRoleDescription{
			Role:                v,
			InstanceProfileArns: profiles,
			InlinePolicies:      policies,
			AttachedPolicyArns:  aPolicies,
		},
	}
	return resource
}
func GetIAMRole(ctx context.Context, cfg aws.Config, fields map[string]string) ([]Resource, error) {
	pathPrefix := fields["path"]
	client := iam.NewFromConfig(cfg)

	out, err := client.ListRoles(ctx, &iam.ListRolesInput{
		Marker:     nil,
		MaxItems:   nil,
		PathPrefix: &pathPrefix,
	})
	if err != nil {
		return nil, err
	}

	var values []Resource
	for _, v := range out.Roles {
		profiles, err := getRoleInstanceProfileArns(ctx, client, v.RoleName)
		if err != nil {
			return nil, err
		}

		policies, err := getRolePolicies(ctx, client, v.RoleName)
		if err != nil {
			return nil, err
		}

		aPolicies, err := getRoleAttachedPolicyArns(ctx, client, v.RoleName)
		if err != nil {
			return nil, err
		}

		resource := iAMRoleHandel(ctx, v, profiles, policies, aPolicies)
		values = append(values, resource)
	}
	return values, nil
}

func getRoleInstanceProfileArns(ctx context.Context, client *iam.Client, rolename *string) ([]string, error) {
	paginator := iam.NewListInstanceProfilesForRolePaginator(client, &iam.ListInstanceProfilesForRoleInput{
		RoleName: rolename,
	})

	var arns []string
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, ip := range page.InstanceProfiles {
			arns = append(arns, *ip.Arn)

		}
	}

	return arns, nil
}
func getRolePolicies(ctx context.Context, client *iam.Client, rolename *string) ([]model.InlinePolicy, error) {
	paginator := iam.NewListRolePoliciesPaginator(client, &iam.ListRolePoliciesInput{
		RoleName: rolename,
	})

	var policies []model.InlinePolicy
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, name := range page.PolicyNames {
			output, err := client.GetRolePolicy(ctx, &iam.GetRolePolicyInput{
				RoleName:   rolename,
				PolicyName: aws.String(name),
			})
			if err != nil {
				return nil, err
			}

			policies = append(policies, model.InlinePolicy{
				PolicyName:     *output.PolicyName,
				PolicyDocument: *output.PolicyDocument,
			})
		}

	}

	return policies, nil
}
func getRoleAttachedPolicyArns(ctx context.Context, client *iam.Client, rolename *string) ([]string, error) {
	paginator := iam.NewListAttachedRolePoliciesPaginator(client, &iam.ListAttachedRolePoliciesInput{
		RoleName: rolename,
	})

	var arns []string
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, p := range page.AttachedPolicies {
			arns = append(arns, *p.PolicyArn)

		}
	}

	return arns, nil
}

func IAMServerCertificate(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	client := iam.NewFromConfig(cfg)
	paginator := iam.NewListServerCertificatesPaginator(client, &iam.ListServerCertificatesInput{})

	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, v := range page.ServerCertificateMetadataList {
			output, err := client.GetServerCertificate(ctx, &iam.GetServerCertificateInput{
				ServerCertificateName: v.ServerCertificateName,
			})
			if err != nil {
				return nil, err
			}

			resource := iAMServerCertificateHandel(ctx, v, output)
			if stream != nil {
				if err := (*stream)(resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, resource)
			}
		}
	}
	return values, nil
}
func iAMServerCertificateHandel(ctx context.Context, v types.ServerCertificateMetadata, output *iam.GetServerCertificateOutput) Resource {
	describeCtx := GetDescribeContext(ctx)
	resource := Resource{
		Region: describeCtx.KaytuRegion,
		ARN:    *v.Arn,
		Name:   *v.ServerCertificateName,
		Description: model.IAMServerCertificateDescription{
			ServerCertificate: *output.ServerCertificate,
		},
	}
	return resource
}
func GetIAMServerCertificate(ctx context.Context, cfg aws.Config, fields map[string]string) ([]Resource, error) {
	var value []Resource
	pathPerfix := fields["path"]

	client := iam.NewFromConfig(cfg)
	listServiceCer, err := client.ListServerCertificates(ctx, &iam.ListServerCertificatesInput{
		PathPrefix: &pathPerfix,
	})
	if err != nil {
		if isErr(err, "ListServerCertificatesNotFound") || isErr(err, "InvalidParameterValue") {
			return nil, nil
		}
		return nil, err
	}

	for _, v := range listServiceCer.ServerCertificateMetadataList {
		output, err := client.GetServerCertificate(ctx, &iam.GetServerCertificateInput{
			ServerCertificateName: v.ServerCertificateName,
		})
		if err != nil {
			if isErr(err, "GetServerCertificateNotFound") || isErr(err, "InvalidParameterValue") {
				return nil, nil
			}
			return nil, err
		}

		resource := iAMServerCertificateHandel(ctx, v, output)
		value = append(value, resource)
	}
	return value, nil
}

func IAMUser(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	client := iam.NewFromConfig(cfg)
	paginator := iam.NewListUsersPaginator(client, &iam.ListUsersInput{})

	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, v := range page.Users {
			policies, err := getUserPolicies(ctx, client, v.UserName)
			if err != nil {
				return nil, err
			}

			aPolicies, err := getUserAttachedPolicyArns(ctx, client, v.UserName)
			if err != nil {
				return nil, err
			}

			groups, err := getUserGroups(ctx, client, v.UserName)
			if err != nil {
				return nil, err
			}

			devices, err := getUserMFADevices(ctx, client, v.UserName)
			if err != nil {
				return nil, err
			}

			resource := iAMUserHandel(ctx, v, policies, aPolicies, groups, devices)
			if stream != nil {
				if err := (*stream)(resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, resource)
			}
		}
	}

	return values, nil
}
func iAMUserHandel(ctx context.Context, v types.User, policies []model.InlinePolicy, aPolicies []string, groups []types.Group, devices []types.MFADevice) Resource {
	describeCtx := GetDescribeContext(ctx)
	resource := Resource{
		Region: describeCtx.KaytuRegion,
		ARN:    *v.Arn,
		Name:   *v.UserName,
		Description: model.IAMUserDescription{
			User:               v,
			Groups:             groups,
			InlinePolicies:     policies,
			AttachedPolicyArns: aPolicies,
			MFADevices:         devices,
		},
	}
	return resource
}
func GetIAMUser(ctx context.Context, cfg aws.Config, fields map[string]string) ([]Resource, error) {
	userName := fields["name"]

	client := iam.NewFromConfig(cfg)
	out, err := client.GetUser(ctx, &iam.GetUserInput{
		UserName: &userName,
	})
	if err != nil {
		return nil, err
	}

	v := out.User
	var values []Resource
	policies, err := getUserPolicies(ctx, client, v.UserName)
	if err != nil {
		return nil, err
	}

	aPolicies, err := getUserAttachedPolicyArns(ctx, client, v.UserName)
	if err != nil {
		return nil, err
	}

	groups, err := getUserGroups(ctx, client, v.UserName)
	if err != nil {
		return nil, err
	}

	devices, err := getUserMFADevices(ctx, client, v.UserName)
	if err != nil {
		return nil, err
	}

	resource := iAMUserHandel(ctx, *v, policies, aPolicies, groups, devices)
	values = append(values, resource)

	return values, nil
}

func IAMPolicyAttachment(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := iam.NewFromConfig(cfg)
	paginator := iam.NewListPoliciesPaginator(client, &iam.ListPoliciesInput{
		OnlyAttached: false,
		Scope:        types.PolicyScopeTypeAll,
	})

	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, policy := range page.Policies {
			attachmentPaginator := iam.NewListEntitiesForPolicyPaginator(client, &iam.ListEntitiesForPolicyInput{
				PolicyArn: policy.Arn,
			})

			var policyGroups []types.PolicyGroup
			var policyRoles []types.PolicyRole
			var policyUsers []types.PolicyUser
			for attachmentPaginator.HasMorePages() {
				attachmentPage, err := attachmentPaginator.NextPage(ctx)
				if err != nil {
					return nil, err
				}

				policyGroups = append(policyGroups, attachmentPage.PolicyGroups...)
				policyRoles = append(policyRoles, attachmentPage.PolicyRoles...)
				policyUsers = append(policyUsers, attachmentPage.PolicyUsers...)
			}
			resource := Resource{
				Region: describeCtx.KaytuRegion,
				Name:   fmt.Sprintf("%s - Attachments", *policy.Arn),
				Description: model.IAMPolicyAttachmentDescription{
					PolicyArn:             *policy.Arn,
					PolicyAttachmentCount: *policy.AttachmentCount,
					IsAttached:            *policy.AttachmentCount > 0,
					PolicyGroups:          policyGroups,
					PolicyRoles:           policyRoles,
					PolicyUsers:           policyUsers,
				},
			}
			if stream != nil {
				if err := (*stream)(resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, resource)
			}

		}
	}

	return values, nil
}
func iAMPolicyAttachmentHandel(ctx context.Context, policy types.Policy, policyGroups []types.PolicyGroup, policyRoles []types.PolicyRole, policyUsers []types.PolicyUser) Resource {
	describeCtx := GetDescribeContext(ctx)
	resource := Resource{
		Region: describeCtx.KaytuRegion,
		Name:   fmt.Sprintf("%s - Attachments", *policy.Arn),
		Description: model.IAMPolicyAttachmentDescription{
			PolicyArn:             *policy.Arn,
			PolicyAttachmentCount: *policy.AttachmentCount,
			IsAttached:            *policy.AttachmentCount > 0,
			PolicyGroups:          policyGroups,
			PolicyRoles:           policyRoles,
			PolicyUsers:           policyUsers,
		},
	}
	return resource
}
func GetIAMPolicyAttachment(ctx context.Context, cfg aws.Config, fields map[string]string) ([]Resource, error) {
	client := iam.NewFromConfig(cfg)
	policyARN := fields["arn"]
	policy, err := client.GetPolicy(ctx, &iam.GetPolicyInput{PolicyArn: &policyARN})
	if err != nil {
		return nil, err
	}

	var values []Resource
	attachmentPaginator := iam.NewListEntitiesForPolicyPaginator(client, &iam.ListEntitiesForPolicyInput{
		PolicyArn: &policyARN,
	})

	var policyGroups []types.PolicyGroup
	var policyRoles []types.PolicyRole
	var policyUsers []types.PolicyUser
	for attachmentPaginator.HasMorePages() {
		attachmentPage, err := attachmentPaginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		policyGroups = append(policyGroups, attachmentPage.PolicyGroups...)
		policyRoles = append(policyRoles, attachmentPage.PolicyRoles...)
		policyUsers = append(policyUsers, attachmentPage.PolicyUsers...)
	}
	resource := iAMPolicyAttachmentHandel(ctx, *policy.Policy, policyGroups, policyRoles, policyUsers)
	values = append(values, resource)
	return values, nil
}

func IAMSamlProvider(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	client := iam.NewFromConfig(cfg)
	output, err := client.ListSAMLProviders(ctx, &iam.ListSAMLProvidersInput{})
	if err != nil {
		return nil, err
	}

	var values []Resource
	for _, v := range output.SAMLProviderList {
		samlProvider, err := client.GetSAMLProvider(ctx, &iam.GetSAMLProviderInput{
			SAMLProviderArn: v.Arn,
		})
		if err != nil {
			return nil, err
		}

		if samlProvider.SAMLMetadataDocument != nil && len(*samlProvider.SAMLMetadataDocument) > 10000 {
			samlProvider.SAMLMetadataDocument = nil
		}

		resource := iAMSamlProviderHandel(ctx, samlProvider, *v.Arn)
		if stream != nil {
			if err := (*stream)(resource); err != nil {
				return nil, err
			}
		} else {
			values = append(values, resource)
		}
	}
	return values, nil
}
func iAMSamlProviderHandel(ctx context.Context, samlProvider *iam.GetSAMLProviderOutput, Arn string) Resource {
	describeCtx := GetDescribeContext(ctx)
	resource := Resource{
		Region: describeCtx.KaytuRegion,
		ARN:    Arn,
		Description: model.IAMSamlProviderDescription{
			SamlProvider: *samlProvider,
		},
	}
	return resource
}
func GetIAMSamlProvider(ctx context.Context, cfg aws.Config, fields map[string]string) ([]Resource, error) {
	samlProviderArn := fields["samlProviderArn"]
	var value []Resource
	client := iam.NewFromConfig(cfg)

	samlProvider, err := client.GetSAMLProvider(ctx, &iam.GetSAMLProviderInput{
		SAMLProviderArn: &samlProviderArn,
	})
	if err != nil {
		if isErr(err, "GetSAMLProviderNotFound") || isErr(err, "InvalidParameterValue") {
			return nil, nil
		}
		return nil, err
	}

	resource := iAMSamlProviderHandel(ctx, samlProvider, samlProviderArn)
	value = append(value, resource)
	return value, nil
}

func IAMServiceSpecificCredential(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	client := iam.NewFromConfig(cfg)
	paginator := iam.NewListUsersPaginator(client, &iam.ListUsersInput{})

	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, user := range page.Users {
			serviceSpecificCredentials, err := client.ListServiceSpecificCredentials(ctx, &iam.ListServiceSpecificCredentialsInput{
				UserName: user.UserName,
			})
			if err != nil {
				return nil, err
			}

			for _, credential := range serviceSpecificCredentials.ServiceSpecificCredentials {
				resource := iAMServiceSpecificCredentialHandle(ctx, credential)
				if stream != nil {
					if err := (*stream)(resource); err != nil {
						return nil, err
					}
				} else {
					values = append(values, resource)
				}
			}
		}
	}

	return values, nil
}
func iAMServiceSpecificCredentialHandle(ctx context.Context, credential types.ServiceSpecificCredentialMetadata) Resource {
	describeCtx := GetDescribeContext(ctx)
	resource := Resource{
		Region: describeCtx.KaytuRegion,
		ID:     *credential.ServiceSpecificCredentialId,
		Description: model.IAMServiceSpecificCredentialDescription{
			ServiceSpecificCredential: credential,
		},
	}
	return resource
}
func GetIAMServiceSpecificCredential(ctx context.Context, cfg aws.Config, fields map[string]string) ([]Resource, error) {
	UserName := fields["userName"]
	var value []Resource
	client := iam.NewFromConfig(cfg)

	serviceSpecificCredentials, err := client.ListServiceSpecificCredentials(ctx, &iam.ListServiceSpecificCredentialsInput{
		UserName: &UserName,
	})
	if err != nil {
		if isErr(err, "ListServiceSpecificCredentialsNotFound") || isErr(err, "InvalidParameterValue") {
			return nil, nil
		}
		return nil, err
	}

	for _, credential := range serviceSpecificCredentials.ServiceSpecificCredentials {
		resource := iAMServiceSpecificCredentialHandle(ctx, credential)
		value = append(value, resource)
	}
	return value, nil
}

func getUserPolicies(ctx context.Context, client *iam.Client, username *string) ([]model.InlinePolicy, error) {
	paginator := iam.NewListUserPoliciesPaginator(client, &iam.ListUserPoliciesInput{
		UserName: username,
	})

	var policies []model.InlinePolicy
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, p := range page.PolicyNames {
			output, err := client.GetUserPolicy(ctx, &iam.GetUserPolicyInput{
				PolicyName: aws.String(p),
				UserName:   username,
			})
			if err != nil {
				return nil, err
			}

			policies = append(policies, model.InlinePolicy{
				PolicyName:     *output.PolicyName,
				PolicyDocument: *output.PolicyDocument,
			})
		}
	}

	return policies, nil
}
func getUserAttachedPolicyArns(ctx context.Context, client *iam.Client, username *string) ([]string, error) {
	paginator := iam.NewListAttachedUserPoliciesPaginator(client, &iam.ListAttachedUserPoliciesInput{
		UserName: username,
	})

	var arns []string
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, p := range page.AttachedPolicies {
			arns = append(arns, *p.PolicyArn)

		}
	}

	return arns, nil
}
func getUserGroups(ctx context.Context, client *iam.Client, username *string) ([]types.Group, error) {
	paginator := iam.NewListGroupsForUserPaginator(client, &iam.ListGroupsForUserInput{
		UserName: username,
	})

	var groups []types.Group
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		groups = append(groups, page.Groups...)
	}

	return groups, nil
}
func getUserMFADevices(ctx context.Context, client *iam.Client, username *string) ([]types.MFADevice, error) {
	paginator := iam.NewListMFADevicesPaginator(client, &iam.ListMFADevicesInput{
		UserName: username,
	})

	var devices []types.MFADevice
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		devices = append(devices, page.MFADevices...)
	}

	return devices, nil
}

func IAMVirtualMFADevice(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	client := iam.NewFromConfig(cfg)
	output, err := client.ListVirtualMFADevices(ctx, &iam.ListVirtualMFADevicesInput{})
	if err != nil {
		return nil, err
	}

	var values []Resource
	for _, v := range output.VirtualMFADevices {
		output, err := client.ListMFADeviceTags(ctx, &iam.ListMFADeviceTagsInput{
			SerialNumber: v.SerialNumber,
		})
		if err != nil {
			output = &iam.ListMFADeviceTagsOutput{}
		}

		resource := iAMVirtualMFADeviceHandle(ctx, v, output)
		if stream != nil {
			if err := (*stream)(resource); err != nil {
				return nil, err
			}
		} else {
			values = append(values, resource)
		}
	}

	return values, nil
}
func iAMVirtualMFADeviceHandle(ctx context.Context, v types.VirtualMFADevice, output *iam.ListMFADeviceTagsOutput) Resource {
	describeCtx := GetDescribeContext(ctx)
	resource := Resource{
		Region: describeCtx.KaytuRegion,
		ARN:    *v.SerialNumber,
		Name:   *v.SerialNumber,
		Description: model.IAMVirtualMFADeviceDescription{
			VirtualMFADevice: v,
			Tags:             output.Tags,
		},
	}
	return resource
}
func GetIAMVirtualMFADevice(ctx context.Context, cfg aws.Config, fields map[string]string) ([]Resource, error) {
	client := iam.NewFromConfig(cfg)
	SerialNumber := fields["serialNumber"]
	output, err := client.ListVirtualMFADevices(ctx, &iam.ListVirtualMFADevicesInput{})
	if err != nil {
		if isErr(err, "ListVirtualMFADevicesNotFound") || isErr(err, "InvalidParameterValue") {
			return nil, nil
		}
		return nil, err
	}

	var values []Resource
	for _, v := range output.VirtualMFADevices {
		if *v.SerialNumber != SerialNumber {
			continue
		}

		output, err := client.ListMFADeviceTags(ctx, &iam.ListMFADeviceTagsInput{
			SerialNumber: &SerialNumber,
		})
		if err != nil {
			output = &iam.ListMFADeviceTagsOutput{}
		}

		resource := iAMVirtualMFADeviceHandle(ctx, v, output)
		values = append(values, resource)
	}
	return values, nil
}
