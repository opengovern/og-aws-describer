package aws

import (
	"context"

	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAwsAMPWorkspace(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_amp_workspace",
		Description: "AWS AMP Workspace",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("workspace_id"),
			Hydrate:    kaytu.GetAMPWorkspace,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListAMPWorkspace,
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "workspace_id",
				Description: "The id of the workspace.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Workspace.WorkspaceId")},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the workspace",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Workspace.Arn")},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Workspace.WorkspaceId")},
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
				Transform:   transform.FromField("Description.Workspace.Arn").Transform(arnToAkas),
			},
		}),
	}
}
