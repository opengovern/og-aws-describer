package describer

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ram"
	"github.com/aws/aws-sdk-go-v2/service/ram/types"
	"github.com/kaytu-io/kaytu-aws-describer/aws/model"
)

func RamPrincipalAssociation(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := ram.NewFromConfig(cfg)

	var values []Resource
	paginator := ram.NewGetResourceShareAssociationsPaginator(client, &ram.GetResourceShareAssociationsInput{AssociationType: types.ResourceShareAssociationTypePrincipal})
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, association := range page.ResourceShareAssociations {
			permissionPaginator := ram.NewListResourceSharePermissionsPaginator(client, &ram.ListResourceSharePermissionsInput{
				ResourceShareArn: association.ResourceShareArn,
			})
			var permissions []types.ResourceSharePermissionSummary
			for permissionPaginator.HasMorePages() {
				permissionPage, err := permissionPaginator.NextPage(ctx)
				if err != nil {
					return nil, err
				}
				permissions = append(permissions, permissionPage.Permissions...)
			}

			resource := Resource{
				Region: describeCtx.KaytuRegion,
				Name:   *association.ResourceShareName,
				ARN:    *association.ResourceShareArn,
				Description: model.RamPrincipalAssociationDescription{
					PrincipalAssociation:    association,
					ResourceSharePermission: permissions,
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

func RamResourceAssociation(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := ram.NewFromConfig(cfg)

	var values []Resource
	paginator := ram.NewGetResourceShareAssociationsPaginator(client, &ram.GetResourceShareAssociationsInput{AssociationType: types.ResourceShareAssociationTypeResource})
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, association := range page.ResourceShareAssociations {
			permissionPaginator := ram.NewListResourceSharePermissionsPaginator(client, &ram.ListResourceSharePermissionsInput{
				ResourceShareArn: association.ResourceShareArn,
			})
			var permissions []types.ResourceSharePermissionSummary
			for permissionPaginator.HasMorePages() {
				permissionPage, err := permissionPaginator.NextPage(ctx)
				if err != nil {
					return nil, err
				}
				permissions = append(permissions, permissionPage.Permissions...)
			}

			resource := Resource{
				Region: describeCtx.KaytuRegion,
				Name:   *association.ResourceShareName,
				ARN:    *association.ResourceShareArn,
				Description: model.RamResourceAssociationDescription{
					ResourceAssociation:     association,
					ResourceSharePermission: permissions,
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
