package aws

import (
	"context"
	"strings"

	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsWafv2IpSet(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_wafv2_ip_set",
		Description: "AWS WAFv2 IP Set",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"id", "name", "scope"}),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"WAFNonexistentItemException", "WAFInvalidParameterException", "InvalidParameter", "ValidationException"}),
			},
			Hydrate: opengovernance.GetWAFv2IPSet,
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListWAFv2IPSet,
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the IP set.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.IPSet.Name")},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the entity.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.IPSet.ARN")},
			{
				Name:        "id",
				Description: "A unique identifier for the IP set.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.IPSet.Id")},
			{
				Name:        "scope",
				Description: "Specifies the scope of the IP Set. Possible values are: 'REGIONAL' and 'CLOUDFRONT'.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.From(ipSetLocation),
			},
			{
				Name:        "description",
				Description: "A description of the IP set that helps with identification.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.IPSet.Description")},
			{
				Name:        "ip_address_version",
				Description: "Specifies the IP address type. Possible values are: 'IPV4' and 'IPV6'.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.IPSet.IPAddressVersion")},
			{
				Name:        "lock_token",
				Description: "A token used for optimistic locking.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.IPSetSummary.LockToken")},
			{
				Name:        "addresses",
				Description: "An array of strings that specify one or more IP addresses or blocks of IP addresses in Classless Inter-Domain Routing (CIDR) notation.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.IPSet.Addresses")},
			{
				Name:        "tags_src",
				Description: "A list of tags associated with the resource.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Tags")},

			// steampipe standard columns

			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.IPSet.Name")},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,

				Transform: transform.FromField("Description.Tags").Transform(ipSetTagListToTurbotTags),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.IPSet.ARN").Transform(arnToAkas),

				// AWS standard columns
			},

			{
				Name:        "partition",
				Description: "The AWS partition in which the resource is located (aws, aws-cn, or aws-us-gov).",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Metadata.Partition")},
			{
				Name:        "region",
				Description: "The AWS Region in which the resource is located.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.From(ipSetRegion),
			},
			{
				Name:        "account_id",
				Description: "The AWS Account ID in which the resource is located.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Metadata.AccountID")},

			//// LIST FUNCTION
		}),
	}
}

//// TRANSFORM FUNCTIONS

func ipSetLocation(_ context.Context, d *transform.TransformData) (interface{}, error) {
	data := d.HydrateItem.(opengovernance.WAFv2IPSet)
	loc := strings.Split(strings.Split(*data.Description.IPSet.ARN, ":")[5], "/")[0]
	if loc == "regional" {
		return "REGIONAL", nil
	}
	return "CLOUDFRONT", nil
}

func ipSetTagListToTurbotTags(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	data := d.HydrateItem.(opengovernance.WAFv2IPSet).Description.Tags
	// Mapping the resource tags inside turbotTags
	var turbotTagsMap map[string]string
	if data != nil {
		turbotTagsMap = map[string]string{}
		for _, i := range data {
			turbotTagsMap[*i.Key] = *i.Value
		}
	}

	return turbotTagsMap, nil
}

func ipSetRegion(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	data := d.HydrateItem.(opengovernance.WAFv2IPSet)
	loc := strings.Split(strings.Split(*data.Description.IPSet.ARN, ":")[5], "/")[0]

	if loc == "global" {
		return "global", nil
	}
	return data.Metadata.Region, nil
}
