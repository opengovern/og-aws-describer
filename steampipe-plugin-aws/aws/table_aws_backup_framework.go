package aws

import (
	"context"

	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

// // TABLE DEFINITION
func tableAwsBackupFramework(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_backup_framework",
		Description: "AWS Backup Framework",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("framework_name"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"InvalidParameterValueException"}),
			},
			Hydrate: kaytu.GetBackupFramework,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListBackupFramework,
		},

		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "framework_name",
				Description: "The unique name of a backup framework.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Framework.FrameworkName")},
			{
				Name:        "arn",
				Description: "An Amazon Resource Name (ARN) that uniquely identifies a backup framework resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Framework.FrameworkArn"),
			},
			{
				Name:        "framework_description",
				Description: "An optional description of the backup framework.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Framework.FrameworkDescription")},
			{
				Name:        "deployment_status",
				Description: "The deployment status of a backup framework.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Framework.DeploymentStatus")},
			{
				Name:        "creation_time",
				Description: "The date and time that a framework was created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Framework.CreationTime")},
			{
				Name:        "number_of_controls",
				Description: "The number of controls contained by the framework.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.From(getNumberOfControls)},
			{
				Name:        "framework_status",
				Description: "The framework status based on recording statuses for resources governed by the framework (ACTIVE | PARTIALLY_ACTIVE | INACTIVE | UNAVAILABLE).",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Framework.FrameworkStatus")},
			{
				Name:        "framework_controls",
				Description: "A list of the controls that make up the framework. Each control in the list has a name, input parameters, and scope.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Framework.FrameworkControls"),
			},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Tags")},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Framework.FrameworkName")},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Framework.FrameworkArn").Transform(transform.EnsureStringArray),
			},
		}),
	}
}

func getNumberOfControls(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	framework := d.HydrateItem.(kaytu.BackupFramework).Description.Framework
	return len(framework.FrameworkControls), nil
}
