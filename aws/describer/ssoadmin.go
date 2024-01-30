package describer

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ssoadmin"
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
			permissionSet, err := client.ProvisionPermissionSet(ctx, &ssoadmin.ProvisionPermissionSetInput{
				InstanceArn: v.InstanceArn,
			})
			if err != nil {
				return nil, err
			}

			accountAssignment, err := client.ListAccountAssignments(ctx, &ssoadmin.ListAccountAssignmentsInput{
				InstanceArn: v.InstanceArn,
			})
			if err != nil {
				return nil, err
			}

			for _, accountA := range accountAssignment.AccountAssignments {
				resource := Resource{
					Region: describeCtx.Region,
					ID:     *v.IdentityStoreId,
					ARN:    *v.InstanceArn,
					Description: model.SSOAdminAccountAssignmentDescription{
						Instance:               v,
						AccountAssignment:      accountA,
						PermissionSetProvision: *permissionSet.PermissionSetProvisioningStatus,
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
