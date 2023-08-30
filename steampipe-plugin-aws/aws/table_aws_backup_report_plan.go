package aws

import (
	"context"
	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"

	backupv1 "github.com/aws/aws-sdk-go/service/backup"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsBackupReportPlan(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_backup_report_plan",
		Description: "AWS Backup Report Plan",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("report_plan_name"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"InvalidParameterValueException", "ResourceNotFoundException"}),
			},
			Hydrate: kaytu.GetBackupReportPlan,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListBackupReportPlan,
		},
		GetMatrixItemFunc: SupportedRegionMatrix(backupv1.EndpointsID),
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "arn",
				Description: "An Amazon Resource Name (ARN) that uniquely identifies a resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ReportPlan.ReportPlanArn"),
			},
			{
				Name:        "report_plan_name",
				Description: "The unique name of the report plan.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ReportPlan.ReportPlanName"),
			},
			{
				Name:        "description",
				Description: "An optional description of the report plan with a maximum 1,024 characters.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ReportPlan.ReportPlanDescription"),
			},
			{
				Name:        "creation_time",
				Description: "The date and time that a report plan is created, in Unix format and Coordinated Universal Time (UTC).",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.ReportPlan.CreationTime"),
			},
			{
				Name:        "deployment_status",
				Description: "The deployment status of a report plan. The statuses are CREATE_IN_PROGRESS, UPDATE_IN_PROGRESS, DELETE_IN_PROGRESS, and COMPLETED.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ReportPlan.DeploymentStatus"),
			},
			{
				Name:        "last_attempted_execution_time",
				Description: "The date and time that a report job associated with this report plan last attempted to run, in Unix format and Coordinated Universal Time (UTC).",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.ReportPlan.LastAttemptedExecutionTime"),
			},
			{
				Name:        "last_successful_execution_time",
				Description: "The date and time that a report job associated with this report plan last successfully ran, in Unix format and Coordinated Universal Time (UTC).",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.ReportPlan.LastSuccessfulExecutionTime"),
			},

			// JSON Columns
			{
				Name:        "report_delivery_channel",
				Description: "Contains information about where and how to deliver your reports, specifically your Amazon S3 bucket name, S3 key prefix, and the formats of your reports.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ReportPlan.ReportDeliveryChannel"),
			},
			{
				Name:        "report_setting",
				Description: "Identifies the report template for the report. Reports are built using a report template.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ReportPlan.ReportSetting"),
			},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ReportPlan.ReportPlanName"),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ReportPlan.ReportPlanArn").Transform(transform.EnsureStringArray),
			},
		}),
	}
}
