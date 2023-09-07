package aws

import (
	"context"
	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAwsEC2LocalGateway(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_ec2_localgateway",
		Description: "AWS EC2 LocalGateway",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("arn"), //TODO: change this to the primary key columns in model.go
			Hydrate:    kaytu.GetEC2LocalGateway,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListEC2LocalGateway,
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "id",
				Description: "The id of the localgateway.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.LocalGateway.LocalGatewayId")},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the localgateway",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("ARN")},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.LocalGateway.LocalGatewayId")},
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
