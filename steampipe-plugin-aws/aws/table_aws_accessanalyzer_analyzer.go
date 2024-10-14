package aws

import (
	"context"
	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsAccessAnalyzer(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_accessanalyzer_analyzer",
		Description: "AWS Access Analyzer",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("name"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"ResourceNotFoundException", "ValidationException", "InvalidParameter"}),
			},
			Hydrate: opengovernance.GetAccessAnalyzerAnalyzer,
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListAccessAnalyzerAnalyzer,
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:    "type",
					Require: plugin.Optional,
				},
			},
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the Analyzer.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Analyzer.Name"),
			},
			{
				Name:        "arn",
				Description: "The ARN of the analyzer.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Analyzer.Arn"),
			},
			{
				Name:        "status",
				Description: "The status of the analyzer.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Analyzer.Status"),
			},
			{
				Name:        "type",
				Description: "The type of analyzer, which corresponds to the zone of trust chosen for the analyzer.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Analyzer.Type"),
			},
			{
				Name:        "created_at",
				Description: "A timestamp for the time at which the analyzer was created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Analyzer.CreatedAt"),
			},
			{
				Name:        "last_resource_analyzed",
				Description: "The resource that was most recently analyzed by the analyzer.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Analyzer.LastResourceAnalyzed"),
			},
			{
				Name:        "last_resource_analyzed_at",
				Description: "The time at which the most recently analyzed resource was analyzed.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Analyzer.LastResourceAnalyzedAt"),
			},
			{
				Name:        "status_reason",
				Description: "The statusReason provides more details about the current status of the analyzer.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Analyzer.StatusReason"),
			},
			{
				Name:        "findings",
				Description: "A list of findings retrieved from the analyzer that match the filter criteria specified, if any.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Findings"),
			},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Analyzer.Name"),
			},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Analyzer.Tags"),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Analyzer.Arn").Transform(arnToAkas),
			},
		}),
	}
}

//// LIST FUNCTION

//// HYDRATE FUNCTIONS
