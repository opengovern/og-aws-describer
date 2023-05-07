package describer

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/kaytu-io/kaytu-aws-describer/aws/model"
)

func SSMManagedInstance(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := ssm.NewFromConfig(cfg)
	paginator := ssm.NewDescribeInstanceInformationPaginator(client, &ssm.DescribeInstanceInformationInput{})

	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, item := range page.InstanceInformationList {
			arn := "arn:" + describeCtx.Partition + ":ssm:" + describeCtx.Region + ":" + describeCtx.AccountID + ":managed-instance/" + *item.InstanceId
			name := ""
			if item.Name != nil {
				name = *item.Name
			} else {
				name = *item.InstanceId
			}
			resource := Resource{
				Region: describeCtx.KaytuRegion,
				ARN:    arn,
				Name:   name,
				Description: model.SSMManagedInstanceDescription{
					InstanceInformation: item,
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

func SSMManagedInstanceCompliance(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := ssm.NewFromConfig(cfg)
	paginator := ssm.NewDescribeInstanceInformationPaginator(client, &ssm.DescribeInstanceInformationInput{})

	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, item := range page.InstanceInformationList {
			cpaginator := ssm.NewListComplianceItemsPaginator(client, &ssm.ListComplianceItemsInput{
				ResourceIds: []string{*item.InstanceId},
			})

			for cpaginator.HasMorePages() {
				cpage, err := cpaginator.NextPage(ctx)
				if err != nil {
					return nil, err
				}

				for _, item := range cpage.ComplianceItems {
					arn := "arn:" + describeCtx.Partition + ":ssm:" + describeCtx.Region + ":" + describeCtx.AccountID + ":managed-instance/" + *item.ResourceId + "/compliance-item/" + *item.Id + ":" + *item.ComplianceType
					resource := Resource{
						Region: describeCtx.KaytuRegion,
						ARN:    arn,
						Name:   *item.Title,
						Description: model.SSMManagedInstanceComplianceDescription{
							ComplianceItem: item,
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

	return values, nil
}

func SSMAssociation(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := ssm.NewFromConfig(cfg)
	paginator := ssm.NewListAssociationsPaginator(client, &ssm.ListAssociationsInput{})

	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, v := range page.Associations {
			out, err := client.DescribeAssociation(ctx, &ssm.DescribeAssociationInput{
				AssociationId: v.AssociationId,
			})
			if err != nil {
				return nil, err
			}

			resource := Resource{
				Region: describeCtx.Region,
				ID:     *v.AssociationId,
				Name:   *v.Name,
				Description: model.SSMAssociationDescription{
					AssociationItem: v,
					Association:     out,
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

func SSMDocument(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := ssm.NewFromConfig(cfg)
	paginator := ssm.NewListDocumentsPaginator(client, &ssm.ListDocumentsInput{})

	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, v := range page.DocumentIdentifiers {
			permissions, err := client.DescribeDocumentPermission(ctx, &ssm.DescribeDocumentPermissionInput{
				Name:           v.Name,
				PermissionType: "Share",
			})
			if err != nil {
				return nil, err
			}

			data, err := client.DescribeDocument(ctx, &ssm.DescribeDocumentInput{
				Name: v.Name,
			})
			if err != nil {
				return nil, err
			}

			resource := Resource{
				Region: describeCtx.Region,
				ID:     *v.Name,
				Name:   *v.Name,
				Description: model.SSMDocumentDescription{
					DocumentIdentifier: v,
					Document:           data,
					Permissions:        permissions,
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

func SSMMaintenanceWindow(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := ssm.NewFromConfig(cfg)
	paginator := ssm.NewDescribeMaintenanceWindowsPaginator(client, &ssm.DescribeMaintenanceWindowsInput{})

	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, v := range page.WindowIdentities {
			resource := Resource{
				Region:      describeCtx.Region,
				ID:          *v.WindowId,
				Name:        *v.Name,
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

func SSMMaintenanceWindowTarget(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	windows, err := SSMMaintenanceWindow(ctx, cfg, nil)
	if err != nil {
		return nil, err
	}

	client := ssm.NewFromConfig(cfg)

	var values []Resource
	for _, w := range windows {
		window := w.Description.(types.MaintenanceWindowIdentity)
		paginator := ssm.NewDescribeMaintenanceWindowTargetsPaginator(client, &ssm.DescribeMaintenanceWindowTargetsInput{
			WindowId: window.WindowId,
		})

		for paginator.HasMorePages() {
			page, err := paginator.NextPage(ctx)
			if err != nil {
				return nil, err
			}

			for _, v := range page.Targets {
				resource := Resource{
					Region:      describeCtx.Region,
					ID:          *v.WindowTargetId,
					Name:        *v.Name,
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
	}

	return values, nil
}

func SSMMaintenanceWindowTask(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	windows, err := SSMMaintenanceWindow(ctx, cfg, nil)
	if err != nil {
		return nil, err
	}

	client := ssm.NewFromConfig(cfg)

	var values []Resource
	for _, w := range windows {
		window := w.Description.(types.MaintenanceWindowIdentity)
		paginator := ssm.NewDescribeMaintenanceWindowTasksPaginator(client, &ssm.DescribeMaintenanceWindowTasksInput{
			WindowId: window.WindowId,
		})

		for paginator.HasMorePages() {
			page, err := paginator.NextPage(ctx)
			if err != nil {
				return nil, err
			}

			for _, v := range page.Tasks {
				resource := Resource{
					Region:      describeCtx.Region,
					ARN:         *v.TaskArn,
					Name:        *v.Name,
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
	}

	return values, nil
}

func SSMParameter(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := ssm.NewFromConfig(cfg)
	paginator := ssm.NewDescribeParametersPaginator(client, &ssm.DescribeParametersInput{})

	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, v := range page.Parameters {
			resource := Resource{
				Region:      describeCtx.Region,
				ID:          *v.Name,
				Name:        *v.Name,
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

func SSMPatchBaseline(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := ssm.NewFromConfig(cfg)
	paginator := ssm.NewDescribePatchBaselinesPaginator(client, &ssm.DescribePatchBaselinesInput{})

	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, v := range page.BaselineIdentities {
			resource := Resource{
				Region:      describeCtx.Region,
				ID:          *v.BaselineId,
				Name:        *v.BaselineName,
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

func SSMResourceDataSync(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := ssm.NewFromConfig(cfg)
	paginator := ssm.NewListResourceDataSyncPaginator(client, &ssm.ListResourceDataSyncInput{})

	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, v := range page.ResourceDataSyncItems {
			resource := Resource{
				Region:      describeCtx.Region,
				ID:          *v.SyncName,
				Name:        *v.SyncName,
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
