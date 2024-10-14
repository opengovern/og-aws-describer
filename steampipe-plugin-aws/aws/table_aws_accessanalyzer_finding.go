package aws

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/accessanalyzer/types"
	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

type accessanalyzerFindingInfo = struct {
	Finding           types.Finding
	AccessAnalyzerArn string
}

func tableAwsAccessAnalyzerFinding(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_accessanalyzer_finding",
		Description: "AWS Access Analyzer Finding",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"id", "access_analyzer_arn"}),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"ResourceNotFoundException", "ValidationException"}),
			},
			Hydrate: opengovernance.GetAccessAnalyzerAnalyzerFinding,
			Tags:    map[string]string{"service": "access-analyzer", "action": "GetFinding"},
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListAccessAnalyzerAnalyzerFinding,
			Tags:    map[string]string{"service": "access-analyzer", "action": "ListFindings"},
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"ResourceNotFoundException", "ValidationException"}),
			},
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "access_analyzer_arn",
				Description: "The Amazon Resource Name (ARN) of the analyzer that generated the finding.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AnalyzerArn"),
			},
			{
				Name:        "id",
				Description: "The ID of the finding.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Finding.Id"),
			},
			{
				Name:        "analyzed_at",
				Description: "The time at which the resource-based policy that generated the finding was analyzed.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Finding.AnalyzedAt"),
			},
			{
				Name:        "created_at",
				Description: "The time at which the finding was created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Finding.CreatedAt"),
			},
			{
				Name:        "error",
				Description: "The error that resulted in an Error finding.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Finding.Error"),
			},
			{
				Name:        "is_public",
				Description: "Indicates whether the finding reports a resource that has a policy that allows public access.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Finding.IsPublic"),
			},
			{
				Name:        "resource",
				Description: "The resource that the external principal has access to.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Finding.Resource"),
			},
			{
				Name:        "resource_owner_account",
				Description: "The Amazon Web Services account ID that owns the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Finding.ResourceOwnerAccount"),
			},
			{
				Name:        "resource_type",
				Description: "The type of the resource that the external principal has access to.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Finding.ResourceType"),
			},
			{
				Name:        "status",
				Description: "The status of the finding.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Finding.Status"),
			},
			{
				Name:        "updated_at",
				Description: "The time at which the finding was most recently updated.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Finding.UpdatedAt"),
			},
			{
				Name:        "action",
				Description: "The action in the analyzed policy statement that an external principal has permission to use.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Finding.Action"),
			},
			{
				Name:        "sources",
				Description: "The sources of the finding, indicating how the access that generated the finding is granted. It is populated for Amazon S3 bucket findings.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Finding.Sources"),
			},
			{
				Name:        "principal",
				Description: "The external principal that has access to a resource within the zone of trust.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Finding.Principal"),
			},
			{
				Name:        "condition",
				Description: "The condition in the analyzed policy statement that resulted in a finding.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Finding.Condition"),
			},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Finding.Id"),
			},
		}),
	}
}
