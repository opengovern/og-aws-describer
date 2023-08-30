package aws

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/route53/types"

	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAwsRoute53HealthCheck(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_route53_health_check",
		Description: "AWS Route53 Health Check",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    kaytu.GetRoute53HealthCheck,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"NoSuchHealthCheck"}),
			},
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListRoute53HealthCheck,
		},
		Columns: awsKaytuColumns([]*plugin.Column{
			{
				Name:        "id",
				Description: "The identifier that Amazon Route 53 assigned to the health check.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.HealthCheck.Id")},
			{
				Name:        "caller_reference",
				Description: "A unique string that you specified when you created the health check.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.HealthCheck.CallerReference")},
			{
				Name:        "health_check_version",
				Description: "The version of the health check.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.HealthCheck.HealthCheckVersion")},
			{
				Name:        "linked_service_principal",
				Description: "If the health check was created by another service, the service that created the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.HealthCheck.LinkedService.ServicePrincipal")},
			{
				Name:        "linked_service_description",
				Description: "If the health check was created by another service, an configurationtional description that can be provided by the other service.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.HealthCheck.LinkedService.Description")},
			{
				Name:        "cloud_watch_alarm_configuration",
				Description: "A complex type that contains information about the CloudWatch alarm that Amazon Route 53 is monitoring for this health check.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.HealthCheck.CloudWatchAlarmConfiguration")},
			{
				Name:        "health_check_config",
				Description: "A complex type that contains detailed information about one health check.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.HealthCheck.HealthCheckConfig")},
			{
				Name:        "health_check_status",
				Description: "A list that contains one HealthCheckObservation element for each Amazon Route 53 health checker that is reporting a status about the health check endpoint.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Status.HealthCheckObservations"),
			},
			{
				Name:        "tags_src",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Tags.ResourceTagSet.Tags")},
			// Steampipe standard columns
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.HealthCheck.Id")},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Tags.ResourceTagSet.Tags").Transform(route53HealthCheckTurbotTags),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Hydrate:     getRoute53HealthCheckTurbotAkas,
				Transform:   transform.FromValue(),
			},
		}),
	}
}

func getRoute53HealthCheckTurbotAkas(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	healthCheck := h.Item.(kaytu.Route53HealthCheck)

	// Get data for turbot defined prconfigurationerties
	akas := []string{"arn:" + healthCheck.Metadata.Partition + ":route53:::" + "healthcheck/" + *healthCheck.Description.HealthCheck.Id}

	return akas, nil
}

//// TRANSFORM FUNCTIONS

func route53HealthCheckTurbotTags(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	tags := d.Value.([]types.Tag)

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
