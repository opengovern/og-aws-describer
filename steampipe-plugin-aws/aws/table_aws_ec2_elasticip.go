package aws

import (
	"context"
	"github.com/kaytu-io/og-aws-describer/pkg/opengovernance-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAwsEC2ElasticIP(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_ec2_elasticip",
		Description: "AWS EC2 ElasticIP",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"), //TODO: change this to the primary key columns in model.go
			Hydrate:    opengovernance.GetEC2ElasticIP,
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListEC2ElasticIP,
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "id",
				Description: "The id of the elasticip.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Address.AllocationId")},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Address.AllocationId")},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Address.AllocationId").Transform(arnToAkas),
			},
		}),
	}
}
