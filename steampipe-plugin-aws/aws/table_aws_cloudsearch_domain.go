package aws

import (
	"context"

	"github.com/kaytu-io/og-aws-describer/pkg/opengovernance-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsCloudSearchDomain(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_cloudsearch_domain",
		Description: "AWS CloudSearch Domain",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("domain_name"),
			Hydrate:    opengovernance.GetCloudSearchDomain,
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListCloudSearchDomain,
		},

		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "domain_name",
				Description: "A string that represents the name of a domain.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DomainStatus.DomainName")},
			{
				Name:        "domain_id",
				Description: "An internally generated unique identifier for a domain.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DomainStatus.DomainId")},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the search domain.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("ARN").Transform(arnToAkas),
			},
			{
				Name:        "created",
				Description: "True if the search domain is created.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.DomainStatus.Created")},
			{
				Name:        "deleted",
				Description: "True if the search domain has been deleted.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.DomainStatus.Deleted")},
			{
				Name:        "processing",
				Description: "True if processing is being done to activate the current domain configuration.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.DomainStatus.Processing")},
			{
				Name:        "requires_index_documents",
				Description: "True if Index Documents need to be called to activate the current domain configuration.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.DomainStatus.RequiresIndexDocuments")},
			{
				Name:        "search_instance_count",
				Description: "The number of search instances that are available to process search requests.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.DomainStatus.SearchInstanceCount")},
			{
				Name:        "search_instance_type",
				Description: "The instance type that is being used to process search requests.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DomainStatus.SearchInstanceType")},
			{
				Name:        "search_partition_count",
				Description: "The number of partitions across which the search index is spread.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.DomainStatus.SearchPartitionCount")},
			{
				Name:        "doc_service",
				Description: "The service endpoint for updating documents in a search domain.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DomainStatus.DocService")},
			{
				Name:        "limits",
				Description: "Limit details for a search domain.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DomainStatus.Limits")},
			{
				Name:        "search_service",
				Description: "The service endpoint for requesting search results from a search domain.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DomainStatus.SearchService")},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("ARN").Transform(arnToAkas),
			},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DomainStatus.DomainName")},
		}),
	}
}
