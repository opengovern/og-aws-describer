package aws

import (
	"context"

	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsAPIGatewayUsagePlan(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_api_gateway_usage_plan",
		Description: "AWS API Gateway Usage Plan",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"NotFoundException"}),
			},
			Hydrate: kaytu.GetApiGatewayUsagePlan,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListApiGatewayUsagePlan,
		},

		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The name of a usage plan",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.UsagePlan.Name")},
			{
				Name:        "id",
				Description: "The identifier of a UsagePlan resource",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.UsagePlan.Id")},
			{
				Name:        "product_code",
				Description: "The AWS Markeplace product identifier to associate with the usage plan as a SaaS product on AWS Marketplace",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.UsagePlan.ProductCode")},
			{
				Name:        "description",
				Description: "The description of a usage plan",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.UsagePlan.Description")},
			{
				Name:        "quota",
				Description: "The maximum number of permitted requests per a given unit time interval",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.UsagePlan.Quota")},
			{
				Name:        "throttle",
				Description: "The request throttle limits of a usage plan",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.UsagePlan.Throttle")},
			{
				Name:        "api_stages",
				Description: "The associated API stages of a usage plan",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.UsagePlan.ApiStages")},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.UsagePlan.Tags")},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.UsagePlan.Name")},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("ARN").Transform(arnToAkas)},
		}),
	}
}
