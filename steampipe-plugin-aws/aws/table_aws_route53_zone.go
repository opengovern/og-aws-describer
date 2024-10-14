package aws

import (
	"context"
	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAwsRoute53Zone(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_route53_zone",
		Description: "AWS Route53 Zone",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    opengovernance.GetRoute53HostedZone,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"NoSuchHostedZone"}),
			},
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListRoute53HostedZone,
		},
		Columns: awsKaytuColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the domain. For public hosted zones, this is the name that is registered with your DNS registrar.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.HostedZone.Name")},
			{
				Name:        "id",
				Description: "The ID that Amazon Route 53 assigned to the hosted zone when it was created.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ID"),
			},
			{
				Name:        "caller_reference",
				Description: "The value that you specified for CallerReference when you created the hosted zone.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.HostedZone.CallerReference")},
			{
				Name:        "comment",
				Description: "A comment for the zone.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.HostedZone.Config.Comment")},
			{
				Name:        "private_zone",
				Description: "If true, the zone is Private hosted Zone, otherwise it is public.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.HostedZone.Config.PrivateZone")},
			{
				Name:        "linked_service_principal",
				Description: "If the health check or hosted zone was created by another service, the service that created the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.HostedZone.LinkedService.ServicePrincipal")},
			{
				Name:        "linked_service_description",
				Description: "If the health check or hosted zone was created by another service, an optional description that can be provided by the other service.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.HostedZone.LinkedService.Description")},
			{
				Name:        "resource_record_set_count",
				Description: "The number of resource record sets in the hosted zone.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.HostedZone.ResourceRecordSetCount")},
			{
				Name:        "query_logging_configs",
				Description: "A list of configuration for DNS query logging that is associated with the current AWS account.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.QueryLoggingConfigs")},
			{
				Name:        "dnssec_key_signing_keys",
				Description: "The key-signing keys (KSKs) in AWS account.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DNSSec.KeySigningKeys")},
			{
				Name:        "dnssec_status",
				Description: "The status of DNSSEC.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DNSSec.Status")},
			{
				Name:        "tags_src",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Tags")},

			{
				Name:        "vpcs",
				Description: "The list of VPCs that are authorized to be associated with the specified hosted zone.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("VPCs"),
			},
			//error
			// i did not find this field

			// Steampipe standard columns
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.HostedZone.Name")},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(route53HostedZoneTurbotTags),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("ARN").Transform(arnToAkas)},
			{
				Name:        "resource_record_set_limit",
				Description: "The maximum number of resource record sets allowed in the hosted zone.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Limit.Value"),
			},
		}),
	}
}

func route53HostedZoneTurbotTags(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	plugin.Logger(ctx).Trace("route53HostedZoneTurbotTags")
	var turbotTagsMap map[string]string
	if d.Value == nil {
		return turbotTagsMap, nil
	}

	tags := d.Value.(opengovernance.Route53HostedZone).Description.Tags

	// Mapping the resource tags inside turbotTags
	if tags != nil {
		turbotTagsMap = map[string]string{}
		for _, i := range tags {
			turbotTagsMap[*i.Key] = *i.Value
		}
	}

	return turbotTagsMap, nil
}
