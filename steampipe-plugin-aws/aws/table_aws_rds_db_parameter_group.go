package aws

import (
	"context"

	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsRDSDBParameterGroup(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_rds_db_parameter_group",
		Description: "AWS RDS DB Parameter Group",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("name"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"DBParameterGroupNotFound"}),
			},
			Hydrate: opengovernance.GetRDSDBParameterGroup,
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListRDSDBParameterGroup,
		},

		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The friendly name to identify the DB parameter group.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DBParameterGroup.DBParameterGroupName")},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) for the DB parameter group.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("ARN"),
			},
			{
				Name:        "description",
				Description: "Provides the customer-specified description for this DB parameter group.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DBParameterGroup.Description")},
			{
				Name:        "db_parameter_group_family",
				Description: "The name of the DB parameter group family that this DB parameter group is compatible with.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DBParameterGroup.DBParameterGroupFamily")},
			{
				Name:        "parameters",
				Description: "A list of detailed parameter for a particular DB parameter group.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Parameters")},
			{
				Name:        "tags_src",
				Description: "A list of tags attached to the DB parameter group.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Tags")},

			// Standard columns
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(getRDSDBParameterGroupTurbotTags),
			},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DBParameterGroup.DBParameterGroupName")},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DBParameterGroup.DBParameterGroupArn").Transform(arnToAkas),
			},
		}),
	}
}

// Create Session

// Limiting the results

// List call

// Context can be cancelled due to manual cancellation or the limit has been hit

//// HYDRATE FUNCTIONS

// Create service

// Create service

// Create service

//// TRANSFORM FUNCTIONS

func getRDSDBParameterGroupTurbotTags(_ context.Context, d *transform.TransformData) (interface{}, error) {
	dbParameterGroup := d.HydrateItem.(opengovernance.RDSDBParameterGroup).Description

	if dbParameterGroup.Tags != nil {
		turbotTagsMap := map[string]string{}
		for _, i := range dbParameterGroup.Tags {
			turbotTagsMap[*i.Key] = *i.Value
		}
		return turbotTagsMap, nil
	}
	return nil, nil
}
