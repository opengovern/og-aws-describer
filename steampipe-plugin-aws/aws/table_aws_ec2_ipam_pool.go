package aws

import (
	"context"

	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAwsEC2IpamPool(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_ec2_ipam_pool",
		Description: "AWS EC2 IpamPool",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("ipam_pool_id"),
			Hydrate:    kaytu.GetEC2IpamPool,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListEC2IpamPool,
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "ipam_pool_id",
				Description: "The id of the ipam pool.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.IpamPool.IpamPoolId")},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the ipam pool",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.IpamPool.IpamPoolArn")},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.IpamPool.IpamPoolId")},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(getEC2IpamPoolTurbotTags),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.IpamPool.IpamPoolArn").Transform(arnToAkas),
			},
		}),
	}
}

//// TRANSFORM FUNCTIONS

func getEC2IpamPoolTurbotTags(_ context.Context, d *transform.TransformData) (interface{}, error) {
	tags := d.HydrateItem.(kaytu.EC2IpamPool).Description.IpamPool.Tags
	return ec2V2TagsToMap(tags)
}
