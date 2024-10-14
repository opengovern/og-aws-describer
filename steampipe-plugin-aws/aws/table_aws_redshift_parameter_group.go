package aws

import (
	"context"
	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"
	"strings"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAwsRedshiftParameterGroup(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_redshift_parameter_group",
		Description: "AWS Redshift Parameter Group",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("name"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"ClusterParameterGroupNotFound"}),
			},
			Hydrate: opengovernance.GetRedshiftClusterParameterGroup,
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListRedshiftClusterParameterGroup,
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the cluster parameter group.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.ClusterParameterGroup.ParameterGroupName")},
			{
				Name:        "description",
				Description: "The description of the parameter group.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ClusterParameterGroup.Description")},
			{
				Name:        "family",
				Description: "The name of the cluster parameter group family that this cluster parameter group is compatible with.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ClusterParameterGroup.ParameterGroupFamily")},
			{
				Name:        "parameters",
				Description: "A list of Parameter instances. Each instance lists the parameters of one cluster parameter group.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Parameters")},
			{
				Name:        "tags_src",
				Description: "A list of tags assigned to the parameter group.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ClusterParameterGroup.Tags")},
			// Standard columns for all tables

			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ClusterParameterGroup.ParameterGroupName")},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ClusterParameterGroup.Tags").Transform(tagListToTurbotTags),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(getAwsRedshiftParameterGroupAkas),
			},
		}),
	}
}

//// LIST FUNCTION

//// HYDRATE FUNCTIONS

//// TRANSFORM FUNCTIONS

func getAwsRedshiftParameterGroupAkas(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	parameterData := d.HydrateItem.(opengovernance.RedshiftClusterParameterGroup).Description.ClusterParameterGroup
	metadata := d.HydrateItem.(opengovernance.RedshiftClusterParameterGroup).Metadata

	aka := "arn:" + metadata.Partition + ":redshift:" + metadata.Region + ":" + metadata.AccountID + ":parametergroup"

	if strings.HasPrefix(*parameterData.ParameterGroupName, ":") {
		aka = aka + *parameterData.ParameterGroupName
	} else {
		aka = aka + ":" + *parameterData.ParameterGroupName
	}

	return []string{aka}, nil
}

func tagListToTurbotTags(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	plugin.Logger(ctx).Trace("tagListToTurbotTags")

	tagList := d.HydrateItem.(opengovernance.RedshiftClusterParameterGroup).Description.ClusterParameterGroup

	// Mapping the resource tags inside turbotTags
	var turbotTagsMap map[string]string
	turbotTagsMap = map[string]string{}
	for _, i := range tagList.Tags {
		turbotTagsMap[*i.Key] = *i.Value
	}

	return turbotTagsMap, nil
}
