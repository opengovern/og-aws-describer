package aws

import (
	"context"

	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAwsElasticBeanstalkApplication(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_elastic_beanstalk_application",
		Description: "AWS Elastic Beanstalk Application",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("name"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"ResourceNotFoundException"}),
			},
			Hydrate: opengovernance.GetElasticBeanstalkApplication,
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListElasticBeanstalkApplication,
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the application.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Application.ApplicationName")},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the application.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Application.ApplicationArn"),
			},
			{
				Name:        "description",
				Description: "User-defined description of the application.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Application.Description")},
			{
				Name:        "date_created",
				Description: "The date when the application was created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Application.DateCreated")},
			{
				Name:        "date_updated",
				Description: "The date when the application was last modified.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Application.DateUpdated")},
			{
				Name:        "configuration_templates",
				Description: "The names of the configuration templates associated with this application.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Application.ConfigurationTemplates")},
			{
				Name:        "versions",
				Description: "The names of the versions for this application.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Application.Versions")},
			{
				Name:        "resource_lifecycle_config",
				Description: "The lifecycle settings for the application.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Application.ResourceLifecycleConfig")},
			{
				Name:        "tags_src",
				Description: "A list of tags assigned to the application.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Tags")},

			// Standard columns for all tables
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Application.ApplicationName")},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(getElasticBeanstalkApplicationTurbotTags),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Application.ApplicationArn").Transform(arnToAkas),
			},
		}),
	}
}

func getElasticBeanstalkApplicationTurbotTags(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	plugin.Logger(ctx).Trace("getElasticBeanstalkApplicationTurbotTags")
	tagList := d.HydrateItem.(opengovernance.ElasticBeanstalkApplication).Description.Tags

	var turbotTagsMap map[string]string
	if tagList != nil {
		turbotTagsMap = map[string]string{}
		for _, i := range tagList {
			turbotTagsMap[*i.Key] = *i.Value
		}
	}

	return turbotTagsMap, nil
}
