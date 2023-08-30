package aws

import (
	"context"

	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsEc2ManagedPrefixList(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_ec2_managed_prefix_list",
		Description: "AWS EC2 Managed Prefix List",
		List: &plugin.ListConfig{
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"InvalidAction", "InvalidRequest"}),
			},
			Hydrate: kaytu.ListEC2ManagedPrefixList,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "name", Require: plugin.Optional},
				{Name: "id", Require: plugin.Optional},
				{Name: "owner_id", Require: plugin.Optional},
			},
		},

		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the prefix list.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ManagedPrefixList.PrefixListName")},
			{
				Name:        "id",
				Description: "The ID of the prefix list.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ManagedPrefixList.PrefixListId"),
			},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) for the prefix list.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ManagedPrefixList.PrefixListArn"),
			},
			{
				Name:        "state",
				Description: "The current state of the prefix list.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ManagedPrefixList.State")},
			{
				Name:        "address_family",
				Description: "The IP address version of the prefix list.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ManagedPrefixList.AddressFamily")},
			{
				Name:        "max_entries",
				Description: "The maximum number of entries for the prefix list.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.ManagedPrefixList.MaxEntries")},
			{
				Name:        "owner_id",
				Description: "The ID of the owner of the prefix list.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ManagedPrefixList.OwnerId")},
			{
				Name:        "state_message",
				Description: "The message regarding the current state of the prefix list.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ManagedPrefixList.StateMessage")},
			{
				Name:        "version",
				Description: "The version of the prefix list.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.ManagedPrefixList.Version")},
			{
				Name:        "tags_src",
				Description: "The tags for the prefix list.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ManagedPrefixList.Tags")},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ManagedPrefixList.PrefixListName")},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(prefixListTags),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ManagedPrefixList.PrefixListArn").Transform(transform.EnsureStringArray),
			},
		}),
	}
}

//// TRANSFORM FUNCTION

func prefixListTags(_ context.Context, d *transform.TransformData) (interface{}, error) {
	prefixList := d.HydrateItem.(kaytu.EC2ManagedPrefixList).Description.ManagedPrefixList

	var turbotTagsMap map[string]string
	if prefixList.Tags != nil {
		turbotTagsMap = map[string]string{}
		for _, i := range prefixList.Tags {
			turbotTagsMap[*i.Key] = *i.Value
		}
		return turbotTagsMap, nil
	}
	return nil, nil
}
