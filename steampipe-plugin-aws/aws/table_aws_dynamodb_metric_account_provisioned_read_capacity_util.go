package aws

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

//// TABLE DEFINITION

func tableAwsDynamoDBMetricAccountProvisionedReadCapacityUtilization(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_dynamodb_metric_account_provisioned_read_capacity_util",
		Description: "AWS DynamoDB Metric Account Provisioned Read Capacity Utilization",
		List:        &plugin.ListConfig{
			//Hydrate: opengovernance.ListDynamoDBMetricAccountProvisionedReadCapacityUtilization,
		},

		Columns: awsKaytuRegionalColumns(kaytuCwMetricColumns([]*plugin.Column{})),
	}
}
