package aws

import (
	"context"
	"fmt"

	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"
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
			Hydrate:    kaytu.GetEC2Host,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListEC2Host,
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
				Transform:   transform.From(getEc2HostArn)},
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
				Transform:   transform.From(getEc2HostArn).Transform(arnToAkas),
			},
		}),
	}
}

//// TRANSFORM FUNCTIONS

func getEc2HostTurbotTags(_ context.Context, d *transform.TransformData) (interface{}, error) {
	tags := d.HydrateItem.(kaytu.EC2Host).Description.Host.Tags
	return ec2V2TagsToMap(tags)
}

func getEc2HostArn(_ context.Context, d *transform.TransformData) (interface{}, error) {
	host := d.HydrateItem.(kaytu.EC2Host).Description.Host
	metadata := d.HydrateItem.(kaytu.EC2Host).Metadata
	arn := fmt.Sprintf("arn:%s:ec2:%s:%s:dedicated-host/%s", metadata.Partition, metadata.Region, metadata.AccountID, *host.HostId)
	return arn, nil
}
