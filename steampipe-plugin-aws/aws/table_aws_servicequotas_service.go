package aws

import (
	"context"
	"fmt"
	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"

	"github.com/aws/aws-sdk-go-v2/service/servicequotas/types"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAwsServiceQuotasService(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_servicequotas_service",
		Description: "AWS Service Quotas Service",
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListServiceQuotasService,
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "service_name",
				Description: "Specifies the service name.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Service.ServiceName"),
			},
			{
				Name:        "service_code",
				Description: "Specifies the service identifier.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Service.ServiceCode"),
			},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Service.ServiceName"),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Hydrate:     getServiceQuotaServiceAkas,
				Transform:   transform.FromField("ARN").Transform(arnToAkas),
			},
		}),
	}
}

func getServiceQuotaServiceAkas(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	region := d.EqualsQualString(matrixKeyRegion)
	data := h.Item.(types.ServiceInfo)

	// Get common columns

	c, err := getCommonColumns(ctx, d, h)
	if err != nil {
		plugin.Logger(ctx).Error("aws_servicequotas_service.getServiceQuotaServiceAkas", "common_data_error", err)
		return nil, err
	}
	commonColumnData := c.(*awsCommonColumnData)
	arn := fmt.Sprintf("arn:%s:servicequotas:%s:%s:%s", commonColumnData.Partition, region, commonColumnData.AccountId, *data.ServiceCode)

	return []string{arn}, nil
}
