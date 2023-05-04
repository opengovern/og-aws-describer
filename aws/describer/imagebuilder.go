package describer

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/imagebuilder"
	"github.com/kaytu-io/kaytu-aws-describer/aws/model"
)

func ImageBuilderImage(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := imagebuilder.NewFromConfig(cfg)
	paginator := imagebuilder.NewListImagesPaginator(client, &imagebuilder.ListImagesInput{})

	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, v := range page.ImageVersionList {
			image, err := client.GetImage(ctx, &imagebuilder.GetImageInput{
				ImageBuildVersionArn: v.Arn,
			})
			if err != nil {
				return nil, err
			}

			resource := Resource{
				Region: describeCtx.KaytuRegion,
				ARN:    *image.Image.Arn,
				Name:   *image.Image.Name,
				Description: model.ImageBuilderImageDescription{
					Image: *image.Image,
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
