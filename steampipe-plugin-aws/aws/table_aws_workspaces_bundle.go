package aws

import (
	"context"
	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"
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
			Hydrate:    opengovernance.GetWorkspacesBundle,
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListWorkspacesBundle,
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
				Transform:   transform.FromField("ARN")},
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
				Transform:   transform.FromField("ARN").Transform(arnToAkas),
			},
		}),
	}
}

//// TRANSFORM FUNCTIONS

func getWorkspacesBundleTurbotTags(_ context.Context, d *transform.TransformData) (interface{}, error) {
	tags := d.HydrateItem.(opengovernance.WorkspacesBundle).Description.Tags
	return workspacesV2TagsToMap(tags)
}
