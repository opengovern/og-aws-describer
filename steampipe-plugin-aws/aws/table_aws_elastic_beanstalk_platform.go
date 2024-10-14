package aws

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAwsElasticBeanstalkPlatform(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_elastic_beanstalk_platform",
		Description: "AWS ElasticBeanstalk Platform",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("platform_name"),
			//Hydrate:    opengovernance.GetElasticBeanstalkPlatform,
		},
		List: &plugin.ListConfig{
			//Hydrate: opengovernance.ListElasticBeanstalkPlatform,
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "platform_name",
				Description: "The name of the platform.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Platform.PlatformName")},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the platform",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Platform.PlatformArn")},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Platform.PlatformName")},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Platform.PlatformArn").Transform(arnToAkas),
			},
		}),
	}
}
