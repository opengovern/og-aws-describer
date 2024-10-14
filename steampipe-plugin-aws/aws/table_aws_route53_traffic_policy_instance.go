package aws

import (
	"context"

	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAwsRoute53TrafficPolicyInstance(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_route53_traffic_policy_instance",
		Description: "AWS Route53 Traffic Policy Instance",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"id"}),
			Hydrate:    opengovernance.GetRoute53TrafficPolicyInstance,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"NoSuchTrafficPolicyInstance"}),
			},
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListRoute53TrafficPolicyInstance,
		},
		Columns: awsKaytuColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The DNS name for which Amazon Route 53 responds to queries.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.TrafficPolicyInstance.Name")},
			{
				Name:        "id",
				Description: "The id that Amazon Route 53 assigned to the new traffic policy instance.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.TrafficPolicyInstance.Id")},
			{
				Name:        "hosted_zone_id",
				Description: "The id of the hosted zone that Amazon Route 53 created resource record sets in.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.TrafficPolicyInstance.HostedZoneId")},
			{
				Name:        "message",
				Description: "If State is Failed, an explanation of the reason for the failure.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.TrafficPolicyInstance.Message")},
			{
				Name:        "state",
				Description: "Current state of the instance.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.TrafficPolicyInstance.State")},
			{
				Name:        "traffic_policy_id",
				Description: "The ID of the traffic policy that Amazon Route 53 used to create resource record sets in the specified hosted zone.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.TrafficPolicyInstance.TrafficPolicyId")},
			{
				Name:        "traffic_policy_type",
				Description: "The DNS type that Amazon Route 53 assigned to all of the resource record sets that it created for this traffic policy instance.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.TrafficPolicyInstance.TrafficPolicyType")},
			{
				Name:        "traffic_policy_version",
				Description: "The version of the traffic policy that Amazon Route 53 used to create resource record sets in the specified hosted zone.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.TrafficPolicyInstance.TrafficPolicyVersion")},
			{
				Name:        "ttl",
				Description: "The TTL that Amazon Route 53 assigned to all of the resource record sets that it created in the specified hosted zone.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("TTL"),
			},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.TrafficPolicyInstance.Name")},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("ARN").Transform(arnToAkas)},
		}),
	}
}
