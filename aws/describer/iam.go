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
	describeCtx := GetDescribeContext(ctx)
	client := iam.NewFromConfig(cfg)
	output, err := client.GetAccountPasswordPolicy(ctx, &iam.GetAccountPasswordPolicyInput{})
	if err != nil {
		if !isErr(err, "NoSuchEntity") {
			return nil, err
		}

		output = &iam.GetAccountPasswordPolicyOutput{}
	}

	if output.PasswordPolicy == nil {
		return nil, nil
	}

	accountId, err := STSAccount(ctx, cfg)
	if err != nil {
		return nil, err
	}

	var values []Resource
	resource := Resource{
		Region: describeCtx.KaytuRegion,
		// No ID or ARN. Per Account Configuration
		Name: accountId + " IAM Password Policy",
		Description: model.IAMAccountPasswordPolicyDescription{
			PasswordPolicy: *output.PasswordPolicy,
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

func IAMAccessKey(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := iam.NewFromConfig(cfg)
	paginator := iam.NewListAccessKeysPaginator(client, &iam.ListAccessKeysInput{})

	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, v := range page.AccessKeyMetadata {
			arn := "arn:" + describeCtx.Partition + ":iam::" + describeCtx.AccountID + ":user/" + *v.UserName + "/accesskey/" + *v.AccessKeyId
			resource := Resource{
				Region: describeCtx.KaytuRegion,
				ARN:    arn,
				Name:   *v.UserName,
				Description: model.IAMAccessKeyDescription{
					AccessKey: v,
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

		for _, v := range page.Policies {
			version, err := client.GetPolicyVersion(ctx, &iam.GetPolicyVersionInput{
				PolicyArn: v.Arn,
				VersionId: v.DefaultVersionId,
			})
			if err != nil {
				return nil, err
			}

			resource := Resource{
				Region: describeCtx.KaytuRegion,
				ARN:    *v.Arn,
				Name:   *v.PolicyName,
				Description: model.IAMPolicyDescription{
					Policy:        v,
					PolicyVersion: *version.PolicyVersion,
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

func GetIAMPolicy(ctx context.Context, cfg aws.Config, arn string) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
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

	values = append(values, Resource{
		Region: describeCtx.KaytuRegion,
		ARN:    *v.Arn,
		Name:   *v.PolicyName,
		Description: model.IAMPolicyDescription{
			Policy:        *v,
			PolicyVersion: *version.PolicyVersion,
		},
	})

	return values, nil
}

func IAMGroup(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
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
	describeCtx := GetDescribeContext(ctx)
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

func GetIAMRole(ctx context.Context, cfg aws.Config, pathPrefix string) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
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

		values = append(values, Resource{
			Region: describeCtx.KaytuRegion,
			ARN:    *v.Arn,
			Name:   *v.RoleName,
			Description: model.IAMRoleDescription{
				Role:                v,
				InstanceProfileArns: profiles,
				InlinePolicies:      policies,
				AttachedPolicyArns:  aPolicies,
			},
		})
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

func IAMSAMLProvider(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := iam.NewFromConfig(cfg)
	output, err := client.ListSAMLProviders(ctx, &iam.ListSAMLProvidersInput{})
	if err != nil {
		return nil, err
	}

	var values []Resource
	for _, v := range output.SAMLProviderList {
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

func IAMServerCertificate(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
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

			resource := Resource{
				Region: describeCtx.KaytuRegion,
				ARN:    *v.Arn,
				Name:   *v.ServerCertificateName,
				Description: model.IAMServerCertificateDescription{
					ServerCertificate: *output.ServerCertificate,
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

func IAMUser(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
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

func GetIAMUser(ctx context.Context, cfg aws.Config, userName string) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
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

	values = append(values, Resource{
		Region: describeCtx.KaytuRegion,
		ARN:    *v.Arn,
		Name:   *v.UserName,
		Description: model.IAMUserDescription{
			User:               *v,
			Groups:             groups,
			InlinePolicies:     policies,
			AttachedPolicyArns: aPolicies,
			MFADevices:         devices,
		},
	})

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

func GetIAMPolicyAttachment(ctx context.Context, cfg aws.Config, policyARN string) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := iam.NewFromConfig(cfg)
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
	values = append(values, Resource{
		Region: describeCtx.KaytuRegion,
		Name:   fmt.Sprintf("%s - Attachments", policyARN),
		Description: model.IAMPolicyAttachmentDescription{
			PolicyArn:             policyARN,
			PolicyAttachmentCount: *policy.Policy.AttachmentCount,
			IsAttached:            *policy.Policy.AttachmentCount > 0,
			PolicyGroups:          policyGroups,
			PolicyRoles:           policyRoles,
			PolicyUsers:           policyUsers,
		},
	})

	return values, nil
}

func IAMSamlProvider(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
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
		resource := Resource{
			Region: describeCtx.KaytuRegion,
			ARN:    *v.Arn,
			Description: model.IAMSamlProviderDescription{
				SamlProvider: *samlProvider,
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

func IAMServiceSpecificCredential(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
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
				resource := Resource{
					Region: describeCtx.KaytuRegion,
					ID:     *credential.ServiceSpecificCredentialId,
					Description: model.IAMServiceSpecificCredentialDescription{
						ServiceSpecificCredential: credential,
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
	}

	return values, nil
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
	describeCtx := GetDescribeContext(ctx)
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
			return nil, err
		}

		resource := Resource{
			Region: describeCtx.KaytuRegion,
			ARN:    *v.SerialNumber,
			Name:   *v.SerialNumber,
			Description: model.IAMVirtualMFADeviceDescription{
				VirtualMFADevice: v,
				Tags:             output.Tags,
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
