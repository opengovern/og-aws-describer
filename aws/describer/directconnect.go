package describer

import (
	"context"
	"fmt"
	"math"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/directconnect"
	"github.com/aws/aws-sdk-go-v2/service/directconnect/types"
	"github.com/kaytu-io/kaytu-aws-describer/aws/model"
)

func DirectConnectConnection(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)

	client := directconnect.NewFromConfig(cfg)

	connections, err := client.DescribeConnections(ctx, &directconnect.DescribeConnectionsInput{})
	if err != nil {
		return nil, err
	}

	var values []Resource
	for _, v := range connections.Connections {
		arn := fmt.Sprintf("arn:%s:directconnect:%s:%s:dxcon/%s", describeCtx.Partition, describeCtx.Region, describeCtx.AccountID, *v.ConnectionId)

		resource := Resource{
			Region: describeCtx.KaytuRegion,
			ARN:    arn,
			Name:   *v.ConnectionId,
			Description: model.DirectConnectConnectionDescription{
				Connection: v,
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

func getDirectConnectGatewayArn(describeCtx DescribeContext, directConnectGatewayId string) string {
	return fmt.Sprintf("arn:%s:directconnect::%s:dx-gateway/%s", describeCtx.Partition, describeCtx.AccountID, directConnectGatewayId)
}

func DirectConnectGateway(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)

	client := directconnect.NewFromConfig(cfg)

	var values []Resource
	err := PaginateRetrieveAll(func(prevToken *string) (nextToken *string, err error) {
		connections, err := client.DescribeDirectConnectGateways(ctx, &directconnect.DescribeDirectConnectGatewaysInput{
			MaxResults: aws.Int32(100),
			NextToken:  prevToken,
		})
		if err != nil {
			return nil, err
		}
		if len(connections.DirectConnectGateways) == 0 {
			return nil, nil
		}
		arns := make([]string, 0, len(connections.DirectConnectGateways))
		for _, v := range connections.DirectConnectGateways {
			arns = append(arns, getDirectConnectGatewayArn(describeCtx, *v.DirectConnectGatewayId))
		}
		// DescribeTags can only handle 20 ARNs at a time
		arnToTagMap := make(map[string][]types.Tag)
		for i := 0; i < len(arns); i += 20 {
			tags, err := client.DescribeTags(ctx, &directconnect.DescribeTagsInput{
				ResourceArns: arns[i:int(math.Min(float64(i+20), float64(len(arns))))],
			})
			if err != nil {
				return nil, err
			}

			for _, tag := range tags.ResourceTags {
				arnToTagMap[*tag.ResourceArn] = tag.Tags
			}
		}

		for _, v := range connections.DirectConnectGateways {
			arn := getDirectConnectGatewayArn(describeCtx, *v.DirectConnectGatewayId)

			tagsList, ok := arnToTagMap[arn]
			if !ok {
				tagsList = []types.Tag{}
			}

			resource := Resource{
				Region: describeCtx.KaytuRegion,
				ARN:    arn,
				Name:   *v.DirectConnectGatewayName,
				Description: model.DirectConnectGatewayDescription{
					Gateway: v,
					Tags:    tagsList,
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

		return connections.NextToken, nil
	})
	if err != nil {
		return nil, err
	}

	return values, nil
}
