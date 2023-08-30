package aws

import (
	"context"
	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsBackupPlan(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_backup_plan",
		Description: "AWS Backup Plan",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("backup_plan_id"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"InvalidParameterValueException"}),
			},
			Hydrate: kaytu.GetBackupPlan,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListBackupPlan,
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The display name of a saved backup plan.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.BackupPlan.BackupPlanName"),
			},
			{
				Name:        "arn",
				Description: "An Amazon Resource Name (ARN) that uniquely identifies a backup plan.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.BackupPlan.BackupPlanArn"),
			},
			{
				Name:        "backup_plan_id",
				Description: "Specifies the id to identify a backup plan uniquely.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.BackupPlan.BackupPlanId"),
			},
			{
				Name:        "creation_date",
				Description: "The date and time a resource backup plan is created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.BackupPlan.CreationDate"),
			},
			{
				Name:        "deletion_date",
				Description: "The date and time a backup plan is deleted.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.BackupPlan.DeletionDate"),
			},
			{
				Name:        "last_execution_date",
				Description: "The last time a job to back up resources was run with this rule.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.BackupPlan.LastExecutionDate"),
			},
			{
				Name:        "creator_request_id",
				Description: "An unique string that identifies the request and allows failed requests to be retried without the risk of running the operation twice.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.BackupPlan.CreatorRequestId"),
			},
			{
				Name:        "version_id",
				Description: "Unique, randomly generated, Unicode, UTF-8 encoded strings that are at most 1,024 bytes long. Version IDs cannot be edited.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.BackupPlan.VersionId"),
			},
			{
				Name:        "backup_plan",
				Description: "Specifies the body of a backup plan.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.BackupPlan"),
			},
			{
				Name:        "advanced_backup_settings",
				Description: "Contains a list of BackupOptions for a resource type.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.BackupPlan.AdvancedBackupSettings"),
			},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.BackupPlan.BackupPlanName"),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.BackupPlan.BackupPlanArn").Transform(transform.EnsureStringArray),
			},
		}),
	}
}

//// LIST FUNCTION

//// HYDRATE FUNCTIONS
