package aws

import (
	"context"

	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAwsNeptuneDatabase(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_neptune_database",
		Description: "AWS Neptune Database",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("db_instance_identifier"),
			Hydrate:    opengovernance.GetNeptuneDatabase,
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListNeptuneDatabase,
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "db_instance_identifier",
				Description: "The id of the database.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Database.DBInstanceIdentifier")},
			{
				Name:        "db_name",
				Description: "The name of the database.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Database.DBName")},
			{
				Name:        "db_instance_arn",
				Description: "The Amazon Resource Name (ARN) of the database",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Database.DBInstanceArn")},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Database.DBName")},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(getNeptuneDatabaseTurbotTags),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Database.DBInstanceArn").Transform(arnToAkas),
			},
		}),
	}
}

//// TRANSFORM FUNCTIONS

func getNeptuneDatabaseTurbotTags(_ context.Context, d *transform.TransformData) (interface{}, error) {
	tags := d.HydrateItem.(opengovernance.NeptuneDatabase).Description.Tags
	return neptuneV2TagsToMap(tags)
}
