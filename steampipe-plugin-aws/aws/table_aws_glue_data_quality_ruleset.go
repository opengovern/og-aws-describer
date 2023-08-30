package aws

import (
	"context"
	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsGlueDataQualityRuleset(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_glue_data_quality_ruleset",
		Description: "AWS Glue Data Quality Ruleset",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("name"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"EntityNotFoundException", "UnknownOperationException"}),
			},
			Hydrate: kaytu.GetGlueDataQualityRuleset,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListGlueDataQualityRuleset,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"UnknownOperationException"}),
			},
			KeyColumns: []*plugin.KeyColumn{
				// We need to pass both database name and table name all together.
				// If the database name and table name are passed as input parameters, the API gives a validation error, thus they are eliminated from the optional quals.
				// api error ValidationException: 1 validation error detected: Value null at 'filter.targetTable.name' failed to satisfy constraint: Member must not be null
				{Name: "created_on", Require: plugin.Optional, Operators: []string{"<=", "<", ">=", ">"}},
				{Name: "last_modified_on", Require: plugin.Optional},
			},
		},

		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the data quality ruleset.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DataQualityRuleset.Name")},
			{
				Name:        "database_name",
				Description: "The name of the database where the glue table exists.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DataQualityRuleset.TargetTable.DatabaseName")},
			{
				Name:        "table_name",
				Description: "The name of the glue table.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DataQualityRuleset.TargetTable.TableName")},
			{
				Name:        "created_on",
				Description: "The date and time the data quality ruleset was created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.DataQualityRuleset.CreatedOn")},
			{
				Name:        "description",
				Description: "A description of the data quality ruleset.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DataQualityRuleset.Description")},
			{
				Name:        "last_modified_on",
				Description: "The date and time the data quality ruleset was last modified.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.DataQualityRuleset.LastModifiedOn")},
			{
				Name:        "recommendation_run_id",
				Description: "When a ruleset was created from a recommendation run, this run ID is generated to link the two together.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DataQualityRuleset.RecommendationRunId")},
			{
				Name:        "rule_count",
				Description: "The number of rules in the ruleset.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.DataQualityRuleset")},
			//error
			//we don't have this field in the struct
			{
				Name:        "rule_set",
				Description: "A Data Quality Definition Language (DQDL) ruleset.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DataQualityRuleset.Ruleset")},
			{
				Name:        "target_table",
				Description: "An object representing a glue table.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DataQualityRuleset.TargetTable")},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DataQualityRuleset.Name")},
		}),
	}
}
