package aws

import (
	"context"

	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsGlueCatalogDatabase(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_glue_catalog_database",
		Description: "AWS Glue Catalog Database",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("name"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"EntityNotFoundException"}),
			},
			Hydrate: kaytu.GetGlueCatalogDatabase,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListGlueCatalogDatabase,
		},

		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the database. For Hive compatibility, this is folded to lowercase when it is stored.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Database.Name")},
			{
				Name:        "catalog_id",
				Description: "The ID of the Data Catalog in which the database resides.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Database.CatalogId")},
			{
				Name:        "create_time",
				Description: "The time at which the metadata database was created in the catalog.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Database.CreateTime")},
			{
				Name:        "description",
				Description: "A description of the database.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Database.Description")},
			{
				Name:        "location_uri",
				Description: "The location of the database (for example, an HDFS path).",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Database.LocationUri")},
			{
				Name:        "create_table_default_permissions",
				Description: "Creates a set of default permissions on the table for principals.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Database.CreateTableDefaultPermissions")},
			{
				Name:        "parameters",
				Description: "These key-value pairs define parameters and properties of the database.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Database.Parameters")},
			{
				Name:        "target_database",
				Description: "A DatabaseIdentifier structure that describes a target database for resource linking.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Database.TargetDatabase")},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Database.Name")},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("ARN").Transform(arnToAkas)},
		}),
	}
}
