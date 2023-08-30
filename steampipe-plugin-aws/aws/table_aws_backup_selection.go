package aws

import (
	"context"
	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsBackupSelection(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_backup_selection",
		Description: "AWS Backup Selection",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"backup_plan_id", "selection_id"}),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"InvalidParameterValue", "InvalidParameterValueException"}),
			},
			Hydrate: kaytu.GetBackupSelection,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListBackupSelection,
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "selection_name",
				Description: "The display name of a resource selection document.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.BackupSelection.SelectionName")},
			{
				Name:        "selection_id",
				Description: "Uniquely identifies a request to assign a set of resources to a backup plan.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.BackupSelection.SelectionId")},
			{
				Name:        "backup_plan_id",
				Description: "An ID that uniquely identifies a backup plan.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.BackupSelection.BackupPlanId")},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) specifying the backup selection.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("ARN").Transform(arnToAkas),
			},
			{
				Name:        "creation_date",
				Description: "The date and time a resource backup plan is created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.BackupSelection.CreationDate")},
			{
				Name:        "creator_request_id",
				Description: "An unique string that identifies the request and allows failed requests to be retried without the risk of running the operation twice.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.BackupSelection.CreatorRequestId")},
			{
				Name:        "iam_role_arn",
				Description: "Specifies the IAM role Amazon Resource Name (ARN) to create the target recovery point.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.BackupSelection.IamRoleArn")},
			{
				Name:        "list_of_tags",
				Description: "An array of conditions used to specify a set of resources to assign to a backup plan.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ListOfTags"),
			},
			{
				Name:        "resources",
				Description: "Contains a list of BackupOptions for a resource type.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Resources"),
			},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.BackupSelection.SelectionName"),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("ARN").Transform(transform.EnsureStringArray),
			},
		}),
	}
}
