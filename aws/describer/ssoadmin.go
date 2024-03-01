package describer

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ssoadmin"
	"github.com/aws/aws-sdk-go-v2/service/ssoadmin/types"
	"github.com/kaytu-io/kaytu-aws-describer/aws/model"
)

func SSOAdminInstance(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := ssoadmin.NewFromConfig(cfg)
	paginator := ssoadmin.NewListInstancesPaginator(client, &ssoadmin.ListInstancesInput{})

	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, v := range page.Instances {
			resource := Resource{
				Region: describeCtx.KaytuRegion,
				ARN:    *v.InstanceArn,
				Name:   *v.InstanceArn,
				Description: model.SSOAdminInstanceDescription{
					Instance: v,
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

func SSOAdminAccountAssignment(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	client := ssoadmin.NewFromConfig(cfg)
	paginator := ssoadmin.NewListInstancesPaginator(client, &ssoadmin.ListInstancesInput{})

	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, v := range page.Instances {
			resources, err := ListSSOAdminInstanceAccountAssignments(ctx, client, v)
			if err != nil {
				return nil, err
			}
			for _, resource := range resources {
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

func ListSSOAdminInstanceAccountAssignments(ctx context.Context, client *ssoadmin.Client, instance types.InstanceMetadata) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	paginator := ssoadmin.NewListPermissionSetsPaginator(client, &ssoadmin.ListPermissionSetsInput{
		InstanceArn: instance.InstanceArn,
	})
	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, v := range page.PermissionSets {
			accountAssignment, err := client.ListAccountAssignments(ctx, &ssoadmin.ListAccountAssignmentsInput{
				InstanceArn:      instance.InstanceArn,
				AccountId:        aws.String(describeCtx.AccountID),
				PermissionSetArn: &v,
			})
			if err != nil {
				return nil, err
			}

			for _, accountA := range accountAssignment.AccountAssignments {
				resource := Resource{
					Region: describeCtx.Region,
					ID:     fmt.Sprintf("%s|%s|%s", *accountA.AccountId, *accountA.PermissionSetArn, *accountA.PrincipalId),
					Description: model.SSOAdminAccountAssignmentDescription{
						Instance:          instance,
						AccountAssignment: accountA,
					},
				}
				values = append(values, resource)
			}
		}
	}
	return values, nil
}

func SSOAdminPermissionSet(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	client := ssoadmin.NewFromConfig(cfg)
	paginator := ssoadmin.NewListInstancesPaginator(client, &ssoadmin.ListInstancesInput{})

	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, v := range page.Instances {
			resources, err := ListSSOAdminInstancePermissionSets(ctx, client, *v.InstanceArn)
			if err != nil {
				return nil, err
			}
			for _, resource := range resources {
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

func ListSSOAdminInstancePermissionSets(ctx context.Context, client *ssoadmin.Client, instanceArn string) ([]Resource, error) {
	paginator := ssoadmin.NewListPermissionSetsPaginator(client, &ssoadmin.ListPermissionSetsInput{
		InstanceArn: &instanceArn,
	})
	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, v := range page.PermissionSets {
			resource, err := GetSSOAdminPermissionSet(ctx, client, instanceArn, v)
			if err != nil {
				return nil, err
			}
			values = append(values, *resource)
		}
	}
	return values, nil
}

func GetSSOAdminPermissionSet(ctx context.Context, client *ssoadmin.Client, instanceArn, permissionSetArn string) (*Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	detail, err := client.DescribePermissionSet(ctx, &ssoadmin.DescribePermissionSetInput{
		InstanceArn:      aws.String(instanceArn),
		PermissionSetArn: aws.String(permissionSetArn),
	})
	if err != nil {
		return nil, err
	}

	tags := []types.Tag{}

	paginator := ssoadmin.NewListTagsForResourcePaginator(client, &ssoadmin.ListTagsForResourceInput{
		InstanceArn: aws.String(instanceArn),
		ResourceArn: aws.String(permissionSetArn),
	})

	for paginator.HasMorePages() {
		output, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		tags = append(tags, output.Tags...)
	}
	tagsMap := map[string]string{}

	for _, tag := range tags {
		tagsMap[*tag.Key] = *tag.Value
	}
	resource := Resource{
		Region: describeCtx.Region,
		ID:     *detail.PermissionSet.PermissionSetArn,
		ARN:    *detail.PermissionSet.PermissionSetArn,
		Description: model.SSOAdminPermissionSetDescription{
			InstanceArn:   instanceArn,
			PermissionSet: *detail.PermissionSet,
			Tags:          tagsMap,
		},
	}
	return &resource, nil
}

func SSOAdminManagedPolicyAttachment(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	client := ssoadmin.NewFromConfig(cfg)
	paginator := ssoadmin.NewListInstancesPaginator(client, &ssoadmin.ListInstancesInput{})

	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, v := range page.Instances {
			resources, err := ListSSOAdminInstanceManagedPolicyAttachments(ctx, client, *v.InstanceArn)
			if err != nil {
				return nil, err
			}
			for _, resource := range resources {
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

func ListSSOAdminInstanceManagedPolicyAttachments(ctx context.Context, client *ssoadmin.Client, instanceArn string) ([]Resource, error) {
	paginator := ssoadmin.NewListPermissionSetsPaginator(client, &ssoadmin.ListPermissionSetsInput{
		InstanceArn: &instanceArn,
	})
	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, v := range page.PermissionSets {
			resources, err := ListSSOAdminPermissionSetPolicyAttachments(ctx, client, instanceArn, v)
			if err != nil {
				return nil, err
			}
			values = append(values, resources...)
		}
	}
	return values, nil
}

func ListSSOAdminPermissionSetPolicyAttachments(ctx context.Context, client *ssoadmin.Client, instanceArn, permissionSetArn string) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	paginator := ssoadmin.NewListManagedPoliciesInPermissionSetPaginator(client, &ssoadmin.ListManagedPoliciesInPermissionSetInput{
		InstanceArn:      aws.String(instanceArn),
		PermissionSetArn: aws.String(permissionSetArn),
	})

	var values []Resource
	for paginator.HasMorePages() {
		output, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, item := range output.AttachedManagedPolicies {
			resource := Resource{
				Region: describeCtx.Region,
				ID:     *item.Arn,
				ARN:    *item.Arn,
				Description: model.SSOAdminPolicyAttachmentDescription{
					InstanceArn:           instanceArn,
					PermissionSetArn:      permissionSetArn,
					AttachedManagedPolicy: item,
				},
			}
			values = append(values, resource)
		}
	}
	return values, nil
}
