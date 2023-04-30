package describer

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/identitystore"
	"github.com/kaytu-io/kaytu-aws-describer/aws/model"
)

func IdentityStoreGroup(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := identitystore.NewFromConfig(cfg)
	paginator := identitystore.NewListGroupsPaginator(client, &identitystore.ListGroupsInput{})

	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, group := range page.Groups {
			resource := Resource{
				Region: describeCtx.Region,
				ID:     *group.GroupId,
				Name:   *group.DisplayName,
				Description: model.IdentityStoreGroupDescription{
					Group: group,
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

func IdentityStoreUser(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := identitystore.NewFromConfig(cfg)
	paginator := identitystore.NewListUsersPaginator(client, &identitystore.ListUsersInput{})

	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, user := range page.Users {
			resource := Resource{
				Region: describeCtx.Region,
				ID:     *user.UserId,
				Name:   *user.UserName,
				Description: model.IdentityStoreUserDescription{
					User: user,
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
