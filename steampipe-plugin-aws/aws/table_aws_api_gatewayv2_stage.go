package aws

import (
	"context"
	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsAPIGatewayV2Stage(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_api_gatewayv2_stage",
		Description: "AWS API Gateway Version 2 Stage",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"api_id", "stage_name"}),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"NotFoundException"}),
			},
			Hydrate: kaytu.GetApiGatewayV2Stage,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListApiGatewayV2Stage,
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "stage_name",
				Description: "The name of the stage",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Stage.StageName")},
			{
				Name:        "api_id",
				Description: "The id of the api which contains this stage",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ApiId")},
			{
				Name:        "api_gateway_managed",
				Description: "Specifies whether a stage is managed by API Gateway",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Stage.ApiGatewayManaged")},
			{
				Name:        "auto_deploy",
				Description: "Specifies whether updates to an API automatically trigger a new deployment",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Stage.AutoDeploy")},
			{
				Name:        "client_certificate_id",
				Description: "The identifier of a client certificate for a Stage. Supported only for WebSocket APIs",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Stage.ClientCertificateId")},
			{
				Name:        "created_date",
				Description: "The timestamp when the stage was created",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Stage.CreatedDate")},
			{
				Name:        "deployment_id",
				Description: "The identifier of the Deployment that the Stage is associated with",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Stage.DeploymentId")},
			{
				Name:        "default_route_data_trace_enabled",
				Description: "Specifies whether (true) or not (false) data trace logging is enabled for this route. This property affects the log entries pushed to Amazon CloudWatch Logs. Supported only for WebSocket APIs",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Stage.DefaultRouteSettings.DataTraceEnabled")},
			{
				Name:        "default_route_detailed_metrics_enabled",
				Description: "Specifies whether detailed metrics are enabled",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Stage.DefaultRouteSettings.DetailedMetricsEnabled")},
			{
				Name:        "default_route_logging_level",
				Description: "Specifies the logging level for this route: INFO, ERROR, or OFF. This property affects the log entries pushed to Amazon CloudWatch Logs. Supported only for WebSocket APIs",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Stage.DefaultRouteSettings.LoggingLevel")},
			{
				Name:        "default_route_throttling_burst_limit",
				Description: "Throttling burst limit for default route settings",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Stage.DefaultRouteSettings.ThrottlingBurstLimit")},
			{
				Name:        "default_route_throttling_rate_limit",
				Description: "Throttling rate limit for default route settings",
				Type:        proto.ColumnType_DOUBLE,
				Transform:   transform.FromField("Description.Stage.DefaultRouteSettings.ThrottlingRateLimit")},
			{
				Name:        "last_deployment_status_message",
				Description: "Describes the status of the last deployment of a stage. Supported only for stages with autoDeploy enabled",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Stage.LastDeploymentStatusMessage")},
			{
				Name:        "last_updated_date",
				Description: "The timestamp when the stage was last updated",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Stage.LastUpdatedDate")},
			{
				Name:        "description",
				Description: "The stage's description",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Stage.Description"),
			},
			{
				Name:        "access_log_settings",
				Description: "Access log settings of the stage.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Stage.AccessLogSettings"),
			},
			{
				Name:        "stage_variables",
				Description: "A map that defines the stage variables for a stage resource",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Stage.StageVariables")},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Stage.Tags")},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Stage.StageName")},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(apiGatewayV2StageAkas),
			},
		}),
	}
}

//// LIST FUNCTION

//// HYDRATE FUNCTIONS

func apiGatewayV2StageAkas(_ context.Context, d *transform.TransformData) (interface{}, error) {
	data := d.HydrateItem.(kaytu.ApiGatewayV2Stage).Description
	metadata := d.HydrateItem.(kaytu.ApiGatewayV2Stage).Metadata
	akas := []string{"arn:" + metadata.Partition + ":apigateway:" + metadata.Region + "::/apis/" + *data.ApiId + "/stages/" + *data.Stage.StageName}

	return akas, nil
}
