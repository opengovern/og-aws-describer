package describer

import (
	"context"
	"fmt"

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
			permissionSetPaginator := ssoadmin.NewListPermissionSetsPaginator(client, &ssoadmin.ListPermissionSetsInput{
				InstanceArn: v.InstanceArn,
			})
			for permissionSetPaginator.HasMorePages() {
				psPage, err := permissionSetPaginator.NextPage(ctx)
				if err != nil {
					return nil, err
				}
				for _, ps := range psPage.PermissionSets {
					aaPaginator := ssoadmin.NewListAccountAssignmentsPaginator(client, &ssoadmin.ListAccountAssignmentsInput{
						InstanceArn:      v.InstanceArn,
						PermissionSetArn: &ps,
					})
					for aaPaginator.HasMorePages() {
						aaPage, err := aaPaginator.NextPage(ctx)
						if err != nil {
							return nil, err
						}
						for _, aa := range aaPage.AccountAssignments {
							id := fmt.Sprintf("%s-%s-%s", *v.InstanceArn, ps, *aa.AccountId)
							resource := Resource{
								Region: describeCtx.Region,
								ID:     id,
								Description: model.SSOAdminAccountAssignmentDescription{
									Instance:          v,
									AccountAssignment: aa,
									PermissionSetArn:  ps,
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
			}
		}
	}
	return values, nil
}
