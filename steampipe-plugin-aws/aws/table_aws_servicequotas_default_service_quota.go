package aws

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsServiceQuotasDefaultServiceQuota(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_servicequotas_default_service_quota",
		Description: "AWS ServiceQuotas Default Service Quota",
		DefaultIgnoreConfig: &plugin.IgnoreConfig{
			ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"NoSuchResourceException"}),
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"service_code", "quota_code"}),
			//Hydrate:    opengovernance.GetServiceQuotasDefaultServiceQuota,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"NoSuchResourceException"}),
			},
		},
		List: &plugin.ListConfig{
			//Hydrate: opengovernance.ListServiceQuotasDefaultServiceQuota,
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
				Transform:   transform.FromField("Description.DefaultServiceQuota.QuotaName")},
			{
				Name:        "quota_code",
				Description: "The quota code.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DefaultServiceQuota.QuotaCode")},
			{
				Name:        "quota_arn",
				Description: "The arn of the service quota.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DefaultServiceQuota.QuotaArn")},
			{
				Name:        "global_quota",
				Description: "Indicates whether the quota is global.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.DefaultServiceQuota.GlobalQuota")},
			{
				Name:        "service_name",
				Description: "The service name.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DefaultServiceQuota.ServiceName")},
			{
				Name:        "service_code",
				Description: "The service identifier.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DefaultServiceQuota.ServiceCode")},
			{
				Name:        "adjustable",
				Description: "Indicates whether the quota value can be increased.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.DefaultServiceQuota.Adjustable")},
			{
				Name:        "unit",
				Description: "The unit of measurement.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DefaultServiceQuota.Unit")},
			{
				Name:        "value",
				Description: "The quota value.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.DefaultServiceQuota.Value")},
			{
				Name:        "error_reason",
				Description: "The error code and error reason.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DefaultServiceQuota.ErrorReason")},
			{
				Name:        "period",
				Description: "The period of time for the quota.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DefaultServiceQuota.Period")},
			{
				Name:        "usage_metric",
				Description: "Information about the measurement.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DefaultServiceQuota.UsageMetric")},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DefaultServiceQuota.QuotaName")},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DefaultServiceQuota.QuotaArn").Transform(transform.EnsureStringArray),
			},
		}),
	}
}
