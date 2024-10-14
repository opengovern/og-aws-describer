package aws

import (
	"context"
	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsAPIGatewayStage(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_api_gateway_stage",
		Description: "AWS API Gateway Stage",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"rest_api_id", "name"}),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"NotFoundException"}),
			},
			Hydrate: opengovernance.GetApiGatewayStage,
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListApiGatewayStage,
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the stage.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Stage.StageName"),
			},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the  stage.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("ARN"),
			},
			{
				Name:        "rest_api_id",
				Description: "The id of the rest api which contains this stage.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.RestApiId"),
			},
			{
				Name:        "deployment_id",
				Description: "The identifier of the Deployment that the stage points to.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Stage.DeploymentId"),
			},
			{
				Name:        "created_date",
				Description: "The timestamp when the stage was created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Stage.CreatedDate"),
			},
			{
				Name:        "cache_cluster_enabled",
				Description: "Specifies whether a cache cluster is enabled for the stage.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Stage.CacheClusterEnabled"),
			},
			{
				Name:        "tracing_enabled",
				Description: "Specifies whether active tracing with X-ray is enabled for the Stage.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Stage.TracingEnabled"),
			},
			{
				Name:        "access_log_settings",
				Description: "Settings for logging access in this stage.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Stage.AccessLogSettings"),
			},
			{
				Name:        "cache_cluster_size",
				Description: "The size of the cache cluster for the stage, if enabled.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Stage.CacheClusterSize"),
			},
			{
				Name:        "cache_cluster_status",
				Description: "The status of the cache cluster for the stage, if enabled.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Stage.CacheClusterStatus"),
			},
			{
				Name:        "client_certificate_id",
				Description: "The identifier of a client certificate for an API stage.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Stage.ClientCertificateId"),
			},
			{
				Name:        "description",
				Description: "The stage's description.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Stage.Description"),
			},
			{
				Name:        "documentation_version",
				Description: "The version of the associated API documentation.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Stage.DocumentationVersion"),
			},
			{
				Name:        "last_updated_date",
				Description: "The timestamp when the stage last updated.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Stage.LastUpdatedDate"),
			},
			{
				Name:        "canary_settings",
				Description: "A map of settings for the canary deployment in this stage.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Stage.CanarySettings"),
			},
			{
				Name:        "method_settings",
				Description: "A map that defines the method settings for a Stage resource.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Stage.MethodSettings"),
			},
			{
				Name:        "variables",
				Description: "A map that defines the stage variables for a Stage resource.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Stage.Variables"),
			},
			{
				Name:        "web_acl_arn",
				Description: "The ARN of the WebAcl associated with the Stage.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Stage.WebAclArn"),
			},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Stage.StageName"),
			},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Stage.Tags"),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("ARN").Transform(transform.EnsureStringArray),
			},
		}),
	}
}
