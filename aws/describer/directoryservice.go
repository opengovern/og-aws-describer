package describer

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/directoryservice"
	"github.com/kaytu-io/kaytu-aws-describer/aws/model"
)

func DirectoryServiceDirectory(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)

	client := directoryservice.NewFromConfig(cfg)
	paginator := directoryservice.NewDescribeDirectoriesPaginator(client, &directoryservice.DescribeDirectoriesInput{})

	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			if !isErr(err, "InvalidParameterValueException") && !isErr(err, "ResourceNotFoundFault") && !isErr(err, "EntityDoesNotExistException") {
				return nil, err
			}
			continue
		}

		for _, v := range page.DirectoryDescriptions {
			arn := fmt.Sprintf("arn:%s:ds:%s:%s:directory/%s", describeCtx.Partition, describeCtx.Region, describeCtx.AccountID, *v.DirectoryId)

			tags, err := client.ListTagsForResource(ctx, &directoryservice.ListTagsForResourceInput{
				ResourceId: v.DirectoryId,
			})
			if err != nil {
				if !isErr(err, "InvalidParameterValueException") && !isErr(err, "ResourceNotFoundFault") && !isErr(err, "EntityDoesNotExistException") {
					return nil, err
				}
				tags = &directoryservice.ListTagsForResourceOutput{}
			}
			sharedDirectory, err := client.DescribeSharedDirectories(ctx, &directoryservice.DescribeSharedDirectoriesInput{
				OwnerDirectoryId: v.DirectoryId,
			})
			if err != nil {
				if isErr(err, "DescribeSharedDirectoriesNotFound") || isErr(err, "InvalidParameterValue") {
					return nil, nil
				}
				return nil, err
			}

			snapshot, err := client.GetSnapshotLimits(ctx, &directoryservice.GetSnapshotLimitsInput{
				DirectoryId: v.DirectoryId,
			})
			if err != nil {
				if isErr(err, "GetSnapshotLimitsNotFound") || isErr(err, "InvalidParameterValue") {
					return nil, nil
				}
				return nil, err
			}

			eventTopic, err := client.DescribeEventTopics(ctx, &directoryservice.DescribeEventTopicsInput{
				DirectoryId: v.DirectoryId,
			})
			if err != nil {
				if isErr(err, "DescribeEventTopicsNotFound") || isErr(err, "InvalidParameterValue") {
					return nil, nil
				}
				return nil, err
			}

			resource := Resource{
				Region: describeCtx.KaytuRegion,
				ARN:    arn,
				Name:   *v.Name,
				Description: model.DirectoryServiceDirectoryDescription{
					Directory:       v,
					Snapshot:        *snapshot.SnapshotLimits,
					EventTopic:      *eventTopic.EventTopics,
					SharedDirectory: sharedDirectory.SharedDirectories,
					Tags:            tags.Tags,
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
