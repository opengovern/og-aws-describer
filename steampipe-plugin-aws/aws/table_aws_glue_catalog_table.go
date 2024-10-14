package aws

import (
	"context"

	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsGlueCatalogTable(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_glue_catalog_table",
		Description: "AWS Glue Catalog Table",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"name", "database_name"}),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"EntityNotFoundException"}),
			},
			Hydrate: opengovernance.GetGlueCatalogTable,
		},
		List: &plugin.ListConfig{
			KeyColumns: []*plugin.KeyColumn{
				{Name: "catalog_id", Require: plugin.Optional},
				{Name: "database_name", Require: plugin.Optional},
			},
			Hydrate: opengovernance.ListGlueCatalogTable,
		},

		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The table name.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Table.Name")},
			{
				Name:        "catalog_id",
				Description: "The ID of the Data Catalog in which the table resides.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Table.CatalogId")},
			{
				Name:        "create_time",
				Description: "The time when the table definition was created in the data catalog.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Table.CreateTime")},
			{
				Name:        "description",
				Description: "A description of the table.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Table.Description")},
			{
				Name:        "created_by",
				Description: "The person or entity who created the table.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Table.CreatedBy")},
			{
				Name:        "database_name",
				Description: "The name of the database where the table metadata resides.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Table.DatabaseName")},
			{
				Name:        "is_registered_with_lake_formation",
				Description: "Indicates whether the table has been registered with lake formation.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Table.IsRegisteredWithLakeFormation")},
			{
				Name:        "last_access_time",
				Description: "The last time that the table was accessed. This is usually taken from HDFS, and might not be reliable.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Table.LastAccessTime")},
			{
				Name:        "last_analyzed_time",
				Description: "The last time that column statistics were computed for this table.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Table.LastAnalyzedTime")},
			{
				Name:        "owner",
				Description: "The owner of the table.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Table.Owner")},
			{
				Name:        "retention",
				Description: "The retention time for this table.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Table.Retention")},
			{
				Name:        "table_type",
				Description: "The type of this table (EXTERNAL_TABLE, VIRTUAL_VIEW, etc.).",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Table.TableType")},
			{
				Name:        "update_time",
				Description: "The last time that the table was updated.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Table.UpdateTime")},
			{
				Name:        "view_expanded_text",
				Description: "If the table is a view, the expanded text of the view otherwise null.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Table.ViewExpandedText")},
			{
				Name:        "view_original_text",
				Description: "If the table is a view, the original text of the view otherwise null.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Table.ViewOriginalText")},
			{
				Name:        "parameters",
				Description: "These key-value pairs define properties associated with the table.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Table.Parameters")},
			{
				Name:        "partition_keys",
				Description: "A list of columns by which the table is partitioned.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Table.PartitionKeys")},
			{
				Name:        "storage_descriptor",
				Description: "A storage descriptor containing information about the physical storage of this table.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Table.StorageDescriptor")},
			{
				Name:        "target_table",
				Description: "A TableIdentifier structure that describes a target table for resource linking.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Table.TargetTable")},
			{
				Name:        "lf_tags",
				Description: "LF-Tags assigned to the table by AWS Lake Formation.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.LfTags")},
			// Steampipe standard columns
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Table.Name")},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("ARN").Transform(arnToAkas)},
		}),
	}
}
