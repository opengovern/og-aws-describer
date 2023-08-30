package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/kinesisanalyticsv2/types"
	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsKinesisAnalyticsV2Application(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_kinesisanalyticsv2_application",
		Description: "AWS Kinesis Analytics V2 Application",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("application_name"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"ResourceNotFoundException"}),
			},
			Hydrate: kaytu.GetKinesisAnalyticsV2Application,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListKinesisAnalyticsV2Application,
		},

		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "application_name",
				Description: "The name of the application.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Application.ApplicationName")},
			{
				Name:        "application_version_id",
				Description: "Provides the current application version.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Application.ApplicationVersionId")},
			{
				Name:        "application_arn",
				Description: "The ARN of the application.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Application.ApplicationARN")},
			{
				Name:        "application_status",
				Description: "The status of the application.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Application.ApplicationStatus")},
			{
				Name:        "create_timestamp",
				Description: "The current timestamp when the application was created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Application.CreateTimestamp")},
			{
				Name:        "application_description",
				Description: "The description of the application.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Application.ApplicationDescription")},
			{
				Name:        "last_update_timestamp",
				Description: "The current timestamp when the application was last updated.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Application.LastUpdateTimestamp")},
			{
				Name:        "runtime_environment",
				Description: "The runtime environment for the application.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Application.RuntimeEnvironment")},
			{
				Name:        "service_execution_role",
				Description: "Specifies the IAM role that the application uses to access external resources.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Application.ServiceExecutionRole")},
			{
				Name:        "application_configuration_description",
				Description: "Provides details about the application's Java, SQL, or Scala code and starting parameters.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Application.ApplicationConfigurationDescription")},
			{
				Name:        "cloud_watch_logging_option_descriptions",
				Description: "Describes the application Amazon CloudWatch logging options.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Application.CloudWatchLoggingOptionDescriptions")},
			{
				Name:        "tags_src",
				Description: "The key-value tags assigned to the application.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Tags")},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Application.ApplicationName")},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Tags").Transform(kinesisAnalyticsV2ApplicationTagListToTurbotTags),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Application.ApplicationARN").Transform(arnToAkas),
			},
		}),
	}
}

func kinesisAnalyticsV2ApplicationTagListToTurbotTags(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	tagList := d.Value.([]types.Tag)

	if tagList == nil {
		return nil, nil
	}

	// Mapping the resource tags inside turbotTags
	var turbotTagsMap map[string]string
	if tagList != nil {
		turbotTagsMap = map[string]string{}
		for _, i := range tagList {
			turbotTagsMap[*i.Key] = *i.Value
		}
	}

	return turbotTagsMap, nil
}
