package aws

import (
	"context"
	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"

	healthv1 "github.com/aws/aws-sdk-go/service/health"
)

func tableAwsHealthAffectedEntity(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_health_affected_entity",
		Description: "AWS Health Affected Entity",
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListHealthAffectedEntity,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"SubscriptionRequiredException"}),
			},
			KeyColumns: []*plugin.KeyColumn{
				{Name: "arn", Require: plugin.Optional},
				{Name: "event_arn", Require: plugin.Optional},
				{Name: "entity_value", Require: plugin.Optional},
				{Name: "status_code", Require: plugin.Optional},
				{Name: "last_updated_time", Require: plugin.Optional, Operators: []string{">", ">=", "<", "<=", "="}},
			},
		},
		GetMatrixItemFunc: SupportedRegionMatrix(healthv1.EndpointsID),
		Columns: awsGlobalRegionColumns([]*plugin.Column{
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the health entity.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Entity.EntityArn"),
			},
			{
				Name:        "entity_url",
				Description: "The URL of the affected entity.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Entity.EntityUrl"),
			},
			{
				Name:        "entity_value",
				Description: "The ID of the affected entity.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Entity.EntityValue"),
			},
			{
				Name:        "event_arn",
				Description: "The Amazon Resource Name (ARN) of the health event.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Entity.EntityArn"),
			},
			{
				Name:        "last_updated_time",
				Description: "The most recent time that the entity was updated.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Entity.LastUpdatedTime"),
			},
			{
				Name:        "status_code",
				Description: "The most recent status of the entity affected by the event. The possible values are IMPAIRED, UNIMPAIRED, and UNKNOWN.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Entity.StatusCode"),
			},

			// Steampipe standard columns
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Entity.Tags"),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(getHealthAffectedEntityAkas),
			},
		}),
	}
}

func getHealthAffectedEntityAkas(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	clusterTags := d.HydrateItem.(kaytu.HealthAffectedEntity)
	if clusterTags.Description.Entity.EntityArn != nil {
		return []string{*clusterTags.Description.Entity.EntityArn}, nil
	}
	return nil, nil
}
