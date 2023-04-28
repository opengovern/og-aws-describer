package describer

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodbstreams"
	"github.com/kaytu-io/kaytu-aws-describer/aws/model"
)

func DynamoDbTable(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	client := dynamodb.NewFromConfig(cfg)
	paginator := dynamodb.NewListTablesPaginator(client, &dynamodb.ListTablesInput{})

	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, table := range page.TableNames {
			// This prevents Implicit memory aliasing in for loop
			table := table
			v, err := client.DescribeTable(ctx, &dynamodb.DescribeTableInput{
				TableName: &table,
			})
			if err != nil {
				return nil, err
			}

			continuousBackup, err := client.DescribeContinuousBackups(ctx, &dynamodb.DescribeContinuousBackupsInput{
				TableName: &table,
			})
			if err != nil {
				return nil, err
			}

			tags, err := client.ListTagsOfResource(ctx, &dynamodb.ListTagsOfResourceInput{
				ResourceArn: v.Table.TableArn,
			})
			if err != nil {
				return nil, err
			}

			resource := Resource{
				ARN:  *v.Table.TableArn,
				Name: *v.Table.TableName,
				Description: model.DynamoDbTableDescription{
					Table:            v.Table,
					ContinuousBackup: continuousBackup.ContinuousBackupsDescription,
					Tags:             tags.Tags,
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

func DynamoDbGlobalSecondaryIndex(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	client := dynamodb.NewFromConfig(cfg)
	paginator := dynamodb.NewListTablesPaginator(client, &dynamodb.ListTablesInput{})

	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, table := range page.TableNames {
			// This prevents Implicit memory aliasing in for loop
			table := table
			tableOutput, err := client.DescribeTable(ctx, &dynamodb.DescribeTableInput{
				TableName: &table,
			})
			if err != nil {
				return nil, err
			}

			for _, v := range tableOutput.Table.GlobalSecondaryIndexes {
				resource := Resource{
					ARN:  *v.IndexArn,
					Name: *v.IndexName,
					Description: model.DynamoDbGlobalSecondaryIndexDescription{
						GlobalSecondaryIndex: v,
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
	}

	return values, nil
}

func GetDynamoDbGlobalSecondaryIndex(ctx context.Context, cfg aws.Config, fields map[string]string) ([]Resource, error) {
	tableName := fields["name"]

	client := dynamodb.NewFromConfig(cfg)

	var values []Resource
	tableOutput, err := client.DescribeTable(ctx, &dynamodb.DescribeTableInput{
		TableName: &tableName,
	})
	if err != nil {
		return nil, err
	}

	for _, v := range tableOutput.Table.GlobalSecondaryIndexes {
		values = append(values, Resource{
			ARN:  *v.IndexArn,
			Name: *v.IndexName,
			Description: model.DynamoDbGlobalSecondaryIndexDescription{
				GlobalSecondaryIndex: v,
			},
		})
	}

	return values, nil
}

func DynamoDbLocalSecondaryIndex(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	client := dynamodb.NewFromConfig(cfg)
	paginator := dynamodb.NewListTablesPaginator(client, &dynamodb.ListTablesInput{})

	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, table := range page.TableNames {
			// This prevents Implicit memory aliasing in for loop
			table := table
			tableOutput, err := client.DescribeTable(ctx, &dynamodb.DescribeTableInput{
				TableName: &table,
			})
			if err != nil {
				return nil, err
			}

			for _, v := range tableOutput.Table.LocalSecondaryIndexes {
				resource := Resource{
					ARN:  *v.IndexArn,
					Name: *v.IndexName,
					Description: model.DynamoDbLocalSecondaryIndexDescription{
						LocalSecondaryIndex: v,
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
	}

	return values, nil
}

func DynamoDbStream(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	client := dynamodbstreams.NewFromConfig(cfg)
	var values []Resource
	var lastArn *string = nil
	for {
		streams, err := client.ListStreams(ctx, &dynamodbstreams.ListStreamsInput{
			ExclusiveStartStreamArn: lastArn,
			Limit:                   aws.Int32(100),
		})
		if len(streams.Streams) == 0 {
			break
		}

		if err != nil {
			return nil, err
		}

		for _, v := range streams.Streams {
			resource := Resource{
				ARN:  *v.StreamArn,
				Name: *v.StreamLabel,
				Description: model.DynamoDbStreamDescription{
					Stream: v,
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
		if streams.LastEvaluatedStreamArn == nil {
			break
		}
		lastArn = streams.LastEvaluatedStreamArn
	}

	return values, nil
}

func DynamoDbBackUp(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	client := dynamodb.NewFromConfig(cfg)
	var values []Resource
	var lastArn *string = nil
	for {
		backups, err := client.ListBackups(ctx, &dynamodb.ListBackupsInput{
			ExclusiveStartBackupArn: lastArn,
			Limit:                   aws.Int32(100),
		})
		if err != nil {
			if isErr(err, "ValidationException") {
				return nil, nil
			}
			return nil, err
		}
		if len(backups.BackupSummaries) == 0 {
			break
		}

		for _, v := range backups.BackupSummaries {
			resource := Resource{
				ARN:  *v.BackupArn,
				Name: *v.BackupName,
				Description: model.DynamoDbBackupDescription{
					Backup: v,
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

		if backups.LastEvaluatedBackupArn == nil {
			break
		}
		lastArn = backups.LastEvaluatedBackupArn
	}

	return values, nil
}

func DynamoDbGlobalTable(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	client := dynamodb.NewFromConfig(cfg)

	var values []Resource
	var last *string = nil
	for {
		globalTables, err := client.ListGlobalTables(ctx, &dynamodb.ListGlobalTablesInput{
			ExclusiveStartGlobalTableName: last,
			Limit:                         aws.Int32(100),
		})
		if err != nil {
			if isErr(err, "ResourceNotFoundException") {
				return nil, nil
			}
			return nil, err
		}
		if len(globalTables.GlobalTables) == 0 {
			break
		}

		for _, table := range globalTables.GlobalTables {
			globalTable, err := client.DescribeGlobalTable(ctx, &dynamodb.DescribeGlobalTableInput{
				GlobalTableName: table.GlobalTableName,
			})
			if err != nil {
				if isErr(err, "ResourceNotFoundException") {
					continue
				}
				return nil, err
			}
			resource := Resource{
				ARN:  *globalTable.GlobalTableDescription.GlobalTableArn,
				Name: *globalTable.GlobalTableDescription.GlobalTableName,
				Description: model.DynamoDbGlobalTableDescription{
					GlobalTable: *globalTable.GlobalTableDescription,
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

		if globalTables.LastEvaluatedGlobalTableName == nil {
			break
		}
		last = globalTables.LastEvaluatedGlobalTableName
	}

	return values, nil
}

func DynamoDbTableExport(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	client := dynamodb.NewFromConfig(cfg)
	paginator := dynamodb.NewListExportsPaginator(client, &dynamodb.ListExportsInput{})

	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, exportSummary := range page.ExportSummaries {
			export, err := client.DescribeExport(ctx, &dynamodb.DescribeExportInput{
				ExportArn: exportSummary.ExportArn,
			})
			if err != nil {
				return nil, err
			}

			resource := Resource{
				ARN: *export.ExportDescription.ExportArn,
				Description: model.DynamoDbTableExportDescription{
					Export: *export.ExportDescription,
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

func DynamoDBMetricAccountProvisionedReadCapacityUtilization(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	metrics, err := listCloudWatchMetricStatistics(ctx, cfg, "5_MIN", "AWS/DynamoDB", "AccountProvisionedWriteCapacityUtilization", "", "")
	if err != nil {
		return nil, err
	}

	var values []Resource
	for _, metric := range metrics {
		resource := Resource{
			ID: fmt.Sprintf("dynamodb-metric-account-provisioned-read-capacity-utilization-%s-%s-%s", *metric.DimensionName, *metric.DimensionValue, metric.Timestamp.Format(time.RFC3339)),
			Description: model.DynamoDBMetricAccountProvisionedReadCapacityUtilizationDescription{
				CloudWatchMetricRow: metric,
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

func DynamoDBMetricAccountProvisionedWriteCapacityUtilization(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	metrics, err := listCloudWatchMetricStatistics(ctx, cfg, "5_MIN", "AWS/DynamoDB", "AccountProvisionedWriteCapacityUtilization", "", "")
	if err != nil {
		return nil, err
	}

	var values []Resource
	for _, metric := range metrics {
		resource := Resource{
			ID: fmt.Sprintf("dynamodb-metric-account-provisioned-write-capacity-utilization-%s-%s-%s", *metric.DimensionName, *metric.DimensionValue, metric.Timestamp.Format(time.RFC3339)),
			Description: model.DynamoDBMetricAccountProvisionedWriteCapacityUtilizationDescription{
				CloudWatchMetricRow: metric,
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
