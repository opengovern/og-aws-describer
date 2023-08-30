package aws

import (
	"context"
	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAwsCloudwatchLogGroup(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_cloudwatch_log_group",
		Description: "AWS CloudWatch Log Group",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("name"),
			Hydrate:    kaytu.GetCloudWatchLogsLogGroup,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListCloudWatchLogsLogGroup,
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:    "name",
					Require: plugin.Optional,
				},
			},
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the log group.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.LogGroup.LogGroupName"),
			},

			// Most CloudWatch APIs' inputs only accept a CloudWatch log group ARN without ":" at the end, but the
			// DescribeLogGroups API returns an ARN with ":*", which we've chosen to keep to better match what AWS shows
			// in their console and documentation.
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the log group.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.LogGroup.Arn"),
			},
			{
				Name:        "creation_time",
				Description: "The creation time of the log group.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.LogGroup.CreationTime").Transform(convertTimestamp),
			},
			{
				Name:        "kms_key_id",
				Description: "The Amazon Resource Name (ARN) of the CMK to use when encrypting log data.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.LogGroup.KmsKeyId"),
			},
			{
				Name:        "metric_filter_count",
				Description: "The number of metric filters.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.LogGroup.MetricFilterCount"),
			},
			{
				Name:        "retention_in_days",
				Description: "The number of days to retain the log events in the specified log group. Possible values are: 1, 3, 5, 7, 14, 30, 60, 90, 120, 150, 180, 365, 400, 545, 731, 1827, and 3653.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.LogGroup.RetentionInDays"),
			},
			{
				Name:        "stored_bytes",
				Description: "The number of bytes stored.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.LogGroup.StoredBytes"),
			},
			{
				Name:        "data_protection",
				Description: "Log group data protection policy information.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DataProtection"),
			},
			{
				Name:        "data_protection_policy",
				Description: "The data protection policy document for a log group.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DataProtection.PolicyDocument"),
			},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.LogGroup.LogGroupName"),
			},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Tags"),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.LogGroup.Arn").Transform(arnToAkas),
			},
		}),
	}
}

//// LIST FUNCTION
//// HYDRATE FUNCTIONS
