package aws

import (
	"context"
	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"

	"github.com/aws/aws-sdk-go-v2/service/databasemigrationservice"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsDmsReplicationTask(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_dms_replication_task",
		Description: "AWS DMS Replication Task",
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListDMSReplicationTask,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"InvalidParameterValueException", "ResourceNotFoundFault"}),
			},
		},
		Columns: awsRegionalColumns([]*plugin.Column{
			{
				Name:        "replication_task_identifier",
				Description: "The user-assigned replication task identifier or name.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ReplicationTask.ReplicationTaskIdentifier"),
			},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the replication task.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ReplicationTask.ReplicationTaskArn"),
			},
			{
				Name:        "cdc_start_position",
				Description: "Indicates when you want a change data capture (CDC) operation to start.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ReplicationTask.CdcStartPosition"),
			},
			{
				Name:        "cdc_stop_position",
				Description: "Indicates when you want a change data capture (CDC) operation to stop.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ReplicationTask.CdcStopPosition"),
			},
			{
				Name:        "last_failure_message",
				Description: "The last error (failure) message generated for the replication task.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ReplicationTask.LastFailureMessage"),
			},
			{
				Name:        "migration_type",
				Description: "The type of migration.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ReplicationTask.MigrationType"),
			},
			{
				Name:        "recovery_checkpoint",
				Description: "Indicates the last checkpoint that occurred during a change data capture (CDC) operation.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ReplicationTask.RecoveryCheckpoint"),
			},
			{
				Name:        "replication_instance_arn",
				Description: "The Amazon Resource Name (ARN) of the replication instance.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ReplicationTask.ReplicationInstanceArn"),
			},
			{
				Name:        "replication_task_creation_date",
				Description: "The date the replication task was created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.ReplicationTask.ReplicationTaskCreationDate"),
			},
			{
				Name:        "replication_task_start_date",
				Description: "The date the replication task is scheduled to start.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.ReplicationTask.ReplicationTaskStartDate"),
			},
			{
				Name:        "source_endpoint_arn",
				Description: "The Amazon Resource Name (ARN) that uniquely identifies the endpoint.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ReplicationTask.SourceEndpointArn"),
			},
			{
				Name:        "status",
				Description: "The status of the replication task.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ReplicationTask.Status"),
			},
			{
				Name:        "stop_reason",
				Description: "The reason the replication task was stopped.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ReplicationTask.StopReason"),
			},
			{
				Name:        "table_mappings",
				Description: "Table mappings specified in the task.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ReplicationTask.TableMappings"),
			},
			{
				Name:        "target_endpoint_arn",
				Description: "The ARN that uniquely identifies the endpoint.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ReplicationTask.TargetEndpointArn"),
			},
			{
				Name:        "target_replication_instance_arn",
				Description: "The ARN of the replication instance to which this task is moved in response to running the MoveReplicationTask operation.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ReplicationTask.TargetReplicationInstanceArn"),
			},
			{
				Name:        "task_data",
				Description: "Supplemental information that the task requires to migrate the data for certain source and target endpoints.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ReplicationTask.TaskData"),
			},
			{
				Name:        "replication_task_settings",
				Description: "The settings for the replication task.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ReplicationTask.ReplicationTaskSettings"),
			},
			{
				Name:        "replication_task_stats",
				Description: "The statistics for the task, including elapsed time, tables loaded, and table errors.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ReplicationTask.ReplicationTaskStats"),
			},
			{
				Name:        "tags_src",
				Description: "A list of tags currently associated with the replication instance.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Tags"),
			},

			// Steampipe Standard Columns
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ReplicationTask.ReplicationTaskIdentifier"),
			},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Tags").Transform(dmsReplicationTaskTagListToTagsMap),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ReplicationTask.ReplicationTaskArn").Transform(transform.EnsureStringArray),
			},
		}),
	}
}

func dmsReplicationTaskTagListToTagsMap(_ context.Context, d *transform.TransformData) (interface{}, error) {
	data := d.HydrateItem.(*databasemigrationservice.ListTagsForResourceOutput)

	// Mapping the resource tags inside turbotTags
	if data.TagList != nil {
		turbotTagsMap := map[string]string{}
		for _, i := range data.TagList {
			turbotTagsMap[*i.Key] = *i.Value
		}
		return turbotTagsMap, nil
	}
	return nil, nil
}
