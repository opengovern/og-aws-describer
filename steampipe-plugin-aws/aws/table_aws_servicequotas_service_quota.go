package aws

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsServiceQuotasServiceQuota(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_servicequotas_service_quota",
		Description: "AWS ServiceQuotas Service Quota",
		DefaultIgnoreConfig: &plugin.IgnoreConfig{
			ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"NoSuchResourceException"}),
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"service_code", "quota_code"}),
			//Hydrate:    opengovernance.GetServiceQuotasServiceQuota,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"NoSuchResourceException"}),
			},
		},
		List: &plugin.ListConfig{
			//Hydrate: opengovernance.ListServiceQuotasServiceQuota,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"NoSuchResourceException"}),
			},
			KeyColumns: []*plugin.KeyColumn{
				{Name: "service_code", Require: plugin.Optional},
			},
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "quota_name",
				Description: "The quota name.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ServiceQuota.QuotaName")},
			{
				Name:        "quota_code",
				Description: "The quota code.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ServiceQuota.QuotaCode")},
			{
				Name:        "quota_arn",
				Description: "The arn of the service quota.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ServiceQuota.QuotaArn")},
			{
				Name:        "global_quota",
				Description: "Indicates whether the quota is global.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.ServiceQuota.GlobalQuota")},
			{
				Name:        "service_name",
				Description: "The service name.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ServiceQuota.ServiceName")},
			{
				Name:        "service_code",
				Description: "The service identifier.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ServiceQuota.ServiceCode")},
			{
				Name:        "adjustable",
				Description: "Indicates whether the quota value can be increased.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.ServiceQuota.Adjustable")},
			{
				Name:        "unit",
				Description: "The unit of measurement.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ServiceQuota.Unit")},
			{
				Name:        "value",
				Description: "The quota value.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.ServiceQuota.Value")},
			{
				Name:        "error_reason",
				Description: "The error code and error reason.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ServiceQuota.ErrorReason")},
			{
				Name:        "period",
				Description: "The period of time for the quota.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ServiceQuota.Period")},
			{
				Name:        "tags_src",
				Description: "The list of tags associated with the service quota.",
				Type:        proto.ColumnType_JSON,

				Transform: transform.FromField("Description.ServiceQuota.Adjustable")},
			{
				Name:        "usage_metric",
				Description: "Information about the measurement.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ServiceQuota.UsageMetric")},

			// Steampipe standard columns
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				//Transform:   transform.From(serviceQuotaTagsToTurbotTags),
			},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ServiceQuota.QuotaName")},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ServiceQuota.QuotaArn").Transform(transform.EnsureStringArray),
			},
		}),
	}
}

//// TRANSFORM FUNCTIONS

//func serviceQuotaTagsToTurbotTags(ctx context.Context, d *transform.TransformData) (interface{}, error) {
//	tags := d.HydrateItem.(opengovernance.ServiceQuotasServiceQuota).Description.Tags
//
//	// Mapping the resource tags inside turbotTags
//	var turbotTagsMap map[string]string
//	if len(tags) > 0 {
//		turbotTagsMap = map[string]string{}
//		for _, i := range tags {
//			turbotTagsMap[*i.Key] = *i.Value
//		}
//	}
//
//	return turbotTagsMap, nil
//}
