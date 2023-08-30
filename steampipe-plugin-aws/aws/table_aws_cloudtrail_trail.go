package aws

import (
	"context"
	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsCloudtrailTrail(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_cloudtrail_trail",
		Description: "AWS CloudTrail Trail",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AnyColumn([]string{"name", "arn"}),
			Hydrate:    kaytu.GetCloudTrailTrail,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"InvalidTrailNameException", "TrailNotFoundException", "CloudTrailARNInvalidException"}),
			},
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListCloudTrailTrail,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"InvalidTrailNameException", "TrailNotFoundException", "CloudTrailARNInvalidException"}),
			}},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the trail.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Trail.Name"),
			},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the trail.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Trail.TrailARN"),
			},
			{
				Name:        "home_region",
				Description: "The region in which the trail was created.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Trail.HomeRegion"),
			},
			{
				Name:        "is_multi_region_trail",
				Description: "Specifies whether the trail exists only in one region or exists in all regions.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Trail.IsMultiRegionTrail"),
			},
			{
				Name:        "log_file_validation_enabled",
				Description: "Specifies whether log file validation is enabled, or not.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Trail.LogFileValidationEnabled"),
			},
			{
				Name:        "is_logging",
				Description: "Specifies whether the CloudTrail is currently logging AWS API calls, or not.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.TrailStatus.IsLogging"),
			},
			{
				Name:        "log_group_arn",
				Description: "Specifies an Amazon Resource Name (ARN), a unique identifier that represents the log group to which CloudTrail logs will be delivered.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Trail.CloudWatchLogsLogGroupArn"),
			},
			{
				Name:        "cloudwatch_logs_role_arn",
				Description: "Specifies the role for the CloudWatch Logs endpoint to assume to write to a user's log group.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Trail.CloudWatchLogsRoleArn"),
			},
			{
				Name:        "has_custom_event_selectors",
				Description: "Specifies whether the trail has custom event selectors, or not.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Trail.HasCustomEventSelectors"),
			},
			{
				Name:        "has_insight_selectors",
				Description: "Specifies whether a trail has insight types specified in an InsightSelector list, or not.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Trail.HasInsightSelectors"),
			},
			{
				Name:        "include_global_service_events",
				Description: "Specifies whether to include AWS API calls from AWS global services, or not.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Trail.IncludeGlobalServiceEvents"),
			},
			{
				Name:        "is_organization_trail",
				Description: "Specifies whether the trail is an organization trail, or not.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Trail.IsOrganizationTrail"),
			},
			{
				Name:        "kms_key_id",
				Description: "Specifies the KMS key ID that encrypts the logs delivered by CloudTrail.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Trail.KmsKeyId"),
			},
			{
				Name:        "s3_bucket_name",
				Description: "Name of the Amazon S3 bucket into which CloudTrail delivers your trail files.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Trail.S3BucketName"),
			},
			{
				Name:        "s3_key_prefix",
				Description: "Specifies the Amazon S3 key prefix that comes after the name of the bucket you have designated for log file delivery.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Trail.S3KeyPrefix"),
			},
			{
				Name:        "sns_topic_arn",
				Description: "Specifies the ARN of the Amazon SNS topic that CloudTrail uses to send notifications when log files are delivered.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Trail.SnsTopicARN"),
			},

			// details of trail status
			{
				Name:        "latest_cloudwatch_logs_delivery_error",
				Description: "Displays any CloudWatch Logs error that CloudTrail encountered when attempting to deliver logs to CloudWatch Logs.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.TrailStatus.LatestCloudWatchLogsDeliveryError"),
			},
			{
				Name:        "latest_cloudwatch_logs_delivery_time",
				Description: "Displays the most recent date and time when CloudTrail delivered logs to CloudWatch Logs.",
				Type:        proto.ColumnType_TIMESTAMP,

				Transform: transform.FromField("Description.TrailStatus.LatestCloudWatchLogsDeliveryTime"),
			},
			{
				Name:        "latest_delivery_error",
				Description: "Displays any Amazon S3 error that CloudTrail encountered when attempting to deliver log files to the designated bucket.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.TrailStatus.LatestDeliveryError"),
			},
			{
				Name:        "latest_delivery_time",
				Description: "Specifies the date and time that CloudTrail last delivered log files to an account's Amazon S3 bucket.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.TrailStatus.LatestDeliveryTime"),
			},
			{
				Name:        "latest_digest_delivery_error",
				Description: "Displays any Amazon S3 error that CloudTrail encountered when attempting to deliver a digest file to the designated bucket.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.TrailStatus.LatestDigestDeliveryError"),
			},
			{
				Name:        "latest_digest_delivery_time",
				Description: "Specifies the date and time that CloudTrail last delivered a digest file to an account's Amazon S3 bucket.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.TrailStatus.LatestDigestDeliveryTime"),
			},
			{
				Name:        "latest_notification_error",
				Description: "Displays any Amazon SNS error that CloudTrail encountered when attempting to send a notification.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.TrailStatus.LatestNotificationError"),
			},
			{
				Name:        "latest_notification_time",
				Description: "Specifies the date and time of the most recent Amazon SNS notification that CloudTrail has written a new log file to an account's Amazon S3 bucket.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.TrailStatus.LatestNotificationTime"),
			},
			{
				Name:        "start_logging_time",
				Description: "Specifies the most recent date and time when CloudTrail started recording API calls for an AWS account.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.TrailStatus.StartLoggingTime"),
			},
			{
				Name:        "stop_logging_time",
				Description: "Specifies the most recent date and time when CloudTrail stopped recording API calls for an AWS account.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.TrailStatus.StopLoggingTime"),
			},
			{
				Name:        "advanced_event_selectors",
				Description: "Describes the advanced event selectors that are configured for the trail.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.AdvancedEventSelectors"),
			},
			{
				Name:        "event_selectors",
				Description: "Describes the event selectors that are configured for the trail.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.EventSelectors"),
			},
			{
				Name:        "insight_selectors",
				Description: "A JSON string that contains the insight types you want to log on a trail.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Trail.HasInsightSelectors"),
			},
			{
				Name:        "tags_src",
				Description: "A list of tags assigned to the trail.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Tags"),
			},

			// steampipe standard columns
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(getCloudtrailTrailTurbotTags),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Trail.TrailARN").Transform(arnToAkas),
			},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Trail.Name"),
			},
		}),
	}
}

//// LIST FUNCTION

//// HYDRATE FUNCTIONS

//// TRANSFORM FUNCTIONS

func getCloudtrailTrailTurbotTags(_ context.Context, d *transform.TransformData) (interface{}, error) {
	tags := d.HydrateItem.(kaytu.CloudTrailTrail).Description.Tags
	if tags == nil {
		return nil, nil
	}

	// Mapping the resource tags inside turbotTags
	turbotTagsMap := map[string]string{}
	for _, i := range tags {
		turbotTagsMap[*i.Key] = *i.Value
	}
	return turbotTagsMap, nil
}
