package aws

import (
	"context"

	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsRDSDBClusterParameterGroup(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_rds_db_cluster_parameter_group",
		Description: "AWS RDS DB Cluster Parameter Group",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("name"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"DBParameterGroupNotFound"}),
			},
			Hydrate: opengovernance.GetRDSDBClusterParameterGroup,
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListRDSDBClusterParameterGroup,
		},

		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The friendly name to identify the DB cluster parameter group.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DBClusterParameterGroup.DBClusterParameterGroupName")},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) for the DB cluster parameter group.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DBClusterParameterGroup.DBClusterParameterGroupArn"),
			},
			{
				Name:        "description",
				Description: "Provides the customer-specified description for this DB cluster parameter group.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DBClusterParameterGroup.Description")},
			{
				Name:        "db_parameter_group_family",
				Description: "The name of the DB parameter group family that this DB cluster parameter group is compatible with.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DBClusterParameterGroup.DBParameterGroupFamily")},
			{
				Name:        "parameters",
				Description: "A list of detailed parameter for a particular DB Cluster parameter group.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Parameters")},
			{
				Name:        "tags_src",
				Description: "A list of tags attached to the DB Cluster parameter group.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Tags")},

			// Standard columns

			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(getRDSDBClusterParameterGroupTurbotTags),
			},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DBClusterParameterGroup.DBClusterParameterGroupName")},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DBClusterParameterGroup.DBClusterParameterGroupArn").Transform(arnToAkas),
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

// List call

// Create service

//// TRANSFORM FUNCTIONS ////

func getRDSDBClusterParameterGroupTurbotTags(_ context.Context, d *transform.TransformData) (interface{}, error) {
	dbClusterParameterGroup := d.HydrateItem.(opengovernance.RDSDBClusterParameterGroup).Description

	if dbClusterParameterGroup.Tags != nil {
		turbotTagsMap := map[string]string{}
		for _, i := range dbClusterParameterGroup.Tags {
			turbotTagsMap[*i.Key] = *i.Value
		}
		return turbotTagsMap, nil
	}
	return nil, nil
}
