package aws

import (
	"context"
	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAwsSESConfigurationSet(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_ses_configuration_set",
		Description: "AWS SES ConfigurationSet",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("name"),
			Hydrate:    kaytu.GetSESConfigurationSet,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListSESConfigurationSet,
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the configuration set.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ConfigurationSet.Name")},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the configuration set",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("ARN")},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ConfigurationSet.Name")},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("ARN").Transform(arnToAkas),
			},
		}),
	}
}
