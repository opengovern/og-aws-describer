package aws

import (
	"context"

	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsSageMakerApp(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_sagemaker_app",
		Description: "AWS Sagemaker App",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"name", "app_type", "domain_id", "user_profile_name"}),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"ValidationException", "NotFoundException", "ResourceNotFound"}),
			},
			Hydrate: opengovernance.GetSageMakerApp,
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListSageMakerApp,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "user_profile_name", Require: plugin.Optional},
				{Name: "domain_id", Require: plugin.Optional},
			},
		},

		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The app name.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AppDetails.AppName")},
			{
				Name:        "app_type",
				Description: "The type of app.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AppDetails.AppType")},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the app.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DescribeAppOutput.AppArn"),
			},
			{
				Name:        "creation_time",
				Description: "A timestamp that indicates when the app was created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.AppDetails.CreationTime")},
			{
				Name:        "domain_id",
				Description: "The domain ID.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AppDetails.DomainId")},
			{
				Name:        "failure_reason",
				Description: "The failure reason.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DescribeAppOutput.FailureReason")},
			{
				Name:        "last_health_check_timestamp",
				Description: "The timestamp of the last health check.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.DescribeAppOutput.LastHealthCheckTimestamp")},
			{
				Name:        "last_user_activity_timestamp",
				Description: "The timestamp of the last user activity.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.DescribeAppOutput.LastUserActivityTimestamp")},
			{
				Name:        "status",
				Description: "The status of the app.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AppDetails.Status")},
			{
				Name:        "user_profile_name",
				Description: "The user profile name.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AppDetails.UserProfileName")},
			{
				Name:        "resource_spec",
				Description: "The instance type and the Amazon Resource Name (ARN) of the SageMaker image created on the instance.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DescribeAppOutput.ResourceSpec")},

			// Steampipe standard columns

			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AppDetails.AppName")},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,

				Transform: transform.FromField("Description.DescribeAppOutput.AppArn").Transform(transform.EnsureStringArray),

				//// LIST FUNCTION
			},
		}),
	}
}
