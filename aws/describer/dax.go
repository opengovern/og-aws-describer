package describer

import (
	"context"
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dax"
	"github.com/kaytu-io/kaytu-aws-describer/aws/model"
)

func DAXCluster(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := dax.NewFromConfig(cfg)
	out, err := client.DescribeClusters(ctx, &dax.DescribeClustersInput{})
	if err != nil {
		if strings.Contains(err.Error(), "InvalidParameterValueException") || strings.Contains(err.Error(), "no such host") {
			return nil, nil
		}
		return nil, err
	}

	var values []Resource
	for _, cluster := range out.Clusters {
		tags, err := client.ListTags(ctx, &dax.ListTagsInput{
			ResourceName: cluster.ClusterArn,
		})
		if err != nil {
			if strings.Contains(err.Error(), "ClusterNotFoundFault") {
				tags = nil
			} else {
				return nil, err
			}
		}

		resource := Resource{
			Region: describeCtx.KaytuRegion,
			ARN:    *cluster.ClusterArn,
			Name:   *cluster.ClusterName,
			Description: model.DAXClusterDescription{
				Cluster: cluster,
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

	return values, nil
}

func DAXParameterGroup(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	//
	client := dax.NewFromConfig(cfg)

	var values []Resource
	err := PaginateRetrieveAll(func(prevToken *string) (nextToken *string, err error) {
		parameterGroups, err := client.DescribeParameterGroups(ctx, &dax.DescribeParameterGroupsInput{
			MaxResults: aws.Int32(100),
			NextToken:  prevToken,
		})
		if err != nil {
			return nil, err
		}

		for _, parameterGroup := range parameterGroups.ParameterGroups {
			resource := Resource{
				Region: describeCtx.KaytuRegion,
				Name:   *parameterGroup.ParameterGroupName,
				Description: model.DAXParameterGroupDescription{
					ParameterGroup: parameterGroup,
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

		return parameterGroups.NextToken, nil
	})
	if err != nil {
		return nil, err
	}

	return values, nil
}

func DAXParameter(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	//
	client := dax.NewFromConfig(cfg)

	var values []Resource
	err := PaginateRetrieveAll(func(prevToken *string) (nextToken *string, err error) {
		parameterGroups, err := client.DescribeParameterGroups(ctx, &dax.DescribeParameterGroupsInput{
			MaxResults: aws.Int32(100),
			NextToken:  prevToken,
		})
		if err != nil {
			return nil, err
		}

		for _, parameterGroup := range parameterGroups.ParameterGroups {
			err := PaginateRetrieveAll(func(prevToken *string) (nextToken *string, err error) {
				parameters, err := client.DescribeParameters(ctx, &dax.DescribeParametersInput{
					ParameterGroupName: parameterGroup.ParameterGroupName,
					MaxResults:         aws.Int32(100),
					NextToken:          prevToken,
				})
				if err != nil {
					return nil, err
				}

				for _, parameter := range parameters.Parameters {
					resource := Resource{
						Region: describeCtx.KaytuRegion,
						Name:   *parameter.ParameterName,
						Description: model.DAXParameterDescription{
							Parameter:          parameter,
							ParameterGroupName: *parameterGroup.ParameterGroupName,
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

				return parameters.NextToken, nil
			})
			if err != nil {
				return nil, err
			}
		}

		return parameterGroups.NextToken, nil
	})
	if err != nil {
		return nil, err
	}

	return values, nil
}

func DAXSubnetGroup(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)

	client := dax.NewFromConfig(cfg)

	var values []Resource
	err := PaginateRetrieveAll(func(prevToken *string) (nextToken *string, err error) {
		subnetGroups, err := client.DescribeSubnetGroups(ctx, &dax.DescribeSubnetGroupsInput{
			MaxResults: aws.Int32(100),
			NextToken:  prevToken,
		})
		if err != nil {
			return nil, err
		}

		for _, subnetGroup := range subnetGroups.SubnetGroups {
			arn := fmt.Sprintf("arn:%s:dax:%s::subnetgroup:%s", describeCtx.Partition, describeCtx.Region, *subnetGroup.SubnetGroupName)
			resource := Resource{
				Region: describeCtx.KaytuRegion,
				Name:   *subnetGroup.SubnetGroupName,
				ARN:    arn,
				Description: model.DAXSubnetGroupDescription{
					SubnetGroup: subnetGroup,
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

		return subnetGroups.NextToken, nil
	})
	if err != nil {
		return nil, err
	}

	return values, nil
}
