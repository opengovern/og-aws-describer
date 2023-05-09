package describer

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/kafka/types"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/kafka"
	"github.com/kaytu-io/kaytu-aws-describer/aws/model"
)

func KafkaCluster(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := kafka.NewFromConfig(cfg)
	paginator := kafka.NewListClustersV2Paginator(client, &kafka.ListClustersV2Input{})

	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, cluster := range page.ClusterInfoList {
			var configArn string
			var operationInfo *types.ClusterOperationInfo
			var configOut *kafka.DescribeConfigurationOutput
			if cluster.ClusterType == types.ClusterTypeProvisioned {
				if cluster.Provisioned.CurrentBrokerSoftwareInfo.ConfigurationArn != nil {
					configArn = *cluster.Provisioned.CurrentBrokerSoftwareInfo.ConfigurationArn
				}
			}

			if len(configArn) >= 1 {
				configOut, err = client.DescribeConfiguration(ctx, &kafka.DescribeConfigurationInput{Arn: &configArn})
				if err != nil {
					return nil, err
				}
			}

			if cluster.ActiveOperationArn != nil {
				op, err := client.DescribeClusterOperation(ctx, &kafka.DescribeClusterOperationInput{
					ClusterOperationArn: cluster.ActiveOperationArn,
				})
				if err != nil {
					return nil, err
				}
				operationInfo = op.ClusterOperationInfo
			}

			resource := Resource{
				Region: describeCtx.KaytuRegion,
				ARN:    *cluster.ClusterArn,
				Name:   *cluster.ClusterName,
				Description: model.KafkaClusterDescription{
					Cluster:              cluster,
					ClusterOperationInfo: operationInfo,
					Configuration:        configOut,
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
