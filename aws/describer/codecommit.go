package describer

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/codecommit"
	"github.com/kaytu-io/kaytu-aws-describer/aws/model"
)

func CodeCommitRepository(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := codecommit.NewFromConfig(cfg)
	paginator := codecommit.NewListRepositoriesPaginator(client, &codecommit.ListRepositoriesInput{})

	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			if !isErr(err, "InvalidParameter") {
				return nil, err
			}
			continue
		}
		var repositoryNames []string
		for _, v := range page.Repositories {
			repositoryNames = append(repositoryNames, *v.RepositoryName)
		}
		if len(repositoryNames) == 0 {
			continue
		}
		repos, err := client.BatchGetRepositories(ctx, &codecommit.BatchGetRepositoriesInput{
			RepositoryNames: repositoryNames,
		})
		if err != nil {
			if !isErr(err, "InvalidParameter") {
				return nil, err
			}
			continue
		}
		for _, v := range repos.Repositories {
			tags, err := client.ListTagsForResource(ctx, &codecommit.ListTagsForResourceInput{
				ResourceArn: v.Arn,
			})
			if err != nil {
				if !isErr(err, "InvalidParameter") {
					return nil, err
				}
				tags = &codecommit.ListTagsForResourceOutput{}
			}

			resource := Resource{
				Region: describeCtx.Region,
				ARN:    *v.Arn,
				Name:   *v.RepositoryName,
				Description: model.CodeCommitRepositoryDescription{
					Repository: v,
					Tags:       tags.Tags,
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
