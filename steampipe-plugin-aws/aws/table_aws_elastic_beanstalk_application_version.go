package aws

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk"
	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAwsElasticBeanstalkApplicationVersion(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_elastic_beanstalk_application_version",
		Description: "AWS Elastic Beanstalk Application Version",
		Get: &plugin.GetConfig{
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"ResourceNotFoundException"}),
			},
			Hydrate: kaytu.GetElasticBeanstalkApplicationVersion,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListElasticBeanstalkApplicationVersion,
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "application_name",
				Description: "The name of the application to which the application version belongs.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ApplicationVersion.ApplicationName"),
			},
			{
				Name:        "application_version_arn",
				Description: "The Amazon Resource Name (ARN) of the application version.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ApplicationVersion.ApplicationVersionArn"),
			},
			{
				Name:        "build_arn",
				Description: "Reference to the artifact from the AWS CodeBuild build.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ApplicationVersion.BuildArn"),
			},
			{
				Name:        "date_created",
				Description: "The creation date of the application version.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.ApplicationVersion.DateCreated"),
			},
			{
				Name:        "date_updated",
				Description: "The last modified date of the application version.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.ApplicationVersion.DateUpdated"),
			},
			{
				Name:        "description",
				Description: "The description of the application version.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ApplicationVersion.Description"),
			},
			{
				Name:        "status",
				Description: "The processing status of the application version. Reflects the state of the application version during its creation.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ApplicationVersion.Status"),
			},
			{
				Name:        "version_label",
				Description: "A unique identifier for the application version.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ApplicationVersion.VersionLabel"),
			},
			{
				Name:        "source_build_information",
				Description: "Information about the source code for the application version if the source code was retrieved from AWS CodeCommit.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ApplicationVersion.SourceBuildInformation"),
			},
			{
				Name:        "source_bundle",
				Description: "The storage location of the application version's source bundle in Amazon S3.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ApplicationVersion.SourceBundle"),
			},
			{
				Name:        "tags_src",
				Description: "A list of tags assigned to the application.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Tags"),
			},
			// Standard columns for all tables
			{
				Name:        "title",
				Description: "A title for the resource, typically the resource name.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ApplicationVersion.VersionLabel"),
			},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Tags").Transform(handleElasticBeanstalkApplicationVersionTurbotTags),
			},
			{
				Name:        "akas",
				Description: "Array of globally unique identifier strings (also known as) for the resource.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ApplicationVersion.ApplicationVersionArn").Transform(transform.EnsureStringArray),
			},
		}),
	}
}

func handleElasticBeanstalkApplicationVersionTurbotTags(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	tagList := d.HydrateItem.(*elasticbeanstalk.ListTagsForResourceOutput)

	// Mapping the resource tags inside turbotTags
	var turbotTagsMap map[string]string
	if tagList != nil && len(tagList.ResourceTags) > 0 {
		turbotTagsMap = map[string]string{}
		for _, i := range tagList.ResourceTags {
			turbotTagsMap[*i.Key] = *i.Value
		}
	}

	return turbotTagsMap, nil
}
