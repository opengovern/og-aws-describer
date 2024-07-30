package aws

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAwsDynamoDBTable(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_dynamodb_table",
		Description: "AWS DynamoDB Table",
		DefaultIgnoreConfig: &plugin.IgnoreConfig{
			ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"TableNotFoundException"}),
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("name"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"ResourceNotFoundException"}),
			},
			Hydrate: kaytu.GetDynamoDbTable,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListDynamoDbTable,
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
				Description: "The name of the table.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Table.TableName")},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) that uniquely identifies the table.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Table.TableArn")},
			{
				Name:        "table_id",
				Description: "Unique identifier for the table.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Table.TableId")},
			{
				Name:        "creation_date_time",
				Description: "The date and time when the table was created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Table.CreationDateTime")},
			{
				Name:        "table_class",
				Description: "The table class of the specified table. Valid values are STANDARD and STANDARD_INFREQUENT_ACCESS.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Table.TableClassSummary.TableClass"),
			},
			{
				Name:        "table_status",
				Description: "The current state of the table.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Table.TableStatus")},
			{
				Name:        "billing_mode",
				Description: "Controls how AWS charges for read and write throughput and manage capacity.",
				Type:        proto.ColumnType_STRING,
				Default:     "PROVISIONED",
				Transform:   transform.FromField("Description.Table.BillingModeSummary.BillingMode").Transform(getTableBillingMode),
				// If it is not available then it should default to  "PROVISIONED"
				// Billing mode can only be PAY_PER_REQUEST or PROVISIONED
			},
			{
				Name:        "item_count",
				Description: "Number of items in the table. Note that this is an approximate value.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Table.ItemCount")},
			{
				Name:        "global_table_version",
				Description: "Represents the version of global tables (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/GlobalTables.html) in use, if the table is replicated across AWS Regions.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Table.GlobalTableVersion")},
			{
				Name:        "read_capacity",
				Description: "The maximum number of strongly consistent reads consumed per second before DynamoDB returns a ThrottlingException.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Table.ProvisionedThroughput.ReadCapacityUnits")},
			{
				Name:        "write_capacity",
				Description: "The maximum number of writes consumed per second before DynamoDB returns a ThrottlingException.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Table.ProvisionedThroughput.WriteCapacityUnits")},
			{
				Name:        "latest_stream_arn",
				Description: "The Amazon Resource Name (ARN) that uniquely identifies the latest stream for this table.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Table.LatestStreamArn")},
			{
				Name:        "latest_stream_label",
				Description: "A timestamp, in ISO 8601 format, for this stream.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Table.LatestStreamLabel")},
			{
				Name:        "table_size_bytes",
				Description: "Size of the table in bytes. Note that this is an approximate value.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Table.TableSizeBytes")},
			{
				Name:        "archival_summary",
				Description: "Contains information about the table archive.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Table.ArchivalSummary")},
			{
				Name:        "attribute_definitions",
				Description: "An array of AttributeDefinition objects. Each of these objects describes one attribute in the table and index key schema.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Table.AttributeDefinitions")},
			{
				Name:        "key_schema",
				Description: "The primary key structure for the table.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Table.KeySchema")},
			{
				Name:        "sse_description",
				Description: "The description of the server-side encryption status on the specified table.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Table.SSEDescription")},
			{
				Name:        "deletion_protection_enabled",
				Description: "Indicates whether deletion protection is enabled (true) or disabled (false) on the table.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Table.DeletionProtectionEnabled"),
			},
			{
				Name:        "continuous_backups_status",
				Description: "The continuous backups status of the table. ContinuousBackupsStatus can be one of the following states: ENABLED, DISABLED.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ContinuousBackup.ContinuousBackupsStatus")},
			{
				Name:        "streaming_destination",
				Description: "Provides information about the status of Kinesis streaming.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.StreamingDestination")},
			{
				Name:        "point_in_time_recovery_description",
				Description: "The description of the point in time recovery settings applied to the table.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ContinuousBackup.PointInTimeRecoveryDescription")},
			{
				Name:        "tags_src",
				Description: "A list of tags assigned to the table.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Tags"),
			},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(getTableTurbotTags),
			},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Table.TableName")},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Table.TableArn").Transform(transform.EnsureStringArray),
			},
		}),
	}
}

//// LIST FUNCTION

//// HYDRATE FUNCTIONS

//// TRANSFORM FUNCTIONS

func getTableBillingMode(_ context.Context, d *transform.TransformData) (interface{}, error) {
	billingMode := types.BillingModeProvisioned
	if d.Value != nil {
		billingMode = d.Value.(types.BillingMode)
	}

	return billingMode, nil
}

func getTableTurbotTags(_ context.Context, d *transform.TransformData) (interface{}, error) {
	tags := d.HydrateItem.(kaytu.DynamoDbTable).Description.Tags

	// Mapping the resource tags inside turbotTags
	var turbotTagsMap map[string]string
	if tags != nil {
		turbotTagsMap = map[string]string{}
		for _, i := range tags {
			turbotTagsMap[*i.Key] = *i.Value
		}
	}

	return turbotTagsMap, nil
}
