package aws

import (
	"context"
	"fmt"

	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAwsGrafanaWorkspace(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_grafana_workspace",
		Description: "AWS Grafana Workspace",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    kaytu.GetGrafanaWorkspace,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListGrafanaWorkspace,
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "id",
				Description: "The id of the workspace.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Workspace.Id")},
			{
				Name:        "name",
				Description: "The name of the workspace.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Workspace.Name")},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the workspace",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.From(getGrafanaWorkspaceArn)},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Workspace.Name")},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Workspace.Tags"),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(getGrafanaWorkspaceArn).Transform(arnToAkas),
			},
		}),
	}
}

//// TRANSFORM FUNCTIONS

func getGrafanaWorkspaceArn(_ context.Context, d *transform.TransformData) (interface{}, error) {
	workspace := d.HydrateItem.(kaytu.GrafanaWorkspace).Description.Workspace
	metadata := d.HydrateItem.(kaytu.GrafanaWorkspace).Metadata

	arn := fmt.Sprintf("arn:%s:grafana:%s:%s:/workspaces/%s", metadata.Partition, metadata.Region, metadata.AccountID, *workspace.Id)
	return arn, nil
}
