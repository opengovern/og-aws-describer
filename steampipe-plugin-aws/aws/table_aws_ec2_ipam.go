package aws

import (
	"context"

	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAwsEC2Ipam(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_ec2_ipam",
		Description: "AWS EC2 Ipam",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("ipam_id"),
			Hydrate:    kaytu.GetEC2Ipam,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListEC2Ipam,
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "ipam_id",
				Description: "The id of the ipam.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Ipam.IpamId")},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the ipam",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Ipam.IpamArn")},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Ipam.IpamId")},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(getEC2IpamTurbotTags),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Ipam.IpamArn").Transform(arnToAkas),
			},
		}),
	}
}

//// TRANSFORM FUNCTIONS

func getEC2IpamTurbotTags(_ context.Context, d *transform.TransformData) (interface{}, error) {
	tags := d.HydrateItem.(kaytu.EC2Ipam).Description.Ipam.Tags
	return ec2V2TagsToMap(tags)
}
