package aws

import (
	"context"
	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAwsEc2Host(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_ec2_host",
		Description: "AWS Ec2 Host",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("host_id"),
			Hydrate:    opengovernance.GetEC2Host,
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListEC2Host,
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "host_id",
				Description: "The id of the host.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Host.HostId")},
			{
				Name:        "host_arn",
				Description: "The Amazon Resource Name (ARN) of the host",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("ARN")},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Host.HostId")},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(getEc2HostTurbotTags),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("ARN").Transform(arnToAkas),
			},
		}),
	}
}

//// TRANSFORM FUNCTIONS

func getEc2HostTurbotTags(_ context.Context, d *transform.TransformData) (interface{}, error) {
	tags := d.HydrateItem.(opengovernance.EC2Host).Description.Host.Tags
	return ec2V2TagsToMap(tags)
}
