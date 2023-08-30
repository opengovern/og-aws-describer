package aws

import (
	"context"
	"fmt"

	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAwsWorkspacesBundle(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_workspaces_bundle",
		Description: "AWS Workspaces Bundle",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("bundle_id"),
			Hydrate:    kaytu.GetWorkspacesBundle,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListWorkspacesBundle,
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "bundle_id",
				Description: "The id of the bundle.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Bundle.BundleId")},
			{
				Name:        "name",
				Description: "The name of the bundle.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Bundle.Name")},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the bundle",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.From(getWorkspacesBundleArn)},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Bundle.Name")},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(getWorkspacesBundleTurbotTags),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(getWorkspacesBundleArn).Transform(arnToAkas),
			},
		}),
	}
}

//// TRANSFORM FUNCTIONS

func getWorkspacesBundleTurbotTags(_ context.Context, d *transform.TransformData) (interface{}, error) {
	tags := d.HydrateItem.(kaytu.WorkspacesBundle).Description.Tags
	return workspacesV2TagsToMap(tags)
}

func getWorkspacesBundleArn(_ context.Context, d *transform.TransformData) (interface{}, error) {
	bundle := d.HydrateItem.(kaytu.WorkspacesBundle).Description.Bundle
	metadata := d.HydrateItem.(kaytu.WorkspacesBundle).Metadata

	arn := fmt.Sprintf("arn:%s:workspaces:%s:%s:workspacebundle/%s", metadata.Partition, metadata.Region, metadata.AccountID, *bundle.BundleId)
	return arn, nil
}
