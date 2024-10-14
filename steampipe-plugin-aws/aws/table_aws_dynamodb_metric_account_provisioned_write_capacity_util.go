package aws

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

//// TABLE DEFINITION

func tableAwsDynamoDBMetricAccountProvisionedWriteCapacityUtilization(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_dynamodb_metric_account_provisioned_write_capacity_util",
		Description: "AWS DynamoDB Metric Account Provisioned Write Capacity Utilization",
		List:        &plugin.ListConfig{
			//Hydrate: opengovernance.ListDynamoDBMetricAccountProvisionedWriteCapacityUtilization,
		},

		Columns: awsKaytuRegionalColumns(kaytuCwMetricColumns([]*plugin.Column{})),
	}
}
