package aws

import (
	"context"

	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsGlueCrawler(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_glue_crawler",
		Description: "AWS Glue Crawler",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("name"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"EntityNotFoundException", "InvalidParameter"}),
			},
			Hydrate: opengovernance.GetGlueCrawler,
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListGlueCrawler,
		},

		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the crawler.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Crawler.Name")},
			{
				Name:        "arn",
				Description: "The ARN of the crawler.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("ARN").Transform(arnToAkas),
			},
			{
				Name:        "database_name",
				Description: "The name of the database in which the crawler's output is stored.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Crawler.DatabaseName")},
			{
				Name:        "state",
				Description: "Indicates whether the crawler is running or pending.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Crawler.State")},
			{
				Name:        "role",
				Description: "The Amazon Resource Name (ARN) of an IAM role that's used to access customer resources, such as Amazon Simple Storage Service (Amazon S3) data.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Crawler.Role")},
			{
				Name:        "creation_time",
				Description: "The time that the crawler was created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Crawler.CreationTime")},
			{
				Name:        "description",
				Description: "A description of the crawler.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Crawler.Description")},
			{
				Name:        "crawl_elapsed_time",
				Description: "If the crawler is running, contains the total time elapsed since the last crawl began.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Crawler.CrawlElapsedTime")},
			{
				Name:        "crawler_lineage_settings",
				Description: "Specifies whether data lineage is enabled for the crawler.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Crawler.LineageConfiguration.CrawlerLineageSettings")},
			{
				Name:        "crawler_security_configuration",
				Description: "The name of the SecurityConfiguration structure to be used by this crawler.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Crawler.CrawlerSecurityConfiguration")},
			{
				Name:        "last_updated",
				Description: "The time that the crawler was last updated.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Crawler.LastUpdated")},
			{
				Name:        "recrawl_behavior",
				Description: "Specifies whether to crawl the entire dataset again or to crawl only folders that were added since the last crawler run. A value of CRAWL_EVERYTHING specifies crawling the entire dataset again. A value of CRAWL_NEW_FOLDERS_ONLY specifies crawling only folders that were added since the last crawler run. A value of CRAWL_EVENT_MODE specifies crawling only the changes identified by Amazon S3 events.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Crawler.RecrawlPolicy.RecrawlBehavior")},
			{
				Name:        "table_prefix",
				Description: "The prefix added to the names of tables that are created.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Crawler.TablePrefix")},
			{
				Name:        "version",
				Description: "The version of the crawler.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Crawler.Version")},
			{
				Name:        "classifiers",
				Description: "A list of UTF-8 strings that specify the custom classifiers that are associated with the crawler.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Crawler.Classifiers")},
			{
				Name:        "configuration",
				Description: "Crawler configuration information.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Crawler.Configuration")},
			{
				Name:        "last_crawl",
				Description: "The status of the last crawl, and potentially error information if an error occurred.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Crawler.LastCrawl")},
			{
				Name:        "schedule",
				Description: "For scheduled crawlers, the schedule when the crawler runs.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Crawler.Schedule")},
			{
				Name:        "schema_change_policy",
				Description: "The policy that specifies update and delete behaviors for the crawler.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Crawler.SchemaChangePolicy")},
			{
				Name:        "targets",
				Description: "A collection of targets to crawl.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Crawler.Targets")},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Crawler.Name")},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("ARN").Transform(transform.EnsureStringArray),
			},
		}),
	}
}
