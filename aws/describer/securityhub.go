package describer

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/securityhub"
	"github.com/kaytu-io/kaytu-aws-describer/aws/model"
)

func SecurityHubHub(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := securityhub.NewFromConfig(cfg)
	out, err := client.DescribeHub(ctx, &securityhub.DescribeHubInput{})
	if err != nil {
		if isErr(err, "InvalidAccessException") {
			return nil, nil
		}
		return nil, err
	}

	var values []Resource

	tags, err := client.ListTagsForResource(ctx, &securityhub.ListTagsForResourceInput{ResourceArn: out.HubArn})
	if err != nil {
		return nil, err
	}

	resource := Resource{
		Region: describeCtx.KaytuRegion,
		ARN:    *out.HubArn,
		Name:   nameFromArn(*out.HubArn),
		Description: model.SecurityHubHubDescription{
			Hub:  out,
			Tags: tags.Tags,
		},
	}
	if stream != nil {
		if err := (*stream)(resource); err != nil {
			return nil, err
		}
	} else {
		values = append(values, resource)
	}

	return values, nil
}
