package aws

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/route53resolver/types"

	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsRoute53ResolverEndpoint(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_route53_resolver_endpoint",
		Description: "AWS Route53 Resolver Endpoint",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"ResourceNotFoundException"}),
			},
			Hydrate: opengovernance.GetRoute53ResolverEndpoint,
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListRoute53ResolverEndpoint,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "creator_request_id", Require: plugin.Optional},
				{Name: "direction", Require: plugin.Optional},
				{Name: "host_vpc_id", Require: plugin.Optional},
				{Name: "ip_address_count", Require: plugin.Optional},
				{Name: "status", Require: plugin.Optional},
				{Name: "name", Require: plugin.Optional},
			},
		},

		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The name that you assigned to the Resolver endpoint when you submitted a CreateResolverEndpoint.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ResolverEndpoint.Name")},
			{
				Name:        "id",
				Description: "The ID of the Resolver endpoint.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ResolverEndpoint.Id")},
			{
				Name:        "arn",
				Description: "The ARN (Amazon Resource Name) for the Resolver endpoint.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ResolverEndpoint.Arn")},
			{
				Name:        "creation_time",
				Description: "The date and time that the endpoint was created, in Unix time format and Coordinated Universal Time (UTC).",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ResolverEndpoint.CreationTime")},
			{
				Name:        "creator_request_id",
				Description: "A unique string that identifies the request that created the Resolver endpoint.The CreatorRequestId allows failed requests to be retried without the risk of executing the operation twice.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ResolverEndpoint.CreatorRequestId")},
			{
				Name:        "direction",
				Description: "Indicates whether the Resolver endpoint allows inbound or outbound DNS queries.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ResolverEndpoint.Direction")},
			{
				Name:        "host_vpc_id",
				Description: "The ID of the VPC that you want to create the Resolver endpoint in.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ResolverEndpoint.HostVPCId")},
			{
				Name:        "ip_address_count",
				Description: "The number of IP addresses that the Resolver endpoint can use for DNS queries.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.ResolverEndpoint.IpAddressCount")},
			{
				Name:        "modification_time",
				Description: "The date and time that the endpoint was last modified, in Unix time format and Coordinated Universal Time (UTC).",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ResolverEndpoint.ModificationTime")},
			{
				Name:        "status",
				Description: "A code that specifies the current status of the Resolver endpoint.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ResolverEndpoint.Status")},
			{
				Name:        "status_message",
				Description: "A detailed description of the status of the Resolver endpoint.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ResolverEndpoint.StatusMessage")},
			{
				Name:        "ip_addresses",
				Description: "Information about the IP addresses in your VPC that DNS queries originate from (for outbound endpoints) or that you forward DNS queries to (for inbound endpoints).",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.IpAddresses")},
			{
				Name:        "security_group_ids",
				Description: "The ID of one or more security groups that control access to this VPC.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ResolverEndpoint.SecurityGroupIds")},
			{
				Name:        "tags_src",
				Description: "A list of tags assigned to the Resolver endpoint.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Tags")},

			// Standard columns for all tables
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ResolverEndpoint.Name")},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Tags").Transform(route53resolverTagListToTurbotTags),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ResolverEndpoint.Arn").Transform(arnToAkas),
			},
		}),
	}
}

//// TRANSFORM FUNCTIONS

func route53resolverTagListToTurbotTags(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	tagList := d.Value.([]types.Tag)

	// Mapping the resource tags inside turbotTags
	var turbotTagsMap map[string]string
	if tagList != nil {
		turbotTagsMap = map[string]string{}
		for _, i := range tagList {
			turbotTagsMap[*i.Key] = *i.Value
		}
	} else {
		return nil, nil
	}

	return turbotTagsMap, nil
}
