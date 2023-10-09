package aws

import (
	"context"
	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAwsResourceGroupsGroup(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_resourcegroups_group",
		Description: "AWS ResourceGroups Group",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("arn"), //TODO: change this to the primary key columns in model.go
			Hydrate:    kaytu.GetResourceGroupsGroup,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListResourceGroupsGroup,
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "arn",
				Description: "The id of the group.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.GroupIdentifier.GroupArn")},
			{
				Name:        "name",
				Description: "The name of the group.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.GroupIdentifier.GroupName")},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.GroupIdentifier.GroupName")},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(getresourceGroupGroupTag), // probably needs a transform function
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.GroupIdentifier.GroupArn").Transform(arnToAkas),
			},
		}),
	}
}

func getresourceGroupGroupTag(_ context.Context, d *transform.TransformData) (interface{}, error) {
	snapshot := d.HydrateItem.(kaytu.ResourceGroupsGroup).Description.Tags

	// Get the resource tags
	if snapshot.Tags != nil {
		turbotTagsMap := map[string]string{}
		for key, value := range snapshot.Tags {
			turbotTagsMap[key] = value
		}
		return turbotTagsMap, nil
	}
	return nil, nil
}
