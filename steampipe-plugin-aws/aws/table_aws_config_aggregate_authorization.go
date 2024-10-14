package aws

import (
	"context"

	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAwsConfigAggregateAuthorization(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_config_aggregate_authorization",
		Description: "AWS Config Aggregate Authorization",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListConfigAggregationAuthorization,
		},

		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the aggregation object.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AggregationAuthorization.AggregationAuthorizationArn"),
			},
			{
				Name:        "authorized_account_id",
				Description: "The 12-digit account ID of the account authorized to aggregate data.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AggregationAuthorization.AuthorizedAccountId")},
			{
				Name:        "authorized_aws_region",
				Description: "The region authorized to collect aggregated data.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AggregationAuthorization.AuthorizedAwsRegion")},
			{
				Name:        "creation_time",
				Description: "The time stamp when the aggregation authorization was created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.AggregationAuthorization.CreationTime")},
			{
				Name:        "tags_src",
				Description: "A list of tags attached to the Cluster.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Tags")},

			// Standard columns
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Tags").Transform(configAggregateAuthorizationsTagListToTurbotTags),
			},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AggregationAuthorization.AggregationAuthorizationArn")},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.AggregationAuthorization.AggregationAuthorizationArn").Transform(transform.EnsureStringArray),
			},
		}),
	}
}

func configAggregateAuthorizationsTagListToTurbotTags(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	tagList := d.HydrateItem.(opengovernance.ConfigAggregationAuthorization).Description.Tags

	// Mapping the resource tags inside turbotTags
	var turbotTagsMap map[string]string
	if len(tagList) > 0 {
		turbotTagsMap = map[string]string{}
		for _, i := range tagList {
			turbotTagsMap[*i.Key] = *i.Value
		}
	}
	return turbotTagsMap, nil
}
