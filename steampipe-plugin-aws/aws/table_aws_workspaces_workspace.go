package aws

import (
	"context"
	"fmt"
	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsWorkspace(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_workspaces_workspace",
		Description: "AWS Workspaces",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("workspace_id"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"ValidationException"}),
			},
			Hydrate: kaytu.GetWorkspacesWorkspace,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListWorkspacesWorkspace,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "bundle_id", Require: plugin.Optional},
				{Name: "directory_id", Require: plugin.Optional},
				{Name: "user_name", Require: plugin.Optional},
			},
		},

		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "workspace_id",
				Description: "The id of the WorkSpace.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Workspace.WorkspaceId")},
			{
				Name:        "name",
				Description: "The name of the WorkSpace.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Workspace.ComputerName")},
			{
				Name:        "arn",
				Description: "The arn of the WorkSpace.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.From(getWorkspacesWorkspaceArn),
			},
			{
				Name:        "bundle_id",
				Description: "The identifier of the bundle used to create the WorkSpace.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Workspace.BundleId")},
			{
				Name:        "directory_id",
				Description: "The identifier of the AWS Directory Service directory for the WorkSpace.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Workspace.DirectoryId")},
			{
				Name:        "state",
				Description: "The operational state of the WorkSpace.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Workspace.State")},
			{
				Name:        "error_code",
				Description: "The error code that is returned if the WorkSpace cannot be created.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Workspace.ErrorCode")},
			{
				Name:        "error_message",
				Description: "The text of the error message that is returned if the WorkSpace cannot be created.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Workspace.ErrorMessage")},
			{
				Name:        "ip_address",
				Description: "The IP address of the WorkSpace.",
				Type:        proto.ColumnType_IPADDR,
				Transform:   transform.FromField("Description.Workspace.IpAddress")},
			{
				Name:        "root_volume_encryption_enabled",
				Description: "Indicates whether the data stored on the root volume is encrypted.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Workspace.RootVolumeEncryptionEnabled")},
			{
				Name:        "subnet_id",
				Description: "The identifier of the subnet for the WorkSpace.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Workspace.SubnetId")},
			{
				Name:        "user_name",
				Description: "The user for the WorkSpace.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Workspace.UserName")},
			{
				Name:        "user_volume_encryption_enabled",
				Description: "Indicates whether the data stored on the user volume is encrypted.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Workspace.UserVolumeEncryptionEnabled")},
			{
				Name:        "volume_encryption_key",
				Description: "The symmetric AWS KMS customer master key (CMK) used to encrypt data stored on your WorkSpace. Amazon WorkSpaces does not support asymmetric CMKs.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Workspace.VolumeEncryptionKey")},
			{
				Name:        "modification_states",
				Description: "The modification states of the WorkSpace.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Workspace.ModificationStates")},
			{
				Name:        "workspace_properties",
				Description: "The properties of the WorkSpace.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Workspace.WorkspaceProperties")},
			{
				Name:        "tags_src",
				Description: "The list of tags for the WorkSpace.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Tags")},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Workspace.ComputerName")},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(workspaceTurbotTags),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(getWorkspacesWorkspaceArn).Transform(arnToAkas),
			},
		}),
	}
}

// // TRANSFORM FUNCTION
func getWorkspacesWorkspaceArn(_ context.Context, d *transform.TransformData) (interface{}, error) {
	workspace := d.HydrateItem.(kaytu.WorkspacesWorkspace).Description.Workspace
	metadata := d.HydrateItem.(kaytu.WorkspacesWorkspace).Metadata

	arn := fmt.Sprintf("arn:%s:workspaces:%s:%s:workspace/%s", metadata.Partition, metadata.Region, metadata.AccountID, *workspace.WorkspaceId)
	return arn, nil
}

func workspaceTurbotTags(_ context.Context, d *transform.TransformData) (interface{}, error) {
	tags := d.HydrateItem.(kaytu.WorkspacesWorkspace).Description.Tags

	if tags != nil {
		turbotTagsMap := map[string]string{}
		for _, i := range tags {
			turbotTagsMap[*i.Key] = *i.Value
		}
		return turbotTagsMap, nil
	}

	return nil, nil
}
