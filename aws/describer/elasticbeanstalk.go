package describer

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk"
	"github.com/kaytu-io/kaytu-aws-describer/aws/model"
)

func ElasticBeanstalkEnvironment(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := elasticbeanstalk.NewFromConfig(cfg)
	out, err := client.DescribeEnvironments(ctx, &elasticbeanstalk.DescribeEnvironmentsInput{})
	if err != nil {
		return nil, err
	}

	var values []Resource
	for _, item := range out.Environments {
		tags, err := client.ListTagsForResource(ctx, &elasticbeanstalk.ListTagsForResourceInput{
			ResourceArn: item.EnvironmentArn,
		})
		if err != nil {
			return nil, err
		}

		resource := Resource{
			Region: describeCtx.KaytuRegion,
			ARN:    *item.EnvironmentArn,
			Name:   *item.EnvironmentName,
			Description: model.ElasticBeanstalkEnvironmentDescription{
				EnvironmentDescription: item,
				Tags:                   tags.ResourceTags,
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

func ElasticBeanstalkApplication(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := elasticbeanstalk.NewFromConfig(cfg)
	out, err := client.DescribeApplications(ctx, &elasticbeanstalk.DescribeApplicationsInput{})
	if err != nil {
		if !isErr(err, "ResourceNotFoundException") {
			return nil, err
		}
		return nil, nil
	}

	var values []Resource
	for _, item := range out.Applications {
		tags, err := client.ListTagsForResource(ctx, &elasticbeanstalk.ListTagsForResourceInput{
			ResourceArn: item.ApplicationArn,
		})
		if err != nil {
			if !isErr(err, "ResourceNotFoundException") {
				return nil, err
			}
			tags = &elasticbeanstalk.ListTagsForResourceOutput{}
		}

		resource := Resource{
			Region: describeCtx.KaytuRegion,
			ARN:    *item.ApplicationArn,
			Name:   *item.ApplicationName,
			Description: model.ElasticBeanstalkApplicationDescription{
				Application: item,
				Tags:        tags.ResourceTags,
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

func ElasticBeanstalkPlatform(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := elasticbeanstalk.NewFromConfig(cfg)
	paginator := elasticbeanstalk.NewListPlatformVersionsPaginator(client, &elasticbeanstalk.ListPlatformVersionsInput{})

	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, item := range page.PlatformSummaryList {
			platform, err := client.DescribePlatformVersion(ctx, &elasticbeanstalk.DescribePlatformVersionInput{
				PlatformArn: item.PlatformArn,
			})
			if err != nil {
				return nil, err
			}

			resource := Resource{
				Region: describeCtx.KaytuRegion,
				ARN:    *platform.PlatformDescription.PlatformArn,
				Name:   *platform.PlatformDescription.PlatformName,
				Description: model.ElasticBeanstalkPlatformDescription{
					Platform: *platform.PlatformDescription,
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
