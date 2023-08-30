package aws

import (
	"context"

	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsDaxParameter(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_dax_parameter",
		Description: "AWS DAX Parameter",
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListDAXParameter,
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
				Name:        "parameter_name",
				Description: "The name of the parameter.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Parameter.ParameterName")},
			{
				Name:        "parameter_group_name",
				Description: "The name of the parameter group.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ParameterGroupName")},
			{
				Name:        "parameter_value",
				Description: "The value of the parameter.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Parameter.ParameterValue")},
			{
				Name:        "description",
				Description: "Description of the parameter.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Parameter.Description")},
			{
				Name:        "allowed_values",
				Description: "A range of values within which the parameter can be set.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Parameter.AllowedValues")},
			{
				Name:        "change_type",
				Description: "The conditions under which changes to this parameter can be applied. Possible values are 'IMMEDIATE', 'REQUIRES_REBOOT'.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Parameter.ChangeType")},
			{
				Name:        "data_type",
				Description: "The data type of the parameter.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Parameter.DataType")},
			{
				Name:        "is_modifiable",
				Description: "Whether the customer is allowed to modify the parameter. Possible values are 'TRUE', 'FALSE' 'CONDITIONAL'.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Parameter.IsModifiable")},
			{
				Name:        "parameter_type",
				Description: "Determines whether the parameter can be applied to any node or only nodes of a particular type. Possible values are 'DEFAULT', 'NODE_TYPE_SPECIFIC'.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Parameter.ParameterType")},
			{
				Name:        "source",
				Description: "How the parameter is defined. For example, system denotes a system-defined parameter.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Parameter.Source")},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Parameter.ParameterName")},
		}),
	}
}
