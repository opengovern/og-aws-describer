package aws

import (
	"context"

	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsDaxSubnetGroup(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_dax_subnet_group",
		Description: "AWS DAX Subnet Group",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListDAXSubnetGroup,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"SubnetGroupNotFoundFault"}),
			},
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:    "subnet_group_name",
					Require: plugin.Optional,
				},
			},
		},

		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "subnet_group_name",
				Description: "The name of the subnet group.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.SubnetGroup.SubnetGroupName")},
			{
				Name:        "description",
				Description: "The description of the subnet group.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.SubnetGroup.Description")},
			{
				Name:        "vpc_id",
				Description: "The Amazon Virtual Private Cloud identifier (VPC ID) of the subnet group.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.SubnetGroup.VpcId")},
			{
				Name:        "subnets",
				Description: "A list of subnets associated with the subnet group.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.SubnetGroup.Subnets")},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.SubnetGroup.SubnetGroupName")},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("ARN").Transform(arnToAkas)},
		}),
	}
}
