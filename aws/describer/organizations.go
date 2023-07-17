package describer

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/organizations"
	"github.com/aws/aws-sdk-go-v2/service/organizations/types"
	"github.com/kaytu-io/kaytu-aws-describer/aws/model"
)

// OrganizationOrganization Retrieves information about the organization that the
// user's account belongs to.
func OrganizationOrganization(ctx context.Context, cfg aws.Config) (*types.Organization, error) {
	client := organizations.NewFromConfig(cfg)

	req, err := client.DescribeOrganization(ctx, &organizations.DescribeOrganizationInput{})
	if err != nil {
		return nil, err
	}

	return req.Organization, nil
}

func OrganizationsOrganization(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	client := organizations.NewFromConfig(cfg)

	req, err := client.DescribeOrganization(ctx, &organizations.DescribeOrganizationInput{})
	if err != nil {
		return nil, err
	}

	var values []Resource
	resource := organizationsOrganizationHandle(ctx, req)
	if stream != nil {
		if err := (*stream)(resource); err != nil {
			return nil, err
		}
	} else {
		values = append(values, resource)
	}

	return values, nil
}
func organizationsOrganizationHandle(ctx context.Context, req *organizations.DescribeOrganizationOutput) Resource {
	describeCtx := GetDescribeContext(ctx)
	resource := Resource{
		Region: describeCtx.KaytuRegion,
		ARN:    *req.Organization.Arn,
		Name:   *req.Organization.Id,
		Description: model.OrganizationsOrganizationDescription{
			Organization: *req.Organization,
		},
	}
	return resource
}
func GetOrganizationsOrganization(ctx context.Context, cfg aws.Config, fields map[string]string) ([]Resource, error) {
	client := organizations.NewFromConfig(cfg)
	var values []Resource
	describes, err := client.DescribeOrganization(ctx, &organizations.DescribeOrganizationInput{})
	if err != nil {
		return nil, err
	}
	resource := organizationsOrganizationHandle(ctx, describes)
	values = append(values, resource)
	return values, nil
}

// OrganizationAccount Retrieves AWS Organizations-related information about
// the specified (ID) account .
func OrganizationAccount(ctx context.Context, cfg aws.Config, id string) (*types.Account, error) {
	svc := organizations.NewFromConfig(cfg)

	req, err := svc.DescribeAccount(ctx, &organizations.DescribeAccountInput{AccountId: aws.String(id)})
	if err != nil {
		return nil, err
	}

	return req.Account, nil
}

// DescribeOrganization Retrieves information about the organization that the
// user's account belongs to.
func OrganizationAccounts(ctx context.Context, cfg aws.Config) ([]types.Account, error) {
	client := organizations.NewFromConfig(cfg)

	paginator := organizations.NewListAccountsPaginator(client, &organizations.ListAccountsInput{})

	var values []types.Account
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		values = append(values, page.Accounts...)
	}

	return values, nil
}

func OrganizationsAccount(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := organizations.NewFromConfig(cfg)

	paginator := organizations.NewListAccountsPaginator(client, &organizations.ListAccountsInput{})

	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, acc := range page.Accounts {
			tags, err := client.ListTagsForResource(ctx, &organizations.ListTagsForResourceInput{
				ResourceId: acc.Id,
			})
			if err != nil {
				return nil, err
			}

			resource := Resource{
				Region: describeCtx.KaytuRegion,
				ARN:    *acc.Arn,
				Name:   *acc.Name,
				Description: model.OrganizationsAccountDescription{
					Account: acc,
					Tags:    tags.Tags,
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
