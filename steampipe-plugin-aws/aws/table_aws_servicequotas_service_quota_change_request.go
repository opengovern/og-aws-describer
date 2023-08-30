package aws

import (
	"context"

	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsServiceQuotasServiceQuotaChangeRequest(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_servicequotas_service_quota_change_request",
		Description: "AWS ServiceQuotas Service Quota Change Request",
		DefaultIgnoreConfig: &plugin.IgnoreConfig{
			ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"NoSuchResourceException"}),
		},
		Get: &plugin.GetConfig{
			Hydrate:    kaytu.GetServiceQuotasServiceQuotaChangeRequest,
			KeyColumns: plugin.SingleColumn("id"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"NoSuchResourceException"}),
			},
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListServiceQuotasServiceQuotaChangeRequest,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"NoSuchResourceException"}),
			},
			KeyColumns: []*plugin.KeyColumn{
				{Name: "service_code", Require: plugin.Optional},
				{Name: "status", Require: plugin.Optional},
			},
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "id",
				Description: "The unique identifier.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ServiceQuotaChangeRequest.Id")},
			{
				Name:        "case_id",
				Description: "The case ID.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ServiceQuotaChangeRequest.CaseId")},
			{
				Name:        "status",
				Description: "The state of the quota increase request.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ServiceQuotaChangeRequest.Status")},
			{
				Name:        "quota_name",
				Description: "The quota name.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ServiceQuotaChangeRequest.QuotaName")},
			{
				Name:        "quota_code",
				Description: "The quota code.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ServiceQuotaChangeRequest.QuotaCode")},
			{
				Name:        "quota_arn",
				Description: "The arn of the service quota.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ServiceQuotaChangeRequest.QuotaArn")},
			{
				Name:        "desired_value",
				Description: "The increased value for the quota.",
				Type:        proto.ColumnType_DOUBLE,
				Transform:   transform.FromField("Description.ServiceQuotaChangeRequest.DesiredValue")},
			{
				Name:        "created",
				Description: "The date and time when the quota increase request was received and the case ID was created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.ServiceQuotaChangeRequest.Created")},
			{
				Name:        "global_quota",
				Description: "Indicates whether the quota is global.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.ServiceQuotaChangeRequest.GlobalQuota")},
			{
				Name:        "last_updated",
				Description: "The date and time of the most recent change.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.ServiceQuotaChangeRequest.LastUpdated")},
			{
				Name:        "requester",
				Description: "The IAM identity of the requester.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ServiceQuotaChangeRequest.Requester")},
			{
				Name:        "service_name",
				Description: "The service name.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ServiceQuotaChangeRequest.ServiceName")},
			{
				Name:        "service_code",
				Description: "The service identifier.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ServiceQuotaChangeRequest.ServiceCode")},
			{
				Name:        "unit",
				Description: "The unit of measurement.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ServiceQuotaChangeRequest.Unit")},
			{
				Name:        "tags_src",
				Description: "The list of tags associated with the change request.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Tags")},

			// Steampipe standard columns
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(serviceQuotaChangeRequestTagsToTurbotTags),
			},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ServiceQuotaChangeRequest.QuotaName")},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("ARN").Transform(arnToAkas),
			},
		}),
	}
}

func serviceQuotaChangeRequestTagsToTurbotTags(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	tags := d.HydrateItem.(kaytu.ServiceQuotasServiceQuotaChangeRequest).Description.Tags

	// Mapping the resource tags inside turbotTags
	var turbotTagsMap map[string]string
	if len(tags) > 0 {
		turbotTagsMap = map[string]string{}
		for _, i := range tags {
			turbotTagsMap[*i.Key] = *i.Value
		}
	}

	return turbotTagsMap, nil
}
