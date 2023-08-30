package aws

import (
	"context"

	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsDaxParameterGroup(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_dax_parameter_group",
		Description: "AWS DAX Parameter Group",
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListDAXParameterGroup,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"ParameterGroupNotFoundFault"}),
			},
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:    "parameter_group_name",
					Require: plugin.Optional,
				},
			},
		},

		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "parameter_group_name",
				Description: "The name of the parameter group.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ParameterGroup.ParameterGroupName")},
			{
				Name:        "description",
				Description: "Description of the parameter group.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ParameterGroup.Description")},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ParameterGroup.ParameterGroupName")},
		}),
	}
}
